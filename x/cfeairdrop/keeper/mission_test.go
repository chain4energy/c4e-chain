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

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})
	testHelper.C4eAirdropUtils.CompleteMission(1, 2, acountsAddresses[0])
}

func TestCompleteMissionCamapignNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})
	testHelper.C4eAirdropUtils.CompleteMissionError(3, 2, acountsAddresses[0], "camapign not found: campaignId 3: not found")
}

func TestCompleteMissionCamapignDisabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	campaigns[0].Enabled = false
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[0], "campaign disabled - campaignId 1: campaignId 1: campaign is disabled")
}

func TestCompleteMissionCamapignNotStarted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	startTime := testHelper.Context.BlockTime().Add(10)
	campaigns[0].StartTime = &startTime
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("campaign disabled - campaignId 1: campaignId 1 not started: time %s < startTime %s: campaign is disabled", testHelper.Context.BlockTime(), campaigns[0].StartTime))
}

func TestCompleteMissionCamapignEnded(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	startTime := testHelper.Context.BlockTime().Add(-10000)
	campaigns[0].StartTime = &startTime
	endTime := testHelper.Context.BlockTime().Add(-1000)
	campaigns[0].EndTime = &endTime
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 2, acountsAddresses[0],
		fmt.Sprintf("campaign disabled - campaignId 1: campaignId 1 ended: time %s > endTime %s: campaign is disabled", testHelper.Context.BlockTime(), campaigns[0].EndTime))
}

func TestCompleteMissionMissionNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, 3, acountsAddresses[0], "mission not found - campaignId 1, missionId 3: not found")
}

func TestCompleteMissionClaimRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, uint64(types.DELEGATION), acountsAddresses[2],
		fmt.Sprintf("claim record not found for address %s: not found", acountsAddresses[2]))
}

func TestCompleteMissionCampaignRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	claimRecord := prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})
	claimRecord.GetCampaignRecords()[0].CampaignId = 2
	testHelper.C4eAirdropUtils.SetClaimRecord(claimRecord)

	testHelper.C4eAirdropUtils.CompleteMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		fmt.Sprintf("campaign record with id: 1 not found for address %s: not found", acountsAddresses[0]))
}

func TestCompleteMissionAlreadeCompleted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL), uint64(types.DELEGATION)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		fmt.Sprintf("mission already completed: address %s, campaignId: 1, missionId: 1: mission already completed", acountsAddresses[0]))
}

func TestCompleteMissionNoInitialClaim(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	_, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: []types.InitialClaim{}, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		"initial claim not found - campaignId 1: not found")
}

func TestCompleteMissionInitialMissionNotClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{})

	testHelper.C4eAirdropUtils.CompleteMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		fmt.Sprintf("initial mission not completed: address %s, campaignId: 1, missionId: 0: mission not completed yet", acountsAddresses[0]))
}

func TestClaimMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1, 2}, []uint64{1})

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[0], sdk.NewCoins(), 12312, 1555565657676576)

	testHelper.C4eAirdropUtils.ClaimMission(1, 2, acountsAddresses[0])

}

func TestClaimMissionCamapignNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})
	testHelper.C4eAirdropUtils.ClaimMissionError(3, uint64(types.DELEGATION), acountsAddresses[0], "camapign not found: campaignId 3: not found")
}

func TestClaimMissionCamapignDisabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	campaigns[0].Enabled = false
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, uint64(types.DELEGATION), acountsAddresses[0], "campaign disabled - campaignId 1: campaignId 1: campaign is disabled")
}

func TestClaimMissionCamapignNotStarted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	startTime := testHelper.Context.BlockTime().Add(10)
	campaigns[0].StartTime = &startTime
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		fmt.Sprintf("campaign disabled - campaignId 1: campaignId 1 not started: time %s < startTime %s: campaign is disabled", testHelper.Context.BlockTime(), campaigns[0].StartTime))
}

func TestClaimMissionCamapignEnded(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	startTime := testHelper.Context.BlockTime().Add(-10000)
	campaigns[0].StartTime = &startTime
	endTime := testHelper.Context.BlockTime().Add(-1000)
	campaigns[0].EndTime = &endTime
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		fmt.Sprintf("campaign disabled - campaignId 1: campaignId 1 ended: time %s > endTime %s: campaign is disabled", testHelper.Context.BlockTime(), campaigns[0].EndTime))
}

func TestClaimMissionMissionNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{1}, []uint64{1})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, 3, acountsAddresses[0], "mission not found - campaignId 1, missionId 3: not found")
}

func TestClaimMissionClaimRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, uint64(types.DELEGATION), acountsAddresses[2],
		fmt.Sprintf("claim record not found for address %s: not found", acountsAddresses[2]))
}

func TestClaimMissionCampaignRecordNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	claimRecord := prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})
	claimRecord.GetCampaignRecords()[0].CampaignId = 2
	testHelper.C4eAirdropUtils.SetClaimRecord(claimRecord)

	testHelper.C4eAirdropUtils.ClaimMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		fmt.Sprintf("campaign record with id: 1 not found for address %s: not found", acountsAddresses[0]))
}

func TestClaimMissionNotCompleted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		fmt.Sprintf("mission not completed: address %s, campaignId: 1, missionId: 1: mission not completed yet", acountsAddresses[0]))
}

func TestClaimMissionAlreadyClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL), uint64(types.DELEGATION)}, []uint64{uint64(types.INITIAL), uint64(types.DELEGATION)})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		fmt.Sprintf("mission already claimed: address %s, campaignId: 1, missionId: 1: mission already claimed", acountsAddresses[0]))
}

func TestClaimMissionAccountNotExists(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL), uint64(types.DELEGATION)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.ClaimMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		fmt.Sprintf("send to claiming address %s error: create airdrop account - account does not exist: %s: entity does not exist: failed to send coins", acountsAddresses[0], acountsAddresses[0]))
}

func TestClaimMissionToAnotherAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	claimRecord := prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL), uint64(types.DELEGATION)}, []uint64{uint64(types.INITIAL)})

	claimRecord.ClaimAddress = acountsAddresses[2].String()
	testHelper.C4eAirdropUtils.SetClaimRecord(claimRecord)

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[2], sdk.NewCoins(), 12312, 1555565657676576)

	testHelper.C4eAirdropUtils.ClaimMissionToAddress(1, uint64(types.DELEGATION), acountsAddresses[0], acountsAddresses[2])

}

func TestClaimMissionToWrongAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})

	claimRecord := prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL), uint64(types.DELEGATION)}, []uint64{uint64(types.INITIAL)})

	claimRecord.ClaimAddress = "wrongAddress"
	testHelper.C4eAirdropUtils.SetClaimRecord(claimRecord)

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[2], sdk.NewCoins(), 12312, 1555565657676576)

	testHelper.C4eAirdropUtils.ClaimMissionError(1, uint64(types.DELEGATION), acountsAddresses[0],
		fmt.Sprintf("wrong claiming address %s: decoding bech32 failed: string not all lowercase or all uppercase: failed to parse", claimRecord.ClaimAddress))
}

func TestCompleteDelegationMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})
	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[0], sdk.NewCoins(), 12312, 1555565657676576)
	delagationAmount := sdk.NewInt(1000000)
	testHelper.BankUtils.AddDefaultDenomCoinToAccount(delagationAmount, acountsAddresses[0])

	testHelper.C4eAirdropUtils.CompleteDelegationMission(1, acountsAddresses[0], delagationAmount)
}

func TestCompleteVoteMission(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[0], sdk.NewCoins(), 12312, 1555565657676576)

	delagationAmount := sdk.NewInt(1000000)
	testHelper.BankUtils.AddDefaultDenomCoinToAccount(delagationAmount, acountsAddresses[0])

	validators := testHelper.StakingUtils.GetValidators()
	valAddr, err := sdk.ValAddressFromBech32(validators[0].OperatorAddress)
	require.NoError(t, err)
	testHelper.StakingUtils.MessageDelegate(1, 0, valAddr, acountsAddresses[0], delagationAmount)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})
	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{uint64(types.INITIAL)}, []uint64{uint64(types.INITIAL)})

	testHelper.C4eAirdropUtils.CompleteVoteMission(1, acountsAddresses[0])

}

func TestFullCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	testHelper.C4eAirdropUtils.CreateAirdropAccout(acountsAddresses[0], sdk.NewCoins(), 12312, 1555565657676576)

	params, campaigns := prepareTestCampaign(testHelper.Context)
	initialClaims, missions := prepareMissions()
	testHelper.C4eAirdropUtils.InitGenesis(types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns})
	prepareClaimRecord(testHelper, acountsAddresses[1], acountsAddresses[0], []uint64{}, []uint64{})

	testHelper.C4eAirdropUtils.ClaimInitial(1, acountsAddresses[0])

	delagationAmount := sdk.NewInt(1000000)
	testHelper.BankUtils.AddDefaultDenomCoinToAccount(delagationAmount, acountsAddresses[0])

	testHelper.C4eAirdropUtils.CompleteDelegationMission(1, acountsAddresses[0], delagationAmount)

	testHelper.C4eAirdropUtils.CompleteVoteMission(1, acountsAddresses[0])

	testHelper.C4eAirdropUtils.ClaimMission(1, uint64(types.DELEGATION), acountsAddresses[0])

	testHelper.C4eAirdropUtils.ClaimMission(1, uint64(types.VOTE), acountsAddresses[0])

}

func createNMission(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mission {
	items := make([]types.Mission, n)
	for i := range items {
		items[i].CampaignId = uint64(i)
		items[i].MissionId = uint64(i)
		items[i].Weight = sdk.NewDec(int64(i))
		items[i].Description = fmt.Sprintf("desc %d", i)

		keeper.SetMission(ctx, items[i])
	}
	return items
}

func prepareTestCampaign(ctx sdk.Context) (types.Params, []*types.Campaign) {
	start := ctx.BlockTime()
	end := ctx.BlockTime().Add(1000)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	return types.Params{
			Denom: commontestutils.DefaultTestDenom},
		[]*types.Campaign{
			{
				Id:            1,
				Enabled:       true,
				StartTime:     &start,
				EndTime:       &end,
				LockupPeriod:  lockupPeriod,
				VestingPeriod: vestingPeriod,
				Description:   "test-campaign",
			},
		}
}

func prepareMissions() ([]types.InitialClaim, []types.Mission) {
	initialClaims := []types.InitialClaim{{CampaignId: 1, MissionId: uint64(types.INITIAL)}}
	missions := []types.Mission{
		{CampaignId: 1, MissionId: uint64(types.INITIAL), Description: "initial-mission", Weight: sdk.MustNewDecFromStr("0.1")},
		{CampaignId: 1, MissionId: uint64(types.DELEGATION), Description: "test-delegation-mission", Weight: sdk.MustNewDecFromStr("0.2")},
		{CampaignId: 1, MissionId: uint64(types.VOTE), Description: "test-vote-mission", Weight: sdk.MustNewDecFromStr("0.3")},
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
