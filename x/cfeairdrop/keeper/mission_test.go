package keeper_test

import (
	"fmt"
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

func TestMissionGet(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNMission(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMission(ctx,
			item.CampaignId,
			item.MissionId,
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
			item.MissionId,
		)
		_, found := keeper.GetMission(ctx,
			item.CampaignId,
			item.MissionId,
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

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})
	testHelper.C4eAirdropUtils.CompleteMission(1, 2, acountsAddresses[0])
}

func TestCompleteMissionCamapignNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})
	testHelper.C4eAirdropUtils.CompleteMissionError(3, 2, acountsAddresses[0], "camapign not found: campaignId 3: not found")
}

func TestCompleteMissionCamapignDisabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	params.Campaigns[0].Enabled = false
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[0], "campaign disabled - campaignId 1: campaignId 1: campaign is disabled")
}

func TestCompleteMissionCamapignNotStarted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	params.Campaigns[0].StartTime = testHelper.Context.BlockTime().Add(10)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("campaign disabled - campaignId 1: campaignId 1 not started: time %s < startTime %s: campaign is disabled", testHelper.Context.BlockTime(), params.Campaigns[0].StartTime))
}

func TestCompleteMissionCamapignEnded(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	params.Campaigns[0].StartTime = testHelper.Context.BlockTime().Add(-10000)
	endTime := testHelper.Context.BlockTime().Add(-1000)
	params.Campaigns[0].EndTime = &endTime
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("campaign disabled - campaignId 1: campaignId 1 ended: time %s > endTime %s: campaign is disabled", testHelper.Context.BlockTime(), params.Campaigns[0].EndTime))
}

func TestCompleteMissionMissionNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 3, acountsAddresses[0], "mission not found - campaignId 1, missionId 3: not found")
}

func TestCompleteMissionClaimRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[2],
		fmt.Sprintf("claim record not found for address %s: not found", acountsAddresses[2]))
}

func TestCompleteMissionCampaignRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	claimRecord := prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})
	claimRecord.GetCampaignRecords()[0].CampaignId = 2
	testHelper.C4eAirdropUtils.SetClaimRecord(claimRecord)

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("campaign record with id: 1 not found for address %s: not found", acountsAddresses[0]))
}

func TestCompleteMissionAlreadeCompleted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1, 2}, []uint64{1})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("mission already completed: address %s, campaignId: 1, missionId: 2: mission already completed", acountsAddresses[0]))
}

func TestCompleteMissionNoInitialClaim(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	_, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: []types.InitialClaim{}, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[0],
		"initial claim not found - campaignId 1: not found")
}

func TestCompleteMissionInitialMissionNotClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("initial mission not completed: address %s, campaignId: 1, missionId: 1: mission not completed yet", acountsAddresses[0]))
}

func TestClaimMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1, 2}, []uint64{1})

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[0], sdk.NewCoins(), 12312, 1555565657676576)

	testHelper.C4eAirdropUtils.ClaimMission(1, 2, acountsAddresses[0])

}

func TestClaimMissionCamapignNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})
	testHelper.C4eAirdropUtils.ClaimMissionError(3, 2, acountsAddresses[0], "camapign not found: campaignId 3: not found")
}

func TestClaimMissionCamapignDisabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	params.Campaigns[0].Enabled = false
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 2, acountsAddresses[0], "campaign disabled - campaignId 1: campaignId 1: campaign is disabled")
}

func TestClaimMissionCamapignNotStarted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	params.Campaigns[0].StartTime = testHelper.Context.BlockTime().Add(10)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("campaign disabled - campaignId 1: campaignId 1 not started: time %s < startTime %s: campaign is disabled", testHelper.Context.BlockTime(), params.Campaigns[0].StartTime))
}

func TestClaimMissionCamapignEnded(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	params.Campaigns[0].StartTime = testHelper.Context.BlockTime().Add(-10000)
	endTime := testHelper.Context.BlockTime().Add(-1000)
	params.Campaigns[0].EndTime = &endTime
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("campaign disabled - campaignId 1: campaignId 1 ended: time %s > endTime %s: campaign is disabled", testHelper.Context.BlockTime(), params.Campaigns[0].EndTime))
}

func TestClaimMissionMissionNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 3, acountsAddresses[0], "mission not found - campaignId 1, missionId 3: not found")
}

func TestClaimMissionClaimRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 2, acountsAddresses[2],
		fmt.Sprintf("claim record not found for address %s: not found", acountsAddresses[2]))
}

func TestClaimMissionCampaignRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	claimRecord := prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})
	claimRecord.GetCampaignRecords()[0].CampaignId = 2
	testHelper.C4eAirdropUtils.SetClaimRecord(claimRecord)

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("campaign record with id: 1 not found for address %s: not found", acountsAddresses[0]))
}

func TestClaimMissionNotCompleted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("mission not completed: address %s, campaignId: 1, missionId: 2: mission not completed yet", acountsAddresses[0]))
}

func TestClaimMissionAlreadyClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1, 2}, []uint64{1, 2})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("mission already claimed: address %s, campaignId: 1, missionId: 2: mission already claimed", acountsAddresses[0]))
}

func TestClaimMissionAccountNotExists(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1, 2}, []uint64{1})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("send to claiming address %s error: create airdrop account - account does not exist: %s: entity does not exist: failed to send coins", acountsAddresses[0], acountsAddresses[0]))
}

func TestClaimMissionToAnotherAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	claimRecord := prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1, 2}, []uint64{1})

	claimRecord.ClaimAddress = acountsAddresses[2].String()
	testHelper.C4eAirdropUtils.SetClaimRecord(claimRecord)

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[2], sdk.NewCoins(), 12312, 1555565657676576)

	testHelper.C4eAirdropUtils.ClaimMissionToAddress(1, 2, acountsAddresses[0], acountsAddresses[2])

}

func TestClaimMissionToWrongAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	params := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions})

	claimRecord := prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1, 2}, []uint64{1})

	claimRecord.ClaimAddress = "wrongAddress"
	testHelper.C4eAirdropUtils.SetClaimRecord(claimRecord)

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[2], sdk.NewCoins(), 12312, 1555565657676576)

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("wrong claiming address %s: decoding bech32 failed: string not all lowercase or all uppercase: failed to parse", claimRecord.ClaimAddress))
}

func createNMission(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mission {
	items := make([]types.Mission, n)
	for i := range items {
		items[i].CampaignId = uint64(i)
		items[i].MissionId = uint64(i)

		keeper.SetMission(ctx, items[i])
	}
	return items
}

func prepareTestCampaign(ctx sdk.Context) types.Params {
	end := ctx.BlockTime().Add(1000)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	return types.Params{
		Denom: commontestutils.DefaultTestDenom,
		Campaigns: []*types.Campaign{
			{
				CampaignId:    1,
				Enabled:       true,
				StartTime:     ctx.BlockTime(),
				EndTime:       &end,
				LockupPeriod:  lockupPeriod,
				VestingPeriod: vestingPeriod,
				Description:   "test-campaign",
			},
		}}
}

func prepareMissions() ([]types.InitialClaim, []types.Mission) {
	initialClaims := []types.InitialClaim{{CampaignId: 1, MissionId: 1}}
	missions := []types.Mission{
		{CampaignId: 1, MissionId: 1, Description: "initial-mission", Weight: sdk.MustNewDecFromStr("0.1")},
		{CampaignId: 1, MissionId: 2, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.2")},
	}
	return initialClaims, missions
}

func prepareClaimRecord(testHelper *testapp.TestHelper, srcAddress sdk.AccAddress, recordAddress sdk.AccAddress,
	completedMissions []uint64, claimedMissions []uint64) *types.ClaimRecord {
	records := map[string]sdk.Int{recordAddress.String(): sdk.NewInt(10000)}
	testHelper.C4eAirdropUtils.AddCampaignRecords(srcAddress, 1, records)
	claimRecord := testHelper.C4eAirdropUtils.GetClaimRecord(recordAddress.String())
	claimRecord.GetCampaignRecords()[0].ClaimedMissions = claimedMissions
	claimRecord.GetCampaignRecords()[0].CompletedMissions = completedMissions
	testHelper.C4eAirdropUtils.SetClaimRecord(claimRecord)
	return claimRecord
}
