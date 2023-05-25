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

func createNEnergyTransferOffer(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EnergyTransferOffer {
	items := make([]types.EnergyTransferOffer, n)
	for i := range items {
		items[i].Id = keeper.AppendEnergyTransferOffer(ctx, items[i])
	}
	return items
}

func TestEnergyTransferOfferGet(t *testing.T) {
	keeper, ctx := keepertest.CfeevKeeper(t)
	items := createNEnergyTransferOffer(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEnergyTransferOffer(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEnergyTransferOfferRemove(t *testing.T) {
	keeper, ctx := keepertest.CfeevKeeper(t)
	items := createNEnergyTransferOffer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEnergyTransferOffer(ctx, item.Id)
		_, found := keeper.GetEnergyTransferOffer(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEnergyTransferOfferGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfeevKeeper(t)
	items := createNEnergyTransferOffer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEnergyTransferOffer(ctx)),
	)
}

func TestEnergyTransferOfferCount(t *testing.T) {
	keeper, ctx := keepertest.CfeevKeeper(t)
	items := createNEnergyTransferOffer(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEnergyTransferOfferCount(ctx))
}
