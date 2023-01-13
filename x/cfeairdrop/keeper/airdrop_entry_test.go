package keeper_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNAirdropEntry(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AirdropEntry {
	items := make([]types.AirdropEntry, n)
	for i := range items {
		items[i].Id = keeper.AppendAirdropEntry(ctx, items[i])
	}
	return items
}

func TestAirdropEntryGet(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNAirdropEntry(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAirdropEntry(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAirdropEntryRemove(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNAirdropEntry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAirdropEntry(ctx, item.Id)
		_, found := keeper.GetAirdropEntry(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAirdropEntryGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNAirdropEntry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAirdropEntry(ctx)),
	)
}

func TestAirdropEntryCount(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNAirdropEntry(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAirdropEntryCount(ctx))
}
