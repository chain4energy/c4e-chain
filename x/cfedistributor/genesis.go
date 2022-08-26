package cfedistributor

import (
	"github.com/chain4energy/c4e-chain/x/cfedistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState, ak types.AccountKeeper) {
	k.SetParams(ctx, genState.Params)
	for _, sb := range k.GetParams(ctx).SubDistributors {
		k.Logger(ctx).Info("SubDistributor: " + sb.Name)
	}
	k.Logger(ctx).Info("Init Genesis module: " + types.ModuleName)

	states := genState.States
	for _, av := range states {
		k.SetState(ctx, *av)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	states := k.GetALlStates(ctx)

	for i := 0; i < len(states); i++ {
		genesis.States = append(genesis.States, &states[i])
	}
	return genesis
}
