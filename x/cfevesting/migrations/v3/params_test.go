package v3_test

import (
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	v2 "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v2"

	v3 "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMigrate(t *testing.T) {
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	testUtil.ParamsStore.WithKeyTable(v2.ParamKeyTable())
	testUtil.ParamsStore.Set(ctx, v2.KeyDenom, v2.DefaultDenom)
	require.NoError(t, v3.MigrateParams(ctx, testUtil.StoreKey, testUtil.ParamsStore, testUtil.Cdc))
	store := ctx.KVStore(testUtil.StoreKey)

	var res v2.Params
	bz := store.Get(v3.ParamsKey)
	require.NoError(t, testUtil.Cdc.Unmarshal(bz, &res))

	var params v2.Params
	testUtil.ParamsStore.GetParamSet(ctx, &params)
	require.Equal(t, params, res)
}
