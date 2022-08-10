package energybank

import (
	"github.com/chain4energy/c4e-chain/x/energybank/keeper"
	"github.com/chain4energy/c4e-chain/x/energybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the energyToken
	for _, elem := range genState.EnergyTokenList {
		k.SetEnergyToken(ctx, elem)
	}

	// Set energyToken count
	k.SetEnergyTokenCount(ctx, genState.EnergyTokenCount)
	// Set all the tokenParams
	for _, elem := range genState.TokenParamsList {
		k.SetTokenParams(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.EnergyTokenList = k.GetAllEnergyToken(ctx)
	genesis.EnergyTokenCount = k.GetEnergyTokenCount(ctx)
	genesis.TokenParamsList = k.GetAllTokenParams(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
