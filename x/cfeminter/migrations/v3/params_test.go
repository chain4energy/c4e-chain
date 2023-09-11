package v3_test

import (
	"cosmossdk.io/math"
	testenv "github.com/chain4energy/c4e-chain/v2/testutil/env"
	testkeeper "github.com/chain4energy/c4e-chain/v2/testutil/keeper"
	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/migrations/v2"
	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/migrations/v3"
	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/types"
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
	Minters := []*v2.Minter{
		createV2Minter(1, nil, v2.NoMintingType, nil, nil),
	}

	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), Minters)
	MigrateParamsV2ToV3(t, ctx, testUtil, false, "")
}

func TestMigrateLinearMinting(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	endTime := testenv.TestEnvTime.Add(time.Hour)
	Minters := []*v2.Minter{
		createV2Minter(1, &endTime, v2.LinearMintingType, nil, createLinearMinting(math.NewInt(1000))),
		createV2Minter(2, nil, v2.NoMintingType, nil, nil),
	}

	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), Minters)
	MigrateParamsV2ToV3(t, ctx, testUtil, false, "")
}

func TestMigrateExponentialStepMinting(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	endTime := testenv.TestEnvTime.Add(time.Hour)
	Minters := []*v2.Minter{
		createV2Minter(1, &endTime, v2.ExponentialStepMintingType,
			createExonentialStepMinting(math.NewInt(1000), time.Hour, sdk.MustNewDecFromStr("0.5")), nil),
		createV2Minter(2, nil, v2.NoMintingType, nil, nil),
	}

	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), Minters)
	MigrateParamsV2ToV3(t, ctx, testUtil, false, "")
}

func TestMigrateExponentialStepMintingAndLinearMinting(t *testing.T) {
	testUtil, ctx := testkeeper.CfeminterKeeperTestUtilWithCdc(t)
	endTime := testenv.TestEnvTime.Add(time.Hour)
	endTime2 := endTime.Add(time.Hour)

	Minters := []*v2.Minter{
		createV2Minter(1, &endTime, v2.LinearMintingType, nil, createLinearMinting(math.NewInt(1000))),
		createV2Minter(2, &endTime2, v2.ExponentialStepMintingType,
			createExonentialStepMinting(math.NewInt(1000), time.Hour, sdk.MustNewDecFromStr("0.5")), nil),
		createV2Minter(3, nil, v2.NoMintingType, nil, nil),
	}

	setV2MinterParams(t, ctx, testUtil, testenv.DefaultTestDenom, time.Now(), Minters)
	MigrateParamsV2ToV3(t, ctx, testUtil, false, "")
}

func MigrateParamsV2ToV3(
	t *testing.T,
	ctx sdk.Context,
	testUtil *testkeeper.ExtendedC4eMinterKeeperUtils,
	expectError bool, errorMessage string,
) {
	var oldMinterConfig v2.MinterConfig
	store := newStore(ctx, testUtil)
	oldMinterConfigRaw := store.Get(v2.KeyMinterConfig)
	err := codec.NewLegacyAmino().UnmarshalJSON(oldMinterConfigRaw, &oldMinterConfig)
	require.NoError(t, err)

	var oldMintDenom string
	oldMintDenomRaw := store.Get(v2.KeyMintDenom)
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
		if oldMinter.Type == v2.ExponentialStepMintingType {
			config, _ = codectypes.NewAnyWithValue(oldMinter.ExponentialStepMinting)
		} else if oldMinter.Type == v2.LinearMintingType {
			config, _ = codectypes.NewAnyWithValue(oldMinter.LinearMinting)
		} else {
			config, _ = codectypes.NewAnyWithValue(&types.NoMinting{})
		}
		require.EqualValues(t, newMinter.Config, config)
	}
}

func setV2MinterParams(t *testing.T, ctx sdk.Context, testUtil *testkeeper.ExtendedC4eMinterKeeperUtils, mintDenom string, startTime time.Time, minters []*v2.Minter) {
	setV2MintDenom(t, ctx, testUtil, mintDenom)
	minterConfig := v2.MinterConfig{
		StartTime: startTime,
		Minters:   minters,
	}

	store := newStore(ctx, testUtil)
	bz, err := codec.NewLegacyAmino().MarshalJSON(minterConfig)
	require.NoError(t, err)
	store.Set(v2.KeyMinterConfig, bz)
}

func setV2MintDenom(t *testing.T, ctx sdk.Context, testUtil *testkeeper.ExtendedC4eMinterKeeperUtils, mintDenom string) {
	store := newStore(ctx, testUtil)
	bz, err := codec.NewLegacyAmino().MarshalJSON(mintDenom)
	require.NoError(t, err)
	store.Set(v2.KeyMintDenom, bz)
}

func newStore(ctx sdk.Context, testUtil *testkeeper.ExtendedC4eMinterKeeperUtils) prefix.Store {
	return prefix.NewStore(ctx.KVStore(testUtil.StoreKey), append([]byte((testUtil.Subspace.Name())), '/'))
}

func createLinearMinting(
	amount math.Int,
) *v2.LinearMinting {
	return &v2.LinearMinting{
		Amount: amount,
	}
}

func createExonentialStepMinting(
	amount math.Int,
	stepDuration time.Duration,
	amountMultiplier sdk.Dec,
) *v2.ExponentialStepMinting {
	return &v2.ExponentialStepMinting{
		StepDuration:     stepDuration,
		Amount:           amount,
		AmountMultiplier: amountMultiplier,
	}
}

func createV2Minter(
	sequenceId uint32,
	endTime *time.Time,
	minterType string,
	exponentialStepMinting *v2.ExponentialStepMinting,
	linearMinting *v2.LinearMinting,
) *v2.Minter {
	return &v2.Minter{
		SequenceId:             sequenceId,
		EndTime:                endTime,
		Type:                   minterType,
		LinearMinting:          linearMinting,
		ExponentialStepMinting: exponentialStepMinting,
	}
}
