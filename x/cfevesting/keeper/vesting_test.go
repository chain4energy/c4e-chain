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
	// end := testutils.CreateTimeFromNumOfHours(110000)
	// unlockingPeriod := testutils.CreateDurationFromNumOfHours(10)
	amount := sdk.NewInt(1000000)
	// withdrawn := sdk.NewInt(500)
	// notDivisibleAmount := sdk.NewInt(1000)

	vesting := types.VestingPool{
		Id:                        1,
		VestingType:               "test",
		LockStart:                 start,
		LockEnd:                   lockEnd,
		Vested:                    amount,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:          start,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
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

// func TestCalculateWithdrawableAfterSendReceivingSide(t *testing.T) {
// 	// vesting start before lock end - no additional tests. It is the same as tests in TestCalculateWithdrawable

// 	// vesting start after lock end
// 	startHeight := testutils.CreateTimeFromNumOfHours(10000)
// 	lockEndHeight := startHeight.Add(testutils.CreateDurationFromNumOfHours(-1000))
// 	endHeight := testutils.CreateTimeFromNumOfHours(110000)
// 	unlockingPeriod := testutils.CreateDurationFromNumOfHours(10)
// 	amount := sdk.NewInt(1000000)
// 	withdrawn := sdk.NewInt(500)
// 	notDivisibleAmount := sdk.NewInt(1000)

// 	vesting := types.Vesting{
// 		Id:                        1,
// 		VestingType:               "test",
// 		LockStart:                 startHeight,
// 		LockEnd:                   lockEndHeight,
// 		Vested:                    amount,
// 		Withdrawn:                 sdk.ZeroInt(),
// 		Sent:                      sdk.ZeroInt(),
// 		LastModification:          startHeight,
// 		LastModificationVested:    amount,
// 		LastModificationWithdrawn: sdk.ZeroInt(),
// 	}

// 	// current block less than vesting start - witdrawable 0
// 	withdrawable := keeper.CalculateWithdrawable(startHeight.Add(testutils.CreateDurationFromNumOfHours(-1)), vesting)
// 	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

// 	// current block equal to vesting start - witdrawable 0
// 	withdrawable = keeper.CalculateWithdrawable(startHeight, vesting)
// 	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

// 	// current block equal to lock end - witdrawable 0
// 	withdrawable = keeper.CalculateWithdrawable(lockEndHeight, vesting)
// 	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

// 	// current block higher than vesting start but lass than one unlocking period higher - witdrawable 0
// 	withdrawable = keeper.CalculateWithdrawable(startHeight.Add(unlockingPeriod).Add(testutils.CreateDurationFromNumOfHours(-1)), vesting)
// 	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

// 	// current block one unlocking period more than vesting start - witdrawable = vested/number of unlocking periods (1000000/((110000-10000)/10) = 100)
// 	withdrawable = keeper.CalculateWithdrawable(startHeight.Add(unlockingPeriod), vesting)
// 	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, startHeight.Add(unlockingPeriod), amount), withdrawable)

// 	// current block 10 unlocking periods more than vesting start - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) = 1000)
// 	withdrawable = keeper.CalculateWithdrawable(startHeight.Add(10*unlockingPeriod), vesting)
// 	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, startHeight.Add(10*unlockingPeriod), amount), withdrawable)

// 	// current block equal to vesting end - witdrawable all = 1000000
// 	withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
// 	require.EqualValues(t, amount, withdrawable)

// 	// current block equal to one block before vesting end - witdrawable all except one unlocking period = 9999*(1000000/((110000-10000)/10) = 999900)
// 	withdrawable = keeper.CalculateWithdrawable(endHeight.Add(testutils.CreateDurationFromNumOfHours(-1)), vesting)
// 	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, endHeight.Add(testutils.CreateDurationFromNumOfHours(-1)), amount), withdrawable)

// 	// setting VestingEndBlock to whole vesting periof not be divisible by unlocking period
// 	// vesting.VestingEnd = endHeight.Add(testutils.CreateDurationFromNumOfHours(-1))

// 	// current block 10 unlocking periods more than vesting start - witdrawable =  10*vested/number of unlocking periods 10*(1000000/(round-up[(109999-10000)/10]) = 1000)
// 	// where number of unlocking periods: (109999-10000)/10 = 9999.9 and round-up[(109999-10000)/10] = 10000
// 	withdrawable = keeper.CalculateWithdrawable(startHeight.Add(10*unlockingPeriod), vesting)
// 	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, startHeight.Add(10*unlockingPeriod), amount), withdrawable)

// 	// current block equal to vesting end - witdrawable all = 1000000
// 	// withdrawable = keeper.CalculateWithdrawable(vesting.VestingEnd, vesting)
// 	// require.EqualValues(t, amount, withdrawable)

// 	// vesting.VestingEnd = endHeight
// 	// setting LastModificationVested to prove that LastModificationVested is used in calculations (Vested field is just historical info)
// 	// and value set to number not divisible by number of periods - some more periods must pass to get some withdrawable amount
// 	vesting.LastModificationVested = notDivisibleAmount

// 	// current block one unlocking period more than vesting start - witdrawable = vested/number of unlocking periods (1000/((110000-10000)/10) = 0.1 this is less than 1 so 0 to wiithdraw)
// 	withdrawable = keeper.CalculateWithdrawable(startHeight.Add(unlockingPeriod), vesting)
// 	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

// 	// current block 9 unlocking period more than vesting start (still not divisible - one block before being divisible) -
// 	// witdrawable = vested/number of unlocking periods 9*(1000/((110000-10000)/10) = 0.9 this is less than 1 so 0 to wiithdraw)
// 	withdrawable = keeper.CalculateWithdrawable(startHeight.Add(10*unlockingPeriod).Add(testutils.CreateDurationFromNumOfHours(-1)), vesting)
// 	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

// 	// current block 10 unlocking period more than vesting start -
// 	// witdrawable = vested/number of unlocking periods 10*(1000/((110000-10000)/10) = 1)
// 	withdrawable = keeper.CalculateWithdrawable(startHeight.Add(10*unlockingPeriod), vesting)
// 	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, startHeight.Add(10*unlockingPeriod), vesting.LastModificationVested), withdrawable)

// 	vesting.LastModificationVested = amount
// 	// setting some coins withdrawn
// 	vesting.LastModificationWithdrawn = withdrawn

// 	// current block 10 unlocking periods more than vesting start - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) - 500 = 500)
// 	withdrawable = keeper.CalculateWithdrawable(startHeight.Add(10*unlockingPeriod), vesting)
// 	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, startHeight.Add(10*unlockingPeriod), vesting.LastModificationVested).Sub(withdrawn), withdrawable)

// 	// current block equal to vesting end - witdrawable all available = 1000000 - 500
// 	withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
// 	require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn), withdrawable)
// }

func TestCalculateWithdrawableAfterSendSendingSideBeforeLockEnd(t *testing.T) {
	startHeight := testutils.CreateTimeFromNumOfHours(1000)
	lockEndHeight := testutils.CreateTimeFromNumOfHours(10000)
	// endHeight := testutils.CreateTimeFromNumOfHours(110000)
	// unlockingPeriod := testutils.CreateDurationFromNumOfHours(10)
	amount := sdk.NewInt(1000000)
	// withdrawn := sdk.NewInt(500)
	// notDivisibleAmount := sdk.NewInt(1000)

	vesting := types.VestingPool{
		Id:                        1,
		VestingType:               "test",
		LockStart:                 startHeight.Add(testutils.CreateDurationFromNumOfHours(-300)),
		LockEnd:                   lockEndHeight,
		Vested:                    amount.AddRaw(50000),
		Withdrawn:                 sdk.NewInt(500000),
		Sent:                      sdk.NewInt(50000),
		LastModification:          startHeight,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	// current block less than lock start - witdrawable 0
	withdrawable := keeper.CalculateWithdrawable(startHeight.Add(testutils.CreateDurationFromNumOfHours(-100)), vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to lock start - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(startHeight, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight, vesting)
	require.EqualValues(t, amount, withdrawable)

	// current block higher than lock start  but lass than lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(-1), vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// // current block higher than lock end  but lass than one unlocking period higher - witdrawable 0
	// withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(unlockingPeriod).Add(testutils.CreateDurationFromNumOfHours(-1)), vesting)
	// require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// // current block one unlocking period more than lock end - witdrawable = vested/number of unlocking periods (1000000/((110000-10000)/10) = 100)
	// withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(unlockingPeriod), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, lockEndHeight.Add(unlockingPeriod), amount), withdrawable)

	// // current block 10 unlocking periods more than lock end - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) = 1000)
	// withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(10*unlockingPeriod), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, lockEndHeight.Add(10*unlockingPeriod), amount), withdrawable)

	// // current block equal to vesting end - witdrawable all = 1000000
	// withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	// require.EqualValues(t, amount, withdrawable)

	// // current block equal to one block before vesting end - witdrawable all except one unlocking period = 9999*(1000000/((110000-10000)/10) = 999900)
	// withdrawable = keeper.CalculateWithdrawable(endHeight.Add(testutils.CreateDurationFromNumOfHours(-1)), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, endHeight.Add(testutils.CreateDurationFromNumOfHours(-1)), amount), withdrawable)

	// // setting VestingEndBlock to whole vesting periof not be divisible by unlocking period
	// // vesting.VestingEnd = endHeight.Add(testutils.CreateDurationFromNumOfHours(-1))

	// // current block 10 unlocking periods more than lock end - witdrawable =  10*vested/number of unlocking periods 10*(1000000/(round-up[(109999-10000)/10]) = 1000)
	// // where number of unlocking periods: (109999-10000)/10 = 9999.9 and round-up[(109999-10000)/10] = 10000
	// withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(10*unlockingPeriod), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, lockEndHeight.Add(10*unlockingPeriod), amount), withdrawable)

	// // current block equal to vesting end - witdrawable all = 1000000
	// // withdrawable = keeper.CalculateWithdrawable(vesting.VestingEnd, vesting)
	// // require.EqualValues(t, amount, withdrawable)

	// // vesting.VestingEnd = endHeight
	// // setting LastModificationVested to prove that LastModificationVested is used in calculations (Vested field is just historical info)
	// // and value set to number not divisible by number of periods - some more periods must pass to get some withdrawable amount
	// vesting.LastModificationVested = notDivisibleAmount

	// // current block one unlocking period more than lock end - witdrawable = vested/number of unlocking periods (1000/((110000-10000)/10) = 0.1 this is less than 1 so 0 to wiithdraw)
	// withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(unlockingPeriod), vesting)
	// require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// // current block 9 unlocking period more than lock end (still not divisible - one block before being divisible) -
	// // witdrawable = vested/number of unlocking periods 9*(1000/((110000-10000)/10) = 0.9 this is less than 1 so 0 to wiithdraw)
	// withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(10*unlockingPeriod).Add(testutils.CreateDurationFromNumOfHours(-1)), vesting)
	// require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// // current block 10 unlocking period more than lock end -
	// // witdrawable = vested/number of unlocking periods 10*(1000/((110000-10000)/10) = 1)
	// withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(10*unlockingPeriod), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, lockEndHeight.Add(10*unlockingPeriod), vesting.LastModificationVested), withdrawable)

	// vesting.LastModificationVested = amount
	// // setting some coins withdrawn
	// vesting.LastModificationWithdrawn = withdrawn

	// // current block 10 unlocking periods more than lock end - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) - 500 = 500)
	// withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(10*unlockingPeriod), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, lockEndHeight.Add(10*unlockingPeriod), vesting.LastModificationVested).Sub(withdrawn), withdrawable)

	// // current block equal to vesting end - witdrawable all available = 1000000 - 500
	// withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	// require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn), withdrawable)
}

func TestCalculateWithdrawableAfterSendSendingSideAfterLockEnd(t *testing.T) {
	startHeight := testutils.CreateTimeFromNumOfHours(10000)
	lockEndHeight := startHeight.Add(testutils.CreateDurationFromNumOfHours(-100))
	// endHeight := testutils.CreateTimeFromNumOfHours(110000)
	// unlockingPeriod := testutils.CreateDurationFromNumOfHours(10)
	amount := sdk.NewInt(1000000)
	// withdrawn := sdk.NewInt(500)
	// notDivisibleAmount := sdk.NewInt(1000)

	vesting := types.VestingPool{
		Id:          1,
		VestingType: "test",
		LockStart:   lockEndHeight.Add(testutils.CreateDurationFromNumOfHours(-300)),
		LockEnd:     lockEndHeight,
		// VestingEnd:                endHeight,
		Vested: amount.AddRaw(50000),
		// ReleasePeriod:             unlockingPeriod,
		// DelegationAllowed:         true,
		Withdrawn:                 sdk.NewInt(500000),
		Sent:                      sdk.NewInt(50000),
		LastModification:          startHeight,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	// current block less than lock start - witdrawable 0
	withdrawable := keeper.CalculateWithdrawable(startHeight.Add(testutils.CreateDurationFromNumOfHours(-100)), vesting)
	require.EqualValues(t, amount, withdrawable)

	// current block equal to vesting start - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(startHeight, vesting)
	require.EqualValues(t, amount, withdrawable)

	// current block equal to lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight, vesting)
	require.EqualValues(t, amount, withdrawable)

	withdrawable = keeper.CalculateWithdrawable(lockEndHeight.Add(-1), vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// // current block higher than last modification height  but lass than one unlocking period higher - witdrawable 0
	// withdrawable = keeper.CalculateWithdrawable(startHeight.Add(unlockingPeriod).Add(testutils.CreateDurationFromNumOfHours(-1)), vesting)
	// require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// // current block one unlocking period more than last modification height - witdrawable = vested/number of unlocking periods (1000000/((110000-10000)/10) = 100)
	// withdrawable = keeper.CalculateWithdrawable(startHeight.Add(unlockingPeriod), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, startHeight.Add(unlockingPeriod), amount), withdrawable)

	// // current block 10 unlocking periods more than last modification height - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) = 1000)
	// withdrawable = keeper.CalculateWithdrawable(startHeight.Add(10*unlockingPeriod), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, startHeight.Add(10*unlockingPeriod), amount), withdrawable)

	// // current block equal to vesting end - witdrawable all = 1000000
	// withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	// require.EqualValues(t, amount, withdrawable)

	// // current block equal to one block before vesting end - witdrawable all except one unlocking period = 9999*(1000000/((110000-10000)/10) = 999900)
	// withdrawable = keeper.CalculateWithdrawable(endHeight.Add(testutils.CreateDurationFromNumOfHours(-1)), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, endHeight.Add(testutils.CreateDurationFromNumOfHours(-1)), amount), withdrawable)

	// // setting VestingEndBlock to whole vesting periof not be divisible by unlocking period
	// // vesting.VestingEnd = endHeight.Add(testutils.CreateDurationFromNumOfHours(-1))

	// // current block 10 unlocking periods more than last modification height- witdrawable =  10*vested/number of unlocking periods 10*(1000000/(round-up[(109999-10000)/10]) = 1000)
	// // where number of unlocking periods: (109999-10000)/10 = 9999.9 and round-up[(109999-10000)/10] = 10000
	// withdrawable = keeper.CalculateWithdrawable(startHeight.Add(10*unlockingPeriod), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, startHeight.Add(10*unlockingPeriod), amount), withdrawable)

	// // current block equal to vesting end - witdrawable all = 1000000
	// // withdrawable = keeper.CalculateWithdrawable(vesting.VestingEnd, vesting)
	// // require.EqualValues(t, amount, withdrawable)

	// // vesting.VestingEnd = endHeight
	// // setting LastModificationVested to prove that LastModificationVested is used in calculations (Vested field is just historical info)
	// // and value set to number not divisible by number of periods - some more periods must pass to get some withdrawable amount
	// vesting.LastModificationVested = notDivisibleAmount

	// // current block one unlocking period more than last modification height - witdrawable = vested/number of unlocking periods (1000/((110000-10000)/10) = 0.1 this is less than 1 so 0 to wiithdraw)
	// withdrawable = keeper.CalculateWithdrawable(startHeight.Add(unlockingPeriod), vesting)
	// require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// // current block 9 unlocking period more than last modification height (still not divisible - one block before being divisible) -
	// // witdrawable = vested/number of unlocking periods 9*(1000/((110000-10000)/10) = 0.9 this is less than 1 so 0 to wiithdraw)
	// withdrawable = keeper.CalculateWithdrawable(startHeight.Add(10*unlockingPeriod).Add(testutils.CreateDurationFromNumOfHours(-1)), vesting)
	// require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// // current block 10 unlocking period more than last modification height -
	// // witdrawable = vested/number of unlocking periods 10*(1000/((110000-10000)/10) = 1)
	// withdrawable = keeper.CalculateWithdrawable(startHeight.Add(10*unlockingPeriod), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, startHeight.Add(10*unlockingPeriod), vesting.LastModificationVested), withdrawable)

	// vesting.LastModificationVested = amount
	// // setting some coins withdrawn
	// vesting.LastModificationWithdrawn = withdrawn

	// // current block 10 unlocking periods more than last modification height - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) - 500 = 500)
	// withdrawable = keeper.CalculateWithdrawable(startHeight.Add(10*unlockingPeriod), vesting)
	// require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, startHeight.Add(10*unlockingPeriod), vesting.LastModificationVested).Sub(withdrawn), withdrawable)

	// // current block equal to vesting end - witdrawable all available = 1000000 - 500
	// withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	// require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn), withdrawable)
}
