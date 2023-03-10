package main

import (
	cliflags "github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sirupsen/logrus"

	"github.com/TessorNetwork/go-broadcaster"
)

type BlockchainOpts struct {
	BlockchainNode               string `long:"blockchain.node" env:"BLOCKCHAIN_NODE" default:"http://zeus.testnet.furya.xyz:26657" description:"furya node address"`
	BlockchainFrom               string `long:"blockchain.from" env:"BLOCKCHAIN_FROM" description:"furya account name to send stakes" required:"true"`
	BlockchainTxMemo             string `long:"blockchain.tx_memo" env:"BLOCKCHAIN_TX_MEMO" description:"furya tx's memo'"`
	BlockchainChainID            string `long:"blockchain.chain_id" env:"BLOCKCHAIN_CHAIN_ID" default:"testnet" description:"furya chain id"`
	BlockchainClientHome         string `long:"blockchain.client_home" env:"BLOCKCHAIN_CLIENT_HOME" default:"~/.furyacli" description:"furyacli home directory"`
	BlockchainKeyringBackend     string `long:"blockchain.keyring_backend" env:"BLOCKCHAIN_KEYRING_BACKEND" default:"test" description:"furyacli keyring backend"`
	BlockchainKeyringPromptInput string `long:"blockchain.keyring_prompt_input" env:"BLOCKCHAIN_KEYRING_PROMPT_INPUT" description:"furyacli keyring prompt input"`
	BlockchainGas                uint64 `long:"blockchain.gas" env:"BLOCKCHAIN_GAS" default:"10" description:"gas amount"`
	BlockchainFee                string `long:"blockchain.fee" env:"BLOCKCHAIN_FEE" default:"1ufury" description:"transaction fee"`
}

func mustGetBroadcaster() broadcaster.Broadcaster {
	fee, err := sdk.ParseCoinNormalized(opts.BlockchainFee)
	if err != nil {
		logrus.WithError(err).Error("failed to parse fee")
	}

	b, err := broadcaster.New(broadcaster.Config{
		KeyringRootDir:     opts.BlockchainClientHome,
		KeyringBackend:     opts.BlockchainKeyringBackend,
		KeyringPromptInput: opts.BlockchainKeyringPromptInput,
		NodeURI:            opts.BlockchainNode,
		BroadcastMode:      cliflags.BroadcastSync,
		From:               opts.BlockchainFrom,
		ChainID:            opts.BlockchainChainID,
		Gas:                opts.BlockchainGas,
		GasAdjust:          1.2,
		Fees:               sdk.Coins{fee},
	})

	if err != nil {
		logrus.WithError(err).Fatal("failed to create broadcaster")
	}

	return b
}
