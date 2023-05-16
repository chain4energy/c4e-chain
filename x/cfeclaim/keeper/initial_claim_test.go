package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"testing"
	"time"
)

func TestCorrectInitialClaim(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
}

func TestCorrectmanyInitialClaims(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)

	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 1, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 1, nil, nil)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 1, claimEntries)

	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[2], 0, 80000002)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[3], 0, 80000003)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[4], 0, 80000004)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 1, 80000001)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[2], 1, 80000002)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[3], 1, 80000003)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[4], 1, 80000004)
}

func TestCorrectmanyInitialClaimsForDifferentCampaigns(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[2], 0, 80000002)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[3], 0, 80000003)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[4], 0, 80000004)
}

func TestInitialClaimAlreadyClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	testHelper.C4eClaimUtils.ClaimInitialError(acountsAddresses[1], 0, fmt.Sprintf("address %s, campaignId: 0, missionId: 0: mission already completed", acountsAddresses[1].String()))
}

func TestInitialClaimRecordDosentExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses[:9], 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	testHelper.C4eClaimUtils.ClaimInitialError(acountsAddresses[10], 0, fmt.Sprintf("user claim entries not found for address %s: not found", acountsAddresses[10].String()))
}

func TestInitialClaimWrongCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses[:9], 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitialError(acountsAddresses[1], 1, "camapign not found: campaignId 1: not found")
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 1, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitialError(acountsAddresses[1], 1, fmt.Sprintf("campaign record with id 1 not found for address %s: not found", acountsAddresses[1].String()))
}

func TestInitialClaimCampaignDidntStartYey(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.StartTime = testHelper.Context.BlockTime().Add(time.Second * 5)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)

	testHelper.C4eClaimUtils.ClaimInitialError(acountsAddresses[1], 0, fmt.Sprintf("campaign 0 not started yet (%s < startTime %s) error: campaign is disabled", testHelper.Context.BlockTime(), campaign.StartTime))
}

func TestInitialClaimCampaignNotEnabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaign(acountsAddresses[0].String(), 0)
	testHelper.C4eClaimUtils.ClaimInitialError(acountsAddresses[1], 0, "campaign 0 error: campaign is disabled")
}

func TestInitialClaimCampaignIsOver(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.ClaimInitialError(acountsAddresses[1], 0, fmt.Sprintf("campaign 0 has already ended (%s > endTime %s) error: campaign is disabled", testHelper.Context.BlockTime(), campaign.EndTime))
}

func TestInitialClaimFreeInitialAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	campaign.InitialClaimFreeAmount = math.NewInt(100000)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
}

func TestInitialClaim2Campaigns(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)

	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 1, mission)

	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 1, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 1, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 1, 80000001)
}

func TestInitialClaimFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	campaign.FeegrantAmount = math.NewInt(100000)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], campaign.FeegrantAmount.MulRaw(int64(len(acountsAddresses))))
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
}
