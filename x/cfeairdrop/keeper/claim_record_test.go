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

func createNClaimRecord(keeper *keeper.Keeper, ctx sdk.Context, numOfClaimRecords int, numOfCampaignRecords int, addClaimAddr bool, addCompletedMissions bool) []types.ClaimRecord {
	claimRecords := make([]types.ClaimRecord, numOfClaimRecords)
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
		keeper.SetClaimRecord(ctx, claimRecords[i])
	}
	return claimRecords
}

func TestClaimRecordGet(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNClaimRecord(keeper, ctx, 10, 0, false, false)
	for _, item := range items {
		rst, found := keeper.GetClaimRecord(ctx,
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
		rst, found := keeper.GetClaimRecord(ctx,
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
		rst, found := keeper.GetClaimRecord(ctx,
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
		rst, found := keeper.GetClaimRecord(ctx,
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
		_, found := keeper.GetClaimRecord(ctx,
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
		nullify.Fill(keeper.GetAllClaimRecord(ctx)),
	)

	items = createNClaimRecord(keeper, ctx, 10, 10, true, true)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClaimRecord(ctx)),
	)
}

func TestNewClaimRecordWithNewCampaignRecords(t *testing.T) {
	k, ctx := keepertest.CfeairdropKeeper(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	campaignRecordsData := make([]*keeper.CampaignRecordData, len(acountsAddresses))
	for i, addr := range acountsAddresses {
		campaignRecordsData[i] = &keeper.CampaignRecordData{Address: addr.String(), Claimable: sdk.NewInt(int64(100 + i))}
	}
	campaignId := uint64(23)
	require.NoError(t, k.AddCampaignRecords(ctx, campaignId, campaignRecordsData))

	allRecords := k.GetAllClaimRecord(ctx)
	require.EqualValues(t, len(acountsAddresses), len(allRecords))
	for _, recordData := range campaignRecordsData {
		claimRecord, found := k.GetClaimRecord(ctx, recordData.Address)
		require.True(t, found)
		require.EqualValues(t, 1, len(claimRecord.CampaignRecords))
		require.EqualValues(t, recordData.Address, claimRecord.Address)
		require.EqualValues(t, "", claimRecord.ClaimAddress)
		require.EqualValues(t, campaignId, claimRecord.CampaignRecords[0].CampaignId)
		require.True(t, recordData.Claimable.Equal(claimRecord.CampaignRecords[0].Claimable))
		require.EqualValues(t, 0, len(claimRecord.CampaignRecords[0].CompletedMissions))

	}
}


func TestAddNewCampaignRecordsToExistingClaimRecords(t *testing.T) {
	k, ctx := keepertest.CfeairdropKeeper(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	campaignRecordsData := make([]*keeper.CampaignRecordData, len(acountsAddresses))
	for i, addr := range acountsAddresses {
		campaignRecordsData[i] = &keeper.CampaignRecordData{Address: addr.String(), Claimable: sdk.NewInt(int64(100 + i))}
	}
	campaignId := uint64(23)
	require.NoError(t, k.AddCampaignRecords(ctx, campaignId, campaignRecordsData))

	campaignRecordsData = make([]*keeper.CampaignRecordData, len(acountsAddresses))
	for i, addr := range acountsAddresses {
		campaignRecordsData[i] = &keeper.CampaignRecordData{Address: addr.String(), Claimable: sdk.NewInt(int64(500 + i))}
	}
	campaignId = uint64(24)
	require.NoError(t, k.AddCampaignRecords(ctx, campaignId, campaignRecordsData))

	allRecords := k.GetAllClaimRecord(ctx)
	require.EqualValues(t, len(acountsAddresses), len(allRecords))
	for _, recordData := range campaignRecordsData {
		claimRecord, found := k.GetClaimRecord(ctx, recordData.Address)
		require.True(t, found)
		require.EqualValues(t, 2, len(claimRecord.CampaignRecords))
		require.EqualValues(t, recordData.Address, claimRecord.Address)
		require.EqualValues(t, "", claimRecord.ClaimAddress)
		require.EqualValues(t, campaignId, claimRecord.CampaignRecords[1].CampaignId)
		require.True(t, recordData.Claimable.Equal(claimRecord.CampaignRecords[1].Claimable))
		require.EqualValues(t, 0, len(claimRecord.CampaignRecords[1].CompletedMissions))

	}
}

func TestAddExistingCampaignRecordsToExistingClaimRecords(t *testing.T) {
	k, ctx := keepertest.CfeairdropKeeper(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	campaignRecordsData := make([]*keeper.CampaignRecordData, len(acountsAddresses))
	for i, addr := range acountsAddresses {
		campaignRecordsData[i] = &keeper.CampaignRecordData{Address: addr.String(), Claimable: sdk.NewInt(int64(100 + i))}
	}
	campaignId := uint64(23)
	require.NoError(t, k.AddCampaignRecords(ctx, campaignId, campaignRecordsData))

	require.EqualError(t, k.AddCampaignRecords(ctx, campaignId, []*keeper.CampaignRecordData{campaignRecordsData[5]}), 
		fmt.Sprintf("campaignId 23 already exists for address: %s: entity already exists", campaignRecordsData[5].Address))

	allRecords := k.GetAllClaimRecord(ctx)
	require.EqualValues(t, len(acountsAddresses), len(allRecords))
	for _, recordData := range campaignRecordsData {
		claimRecord, found := k.GetClaimRecord(ctx, recordData.Address)
		require.True(t, found)
		require.EqualValues(t, 1, len(claimRecord.CampaignRecords))
		require.EqualValues(t, recordData.Address, claimRecord.Address)
		require.EqualValues(t, "", claimRecord.ClaimAddress)
		require.EqualValues(t, campaignId, claimRecord.CampaignRecords[0].CampaignId)
		require.True(t, recordData.Claimable.Equal(claimRecord.CampaignRecords[0].Claimable))
		require.EqualValues(t, 0, len(claimRecord.CampaignRecords[0].CompletedMissions))

	}
}
