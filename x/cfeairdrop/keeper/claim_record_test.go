package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
// var _ = strconv.IntSize

func createNClaimRecord(keeper *keeper.Keeper, ctx sdk.Context, numOfClaimRecords int, numOfCampaignRecords int, addClaimAddr bool, addCompletedMissions bool) []types.UserAirdropEntries {
	claimRecords := make([]types.UserAirdropEntries, numOfClaimRecords)
	for i := range claimRecords {
		claimRecords[i].Address = strconv.Itoa(i)
		if addClaimAddr {
			claimRecords[i].ClaimAddress = strconv.Itoa(1000000 + i)
		}
		campaignRecords := make([]types.CampaignRecord, numOfCampaignRecords)
		for j := range campaignRecords {
			campaignRecords[j].CampaignId = uint64(2000000 + i)
			campaignRecords[j].Claimable = sdk.NewInt(int64(3000000 + i))
			if addCompletedMissions {
				campaignRecords[j].CompletedMissions = []uint64{uint64(4000000 + i), uint64(5000000 + i), uint64(6000000 + i)}
			}

		}
		keeper.SetUserAirdropEntries(ctx, claimRecords[i])
	}
	return claimRecords
}

func TestClaimRecordGet(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNClaimRecord(keeper, ctx, 10, 0, false, false)
	for _, item := range items {
		rst, found := keeper.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNClaimRecord(keeper, ctx, 10, 10, false, false)
	for _, item := range items {
		rst, found := keeper.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNClaimRecord(keeper, ctx, 10, 10, true, false)
	for _, item := range items {
		rst, found := keeper.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNClaimRecord(keeper, ctx, 10, 10, false, true)
	for _, item := range items {
		rst, found := keeper.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestClaimRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNClaimRecord(keeper, ctx, 10, 0, false, false)
	for _, item := range items {
		keeper.RemoveClaimRecord(ctx,
			item.Address,
		)
		_, found := keeper.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestClaimRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNClaimRecord(keeper, ctx, 10, 0, false, false)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetUserAirdropEntries(ctx)),
	)

	items = createNClaimRecord(keeper, ctx, 10, 10, true, true)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetUserAirdropEntries(ctx)),
	)
}

func TestNewClaimRecordWithNewCampaignRecords(t *testing.T) {
	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	srcAddr := commontestutils.CreateIncrementalAccounts(1, 100)[0]

	campaignRecordsData := map[string]sdk.Int{}
	for i, addr := range acountsAddresses {
		campaignRecordsData[addr.String()] = sdk.NewInt(int64(100 + i))
	}
	campaignId := uint64(23)
	testUtil.AddCampaignRecords(ctx, srcAddr, campaignId, campaignRecordsData)

}

func TestAddNewCampaignRecordsToExistingClaimRecords(t *testing.T) {
	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	srcAddr := commontestutils.CreateIncrementalAccounts(1, 100)[0]
	campaignRecordsData := map[string]sdk.Int{}
	for i, addr := range acountsAddresses {
		campaignRecordsData[addr.String()] = sdk.NewInt(int64(100 + i))
	}
	campaignId := uint64(23)
	testUtil.AddCampaignRecords(ctx, srcAddr, campaignId, campaignRecordsData)

	campaignRecordsData = map[string]sdk.Int{}
	for i, addr := range acountsAddresses {
		campaignRecordsData[addr.String()] = sdk.NewInt(int64(500 + i))
	}
	campaignId = uint64(24)
	testUtil.AddCampaignRecords(ctx, srcAddr, campaignId, campaignRecordsData)
}

func TestAddExistingCampaignRecordsToExistingClaimRecords(t *testing.T) {
	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	srcAddr := commontestutils.CreateIncrementalAccounts(1, 100)[0]
	campaignRecordsData := map[string]sdk.Int{}
	for i, addr := range acountsAddresses {
		campaignRecordsData[addr.String()] = sdk.NewInt(int64(100 + i))
	}
	campaignId := uint64(23)
	testUtil.AddCampaignRecords(ctx, srcAddr, campaignId, campaignRecordsData)

	testUtil.AddCampaignRecordsError(ctx, srcAddr, campaignId, map[string]sdk.Int{acountsAddresses[5].String(): campaignRecordsData[acountsAddresses[5].String()]},
		fmt.Sprintf("campaignId 23 already exists for address: %s: entity already exists", acountsAddresses[5]), true)
}

func TestAddCampaignRecordsSendError(t *testing.T) {
	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	srcAddr := commontestutils.CreateIncrementalAccounts(1, 100)[0]
	campaignRecordsData := map[string]sdk.Int{}
	for i, addr := range acountsAddresses {
		campaignRecordsData[addr.String()] = sdk.NewInt(int64(100 + i))
	}
	campaignId := uint64(23)

	testUtil.AddCampaignRecordsError(ctx, srcAddr, campaignId, map[string]sdk.Int{acountsAddresses[5].String(): campaignRecordsData[acountsAddresses[5].String()]},
		"0uc4e is smaller than 105uc4e: insufficient funds", false)
}
