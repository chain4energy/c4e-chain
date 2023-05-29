package v2_test

import (
	"cosmossdk.io/math"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	v1 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v1"
	v2 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v2"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/stretchr/testify/require"

	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationLinearMinting(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	timeLinearMinter := createV1TimeLinearMinter(math.NewInt(10000))
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)
	V1MintingPeriods := []*v1.MintingPeriod{
		createV100MinterPeriod(1, &endTime, "TIME_LINEAR_MINTER", nil, timeLinearMinter),
		createV100MinterPeriod(2, nil, "NO_MINTING", nil, nil),
	}
	setV1MinterConfig(t, ctx, &keeperData, startTime, V1MintingPeriods)
	MigrateParamsV100ToV1(t, ctx, *k, &keeperData, false, "")
}

func TestMigrationExponentialStepMinting(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	periodicReductionMinter := createV1TimePeriodicReductionMinter(4, 100000, sdk.MustNewDecFromStr("0.5"), math.NewInt(10000))
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)
	V1MintingPeriods := []*v1.MintingPeriod{
		createV100MinterPeriod(1, &endTime, "PERIODIC_REDUCTION_MINTER", periodicReductionMinter, nil),
		createV100MinterPeriod(2, nil, "NO_MINTING", nil, nil),
	}
	setV1MinterConfig(t, ctx, &keeperData, startTime, V1MintingPeriods)
	MigrateParamsV100ToV1(t, ctx, *k, &keeperData, false, "")
}

func TestMigrationLinearMintingAndExponentialStepMinting(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	timeLinearMinter := createV1TimeLinearMinter(math.NewInt(10000))
	periodicReductionMinter := createV1TimePeriodicReductionMinter(4, 100000, sdk.MustNewDecFromStr("0.5"), math.NewInt(10000))
	startTime := time.Now()
	endTime1 := startTime.Add(time.Hour)
	endTime2 := endTime1.Add(time.Hour)
	V1MintingPeriods := []*v1.MintingPeriod{
		createV100MinterPeriod(1, &endTime1, "TIME_LINEAR_MINTER", nil, timeLinearMinter),
		createV100MinterPeriod(2, &endTime2, "PERIODIC_REDUCTION_MINTER", periodicReductionMinter, nil),
		createV100MinterPeriod(3, nil, "NO_MINTING", nil, nil),
	}
	setV1MinterConfig(t, ctx, &keeperData, startTime, V1MintingPeriods)
	MigrateParamsV100ToV1(t, ctx, *k, &keeperData, false, "")
}

func TestMigrationWrongMinterType(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	timeLinearMinter := createV1TimeLinearMinter(math.NewInt(10000))
	startTime := time.Now()
	endTime1 := startTime.Add(time.Hour)
	V1MintingPeriods := []*v1.MintingPeriod{
		createV100MinterPeriod(1, &endTime1, "WRONG_MINTER_TYPE", nil, timeLinearMinter),
	}
	setV1MinterConfig(t, ctx, &keeperData, startTime, V1MintingPeriods)
	MigrateParamsV100ToV1(t, ctx, *k, &keeperData, true, "wrong minting period type")
}

func setV1MinterConfig(t *testing.T, ctx sdk.Context, keeperData *testenv.AdditionalKeeperData, startTime time.Time, mintingPeriods []*v1.MintingPeriod) {
	minter := v1.Minter{
		Start:   startTime,
		Periods: mintingPeriods,
	}
	store := newStore(ctx, keeperData)
	bz, err := codec.NewLegacyAmino().MarshalJSON(minter)
	require.NoError(t, err)
	store.Set(v1.KeyMinter, bz)
}

func newStore(ctx sdk.Context, testUtil *testenv.AdditionalKeeperData) prefix.Store {
	return prefix.NewStore(ctx.KVStore(testUtil.StoreKey), append([]byte((testUtil.Subspace.Name())), '/'))
}

func getV1MinterConfig(ctx sdk.Context, keeperData *testenv.AdditionalKeeperData) (oldMinterConfig v1.Minter) {
	oldMinterConfigRaw := keeperData.Subspace.GetRaw(ctx, v1.KeyMinter)
	if err := codec.NewLegacyAmino().UnmarshalJSON(oldMinterConfigRaw, &oldMinterConfig); err != nil {
		panic(err)
	}
	return
}

func MigrateParamsV100ToV1(
	t *testing.T,
	ctx sdk.Context,
	keeper keeper.Keeper,
	keeperData *testenv.AdditionalKeeperData,
	expectError bool, errorMessage string,
) {
	oldMinterConfig := getV1MinterConfig(ctx, keeperData)
	store := newStore(ctx, keeperData)
	err := v2.MigrateParams(ctx, &keeperData.Subspace)
	if expectError {
		require.EqualError(t, err, errorMessage)
		return
	}
	require.NoError(t, err)
	var newMinterConfig v2.MinterConfig
	newMinterConfigRaw := store.Get(v2.KeyMinterConfig)
	err = codec.NewLegacyAmino().UnmarshalJSON(newMinterConfigRaw, &newMinterConfig)
	require.NoError(t, err)

	require.EqualValues(t, len(newMinterConfig.Minters), len(oldMinterConfig.Periods))
	newMinters := newMinterConfig.Minters
	for i, oldMinterPeriod := range oldMinterConfig.Periods {
		require.EqualValues(t, newMinters[i].SequenceId, oldMinterPeriod.Position)
		require.EqualValues(t, newMinters[i].EndTime, oldMinterPeriod.PeriodEnd)

		switch oldMinterPeriod.Type {
		case "TIME_LINEAR_MINTER":
			require.EqualValues(t, newMinters[i].Type, v2.LinearMintingType)
			break
		case "PERIODIC_REDUCTION_MINTER":
			require.EqualValues(t, newMinters[i].Type, v2.ExponentialStepMintingType)
			break
		case "NO_MINTING":
			require.EqualValues(t, newMinters[i].Type, v2.NoMintingType)
			break
		}

		if oldMinterPeriod.TimeLinearMinter == nil {
			require.Nil(t, newMinters[i].LinearMinting)
		} else {
			require.EqualValues(t, newMinters[i].LinearMinting.Amount, oldMinterPeriod.TimeLinearMinter.Amount)
		}

		if oldMinterPeriod.PeriodicReductionMinter == nil {
			require.Nil(t, newMinters[i].ExponentialStepMinting)
		} else {
			require.Equal(t,
				newMinters[i].ExponentialStepMinting.Amount,
				oldMinterPeriod.PeriodicReductionMinter.MintAmount.MulRaw(int64(oldMinterPeriod.PeriodicReductionMinter.ReductionPeriodLength)),
			)
			require.Equal(t,
				newMinters[i].ExponentialStepMinting.StepDuration.Seconds(),
				float64(oldMinterPeriod.PeriodicReductionMinter.MintPeriod*oldMinterPeriod.PeriodicReductionMinter.ReductionPeriodLength),
			)
			require.Equal(t, newMinters[i].ExponentialStepMinting.AmountMultiplier, oldMinterPeriod.PeriodicReductionMinter.ReductionFactor)
		}
	}
}

func createV100MinterPeriod(
	position int32,
	endTime *time.Time,
	minterType string,
	periodicReductionMinter *v1.PeriodicReductionMinter,
	timeLinearMinter *v1.TimeLinearMinter,
) *v1.MintingPeriod {
	return &v1.MintingPeriod{
		Position:                position,
		PeriodicReductionMinter: periodicReductionMinter,
		TimeLinearMinter:        timeLinearMinter,
		PeriodEnd:               endTime,
		Type:                    minterType,
	}
}

func createV1TimeLinearMinter(
	amount math.Int,
) *v1.TimeLinearMinter {
	return &v1.TimeLinearMinter{
		Amount: amount,
	}
}

func createV1TimePeriodicReductionMinter(
	reductionPeriodLength int32,
	mintPeriod int32,
	reductionFactor sdk.Dec,
	mintAmount math.Int,
) *v1.PeriodicReductionMinter {
	return &v1.PeriodicReductionMinter{
		ReductionPeriodLength: reductionPeriodLength,
		ReductionFactor:       reductionFactor,
		MintAmount:            mintAmount,
		MintPeriod:            mintPeriod,
	}
}
