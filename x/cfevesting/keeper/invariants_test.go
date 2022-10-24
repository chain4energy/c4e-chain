package keeper_test

import (
	"fmt"
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestNonNegativeVestingPoolAmountsInvariantCorrect(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)
	testUtil.SetupAccountsVestingPools(ctx, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, false,
		"cfevesting: nonnegative vesting pool amounts invariant\n\tno negative amounts in vesting pools\n")
}

func TestNonNegativeVestingPoolAmountsInvariantCorrectZeros(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)
	testUtil.SetupAccountsVestingPools(ctx, accounts[0].String(), 1, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, false,
		"cfevesting: nonnegative vesting pool amounts invariant\n\tno negative amounts in vesting pools\n")

}

func TestNonNegativeVestingPoolAmountsInvariantNagativeLastModificationVested(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.LastModificationVested = sdk.NewInt(-1)
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, true,
		fmt.Sprintf("cfevesting: nonnegative vesting pool amounts invariant\n\tnegative LastModificationVested -1 in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestNonNegativeVestingPoolAmountsInvariantNagativeLastModificationWithdrawn(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.LastModificationWithdrawn = sdk.NewInt(-1)
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, true,
		fmt.Sprintf("cfevesting: nonnegative vesting pool amounts invariant\n\tnegative LastModificationWithdrawn -1 in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestNonNegativeVestingPoolAmountsInvariantNagativeVested(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Vested = sdk.NewInt(-1)
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, true,
		fmt.Sprintf("cfevesting: nonnegative vesting pool amounts invariant\n\tnegative Vested -1 in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestNonNegativeVestingPoolAmountsInvariantNagativeWithdrawn(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Withdrawn = sdk.NewInt(-1)
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckNonNegativeVestingPoolAmountsInvariant(ctx, true,
		fmt.Sprintf("cfevesting: nonnegative vesting pool amounts invariant\n\tnegative Withdrawn -1 in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestVestingPoolConsistentDataInvariantCorrect(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)
	testUtil.SetupAccountsVestingPools(ctx, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, false,
		"cfevesting: vesting pool consistent data invariant\n\tno inconsistent vesting pools\n")
}

func TestVestingPoolConsistentDataInvariantCorrectAllWithdrawn(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)
	testUtil.SetupAccountsVestingPools(ctx, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(100))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, false,
		"cfevesting: vesting pool consistent data invariant\n\tno inconsistent vesting pools\n")
}

func TestVestingPoolConsistentDataInvariantCorrectAfterMidifification(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.LastModificationVested = sdk.NewInt(80)
		pool.LastModificationWithdrawn = sdk.NewInt(30)
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, false,
		"cfevesting: vesting pool consistent data invariant\n\tno inconsistent vesting pools\n")
}

func TestVestingPoolConsistentDataInvariantCorrectAfterMidifificationAndSent(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Sent = sdk.NewInt(10)
		pool.LastModificationVested = sdk.NewInt(70)
		pool.LastModificationWithdrawn = sdk.NewInt(30)
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, false,
		"cfevesting: vesting pool consistent data invariant\n\tno inconsistent vesting pools\n")
}

func TestVestingPoolConsistentDataInvariantWrongLastModificationWithdrawn(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.LastModificationWithdrawn = sdk.NewInt(101)
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, true,
		fmt.Sprintf("cfevesting: vesting pool consistent data invariant\n\tLastModificationWithdrawn (101) GT LastModificationVested (100) in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestVestingPoolConsistentDataInvariantWrongWithdrawn(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Withdrawn = sdk.NewInt(101)
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, true,
		fmt.Sprintf("cfevesting: vesting pool consistent data invariant\n\tWithdrawn (101) + Sent (0) GT Vested (100) in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestVestingPoolConsistentDataInvariantWrongSent(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Sent = sdk.NewInt(101)
		pool.Withdrawn = sdk.ZeroInt()
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, true,
		fmt.Sprintf("cfevesting: vesting pool consistent data invariant\n\tWithdrawn (0) + Sent (101) GT Vested (100) in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestVestingPoolConsistentDataInvariantWrongWithdrawnPlusSent(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Sent = sdk.NewInt(30)
		pool.Withdrawn = sdk.NewInt(71)
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, true,
		fmt.Sprintf("cfevesting: vesting pool consistent data invariant\n\tWithdrawn (71) + Sent (30) GT Vested (100) in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestVestingPoolConsistentDataInvariantWrongMainAmounVsLastModificationAmount(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtil(t)

	modification := func(pool *types.VestingPool) {
		pool.Sent = sdk.NewInt(9)
		pool.LastModificationVested = sdk.NewInt(70)
		pool.LastModificationWithdrawn = sdk.NewInt(30)
	}
	testUtil.SetupAccountsVestingPoolsWithModification(ctx, modification, accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))
	testUtil.CheckVestingPoolConsistentDataInvariant(ctx, true,
		fmt.Sprintf("cfevesting: vesting pool consistent data invariant\n\t Vested (100) - Withdrawn (50) - Sent (9) <> LastModificationVested (70) - LastModificationWithdrawn (30) in vesting pool: test-vesting-account-name1-1 for address: %s\n", accounts[0].String()))
}

func TestCheckModuleAccountInvariantCorrect(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eVestingUtils.SetupAccountsVestingPools(accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(30))
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(70), types.ModuleName)
	testHelper.C4eVestingUtils.CheckModuleAccountInvariant(false,
		"cfevesting: module account invariant\n\tamount consistent with vesting pools\n")

}

func TestCheckModuleAccountInvariantNoModuleAccount(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eVestingUtils.SetupAccountsVestingPools(accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(50))

	testHelper.C4eVestingUtils.CheckModuleAccountInvariant(true,
		"cfevesting: module account invariant\n\tamount (0) inconsistent with vesting pools (50)\n")

}

func TestCheckModuleAccountInvariantWrongBalance(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(1, 0)
	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eVestingUtils.SetupAccountsVestingPools(accounts[0].String(), 1, sdk.NewInt(100), sdk.NewInt(30))
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(71), types.ModuleName)
	testHelper.C4eVestingUtils.CheckModuleAccountInvariant(true,
		"cfevesting: module account invariant\n\tamount (71) inconsistent with vesting pools (70)\n")

}
