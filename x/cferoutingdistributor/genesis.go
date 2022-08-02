package cferoutingdistributor

import (
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState, ak types.AccountKeeper) {
	k.SetParams(ctx, genState.Params)

	k.Logger(ctx).Info("Init Genesis module: " + types.ModuleName)
	for _, account := range genState.Params.RoutingDistributor.ModuleAccounts {
		k.Logger(ctx).Info("Load module account name: " + account)
		ak.GetModuleAccount(ctx, account)
	}

	allRemains := genState.RemainsList.Remains

	for _, av := range allRemains {
		k.SetRemains(ctx, *av)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	allRemains := k.GetAllRemains(ctx)

	for i := 0; i < len(allRemains); i++ {
		genesis.RemainsList.Remains = append(genesis.RemainsList.Remains, &allRemains[i])
	}

	return genesis
}
