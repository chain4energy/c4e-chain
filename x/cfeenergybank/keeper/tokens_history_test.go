package keeper_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeenergybank/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNTokensHistory(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TokensHistory {
	items := make([]types.TokensHistory, n)
	for i := range items {
		items[i].Id = keeper.AppendTokensHistory(ctx, items[i])
	}
	return items
}

func TestTokensHistoryGet(t *testing.T) {
	keeper, ctx := keepertest.CfeenergybankKeeper(t)
	items := createNTokensHistory(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTokensHistory(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTokensHistoryRemove(t *testing.T) {
	keeper, ctx := keepertest.CfeenergybankKeeper(t)
	items := createNTokensHistory(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTokensHistory(ctx, item.Id)
		_, found := keeper.GetTokensHistory(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTokensHistoryGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfeenergybankKeeper(t)
	items := createNTokensHistory(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTokensHistory(ctx)),
	)
}

func TestTokensHistoryCount(t *testing.T) {
	keeper, ctx := keepertest.CfeenergybankKeeper(t)
	items := createNTokensHistory(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTokensHistoryCount(ctx))
}
