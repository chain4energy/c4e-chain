package cfeminter_test

import (
	"github.com/chain4energy/c4e-chain/app"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

//func TestGenesis(t *testing.T) {
//	genesisState := types.GenesisState{
//		Params: types.DefaultParams(),
//
//		// this line is used by starport scaffolding # genesis/test/state
//	}
//
//	k, ctx := keepertest.CfeminterKeeper(t)
//	cfeminter.InitGenesis(ctx, *k, genesisState)
//	got := cfeminter.ExportGenesis(ctx, *k)
//	require.NotNil(t, got)
//
//	nullify.Fill(&genesisState)
//	nullify.Fill(got)
//
//	// this line is used by starport scaffolding # genesis/test/assert
//}

func TestGenesis2(t *testing.T) {

	// Setup main app
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	//Setup minter params
	minterNew := app.CfeminterKeeper.GetHalvingMinter(ctx)
	minterNew.MintDenom = "uC4E"
	minterNew.NewCoinsMint = 20596877
	minterNew.BlocksPerYear = 4855105
	app.CfeminterKeeper.SetHalvingMinter(ctx, minterNew)

	for i := 1; i < 1; i++ {
		ctx = ctx.WithBlockHeight(int64(i))
		app.BeginBlocker(ctx, abci.RequestBeginBlock{})
		app.EndBlocker(ctx, abci.RequestEndBlock{})
	}

	//app.BankKeeper.
	require.Equal(t, app.BankKeeper.GetSupply(ctx, "uC4E").Amount.String(), 20596877, "asd")
}
