package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
)

func TestCreateVestingAccount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := sdk.NewInt(100000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(10000)
	coins := sdk.Coins{{Amount: sendAmount, Denom: commontestutils.DefaultTestDenom}}
	startTime := time.Date(2025, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2035, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestCreateVestingAccountAccountExists(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]
	accAddr3 := acountsAddresses[2]

	accBalance := sdk.NewInt(100000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(10000)
	coins := sdk.Coins{{Amount: sendAmount, Denom: commontestutils.DefaultTestDenom}}
	startTime := time.Date(2025, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2035, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)

	accBalance = accBalance.Sub(coins.AmountOf(commontestutils.DefaultTestDenom))
	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr3,
		coins,
		startTime,
		endTime,
		accBalance,
	)

	accBalance = accBalance.Sub(coins.AmountOf(commontestutils.DefaultTestDenom))
	testHelper.C4eVestingUtils.MessageCreateVestingAccountError(
		accAddr1,
		accAddr1,
		coins,
		startTime,
		endTime,
		accBalance,
		"create vesting account - account address: "+accAddr1.String()+": entity already exists",
	)

	testHelper.C4eVestingUtils.MessageCreateVestingAccountError(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
		"create vesting account - account address: "+accAddr2.String()+": entity already exists",
	)

	testHelper.C4eVestingUtils.MessageCreateVestingAccountError(
		accAddr1,
		accAddr3,
		coins,
		startTime,
		endTime,
		accBalance,
		"create vesting account - account address: "+accAddr3.String()+": entity already exists",
	)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestCreateVestingAccountNotEnoughFunds(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]
	accAddr3 := acountsAddresses[2]

	accBalance := sdk.NewInt(15000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(10000)
	coins := sdk.Coins{{Amount: sendAmount, Denom: commontestutils.DefaultTestDenom}}
	startTime := time.Date(2025, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2035, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)

	accBalance = accBalance.Sub(coins.AmountOf(commontestutils.DefaultTestDenom))
	testHelper.C4eVestingUtils.MessageCreateVestingAccountError(
		accAddr1,
		accAddr3,
		coins,
		startTime,
		endTime,
		accBalance,
		"create vesting account - send coins to vesting account error (from: "+accAddr1.String()+", to: "+accAddr3.String()+", amount: "+coins.String()+"): "+accBalance.String()+commontestutils.DefaultTestDenom+" is smaller than "+coins.String()+": insufficient funds: failed to send coins",
	)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestCreateVestingAccountStartTimeAfterEndTime(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := sdk.NewInt(15000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(10000)
	coins := sdk.Coins{{Amount: sendAmount, Denom: commontestutils.DefaultTestDenom}}
	startTime := time.Date(2035, 2, 3, 0, 0, 0, 0, time.Local)
	endTime := time.Date(2025, 2, 3, 0, 0, 0, 0, time.Local)

	testHelper.C4eVestingUtils.MessageCreateVestingAccountError(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
		"create vesting account - start time is after end time error ("+startTime.String()+" > "+endTime.String()+"): start time cannot be after end time",
	)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}
