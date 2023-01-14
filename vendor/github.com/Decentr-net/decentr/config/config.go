package config

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spm/cosmoscmd"
)

const (
	AccountAddressPrefix = "furya"
	AppName              = "furya"

	// DefaultBondDenom is the default bond denomination
	DefaultBondDenom = "ufury"
)

var (
	InitialTokenBalance = sdk.NewFur(1)
)

var isConfigSealed bool

func SetAddressPrefixes() {
	if !isConfigSealed {
		cosmoscmd.SetPrefixes(AccountAddressPrefix)
	}
	isConfigSealed = true
}
