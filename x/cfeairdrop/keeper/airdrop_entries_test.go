package keeper_test

import (
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/module/cfeairdrop"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUserAirdropEntries(keeper *keeper.Keeper, ctx sdk.Context, numberOfUserAirdropEntries int, numberOfAirdropEntreis int, addClaimAddress bool, addCompletedMissions bool) []types.UserAirdropEntries {
	userAirdropEntries := make([]types.UserAirdropEntries, numberOfUserAirdropEntries)
	for i := range userAirdropEntries {
		userAirdropEntries[i].Address = strconv.Itoa(i)
		if addClaimAddress {
			userAirdropEntries[i].ClaimAddress = strconv.Itoa(1000000 + i)
		}
		airdropEntryStates := make([]types.AirdropEntry, numberOfAirdropEntreis)
		for j := range airdropEntryStates {
			airdropEntryStates[j].CampaignId = uint64(2000000 + i)
			airdropEntryStates[j].AirdropCoins = sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(int64(3000000+i))))
			if addCompletedMissions {
				airdropEntryStates[j].CompletedMissions = []uint64{uint64(4000000 + i), uint64(5000000 + i), uint64(6000000 + i)}
			}

		}
		keeper.SetUserAirdropEntries(ctx, userAirdropEntries[i])
	}
	return userAirdropEntries
}

func TestUserAirdropEntriesGet(t *testing.T) {
	k, ctx := keepertest.CfeairdropKeeper(t)
	items := createNUserAirdropEntries(k, ctx, 10, 0, false, false)
	for _, item := range items {
		rst, found := k.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNUserAirdropEntries(k, ctx, 10, 10, false, false)
	for _, item := range items {
		rst, found := k.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNUserAirdropEntries(k, ctx, 10, 10, true, false)
	for _, item := range items {
		rst, found := k.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNUserAirdropEntries(k, ctx, 10, 10, false, true)
	for _, item := range items {
		rst, found := k.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestUserAirdropEntriesGetAll(t *testing.T) {
	k, ctx := keepertest.CfeairdropKeeper(t)
	items := createNUserAirdropEntries(k, ctx, 10, 0, false, false)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetUsersAirdropEntries(ctx)),
	)

	items = createNUserAirdropEntries(k, ctx, 10, 10, true, true)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetUsersAirdropEntries(ctx)),
	)
}

func TestAddUserAirdropEntries(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries := createTestAirdropEntries(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, airdropEntries)
}

func TestAddUserAirdropEntriesEmptyAirdropCoins(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries := createTestAirdropEntries(acountsAddresses, 100000000)
	airdropEntries[0].AirdropCoins = sdk.NewCoins()
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, airdropEntries)
}

//func TestAddNewCampaignRecordsToExistingUserAirdropEntriess(t *testing.T) {
//	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
//	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
//	srcAddr := testcosmos.CreateIncrementalAccounts(1, 100)[0]
//	airdropEntries := generateAirdropEntries(acountsAddresses, 100000000)
//
//	start := ctx.BlockTime()
//	end := ctx.BlockTime().Add(1000)
//	lockupPeriod := time.Hour
//	vestingPeriod := 3 * time.Hour
//	campaign := types.Campaign{
//		Owner:         srcAddr.String(),
//		Enabled:       true,
//		Name:          "NewCampaign",
//		StartTime:     start,
//		EndTime:       end,
//		LockupPeriod:  lockupPeriod,
//		VestingPeriod: vestingPeriod,
//		Description:   "test-campaign",
//	}
//	testUtil.CreateAirdropCampaign(ctx, campaign.Owner, campaign.Name, campaign.Description, campaign.FeegrantAmount, campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)
//
//	testUtil.AddAirdropEntries(ctx, srcAddr, 0, airdropEntries)
//	testUtil.CreateAirdropCampaign(ctx, campaign.Owner, campaign.Name, campaign.Description, campaign.FeegrantAmount, campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)
//
//	airdropEntries = generateAirdropEntries(acountsAddresses, 500000000)
//	testUtil.AddAirdropEntries(ctx, srcAddr, 1, airdropEntries)
//}
//
//func TestAddExistingCampaignRecordsToExistingUserAirdropEntriess(t *testing.T) {
//	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
//	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
//	srcAddr := testcosmos.CreateIncrementalAccounts(1, 100)[0]
//	airdropEntries := generateAirdropEntries(acountsAddresses, 100000000)
//	start := ctx.BlockTime()
//	end := ctx.BlockTime().Add(1000)
//	lockupPeriod := time.Hour
//	vestingPeriod := 3 * time.Hour
//	campaign := types.Campaign{
//		Owner:         srcAddr.String(),
//		Enabled:       true,
//		Name:          "NewCampaign",
//		StartTime:     start,
//		EndTime:       end,
//		LockupPeriod:  lockupPeriod,
//		VestingPeriod: vestingPeriod,
//		Description:   "test-campaign",
//	}
//	testUtil.CreateAirdropCampaign(ctx, campaign.Owner, campaign.Name, campaign.Description, campaign.FeegrantAmount, campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)
//
//	testUtil.AddAirdropEntries(ctx, srcAddr, 0, airdropEntries)
//
//	testUtil.AddCampaignRecordsError(ctx, srcAddr, 0, []*types.AirdropEntry{
//		{
//			Address:      airdropEntries[5].Address,
//			AirdropCoins: airdropEntries[5].AirdropCoins,
//		},
//	},
//		fmt.Sprintf("campaignId 0 already exists for address: %s: entity already exists", acountsAddresses[5]), true)
//}
//
//func TestAddCampaignRecordsSendError(t *testing.T) {
//	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
//	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
//	srcAddr := testcosmos.CreateIncrementalAccounts(1, 100)[0]
//	airdropEntries := generateAirdropEntries(acountsAddresses, 100000000)
//	start := ctx.BlockTime()
//	end := ctx.BlockTime().Add(1000)
//	lockupPeriod := time.Hour
//	vestingPeriod := 3 * time.Hour
//	campaign := types.Campaign{
//		Owner:         srcAddr.String(),
//		Enabled:       true,
//		Name:          "NewCampaign",
//		StartTime:     start,
//		EndTime:       end,
//		LockupPeriod:  lockupPeriod,
//		VestingPeriod: vestingPeriod,
//		Description:   "test-campaign",
//	}
//	testUtil.CreateAirdropCampaign(ctx, campaign.Owner, campaign.Name, campaign.Description, campaign.FeegrantAmount, campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)
//
//	testUtil.AddCampaignRecordsError(ctx, srcAddr, 0, []*types.AirdropEntry{
//		{
//			Address:      airdropEntries[5].Address,
//			AirdropCoins: airdropEntries[5].AirdropCoins,
//		},
//	},
//		"0uc4e is smaller than 100000005uc4e: insufficient funds", false)
//}

func createTestAirdropEntries(addresses []sdk.AccAddress, startAmount int) (airdropEntries []*types.AirdropEntry) {
	for i, addr := range addresses {
		newAirdropEntry := types.AirdropEntry{
			Address:      addr.String(),
			AirdropCoins: sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(int64(startAmount+i)))),
		}
		airdropEntries = append(airdropEntries, &newAirdropEntry)
	}
	return
}

func addCampaignsAndMissions(utils *cfeairdrop.ContextC4eAirdropUtils, ownerAddress string, campaigns []types.Campaign, missions []types.Mission) {
	for _, campaign := range campaigns {
		utils.CreateAirdropCampaign(ownerAddress, campaign)
		if campaign.Enabled == true {
			utils.StartAirdropCampaign(ownerAddress, campaign.Id)
		}
	}
	for _, mission := range missions {
		utils.AddMissionToAirdropCampaign(ownerAddress, mission.CampaignId, mission)
	}
}

func createCampaignMissionAndStart(testHelper *testapp.TestHelper, ownerAddress string) {
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(ownerAddress, campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(ownerAddress, 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(ownerAddress, 0)
}
