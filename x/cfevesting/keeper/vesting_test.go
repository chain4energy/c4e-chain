package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
)

func TestCalculateWithdrawable(t *testing.T) {
	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := sdk.NewInt(1000000)

	vesting := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       sdk.ZeroInt(),
		Sent:            sdk.ZeroInt(),
	}

	// current block less than lock start - witdrawable 0
	withdrawable := keeper.CalculateWithdrawable(start.Add(testutils.CreateDurationFromNumOfHours(-100)), vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to vesting start - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(start, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to lock end - witdrawable all
	withdrawable = keeper.CalculateWithdrawable(lockEnd, vesting)
	require.EqualValues(t, amount, withdrawable)

	// current block higher than lock start  but lass than lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEnd.Add(-1), vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

}

func TestCalculateWithdrawableAfterSendSendingSideBeforeLockEnd(t *testing.T) {
	startHeight := testutils.CreateTimeFromNumOfHours(1000)
	lockEndHeight := testutils.CreateTimeFromNumOfHours(10000)
	amount := sdk.NewInt(1000000)
	withdrawn := sdk.NewInt(500000)

	vesting := types.VestingPool{
		VestingType:     "test",
		LockStart:       startHeight.Add(testutils.CreateDurationFromNumOfHours(-300)),
		LockEnd:         lockEndHeight,
		InitiallyLocked: amount.AddRaw(50000),
		Withdrawn:       withdrawn,
		Sent:            sdk.NewInt(50000),
	}

	// current block less than lock start - witdrawable 0
	withdrawable := keeper.CalculateWithdrawable(startHeight.Add(testutils.CreateDurationFromNumOfHours(-100)), vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to lock start - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(startHeight, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight, vesting)
	require.EqualValues(t, amount.Sub(withdrawn), withdrawable)

	// current block higher than lock start  but lass than lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(-1), vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

}

func TestCalculateWithdrawableAfterSendSendingSideAfterLockEnd(t *testing.T) {
	startHeight := testutils.CreateTimeFromNumOfHours(10000)
	lockEndHeight := startHeight.Add(testutils.CreateDurationFromNumOfHours(-100))
	amount := sdk.NewInt(1000000)
	withdrawn := sdk.NewInt(500000)
	vesting := types.VestingPool{
		VestingType:     "test",
		LockStart:       lockEndHeight.Add(testutils.CreateDurationFromNumOfHours(-300)),
		LockEnd:         lockEndHeight,
		InitiallyLocked: amount.AddRaw(50000),
		Withdrawn:       withdrawn,
		Sent:            sdk.NewInt(50000),
	}

	// current block less than lock start - witdrawable 0
	withdrawable := keeper.CalculateWithdrawable(startHeight.Add(testutils.CreateDurationFromNumOfHours(-100)), vesting)
	require.EqualValues(t, amount.Sub(withdrawn), withdrawable)

	// current block equal to vesting start - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(startHeight, vesting)
	require.EqualValues(t, amount.Sub(withdrawn), withdrawable)

	// current block equal to lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight, vesting)
	require.EqualValues(t, amount.Sub(withdrawn), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(-1), vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

}
