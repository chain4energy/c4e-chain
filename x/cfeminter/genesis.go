package cfeminter

import (
	"github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, ak types.AccountKeeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.SetMinterState(ctx, genState.MinterState)
	if genState.StateHistory != nil {
		for _, hist := range genState.StateHistory {
			k.SetMinterStateHistory(ctx, *hist)
		}
	}
	ak.GetModuleAccount(ctx, types.ModuleName)
	ak.GetModuleAccount(ctx, k.GetCollectorName())
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.MinterState = k.GetMinterState(ctx)
	genesis.StateHistory = k.ConvertMinterStateHistory(ctx)
	return genesis
}
