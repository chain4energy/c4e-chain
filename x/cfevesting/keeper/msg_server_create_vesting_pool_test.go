package keeper_test

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"github.com/chain4energy/c4e-chain/testutil/testapp"
)

func TestCreateVestingPool(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*usedVestingType}, []math.Int{vested}, []math.Int{sdk.ZeroInt()})

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, true, true, vPool2, 1200, *usedVestingType, vested, accInitBalance.Sub(vested) /*0,*/, vested, accInitBalance.Sub(vested.MulRaw(2)) /*0,*/, vested.MulRaw(2))

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1, vPool2}, []time.Duration{1000, 1200}, []types.VestingType{*usedVestingType, *usedVestingType}, []math.Int{vested, vested}, []math.Int{sdk.ZeroInt(), sdk.ZeroInt()})

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestCreateVestingPoolUnknownVestingType(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(sdk.NewInt(10000), accAddr)

	testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
	testHelper.C4eVestingUtils.MessageCreateVestingPoolError(accAddr, "pool", 1000, types.VestingType{Name: "unknown"}, vested, "create vesting pool - get vesting type error: vesting type not found: unknown: not found")

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestCreateVestingPoolNameDuplication(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*usedVestingType}, []math.Int{vested}, []math.Int{sdk.ZeroInt()})

	testHelper.C4eVestingUtils.MessageCreateVestingPoolError(accAddr, vPool1, 1000, *usedVestingType, vested, "add vesting pool - vesting pool name: "+vPool1+": entity already exists")

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestCreateVestingPoolEmptyName(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPoolError(accAddr, "", 1000, *usedVestingType, vested, "add vesting pool empty name: wrong param value")
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}
