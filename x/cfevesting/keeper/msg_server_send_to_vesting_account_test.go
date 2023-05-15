package keeper_test

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/testutil/app"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
)

func TestSendVestingAccount(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), true, sdk.NewInt(95))
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountFromGenesisPool(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
	testHelper.C4eVestingUtils.MessageCreateGenesisVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), true, sdk.NewInt(95))
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountJustBeforeLockEnd(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	lockupDuration := time.Duration(1000)
	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, lockupDuration, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)
	testHelper.SetContextBlockTime(testHelper.Context.BlockTime().Add(lockupDuration - 1))

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), true, sdk.NewInt(95))
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountNoRestartVesting(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), false, sdk.NewInt(95))
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountMultiple(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(4, 0)

	accAddr := acountsAddresses[0]
	vestAccAddr1 := acountsAddresses[1]
	vestAccAddr2 := acountsAddresses[2]
	vestAccAddr3 := acountsAddresses[3]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, vestAccAddr1, vPool1, sdk.NewInt(100), true, sdk.NewInt(95))
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, vestAccAddr2, vPool1, sdk.NewInt(34), true, sdk.NewInt(32))
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, vestAccAddr3, vPool1, sdk.NewInt(345), true, sdk.NewInt(327))
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountVestingPoolNotExistsForAddress(t *testing.T) {
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, "pool", sdk.NewInt(100), true,
		fmt.Sprintf("send locked to new vesting account: no vesting pool pool found for address %s: not found", accAddr))
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountVestingPoolNotFound(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)
	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, "wrongpool", sdk.NewInt(100), true, fmt.Sprintf("send locked to new vesting account: no vesting pool wrongpool found for address %s: not found", accAddr))
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccounNotEnoughToSend(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)
	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, vPool1, sdk.NewInt(1100), true, "send to new vesting account - vesting available: 1000 is smaller than requested amount: 1100: insufficient funds")

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountNotEnoughToSendAferSuccesfulSend(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), true, sdk.NewInt(95))
	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, vPool1, sdk.NewInt(950), true, "send to new vesting account - vesting available: 900 is smaller than requested amount: 950: insufficient funds")
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountAlreadyExists(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(100), true, sdk.NewInt(95))
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sdk.NewInt(300), true, sdk.NewInt(380))
	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountVestingTypesFreeZeroFree(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := types.VestingTypes{}

	vestingTypeZeroFree := types.VestingType{
		Name:          "vestingTypeZeroFree",
		LockupPeriod:  2324,
		VestingPeriod: 42423,
		Free:          sdk.ZeroDec(),
	}

	vestingTypesArray := []*types.VestingType{&vestingTypeZeroFree}
	vestingTypes.VestingTypes = vestingTypesArray
	sendToVestingAccountAmount := sdk.NewInt(100)
	testHelper.C4eVestingUtils.SetVestingTypes(vestingTypes)
	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, vestingTypeZeroFree, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sendToVestingAccountAmount, true, sdk.NewInt(100))
}

func TestSendVestingAccountVestingTypesFreeMaxFree(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := app.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := types.VestingTypes{}

	vestingTypeMaxFree := types.VestingType{
		Name:          "vestingTypeMaxFree",
		LockupPeriod:  2324,
		VestingPeriod: 42423,
		Free:          sdk.MustNewDecFromStr("1"),
	}

	vestingTypesArray := []*types.VestingType{&vestingTypeMaxFree}
	vestingTypes.VestingTypes = vestingTypesArray
	sendToVestingAccountAmount := sdk.NewInt(100)
	testHelper.C4eVestingUtils.SetVestingTypes(vestingTypes)
	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, vestingTypeMaxFree, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, vPool1, sendToVestingAccountAmount, true, sdk.NewInt(0))
}
