package types_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
)

const PeriodDuration = time.Duration(345600000000 * 1000000)
const SecondsInYear = int32(3600 * 24 * 365)
const Year = time.Hour * 24 * 365

func TestLinearMinting(t *testing.T) {
	minter := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	minterState := types.MinterState{SequenceId: 1, AmountMinted: sdk.ZeroInt()}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(time.Duration(345600000000 * 1000000))
	blockTime := startTime.Add(time.Duration(345600000000 * 1000000 / 2))

	period := types.Minter{SequenceId: 1, EndTime: &endTime, Type: types.TIME_LINEAR_MINTER, LinearMinting: &minter}
	amount := period.AmountToMint(log.TestingLogger(), &minterState, startTime, blockTime)
	require.EqualValues(t, sdk.NewDec(500000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, endTime)
	require.EqualValues(t, sdk.NewDec(1000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, endTime.Add(time.Duration(10*1000000)))
	require.EqualValues(t, sdk.NewDec(1000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(time.Duration(345600000000*1000000*3/4)))
	require.EqualValues(t, sdk.NewDec(750000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(time.Duration(345600000000*1000000/4)))
	require.EqualValues(t, sdk.NewDec(250000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime)
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(time.Duration(-10*1000000)))
	require.EqualValues(t, sdk.NewDec(0), amount)

}

func TestNoMinting(t *testing.T) {
	minterState := types.MinterState{SequenceId: 1, AmountMinted: sdk.ZeroInt()}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(time.Duration(345600000000 * 1000000))
	blockTime := startTime.Add(time.Duration(345600000000 * 1000000 / 2))

	period := types.Minter{SequenceId: 1, EndTime: &endTime, Type: types.NO_MINTING}
	amount := period.AmountToMint(log.TestingLogger(), &minterState, startTime, blockTime)
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, endTime)
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, endTime.Add(time.Duration(10*1000000)))
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(time.Duration(345600000000*1000000*3/4)))
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(time.Duration(345600000000*1000000/4)))
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime)
	require.EqualValues(t, sdk.NewDec(0), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(time.Duration(-10*1000000)))
	require.EqualValues(t, sdk.NewDec(0), amount)
}

func TestValidateMinterPariodsOrder(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period1, &period2, &period3}
	params := types.Params{Minters: Minters}
	require.NoError(t, params.Validate())

}

func TestValidateMinterPariodsOrderInitialyNotOrdered(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.NoError(t, params.Validate())

}

func TestValidateMinterPariodsOrderInitialyNotFromOne(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 5, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 6, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 7, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.NoError(t, params.Validate())

}

func TestValidateMinterPariodsOrderWrongFirstId(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 0, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "first period ordering id must be bigger than 0, but is 0")

}

func TestValidateMinterPariodsOrderWrongNotIncrementByOne(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 3, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 4, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "missing period with ordering id 2")

}

func TestValidateMinterNoMinters(t *testing.T) {
	startTime := time.Now()

	params := types.Params{StartTime: startTime}
	require.EqualError(t, params.Validate(), "no minter Minters defined")

}

func TestValidateMinterLastPeriodWithEndDate(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	Minters := []*types.Minter{&period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "last period cannot have EndTime set, but is set to 2043-12-30 00:00:00 +0000 UTC")

}

func TestValidateMinterLastPeriodWithEndDateOnePeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}

	Minters := []*types.Minter{&period1}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "last period cannot have EndTime set, but is set to 2033-01-16 00:00:00 +0000 UTC")

}

func TestValidateMinterFirstPeriodWrongEnd(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime
	endTime2 := endTime1.Add(2 * time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 4, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "first period end must be bigger than minter start")

}

func TestValidateMinterNextPeriodWrongEnd(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 4, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "period with Id 2 mast have EndTime bigger than period with id 1")

}

func TestValidateMinterNoMintigTypeWithLinearMinting(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING, LinearMinting: &LinearMinting2}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "period id: 3 - for NO_MINTING type (0) LinearMinting must not be set")

}

func TestValidateMinterTimeLineraMinterTypeWithNoLinearMintingDefinition(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "period id: 2 - for TIME_LINEAR_MINTER type (1) LinearMinting must be set")

}

func TestValidateMinterTimeLineraMinterTypeWithNoEndTimeInNotLastPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: nil, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "only last period can have EndTime empty")

}

func TestValidateMinterTimeLineraMinterTypeWithNoEndTime(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: nil, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}

	Minters := []*types.Minter{&period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "period id: 2 - for TIME_LINEAR_MINTER type (1) EndTime must be set")

}

func TestValidateMinterUnknownType(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: "Unknown"}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "period id: 2 - unknow minting period type: Unknown")

}

func TestValidateMinterTimeLinearAmountLessThanZero(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(-100000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "period id: 2 - LinearMinting amount cannot be less than 0")

}

func TestCointainsIdTrue(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.True(t, params.ContainsId(3))

}

func TestCointainsIdFalse(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting2}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.False(t, params.ContainsId(6))

}

func TestValidateMinterState(t *testing.T) {

	minterState := types.MinterState{SequenceId: 1, AmountMinted: sdk.ZeroInt(), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousPeriod: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{SequenceId: 1, AmountMinted: sdk.NewInt(123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousPeriod: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{SequenceId: 1, AmountMinted: sdk.NewInt(-123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousPeriod: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.EqualError(t, minterState.Validate(), "minter state amount cannot be less than 0")

	minterState = types.MinterState{SequenceId: 1, AmountMinted: sdk.NewInt(123), RemainderToMint: sdk.MustNewDecFromStr("231321.1234"), RemainderFromPreviousPeriod: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{SequenceId: 1, AmountMinted: sdk.NewInt(123), RemainderToMint: sdk.MustNewDecFromStr("-231321.1234"), RemainderFromPreviousPeriod: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.EqualError(t, minterState.Validate(), "minter remainder to mint amount cannot be less than 0")

	minterState = types.MinterState{SequenceId: 1, AmountMinted: sdk.NewInt(123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("231321.1234"), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{SequenceId: 1, AmountMinted: sdk.NewInt(123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("-231321.1234"), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.EqualError(t, minterState.Validate(), "minter remainder from previous period amount cannot be less than 0")
}

func TestLinearMintingInfation(t *testing.T) {
	startTime := time.Now()
	duration := time.Hour * 24 * 365
	endTime := startTime.Add(duration)
	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime, Type: types.TIME_LINEAR_MINTER, LinearMinting: &LinearMinting1}

	inflation := period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime.Add(-1000))
	require.EqualValues(t, sdk.ZeroDec(), inflation)

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	expected, _ := sdk.NewDecFromStr("0.1")
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 73
	endTime = startTime.Add(duration)
	period1.EndTime = &endTime

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 365 * 5
	endTime = startTime.Add(duration)
	period1.EndTime = &endTime

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	expected, _ = sdk.NewDecFromStr("0.02")
	require.EqualValues(t, expected, inflation)
}

func TestNoMintingInfation(t *testing.T) {
	startTime := time.Now()
	duration := time.Hour * 24 * 365
	endTime := startTime.Add(duration)

	period1 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}

	inflation := period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime.Add(-1000))
	expected := sdk.ZeroDec()
	require.EqualValues(t, expected, inflation)

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	expected = sdk.ZeroDec()
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 73
	endTime = startTime.Add(duration)
	period1.EndTime = &endTime

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 365 * 5
	endTime = startTime.Add(duration)
	period1.EndTime = &endTime

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	require.EqualValues(t, expected, inflation)
}

func TestUnlimitedExponentialStepMinting(t *testing.T) {
	minter := types.ExponentialStepMinting{Amount: sdk.NewInt(40000000000000), StepDuration: time.Duration(SecondsInYear), AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	minterState := types.MinterState{SequenceId: 1, AmountMinted: sdk.ZeroInt()}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	period := types.Minter{SequenceId: 1, EndTime: nil, Type: types.PERIODIC_REDUCTION_MINTER, ExponentialStepMinting: &minter}

	amount := period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(Year/2))
	require.EqualValues(t, sdk.NewDec(20000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(Year))
	require.EqualValues(t, sdk.NewDec(40000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(2*Year))
	require.EqualValues(t, sdk.NewDec(80000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(3*Year))
	require.EqualValues(t, sdk.NewDec(120000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(4*Year))
	require.EqualValues(t, sdk.NewDec(160000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(5*Year))
	require.EqualValues(t, sdk.NewDec(180000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(6*Year))
	require.EqualValues(t, sdk.NewDec(200000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(7*Year))
	require.EqualValues(t, sdk.NewDec(220000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(8*Year))
	require.EqualValues(t, sdk.NewDec(240000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(9*Year))
	require.EqualValues(t, sdk.NewDec(250000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(10*Year))
	require.EqualValues(t, sdk.NewDec(260000000000000), amount)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(11*Year))
	require.EqualValues(t, sdk.NewDec(270000000000000), amount)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(12*Year))
	require.EqualValues(t, sdk.NewDec(280000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(13*Year))
	require.EqualValues(t, sdk.NewDec(285000000000000), amount)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(14*Year))
	require.EqualValues(t, sdk.NewDec(290000000000000), amount)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(15*Year))
	require.EqualValues(t, sdk.NewDec(295000000000000), amount)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(16*Year))
	require.EqualValues(t, sdk.NewDec(300000000000000), amount)

	beforeAmount := sdk.NewDec(300000000000000)
	amountToAdd := sdk.NewDec(10000000000000)
	expected := beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(20*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(24*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(28*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(32*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(36*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(40*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(44*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(48*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(52*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(56*Year))
	require.EqualValues(t, expected, amount)

	beforeAmount = expected
	amountToAdd = amountToAdd.QuoInt64(2)
	expected = beforeAmount.Add(amountToAdd)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(60*Year))
	require.EqualValues(t, expected, amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(64*Year))
	require.EqualValues(t, sdk.NewDec(319995117187500), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(250*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("319999999999999.999947958295720691"), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(250*Year).Add(50*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("319999999999999.999999999999999996"), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(250*Year).Add(51*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("320000000000000.000000004235164732"), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(250*Year).Add(250*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("320000000000000.000000847032947246"), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(250*Year).Add(250*Year).Add(250*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("320000000000000.000001204782431465"), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year).Add(250*Year))
	require.EqualValues(t, sdk.MustNewDecFromStr("320000000000000.000001204782431465"), amount)

}

func TestLimitedExponentialStepMinting(t *testing.T) {
	minter := types.ExponentialStepMinting{Amount: sdk.NewInt(40000000000000), StepDuration: time.Duration(SecondsInYear), AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	minterState := types.MinterState{SequenceId: 1, AmountMinted: sdk.ZeroInt()}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(7 * Year)
	period := types.Minter{SequenceId: 1, EndTime: &endTime, Type: types.PERIODIC_REDUCTION_MINTER, ExponentialStepMinting: &minter}

	amount := period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(Year/2))
	require.EqualValues(t, sdk.NewDec(20000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(time.Hour))
	require.EqualValues(t, sdk.MustNewDecFromStr("4566210045.662100456621004566"), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(Year))
	require.EqualValues(t, sdk.NewDec(40000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(2*Year))
	require.EqualValues(t, sdk.NewDec(80000000000000), amount)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(3*Year))
	require.EqualValues(t, sdk.NewDec(120000000000000), amount)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(4*Year))
	require.EqualValues(t, sdk.NewDec(160000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(5*Year))
	require.EqualValues(t, sdk.NewDec(180000000000000), amount)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(6*Year))
	require.EqualValues(t, sdk.NewDec(200000000000000), amount)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(7*Year))
	require.EqualValues(t, sdk.NewDec(220000000000000), amount)
	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(8*Year))
	require.EqualValues(t, sdk.NewDec(220000000000000), amount)

	amount = period.AmountToMint(log.TestingLogger(), &minterState, startTime, startTime.Add(16*Year))
	require.EqualValues(t, sdk.NewDec(220000000000000), amount)

}

func TestValidateExponentialStepMintingMinterNotSet(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	pminter := types.ExponentialStepMinting{Amount: sdk.NewInt(40000000000000), StepDuration: time.Duration(SecondsInYear), AmountMultiplier: sdk.MustNewDecFromStr("0.5")}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.PERIODIC_REDUCTION_MINTER, ExponentialStepMinting: &pminter}
	period2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.PERIODIC_REDUCTION_MINTER}

	period3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period3, &period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "period id: 2 - for PERIODIC_REDUCTION_MINTER type (1) ExponentialStepMinting must be set")

}

func TestValidateExponentialStepMintingAmountBelowZero(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	pminter := types.ExponentialStepMinting{Amount: sdk.NewInt(-40000000000000), StepDuration: time.Duration(SecondsInYear), AmountMultiplier: sdk.MustNewDecFromStr("0.5")}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.PERIODIC_REDUCTION_MINTER, ExponentialStepMinting: &pminter}

	period2 := types.Minter{SequenceId: 2, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "period id: 1 - ExponentialStepMinting Amount cannot be less than 0")

}

func TestValidateExponentialStepMinterLessThanZeror(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	pminter := types.ExponentialStepMinting{Amount: sdk.NewInt(40000000000000), StepDuration: time.Duration(-SecondsInYear), AmountMultiplier: sdk.MustNewDecFromStr("0.5")}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.PERIODIC_REDUCTION_MINTER, ExponentialStepMinting: &pminter}

	period2 := types.Minter{SequenceId: 2, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "period id: 1 - ExponentialStepMinting StepDuration must be bigger than 0")

}

func TestValidateExponentialStepMintingLengthLessThanZeror(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	pminter := types.ExponentialStepMinting{Amount: sdk.NewInt(40000000000000), StepDuration: time.Duration(SecondsInYear), AmountMultiplier: sdk.MustNewDecFromStr("0.5")}

	period1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.PERIODIC_REDUCTION_MINTER, ExponentialStepMinting: &pminter}

	period2 := types.Minter{SequenceId: 2, Type: types.NO_MINTING}
	Minters := []*types.Minter{&period1, &period2}
	params := types.Params{StartTime: startTime, Minters: Minters}
	require.EqualError(t, params.Validate(), "period id: 1 - ExponentialStepMinting AmountMultiplier must be bigger than 0")

}

func TestExponentialStepMintingInfationNotLimted(t *testing.T) {

	minter := types.ExponentialStepMinting{Amount: sdk.NewInt(40000000000000), StepDuration: time.Duration(SecondsInYear), AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	period := types.Minter{SequenceId: 1, EndTime: nil, Type: types.PERIODIC_REDUCTION_MINTER, ExponentialStepMinting: &minter}

	inflation := period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(-1000))
	require.EqualValues(t, sdk.ZeroDec(), inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime)
	expected, _ := sdk.NewDecFromStr("1")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(80000000000000), startTime, startTime.Add(Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(4*Year-1))
	expected, _ = sdk.NewDecFromStr("1")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(4*Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(6*Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(8*Year))
	expected, _ = sdk.NewDecFromStr("0.25")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(12*Year))
	expected, _ = sdk.NewDecFromStr("0.125")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(16*Year))
	expected, _ = sdk.NewDecFromStr("0.0625")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(20*Year))
	expected, _ = sdk.NewDecFromStr("0.03125")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(24*Year))
	expected, _ = sdk.NewDecFromStr("0.015625")
	require.EqualValues(t, expected, inflation)
}

func TestExponentialStepMintingInfationLimted(t *testing.T) {

	minter := types.ExponentialStepMinting{Amount: sdk.NewInt(40000000000000), StepDuration: time.Duration(SecondsInYear), AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(10 * Year)
	period := types.Minter{SequenceId: 1, EndTime: &endTime, Type: types.PERIODIC_REDUCTION_MINTER, ExponentialStepMinting: &minter}

	inflation := period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime)
	expected, _ := sdk.NewDecFromStr("1")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(80000000000000), startTime, startTime.Add(Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(4*Year-1))
	expected, _ = sdk.NewDecFromStr("1")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(4*Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(6*Year))
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(8*Year))
	expected, _ = sdk.NewDecFromStr("0.25")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(12*Year))
	expected, _ = sdk.NewDecFromStr("0")
	require.EqualValues(t, expected, inflation)

	inflation = period.CalculateInfation(sdk.NewInt(40000000000000), startTime, startTime.Add(24*Year))
	expected, _ = sdk.NewDecFromStr("0")
	require.EqualValues(t, expected, inflation)
}
