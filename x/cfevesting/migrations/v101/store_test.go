package v101_test

import (
	"testing"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationManyAccountVestingPoolsWithManyPools(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	testUtil.SetupAccountVestingPools(ctx, accounts[0].String(), 10, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.SetupAccountVestingPools(ctx, accounts[1].String(), 10, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.SetupAccountVestingPools(ctx, accounts[2].String(), 10, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.SetupAccountVestingPools(ctx, accounts[3].String(), 10, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.SetupAccountVestingPools(ctx, accounts[4].String(), 10, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.MigrateV100ToV101(t, ctx)
}

func TestMigrationNoAccountVestingPools(t *testing.T) {
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	testUtil.MigrateV100ToV101(t, ctx)
}

func TestMigrationManyAccountVestingPoolsWithNoPools(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	testUtil.SetupAccountVestingPools(ctx, accounts[0].String(), 0, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.SetupAccountVestingPools(ctx, accounts[1].String(), 0, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.SetupAccountVestingPools(ctx, accounts[2].String(), 0, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.SetupAccountVestingPools(ctx, accounts[3].String(), 0, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.SetupAccountVestingPools(ctx, accounts[4].String(), 0, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.MigrateV100ToV101(t, ctx)
}

func TestMigrationOneAccountVestingPoolsWithOnePool(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	testUtil.SetupAccountVestingPools(ctx, accounts[0].String(), 1, sdk.ZeroInt(), sdk.ZeroInt())
	testUtil.MigrateV100ToV101(t, ctx)
}