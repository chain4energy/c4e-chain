package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
	"github.com/chain4energy/c4e-chain/x/cfevesting/internal/testutils"
)

func TestCalculateWithdrawable(t *testing.T) {
	startHeight := int64(1000)
	lockEndHeight := int64(10000)
	endHeight := int64(110000)
	unlockingPeriod := int64(10)
	amount := sdk.NewInt(1000000)
	withdrawn := sdk.NewInt(500)
	notDivisibleAmount := sdk.NewInt(1000)

	vesting := types.Vesting{
		Id:                        1,
		VestingType:               "test",
		VestingStartBlock:         startHeight,
		LockEndBlock:              lockEndHeight,
		VestingEndBlock:           endHeight,
		Vested:                    amount,
		FreeCoinsBlockPeriod:      unlockingPeriod,
		DelegationAllowed:         true,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModificationBlock:     startHeight,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	// current block less than vesting start - witdrawable 0
	withdrawable := keeper.CalculateWithdrawable(startHeight-100, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to vesting start - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(startHeight, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block higher than lock end  but lass than one unlocking period higher - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+unlockingPeriod-1, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block one unlocking period more than lock end - witdrawable = vested/number of unlocking periods (1000000/((110000-10000)/10) = 100)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, lockEndHeight+unlockingPeriod, amount), withdrawable)

	// current block 10 unlocking periods more than lock end - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) = 1000)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, lockEndHeight+10*unlockingPeriod, amount), withdrawable)

	// current block equal to vesting end - witdrawable all = 1000000
	withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	require.EqualValues(t, amount, withdrawable)

	// current block equal to one block before vesting end - witdrawable all except one unlocking period = 9999*(1000000/((110000-10000)/10) = 999900)
	withdrawable = keeper.CalculateWithdrawable(endHeight-1, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, endHeight-1, amount), withdrawable)

	// setting VestingEndBlock to whole vesting periof not be divisible by unlocking period
	vesting.VestingEndBlock = endHeight - 1

	// current block 10 unlocking periods more than lock end - witdrawable =  10*vested/number of unlocking periods 10*(1000000/(round-up[(109999-10000)/10]) = 1000)
	// where number of unlocking periods: (109999-10000)/10 = 9999.9 and round-up[(109999-10000)/10] = 10000
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, lockEndHeight+10*unlockingPeriod, amount), withdrawable)

	// current block equal to vesting end - witdrawable all = 1000000
	withdrawable = keeper.CalculateWithdrawable(vesting.VestingEndBlock, vesting)
	require.EqualValues(t, amount, withdrawable)

	vesting.VestingEndBlock = endHeight
	// setting LastModificationVested to prove that LastModificationVested is used in calculations (Vested field is just historical info)
	// and value set to number not divisible by number of periods - some more periods must pass to get some withdrawable amount
	vesting.LastModificationVested = notDivisibleAmount

	// current block one unlocking period more than lock end - witdrawable = vested/number of unlocking periods (1000/((110000-10000)/10) = 0.1 this is less than 1 so 0 to wiithdraw)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+unlockingPeriod, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block 9 unlocking period more than lock end (still not divisible - one block before being divisible) -
	// witdrawable = vested/number of unlocking periods 9*(1000/((110000-10000)/10) = 0.9 this is less than 1 so 0 to wiithdraw)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+10*unlockingPeriod-1, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block 10 unlocking period more than lock end -
	// witdrawable = vested/number of unlocking periods 10*(1000/((110000-10000)/10) = 1)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, lockEndHeight+10*unlockingPeriod, vesting.LastModificationVested), withdrawable)

	vesting.LastModificationVested = amount
	// setting some coins withdrawn
	vesting.LastModificationWithdrawn = withdrawn

	// current block 10 unlocking periods more than lock end - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) - 500 = 500)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, lockEndHeight+10*unlockingPeriod, vesting.LastModificationVested).Sub(withdrawn), withdrawable)

	// current block equal to vesting end - witdrawable all available = 1000000 - 500
	withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn), withdrawable)
}

func TestCalculateWithdrawableAfterSendReceivingSide(t *testing.T) {
	// vesting start before lock end - no additional tests. It is the same as tests in TestCalculateWithdrawable

	// vesting start after lock end
	startHeight := int64(10000)
	lockEndHeight := int64(startHeight - 1000)
	endHeight := int64(110000)
	unlockingPeriod := int64(10)
	amount := sdk.NewInt(1000000)
	withdrawn := sdk.NewInt(500)
	notDivisibleAmount := sdk.NewInt(1000)

	vesting := types.Vesting{
		Id:                        1,
		VestingType:               "test",
		VestingStartBlock:         startHeight,
		LockEndBlock:              lockEndHeight,
		VestingEndBlock:           endHeight,
		Vested:                    amount,
		FreeCoinsBlockPeriod:      unlockingPeriod,
		DelegationAllowed:         true,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModificationBlock:     startHeight,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	// current block less than vesting start - witdrawable 0
	withdrawable := keeper.CalculateWithdrawable(startHeight-1, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to vesting start - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(startHeight, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block higher than vesting start but lass than one unlocking period higher - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(startHeight+unlockingPeriod-1, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block one unlocking period more than vesting start - witdrawable = vested/number of unlocking periods (1000000/((110000-10000)/10) = 100)
	withdrawable = keeper.CalculateWithdrawable(startHeight+unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, startHeight+unlockingPeriod, amount), withdrawable)

	// current block 10 unlocking periods more than vesting start - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) = 1000)
	withdrawable = keeper.CalculateWithdrawable(startHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, startHeight+10*unlockingPeriod, amount), withdrawable)

	// current block equal to vesting end - witdrawable all = 1000000
	withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	require.EqualValues(t, amount, withdrawable)

	// current block equal to one block before vesting end - witdrawable all except one unlocking period = 9999*(1000000/((110000-10000)/10) = 999900)
	withdrawable = keeper.CalculateWithdrawable(endHeight-1, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, endHeight-1, amount), withdrawable)

	// setting VestingEndBlock to whole vesting periof not be divisible by unlocking period
	vesting.VestingEndBlock = endHeight - 1

	// current block 10 unlocking periods more than vesting start - witdrawable =  10*vested/number of unlocking periods 10*(1000000/(round-up[(109999-10000)/10]) = 1000)
	// where number of unlocking periods: (109999-10000)/10 = 9999.9 and round-up[(109999-10000)/10] = 10000
	withdrawable = keeper.CalculateWithdrawable(startHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, startHeight+10*unlockingPeriod, amount), withdrawable)

	// current block equal to vesting end - witdrawable all = 1000000
	withdrawable = keeper.CalculateWithdrawable(vesting.VestingEndBlock, vesting)
	require.EqualValues(t, amount, withdrawable)

	vesting.VestingEndBlock = endHeight
	// setting LastModificationVested to prove that LastModificationVested is used in calculations (Vested field is just historical info)
	// and value set to number not divisible by number of periods - some more periods must pass to get some withdrawable amount
	vesting.LastModificationVested = notDivisibleAmount

	// current block one unlocking period more than vesting start - witdrawable = vested/number of unlocking periods (1000/((110000-10000)/10) = 0.1 this is less than 1 so 0 to wiithdraw)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+unlockingPeriod, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block 9 unlocking period more than vesting start (still not divisible - one block before being divisible) -
	// witdrawable = vested/number of unlocking periods 9*(1000/((110000-10000)/10) = 0.9 this is less than 1 so 0 to wiithdraw)
	withdrawable = keeper.CalculateWithdrawable(startHeight+10*unlockingPeriod-1, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block 10 unlocking period more than vesting start -
	// witdrawable = vested/number of unlocking periods 10*(1000/((110000-10000)/10) = 1)
	withdrawable = keeper.CalculateWithdrawable(startHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, startHeight+10*unlockingPeriod, vesting.LastModificationVested), withdrawable)

	vesting.LastModificationVested = amount
	// setting some coins withdrawn
	vesting.LastModificationWithdrawn = withdrawn

	// current block 10 unlocking periods more than vesting start - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) - 500 = 500)
	withdrawable = keeper.CalculateWithdrawable(startHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, startHeight+10*unlockingPeriod, vesting.LastModificationVested).Sub(withdrawn), withdrawable)

	// current block equal to vesting end - witdrawable all available = 1000000 - 500
	withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn), withdrawable)
}

func TestCalculateWithdrawableAfterSendSendingSideBeforeLockEnd(t *testing.T) {
	startHeight := int64(1000)
	lockEndHeight := int64(10000)
	endHeight := int64(110000)
	unlockingPeriod := int64(10)
	amount := sdk.NewInt(1000000)
	withdrawn := sdk.NewInt(500)
	notDivisibleAmount := sdk.NewInt(1000)

	vesting := types.Vesting{
		Id:                        1,
		VestingType:               "test",
		VestingStartBlock:         startHeight - 300,
		LockEndBlock:              lockEndHeight,
		VestingEndBlock:           endHeight,
		Vested:                    amount.AddRaw(50000),
		FreeCoinsBlockPeriod:      unlockingPeriod,
		DelegationAllowed:         true,
		Withdrawn:                 sdk.NewInt(500000),
		Sent:                      sdk.NewInt(50000),
		LastModificationBlock:     startHeight,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	// current block less than vesting start - witdrawable 0
	withdrawable := keeper.CalculateWithdrawable(startHeight-100, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to vesting start - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(startHeight, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block higher than lock end  but lass than one unlocking period higher - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+unlockingPeriod-1, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block one unlocking period more than lock end - witdrawable = vested/number of unlocking periods (1000000/((110000-10000)/10) = 100)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, lockEndHeight+unlockingPeriod, amount), withdrawable)

	// current block 10 unlocking periods more than lock end - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) = 1000)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, lockEndHeight+10*unlockingPeriod, amount), withdrawable)

	// current block equal to vesting end - witdrawable all = 1000000
	withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	require.EqualValues(t, amount, withdrawable)

	// current block equal to one block before vesting end - witdrawable all except one unlocking period = 9999*(1000000/((110000-10000)/10) = 999900)
	withdrawable = keeper.CalculateWithdrawable(endHeight-1, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, endHeight-1, amount), withdrawable)

	// setting VestingEndBlock to whole vesting periof not be divisible by unlocking period
	vesting.VestingEndBlock = endHeight - 1

	// current block 10 unlocking periods more than lock end - witdrawable =  10*vested/number of unlocking periods 10*(1000000/(round-up[(109999-10000)/10]) = 1000)
	// where number of unlocking periods: (109999-10000)/10 = 9999.9 and round-up[(109999-10000)/10] = 10000
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, lockEndHeight+10*unlockingPeriod, amount), withdrawable)

	// current block equal to vesting end - witdrawable all = 1000000
	withdrawable = keeper.CalculateWithdrawable(vesting.VestingEndBlock, vesting)
	require.EqualValues(t, amount, withdrawable)

	vesting.VestingEndBlock = endHeight
	// setting LastModificationVested to prove that LastModificationVested is used in calculations (Vested field is just historical info)
	// and value set to number not divisible by number of periods - some more periods must pass to get some withdrawable amount
	vesting.LastModificationVested = notDivisibleAmount

	// current block one unlocking period more than lock end - witdrawable = vested/number of unlocking periods (1000/((110000-10000)/10) = 0.1 this is less than 1 so 0 to wiithdraw)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+unlockingPeriod, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block 9 unlocking period more than lock end (still not divisible - one block before being divisible) -
	// witdrawable = vested/number of unlocking periods 9*(1000/((110000-10000)/10) = 0.9 this is less than 1 so 0 to wiithdraw)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+10*unlockingPeriod-1, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block 10 unlocking period more than lock end -
	// witdrawable = vested/number of unlocking periods 10*(1000/((110000-10000)/10) = 1)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, lockEndHeight+10*unlockingPeriod, vesting.LastModificationVested), withdrawable)

	vesting.LastModificationVested = amount
	// setting some coins withdrawn
	vesting.LastModificationWithdrawn = withdrawn

	// current block 10 unlocking periods more than lock end - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) - 500 = 500)
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, lockEndHeight+10*unlockingPeriod, vesting.LastModificationVested).Sub(withdrawn), withdrawable)

	// current block equal to vesting end - witdrawable all available = 1000000 - 500
	withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn), withdrawable)
}

func TestCalculateWithdrawableAfterSendSendingSideAfterLockEnd(t *testing.T) {
	startHeight := int64(10000)
	lockEndHeight := startHeight - 100
	endHeight := int64(110000)
	unlockingPeriod := int64(10)
	amount := sdk.NewInt(1000000)
	withdrawn := sdk.NewInt(500)
	notDivisibleAmount := sdk.NewInt(1000)

	vesting := types.Vesting{
		Id:                        1,
		VestingType:               "test",
		VestingStartBlock:         lockEndHeight - 300,
		LockEndBlock:              lockEndHeight,
		VestingEndBlock:           endHeight,
		Vested:                    amount.AddRaw(50000),
		FreeCoinsBlockPeriod:      unlockingPeriod,
		DelegationAllowed:         true,
		Withdrawn:                 sdk.NewInt(500000),
		Sent:                      sdk.NewInt(50000),
		LastModificationBlock:     startHeight,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	// current block less than vesting start - witdrawable 0
	withdrawable := keeper.CalculateWithdrawable(startHeight-100, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to vesting start - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(startHeight, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block equal to lock end - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(lockEndHeight, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block higher than last modification height  but lass than one unlocking period higher - witdrawable 0
	withdrawable = keeper.CalculateWithdrawable(startHeight+unlockingPeriod-1, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block one unlocking period more than last modification height - witdrawable = vested/number of unlocking periods (1000000/((110000-10000)/10) = 100)
	withdrawable = keeper.CalculateWithdrawable(startHeight+unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, startHeight+unlockingPeriod, amount), withdrawable)

	// current block 10 unlocking periods more than last modification height - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) = 1000)
	withdrawable = keeper.CalculateWithdrawable(startHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, startHeight+10*unlockingPeriod, amount), withdrawable)

	// current block equal to vesting end - witdrawable all = 1000000
	withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	require.EqualValues(t, amount, withdrawable)

	// current block equal to one block before vesting end - witdrawable all except one unlocking period = 9999*(1000000/((110000-10000)/10) = 999900)
	withdrawable = keeper.CalculateWithdrawable(endHeight-1, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(lockEndHeight, endHeight, unlockingPeriod, endHeight-1, amount), withdrawable)

	// setting VestingEndBlock to whole vesting periof not be divisible by unlocking period
	vesting.VestingEndBlock = endHeight - 1

	// current block 10 unlocking periods more than last modification height- witdrawable =  10*vested/number of unlocking periods 10*(1000000/(round-up[(109999-10000)/10]) = 1000)
	// where number of unlocking periods: (109999-10000)/10 = 9999.9 and round-up[(109999-10000)/10] = 10000
	withdrawable = keeper.CalculateWithdrawable(startHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, startHeight+10*unlockingPeriod, amount), withdrawable)

	// current block equal to vesting end - witdrawable all = 1000000
	withdrawable = keeper.CalculateWithdrawable(vesting.VestingEndBlock, vesting)
	require.EqualValues(t, amount, withdrawable)

	vesting.VestingEndBlock = endHeight
	// setting LastModificationVested to prove that LastModificationVested is used in calculations (Vested field is just historical info)
	// and value set to number not divisible by number of periods - some more periods must pass to get some withdrawable amount
	vesting.LastModificationVested = notDivisibleAmount

	// current block one unlocking period more than last modification height - witdrawable = vested/number of unlocking periods (1000/((110000-10000)/10) = 0.1 this is less than 1 so 0 to wiithdraw)
	withdrawable = keeper.CalculateWithdrawable(startHeight+unlockingPeriod, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block 9 unlocking period more than last modification height (still not divisible - one block before being divisible) -
	// witdrawable = vested/number of unlocking periods 9*(1000/((110000-10000)/10) = 0.9 this is less than 1 so 0 to wiithdraw)
	withdrawable = keeper.CalculateWithdrawable(startHeight+10*unlockingPeriod-1, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	// current block 10 unlocking period more than last modification height -
	// witdrawable = vested/number of unlocking periods 10*(1000/((110000-10000)/10) = 1)
	withdrawable = keeper.CalculateWithdrawable(startHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, startHeight+10*unlockingPeriod, vesting.LastModificationVested), withdrawable)

	vesting.LastModificationVested = amount
	// setting some coins withdrawn
	vesting.LastModificationWithdrawn = withdrawn

	// current block 10 unlocking periods more than last modification height - witdrawable =  10*vested/number of unlocking periods 10*(1000000/((110000-10000)/10) - 500 = 500)
	withdrawable = keeper.CalculateWithdrawable(startHeight+10*unlockingPeriod, vesting)
	require.EqualValues(t, testutils.GetExpectedWithdrawable(startHeight, endHeight, unlockingPeriod, startHeight+10*unlockingPeriod, vesting.LastModificationVested).Sub(withdrawn), withdrawable)

	// current block equal to vesting end - witdrawable all available = 1000000 - 500
	withdrawable = keeper.CalculateWithdrawable(endHeight, vesting)
	require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn), withdrawable)
}

