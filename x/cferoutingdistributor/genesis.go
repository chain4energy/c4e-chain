package cferoutingdistributor

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState, ak types.AccountKeeper) {
	k.SetParams(ctx, genState.Params)

	fmt.Printf("%+v\n", genState.RoutingDistributor) // Print with Variable Name

	k.SetRoutingDistributor(ctx, genState.RoutingDistributor)

	k.Logger(ctx).Info("Init Genesis module: " + types.ModuleName)
	for _, account := range genState.RoutingDistributor.ModuleAccounts {
		k.Logger(ctx).Info("Load module account name: " + account)
		ak.GetModuleAccount(ctx, account)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.RoutingDistributor = k.GetRoutingDistributorr(ctx)
	return genesis
}
