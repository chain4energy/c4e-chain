package keeper_test

import (
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	cfeairdroptypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"
)

func TestCompleteDelegationMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.CompleteDelegationMission(0, 1, acountsAddresses[1], delagationAmount)
	testHelper.C4eAirdropUtils.ClaimMission(0, 1, acountsAddresses[1])
}

func TestCompleteVoteMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeairdroptypes.MissionVote
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.CompleteVoteMission(0, 1, acountsAddresses[1])
	testHelper.C4eAirdropUtils.ClaimMission(0, 1, acountsAddresses[1])
}

func TestClaimMissionDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.ClaimMissionError(0, 2, acountsAddresses[1], "mission not found - campaignId 0, missionId 2: not found")
}

func TestClaimCampaignDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 0, acountsAddresses[1], "camapign not found: campaignId 1: not found")
}

func TestClaimNoInitialClaimError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.ClaimMissionError(0, 1, acountsAddresses[1], fmt.Sprintf("initial mission not completed: address %s, campaignId: 0: mission not completed yet", acountsAddresses[1].String()))
}

func TestClaimMissionCampaignHasEnded(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.CompleteDelegationMission(0, 1, acountsAddresses[1], delagationAmount)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.ClaimMissionError(0, 1, acountsAddresses[1], fmt.Sprintf("campaign 0 has already ended (%s > endTime %s) error: campaign is disabled", testHelper.Context.BlockTime(), campaign.EndTime))
}

func TestClaimMissionWithTypeClaim(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeairdroptypes.MissionClaim
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.ClaimMission(0, 1, acountsAddresses[1])
}

func TestClaimMissionAlreadyClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeairdroptypes.MissionClaim
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eAirdropUtils.ClaimMissionError(0, 1, acountsAddresses[1], fmt.Sprintf("mission already completed: address %s, campaignId: 0, missionId: 1: mission already completed", acountsAddresses[1].String()))
}

func TestFullCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	mission.MissionType = cfeairdroptypes.MissionVote
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 60000001)

	delagationAmount := sdk.NewInt(1000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.CompleteDelegationMission(0, 1, acountsAddresses[1], delagationAmount)

	testHelper.C4eAirdropUtils.CompleteVoteMission(0, 2, acountsAddresses[1])

	testHelper.C4eAirdropUtils.ClaimMission(0, 1, acountsAddresses[1])

	testHelper.C4eAirdropUtils.ClaimMission(0, 2, acountsAddresses[1])
}

func TestClaimMissionWithTypeClaimRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses[:10], 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)

	testHelper.C4eAirdropUtils.ClaimMissionError(0, 1, acountsAddresses[10], fmt.Sprintf("user airdrop entries not found for address %s: not found", acountsAddresses[10].String()))
}
