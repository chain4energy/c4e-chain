package cfeairdrop

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the claimRecordXX
	for _, elem := range genState.ClaimRecords {
		k.SetClaimRecord(ctx, elem)
	}
	// Set all the initialClaim
	for _, elem := range genState.InitialClaims {
		k.SetInitialClaim(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	k.CreateAirdropAccount(ctx, "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
		"c4e1497khscv809hzwflml2wztryvsvrftu38ea43p",
		sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(1000000))), ctx.BlockTime().Unix(), ctx.BlockTime().Unix()+100000)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ClaimRecords = k.GetAllClaimRecord(ctx)
	genesis.InitialClaims = k.GetAllInitialClaim(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
