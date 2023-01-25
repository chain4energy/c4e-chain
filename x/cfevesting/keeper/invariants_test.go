package keeper_test

import (
	"fmt"
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestNonNegativeVestingPoolAmountsInvariantCorrect(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)
	testUtil.SetupAccountVestingPools(ctx, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, false,
		"cfevesting: nonnegative vesting pool amounts invariant\n\tno negative amounts in vesting pools\n")
}

func TestNonNegativeVestingPoolAmountsInvariantManyyVestingPoolsCorrect(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(3, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)
	testUtil.SetupAccountVestingPools(ctx, accounts[0].String(), 10, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.SetupAccountVestingPools(ctx, accounts[1].String(), 10, sdk.NewInt(102), sdk.NewInt(23))
	testUtil.SetupAccountVestingPools(ctx, accounts[2].String(), 10, sdk.NewInt(89), sdk.NewInt(89))

	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, false,
		"cfevesting: nonnegative vesting pool amounts invariant\n\tno negative amounts in vesting pools\n")
}

func TestNonNegativeVestingPoolAmountsInvariantCorrectZeros(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)
	testUtil.SetupAccountVestingPools(ctx, accounts[0].String(), 1, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, false,
		"cfevesting: nonnegative vesting pool amounts invariant\n\tno negative amounts in vesting pools\n")

}

func TestNonNegativeVestingPoolAmountsInvariantNagativeInitiallyLocked(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.InitiallyLocked = sdk.NewInt(-1)
	}
	testUtil.SetupAccountVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, true,
		fmt.Sprintf("cfevesting: nonnegative vesting pool amounts invariant\n\tnegative InitiallyLocked -1 in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestNonNegativeVestingPoolAmountsInvariantNagativeSent(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Sent = sdk.NewInt(-1)
	}
	testUtil.SetupAccountVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, true,
		fmt.Sprintf("cfevesting: nonnegative vesting pool amounts invariant\n\tnegative Sent -1 in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestNonNegativeVestingPoolAmountsInvariantNagativeWithdrawn(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Withdrawn = sdk.NewInt(-1)
	}
	testUtil.SetupAccountVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, true,
		fmt.Sprintf("cfevesting: nonnegative vesting pool amounts invariant\n\tnegative Withdrawn -1 in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestVestingPoolConsistentDataInvariantCorrect(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)
	testUtil.SetupAccountVestingPools(ctx, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, false,
		"cfevesting: vesting pool consistent data invariant\n\tno inconsistent vesting pools\n")
}

func TestVestingPoolConsistentDataInvariantManyVestingPoolCorrect(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(3, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)
	testUtil.SetupAccountVestingPools(ctx, accounts[0].String(), 10, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.SetupAccountVestingPools(ctx, accounts[1].String(), 10, sdk.NewInt(89), sdk.NewInt(34))
	testUtil.SetupAccountVestingPools(ctx, accounts[2].String(), 10, sdk.NewInt(23), sdk.NewInt(21))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, false,
		"cfevesting: vesting pool consistent data invariant\n\tno inconsistent vesting pools\n")
}

func TestVestingPoolConsistentDataInvariantCorrectAllWithdrawn(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)
	testUtil.SetupAccountVestingPools(ctx, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(100))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, false,
		"cfevesting: vesting pool consistent data invariant\n\tno inconsistent vesting pools\n")
}

func TestVestingPoolConsistentDataInvariantWrongWithdrawn(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Withdrawn = sdk.NewInt(101)
	}
	testUtil.SetupAccountVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, true,
		fmt.Sprintf("cfevesting: vesting pool consistent data invariant\n\tWithdrawn (101) + Sent (0) GT InitiallyLocked (100) in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestVestingPoolConsistentDataInvariantWrongSent(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Sent = sdk.NewInt(101)
		pool.Withdrawn = sdk.ZeroInt()
	}
	testUtil.SetupAccountVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, true,
		fmt.Sprintf("cfevesting: vesting pool consistent data invariant\n\tWithdrawn (0) + Sent (101) GT InitiallyLocked (100) in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestVestingPoolConsistentDataInvariantWrongWithdrawnPlusSent(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Sent = sdk.NewInt(30)
		pool.Withdrawn = sdk.NewInt(71)
	}
	testUtil.SetupAccountVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, true,
		fmt.Sprintf("cfevesting: vesting pool consistent data invariant\n\tWithdrawn (71) + Sent (30) GT InitiallyLocked (100) in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestCheckModuleAccountInvariantCorrect(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eVestingUtils.SetupAccountVestingPools(accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(30))
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(70), types.ModuleName)
	testHelper.C4eVestingUtils.CheckModuleAccountInvariant(false,
		"cfevesting: module account invariant\n\tamount consistent with vesting pools\n")

}

func TestCheckModuleAccountInvariantMamyAccountVestingPoolsCorrect(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(3, 0)
	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eVestingUtils.SetupAccountVestingPools(accounts[0].String(), 10, sdk.NewInt(100), sdk.NewInt(30))
	testHelper.C4eVestingUtils.SetupAccountVestingPools(accounts[1].String(), 10, sdk.NewInt(50), sdk.NewInt(20))
	testHelper.C4eVestingUtils.SetupAccountVestingPools(accounts[2].String(), 10, sdk.NewInt(80), sdk.NewInt(54))
	vestingPoolsAmount := int64(10*(100-30) + 10*(50-20) + 10*(80-54))
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(vestingPoolsAmount), types.ModuleName)
	testHelper.C4eVestingUtils.CheckModuleAccountInvariant(false,
		"cfevesting: module account invariant\n\tamount consistent with vesting pools\n")

}

func TestCheckModuleAccountInvariantNoModuleAccount(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eVestingUtils.SetupAccountVestingPools(accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))

	testHelper.C4eVestingUtils.CheckModuleAccountInvariant(true,
		"cfevesting: module account invariant\n\tamount (0) inconsistent with vesting pools (50)\n")

}

func TestCheckModuleAccountInvariantWrongBalance(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(1, 0)
	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eVestingUtils.SetupAccountVestingPools(accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(30))
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(71), types.ModuleName)
	testHelper.C4eVestingUtils.CheckModuleAccountInvariant(true,
		"cfevesting: module account invariant\n\tamount (71) inconsistent with vesting pools (70)\n")

}
