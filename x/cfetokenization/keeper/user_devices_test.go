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

func createNUserDevices(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UserDevices {
	items := make([]types.UserDevices, n)
	for i := range items {
		items[i].Id = keeper.AppendUserDevices(ctx, items[i])
	}
	return items
}

func TestUserDevicesGet(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNUserDevices(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetUserDevices(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestUserDevicesRemove(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNUserDevices(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserDevices(ctx, item.Id)
		_, found := keeper.GetUserDevices(ctx, item.Id)
		require.False(t, found)
	}
}

func TestUserDevicesGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNUserDevices(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserDevices(ctx)),
	)
}

func TestUserDevicesCount(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	items := createNUserDevices(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetUserDevicesCount(ctx))
}
