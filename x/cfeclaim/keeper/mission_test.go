package keeper_test

import (
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMissionGet(t *testing.T) {
	k, ctx := keepertest.CfeclaimKeeper(t)
	items := createAndSaveNTestMissions(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetMission(ctx,
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

func TestAllCampaignMissionsRemove(t *testing.T) {
	k, ctx := keepertest.CfeclaimKeeper(t)
	items := createAndSaveNTestMissions(k, ctx, 10)
	for _, item := range items {
		k.RemoveAllMissionForCampaign(ctx,
			item.CampaignId,
		)
		_, found := k.GetMission(ctx,
			item.CampaignId,
			item.Id,
		)
		require.False(t, found)
	}
}

func TestMissionGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfeclaimKeeper(t)
	items := createAndSaveNTestMissions(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMission(ctx)),
	)
}

func TestAddMissionToCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
}

func TestAddManyMissionToCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
}

func TestAddMissionToCampaignDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaignError(acountsAddresses[0].String(), 1, mission, "campaign with id 1 not found: entity does not exist")
}

func TestAddMissionToCampaignWrongWeightError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.Weight = sdk.NewDec(-2)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaignError(acountsAddresses[0].String(), 0, mission, fmt.Sprintf("weight (%s) is not between 0 and 1 error: wrong param value", mission.Weight.String()))
	mission.Weight = sdk.NewDec(2)
	testHelper.C4eClaimUtils.AddMissionToCampaignError(acountsAddresses[0].String(), 0, mission, fmt.Sprintf("weight (%s) is not between 0 and 1 error: wrong param value", mission.Weight.String()))

	mission.Weight = sdk.MustNewDecFromStr("0.6")
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.AddMissionToCampaignError(acountsAddresses[0].String(), 0, mission, fmt.Sprintf("all campaign missions weight sum is >= 1 (%s > 1) error: wrong param value", mission.Weight.Mul(sdk.NewDec(2)).String()))
}

func TestAddMissionToCampaignEmptyName(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.Name = ""
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaignError(acountsAddresses[0].String(), 0, mission, "empty name error: wrong param value")
}

func TestAddMissionToCampaignEmptyDescription(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	mission.Description = ""
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaignError(acountsAddresses[0].String(), 0, mission, "mission empty description error: wrong param value")
}

func TestAddMissionToCampaignWrongOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaignError(acountsAddresses[1].String(), 0, mission, "you are not the campaign owner: wrong transaction signer")
}

func TestAddMissionToCampaignAlreadyEnabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.AddMissionToCampaignError(acountsAddresses[0].String(), 0, mission, "campaign is enabled")
}

func TestAddMissionToCampaignAlreadyOver(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaign(acountsAddresses[0].String(), 0)
	testHelper.C4eClaimUtils.AddMissionToCampaignError(acountsAddresses[0].String(), 0, mission, fmt.Sprintf("campaign with id 0 campaign is over (end time - %s < %s): wrong param value", campaign.EndTime, blockTime))
}

func createAndSaveNTestMissions(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mission {
	items := make([]types.Mission, n)
	for i := range items {
		weight := sdk.NewDec(int64(i))
		items[i].CampaignId = uint64(i)
		items[i].Id = uint64(i)
		items[i].Weight = weight
		items[i].Description = fmt.Sprintf("desc %d", i)

		keeper.SetMission(ctx, items[i])
	}
	return items
}

func prepareTestMission() types.Mission {
	delegationMissionWeight := sdk.MustNewDecFromStr("0.2")
	return types.Mission{
		CampaignId:  0,
		Name:        "Name",
		MissionType: types.MissionDelegate,
		Description: "test-delegation-mission",
		Weight:      delegationMissionWeight,
	}
}
