package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/energybank/keeper"
	"github.com/chain4energy/c4e-chain/x/energybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTokenParams(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TokenParams {
	items := make([]types.TokenParams, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTokenParams(ctx, items[i])
	}
	return items
}

func TestTokenParamsGet(t *testing.T) {
	keeper, ctx := keepertest.EnergybankKeeper(t)
	items := createNTokenParams(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTokenParams(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTokenParamsRemove(t *testing.T) {
	keeper, ctx := keepertest.EnergybankKeeper(t)
	items := createNTokenParams(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTokenParams(ctx,
			item.Index,
		)
		_, found := keeper.GetTokenParams(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTokenParamsGetAll(t *testing.T) {
	keeper, ctx := keepertest.EnergybankKeeper(t)
	items := createNTokenParams(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTokenParams(ctx)),
	)
}
