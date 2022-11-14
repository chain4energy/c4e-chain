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

func createNInitialClaim(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.InitialClaim {
	items := make([]types.InitialClaim, n)
	for i := range items {
		items[i].CampaignId = uint64(i)

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
