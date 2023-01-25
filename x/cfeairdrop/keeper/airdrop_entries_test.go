package keeper_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

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
	userAirdropEntries := make([]types.UserAirdropEntries, numOfClaimRecords)
	for i := range userAirdropEntries {
		userAirdropEntries[i].Address = strconv.Itoa(i)
		if addClaimAddr {
			userAirdropEntries[i].ClaimAddress = strconv.Itoa(1000000 + i)
		}
		airdropEntryStates := make([]types.AirdropEntry, numOfCampaignRecords)
		for j := range airdropEntryStates {
			airdropEntryStates[j].CampaignId = uint64(2000000 + i)
			airdropEntryStates[j].Amount = sdk.NewInt(int64(3000000 + i))
			if addCompletedMissions {
				airdropEntryStates[j].CompletedMissions = []uint64{uint64(4000000 + i), uint64(5000000 + i), uint64(6000000 + i)}
			}

		}
		keeper.SetUserAirdropEntries(ctx, userAirdropEntries[i])
	}
	return userAirdropEntries
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

func TestClaimRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNClaimRecord(keeper, ctx, 10, 0, false, false)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetUsersAirdropEntries(ctx)),
	)

	items = createNClaimRecord(keeper, ctx, 10, 10, true, true)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetUsersAirdropEntries(ctx)),
	)
}

func TestNewClaimRecordWithNewCampaignRecords(t *testing.T) {
	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	srcAddr := commontestutils.CreateIncrementalAccounts(1, 100)[0]

	airdropEntries := generateAirdropEntries(acountsAddresses, 100000000)

	start := ctx.BlockTime()
	end := ctx.BlockTime().Add(1000)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	campaign := types.Campaign{
		Owner:         srcAddr.String(),
		Denom:         commontestutils.DefaultTestDenom,
		Enabled:       true,
		Name:          "NewCampaign",
		StartTime:     &start,
		EndTime:       &end,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
		Description:   "test-campaign",
	}
	testUtil.CreateAirdropCampaign(ctx, campaign.Owner, campaign.Name, campaign.Description, campaign.Denom, *campaign.StartTime, *campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)
	testUtil.AddAirdropEntries(ctx, srcAddr, 0, airdropEntries)

}

func TestAddNewCampaignRecordsToExistingClaimRecords(t *testing.T) {
	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	srcAddr := commontestutils.CreateIncrementalAccounts(1, 100)[0]
	airdropEntries := generateAirdropEntries(acountsAddresses, 100000000)

	start := ctx.BlockTime()
	end := ctx.BlockTime().Add(1000)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	campaign := types.Campaign{
		Owner:         srcAddr.String(),
		Denom:         commontestutils.DefaultTestDenom,
		Enabled:       true,
		Name:          "NewCampaign",
		StartTime:     &start,
		EndTime:       &end,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
		Description:   "test-campaign",
	}
	testUtil.CreateAirdropCampaign(ctx, campaign.Owner, campaign.Name, campaign.Description, campaign.Denom, *campaign.StartTime, *campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)

	testUtil.AddAirdropEntries(ctx, srcAddr, 0, airdropEntries)
	testUtil.CreateAirdropCampaign(ctx, campaign.Owner, campaign.Name, campaign.Description, campaign.Denom, *campaign.StartTime, *campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)

	airdropEntries = generateAirdropEntries(acountsAddresses, 500000000)
	testUtil.AddAirdropEntries(ctx, srcAddr, 1, airdropEntries)
}

func TestAddExistingCampaignRecordsToExistingClaimRecords(t *testing.T) {
	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	srcAddr := commontestutils.CreateIncrementalAccounts(1, 100)[0]
	airdropEntries := generateAirdropEntries(acountsAddresses, 100000000)
	start := ctx.BlockTime()
	end := ctx.BlockTime().Add(1000)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	campaign := types.Campaign{
		Owner:         srcAddr.String(),
		Denom:         commontestutils.DefaultTestDenom,
		Enabled:       true,
		Name:          "NewCampaign",
		StartTime:     &start,
		EndTime:       &end,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
		Description:   "test-campaign",
	}
	testUtil.CreateAirdropCampaign(ctx, campaign.Owner, campaign.Name, campaign.Description, campaign.Denom, *campaign.StartTime, *campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)

	testUtil.AddAirdropEntries(ctx, srcAddr, 0, airdropEntries)

	testUtil.AddCampaignRecordsError(ctx, srcAddr, 0, []*types.AirdropEntry{
		{
			Address: airdropEntries[5].Address,
			Amount:  airdropEntries[5].Amount,
		},
	},
		fmt.Sprintf("campaignId 0 already exists for address: %s: entity already exists", acountsAddresses[5]), true)
}

func TestAddCampaignRecordsSendError(t *testing.T) {
	testUtil, _, ctx := keepertest.CfeairdropKeeperTestUtilWithCdc(t)
	acountsAddresses, _ := commontestutils.CreateAccounts(10, 0)
	srcAddr := commontestutils.CreateIncrementalAccounts(1, 100)[0]
	airdropEntries := generateAirdropEntries(acountsAddresses, 100000000)
	start := ctx.BlockTime()
	end := ctx.BlockTime().Add(1000)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	campaign := types.Campaign{
		Owner:         srcAddr.String(),
		Denom:         commontestutils.DefaultTestDenom,
		Enabled:       true,
		Name:          "NewCampaign",
		StartTime:     &start,
		EndTime:       &end,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
		Description:   "test-campaign",
	}
	testUtil.CreateAirdropCampaign(ctx, campaign.Owner, campaign.Name, campaign.Description, campaign.Denom, *campaign.StartTime, *campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)

	testUtil.AddCampaignRecordsError(ctx, srcAddr, 0, []*types.AirdropEntry{
		{
			Address: airdropEntries[5].Address,
			Amount:  airdropEntries[5].Amount,
		},
	},
		"0uc4e is smaller than 100000005uc4e: insufficient funds", false)
}

func generateAirdropEntries(addresses []sdk.AccAddress, startAmount int) []*types.AirdropEntry {
	var airdropEntries []*types.AirdropEntry
	for i, addr := range addresses {
		newAirdropEntry := types.AirdropEntry{
			Address: addr.String(),
			Amount:  sdk.NewInt(int64(startAmount + i)),
		}
		airdropEntries = append(airdropEntries, &newAirdropEntry)
	}
	return airdropEntries
}