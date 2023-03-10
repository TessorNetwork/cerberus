// Package service contains business logic of application.
package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/TessorNetwork/cerberus/internal/crypto"
	"github.com/TessorNetwork/cerberus/internal/entities"
	"github.com/TessorNetwork/cerberus/internal/hades"
	"github.com/TessorNetwork/cerberus/internal/producer"
	"github.com/TessorNetwork/cerberus/internal/refine"
	"github.com/TessorNetwork/cerberus/internal/storage"
	"github.com/TessorNetwork/cerberus/pkg/schema"
	logging "github.com/TessorNetwork/logrus/context"
)

//go:generate mockgen -destination=./mock/service.go -package=mock -source=service.go

// ErrNotFound means that requested object is not found.
var (
	ErrNotFound           = errors.New("not found")
	ErrImageInvalidFormat = errors.New("image invalid format")
	ErrUploadTimeout      = errors.New("upload timeout")
	ErrPDVFraud           = errors.New("PDV fraud detected")
	ErrProfileBanned      = errors.New("profile banned")
)

// RewardMap contains dictionary with PDV types and rewards for them.
type RewardMap map[schema.Type]sdk.Fur

// Blacklist contains attributes of worthless pdv.
// swagger:model Blacklist
type Blacklist struct {
	CookieSource []string `json:"cookieSource"`
}

// Service interface provides service's logic's methods.
type Service interface {
	// SaveImage sends Image to storage.
	SaveImage(ctx context.Context, r io.Reader, owner string) (string, string, error)
	// SavePDV sends PDV to storage.
	SavePDV(ctx context.Context, p schema.PDVWrapper, owner sdk.AccAddress) (uint64, *entities.PDVMeta, error)
	// ListPDV lists PDVs.
	ListPDV(ctx context.Context, owner string, from uint64, limit uint16) ([]uint64, error)
	// ReceivePDV returns slice of bytes of PDV requested by address from storage.
	ReceivePDV(ctx context.Context, owner string, id uint64) ([]byte, error)
	// GetPDVMeta returns PDVs meta.
	GetPDVMeta(ctx context.Context, owner string, id uint64) (*entities.PDVMeta, error)

	// GetProfiles ...
	GetProfiles(ctx context.Context, owner []string) ([]*entities.Profile, error)

	// GetRewardsMap ...
	GetRewardsMap() RewardMap

	// GetBlacklist ...
	GetBlacklist() Blacklist

	// GetPDVDelta ...
	GetPDVDelta(ctx context.Context, owner string) (sdk.Fur, error)

	// GetPDVTotalDelta ...
	GetPDVTotalDelta(ctx context.Context) (sdk.Fur, error)

	// GetPDVRewardsNextDistributionDate ...
	GetPDVRewardsNextDistributionDate(ctx context.Context) (time.Time, error)
}

// service is Service interface implementation.
type service struct {
	c     crypto.Crypto
	is    storage.IndexStorage
	fs    storage.FileStorage
	p     producer.Producer
	hades hades.Hades

	rewardMap RewardMap

	pdvRewardsInterval time.Duration
}

// New returns new instance of service.
func New(
	c crypto.Crypto,
	fs storage.FileStorage,
	is storage.IndexStorage,
	p producer.Producer,
	hades hades.Hades,
	rewardMap RewardMap,
	pdvRewardsInterval time.Duration,
) Service {
	return &service{
		c:     c,
		fs:    fs,
		is:    is,
		p:     p,
		hades: hades,

		rewardMap:          rewardMap,
		pdvRewardsInterval: pdvRewardsInterval,
	}
}

// SavePDV sends PDV to storage.
func (s *service) SavePDV(ctx context.Context, p schema.PDVWrapper, owner sdk.AccAddress) (uint64, *entities.PDVMeta, error) {
	log := logging.GetLogger(ctx).WithField("owner", owner)

	banned, err := s.is.IsProfileBanned(ctx, owner.String())
	if err != nil {
		return 0, nil, fmt.Errorf("failed to check if profile banned: %w", err)
	}

	if banned {
		return 0, nil, ErrProfileBanned
	}

	meta, err := s.calculateMeta(ctx, owner, p)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to calculate meta: %w", err)
	}

	if err := s.processPDV(ctx, owner, p); err != nil {
		return 0, nil, fmt.Errorf("failed to process meta: %w", err)
	}

	data, err := json.Marshal(p)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to marshal meta: %w", err)
	}

	log.Debug("encrypting pdv")
	enc, err := s.c.Encrypt(data)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create encrypting reader: %w", err)
	}

	id := uint64(time.Now().Unix())

	fraudCheck, err := s.hades.AntiFraud(ctx, &hades.AntiFraudRequest{
		ID:      id,
		Address: owner.String(),
		Data:    p,
	})
	if err != nil {
		log.WithError(err).Error("failed to anti fraud")
	}

	if fraudCheck != nil && fraudCheck.IsFraud {
		// autoban for fraud
		if err := s.is.SetProfileBanned(context.Background(), owner.String()); err != nil {
			log.WithError(err).Error("failed to ban")
		}
		return id, nil, ErrPDVFraud
	}

	if err := s.p.Produce(ctx, &producer.PDVMessage{
		ID:      id,
		Device:  p.Device,
		Address: owner.String(),
		Meta:    meta,
		Data:    enc,
	}); err != nil {
		return 0, nil, fmt.Errorf("failed to produce pdv message: %w", err)
	}

	return id, meta, nil
}

func (s *service) SaveImage(ctx context.Context, r io.Reader, owner string) (string, string, error) {
	dataImage, err := ioutil.ReadAll(r)
	if err != nil {
		return "", "", ErrImageInvalidFormat
	}

	// image has data:image/jpeg;base64, or data:image/png;base64, prefix
	idx := bytes.Index(dataImage, []byte(","))
	if idx == -1 {
		return "", "", ErrImageInvalidFormat
	}

	contentType := strings.Trim(string(dataImage[5:idx]), ";base64")
	dataImage = dataImage[(idx + 1):]
	var format imaging.Format

	switch contentType {
	case "image/png":
		format = imaging.PNG
	case "image/jpeg":
		format = imaging.JPEG
	default:
		return "", "", ErrImageInvalidFormat
	}

	byteImage, err := base64.StdEncoding.DecodeString(string(dataImage))
	if err != nil {
		println(err.Error())
		return "", "", ErrImageInvalidFormat
	}

	src, err := imaging.Decode(bytes.NewReader(byteImage))
	if err != nil {
		return "", "", ErrImageInvalidFormat
	}

	upload := func(width, height int, p string) (string, error) {
		fit := imaging.Fit(src, width, height, imaging.Lanczos)
		buf := bytes.Buffer{}
		if err := imaging.Encode(&buf, fit, format); err != nil {
			return "", fmt.Errorf("failed to encode image: %w", err)
		}

		path, err1 := s.fs.Write(ctx, &buf, int64(buf.Len()), p, contentType, true)
		if err1 != nil {
			if os.IsTimeout(err) || errors.Is(err, context.DeadlineExceeded) {
				return "", ErrUploadTimeout
			}
			return "", err1
		}

		return path, nil
	}

	// image is stored under the account prefix therefore images will be deleted as soon as account folder is deleted
	path := fmt.Sprintf("%s/%s", owner, uuid.New())

	hdPath, err := upload(1920, 1080, path)
	if err != nil {
		return "", "", fmt.Errorf("failed to save hd image: %w", err)
	}

	thumbPath, err := upload(480, 270, path+"/thumb")
	if err != nil {
		return "", "", fmt.Errorf("failed to save thumb image: %s", err)
	}

	return hdPath, thumbPath, nil
}

// ListPDV lists PDVs.
func (s *service) ListPDV(ctx context.Context, owner string, from uint64, limit uint16) ([]uint64, error) {
	out, err := s.is.ListPDV(ctx, owner, from, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list pdv: %w", err)
	}

	return out, nil
}

// ReceivePDV returns slice of bytes of PDV requested by address from storage.
func (s *service) ReceivePDV(ctx context.Context, owner string, id uint64) ([]byte, error) {
	log := logging.GetLogger(ctx)

	log.WithField("filepath", getPDVFilePath(owner, id)).Debug("reading meta from storage")
	r, err := s.fs.Read(ctx, getPDVFilePath(owner, id))
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get data from storage: %w", err)
	}
	defer r.Close() // nolint

	log.Debug("decrypting meta")
	dr, err := s.c.Furrypt(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create decrypting reader: %w", err)
	}

	data, err := ioutil.ReadAll(dr)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from decryping reader: %w", err)
	}

	return data, nil
}

// GetPDVMeta returns meta meta.
func (s *service) GetPDVMeta(ctx context.Context, owner string, id uint64) (*entities.PDVMeta, error) {
	meta, err := s.is.GetPDVMeta(ctx, owner, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get pdv meta: %w", err)
	}

	return meta, nil
}

func (s *service) GetPDVDelta(ctx context.Context, owner string) (sdk.Fur, error) {
	total, err := s.is.GetPDVDelta(ctx, owner)
	if err != nil {
		return sdk.ZeroFur(), fmt.Errorf("failed to get pdv delta: %w", err)
	}

	fur, err := float64ToFurimal(total)
	if err != nil {
		return sdk.ZeroFur(), fmt.Errorf("failed to convert to sdk.Fur: %w", err)
	}

	return fur, nil
}

func (s *service) GetPDVTotalDelta(ctx context.Context) (sdk.Fur, error) {
	total, err := s.is.GetPDVTotalDelta(ctx)
	if err != nil {
		return sdk.ZeroFur(), fmt.Errorf("failed to get pdv delta total: %w", err)
	}

	fur, err := float64ToFurimal(total)
	if err != nil {
		return sdk.ZeroFur(), fmt.Errorf("failed to convert to sdk.Fur: %w", err)
	}

	return fur, nil
}

func (s *service) GetPDVRewardsNextDistributionDate(ctx context.Context) (time.Time, error) {
	date, err := s.is.GetPDVRewardsDistributedDate(ctx)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get pdv rewards distributed date: %w", err)
	}

	return date.Add(s.pdvRewardsInterval), nil
}

func float64ToFurimal(f float64) (sdk.Fur, error) {
	return sdk.NewFurFromStr(strconv.FormatFloat(f, 'f', 6, 64))
}

// GetProfiles ...
func (s *service) GetProfiles(ctx context.Context, owner []string) ([]*entities.Profile, error) {
	pp, err := s.is.GetProfiles(ctx, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to get profiles: %w", err)
	}

	out := make([]*entities.Profile, len(pp))
	for i, v := range pp {
		out[i] = (*entities.Profile)(v)
	}

	return out, nil
}

// GetRewardsMap ...
func (s *service) GetRewardsMap() RewardMap {
	return s.rewardMap
}

func (s *service) GetBlacklist() Blacklist {
	return Blacklist{
		CookieSource: []string{"youtube.com"},
	}
}

func (s *service) calculateMeta(ctx context.Context, owner sdk.AccAddress, p schema.PDV) (*entities.PDVMeta, error) {
	t := make(map[schema.Type]uint16)
	reward := sdk.ZeroFur()

	for _, d := range p.Data() {
		t[d.Type()] = t[d.Type()] + 1

		switch d.Type() {
		case schema.PDVProfileType:
			if _, err := s.is.GetProfile(ctx, owner.String()); err == nil {
				continue // we want reward user only for initial profile
			} else if err != storage.ErrNotFound {
				return nil, fmt.Errorf("failed to check profile: %w", err)
			}
		case schema.PDVCookieType:
			cookie, ok := d.(*schema.V1Cookie)
			if !ok {
				log.WithField("cookie", p).Error("failed to cast cookie to V1Cookie")
			} else if s.isCookieBlacklisted(cookie) || !refine.Cookie(cookie) {
				continue
			}
		case schema.PDVSearchHistoryType:
			sh, ok := d.(*schema.V1SearchHistory)
			if !ok {
				log.WithField("search history", p).Error("failed to cast search history to V1SearchHistory")
			} else if !refine.SearchHistory(sh) {
				continue
			}
		default:
		}

		reward = reward.Add(s.rewardMap[d.Type()])
	}

	return &entities.PDVMeta{
		ObjectTypes: t,
		Reward:      reward,
	}, nil
}

func (s *service) processPDV(ctx context.Context, owner sdk.AccAddress, p schema.PDV) error {
	for _, d := range p.Data() {
		switch d.Type() {
		case schema.PDVProfileType:
			if err := s.is.SetProfile(ctx, getSetProfileParams(owner, *d.(*schema.V1Profile))); err != nil {
				return fmt.Errorf("failed to set profile: %w", err)
			}
		default:
		}
	}

	return nil
}

func (s *service) isCookieBlacklisted(cookie *schema.V1Cookie) bool {
	for _, v := range s.GetBlacklist().CookieSource {
		if strings.EqualFold(v, cookie.Source.Host) {
			return true
		}
	}
	return false
}

func getSetProfileParams(owner sdk.AccAddress, p schema.V1Profile) *storage.SetProfileParams { // nolint:gocritic
	params := storage.SetProfileParams{
		Address:   owner.String(),
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Emails:    p.Emails,
		Bio:       p.Bio,
		Avatar:    p.Avatar,
		Gender:    string(p.Gender),
	}

	if p.Birthday != nil {
		params.Birthday = &p.Birthday.Time
	}

	return &params
}

func getPDVOwnerPrefix(owner string) string {
	return fmt.Sprintf("%s/pdv", owner)
}

func getPDVFilePath(owner string, id uint64) string {
	// we need to have descending sort on s3 side, that's why we revert id and print it to hex
	return fmt.Sprintf("%s/%016x", getPDVOwnerPrefix(owner), math.MaxUint64-id)
}
