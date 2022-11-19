package keeper_test

import (
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const PeriodDuration = time.Duration(345600000000 * 1000000)
const MyDenom = "myc4e"

func TestMintFirstPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, createLinearMinters(startTime))

	testHelper.C4eMinterUtils.SetMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	testHelper.SetContextBlockTime(startTime)
	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 1, sdk.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec(), sdk.ZeroInt())

	newTime := startTime.Add(PeriodDuration / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(250000), 1, sdk.NewInt(250000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), sdk.NewInt(250000))

	newTime = startTime.Add(PeriodDuration * 3 / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(500000), 1, sdk.NewInt(750000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), sdk.NewInt(750000))

	newTime = startTime.Add(PeriodDuration)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(1000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(250000), 2, sdk.ZeroInt(), sdk.ZeroDec(), newTime, sdk.ZeroDec(), sdk.NewInt(1000000), expectedHist)

}

func TestMintSecondPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, createLinearMinters(startTime))

	testHelper.C4eMinterUtils.SetMinterState(2, sdk.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	periodStart := startTime.Add(PeriodDuration)

	testHelper.SetContextBlockTime(periodStart)
	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 2, sdk.ZeroInt(), sdk.ZeroDec(), periodStart, sdk.ZeroDec(), sdk.ZeroInt())

	newTime := periodStart.Add(PeriodDuration / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(25000), 2, sdk.NewInt(25000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), sdk.NewInt(25000))

	newTime = periodStart.Add(PeriodDuration * 3 / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(50000), 2, sdk.NewInt(75000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), sdk.NewInt(75000))

	newTime = periodStart.Add(PeriodDuration)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		Position:                    2,
		AmountMinted:                sdk.NewInt(100000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(25000), 3, sdk.ZeroInt(), sdk.ZeroDec(), newTime, sdk.ZeroDec(), sdk.NewInt(100000), expectedHist)
}

func TestMintBetweenFirstAndSecondPeriods(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, createLinearMinters(startTime))

	testHelper.C4eMinterUtils.SetMinterState(1, sdk.NewInt(750000), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	newTime := startTime.Add(PeriodDuration + PeriodDuration/4)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(1000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(275000), 2, sdk.NewInt(25000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), sdk.NewInt(275000), expectedHist)
}

func TestMintBetweenSecondAndThirdPeriods(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, createLinearMinters(startTime))

	testHelper.C4eMinterUtils.SetMinterState(2, sdk.NewInt(75000), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	newTime := startTime.Add(2*PeriodDuration + PeriodDuration/4)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		Position:                    2,
		AmountMinted:                sdk.NewInt(100000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(25000), 3, sdk.NewInt(0), sdk.ZeroDec(), newTime, sdk.ZeroDec(), sdk.NewInt(25000), expectedHist)
}

func TestMintPeriodNotFound(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	testHelper := prepareApp(t, startTime, createLinearMinters(startTime))

	testHelper.C4eMinterUtils.SetMinterState(9, sdk.NewInt(0), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	newTime := startTime.Add(10)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.MintError("minter - mint - current period for position 9 not found: not found")
}

func TestMintSecondPeriodWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, createLinearMinters(startTime))

	testHelper.C4eMinterUtils.SetMinterState(2, sdk.NewInt(0), sdk.ZeroDec(), startTime, sdk.MustNewDecFromStr("0.5"))

	periodStart := startTime.Add(PeriodDuration)
	testHelper.SetContextBlockTime(periodStart)
	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 2, sdk.NewInt(0), sdk.MustNewDecFromStr("0.5"), periodStart, sdk.MustNewDecFromStr("0.5"), sdk.ZeroInt())

	newTime := periodStart.Add(PeriodDuration / 3)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(33333), 2, sdk.NewInt(33333), sdk.MustNewDecFromStr("0.833333333333333333"), newTime, sdk.MustNewDecFromStr("0.5"), sdk.NewInt(33333))

	newTime = periodStart.Add(PeriodDuration * 2 / 3)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(33334), 2, sdk.NewInt(66667), sdk.MustNewDecFromStr("0.166666666666666666"), newTime, sdk.MustNewDecFromStr("0.5"), sdk.NewInt(66667))

	newTime = periodStart.Add(PeriodDuration)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		Position:                    2,
		AmountMinted:                sdk.NewInt(100000),
		RemainderToMint:             sdk.MustNewDecFromStr("0.5"),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("0.5"),
	}
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(33333), 3, sdk.ZeroInt(), sdk.MustNewDecFromStr("0.5"), newTime, sdk.MustNewDecFromStr("0.5"), sdk.NewInt(100000), expectedHist)
}

func TestMintFirstPeriodWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, createReductionMinterWithRemainingPassing(startTime))

	testHelper.C4eMinterUtils.SetMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	testHelper.SetContextBlockTime(startTime)
	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 1, sdk.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec(), sdk.ZeroInt())

	newTime := startTime.Add(PeriodDuration / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(2739726), 1, sdk.NewInt(2739726), sdk.MustNewDecFromStr("0.027397260273972602"), newTime, sdk.ZeroDec(), sdk.NewInt(2739726))

	newTime = startTime.Add(PeriodDuration * 3 / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(3315068), 1, sdk.NewInt(2739726+3315068), sdk.MustNewDecFromStr("0.520547945205479452"), newTime, sdk.ZeroDec(), sdk.NewInt(2739726+3315068))

	newTime = startTime.Add(PeriodDuration)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(2739726 + 3315068 + 684932),
		RemainderToMint:             sdk.MustNewDecFromStr("0.027397260273972602"),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(684932), 2, sdk.ZeroInt(), sdk.MustNewDecFromStr("0.027397260273972602"), newTime, sdk.MustNewDecFromStr("0.027397260273972602"), sdk.NewInt(2739726+3315068+684932), expectedHist)
}

func TestMintBetweenFirstAndSecondPeriodsWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, createReductionMinterWithRemainingPassing(startTime))
	testHelper.C4eMinterUtils.SetMinterState(1, sdk.NewInt(750000), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	newTime := startTime.Add(PeriodDuration + PeriodDuration/4)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(6014726 - 25000 + 750000),
		RemainderToMint:             sdk.MustNewDecFromStr("0.027397260273972602"),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(6014726), 2, sdk.NewInt(25000), sdk.MustNewDecFromStr("0.027397260273972602"), newTime, sdk.MustNewDecFromStr("0.027397260273972602"), sdk.NewInt(6014726), expectedHist)
}

func TestMintWithReductionMinterOnGenesisMinterStateAfterBlockTime(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime.Add(time.Hour), createReductionMinter(startTime))

	testHelper.C4eMinterUtils.SetMinterState(1, sdk.NewInt(1000000), sdk.ZeroDec(), startTime.Add(2*time.Hour), sdk.ZeroDec())

	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 1, sdk.NewInt(1000000), sdk.ZeroDec(), startTime.Add(2*time.Hour), sdk.ZeroDec(), sdk.ZeroInt())
}

func TestMintWithReductionMinterOnGenesisStartInTheFuture(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime.Add(time.Hour), createReductionMinter(startTime.Add(2*time.Hour)))

	testHelper.C4eMinterUtils.SetMinterState(1, sdk.NewInt(1000000), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 1, sdk.NewInt(1000000), sdk.ZeroDec(), startTime, sdk.ZeroDec(), sdk.ZeroInt())
}

func prepareApp(t *testing.T, startTime time.Time, minter types.Minter) *testapp.TestHelper {
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)
	params := types.DefaultParams()
	params.MintDenom = commontestutils.DefaultTestDenom
	params.Minter = minter

	k := testHelper.App.CfeminterKeeper
	k.SetParams(testHelper.Context, params)
	return testHelper
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

	pminter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(1000000), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &pminter}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2, &period3}
	return types.Minter{Start: startTime, Periods: periods}
}

func createReductionMinter(startTime time.Time) types.Minter {
	pminter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(1000000), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}
	period1 := types.MintingPeriod{Position: 1, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &pminter}
	periods := []*types.MintingPeriod{&period1}
	return types.Minter{Start: startTime, Periods: periods}

}
