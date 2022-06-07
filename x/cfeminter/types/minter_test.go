package types_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

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

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(345600000000 * 1000000*3 / 4)))
	require.EqualValues(t, sdk.NewInt(750000), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(345600000000 * 1000000 / 4)))
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

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(345600000000 * 1000000*3 / 4)))
	require.EqualValues(t, sdk.NewInt(0), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(345600000000 * 1000000 / 4)))
	require.EqualValues(t, sdk.NewInt(0), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime)
	require.EqualValues(t, sdk.NewInt(0), amount)

	amount = period.AmountToMint(&minterState, startTime, startTime.Add(time.Duration(-10*1000000)))
	require.EqualValues(t, sdk.NewInt(0), amount)
}
