package keeper_test

import (
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"
)

func TestCorrectInitialClaim(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
}

func TestCorrectmanyInitialClaims(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[2], 0, 80000002)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[3], 0, 80000003)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[4], 0, 80000004)
}

func TestInitialClaimAlreadyClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	testHelper.C4eAirdropUtils.ClaimInitialError(acountsAddresses[1], 0, fmt.Sprintf("mission already completed: address %s, campaignId: 0, missionId: 0: mission already completed", acountsAddresses[1].String()))
}

func TestInitialClaimRecordDosentExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses[:9], 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	testHelper.C4eAirdropUtils.ClaimInitialError(acountsAddresses[10], 0, fmt.Sprintf("user airdrop entries not found for address %s: not found", acountsAddresses[10].String()))
}

func TestInitialClaimWrongCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses[:9], 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitialError(acountsAddresses[1], 1, "camapign not found: campaignId 1: not found")
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 1)
	testHelper.C4eAirdropUtils.ClaimInitialError(acountsAddresses[1], 1, fmt.Sprintf("campaign record with id 1 not found for address %s: not found", acountsAddresses[1].String()))
}

func TestInitialClaimCampaignDidntStartYey(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.StartTime = testHelper.Context.BlockTime().Add(time.Second * 5)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)

	testHelper.C4eAirdropUtils.ClaimInitialError(acountsAddresses[1], 0, fmt.Sprintf("campaign 0 not started yet (%s < startTime %s) error: campaign is disabled", testHelper.Context.BlockTime(), campaign.StartTime))
}

func TestInitialClaimCampaignNotEnabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseAirdropCampaign(acountsAddresses[0].String(), 0, types.AirdropCloseAction_AIRDROP_CLOSE_ACTION_UNSPECIFIED)
	testHelper.C4eAirdropUtils.ClaimInitialError(acountsAddresses[1], 0, "campaign 0 error: campaign is disabled")
}

func TestInitialClaimCampaignIsOver(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.ClaimInitialError(acountsAddresses[1], 0, fmt.Sprintf("campaign 0 has already ended (%s > endTime %s) error: campaign is disabled", testHelper.Context.BlockTime(), campaign.EndTime))
}

func TestInitialClaimFreeInitialAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	campaign.InitialClaimFreeAmount = sdk.NewInt(100000)
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
}

func TestInitialClaim2Campaigns(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)

	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 1, mission)

	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 1)

	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 1, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 1, 80000001)
}

func TestInitialClaimFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	campaign.FeegrantAmount = sdk.NewInt(100000)
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], campaign.FeegrantAmount.MulRaw(int64(len(acountsAddresses))))
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)

	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
}
