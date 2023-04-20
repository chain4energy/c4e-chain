package keeper_test

import (
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"
)

func TestCompleteDelegationMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 1)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
}

func TestCompleteVoteMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionVote
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteVoteMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
}

func TestClaimMissionDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMissionError(0, 2, acountsAddresses[1], "mission not found - campaignId 0, missionId 2: not found")
}

func TestClaimCampaignDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMissionError(1, 0, acountsAddresses[1], "camapign not found: campaignId 1: not found")
}

func TestClaimNoInitialClaimError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[1], fmt.Sprintf("initial mission not completed: address %s, campaignId: 0: mission not completed yet", acountsAddresses[1].String()))
}

func TestClaimMissionCampaignHasEnded(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 1)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[1], fmt.Sprintf("campaign 0 has already ended (%s > endTime %s) error: campaign is disabled", testHelper.Context.BlockTime(), campaign.EndTime))
}

func TestClaimMissionWithTypeClaim(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
}

func TestClaimMissionAlreadyClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 80000001)
	delagationAmount := sdk.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[1], fmt.Sprintf("address %s, campaignId: 0, missionId: 1: mission already completed", acountsAddresses[1].String()))
}

func TestFullCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 1)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	mission.MissionType = cfeclaimtypes.MissionVote
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 60000001)

	delagationAmount := sdk.NewInt(1000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])

	testHelper.C4eClaimUtils.CompleteVoteMission(0, 2, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMission(0, 2, acountsAddresses[1])
}

func TestClaimMissionWithTypeClaimRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses[:10], 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)

	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[10], fmt.Sprintf("user claim entries not found for address %s: not found", acountsAddresses[10].String()))
}
