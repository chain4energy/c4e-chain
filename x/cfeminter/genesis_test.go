package cfeminter_test

import (
	testapp "github.com/chain4energy/c4e-chain/app"
	// "github.com/stretchr/testify/require"
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
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
	perms := []string{authtypes.Minter}
	testapp.AddMaccPerms("fee_collector", perms)
	testapp.AddMaccPerms("payment_collector", perms)
	// Setup main app
	app := testapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	//Setup minter params
	// minterNew := app.CfeminterKeeper.GetHalvingMinter(ctx)
	// minterNew.MintDenom = "uc4e"
	// minterNew.NewCoinsMint = 20596877
	// minterNew.BlocksPerYear = 100
	// app.CfeminterKeeper.SetHalvingMinter(ctx, minterNew)

	for i := 1; i < 4000; i++ {
		ctx = ctx.WithBlockHeight(int64(i))
		app.BeginBlocker(ctx, abci.RequestBeginBlock{})
		app.EndBlocker(ctx, abci.RequestEndBlock{})
	}

	//app.BankKeeper.
	//require.Equal(t, app.BankKeeper.GetSupply(ctx, "uC4E").Amount.String(), 20596877, "asd")
}
