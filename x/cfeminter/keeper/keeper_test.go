package keeper_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	routingdistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const PeriodDuration = time.Duration(345600000000 * 1000000)
const MyDenom = "myc4e"

func TestMintFirstPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	app, ctx := prepareApp(startTime, createLinearMinters(startTime))
	k := app.CfeminterKeeper

	minterState := types.MinterState{Position: 1, AmountMinted: sdk.NewInt(0)}
	k.SetMinterState(ctx, minterState)
	minterState.LastMintBlockTime = startTime
	minterState.RemainderToMint = sdk.ZeroDec()
	minterState.RemainderFromPreviousPeriod = sdk.ZeroDec()

	ctx = ctx.WithBlockTime(startTime)
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(0), amount)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.ZeroInt())

	history := k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime := startTime.Add(PeriodDuration / 4)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(250000), amount)
	minterState.AmountMinted = sdk.NewInt(250000)
	minterState.LastMintBlockTime = newTime
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(250000))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime = startTime.Add(PeriodDuration * 3 / 4)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(500000), amount)
	minterState.AmountMinted = sdk.NewInt(750000)
	minterState.LastMintBlockTime = newTime
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(750000))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime = startTime.Add(PeriodDuration)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(250000), amount)
	minterState.AmountMinted = sdk.NewInt(0)
	minterState.LastMintBlockTime = newTime
	minterState.Position = 2
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(1000000))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 1, len(history))

	expectedHist := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(1000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	require.EqualValues(t, expectedHist, history[0])
}

func TestMintSecondPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	app, ctx := prepareApp(startTime, createLinearMinters(startTime))
	k := app.CfeminterKeeper

	minterState := types.MinterState{Position: 2, AmountMinted: sdk.NewInt(0)}
	k.SetMinterState(ctx, minterState)

	periodStart := startTime.Add(PeriodDuration)

	minterState.LastMintBlockTime = periodStart
	minterState.RemainderToMint = sdk.ZeroDec()
	minterState.RemainderFromPreviousPeriod = sdk.ZeroDec()

	ctx = ctx.WithBlockTime(periodStart)
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(0), amount)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(0))

	history := k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime := periodStart.Add(PeriodDuration / 4)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(25000), amount)
	minterState.AmountMinted = sdk.NewInt(25000)
	minterState.LastMintBlockTime = newTime
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(25000))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime = periodStart.Add(PeriodDuration * 3 / 4)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(50000), amount)
	minterState.AmountMinted = sdk.NewInt(75000)
	minterState.LastMintBlockTime = newTime
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(75000))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime = periodStart.Add(PeriodDuration)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(25000), amount)
	minterState.AmountMinted = sdk.NewInt(0)
	minterState.LastMintBlockTime = newTime
	minterState.Position = 3
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(100000))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 1, len(history))

	expectedHist := types.MinterState{
		Position:                    2,
		AmountMinted:                sdk.NewInt(100000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	require.EqualValues(t, expectedHist, history[0])
}

func TestMintBetweenFirstAndSecondPeriods(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	app, ctx := prepareApp(startTime, createLinearMinters(startTime))
	k := app.CfeminterKeeper

	minterState := types.MinterState{Position: 1, AmountMinted: sdk.NewInt(750000)}
	k.SetMinterState(ctx, minterState)

	newTime := startTime.Add(PeriodDuration + PeriodDuration/4)
	minterState.LastMintBlockTime = newTime
	minterState.RemainderToMint = sdk.ZeroDec()
	minterState.RemainderFromPreviousPeriod = sdk.ZeroDec()

	ctx = ctx.WithBlockTime(newTime)
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(275000), amount)
	minterState.AmountMinted = sdk.NewInt(25000)
	minterState.Position = 2
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(275000))

	history := k.GetAllMinterStateHistory(ctx)

	require.EqualValues(t, 1, len(history))

	expectedHist := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(1000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	require.EqualValues(t, expectedHist, history[0])

}

func TestMintBetweenSecondAndThirdPeriods(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	app, ctx := prepareApp(startTime, createLinearMinters(startTime))
	k := app.CfeminterKeeper

	minterState := types.MinterState{Position: 2, AmountMinted: sdk.NewInt(75000)}
	k.SetMinterState(ctx, minterState)

	newTime := startTime.Add(2*PeriodDuration + PeriodDuration/4)
	minterState.LastMintBlockTime = newTime
	minterState.RemainderToMint = sdk.ZeroDec()
	minterState.RemainderFromPreviousPeriod = sdk.ZeroDec()

	ctx = ctx.WithBlockTime(newTime)
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(25000), amount)
	minterState.AmountMinted = sdk.NewInt(0)
	minterState.Position = 3
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(25000))

	history := k.GetAllMinterStateHistory(ctx)

	require.EqualValues(t, 1, len(history))

	expectedHist := types.MinterState{
		Position:                    2,
		AmountMinted:                sdk.NewInt(100000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	require.EqualValues(t, expectedHist, history[0])

}

func TestMintPeriodNotFound(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	app, ctx := prepareApp(startTime, createLinearMinters(startTime))
	k := app.CfeminterKeeper

	minterState := types.MinterState{Position: 9, AmountMinted: sdk.NewInt(0)}
	k.SetMinterState(ctx, minterState)

	ctx = ctx.WithBlockTime(startTime)
	_, err := k.Mint(ctx)
	require.EqualError(t, err, "minter current period for position 9 not found: not found")

}



func TestMintSecondPeriodWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	app, ctx := prepareApp(startTime, createLinearMinters(startTime))
	k := app.CfeminterKeeper

	minterState := types.MinterState{Position: 2, AmountMinted: sdk.NewInt(0), RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("0.5")}
	k.SetMinterState(ctx, minterState)

	periodStart := startTime.Add(PeriodDuration)

	minterState.LastMintBlockTime = periodStart
	minterState.RemainderToMint = sdk.MustNewDecFromStr("0.5")

	ctx = ctx.WithBlockTime(periodStart)
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(0), amount)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(0))

	history := k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime := periodStart.Add(PeriodDuration / 3)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(33333), amount)
	minterState.AmountMinted = sdk.NewInt(33333)
	minterState.LastMintBlockTime = newTime
	minterState.RemainderToMint = sdk.MustNewDecFromStr("0.833333333333333333")

	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(33333))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime = periodStart.Add(PeriodDuration * 2 / 3)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(33334), amount)
	minterState.AmountMinted = sdk.NewInt(66667)
	minterState.LastMintBlockTime = newTime
	minterState.RemainderToMint = sdk.MustNewDecFromStr("0.166666666666666666")


	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(66667))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime = periodStart.Add(PeriodDuration)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(33333), amount)
	minterState.AmountMinted = sdk.NewInt(0)
	minterState.LastMintBlockTime = newTime
	minterState.RemainderToMint = sdk.MustNewDecFromStr("0.500000000000000000")

	minterState.Position = 3
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(100000))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 1, len(history))

	expectedHist := types.MinterState{
		Position:                    2,
		AmountMinted:                sdk.NewInt(100000),
		RemainderToMint:             sdk.MustNewDecFromStr("0.500000000000000000"),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("0.500000000000000000"),
	}
	require.EqualValues(t, expectedHist, history[0])
}

func TestMintFirstPeriodWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	app, ctx := prepareApp(startTime, createReductionMinterWithRemainingPassing(startTime))
	k := app.CfeminterKeeper

	minterState := types.MinterState{Position: 1, AmountMinted: sdk.NewInt(0)}
	k.SetMinterState(ctx, minterState)
	minterState.LastMintBlockTime = startTime
	minterState.RemainderToMint = sdk.ZeroDec()
	minterState.RemainderFromPreviousPeriod = sdk.ZeroDec()

	ctx = ctx.WithBlockTime(startTime)
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(0), amount)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.ZeroInt())

	history := k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime := startTime.Add(PeriodDuration / 4)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(2739726), amount)
	minterState.AmountMinted = sdk.NewInt(2739726)
	minterState.LastMintBlockTime = newTime
	minterState.RemainderToMint = sdk.MustNewDecFromStr("0.027397260273972602")

	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(2739726))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime = startTime.Add(PeriodDuration * 3 / 4)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(3315068), amount)
	minterState.AmountMinted = sdk.NewInt(2739726 + 3315068)
	minterState.LastMintBlockTime = newTime
	minterState.RemainderToMint = sdk.MustNewDecFromStr("0.520547945205479452")

	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(2739726 + 3315068))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 0, len(history))

	newTime = startTime.Add(PeriodDuration)
	ctx = ctx.WithBlockTime(newTime)
	amount, err = k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(684932), amount)
	minterState.AmountMinted = sdk.NewInt(0)
	minterState.LastMintBlockTime = newTime
	minterState.Position = 2
	minterState.RemainderToMint = sdk.MustNewDecFromStr("0.027397260273972602")
	minterState.RemainderFromPreviousPeriod = sdk.MustNewDecFromStr("0.027397260273972602")

	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(2739726 + 3315068 + 684932))

	history = k.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 1, len(history))

	expectedHist := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(2739726 + 3315068 + 684932),
		RemainderToMint:             sdk.MustNewDecFromStr("0.027397260273972602"),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	require.EqualValues(t, expectedHist, history[0])
}

func TestMintBetweenFirstAndSecondPeriodsWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	app, ctx := prepareApp(startTime, createReductionMinterWithRemainingPassing(startTime))
	k := app.CfeminterKeeper

	minterState := types.MinterState{Position: 1, AmountMinted: sdk.NewInt(750000)}
	k.SetMinterState(ctx, minterState)

	newTime := startTime.Add(PeriodDuration + PeriodDuration/4)
	minterState.LastMintBlockTime = newTime
	minterState.RemainderToMint = sdk.ZeroDec()
	minterState.RemainderFromPreviousPeriod = sdk.ZeroDec()

	ctx = ctx.WithBlockTime(newTime)
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(6014726), amount)
	minterState.AmountMinted = sdk.NewInt(25000)
	minterState.RemainderFromPreviousPeriod = sdk.MustNewDecFromStr("0.027397260273972602")
	minterState.RemainderToMint = sdk.MustNewDecFromStr("0.027397260273972602")

	minterState.Position = 2
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
	commontestutils.VerifyModuleAccountDenomBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, MyDenom, sdk.NewInt(6014726))

	history := k.GetAllMinterStateHistory(ctx)

	require.EqualValues(t, 1, len(history))

	expectedHist := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(6014726 - 25000 + 750000 ),
		RemainderToMint:             sdk.MustNewDecFromStr("0.027397260273972602"),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	require.EqualValues(t, expectedHist, history[0])

}

func prepareApp(startTime time.Time, minter types.Minter) (*app.App, sdk.Context) {
	app, ctx := commontestutils.SetupAppWithTime(1000, startTime)
	params := types.DefaultParams()
	params.MintDenom = MyDenom
	params.Minter = minter

	k := app.CfeminterKeeper
	k.SetParams(ctx, params)
	return app, ctx
}

func createLinearMinters(startTime time.Time) types.Minter {
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2, &period3}
	minter := types.Minter{Start: startTime, Periods: periods}
	return minter
}
const SecondsInYear = int32(3600 * 24 * 365)

func createReductionMinterWithRemainingPassing(startTime time.Time) types.Minter {
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	// linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}

	pminter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(1000000), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &pminter}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2, &period3}
	minter := types.Minter{Start: startTime, Periods: periods}
	return minter
}
