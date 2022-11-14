package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMission(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mission {
	items := make([]types.Mission, n)
	for i := range items {
		items[i].CampaignId = uint64(i)
		items[i].MissionId = uint64(i)

		keeper.SetMission(ctx, items[i])
	}
	return items
}

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
