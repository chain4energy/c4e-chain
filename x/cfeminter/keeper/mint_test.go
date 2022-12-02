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

func TestMintFirstPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

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

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

	testHelper.C4eMinterUtils.SetMinterState(2, sdk.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	Minterstart := startTime.Add(PeriodDuration)

	testHelper.SetContextBlockTime(Minterstart)
	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 2, sdk.ZeroInt(), sdk.ZeroDec(), Minterstart, sdk.ZeroDec(), sdk.ZeroInt())

	newTime := Minterstart.Add(PeriodDuration / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(25000), 2, sdk.NewInt(25000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), sdk.NewInt(25000))

	newTime = Minterstart.Add(PeriodDuration * 3 / 4)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(50000), 2, sdk.NewInt(75000), sdk.ZeroDec(), newTime, sdk.ZeroDec(), sdk.NewInt(75000))

	newTime = Minterstart.Add(PeriodDuration)
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

func TestMintBetweenFirstAndSecondMinters(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

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

func TestMintBetweenSecondAndThirdMinters(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

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

func TestStepDurationNotFound(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

	testHelper.C4eMinterUtils.SetMinterState(9, sdk.NewInt(0), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	newTime := startTime.Add(10)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.MintError("minter - mint - current period for SequenceId 9 not found: not found")
}

func TestMintSecondPeriodWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))

	testHelper.C4eMinterUtils.SetMinterState(2, sdk.NewInt(0), sdk.ZeroDec(), startTime, sdk.MustNewDecFromStr("0.5"))

	Minterstart := startTime.Add(PeriodDuration)
	testHelper.SetContextBlockTime(Minterstart)
	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 2, sdk.NewInt(0), sdk.MustNewDecFromStr("0.5"), Minterstart, sdk.MustNewDecFromStr("0.5"), sdk.ZeroInt())

	newTime := Minterstart.Add(PeriodDuration / 3)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(33333), 2, sdk.NewInt(33333), sdk.MustNewDecFromStr("0.833333333333333333"), newTime, sdk.MustNewDecFromStr("0.5"), sdk.NewInt(33333))

	newTime = Minterstart.Add(PeriodDuration * 2 / 3)
	testHelper.SetContextBlockTime(newTime)
	testHelper.C4eMinterUtils.Mint(sdk.NewInt(33334), 2, sdk.NewInt(66667), sdk.MustNewDecFromStr("0.166666666666666666"), newTime, sdk.MustNewDecFromStr("0.5"), sdk.NewInt(66667))

	newTime = Minterstart.Add(PeriodDuration)
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

	testHelper := prepareApp(t, startTime, startTime, createExponentialStepMintingWithRemainingPassing(startTime))

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

func TestMintBetweenFirstAndSecondMintersWithRemaining(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime, startTime, createExponentialStepMintingWithRemainingPassing(startTime))
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

func TestMintWithExponentialStepMintingOnGenesisMinterStateAfterBlockTime(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime.Add(time.Hour), startTime, createExponentialStepMinting())

	testHelper.C4eMinterUtils.SetMinterState(1, sdk.NewInt(1000000), sdk.ZeroDec(), startTime.Add(2*time.Hour), sdk.ZeroDec())
	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 1, sdk.NewInt(1000000), sdk.ZeroDec(), startTime.Add(2*time.Hour), sdk.ZeroDec(), sdk.ZeroInt())
}

func TestMintWithExponentialStepMintingOnGenesisStartInTheFuture(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper := prepareApp(t, startTime.Add(time.Hour), startTime.Add(2*time.Hour), createExponentialStepMinting())

	testHelper.C4eMinterUtils.SetMinterState(1, sdk.NewInt(1000000), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 1, sdk.NewInt(1000000), sdk.ZeroDec(), startTime, sdk.ZeroDec(), sdk.ZeroInt())
}

//func TestMintWithExponentialStepMintingOnGenesisStartInTheFutureXXXX(t *testing.T) {
//	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
//
//	testHelper := prepareApp(t, startTime.Add(time.Hour), startTime.Add(time.Hour), createExponentialStepMinting())
//
//	testHelper.C4eMinterUtils.SetMinterState(1, sdk.NewInt(1000000), sdk.ZeroDec(), startTime, sdk.ZeroDec())
//
//	testHelper.C4eMinterUtils.Mint(sdk.ZeroInt(), 1, sdk.NewInt(1000000), sdk.ZeroDec(), startTime, sdk.ZeroDec(), sdk.ZeroInt())
//}

func prepareApp(t *testing.T, initialBlockTime time.Time, mintingStartTime time.Time, minters []*types.Minter) *testapp.TestHelper {
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, initialBlockTime)
	params := types.DefaultParams()
	params.MintDenom = commontestutils.DefaultTestDenom
	params.StartTime = mintingStartTime
	params.Minters = minters

	k := testHelper.App.CfeminterKeeper
	k.SetParams(testHelper.Context, params)
	return testHelper
}

func createLinearMintings(startTime time.Time) []*types.Minter {
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.LINEAR_MINTING, LinearMinting: &LinearMinting1}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.LINEAR_MINTING, LinearMinting: &LinearMinting2}

	minter3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	return []*types.Minter{&minter1, &minter2, &minter3}
}

const NanoSecondsInFourYears = 3600 * 24 * 365 * 4 * time.Second

func createExponentialStepMintingWithRemainingPassing(startTime time.Time) []*types.Minter {
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	pminter := types.ExponentialStepMinting{Amount: sdk.NewInt(4000000), StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}

	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.EXPONENTIAL_STEP_MINTING, ExponentialStepMinting: &pminter}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.LINEAR_MINTING, LinearMinting: &LinearMinting2}

	minter3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	return []*types.Minter{&minter1, &minter2, &minter3}
}

func createExponentialStepMinting() []*types.Minter {
	exponentialStepMinting := types.ExponentialStepMinting{Amount: sdk.NewInt(1000000), StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	minter1 := types.Minter{SequenceId: 1, Type: types.EXPONENTIAL_STEP_MINTING, ExponentialStepMinting: &exponentialStepMinting}
	return []*types.Minter{&minter1}
}
