package types_test

import (
	"cosmossdk.io/math"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"testing"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const PeriodDuration = time.Duration(345600000000 * 1000000)
const Year = time.Hour * 24 * 365
const NanoSecondsInFourYears = Year * 4
const customDenom = "uc4e"

func TestLinearMinting(t *testing.T) {
	linearMinting := types.LinearMinting{Amount: math.NewInt(1000000)}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(time.Duration(345600000000 * 1000000))
	blockTime := startTime.Add(time.Duration(345600000000 * 1000000 / 2))
	config, _ := codectypes.NewAnyWithValue(&linearMinting)

	minter := types.Minter{SequenceId: 1, EndTime: &endTime, Config: config}
	amount := minter.AmountToMint(log.TestingLogger(), startTime, blockTime)
	require.EqualValues(t, sdk.NewDec(500000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, endTime)
	require.EqualValues(t, sdk.NewDec(1000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, endTime.Add(time.Duration(10*1000000)))
	require.EqualValues(t, sdk.NewDec(1000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(time.Duration(345600000000*1000000*3/4)))
	require.EqualValues(t, sdk.NewDec(750000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(time.Duration(345600000000*1000000/4)))
	require.EqualValues(t, sdk.NewDec(250000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime)
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(time.Duration(-10*1000000)))
	require.EqualValues(t, sdk.NewDec(0), amount)

}

func TestNoMinting(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(time.Duration(345600000000 * 1000000))
	blockTime := startTime.Add(time.Duration(345600000000 * 1000000 / 2))

	minter := types.Minter{SequenceId: 1, EndTime: &endTime, Config: testenv.NoMintingConfig}
	amount := minter.AmountToMint(log.TestingLogger(), startTime, blockTime)
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, endTime)
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, endTime.Add(time.Duration(10*1000000)))
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(time.Duration(345600000000*1000000*3/4)))
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(time.Duration(345600000000*1000000/4)))
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime)
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(time.Duration(-10*1000000)))
	require.EqualValues(t, sdk.NewDec(0), amount)
}

func TestValidateMinterPariodsOrder(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3, Config: testenv.NoMintingConfig}
	minters := []*types.Minter{&minter1, &minter2, &minter3}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.NoError(t, params.Validate())
}

func TestValidateMinterPariodsOrderInitialyNotOrdered(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3, Config: testenv.NoMintingConfig}
	minters := []*types.Minter{&minter3, &minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.NoError(t, params.Validate())
}

func TestValidateMinterPariodsOrderInitialyNotFromOne(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 5, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 6, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 7, Config: testenv.NoMintingConfig}
	minters := []*types.Minter{&minter3, &minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.NoError(t, params.Validate())
}

func TestValidateMinterPariodsOrderWrongFirstId(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 0, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3}
	minters := []*types.Minter{&minter3, &minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "first minter sequence id must be bigger than 0, but is 0")
}

func TestValidateMinterPariodsOrderWrongNotIncrementByOne(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 3, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 4}
	minters := []*types.Minter{&minter3, &minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "missing minter with sequence id 2")
}

func TestValidateMinterNoMinters(t *testing.T) {
	params := types.Params{MintDenom: customDenom, StartTime: time.Now(), Minters: []*types.Minter{}}
	require.EqualError(t, params.Validate(), "no minters defined")
}

func TestValidateMinterSecondMinterIsNil(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(PeriodDuration)
	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minters := []*types.Minter{&minter1, nil}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "minter on position 2 cannot be nil")
}

func TestValidateMinterfirstMinterNil(t *testing.T) {
	startTime := time.Now()
	minters := []*types.Minter{nil}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "minter on position 1 cannot be nil")
}

func TestValidateMinterLastPeriodWithEndDate(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minters := []*types.Minter{&minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "last minter cannot have EndTime set, but is set to 2043-12-30 00:00:00 +0000 UTC")
}

func TestValidateMinterLastPeriodWithEndDateOnePeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minters := []*types.Minter{&minter1}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "last minter cannot have EndTime set, but is set to 2033-01-16 00:00:00 +0000 UTC")
}

func TestValidateMinterFirstPeriodWrongEnd(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime
	endTime2 := endTime1.Add(2 * PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 4}
	minters := []*types.Minter{&minter3, &minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "first minter end must be bigger than minter start")
}

func TestValidateMinterNextPeriodWrongEnd(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 4}
	minters := []*types.Minter{&minter3, &minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "minter with sequence id 2 mast have EndTime bigger than minter with sequence id 1")
}

func TestValidateMinterNoMintigType(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	minter1 := types.Minter{SequenceId: 1, Config: testenv.NoMintingConfig}
	minters := []*types.Minter{&minter1}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.NoError(t, params.Validate())
}

func TestValidateMinterTimeLineraMinterTypeWithNoEndTimeInNotLastPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: nil, Config: config}
	minter3 := types.Minter{SequenceId: 3}
	minters := []*types.Minter{&minter3, &minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "only last minter can have EndTime empty")
}

func TestValidateMinterTimeLineraMinterTypeWithNoEndTime(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: nil, Config: config}
	minters := []*types.Minter{&minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "minter with id 2 validation error: for LinearMinting EndTime must be set")
}

func TestValidateExponentialStepMintingAmountIsNil(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(PeriodDuration)

	exponentialStepMinting := types.ExponentialStepMinting{Amount: math.Int{}, StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2}
	minters := []*types.Minter{&minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "minter with id 1 validation error: minter config validation error: amount must be positive")
}

func TestValidateExponentialStepMintingAmountMultiplierIsLowerThan0(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(PeriodDuration)

	exponentialStepMinting := types.ExponentialStepMinting{Amount: math.NewInt(1000), StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("-1")}
	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2}
	minters := []*types.Minter{&minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "minter with id 1 validation error: minter config validation error: amountMultiplier cannot be less than 0")
}

func TestValidateMinterTimeLinearAmountLessThanZero(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(-100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3}
	minters := []*types.Minter{&minter3, &minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "minter with id 2 validation error: minter config validation error: amount cannot be less than 0")
}

func TestCointainsIdTrue(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3}
	minters := []*types.Minter{&minter3, &minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.True(t, params.ContainsMinter(3))
}

func TestCointainsIdFalse(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3}
	minters := []*types.Minter{&minter3, &minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.False(t, params.ContainsMinter(6))
}

func TestValidateMinterState(t *testing.T) {
	minterState := types.MinterState{SequenceId: 1, AmountMinted: math.ZeroInt(), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousMinter: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{SequenceId: 1, AmountMinted: math.NewInt(123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousMinter: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{SequenceId: 1, AmountMinted: math.NewInt(-123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousMinter: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.EqualError(t, minterState.Validate(), "minter state validation error: amountMinted cannot be less than 0")

	minterState = types.MinterState{SequenceId: 1, AmountMinted: math.NewInt(123), RemainderToMint: sdk.MustNewDecFromStr("231321.1234"), RemainderFromPreviousMinter: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{SequenceId: 1, AmountMinted: math.NewInt(123), RemainderToMint: sdk.MustNewDecFromStr("-231321.1234"), RemainderFromPreviousMinter: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.EqualError(t, minterState.Validate(), "minter state validation error: remainderToMint cannot be less than 0")

	minterState = types.MinterState{SequenceId: 1, AmountMinted: math.NewInt(123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousMinter: sdk.MustNewDecFromStr("231321.1234"), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{SequenceId: 1, AmountMinted: math.NewInt(123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousMinter: sdk.MustNewDecFromStr("-231321.1234"), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.EqualError(t, minterState.Validate(), "minter state validation error: remainderFromPreviousMinter cannot be less than 0")
}

func TestLinearMintingInfation(t *testing.T) {
	startTime := time.Now()
	duration := time.Hour * 24 * 365
	endTime := startTime.Add(duration)
	linearMinting := types.LinearMinting{Amount: math.NewInt(1000000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime, Config: config}

	inflation := minter1.CalculateInflation(math.NewInt(10000000), startTime, startTime.Add(-1000))
	require.EqualValues(t, sdk.ZeroDec(), inflation)

	inflation = minter1.CalculateInflation(math.NewInt(10000000), startTime, startTime)
	expected, _ := sdk.NewDecFromStr("0.1")
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 73
	endTime = startTime.Add(duration)
	minter1.EndTime = &endTime

	inflation = minter1.CalculateInflation(math.NewInt(10000000), startTime, startTime)
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 365 * 5
	endTime = startTime.Add(duration)
	minter1.EndTime = &endTime

	inflation = minter1.CalculateInflation(math.NewInt(10000000), startTime, startTime)
	expected, _ = sdk.NewDecFromStr("0.02")
	require.EqualValues(t, expected, inflation)
}

func TestNoMintingInfation(t *testing.T) {
	startTime := time.Now()
	duration := time.Hour * 24 * 365
	endTime := startTime.Add(duration)

	minter1 := types.Minter{SequenceId: 3, Config: testenv.NoMintingConfig}

	inflation := minter1.CalculateInflation(math.NewInt(10000000), startTime, startTime.Add(-1000))
	expected := sdk.ZeroDec()
	require.EqualValues(t, expected, inflation)

	inflation = minter1.CalculateInflation(math.NewInt(10000000), startTime, startTime)
	expected = sdk.ZeroDec()
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 73
	endTime = startTime.Add(duration)
	minter1.EndTime = &endTime

	inflation = minter1.CalculateInflation(math.NewInt(10000000), startTime, startTime)
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 365 * 5
	endTime = startTime.Add(duration)
	minter1.EndTime = &endTime

	inflation = minter1.CalculateInflation(math.NewInt(10000000), startTime, startTime)
	require.EqualValues(t, expected, inflation)
}

func TestUnlimitedExponentialStepMinting(t *testing.T) {
	exponentialStepMinting := types.ExponentialStepMinting{Amount: math.NewInt(160000000000000), StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}

	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	minter := types.Minter{SequenceId: 1, EndTime: nil, Config: config}

	amount := minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(Year/2))
	require.EqualValues(t, sdk.NewDec(20000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(Year))
	require.EqualValues(t, sdk.NewDec(40000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(2*Year))
	require.EqualValues(t, sdk.NewDec(80000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(3*Year))
	require.EqualValues(t, sdk.NewDec(120000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(4*Year))
	require.EqualValues(t, sdk.NewDec(160000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(5*Year))
	require.EqualValues(t, sdk.NewDec(180000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(6*Year))
	require.EqualValues(t, sdk.NewDec(200000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(7*Year))
	require.EqualValues(t, sdk.NewDec(220000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(8*Year))
	require.EqualValues(t, sdk.NewDec(240000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(9*Year))
	require.EqualValues(t, sdk.NewDec(250000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(10*Year))
	require.EqualValues(t, sdk.NewDec(260000000000000), amount)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(11*Year))
	require.EqualValues(t, sdk.NewDec(270000000000000), amount)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(12*Year))
	require.EqualValues(t, sdk.NewDec(280000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(13*Year))
	require.EqualValues(t, sdk.NewDec(285000000000000), amount)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(14*Year))
	require.EqualValues(t, sdk.NewDec(290000000000000), amount)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(15*Year))
	require.EqualValues(t, sdk.NewDec(295000000000000), amount)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(16*Year))
	require.EqualValues(t, sdk.NewDec(300000000000000), amount)

	beforeAmount := sdk.NewDec(300000000000000)
	amountToAdd := sdk.NewDec(10000000000000)
	expected := beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(20*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(24*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(28*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(32*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(36*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(40*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(44*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(48*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(52*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(56*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(60*Year))
	require.EqualValues(t, expected, amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(64*Year))
	require.EqualValues(t, sdk.NewDec(319995117187500), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(250*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("319999999999999.999947958295720691"), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(250*Year).Add(50*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("319999999999999.999999999999999996"), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(250*Year).Add(51*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("320000000000000.000000004235164732"), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(250*Year).Add(250*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("320000000000000.000000847032947246"), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(250*Year).Add(250*Year).Add(250*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("320000000000000.000001204782431465"), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("320000000000000.000001204782431465"), amount)
}

func TestLimitedExponentialStepMinting(t *testing.T) {
	exponentialStepMinting := types.ExponentialStepMinting{Amount: math.NewInt(160000000000000), StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(7 * Year)
	minter := types.Minter{SequenceId: 1, EndTime: &endTime, Config: config}

	amount := minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(Year/2))
	require.EqualValues(t, sdk.NewDec(20000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(time.Hour))
	require.EqualValues(t, sdk.MustNewDecFromStr("4566210045.662100456621004566"), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(Year))
	require.EqualValues(t, sdk.NewDec(40000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(2*Year))
	require.EqualValues(t, sdk.NewDec(80000000000000), amount)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(3*Year))
	require.EqualValues(t, sdk.NewDec(120000000000000), amount)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(4*Year))
	require.EqualValues(t, sdk.NewDec(160000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(5*Year))
	require.EqualValues(t, sdk.NewDec(180000000000000), amount)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(6*Year))
	require.EqualValues(t, sdk.NewDec(200000000000000), amount)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(7*Year))
	require.EqualValues(t, sdk.NewDec(220000000000000), amount)
	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(8*Year))
	require.EqualValues(t, sdk.NewDec(220000000000000), amount)

	amount = minter.AmountToMint(log.TestingLogger(), startTime, startTime.Add(16*Year))
	require.EqualValues(t, sdk.NewDec(220000000000000), amount)
}

func TestValidateExponentialStepMintingAmountBelowZero(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(PeriodDuration)

	exponentialStepMinting := types.ExponentialStepMinting{Amount: math.NewInt(-160000000000000), StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2}
	minters := []*types.Minter{&minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "minter with id 1 validation error: minter config validation error: amount cannot be less than 0")
}

func TestValidateMinterNoMinterConfigSet(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)

	minter1 := types.Minter{SequenceId: 1, Config: nil}
	minters := []*types.Minter{&minter1}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "minter with id 1 validation error: minter config is nil")
}

func TestValidateExponentialStepMinterLessThanZeror(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(PeriodDuration)

	exponentialStepMinting := types.ExponentialStepMinting{Amount: math.NewInt(140000000000000), StepDuration: -NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2}
	minters := []*types.Minter{&minter1, &minter2}

	params := types.Params{MintDenom: customDenom, StartTime: startTime, Minters: minters}
	require.EqualError(t, params.Validate(), "minter with id 1 validation error: minter config validation error: stepDuration must be bigger than 0")
}

func TestExponentialStepMintingInfationNotLimted(t *testing.T) {
	exponentialStepMinting := types.ExponentialStepMinting{Amount: math.NewInt(160000000000000), StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)

	minter := types.Minter{SequenceId: 1, EndTime: nil, Config: config}

	inflation := minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(-1000))
	require.EqualValues(t, sdk.ZeroDec(), inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime)
	expected, _ := sdk.NewDecFromStr("1")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(80000000000000), startTime, startTime.Add(Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(4*Year-1))
	expected, _ = sdk.NewDecFromStr("1")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(4*Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(6*Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(8*Year))
	expected, _ = sdk.NewDecFromStr("0.25")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(12*Year))
	expected, _ = sdk.NewDecFromStr("0.125")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(16*Year))
	expected, _ = sdk.NewDecFromStr("0.0625")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(20*Year))
	expected, _ = sdk.NewDecFromStr("0.03125")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(24*Year))
	expected, _ = sdk.NewDecFromStr("0.015625")
	require.EqualValues(t, expected, inflation)
}

func TestExponentialStepMintingInfationLimted(t *testing.T) {
	exponentialStepMinting := types.ExponentialStepMinting{Amount: math.NewInt(160000000000000), StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(10 * Year)
	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)

	minter := types.Minter{SequenceId: 1, EndTime: &endTime, Config: config}

	inflation := minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime)
	expected, _ := sdk.NewDecFromStr("1")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(80000000000000), startTime, startTime.Add(Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(4*Year-1))
	expected, _ = sdk.NewDecFromStr("1")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(4*Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(6*Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(8*Year))
	expected, _ = sdk.NewDecFromStr("0.25")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(12*Year))
	expected, _ = sdk.NewDecFromStr("0")
	require.EqualValues(t, expected, inflation)

	inflation = minter.CalculateInflation(math.NewInt(40000000000000), startTime, startTime.Add(24*Year))
	expected, _ = sdk.NewDecFromStr("0")
	require.EqualValues(t, expected, inflation)
}
