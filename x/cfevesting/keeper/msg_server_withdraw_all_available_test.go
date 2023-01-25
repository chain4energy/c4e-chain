package keeper_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"

	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
)

func TestWithdrawAllAvailableOnLockStart(t *testing.T) {
	vested := sdk.NewInt(1000000)

	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(vested, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 1, vested, sdk.ZeroInt())
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, sdk.ZeroInt(), vested, sdk.ZeroInt())
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableManyVestingPoolsOnLockStart(t *testing.T) {
	vested := sdk.NewInt(1000000)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(vested.MulRaw(3), types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 3, vested, sdk.ZeroInt())
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, sdk.ZeroInt(), vested.MulRaw(3), sdk.ZeroInt())
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableDuringLock(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := sdk.ZeroInt()
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(vested, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 1, vested, sdk.ZeroInt())
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, sdk.ZeroInt(), vested, withdrawable)
	accountVestingPools.VestingPools[0].Withdrawn = withdrawable
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableManyLockedDuringLock(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := sdk.ZeroInt()
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(vested.MulRaw(3), types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 3, vested, sdk.ZeroInt())
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, sdk.ZeroInt(), vested.MulRaw(3), withdrawable.MulRaw(3))
	for _, vesting := range accountVestingPools.VestingPools {
		vesting.Withdrawn = withdrawable
	}
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableAllToWithdrawAndSomeWithdrawn(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := vested
	withdrawn := sdk.NewInt(300)
	balance := vested.Sub(withdrawn)

	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 110000, testutils.CreateTimeFromNumOfHours(110000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(balance, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 1, vested, withdrawn)
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, sdk.ZeroInt(), balance, withdrawable.Sub(withdrawn))
	accountVestingPools.VestingPools[0].Withdrawn = withdrawable
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

}

func TestWithdrawAllAvailableManyVestedAllToWithdrawAndSomeWithdrawn(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := vested
	withdrawn := sdk.NewInt(300)
	balance := vested.Sub(withdrawn).MulRaw(3)

	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 110000, testutils.CreateTimeFromNumOfHours(110000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(balance, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 3, vested, withdrawn)
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, sdk.ZeroInt(), balance, withdrawable.MulRaw(3).Sub(withdrawn.MulRaw(3)))
	for _, vesting := range accountVestingPools.VestingPools {
		vesting.Withdrawn = withdrawable
	}
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestVestAndWithdrawAllAvailable(t *testing.T) {
	vested := sdk.NewInt(1000000)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	accAddr := acountsAddresses[0]

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(vested, accAddr)

	modifyVestingType := func(vt *types.VestingType) {
		vt.LockupPeriod = testutils.CreateDurationFromNumOfHours(9000)
		vt.VestingPeriod = testutils.CreateDurationFromNumOfHours(100000)
	}
	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypesWithModification(modifyVestingType, 1, 1, 1)

	startTime := testHelper.Context.BlockTime()

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *vestingTypes.VestingTypes[0], vested, vested /*0,*/, sdk.ZeroInt(), sdk.ZeroInt() /*0,*/, vested)
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, sdk.ZeroInt(), vested, sdk.ZeroInt())

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []sdk.Int{vested}, []sdk.Int{sdk.ZeroInt()})

	testHelper.SetContextBlockHeightAndTime(int64(110000), testutils.CreateTimeFromNumOfHours(110000))

	withdrawn := vested
	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, sdk.ZeroInt(), vested, withdrawn)

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []sdk.Int{vested}, []sdk.Int{withdrawn}, startTime)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableBadAddress(t *testing.T) {

	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailableError("badaddress", "withdraw all available address parsing error: badaddress: decoding bech32 failed: invalid separator index -1: failed to parse")
}
