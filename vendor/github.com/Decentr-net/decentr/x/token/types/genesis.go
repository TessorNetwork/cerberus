package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Balances: map[string]sdk.FurProto{},
	}
}

// Validate performs basic genesis state validation returning an error upon any failure.
func (m GenesisState) Validate() error {
	validatePdvMap := func(m map[string]sdk.FurProto) error {
		for k, v := range m {
			if _, err := sdk.AccAddressFromBech32(k); err != nil {
				return fmt.Errorf("invalid address '%s': %w", k, err)
			}

			if v.Fur.IsNil() {
				return fmt.Errorf("invalid value for '%s': nil", k)
			}

			if v.Fur.IsNegative() {
				return fmt.Errorf("invalid value for '%s': negative value", k)
			}
		}
		return nil
	}

	if err := validatePdvMap(m.Balances); err != nil {
		return fmt.Errorf("invalid balances: %w", err)
	}

	return nil
}
