package keeper_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNEnergyTransfer(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EnergyTransfer {
	items := make([]types.EnergyTransfer, n)
	for i := range items {
		items[i].Id = keeper.AppendEnergyTransfer(ctx, items[i])
	}
	return items
}

func TestEnergyTransferGet(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createNEnergyTransfer(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEnergyTransfer(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEnergyTransferRemove(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createNEnergyTransfer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEnergyTransfer(ctx, item.Id)
		_, found := keeper.GetEnergyTransfer(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEnergyTransferGetAll(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createNEnergyTransfer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEnergyTransfer(ctx)),
	)
}

func TestEnergyTransferCount(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createNEnergyTransfer(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEnergyTransferCount(ctx))
}
