package keeper_test

import (
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"
)

func TestAddClaimRecordsFromWhitelistedAccountAllCoinsLocked(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := sdk.NewInt(100000000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(1000000045)
	coins := sdk.Coins{{Amount: sendAmount, Denom: testenv.DefaultTestDenom}}
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(accAddr2.String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(accAddr2.String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(accAddr2.String(), 0)
	airdropEntries, _ := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eAirdropUtils.AddClaimRecords(accAddr2, 0, airdropEntries)
}

func TestAddClaimRecordsFromWhitelistedAccountAllCoinsLockedAndAdditionalBalance(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := sdk.NewInt(100000000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(999990045)
	coins := sdk.Coins{{Amount: sendAmount, Denom: testenv.DefaultTestDenom}}
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(accAddr2.String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(accAddr2.String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(accAddr2.String(), 0)
	airdropEntries, _ := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(accAddr2, sdk.NewInt(10000))
	testHelper.C4eAirdropUtils.AddClaimRecords(accAddr2, 0, airdropEntries)
}

func TestAddClaimRecordsFromWhitelistedAccountHalfVested(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := sdk.NewInt(100000000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(999990045)
	coins := sdk.Coins{{Amount: sendAmount, Denom: testenv.DefaultTestDenom}}
	startTime := testHelper.Context.BlockTime()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.EndTime = endTime
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(accAddr2.String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(accAddr2.String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(accAddr2.String(), 0)
	airdropEntries, _ := createTestClaimRecords(acountsAddresses, 100000000)
	blockTime := startTime.Add(time.Hour * 5)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(accAddr2, sdk.NewInt(10000))
	testHelper.C4eAirdropUtils.AddClaimRecords(accAddr2, 0, airdropEntries)
}

func TestAddClaimRecordsFromWhitelistedAccountAllVested(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := sdk.NewInt(100000000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(1000000045)
	coins := sdk.Coins{{Amount: sendAmount, Denom: testenv.DefaultTestDenom}}
	startTime := testHelper.Context.BlockTime()
	endTime := startTime.Add(time.Hour * 11)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.EndTime = endTime
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(accAddr2.String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(accAddr2.String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(accAddr2.String(), 0)
	airdropEntries, _ := createTestClaimRecords(acountsAddresses, 100000000)
	blockTime := startTime.Add(time.Hour * 10)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.AddClaimRecords(accAddr2, 0, airdropEntries)
}

func TestAddClaimRecordsFromWhitelistedAccountTimeInFutureBiggerVestingAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := sdk.NewInt(100000000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(1599990045)
	coins := sdk.Coins{{Amount: sendAmount, Denom: testenv.DefaultTestDenom}}
	startTime := testHelper.Context.BlockTime()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.EndTime = endTime
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(accAddr2.String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(accAddr2.String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(accAddr2.String(), 0)
	airdropEntries, _ := createTestClaimRecords(acountsAddresses, 100000000)
	blockTime := startTime.Add(time.Hour * 5)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(accAddr2, sdk.NewInt(10000))
	testHelper.C4eAirdropUtils.AddClaimRecords(accAddr2, 0, airdropEntries)
}
