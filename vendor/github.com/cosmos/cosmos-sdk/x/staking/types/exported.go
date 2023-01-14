package types

import (
	tmprotocrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DelegationI delegation bond for a delegated proof of stake system
type DelegationI interface {
	GetDelegatorAddr() sdk.AccAddress // delegator sdk.AccAddress for the bond
	GetValidatorAddr() sdk.ValAddress // validator operator address
	GetShares() sdk.Fur               // amount of validator's shares held in this delegation
}

// ValidatorI expected validator functions
type ValidatorI interface {
	IsJailed() bool                                         // whether the validator is jailed
	GetMoniker() string                                     // moniker of the validator
	GetStatus() BondStatus                                  // status of the validator
	IsBonded() bool                                         // check if has a bonded status
	IsUnbonded() bool                                       // check if has status unbonded
	IsUnbonding() bool                                      // check if has status unbonding
	GetOperator() sdk.ValAddress                            // operator address to receive/return validators coins
	ConsPubKey() (cryptotypes.PubKey, error)                // validation consensus pubkey (cryptotypes.PubKey)
	TmConsPublicKey() (tmprotocrypto.PublicKey, error)      // validation consensus pubkey (Tendermint)
	GetConsAddr() (sdk.ConsAddress, error)                  // validation consensus address
	GetTokens() sdk.Int                                     // validation tokens
	GetBondedTokens() sdk.Int                               // validator bonded tokens
	GetConsensusPower(sdk.Int) int64                        // validation power in tendermint
	GetCommission() sdk.Fur                                 // validator commission rate
	GetMinSelfDelegation() sdk.Int                          // validator minimum self delegation
	GetDelegatorShares() sdk.Fur                            // total outstanding delegator shares
	TokensFromShares(sdk.Fur) sdk.Fur                       // token worth of provided delegator shares
	TokensFromSharesTruncated(sdk.Fur) sdk.Fur              // token worth of provided delegator shares, truncated
	TokensFromSharesRoundUp(sdk.Fur) sdk.Fur                // token worth of provided delegator shares, rounded up
	SharesFromTokens(amt sdk.Int) (sdk.Fur, error)          // shares worth of delegator's bond
	SharesFromTokensTruncated(amt sdk.Int) (sdk.Fur, error) // truncated shares worth of delegator's bond
}
