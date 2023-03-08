package v3_test

import (
	"cosmossdk.io/math"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v3"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMigrateNoMinting(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	legacyMinters := []*types.LegacyMinter{
		createV2Minter(1, nil, types.NoMintingType, nil, nil),
	}

	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), legacyMinters)
	MigrateParamsV2ToV3(t, ctx, testUtil, false, "")
}

func TestMigrateLinearMinting(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	endTime := testenv.TestEnvTime.Add(time.Hour)
	legacyMinters := []*types.LegacyMinter{
		createV2Minter(1, &endTime, types.LinearMintingType, nil, createLinearMinting(sdk.NewInt(1000))),
		createV2Minter(2, nil, types.NoMintingType, nil, nil),
	}

	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), legacyMinters)
	MigrateParamsV2ToV3(t, ctx, testUtil, false, "")
}

func TestMigrateExponentialStepMinting(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	endTime := testenv.TestEnvTime.Add(time.Hour)
	legacyMinters := []*types.LegacyMinter{
		createV2Minter(1, &endTime, types.ExponentialStepMintingType,
			createExonentialStepMinting(sdk.NewInt(1000), time.Hour, sdk.MustNewDecFromStr("0.5")), nil),
		createV2Minter(2, nil, types.NoMintingType, nil, nil),
	}

	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), legacyMinters)
	MigrateParamsV2ToV3(t, ctx, testUtil, false, "")
}

func TestMigrateExponentialStepMintingAndLinearMinting(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	endTime := testenv.TestEnvTime.Add(time.Hour)
	endTime2 := endTime.Add(time.Hour)

	legacyMinters := []*types.LegacyMinter{
		createV2Minter(1, &endTime, types.LinearMintingType, nil, createLinearMinting(sdk.NewInt(1000))),
		createV2Minter(2, &endTime2, types.ExponentialStepMintingType,
			createExonentialStepMinting(sdk.NewInt(1000), time.Hour, sdk.MustNewDecFromStr("0.5")), nil),
		createV2Minter(3, nil, types.NoMintingType, nil, nil),
	}

	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), legacyMinters)
	MigrateParamsV2ToV3(t, ctx, testUtil, false, "")
}

func TestMigrateWrongSequenceId(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	endTime := testenv.TestEnvTime.Add(time.Hour)
	endTime2 := endTime.Add(time.Hour)

	legacyMinters := []*types.LegacyMinter{
		createV2Minter(1, &endTime, types.LinearMintingType, nil, createLinearMinting(sdk.NewInt(1000))),
		createV2Minter(1, &endTime2, types.ExponentialStepMintingType,
			createExonentialStepMinting(sdk.NewInt(1000), time.Hour, sdk.MustNewDecFromStr("0.5")), nil),
		createV2Minter(3, nil, types.NoMintingType, nil, nil),
	}

	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), legacyMinters)
	MigrateParamsV2ToV3(t, ctx, testUtil, true, "missing minter with sequence id 2")
}

func TestMigrateWrongMinterType(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	endTime := testenv.TestEnvTime.Add(time.Hour)
	endTime2 := endTime.Add(time.Hour)

	legacyMinters := []*types.LegacyMinter{
		createV2Minter(1, &endTime, types.LinearMintingType, nil, createLinearMinting(sdk.NewInt(1000))),
		createV2Minter(2, &endTime2, types.ExponentialStepMintingType,
			nil, nil),
		createV2Minter(3, nil, types.NoMintingType, nil, nil),
	}

	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), legacyMinters)
	MigrateParamsV2ToV3(t, ctx, testUtil, true, "minter with id 2 validation error: ExponentialStepMintingType error: for ExponentialStepMintingType type (2) ExponentialStepMinting must be set")
}

func TestMigrateNoMintersSet(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	var legacyMinters []*types.LegacyMinter
	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), legacyMinters)
	MigrateParamsV2ToV3(t, ctx, testUtil, true, "no minters defined")
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
	require.EqualValues(t, newParams.MintDenom, oldMintDenom)
	require.True(t, newParams.StartTime.Equal(oldMinterConfig.StartTime))
	for i, oldMinter := range oldMinterConfig.Minters {
		newMinter := newParams.Minters[i]
		var config *codectypes.Any
		if oldMinter.Type == types.ExponentialStepMintingType {
			config, _ = codectypes.NewAnyWithValue(oldMinter.ExponentialStepMinting)
		} else if oldMinter.Type == types.LinearMintingType {
			config, _ = codectypes.NewAnyWithValue(oldMinter.LinearMinting)
		} else {
			config, _ = codectypes.NewAnyWithValue(&types.NoMinting{})
		}
		require.EqualValues(t, newMinter.Config, config)
	}
}

func setV2MinterParams(t *testing.T, ctx sdk.Context, testUtil *testkeeper.ExtendedC4eMinterKeeperUtils, mintDenom string, startTime time.Time, minters []*types.LegacyMinter) {
	setV2MintDenom(t, ctx, testUtil, mintDenom)
	minterConfig := types.MinterConfig{
		StartTime: startTime,
		Minters:   minters,
	}

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

func createLinearMinting(
	amount math.Int,
) *types.LinearMinting {
	return &types.LinearMinting{
		Amount: amount,
	}
}

func createExonentialStepMinting(
	amount math.Int,
	stepDuration time.Duration,
	amountMultiplier sdk.Dec,
) *types.ExponentialStepMinting {
	return &types.ExponentialStepMinting{
		StepDuration:     stepDuration,
		Amount:           amount,
		AmountMultiplier: amountMultiplier,
	}
}

func createV2Minter(
	sequenceId uint32,
	endTime *time.Time,
	minterType string,
	exponentialStepMinting *types.ExponentialStepMinting,
	linearMinting *types.LinearMinting,
) *types.LegacyMinter {
	return &types.LegacyMinter{
		SequenceId:             sequenceId,
		EndTime:                endTime,
		Type:                   minterType,
		LinearMinting:          linearMinting,
		ExponentialStepMinting: exponentialStepMinting,
	}
}
