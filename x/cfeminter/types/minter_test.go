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

func TestTimeLinearMinter(t *testing.T) {
	minter := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	minterState := types.MinterState{Position: 1, AmountMinted: sdk.ZeroInt()}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(time.Duration(345600000000 * 1000000))
	blockTime := startTime.Add(time.Duration(345600000000 * 1000000 / 2))

	period := types.MintingPeriod{Position: 1, PeriodEnd: &endTime, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &minter}
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
	minterState := types.MinterState{Position: 1, AmountMinted: sdk.ZeroInt()}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(time.Duration(345600000000 * 1000000))
	blockTime := startTime.Add(time.Duration(345600000000 * 1000000 / 2))

	period := types.MintingPeriod{Position: 1, PeriodEnd: &endTime, Type: types.NO_MINTING}
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

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2, &period3}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.NoError(t, minter.Validate())

}

func TestValidateMinterPariodsOrderInitialyNotOrdered(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.NoError(t, minter.Validate())

}

func TestValidateMinterPariodsOrderInitialyNotFromOne(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 5, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 6, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 7, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.NoError(t, minter.Validate())

}

func TestValidateMinterPariodsOrderWrongFirstId(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 0, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "first period ordering id must be bigger than 0, but is 0")

}

func TestValidateMinterPariodsOrderWrongNotIncrementByOne(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 3, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 4, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "missing period with ordering id 2")

}

func TestValidateMinterNoPeriods(t *testing.T) {
	startTime := time.Now()

	minter := types.Minter{Start: startTime}
	require.EqualError(t, minter.Validate(), "no minter periods defined")

}

func TestValidateMinterLastPeriodWithEndDate(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	periods := []*types.MintingPeriod{&period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "last period cannot have PeriodEnd set, but is set to 2043-12-30 00:00:00 +0000 UTC")

}

func TestValidateMinterLastPeriodWithEndDateOnePeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}

	periods := []*types.MintingPeriod{&period1}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "last period cannot have PeriodEnd set, but is set to 2033-01-16 00:00:00 +0000 UTC")

}

func TestValidateMinterFirstPeriodWrongEnd(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime
	endTime2 := endTime1.Add(2 * time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 4, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "first period end must be bigger than minter start")

}

func TestValidateMinterNextPeriodWrongEnd(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 4, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period with Id 2 mast have PeriodEnd bigger than period with id 1")

}

func TestValidateMinterNoMintigTypeWithTimeLinearMinter(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING, TimeLinearMinter: &linearMinter2}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 3 - for NO_MINTING type (0) TimeLinearMinter must not be set")

}

func TestValidateMinterTimeLineraMinterTypeWithNoTimeLinearMinterDefinition(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 2 - for TIME_LINEAR_MINTER type (1) TimeLinearMinter must be set")

}

func TestValidateMinterTimeLineraMinterTypeWithNoPeriodEndInNotLastPeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: nil, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "only last period can have PeriodEnd empty")

}

func TestValidateMinterTimeLineraMinterTypeWithNoPeriodEnd(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: nil, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}

	periods := []*types.MintingPeriod{&period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 2 - for TIME_LINEAR_MINTER type (1) PeriodEnd must be set")

}

func TestValidateMinterUnknownType(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: "Unknown"}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 2 - unknow minting period type: Unknown")

}

func TestValidateMinterTimeLinearAmountLessThanZero(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(-100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 2 - TimeLinearMinter amount cannot be less than 0")

}

func TestCointainsIdTrue(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.True(t, minter.ContainsId(3))

}

func TestCointainsIdFalse(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.False(t, minter.ContainsId(6))

}

func TestValidateMinterState(t *testing.T) {

	minterState := types.MinterState{Position: 1, AmountMinted: sdk.ZeroInt(), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousPeriod: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{Position: 1, AmountMinted: sdk.NewInt(123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousPeriod: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{Position: 1, AmountMinted: sdk.NewInt(-123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousPeriod: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.EqualError(t, minterState.Validate(), "minter state amount cannot be less than 0")

	minterState = types.MinterState{Position: 1, AmountMinted: sdk.NewInt(123), RemainderToMint: sdk.MustNewDecFromStr("231321.1234"), RemainderFromPreviousPeriod: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{Position: 1, AmountMinted: sdk.NewInt(123), RemainderToMint: sdk.MustNewDecFromStr("-231321.1234"), RemainderFromPreviousPeriod: sdk.ZeroDec(), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.EqualError(t, minterState.Validate(), "minter remainder to mint amount cannot be less than 0")

	minterState = types.MinterState{Position: 1, AmountMinted: sdk.NewInt(123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("231321.1234"), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{Position: 1, AmountMinted: sdk.NewInt(123), RemainderToMint: sdk.ZeroDec(), RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("-231321.1234"), LastMintBlockTime: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}
	require.EqualError(t, minterState.Validate(), "minter remainder from previous period amount cannot be less than 0")
}

func TestTimeLinearMinterInfation(t *testing.T) {
	startTime := time.Now()
	duration := time.Hour * 24 * 365
	endTime := startTime.Add(duration)
	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}

	inflation := period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime.Add(-1000))
	require.EqualValues(t, sdk.ZeroDec(), inflation)

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	expected, _ := sdk.NewDecFromStr("0.1")
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 73
	endTime = startTime.Add(duration)
	period1.PeriodEnd = &endTime

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	expected, _ = sdk.NewDecFromStr("0.5")
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 365 * 5
	endTime = startTime.Add(duration)
	period1.PeriodEnd = &endTime

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	expected, _ = sdk.NewDecFromStr("0.02")
	require.EqualValues(t, expected, inflation)
}

func TestNoMintingInfation(t *testing.T) {
	startTime := time.Now()
	duration := time.Hour * 24 * 365
	endTime := startTime.Add(duration)

	period1 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}

	inflation := period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime.Add(-1000))
	expected := sdk.ZeroDec()
	require.EqualValues(t, expected, inflation)

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	expected = sdk.ZeroDec()
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 73
	endTime = startTime.Add(duration)
	period1.PeriodEnd = &endTime

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	require.EqualValues(t, expected, inflation)

	duration = time.Hour * 24 * 365 * 5
	endTime = startTime.Add(duration)
	period1.PeriodEnd = &endTime

	inflation = period1.CalculateInfation(sdk.NewInt(10000000), startTime, startTime)
	require.EqualValues(t, expected, inflation)
}

func TestUnlimitedPeriodicReductionMinter(t *testing.T) {
	minter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(40000000000000), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}
	minterState := types.MinterState{Position: 1, AmountMinted: sdk.ZeroInt()}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)

	period := types.MintingPeriod{Position: 1, PeriodEnd: nil, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &minter}

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

func TestLimitedPeriodicReductionMinter(t *testing.T) {
	minter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(40000000000000), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}
	minterState := types.MinterState{Position: 1, AmountMinted: sdk.ZeroInt()}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(7 * Year)
	period := types.MintingPeriod{Position: 1, PeriodEnd: &endTime, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &minter}

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

func TestValidatePeriodicReductionMinterMinterNotSet(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	pminter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(40000000000000), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &pminter}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.PERIODIC_REDUCTION_MINTER}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 2 - for PERIODIC_REDUCTION_MINTER type (1) PeriodicReductionMinter must be set")

}

func TestValidatePeriodicReductionMinterAmountBelowZero(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	pminter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(-40000000000000), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &pminter}

	period2 := types.MintingPeriod{Position: 2, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 1 - PeriodicReductionMinter MintAmount cannot be less than 0")

}

func TestValidatePeriodicReductionMinterPeriodLessThanZeror(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	pminter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(40000000000000), MintPeriod: -SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &pminter}

	period2 := types.MintingPeriod{Position: 2, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 1 - PeriodicReductionMinter MintPeriod must be bigger than 0")

}

func TestValidatePeriodicReductionMinterLengthLessThanZeror(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))

	pminter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(40000000000000), MintPeriod: SecondsInYear, ReductionPeriodLength: -4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &pminter}

	period2 := types.MintingPeriod{Position: 2, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 1 - PeriodicReductionMinter ReductionPeriodLength must be bigger than 0")

}

func TestPeriodicReductionMinterInfationNotLimted(t *testing.T) {

	minter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(40000000000000), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	period := types.MintingPeriod{Position: 1, PeriodEnd: nil, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &minter}

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

func TestPeriodicReductionMinterInfationLimted(t *testing.T) {

	minter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(40000000000000), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(10 * Year)
	period := types.MintingPeriod{Position: 1, PeriodEnd: &endTime, Type: types.PERIODIC_REDUCTION_MINTER, PeriodicReductionMinter: &minter}

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
