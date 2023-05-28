package cfefingerprint

import (
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/keeper"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.Logger(ctx).Debug("init genesis", "genState", genState)
	k.SetParams(ctx, genState.Params)

	links := genState.PayloadLinks
	for _, av := range links {
		k.AppendPayloadLink(ctx, av.ReferenceKey, av.ReferenceValue)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	links := k.GetAllPayloadLinks(ctx)

	for i := 0; i < len(links); i++ {
		genesisPayloadLink := &types.GenesisPayloadLink{ReferenceKey: links[i].ReferenceKey, ReferenceValue: links[i].ReferenceValue}
		genesis.PayloadLinks = append(genesis.PayloadLinks, genesisPayloadLink)
	}
	return genesis
}
