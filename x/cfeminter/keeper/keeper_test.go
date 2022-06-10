package keeper_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

const PeriodDuration = time.Duration(345600000000 * 1000000)
const MyDenom = "myc4e"

func TestMintFirstPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	app, ctx := prepareApp(startTime)
	k := app.CfeminterKeeper

	minterState := types.MinterState{CurrentOrderingId: 1, AmountMinted: sdk.NewInt(0)}
	k.SetMinterState(ctx, minterState)

	minter := createMinter(startTime)

	k.SetMinter(ctx, minter)

	ctx = ctx.WithBlockTime(startTime)
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(0), amount)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(authtypes.FeeCollectorName, ctx, app, t, MyDenom, sdk.ZeroInt())

	ctx = ctx.WithBlockTime(startTime.Add(PeriodDuration / 4))
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(250000), amount)
	minterState.AmountMinted = sdk.NewInt(250000)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(authtypes.FeeCollectorName, ctx, app, t, MyDenom, sdk.NewInt(250000))

	ctx = ctx.WithBlockTime(startTime.Add(PeriodDuration * 3 / 4))
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(500000), amount)
	minterState.AmountMinted = sdk.NewInt(750000)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(authtypes.FeeCollectorName, ctx, app, t, MyDenom, sdk.NewInt(750000))

	ctx = ctx.WithBlockTime(startTime.Add(PeriodDuration))
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(250000), amount)
	minterState.AmountMinted = sdk.NewInt(0)
	minterState.CurrentOrderingId = 2
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(authtypes.FeeCollectorName, ctx, app, t, MyDenom, sdk.NewInt(1000000))

}

func TestMintSecondPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	app, ctx := prepareApp(startTime)
	k := app.CfeminterKeeper

	minterState := types.MinterState{CurrentOrderingId: 2, AmountMinted: sdk.NewInt(0)}
	k.SetMinterState(ctx, minterState)

	minter := createMinter(startTime)

	k.SetMinter(ctx, minter)
	periodStart := startTime.Add(PeriodDuration)
	ctx = ctx.WithBlockTime(periodStart)
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(0), amount)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(authtypes.FeeCollectorName, ctx, app, t, MyDenom, sdk.NewInt(0))

	ctx = ctx.WithBlockTime(periodStart.Add(PeriodDuration / 4))
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(25000), amount)
	minterState.AmountMinted = sdk.NewInt(25000)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(authtypes.FeeCollectorName, ctx, app, t, MyDenom, sdk.NewInt(25000))

	ctx = ctx.WithBlockTime(periodStart.Add(PeriodDuration * 3 / 4))
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(50000), amount)
	minterState.AmountMinted = sdk.NewInt(75000)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(authtypes.FeeCollectorName, ctx, app, t, MyDenom, sdk.NewInt(75000))

	ctx = ctx.WithBlockTime(periodStart.Add(PeriodDuration))
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(25000), amount)
	minterState.AmountMinted = sdk.NewInt(0)
	minterState.CurrentOrderingId = 3
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(authtypes.FeeCollectorName, ctx, app, t, MyDenom, sdk.NewInt(100000))

}

func TestMintBetweenFirstAndSecondPeriods(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	app, ctx := prepareApp(startTime)
	k := app.CfeminterKeeper

	minterState := types.MinterState{CurrentOrderingId: 1, AmountMinted: sdk.NewInt(750000)}
	k.SetMinterState(ctx, minterState)

	minter := createMinter(startTime)

	k.SetMinter(ctx, minter)

	ctx = ctx.WithBlockTime(startTime.Add(PeriodDuration + PeriodDuration/4))
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(275000), amount)
	minterState.AmountMinted = sdk.NewInt(25000)
	minterState.CurrentOrderingId = 2
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(authtypes.FeeCollectorName, ctx, app, t, MyDenom, sdk.NewInt(275000))

}

func TestMintBetweenSecondAndThirdPeriods(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	app, ctx := prepareApp(startTime)
	k := app.CfeminterKeeper

	minterState := types.MinterState{CurrentOrderingId: 2, AmountMinted: sdk.NewInt(75000)}
	k.SetMinterState(ctx, minterState)

	minter := createMinter(startTime)

	k.SetMinter(ctx, minter)

	ctx = ctx.WithBlockTime(startTime.Add(2*PeriodDuration + PeriodDuration/4))
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(25000), amount)
	minterState.AmountMinted = sdk.NewInt(0)
	minterState.CurrentOrderingId = 3
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(authtypes.FeeCollectorName, ctx, app, t, MyDenom, sdk.NewInt(25000))

}

func prepareApp(startTime time.Time) (*app.App, sdk.Context) {
	app, ctx := commontestutils.SetupAppWithTime(1000, startTime)
	params := types.DefaultParams()
	params.MintDenom = MyDenom

	k := app.CfeminterKeeper
	k.SetParams(ctx, params)
	return app, ctx
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
