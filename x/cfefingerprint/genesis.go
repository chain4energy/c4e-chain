package cfefingerprint

import (
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/keeper"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, payloadLink := range genState.PayloadLinks {
		k.Logger(ctx).Debug("set payload link", "payloadLink", payloadLink)
		k.AppendPayloadLink(ctx, payloadLink.ReferenceKey, payloadLink.ReferenceValue)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.PayloadLinks = k.GetAllPayloadLinks(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
