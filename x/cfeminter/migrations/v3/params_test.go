package v3_test

import (
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v3"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMigrateLinearMinter(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	minterConfigNoMinting := types.MinterConfig{
		StartTime: time.Now(),
		Minters: []*types.LegacyMinter{
			{
				SequenceId: 1,
				Type:       types.NoMintingType,
			},
		},
	}
	setV2MinterConfig(t, ctx, testUtil, minterConfigNoMinting)
	setV2MintDenom(t, ctx, testUtil, "uc4e")
	MigrateParamsV2ToV3(t, ctx, testUtil, false, "")
}

func setV2MinterConfig(t *testing.T, ctx sdk.Context, testUtil *testkeeper.ExtendedC4eMinterKeeperUtils, minterConfig types.MinterConfig) {
	store := newStore(ctx, testUtil)
	bz, err := codec.NewLegacyAmino().MarshalJSON(minterConfig)
	require.NoError(t, err)
	store.Set(types.KeyMinterConfig, bz)
}

func setV2MintDenom(t *testing.T, ctx sdk.Context, testUtil *testkeeper.ExtendedC4eMinterKeeperUtils, mintDenom string) {
	store := newStore(ctx, testUtil)
	bz, err := codec.NewLegacyAmino().MarshalJSON(mintDenom)
	require.NoError(t, err)
	store.Set(types.KeyMintDenom, bz)
}

func newStore(ctx sdk.Context, testUtil *testkeeper.ExtendedC4eMinterKeeperUtils) prefix.Store {
	return prefix.NewStore(ctx.KVStore(testUtil.StoreKey), append([]byte((testUtil.Subspace.Name())), '/'))
}

func MigrateParamsV2ToV3(
	t *testing.T,
	ctx sdk.Context,
	testUtil *testkeeper.ExtendedC4eMinterKeeperUtils,
	expectError bool, errorMessage string,
) {
	var oldMinterConfig types.MinterConfig
	store := newStore(ctx, testUtil)
	oldMinterConfigRaw := store.Get(types.KeyMinterConfig)
	err := codec.NewLegacyAmino().UnmarshalJSON(oldMinterConfigRaw, &oldMinterConfig)

	var oldMintDenom string
	oldMintDenomRaw := store.Get(types.KeyMintDenom)
	err = codec.NewLegacyAmino().UnmarshalJSON(oldMintDenomRaw, &oldMintDenom)
	require.NoError(t, err)

	err = v3.MigrateParams(ctx, testUtil.StoreKey, testUtil.Subspace, testUtil.Cdc)
	if expectError {
		require.EqualError(t, err, errorMessage)
		return
	}
	require.NoError(t, err)
	var newParams types.Params
	bz := store.Get(types.ParamsKey)
	err = codec.NewLegacyAmino().UnmarshalJSON(bz, &newParams)
	if err != nil {
		return
	}
	require.EqualValues(t, len(newParams.Minters), len(oldMinterConfig.Minters))
}
