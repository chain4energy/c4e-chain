package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
)

func TestSendVestingAccount(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), true)

}

func TestSendVestingAccountJustBeforeLockEnd(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	lockupDuration := time.Duration(1000)
	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, lockupDuration, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)
	testHelper.SetContextBlockTime(testHelper.Context.BlockTime().Add(lockupDuration -1 ))

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), true)

}

func TestSendVestingAccountNoRestartVesting(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), false)

}

func TestSendVestingAccountOnPoolLockEnd(t *testing.T) {
	sendVestingAccountPoolLockEndedTest(t, 0, true)
}

func TestSendVestingAccountNoRestartVestingOnPoolLockEnd(t *testing.T) {
	sendVestingAccountPoolLockEndedTest(t, 0, false)
}

func TestSendVestingAccountAfterPoolLockEnd(t *testing.T) {
	sendVestingAccountPoolLockEndedTest(t, 1, true)
}

func TestSendVestingAccountNoRestartVestingAfterPoolLockEnd(t *testing.T) {
	sendVestingAccountPoolLockEndedTest(t, 1, false)
}

func sendVestingAccountPoolLockEndedTest(t *testing.T, afterEnd time.Duration, restartVesting bool) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	lockupDuration := time.Duration(1000)
	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, lockupDuration, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)
	
	testHelper.SetContextBlockTime(testHelper.Context.BlockTime().Add(lockupDuration + afterEnd))
	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, vPool1, sdk.NewInt(100), restartVesting,
	"send to new vesting account - vesting available: 0 is smaller than requested amount: 100: insufficient funds")
	testHelper.BankUtils.VerifyAccountDefultDenomBalance(accAddr, accInitBalance)

}

func TestSendVestingAccountMultiple(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(4, 0)

	accAddr := acountsAddresses[0]
	vestAccAddr1 := acountsAddresses[1]
	vestAccAddr2 := acountsAddresses[2]
	vestAccAddr3 := acountsAddresses[3]
	
	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, vestAccAddr1, vPool1, sdk.NewInt(100), true)
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, vestAccAddr2, vPool1, sdk.NewInt(34), true)
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, vestAccAddr3, vPool1, sdk.NewInt(345), true)
}

func TestSendVestingAccountVestingPoolNotExistsForAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, "pool", sdk.NewInt(100), true,
		"send to new vesting account - withdraw all available error: withdraw all available - no vesting pools found error: address: cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq: not found")
}

func TestSendVestingAccountVestingPoolNotFound(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, "wrongpool", sdk.NewInt(100), true, "send to new vesting account - vesting pool with name wrongpool not found: not found")

}

func TestSendVestingAccounNotEnoughToSend(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, vPool1, sdk.NewInt(1100), true, "send to new vesting account - vesting available: 1000 is smaller than requested amount: 1100: insufficient funds")

}

func TestSendVestingAccountNotEnoughToSendAferSuccesfulSend(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), true)
	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, vPool1, sdk.NewInt(950), true, "send to new vesting account - vesting available: 900 is smaller than requested amount: 950: insufficient funds")

}

func TestSendVestingAccountAlreadyExists(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), true)
	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, vPool1, sdk.NewInt(300), true, "new vesting account - account address: "+accAddr2.String()+": entity already exists")

}
