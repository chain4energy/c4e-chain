package cfedistributor

import (
	"github.com/chain4energy/c4e-chain/x/cfedistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState, ak types.AccountKeeper) {
	k.Logger(ctx).Debug("init genesis", "genState", genState)
	k.SetParams(ctx, genState.Params)
	states := genState.States
	for _, av := range states {
		k.SetState(ctx, *av)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	states := k.GetAllStates(ctx)
	for i, state := range states {
		if state.Burn {
			states[i].Account = nil
		}
	}

	for i := 0; i < len(states); i++ {
		genesis.States = append(genesis.States, &states[i])
	}
	return genesis
}
