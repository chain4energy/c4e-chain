package v3_test

import (
	"github.com/chain4energy/c4e-chain/app"
	"github.com/chain4energy/c4e-chain/x/cfevesting/exported"
	v3 "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v3"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type mockSubspace struct {
	ps v3.Params
}

func newMockSubspace(ps v3.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps exported.ParamSet) {
	*ps.(*v3.Params) = ms.ps
}

func TestMigrate(t *testing.T) {
	encCfg := app.MakeEncodingConfig()
	cdc := encCfg.Codec

	storeKey := sdk.NewKVStoreKey(types.ModuleName)
	tKey := sdk.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	store := ctx.KVStore(storeKey)

	legacySubspace := newMockSubspace(v3.DefaultParams())
	require.NoError(t, v3.MigrateParams(ctx, storeKey, legacySubspace, cdc))

	var res v3.Params
	bz := store.Get(types.ParamsKey)
	require.NoError(t, cdc.Unmarshal(bz, &res))
	require.Equal(t, legacySubspace.ps, res)
}
