package keeper_test

import (
	"fmt"
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNVestingAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VestingAccountTrace {
	items := make([]types.VestingAccountTrace, n)
	for i := range items {
		items[i].Address = fmt.Sprintf("Address%d", i)
	}
	for i := range items {
		items[i].Id = keeper.AppendVestingAccountTrace(ctx, items[i])
	}
	return items
}

func TestVestingAccountGet(t *testing.T) {
	keeper, ctx := keepertest.CfevestingKeeper(t)
	items := createNVestingAccount(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetVestingAccountTraceById(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
		got, found = keeper.GetVestingAccountTrace(ctx, item.Address)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestVestingAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.CfevestingKeeper(t)
	items := createNVestingAccount(keeper, ctx, 10)
	for _, item := range items {
		_, found := keeper.GetVestingAccountTrace(ctx, item.Address)
		require.True(t, found)
		keeper.RemoveVestingAccountTrace(ctx, item.Address)
		_, found = keeper.GetVestingAccountTraceById(ctx, item.Id)
		require.False(t, found)
		_, found = keeper.GetVestingAccountTrace(ctx, item.Address)
		require.False(t, found)
	}
}

func TestVestingAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfevestingKeeper(t)
	items := createNVestingAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVestingAccountTrace(ctx)),
	)
}

func TestVestingAccountCount(t *testing.T) {
	keeper, ctx := keepertest.CfevestingKeeper(t)
	items := createNVestingAccount(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetVestingAccountTraceCount(ctx))
}
