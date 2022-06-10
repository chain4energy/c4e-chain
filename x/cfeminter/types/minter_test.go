package types_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const PeriodDuration = time.Duration(345600000000 * 1000000)

func TestTimeLinearMinter(t *testing.T) {
	minter := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	minterState := types.MinterState{CurrentOrderingId: 1, AmountMinted: sdk.ZeroInt()}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(time.Duration(345600000000 * 1000000))
	blockTime := startTime.Add(time.Duration(345600000000 * 1000000 / 2))

	period := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &minter}
	amount := period.AmountToMint(&minterState, startTime, blockTime)
	require.EqualValues(t, sdk.NewInt(500000), amount)

	amount = period.AmountToMint(&minterState, startTime, endTime)
	require.EqualValues(t, sdk.NewInt(1000000), amount)

	amount = period.AmountToMint(&minterState, startTime, endTime.Add(time.Duration(10*1000000)))
	require.EqualValues(t, sdk.NewInt(1000000), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(345600000000*1000000*3/4)))
	require.EqualValues(t, sdk.NewInt(750000), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(345600000000*1000000/4)))
	require.EqualValues(t, sdk.NewInt(250000), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime)
	require.EqualValues(t, sdk.NewInt(0), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(-10*1000000)))
	require.EqualValues(t, sdk.NewInt(0), amount)
}

func TestNoMinting(t *testing.T) {
	minterState := types.MinterState{CurrentOrderingId: 1, AmountMinted: sdk.ZeroInt()}

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := startTime.Add(time.Duration(345600000000 * 1000000))
	blockTime := startTime.Add(time.Duration(345600000000 * 1000000 / 2))

	period := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime, Type: types.MintingPeriod_NO_MINTING}
	amount := period.AmountToMint(&minterState, startTime, blockTime)
	require.EqualValues(t, sdk.NewInt(0), amount)

	amount = period.AmountToMint(&minterState, startTime, endTime)
	require.EqualValues(t, sdk.NewInt(0), amount)

	amount = period.AmountToMint(&minterState, startTime, endTime.Add(time.Duration(10*1000000)))
	require.EqualValues(t, sdk.NewInt(0), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(345600000000*1000000*3/4)))
	require.EqualValues(t, sdk.NewInt(0), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(345600000000*1000000/4)))
	require.EqualValues(t, sdk.NewInt(0), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime)
	require.EqualValues(t, sdk.NewInt(0), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(-10*1000000)))
	require.EqualValues(t, sdk.NewInt(0), amount)
}

func TestValidateMinterPariodsOrder(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING}
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

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING}
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

	period1 := types.MintingPeriod{OrderingId: 5, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 6, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 7, Type: types.MintingPeriod_NO_MINTING}
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

	period1 := types.MintingPeriod{OrderingId: 0, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING}
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

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 3, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 4, Type: types.MintingPeriod_NO_MINTING}
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

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	// period3 := types.MintingPeriod{OrderingId: 4, Type: types.MintingPeriod_NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "last period cannot have PeriodEnd set, but is set to 2043-12-30 00:00:00 +0000 UTC")

}

func TestValidateMinterLastPeriodWithEndDateOnePeriod(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	// endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	// linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	// period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	// period3 := types.MintingPeriod{OrderingId: 4, Type: types.MintingPeriod_NO_MINTING}
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

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 4, Type: types.MintingPeriod_NO_MINTING}
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

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 4, Type: types.MintingPeriod_NO_MINTING}
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

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING, TimeLinearMinter: &linearMinter2}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 3 - for NO_MINTING type (0) TimeLinearMinter must not be set")

}

func TestValidateMinterTimeLineraMinterTypeWithNoTimeLinearMinterDefinition(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	// linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 2 - for MintingPeriod_TIME_LINEAR_MINTER type (1) TimeLinearMinter must be set")

}

func TestValidateMinterUnknownType(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	// linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: 6}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.EqualError(t, minter.Validate(), "period id: 2 - unknow minting period type: 6")

}

func TestValidateMinterTimeLinearAmountLessThanZero(t *testing.T) {
	startTime := time.Now()
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(-100000)}

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING}
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

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING}
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

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING}
	periods := []*types.MintingPeriod{&period3, &period1, &period2}
	minter := types.Minter{Start: startTime, Periods: periods}
	require.False(t, minter.ContainsId(6))

}

func TestValidateMinterState(t *testing.T) {

	minterState := types.MinterState{CurrentOrderingId: 1, AmountMinted: sdk.ZeroInt()}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{CurrentOrderingId: 1, AmountMinted: sdk.NewInt(123)}
	require.NoError(t, minterState.Validate())

	minterState = types.MinterState{CurrentOrderingId: 1, AmountMinted: sdk.NewInt(-123)}
	require.EqualError(t, minterState.Validate(), "minter state amount cannot be less than 0")
}
