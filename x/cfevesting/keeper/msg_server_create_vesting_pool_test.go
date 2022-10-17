package keeper_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
)

func TestCreateVestingPool(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*usedVestingType}, []sdk.Int{vested}, []sdk.Int{sdk.ZeroInt()})

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, true, true, vPool2, 1200, *usedVestingType, vested, accInitBalance.Sub(vested) /*0,*/, vested, accInitBalance.Sub(vested.MulRaw(2)) /*0,*/, vested.MulRaw(2))

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1, vPool2}, []time.Duration{1000, 1200}, []types.VestingType{*usedVestingType, *usedVestingType}, []sdk.Int{vested, vested}, []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt()})

}

func TestCreateVestingPoolUnknownVestingType(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(sdk.NewInt(10000), accAddr)

	testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)

	testHelper.C4eVestingUtils.MessageCreateVestingPoolError(accAddr, "pool", 1000, types.VestingType{Name: "unknown"}, vested, "vesting type not found: unknown: not found")

}

func TestCreateVestingPoolNameDuplication(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*usedVestingType}, []sdk.Int{vested}, []sdk.Int{sdk.ZeroInt()})

	testHelper.C4eVestingUtils.MessageCreateVestingPoolError(accAddr, vPool1, 1000, *usedVestingType, vested, "vesting pool name already exists: "+vPool1+": invalid request")

}

func TestVestingId(t *testing.T) {
	vested := sdk.NewInt(1000)
	accInitBalance := sdk.NewInt(10000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.SetVestingTypes(vestingTypes)

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested, 1)

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, true, true, vPool2, 1000, *usedVestingType, vested, accInitBalance.Sub(vested), vested, accInitBalance.Sub(vested.MulRaw(2)), vested.MulRaw(2), 1, 2)

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, true, true, "v-pool-3", 1000, *usedVestingType, vested, accInitBalance.Sub(vested.MulRaw(2)), vested.MulRaw(2), accInitBalance.Sub(vested.MulRaw(3)), vested.MulRaw(3), 1, 2, 3)
}
