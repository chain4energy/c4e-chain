package cfeminter_test

import (
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/app"
	// "github.com/chain4energy/c4e-chain/testutil/nullify"

	"github.com/chain4energy/c4e-chain/x/cfeminter"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	testminter "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"


)

const PeriodDuration = time.Duration(345600000000 * 1000000)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.NewParams("myc4e"),
		Minter: createMinter(time.Now()),
		MinterState: types.MinterState{CurrentOrderingId: 9, AmountMinted: sdk.NewInt(12312)},

		// this line is used by starport scaffolding # genesis/test/state
		
	}

	app := testapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	cfeminter.InitGenesis(ctx, app.CfeminterKeeper, app.AccountKeeper, genesisState)
	got := cfeminter.ExportGenesis(ctx, app.CfeminterKeeper)
	require.NotNil(t, got)
	// nullify.Fill(&genesisState)
	// nullify.Fill(got)

	require.EqualValues(t, genesisState.Params, got.Params)
	testminter.CompareMinters(t, genesisState.Minter, got.Minter)
	require.EqualValues(t, genesisState.MinterState, got.MinterState)

	// this line is used by starport scaffolding # genesis/test/assert
}

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


func createMinter(startTime time.Time) types.Minter {
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2, &period3}
	minter := types.Minter{Start: startTime, Periods: periods}
	return minter
}
