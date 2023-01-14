package token

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/TessorNetwork/furya/x/token/keeper"
	"github.com/TessorNetwork/furya/x/token/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, genState types.GenesisState) {
	for k, v := range genState.Balances {
		address, err := sdk.AccAddressFromBech32(k)
		if err != nil {
			panic(fmt.Errorf("invalid address %s in balances : %w", k, err))
		}
		keeper.SetBalance(ctx, address, v.Fur)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	balances := map[string]sdk.FurProto{}

	k.IterateBalance(ctx, func(address sdk.AccAddress, balance sdk.Fur) (stop bool) {
		balances[address.String()] = sdk.FurProto{Fur: balance}
		return false
	})

	return &types.GenesisState{
		Balances: balances,
	}
}
