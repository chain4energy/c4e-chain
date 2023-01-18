package keeper_test

import (
	"strconv"
)

// Prevent strconv unused error
var _ = strconv.IntSize

//
//func createNInitialClaim(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.InitialClaim {
//	items := make([]types.InitialClaim, n)
//	for i := range items {
//		items[i].CampaignId = uint64(i)
//		items[i].MissionId = uint64(1000 + i)
//		keeper.SetInitialClaim(ctx, items[i])
//	}
//	return items
//}
//
//func TestInitialClaimGet(t *testing.T) {
//	keeper, ctx := keepertest.CfeairdropKeeper(t)
//	items := createNInitialClaim(keeper, ctx, 10)
//	for _, item := range items {
//		rst, found := keeper.GetInitialClaim(ctx,
//			item.CampaignId,
//		)
//		require.True(t, found)
//		require.Equal(t,
//			nullify.Fill(&item),
//			nullify.Fill(&rst),
//		)
//	}
//}
//func TestInitialClaimRemove(t *testing.T) {
//	keeper, ctx := keepertest.CfeairdropKeeper(t)
//	items := createNInitialClaim(keeper, ctx, 10)
//	for _, item := range items {
//		keeper.RemoveInitialClaim(ctx,
//			item.CampaignId,
//		)
//		_, found := keeper.GetInitialClaim(ctx,
//			item.CampaignId,
//		)
//		require.False(t, found)
//	}
//}
//
//func TestInitialClaimGetAll(t *testing.T) {
//	keeper, ctx := keepertest.CfeairdropKeeper(t)
//	items := createNInitialClaim(keeper, ctx, 10)
//	require.ElementsMatch(t,
//		nullify.Fill(items),
//		nullify.Fill(keeper.GetAllInitialClaim(ctx)),
//	)
//}
//
//func TestClaimInitial(t *testing.T) {
//	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
//
//	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
//
//	// ctx := testHelper.Context
//	start := testHelper.Context.BlockTime()
//
//	end := testHelper.Context.BlockTime().Add(1000)
//	lockupPeriod := time.Hour
//	vestingPeriod := 3 * time.Hour
//
//	campaigns := []*types.Campaign{
//		{
//			Id:            1,
//			Enabled:       true,
//			StartTime:     &start,
//			EndTime:       &end,
//			LockupPeriod:  lockupPeriod,
//			VestingPeriod: vestingPeriod,
//			Description:   "test-campaign",
//		},
//	}
//
//	params := types.Params{Denom: commontestutils.DefaultTestDenom}
//	initialClaims := []types.InitialClaim{{CampaignId: 1, MissionId: 3}}
//	missions := []types.Mission{{CampaignId: 1, Id: 3, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.2")}}
//	genesisState := types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns}
//	testHelper.C4eAirdropUtils.InitGenesis(genesisState)
//
//	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[1], 1, prepareAidropEntries(acountsAddresses[0].String()))
//
//	testHelper.C4eAirdropUtils.ClaimInitial(1, acountsAddresses[0])
//
//}
//
//func TestClaimInitialCampaignNotFound(t *testing.T) {
//	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
//
//	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
//
//	start := testHelper.Context.BlockTime()
//	end := testHelper.Context.BlockTime().Add(1000)
//	campaigns := []*types.Campaign{
//		{
//			Id:            1,
//			Enabled:       true,
//			StartTime:     &start,
//			EndTime:       &end,
//			LockupPeriod:  1000,
//			VestingPeriod: 2000,
//			Description:   "test-campaign",
//		},
//	}
//	params := types.Params{Denom: commontestutils.DefaultTestDenom}
//	initialClaims := []types.InitialClaim{{CampaignId: 1, MissionId: 3}}
//	missions := []types.Mission{{CampaignId: 1, Id: 3, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.2")}}
//	genesisState := types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns}
//	testHelper.C4eAirdropUtils.InitGenesis(genesisState)
//
//	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[1], 1, prepareAidropEntries(acountsAddresses[0].String()))
//
//	testHelper.C4eAirdropUtils.ClaimInitialError(2, acountsAddresses[0], "campaign not found: campaign id: 2 : not found")
//
//}
//
//func TestClaimInitialCampaignClaimError(t *testing.T) {
//	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
//
//	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
//	start := testHelper.Context.BlockTime()
//
//	end := testHelper.Context.BlockTime().Add(1000)
//	campaigns := []*types.Campaign{
//		{
//			Id:            1,
//			Enabled:       true,
//			StartTime:     &start,
//			EndTime:       &end,
//			LockupPeriod:  1000,
//			VestingPeriod: 2000,
//			Description:   "test-campaign",
//		},
//	}
//
//	params := types.Params{Denom: commontestutils.DefaultTestDenom}
//	initialClaims := []types.InitialClaim{{CampaignId: 1, MissionId: 3}}
//	missions := []types.Mission{{CampaignId: 1, Id: 3, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.2")}}
//	genesisState := types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns}
//	testHelper.C4eAirdropUtils.InitGenesis(genesisState)
//
//	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[1], 1, prepareAidropEntries(acountsAddresses[0].String()))
//	userAirdropEntries := testHelper.C4eAirdropUtils.GetUserAirdropEntries(acountsAddresses[0].String())
//	userAirdropEntries.GetAirdropEntries()[0].ClaimedMissions = []uint64{3}
//	testHelper.C4eAirdropUtils.SetUserAirdropEntries(userAirdropEntries)
//
//	testHelper.C4eAirdropUtils.ClaimInitialError(1, acountsAddresses[0], fmt.Sprintf("mission already claimed: address %s, campaignId: 1, missionId: 3: mission already claimed", acountsAddresses[0]))
//
//}
//
//func TestClaimInitialTwoCampaigns(t *testing.T) {
//	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
//
//	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
//
//	start := testHelper.Context.BlockTime()
//	end := testHelper.Context.BlockTime().Add(1000)
//	lockupPeriod := time.Hour
//	vestingPeriod := 3 * time.Hour
//	campaigns := []*types.Campaign{
//		{
//			Id:            1,
//			Enabled:       true,
//			StartTime:     &start,
//			EndTime:       &end,
//			LockupPeriod:  lockupPeriod,
//			VestingPeriod: vestingPeriod,
//			Description:   "test-campaign",
//		},
//		{
//			Id:            2,
//			Enabled:       true,
//			StartTime:     &start,
//			EndTime:       &end,
//			LockupPeriod:  lockupPeriod,
//			VestingPeriod: vestingPeriod,
//			Description:   "test-campaign-1",
//		},
//	}
//
//	params := types.Params{Denom: commontestutils.DefaultTestDenom}
//	initialClaims := []types.InitialClaim{{CampaignId: 1, MissionId: 3}, {CampaignId: 2, MissionId: 4}}
//
//	missions := []types.Mission{
//		{CampaignId: 1, Id: 3, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.2")},
//		{CampaignId: 2, Id: 4, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.3")},
//	}
//	genesisState := types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions, Campaigns: campaigns}
//	testHelper.C4eAirdropUtils.InitGenesis(genesisState)
//
//	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[1], 1, prepareAidropEntries(acountsAddresses[0].String()))
//	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[1], 2, prepareAidropEntries(acountsAddresses[0].String()))
//
//	testHelper.C4eAirdropUtils.ClaimInitial(1, acountsAddresses[0])
//	testHelper.C4eAirdropUtils.ClaimInitial(2, acountsAddresses[0])
//}
//
//// TODO test with 2 initial claims for different camapaigns for same address
