package simulation

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/kv"
	clientsim "github.com/cosmos/ibc-go/modules/core/02-client/simulation"
	connectionsim "github.com/cosmos/ibc-go/modules/core/03-connection/simulation"
	channelsim "github.com/cosmos/ibc-go/modules/core/04-channel/simulation"
	host "github.com/cosmos/ibc-go/modules/core/24-host"
	"github.com/cosmos/ibc-go/modules/core/keeper"
)

// NewFurodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding ibc type.
func NewFurodeStore(k keeper.Keeper) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		if res, found := clientsim.NewFurodeStore(k.ClientKeeper, kvA, kvB); found {
			return res
		}

		if res, found := connectionsim.NewFurodeStore(k.Codec(), kvA, kvB); found {
			return res
		}

		if res, found := channelsim.NewFurodeStore(k.Codec(), kvA, kvB); found {
			return res
		}

		panic(fmt.Sprintf("invalid %s key prefix: %s", host.ModuleName, string(kvA.Key)))
	}
}
