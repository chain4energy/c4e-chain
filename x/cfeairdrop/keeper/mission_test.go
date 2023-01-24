package keeper_test

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/testutil/module/cfeairdrop"
	"strconv"
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

const IntialMissionId = 0
const DelegateMissionId = 1
const VoteMissionId = 2

func TestMissionGet(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNMission(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMission(ctx,
			item.CampaignId,
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMissionRemove(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNMission(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMission(ctx,
			item.CampaignId,
			item.Id,
		)
		_, found := keeper.GetMission(ctx,
			item.CampaignId,
			item.Id,
		)
		require.False(t, found)
	}
}

func TestMissionGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNMission(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMission(ctx)),
	)
}

func TestCompleteMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)
	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0}, []uint64{0})

	testHelper.C4eAirdropUtils.CompleteMission(0, 1, acountsAddresses[1])
}

func TestCompleteMissionCamapignNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0}, []uint64{0})
	testHelper.C4eAirdropUtils.CompleteMissionError(3, 2, acountsAddresses[1], "camapign not found: campaignId 3: not found")
}

func TestCompleteMissionCamapignDisabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	campaigns[0].Enabled = false
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0}, []uint64{0})

	testHelper.C4eAirdropUtils.CompleteMissionError(0, 2, acountsAddresses[1], "campaign disabled - campaignId 0: campaignId 0: campaign is disabled")
}

func TestCompleteMissionCamapignNotStarted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	startTime := testHelper.Context.BlockTime().Add(10)
	campaigns[0].StartTime = &startTime
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0}, []uint64{0})

	testHelper.C4eAirdropUtils.CompleteMissionError(0, 2, acountsAddresses[1],
		fmt.Sprintf("campaign disabled - campaignId 0: campaignId 0 not started: time %s < startTime %s: campaign is disabled", testHelper.Context.BlockTime(), campaigns[0].StartTime))
}

func TestCompleteMissionCamapignEnded(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	campaigns[0].Owner = acountsAddresses[0].String()
	startTime := testHelper.Context.BlockTime().Add(-10000)
	campaigns[0].StartTime = &startTime
	endTime := testHelper.Context.BlockTime().Add(-1000)
	campaigns[0].EndTime = &endTime
	missions := prepareMissions()
	genesisState := types.GenesisState{Missions: missions, Campaigns: campaigns}
	testHelper.C4eAirdropUtils.InitGenesis(genesisState)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0}, []uint64{0})

	testHelper.C4eAirdropUtils.CompleteMissionError(0, 2, acountsAddresses[1],
		fmt.Sprintf("campaign disabled - campaignId 0: campaignId 0 ended: time %s > endTime %s: campaign is disabled", testHelper.Context.BlockTime(), campaigns[0].EndTime))
}

func TestCompleteMissionMissionNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0}, []uint64{0})

	testHelper.C4eAirdropUtils.CompleteMissionError(0, 4, acountsAddresses[1], "mission not found - campaignId 0, missionId 4: not found")
}

func TestCompleteMissionClaimRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0}, []uint64{0})

	testHelper.C4eAirdropUtils.CompleteMissionError(0, uint64(types.MissionDelegation), acountsAddresses[2],
		fmt.Sprintf("claim record not found for address %s: not found", acountsAddresses[2]))
}

func TestCompleteMissionCampaignRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	userAirdropEntries := prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0}, []uint64{0})
	userAirdropEntries.GetAirdropEntries()[0].CampaignId = 2
	testHelper.C4eAirdropUtils.SetUserAirdropEntries(userAirdropEntries)

	testHelper.C4eAirdropUtils.CompleteMissionError(0, 1, acountsAddresses[1],
		fmt.Sprintf("campaign record with id: 0 not found for address %s: not found", acountsAddresses[1]))
}

func TestCompleteMissionAlreadeCompleted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0, 1}, []uint64{0})

	testHelper.C4eAirdropUtils.CompleteMissionError(0, 1, acountsAddresses[1],
		fmt.Sprintf("mission already completed: address %s, campaignId: 0, missionId: 1: mission already completed", acountsAddresses[1]))
}

func TestCompleteMissionNoInitialClaim(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim)}, []uint64{uint64(types.MissionInitialClaim)})

	testHelper.C4eAirdropUtils.CompleteMissionError(0, uint64(types.MissionDelegation), acountsAddresses[1],
		fmt.Sprintf("initial mission not completed: address %s, campaignId: 0, missionId: 0: mission not completed yet", acountsAddresses[1]))
}

func TestCompleteMissionInitialMissionNotClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0}, []uint64{})

	testHelper.C4eAirdropUtils.CompleteMissionError(0, 1, acountsAddresses[1],
		fmt.Sprintf("initial mission not completed: address %s, campaignId: 0, missionId: 0: mission not completed yet", acountsAddresses[1]))
}

func TestClaimMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{1, 2}, []uint64{1})

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[1], sdk.NewCoins(), 12312, 1555565657676576)

	testHelper.C4eAirdropUtils.ClaimMission(0, 2, acountsAddresses[1])

}

func TestClaimMissionCamapignNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim)}, []uint64{uint64(types.MissionInitialClaim)})
	testHelper.C4eAirdropUtils.ClaimMissionError(3, uint64(types.MissionDelegation), acountsAddresses[1], "camapign not found: campaignId 3: not found")
}

func TestClaimMissionCamapignDisabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	campaigns[0].Enabled = false
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim)}, []uint64{uint64(types.MissionInitialClaim)})

	testHelper.C4eAirdropUtils.ClaimMissionError(0, uint64(types.MissionDelegation), acountsAddresses[1], "campaign disabled - campaignId 0: campaignId 0: campaign is disabled")
}

func TestClaimMissionCamapignNotStarted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	startTime := testHelper.Context.BlockTime().Add(10)
	campaigns[0].StartTime = &startTime
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim)}, []uint64{uint64(types.MissionInitialClaim)})

	testHelper.C4eAirdropUtils.ClaimMissionError(0, uint64(types.MissionDelegation), acountsAddresses[1],
		fmt.Sprintf("campaign disabled - campaignId 0: campaignId 0 not started: time %s < startTime %s: campaign is disabled", testHelper.Context.BlockTime(), campaigns[0].StartTime))
}

func TestClaimMissionCamapignEnded(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	campaigns[0].Owner = acountsAddresses[0].String()
	startTime := testHelper.Context.BlockTime().Add(-10000)
	campaigns[0].StartTime = &startTime
	endTime := testHelper.Context.BlockTime().Add(-1000)
	campaigns[0].EndTime = &endTime
	missions := prepareMissions()
	genesisState := types.GenesisState{Missions: missions, Campaigns: campaigns}
	testHelper.C4eAirdropUtils.InitGenesis(genesisState)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim)}, []uint64{uint64(types.MissionInitialClaim)})

	testHelper.C4eAirdropUtils.ClaimMissionError(0, uint64(types.MissionDelegation), acountsAddresses[1],
		fmt.Sprintf("campaign disabled - campaignId 0: campaignId 0 ended: time %s > endTime %s: campaign is disabled", testHelper.Context.BlockTime(), campaigns[0].EndTime))
}

func TestClaimMissionMissionNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.ClaimMissionError(0, 4, acountsAddresses[1], "mission not found - campaignId 0, missionId 4: not found")
}

func TestClaimMissionClaimRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim)}, []uint64{uint64(types.MissionInitialClaim)})

	testHelper.C4eAirdropUtils.ClaimMissionError(0, uint64(types.MissionDelegation), acountsAddresses[2],
		fmt.Sprintf("claim record not found for address %s: not found", acountsAddresses[2]))
}

func TestClaimMissionCampaignRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	userAirdropEntries := prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim)}, []uint64{uint64(types.MissionInitialClaim)})
	userAirdropEntries.GetAirdropEntries()[0].CampaignId = 2
	testHelper.C4eAirdropUtils.SetUserAirdropEntries(userAirdropEntries)

	testHelper.C4eAirdropUtils.ClaimMissionError(0, uint64(types.MissionDelegation), acountsAddresses[1],
		fmt.Sprintf("campaign record with id: 0 not found for address %s: not found", acountsAddresses[1]))
}

func TestClaimMissionNotCompleted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim)}, []uint64{uint64(types.MissionInitialClaim)})

	testHelper.C4eAirdropUtils.ClaimMissionError(0, uint64(types.MissionDelegation), acountsAddresses[1],
		fmt.Sprintf("mission not completed: address %s, campaignId: 0, missionId: 2: mission not completed yet", acountsAddresses[1]))
}

func TestClaimMissionAlreadyClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim), uint64(types.MissionDelegation)}, []uint64{uint64(types.MissionInitialClaim), uint64(types.MissionDelegation)})

	testHelper.C4eAirdropUtils.ClaimMissionError(0, uint64(types.MissionDelegation), acountsAddresses[1],
		fmt.Sprintf("mission already claimed: address %s, campaignId: 0, missionId: 2: mission already claimed", acountsAddresses[1]))
}

func TestClaimMissionAccountNotExists(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim), uint64(types.MissionDelegation)}, []uint64{uint64(types.MissionInitialClaim)})

	testHelper.C4eAirdropUtils.ClaimMissionError(0, uint64(types.MissionDelegation), acountsAddresses[1],
		fmt.Sprintf("send to claiming address %s error: create airdrop account - account does not exist: %s: entity does not exist: failed to send coins", acountsAddresses[1], acountsAddresses[1]))
}

func TestClaimMissionToAnotherAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	userAirdropEntries := prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim), uint64(types.MissionDelegation)}, []uint64{uint64(types.MissionInitialClaim)})

	userAirdropEntries.ClaimAddress = acountsAddresses[2].String()
	testHelper.C4eAirdropUtils.SetUserAirdropEntries(userAirdropEntries)

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[2], sdk.NewCoins(), 12312, 1555565657676576)

	testHelper.C4eAirdropUtils.ClaimMissionToAddress(0, uint64(types.MissionDelegation), acountsAddresses[1], acountsAddresses[2])

}

func TestClaimMissionToWrongAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	userAirdropEntries := prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{uint64(types.MissionInitialClaim), uint64(types.MissionDelegation)}, []uint64{uint64(types.MissionInitialClaim)})

	userAirdropEntries.ClaimAddress = "wrongAddress"
	testHelper.C4eAirdropUtils.SetUserAirdropEntries(userAirdropEntries)

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[2], sdk.NewCoins(), 12312, 1555565657676576)

	testHelper.C4eAirdropUtils.ClaimMissionError(0, uint64(types.MissionDelegation), acountsAddresses[1],
		fmt.Sprintf("send to claiming address %s error: wrong claiming address %s: decoding bech32 failed: string not all lowercase or all uppercase: failed to parse: failed to send coins", userAirdropEntries.ClaimAddress, userAirdropEntries.ClaimAddress))
}

func TestCompleteDelegationMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{0}, []uint64{0})

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[1], sdk.NewCoins(), 12312, 1555565657676576)
	delagationAmount := sdk.NewInt(1000000)
	testHelper.BankUtils.AddDefaultDenomCoinToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.CompleteDelegationMission(0, acountsAddresses[1], delagationAmount)
}

func TestCompleteVoteMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[1], sdk.NewCoins(), 12312, 1555565657676576)

	delagationAmount := sdk.NewInt(1000000)
	testHelper.BankUtils.AddDefaultDenomCoinToAccount(delagationAmount, acountsAddresses[1])

	validators := testHelper.StakingUtils.GetValidators()
	valAddr, err := sdk.ValAddressFromBech32(validators[0].OperatorAddress)
	require.NoError(t, err)
	testHelper.StakingUtils.MessageDelegate(1, 0, valAddr, acountsAddresses[1], delagationAmount)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{types.InitialMissionId}, []uint64{types.InitialMissionId})

	testHelper.C4eAirdropUtils.CompleteVoteMission(0, acountsAddresses[1])

}

func TestFullCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[1], sdk.NewCoins(), 12312, 1555565657676576)

	campaigns := prepareTestCampaign(testHelper.Context)
	missions := prepareMissions()
	addCampaignsAndMissions(testHelper.C4eAirdropUtils, acountsAddresses[0].String(), campaigns, missions)

	prepareUserAirdropEntries(testHelper, acountsAddresses[0], acountsAddresses[1], []uint64{}, []uint64{})

	testHelper.C4eAirdropUtils.ClaimInitial(0, acountsAddresses[1], 500000000)

	delagationAmount := sdk.NewInt(1000000)
	testHelper.BankUtils.AddDefaultDenomCoinToAccount(delagationAmount, acountsAddresses[1])

	testHelper.C4eAirdropUtils.CompleteDelegationMission(0, acountsAddresses[1], delagationAmount)

	testHelper.C4eAirdropUtils.CompleteVoteMission(0, acountsAddresses[1])

	testHelper.C4eAirdropUtils.ClaimMission(0, 1, acountsAddresses[1])

	testHelper.C4eAirdropUtils.ClaimMission(0, 2, acountsAddresses[1])

}

func createNMission(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mission {
	items := make([]types.Mission, n)
	for i := range items {
		weight := sdk.NewDec(int64(i))
		items[i].CampaignId = uint64(i)
		items[i].Id = uint64(i)
		items[i].Weight = &weight
		items[i].Description = fmt.Sprintf("desc %d", i)

		keeper.SetMission(ctx, items[i])
	}
	return items
}

func prepareTestCampaign(ctx sdk.Context) []types.Campaign {
	start := ctx.BlockTime()
	end := ctx.BlockTime().Add(time.Second * 10)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	return []types.Campaign{
		{
			Id:            0,
			Name:          "Name",
			Description:   "test-campaign",
			Enabled:       true,
			StartTime:     &start,
			EndTime:       &end,
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
		},
	}
}

func prepareMissions() []types.Mission {
	initialMissionWeight := sdk.MustNewDecFromStr("0")
	delegationMissionWeight := sdk.MustNewDecFromStr("0.2")
	voteMissionWeight := sdk.MustNewDecFromStr("0.3")
	missions := []types.Mission{
		{
			CampaignId:  0,
			Name:        "Name",
			Description: "initial-mission",
			MissionType: types.MissionInitialClaim,
			Weight:      &initialMissionWeight,
		},
		{
			CampaignId:  0,
			Name:        "Name",
			MissionType: types.MissionDelegation,
			Description: "test-delegation-mission",
			Weight:      &delegationMissionWeight,
		},
		{
			CampaignId:  0,
			Name:        "Name",
			MissionType: types.MissionVote,
			Description: "test-vote-mission",
			Weight:      &voteMissionWeight,
		},
	}
	return missions
}

func prepareUserAirdropEntries(testHelper *testapp.TestHelper, srcAddress sdk.AccAddress, recordAddress sdk.AccAddress,
	completedMissions []uint64, claimedMissions []uint64) *types.UserAirdropEntries {
	testHelper.C4eAirdropUtils.AddAirdropEntries(srcAddress, 0, prepareAidropEntries(recordAddress.String()))
	userAirdropEntries := testHelper.C4eAirdropUtils.GetUserAirdropEntries(recordAddress.String())
	userAirdropEntries.ClaimAddress = userAirdropEntries.Address
	userAirdropEntries.GetAirdropEntries()[0].ClaimedMissions = claimedMissions
	userAirdropEntries.GetAirdropEntries()[0].CompletedMissions = completedMissions
	testHelper.C4eAirdropUtils.SetUserAirdropEntries(userAirdropEntries)
	return userAirdropEntries
}

func prepareAidropEntries(address string) []*types.AirdropEntry {
	airdropEntries := []*types.AirdropEntry{
		{
			AirdropCoins: sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, sdk.NewInt(1000000000))),
			Address:      address,
		},
	}

	return airdropEntries
}

func addCampaignsAndMissions(utils *cfeairdrop.ContextC4eAirdropUtils, ownerAddress string, campaigns []types.Campaign, missions []types.Mission) {
	for _, campaign := range campaigns {
		utils.CreateAirdropCampaign(ownerAddress, campaign.Name, campaign.Description, campaign.AllowFeegrant, campaign.InitialClaimFreeAmount, *campaign.StartTime, *campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)
		if campaign.Enabled == true {
			utils.StartAirdropCampaign(ownerAddress, campaign.Id)
		}
	}
	for _, mission := range missions {
		utils.AddMissionToAirdropCampaign(ownerAddress, mission.CampaignId, mission.Name, mission.Description, mission.MissionType, *mission.Weight)
	}
}
