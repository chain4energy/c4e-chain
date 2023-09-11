package keeper_test

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/v2/testutil/app"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"testing"
	"time"

	testcosmos "github.com/chain4energy/c4e-chain/v2/testutil/cosmossdk"
	testutils "github.com/chain4energy/c4e-chain/v2/testutil/module/cfevesting"
)

func TestWithdrawAllAvailableOnLockStart(t *testing.T) {
	vested := math.NewInt(1000000)

	testHelper := app.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(vested, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 1, vested, math.ZeroInt())
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, math.ZeroInt(), vested, math.ZeroInt())
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableManyVestingPoolsOnLockStart(t *testing.T) {
	vested := math.NewInt(1000000)
	testHelper := app.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(vested.MulRaw(3), types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 3, vested, math.ZeroInt())
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, math.ZeroInt(), vested.MulRaw(3), math.ZeroInt())
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableDuringLock(t *testing.T) {
	vested := math.NewInt(1000000)
	withdrawable := math.ZeroInt()
	testHelper := app.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(vested, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 1, vested, math.ZeroInt())
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, math.ZeroInt(), vested, withdrawable)
	accountVestingPools.VestingPools[0].Withdrawn = withdrawable
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableManyLockedDuringLock(t *testing.T) {
	vested := math.NewInt(1000000)
	withdrawable := math.ZeroInt()
	testHelper := app.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(vested.MulRaw(3), types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 3, vested, math.ZeroInt())
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, math.ZeroInt(), vested.MulRaw(3), withdrawable.MulRaw(3))
	for _, vesting := range accountVestingPools.VestingPools {
		vesting.Withdrawn = withdrawable
	}
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableAllToWithdrawAndSomeWithdrawn(t *testing.T) {
	vested := math.NewInt(1000000)
	withdrawable := vested
	withdrawn := math.NewInt(300)
	balance := vested.Sub(withdrawn)

	testHelper := app.SetupTestAppWithHeightAndTime(t, 110000, testutils.CreateTimeFromNumOfHours(110000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(balance, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 1, vested, withdrawn)
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, math.ZeroInt(), balance, withdrawable.Sub(withdrawn))
	accountVestingPools.VestingPools[0].Withdrawn = withdrawable
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

}

func TestVestAndReservationWithdrawAllAvailable(t *testing.T) {
	vested := math.NewInt(1000000)
	testHelper := app.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	accAddr := acountsAddresses[0]

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(vested, accAddr)

	modifyVestingType := func(vt *types.VestingType) {
		vt.LockupPeriod = testutils.CreateDurationFromNumOfHours(9000)
		vt.VestingPeriod = testutils.CreateDurationFromNumOfHours(100000)
	}
	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypesWithModification(modifyVestingType, 1, 1, 1)

	startTime := testHelper.Context.BlockTime()

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *vestingTypes.VestingTypes[0], vested, vested /*0,*/, math.ZeroInt(), math.ZeroInt() /*0,*/, vested)
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
	reservationAmount := vested.QuoRaw(2)
	testHelper.C4eVestingUtils.AddReservationToVestingPool(accAddr, vPool1, 0, reservationAmount.QuoRaw(2))
	testHelper.C4eVestingUtils.AddReservationToVestingPool(accAddr, vPool1, 1, reservationAmount.QuoRaw(2))
	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, math.ZeroInt(), vested, math.ZeroInt())

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []math.Int{vested}, []math.Int{math.ZeroInt()})

	testHelper.SetContextBlockHeightAndTime(int64(110000), testutils.CreateTimeFromNumOfHours(110000))

	withdrawn := vested.Sub(reservationAmount)
	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, math.ZeroInt(), vested, withdrawn)

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []math.Int{vested}, []math.Int{withdrawn}, startTime)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableManyVestedAllToWithdrawAndSomeWithdrawn(t *testing.T) {
	vested := math.NewInt(1000000)
	withdrawable := vested
	withdrawn := math.NewInt(300)
	balance := vested.Sub(withdrawn).MulRaw(3)

	testHelper := app.SetupTestAppWithHeightAndTime(t, 110000, testutils.CreateTimeFromNumOfHours(110000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(balance, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestingPools := testHelper.C4eVestingUtils.SetupAccountVestingPools(accAddr.String(), 3, vested, withdrawn)
	testHelper.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools()
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, math.ZeroInt(), balance, withdrawable.MulRaw(3).Sub(withdrawn.MulRaw(3)))
	for _, vesting := range accountVestingPools.VestingPools {
		vesting.Withdrawn = withdrawable
	}
	testHelper.C4eVestingUtils.CompareStoredAcountVestingPools(accAddr, accountVestingPools)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestVestAndWithdrawAllAvailable(t *testing.T) {
	vested := math.NewInt(1000000)
	testHelper := app.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	accAddr := acountsAddresses[0]

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(vested, accAddr)

	modifyVestingType := func(vt *types.VestingType) {
		vt.LockupPeriod = testutils.CreateDurationFromNumOfHours(9000)
		vt.VestingPeriod = testutils.CreateDurationFromNumOfHours(100000)
	}
	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypesWithModification(modifyVestingType, 1, 1, 1)

	startTime := testHelper.Context.BlockTime()

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *vestingTypes.VestingTypes[0], vested, vested /*0,*/, math.ZeroInt(), math.ZeroInt() /*0,*/, vested)
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, math.ZeroInt(), vested, math.ZeroInt())

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []math.Int{vested}, []math.Int{math.ZeroInt()})

	testHelper.SetContextBlockHeightAndTime(int64(110000), testutils.CreateTimeFromNumOfHours(110000))

	withdrawn := vested
	testHelper.C4eVestingUtils.MessageWithdrawAllAvailable(accAddr, math.ZeroInt(), vested, withdrawn)

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []math.Int{vested}, []math.Int{withdrawn}, startTime)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestWithdrawAllAvailableBadAddress(t *testing.T) {

	testHelper := app.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))

	testHelper.C4eVestingUtils.MessageWithdrawAllAvailableError("badaddress", "withdraw all available owner parsing error: badaddress: decoding bech32 failed: invalid separator index -1: failed to parse")
}
