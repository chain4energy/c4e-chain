package keeper_test

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/v2/testutil/app"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"testing"
	"time"

	testenv "github.com/chain4energy/c4e-chain/v2/testutil/env"
	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const PeriodDuration = time.Duration(345600000000 * 1000000)

func TestMintFirstPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

	testHelper.C4eMinterUtils.SetMinterState(1, math.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	testHelper.SetContextBlockTime(startTime)
	testHelper.C4eMinterUtils.Mint(math.ZeroInt(), 1, math.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec(), math.ZeroInt())

	newTime := startTime.Add(PeriodDuration / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(math.NewInt(250000), 1, math.NewInt(250000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), math.NewInt(250000))

	newTime = startTime.Add(PeriodDuration * 3 / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(math.NewInt(500000), 1, math.NewInt(750000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), math.NewInt(750000))

	newTime = startTime.Add(PeriodDuration)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		SequenceId:                  1,
		AmountMinted:                math.NewInt(1000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousMinter: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(math.NewInt(250000), 2, math.ZeroInt(), sdk.ZeroDec(), newTime, sdk.ZeroDec(), math.NewInt(1000000), expectedHist)
}

func TestMintSecondPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

	testHelper.C4eMinterUtils.SetMinterState(2, math.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	minterStart := startTime.Add(PeriodDuration)

	testHelper.SetContextBlockTime(minterStart)
	testHelper.C4eMinterUtils.Mint(math.ZeroInt(), 2, math.ZeroInt(), sdk.ZeroDec(), minterStart, sdk.ZeroDec(), math.ZeroInt())

	newTime := minterStart.Add(PeriodDuration / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(math.NewInt(25000), 2, math.NewInt(25000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), math.NewInt(25000))

	newTime = minterStart.Add(PeriodDuration * 3 / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(math.NewInt(50000), 2, math.NewInt(75000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), math.NewInt(75000))

	newTime = minterStart.Add(PeriodDuration)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		SequenceId:                  2,
		AmountMinted:                math.NewInt(100000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousMinter: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(math.NewInt(25000), 3, math.ZeroInt(), sdk.ZeroDec(), newTime, sdk.ZeroDec(), math.NewInt(100000), expectedHist)
}

func TestMintBetweenFirstAndSecondMinters(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

	testHelper.C4eMinterUtils.SetMinterState(1, math.NewInt(750000), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	newTime := startTime.Add(PeriodDuration + PeriodDuration/4)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		SequenceId:                  1,
		AmountMinted:                math.NewInt(1000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousMinter: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(math.NewInt(275000), 2, math.NewInt(25000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), math.NewInt(275000), expectedHist)
}

func TestMintBetweenSecondAndThirdMinters(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

	testHelper.C4eMinterUtils.SetMinterState(2, math.NewInt(75000), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	newTime := startTime.Add(2*PeriodDuration + PeriodDuration/4)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		SequenceId:                  2,
		AmountMinted:                math.NewInt(100000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousMinter: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(math.NewInt(25000), 3, math.NewInt(0), sdk.ZeroDec(), newTime, sdk.ZeroDec(), math.NewInt(25000), expectedHist)
}

func TestStepDurationNotFound(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

	testHelper.C4eMinterUtils.SetMinterState(9, math.NewInt(0), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	newTime := startTime.Add(10)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.MintError("minter - mint - current minter for sequence id 9 not found: not found")
}

func TestMintSecondPeriodWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

	testHelper.C4eMinterUtils.SetMinterState(2, math.NewInt(0), sdk.ZeroDec(), startTime, sdk.MustNewDecFromStr("0.5"))

	minterStart := startTime.Add(PeriodDuration)
	testHelper.SetContextBlockTime(minterStart)
	testHelper.C4eMinterUtils.Mint(math.ZeroInt(), 2, math.NewInt(0), sdk.MustNewDecFromStr("0.5"), minterStart, sdk.MustNewDecFromStr("0.5"), math.ZeroInt())

	newTime := minterStart.Add(PeriodDuration / 3)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(math.NewInt(33333), 2, math.NewInt(33333), sdk.MustNewDecFromStr("0.833333333333333333"), newTime, sdk.MustNewDecFromStr("0.5"), math.NewInt(33333))

	newTime = minterStart.Add(PeriodDuration * 2 / 3)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(math.NewInt(33334), 2, math.NewInt(66667), sdk.MustNewDecFromStr("0.166666666666666666"), newTime, sdk.MustNewDecFromStr("0.5"), math.NewInt(66667))

	newTime = minterStart.Add(PeriodDuration)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		SequenceId:                  2,
		AmountMinted:                math.NewInt(100000),
		RemainderToMint:             sdk.MustNewDecFromStr("0.5"),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousMinter: sdk.MustNewDecFromStr("0.5"),
	}
	testHelper.C4eMinterUtils.Mint(math.NewInt(33333), 3, math.ZeroInt(), sdk.MustNewDecFromStr("0.5"), newTime, sdk.MustNewDecFromStr("0.5"), math.NewInt(100000), expectedHist)
}

func TestMintFirstPeriodWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createExponentialStepMintingWithRemainingPassing(startTime))

	testHelper.C4eMinterUtils.SetMinterState(1, math.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	testHelper.SetContextBlockTime(startTime)
	testHelper.C4eMinterUtils.Mint(math.ZeroInt(), 1, math.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec(), math.ZeroInt())

	newTime := startTime.Add(PeriodDuration / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(math.NewInt(2739726), 1, math.NewInt(2739726), sdk.MustNewDecFromStr("0.027397260273972602"), newTime, sdk.ZeroDec(), math.NewInt(2739726))

	newTime = startTime.Add(PeriodDuration * 3 / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(math.NewInt(3315068), 1, math.NewInt(2739726+3315068), sdk.MustNewDecFromStr("0.520547945205479452"), newTime, sdk.ZeroDec(), math.NewInt(2739726+3315068))

	newTime = startTime.Add(PeriodDuration)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		SequenceId:                  1,
		AmountMinted:                math.NewInt(2739726 + 3315068 + 684932),
		RemainderToMint:             sdk.MustNewDecFromStr("0.027397260273972602"),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousMinter: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(math.NewInt(684932), 2, math.ZeroInt(), sdk.MustNewDecFromStr("0.027397260273972602"), newTime, sdk.MustNewDecFromStr("0.027397260273972602"), math.NewInt(2739726+3315068+684932), expectedHist)
}

func TestMintBetweenFirstAndSecondMintersWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createExponentialStepMintingWithRemainingPassing(startTime))
	testHelper.C4eMinterUtils.SetMinterState(1, math.NewInt(750000), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	newTime := startTime.Add(PeriodDuration + PeriodDuration/4)
	testHelper.SetContextBlockTime(newTime)
	expectedHist := types.MinterState{
		SequenceId:                  1,
		AmountMinted:                math.NewInt(6014726 - 25000 + 750000),
		RemainderToMint:             sdk.MustNewDecFromStr("0.027397260273972602"),
		LastMintBlockTime:           newTime,
		RemainderFromPreviousMinter: sdk.ZeroDec(),
	}
	testHelper.C4eMinterUtils.Mint(math.NewInt(6014726), 2, math.NewInt(25000), sdk.MustNewDecFromStr("0.027397260273972602"), newTime, sdk.MustNewDecFromStr("0.027397260273972602"), math.NewInt(6014726), expectedHist)
}

func TestMintWithExponentialStepMintingOnGenesisMinterStateAfterBlockTime(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime.Add(time.Hour), startTime, createExponentialStepMinting())

	testHelper.C4eMinterUtils.SetMinterState(1, math.NewInt(1000000), sdk.ZeroDec(), startTime.Add(2*time.Hour), sdk.ZeroDec())
	testHelper.C4eMinterUtils.Mint(math.ZeroInt(), 1, math.NewInt(1000000), sdk.ZeroDec(), startTime.Add(2*time.Hour), sdk.ZeroDec(), math.ZeroInt())
}

func TestMintWithExponentialStepMintingOnGenesisStartInTheFuture(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime.Add(time.Hour), startTime.Add(2*time.Hour), createExponentialStepMinting())

	testHelper.C4eMinterUtils.SetMinterState(1, math.NewInt(1000000), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	testHelper.C4eMinterUtils.Mint(math.ZeroInt(), 1, math.NewInt(1000000), sdk.ZeroDec(), startTime, sdk.ZeroDec(), math.ZeroInt())
}

func TestMintWithExponentialStepMintingMinterStateAmountTooBig(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime.Add(time.Hour), startTime.Add(time.Hour), createExponentialStepMinting())

	testHelper.C4eMinterUtils.SetMinterState(1, math.NewInt(1000000), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	testHelper.C4eMinterUtils.Mint(math.ZeroInt(), 1, math.NewInt(1000000), sdk.ZeroDec(), startTime, sdk.ZeroDec(), math.ZeroInt())
}

func prepareApp(t *testing.T, initialBlockTime time.Time, mintingStartTime time.Time, minters []*types.Minter) *app.TestHelper {
	testHelper := app.SetupTestAppWithHeightAndTime(t, 1000, initialBlockTime)
	params := types.Params{
		MintDenom: testenv.DefaultTestDenom,
		StartTime: mintingStartTime,
		Minters:   minters,
	}

	k := testHelper.App.CfeminterKeeper
	k.SetParams(testHelper.Context, params)
	return testHelper
}

func createLinearMintings(startTime time.Time) []*types.Minter {
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3, Config: testenv.NoMintingConfig}

	return []*types.Minter{&minter1, &minter2, &minter3}
}

const NanoSecondsInFourYears = 3600 * 24 * 365 * 4 * time.Second

func createExponentialStepMintingWithRemainingPassing(startTime time.Time) []*types.Minter {
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	exponentialStepMinting := types.ExponentialStepMinting{Amount: math.NewInt(4000000), StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	linearMinting := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3, Config: testenv.NoMintingConfig}

	return []*types.Minter{&minter1, &minter2, &minter3}
}

func createExponentialStepMinting() []*types.Minter {
	exponentialStepMinting := types.ExponentialStepMinting{Amount: math.NewInt(1000000), StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)

	minter := types.Minter{SequenceId: 1, Config: config}
	return []*types.Minter{&minter}
}
