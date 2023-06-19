package v3_test

import (
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	v2 "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v2"
	v3 "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v3"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMigrate(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	testUtil.Subspace.Set(ctx, v2.KeySubDistributors, v2.DefaultSubDistributors)
	require.NoError(t, v3.MigrateParams(ctx, testUtil.StoreKey, testUtil.Subspace, testUtil.Cdc))
	store := ctx.KVStore(testUtil.StoreKey)
	var res v2.Params
	bz := store.Get(v3.ParamsKey)
	require.NoError(t, testUtil.Cdc.Unmarshal(bz, &res))
	var params v2.Params
	testUtil.Subspace.GetParamSet(ctx, &params)
	require.Equal(t, params, res)
}
