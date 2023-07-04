package keeper_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfetokenization/keeper"
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNUserCertificates(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UserCertificates {
	items := make([]types.UserCertificates, n)
	for i := range items {
		items[i].Id = keeper.AppendUserCertificates(ctx, items[i])
	}
	return items
}

func TestUserCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNUserCertificates(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetUserCertificates(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestUserCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNUserCertificates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserCertificates(ctx, item.Id)
		_, found := keeper.GetUserCertificates(ctx, item.Id)
		require.False(t, found)
	}
}

func TestUserCertificatesGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNUserCertificates(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserCertificates(ctx)),
	)
}

func TestUserCertificatesCount(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNUserCertificates(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetUserCertificatesCount(ctx))
}
