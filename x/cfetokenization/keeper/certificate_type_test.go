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

func createNCertificateType(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CertificateType {
	items := make([]types.CertificateType, n)
	for i := range items {
		items[i].Id = keeper.AppendCertificateType(ctx, items[i])
	}
	return items
}

func TestCertificateTypeGet(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNCertificateType(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCertificateType(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCertificateTypeRemove(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNCertificateType(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCertificateType(ctx, item.Id)
		_, found := keeper.GetCertificateType(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCertificateTypeGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNCertificateType(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCertificateType(ctx)),
	)
}

func TestCertificateTypeCount(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNCertificateType(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCertificateTypeCount(ctx))
}
