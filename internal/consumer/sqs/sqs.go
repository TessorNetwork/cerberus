// Package sqs is an aws sqs implementation of consumer
package sqs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/TessorNetwork/cerberus/internal/blockchain"
	"github.com/TessorNetwork/cerberus/internal/consumer"
	"github.com/TessorNetwork/cerberus/internal/producer"
	"github.com/TessorNetwork/cerberus/internal/storage"
)

var _ consumer.Consumer = &impl{}

var log = logrus.WithField("package", "sqs")

const (
	// how long the message is locked from other consumers in seconds
	visibilityTimeout int64 = 600
	// how long consumer will wait for the next messages in seconds
	waitTimeSeconds int64 = 20
	// size of bulk
	maxNumberOfMessages int64 = 10
)

type impl struct {
	fs storage.FileStorage
	is storage.IndexStorage
	b  blockchain.Blockchain

	sqs      *sqs.SQS
	queueURL string

	msgsCh chan *sqs.Message
}

// New return new instance of impl.
func New(fs storage.FileStorage,
	is storage.IndexStorage,
	b blockchain.Blockchain,
	sqsClient *sqs.SQS,
	queueURL string,
) *impl { // nolint:golint
	return &impl{
		fs:       fs,
		is:       is,
		b:        b,
		sqs:      sqsClient,
		queueURL: queueURL,

		msgsCh: make(chan *sqs.Message, 200),
	}
}

// Run consumes messages from SQS and sends pdv to users.
func (i *impl) Run(ctx context.Context) error {
	wg, ctx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		return i.runConsumer(ctx)
	})
	wg.Go(func() error {
		return i.runProcessor(ctx)
	})

	return wg.Wait()
}

// runConsumer consumes messages from SQS, and put it into the channel.
func (i *impl) runConsumer(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		out, err := i.sqs.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
			MaxNumberOfMessages: aws.Int64(maxNumberOfMessages),
			QueueUrl:            aws.String(i.queueURL),
			VisibilityTimeout:   aws.Int64(visibilityTimeout),
			WaitTimeSeconds:     aws.Int64(waitTimeSeconds),
		})
		if err != nil {
			log.WithError(err).Error("failed to receive messages")
			continue
		}

		for _, v := range out.Messages {
			i.msgsCh <- v
		}
	}
}

// runProcessor consumes messages from the channel, and process it.
func (i *impl) runProcessor(ctx context.Context) error {
	for {
		var messages []*sqs.Message

	collectMessages:
		for {
			select {
			case m := <-i.msgsCh:
				messages = append(messages, m)
			default:
				break collectMessages
			}
		}

		if len(messages) > 0 {
			log.WithField("msgs", len(messages)).Info("start processing messages")

			if err := i.processMessages(messages); err != nil {
				log.WithError(err).Error("failed to process messages")
			}
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(time.Second)
		}
	}
}

func (i *impl) processMessages(msgs []*sqs.Message) error {
	// Background context is used to gracefully shutdown processor
	ctx := context.Background()

	var (
		toDelete []*sqs.DeleteMessageBatchRequestEntry
		toReward []producer.PDVMessage

		mu sync.Mutex
	)

	parallel(8, func(m *sqs.Message) {
		var pdv producer.PDVMessage
		if err := json.Unmarshal([]byte(*m.Body), &pdv); err != nil {
			log.WithError(err).Error("failed to unmarshal message")
			return
		}

		savePDV, deleteMsg := i.storePDV(ctx, i.is, &pdv)
		mu.Lock()
		if savePDV && pdv.Meta.Reward.IsPositive() {
			toReward = append(toReward, pdv)
		}

		if deleteMsg {
			toDelete = append(toDelete, &sqs.DeleteMessageBatchRequestEntry{
				Id:            m.MessageId,
				ReceiptHandle: m.ReceiptHandle,
			})
		}
		mu.Unlock()
	}, msgs)

	if err := i.is.InTx(ctx, func(s storage.IndexStorage) error {
		if len(toReward) > 0 {
			rr := make([]blockchain.Reward, 0, len(toReward))
			for _, v := range toReward { // nolint:gocritic
				rr = append(rr, blockchain.Reward{
					Receiver: v.Address,
					ID:       v.ID,
					Reward:   v.Meta.Reward,
				})
			}

			tx, err := i.b.DistributeRewards(rr)
			if err != nil {
				return fmt.Errorf("failed to broadcast MsgDistributeRewards: %w", err)
			}

			for _, v := range toReward { // nolint:gocritic
				if err := i.is.SetPDVMeta(ctx, v.Address, v.ID, tx, v.Device, v.Meta); err != nil {
					return fmt.Errorf("failed to set meta in pg: %w", err)
				}
			}
		}

		return nil
	}); err != nil {
		return fmt.Errorf("failed to process messages bulk: %w", err)
	}

	for len(toDelete) > 0 {
		// sqs allows to delete 10 or less messages
		splitPos := len(toDelete)
		if splitPos > 9 {
			splitPos = 9
		}

		s := toDelete[0:splitPos]
		toDelete = toDelete[splitPos:]

		if _, err := i.sqs.DeleteMessageBatch(&sqs.DeleteMessageBatchInput{
			Entries:  s,
			QueueUrl: &i.queueURL,
		}); err != nil {
			return fmt.Errorf("failed to delete messages from sqs: %w", err)
		}
	}

	return nil
}

func (i *impl) storePDV(ctx context.Context, s storage.IndexStorage, pdv *producer.PDVMessage) (reward bool, deleteMsg bool) {
	log := log.WithFields(logrus.Fields{
		"id":   pdv.ID,
		"meta": pdv.Meta,
	})

	if _, err := s.GetPDVMeta(ctx, pdv.Address, pdv.ID); err == nil {
		return false, true
	} else if !errors.Is(err, storage.ErrNotFound) {
		log.WithError(err).Error("failed to check pdv existence")
		return false, false
	}

	if _, err := i.fs.Write(
		ctx,
		bytes.NewReader(pdv.Data),
		int64(len(pdv.Data)),
		getPDVFilePath(pdv.Address, pdv.ID),
		"binary/octet-stream",
		false,
	); err != nil {
		log.WithError(err).Error("failed to write data to storage")
		return false, false
	}

	return true, true
}

func parallel(routines int, f func(m *sqs.Message), batch []*sqs.Message) {
	var wg sync.WaitGroup

	ch := make(chan *sqs.Message)

	for i := 0; i < routines; i++ {
		wg.Add(1)

		go func() {
			for m := range ch {
				f(m)
			}
			wg.Done()
		}()
	}

	for _, v := range batch {
		ch <- v
	}
	close(ch)

	wg.Wait()
}

func getPDVOwnerPrefix(owner string) string {
	return fmt.Sprintf("%s/pdv", owner)
}

func getPDVFilePath(owner string, id uint64) string {
	// once we needed to have descending sort on s3 side, that's why we revert id and print it to hex
	// now we have to support this or do a bit complicated migration
	return fmt.Sprintf("%s/%016x", getPDVOwnerPrefix(owner), math.MaxUint64-id)
}
