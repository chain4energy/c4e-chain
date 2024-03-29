package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCompleteDelegationMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 1)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(2, 0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestClaimToDefineMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 1)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionToDefine
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)

	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[1], "cannot claim mission with type TO_DEFINE: mission claiming error")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestClaimToDefineMissionMultipleMissions(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 1)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionToDefine
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	mission.MissionType = cfeclaimtypes.MissionClaim
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)

	testHelper.C4eClaimUtils.ClaimMission(0, 2, acountsAddresses[1])
	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[1], "cannot claim mission with type TO_DEFINE: mission claiming error")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestCompleteDelegationMissionDelegateTwiceToTheSameValidator(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 1)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(2, 0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 1, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 1, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 1, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 1)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(3, 1, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])
	testHelper.C4eClaimUtils.ClaimMission(1, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestCompleteDelegationMissionDelegateTwiceToDifferentValidator(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 3)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(2, 0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 1, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 1, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 1, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 1)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(4, 1, 1, acountsAddresses[1], delagationAmount, validatorAddresses[1])
	testHelper.C4eClaimUtils.ClaimMission(1, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestCompleteDelegationMissionTwiceUndelegateAll(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 3)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(2, 0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 1, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 1, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 1, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 1)
	testHelper.StakingUtils.MessageUndelegate(3, 0, validatorAddresses[0], acountsAddresses[1], delagationAmount)
	testHelper.C4eClaimUtils.ClaimMissionError(1, 1, acountsAddresses[1], "campaignId: 1, missionId: 1: mission not completed yet: mission claiming error")

	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestCompleteDelegationMissionTwiceUndelegatePartial(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 3)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(2, 0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 1, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 1, nil, nil)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 1, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 1)
	testHelper.StakingUtils.MessageUndelegate(3, 0, validatorAddresses[0], acountsAddresses[1], delagationAmount.SubRaw(1))
	testHelper.C4eClaimUtils.ClaimMission(1, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestCompleteVoteMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionVote
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)

	testHelper.C4eClaimUtils.CompleteVoteMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestClaimMissionDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)

	testHelper.C4eClaimUtils.ClaimMissionError(0, 2, acountsAddresses[1], "mission not found - campaignId 0, missionId 2: not found")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestClaimCampaignDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)

	testHelper.C4eClaimUtils.ClaimMissionError(1, 0, acountsAddresses[1], "campaign with id 1 not found: not found")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestClaimNoInitialClaimError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)

	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[1], fmt.Sprintf("initial mission not completed: address %s, campaignId: 0: mission not completed yet", acountsAddresses[1].String()))
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestClaimMissionCampaignHasEnded(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 1)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(2, 0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[1], fmt.Sprintf("campaign 0 has already ended (%s > endTime %s): campaign is disabled", testHelper.Context.BlockTime(), campaign.EndTime))
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestInitialClaimMissionInititalClaimAmountBiggerThanInititalClaimAMount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = math.NewInt(100)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestInitialClaimMissionInititalClaimAmountBiggerThanInititalClaimAmountAndFree(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = math.NewInt(100)
	campaign.Free = sdk.MustNewDecFromStr("0.2")
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestClaimMissionWithTypeClaim(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestInitialClaim0Amount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	mission.Weight = sdk.NewDec(1)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestInitialClaim0AmountAndInititalClaimFreeAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = sdk.NewInt(100000000)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	mission.Weight = sdk.NewDec(1)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestInitialClaim0AmountAndInititalClaimFreeAmountAndFree(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.Free = sdk.MustNewDecFromStr("0.5")
	campaign.InitialClaimFreeAmount = sdk.NewInt(100000000)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	mission.Weight = sdk.NewDec(1)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestInitialClaimFree100(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.Free = sdk.NewDec(1)
	campaign.InitialClaimFreeAmount = sdk.NewInt(100)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	mission.Weight = sdk.MustNewDecFromStr("0.5")
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestClaimMissionAlreadyClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[1], "campaignId: 0, missionId: 1: mission already completed")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestFullCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 1)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	mission.MissionType = cfeclaimtypes.MissionVote
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)

	delagationAmount := math.NewInt(1000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(2, 0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])

	testHelper.C4eClaimUtils.CompleteVoteMission(0, 2, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])

	testHelper.C4eClaimUtils.ClaimMission(0, 2, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestClaimMissionWithTypeClaimRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses[:10], 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)

	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)

	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[10], fmt.Sprintf("userEntry %s doesn't exist: not found", acountsAddresses[10].String()))
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestVestingPoolCampaignClaimMissionClaim(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, _ := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	campaign := prepareTestCampaign(testHelper.Context)
	ownerAddress := acountsAddresses[0]

	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	campaign.CampaignType = cfeclaimtypes.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestVestingPoolCampaignClaimMissionClaimOptionalClaimStartDate(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, _ := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	campaign := prepareTestCampaign(testHelper.Context)
	ownerAddress := acountsAddresses[0]

	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	campaign.CampaignType = cfeclaimtypes.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	claimStartDate := campaign.StartTime.Add(time.Minute)
	mission.ClaimStartDate = &claimStartDate
	campaign.EndTime = claimStartDate.Add(time.Minute)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	testHelper.C4eClaimUtils.ClaimMissionError(0, 1, acountsAddresses[1],
		fmt.Sprintf("mission 1 not started yet (blocktime %s < mission start time %s): mission is disabled", testHelper.Context.BlockTime(), claimStartDate))
	testHelper.SetContextBlockTime(claimStartDate)
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestVestingPoolCampaignClaimMissionVote(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, _ := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	campaign := prepareTestCampaign(testHelper.Context)
	ownerAddress := acountsAddresses[0]

	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	campaign.CampaignType = cfeclaimtypes.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionVote
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	testHelper.C4eClaimUtils.CompleteVoteMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestVestingPoolCampaignManyCoins(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, _ := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	campaign := prepareTestCampaign(testHelper.Context)
	ownerAddress := acountsAddresses[0]

	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	campaign.CampaignType = cfeclaimtypes.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionVote
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	testHelper.C4eClaimUtils.CompleteVoteMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestVestingPoolCampaignEverythingClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	claimEntries, _ := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	campaign := prepareTestCampaign(testHelper.Context)
	ownerAddress := acountsAddresses[0]

	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	campaign.CampaignType = cfeclaimtypes.VestingPoolCampaign
	campaign.VestingPoolName = vPool1

	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)

	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[0], 0)
	testHelper.SetContextBlockTime(campaign.EndTime.Add(time.Minute))
	testHelper.C4eClaimUtils.CloseCampaign(ownerAddress.String(), 0)

	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestVestingPoolCampaignClaimMissionDelegate(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorAddresses := testcosmos.CreateAccounts(11, 1)
	claimEntries, _ := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	campaign := prepareTestCampaign(testHelper.Context)
	ownerAddress := acountsAddresses[0]

	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	campaign.CampaignType = cfeclaimtypes.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionDelegate
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	delagationAmount := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eClaimUtils.CompleteDelegationMission(2, 0, 1, acountsAddresses[1], delagationAmount, validatorAddresses[0])
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestVestingPoolCampaignClaimWrongAccountType(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(11, 0)
	claimEntries, _ := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	campaign := prepareTestCampaign(testHelper.Context)
	ownerAddress := acountsAddresses[0]

	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	campaign.CampaignType = cfeclaimtypes.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	mission := prepareTestMission()
	mission.MissionType = cfeclaimtypes.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	err := testHelper.AuthUtils.CreateVestingAccount(acountsAddresses[1].String(), sdk.NewCoins(), time.Now(), time.Now().Add(time.Hour))
	require.NoError(t, err)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.ClaimInitialError(acountsAddresses[1], 0, "account already exists and is not of PeriodicContinuousVestingAccount nor BaseAccount type, got: *types.ContinuousVestingAccount: invalid account type")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}
