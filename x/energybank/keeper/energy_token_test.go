package keeper_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/energybank/keeper"
	"github.com/chain4energy/c4e-chain/x/energybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNEnergyToken(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EnergyToken {
	items := make([]types.EnergyToken, n)
	for i := range items {
		items[i].Id = keeper.AppendEnergyToken(ctx, items[i])
	}
	return items
}

func TestEnergyTokenGet(t *testing.T) {
	keeper, ctx := keepertest.EnergybankKeeper(t)
	items := createNEnergyToken(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEnergyToken(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEnergyTokenRemove(t *testing.T) {
	keeper, ctx := keepertest.EnergybankKeeper(t)
	items := createNEnergyToken(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEnergyToken(ctx, item.Id)
		_, found := keeper.GetEnergyToken(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEnergyTokenGetAll(t *testing.T) {
	keeper, ctx := keepertest.EnergybankKeeper(t)
	items := createNEnergyToken(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEnergyToken(ctx)),
	)
}

func TestEnergyTokenCount(t *testing.T) {
	keeper, ctx := keepertest.EnergybankKeeper(t)
	items := createNEnergyToken(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEnergyTokenCount(ctx))
}
