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

func createNInitialClaim(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.InitialClaim {
	items := make([]types.InitialClaim, n)
	for i := range items {
		items[i].CampaignId = uint64(i)
		items[i].MissionId = uint64(1000 + i)
		keeper.SetInitialClaim(ctx, items[i])
	}
	return items
}

func TestInitialClaimGet(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNInitialClaim(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetInitialClaim(ctx,
			item.CampaignId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestInitialClaimRemove(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNInitialClaim(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveInitialClaim(ctx,
			item.CampaignId,
		)
		_, found := keeper.GetInitialClaim(ctx,
			item.CampaignId,
		)
		require.False(t, found)
	}
}

func TestInitialClaimGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNInitialClaim(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllInitialClaim(ctx)),
	)
}

func TestClaimInitial(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	// ctx := testHelper.Context

	end := testHelper.Context.BlockTime().Add(1000)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	params := types.Params{Denom: commontestutils.DefaultTestDenom, Campaigns: []*types.Campaign{
		{
			CampaignId:    1,
			Enabled:       true,
			StartTime:     testHelper.Context.BlockTime(),
			EndTime:       &end,
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
			Description:   "test-campaign",
		},
	}}
	initialClaims := []types.InitialClaim{{CampaignId: 1, MissionId: 3}}
	missions := []types.Mission{{CampaignId: 1, MissionId: 3, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.2")}}
	genesisState := types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions}
	testHelper.C4eAirdropUtils.InitGenesis(genesisState)

	records := map[string]sdk.Int{acountsAddresses[0].String(): sdk.NewInt(10000)}
	testHelper.C4eAirdropUtils.AddCampaignRecords(acountsAddresses[1], 1, records)

	testHelper.C4eAirdropUtils.ClaimInitial(1, acountsAddresses[0])

}

func TestClaimInitialCampaignNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	end := testHelper.Context.BlockTime().Add(1000)
	params := types.Params{Denom: commontestutils.DefaultTestDenom, Campaigns: []*types.Campaign{
		{
			CampaignId:    1,
			Enabled:       true,
			StartTime:     testHelper.Context.BlockTime(),
			EndTime:       &end,
			LockupPeriod:  1000,
			VestingPeriod: 2000,
			Description:   "test-campaign",
		},
	}}
	initialClaims := []types.InitialClaim{{CampaignId: 1, MissionId: 3}}
	missions := []types.Mission{{CampaignId: 1, MissionId: 3, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.2")}}
	genesisState := types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions}
	testHelper.C4eAirdropUtils.InitGenesis(genesisState)

	records := map[string]sdk.Int{acountsAddresses[0].String(): sdk.NewInt(10000)}
	testHelper.C4eAirdropUtils.AddCampaignRecords(acountsAddresses[1], 1, records)

	testHelper.C4eAirdropUtils.ClaimInitialError(2, acountsAddresses[0], "campaign not found: campaign id: 2 : not found")

}

func TestClaimInitialCampaignClaimError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	end := testHelper.Context.BlockTime().Add(1000)
	params := types.Params{Denom: commontestutils.DefaultTestDenom, Campaigns: []*types.Campaign{
		{
			CampaignId:    1,
			Enabled:       true,
			StartTime:     testHelper.Context.BlockTime(),
			EndTime:       &end,
			LockupPeriod:  1000,
			VestingPeriod: 2000,
			Description:   "test-campaign",
		},
	}}
	initialClaims := []types.InitialClaim{{CampaignId: 1, MissionId: 3}}
	missions := []types.Mission{{CampaignId: 1, MissionId: 3, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.2")}}
	genesisState := types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions}
	testHelper.C4eAirdropUtils.InitGenesis(genesisState)

	records := map[string]sdk.Int{acountsAddresses[0].String(): sdk.NewInt(10000)}
	testHelper.C4eAirdropUtils.AddCampaignRecords(acountsAddresses[1], 1, records)
	claimRecord := testHelper.C4eAirdropUtils.GetClaimRecord(acountsAddresses[0].String())
	claimRecord.GetCampaignRecords()[0].ClaimedMissions = []uint64{3}
	testHelper.C4eAirdropUtils.SetClaimRecord(claimRecord)

	testHelper.C4eAirdropUtils.ClaimInitialError(1, acountsAddresses[0], fmt.Sprintf("mission already claimed: address %s, campaignId: 1, missionId: 3: mission already claimed", acountsAddresses[0]))

}

func TestClaimInitialTwoCampaigns(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	end := testHelper.Context.BlockTime().Add(1000)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	params := types.Params{Denom: commontestutils.DefaultTestDenom, Campaigns: []*types.Campaign{
		{
			CampaignId:    1,
			Enabled:       true,
			StartTime:     testHelper.Context.BlockTime(),
			EndTime:       &end,
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
			Description:   "test-campaign",
		},
		{
			CampaignId:    2,
			Enabled:       true,
			StartTime:     testHelper.Context.BlockTime(),
			EndTime:       &end,
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
			Description:   "test-campaign-1",
		},
	}}
	initialClaims := []types.InitialClaim{{CampaignId: 1, MissionId: 3}, {CampaignId: 2, MissionId: 4}}

	missions := []types.Mission{
		{CampaignId: 1, MissionId: 3, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.2")},
		{CampaignId: 2, MissionId: 4, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.3")},
	}
	genesisState := types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions}
	testHelper.C4eAirdropUtils.InitGenesis(genesisState)

	records := map[string]sdk.Int{acountsAddresses[0].String(): sdk.NewInt(10000)}
	testHelper.C4eAirdropUtils.AddCampaignRecords(acountsAddresses[1], 1, records)
	testHelper.C4eAirdropUtils.AddCampaignRecords(acountsAddresses[1], 2, records)

	testHelper.C4eAirdropUtils.ClaimInitial(1, acountsAddresses[0])
	testHelper.C4eAirdropUtils.ClaimInitial(2, acountsAddresses[0])
}

// TODO test with 2 initial claims for different camapaigns for same address
