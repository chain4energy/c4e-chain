package v110_test

import (
	"github.com/chain4energy/c4e-chain/testutil/common"
	"github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v101"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v110"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/stretchr/testify/require"
	"time"

	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationLinearMinting(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	timeLinearMinter := createV101TimeLinearMinter(sdk.NewInt(10000))
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)
	V101MintingPeriods := []*v101.MintingPeriod{
		createV100MinterPeriod(1, &endTime, "TIME_LINEAR_MINTER", nil, timeLinearMinter),
		createV100MinterPeriod(2, nil, "NO_MINTING", nil, nil),
	}
	setV101MinterConfig(t, ctx, &keeperData, startTime, V101MintingPeriods)
	MigrateParamsV100ToV101(t, ctx, *k, &keeperData, false, "")
}

func TestMigrationExponentialStepMinting(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	periodicReductionMinter := createV101TimePeriodicReductionMinter(4, 100000, sdk.MustNewDecFromStr("0.5"), sdk.NewInt(10000))
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)
	V101MintingPeriods := []*v101.MintingPeriod{
		createV100MinterPeriod(1, &endTime, "PERIODIC_REDUCTION_MINTER", periodicReductionMinter, nil),
		createV100MinterPeriod(2, nil, "NO_MINTING", nil, nil),
	}
	setV101MinterConfig(t, ctx, &keeperData, startTime, V101MintingPeriods)
	MigrateParamsV100ToV101(t, ctx, *k, &keeperData, false, "")
}

func TestMigrationLinearMintingAndExponentialStepMinting(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	timeLinearMinter := createV101TimeLinearMinter(sdk.NewInt(10000))
	periodicReductionMinter := createV101TimePeriodicReductionMinter(4, 100000, sdk.MustNewDecFromStr("0.5"), sdk.NewInt(10000))
	startTime := time.Now()
	endTime1 := startTime.Add(time.Hour)
	endTime2 := endTime1.Add(time.Hour)
	V101MintingPeriods := []*v101.MintingPeriod{
		createV100MinterPeriod(1, &endTime1, "TIME_LINEAR_MINTER", nil, timeLinearMinter),
		createV100MinterPeriod(2, &endTime2, "PERIODIC_REDUCTION_MINTER", periodicReductionMinter, nil),
		createV100MinterPeriod(3, nil, "NO_MINTING", nil, nil),
	}
	setV101MinterConfig(t, ctx, &keeperData, startTime, V101MintingPeriods)
	MigrateParamsV100ToV101(t, ctx, *k, &keeperData, false, "")
}

func TestMigrationNoMinters(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	startTime := time.Now()
	V101MintingPeriods := []*v101.MintingPeriod{}
	setV101MinterConfig(t, ctx, &keeperData, startTime, V101MintingPeriods)
	MigrateParamsV100ToV101(t, ctx, *k, &keeperData, true, "no minters defined")
}

func TestMigrationWrongMinterPosition(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	timeLinearMinter := createV101TimeLinearMinter(sdk.NewInt(10000))
	periodicReductionMinter := createV101TimePeriodicReductionMinter(4, 100000, sdk.MustNewDecFromStr("0.5"), sdk.NewInt(10000))
	startTime := time.Now()
	endTime1 := startTime.Add(time.Hour)
	endTime2 := endTime1.Add(time.Hour)
	V101MintingPeriods := []*v101.MintingPeriod{
		createV100MinterPeriod(1, &endTime1, "TIME_LINEAR_MINTER", nil, timeLinearMinter),
		createV100MinterPeriod(1, &endTime2, "PERIODIC_REDUCTION_MINTER", periodicReductionMinter, nil),
		createV100MinterPeriod(3, nil, "NO_MINTING", nil, nil),
	}
	setV101MinterConfig(t, ctx, &keeperData, startTime, V101MintingPeriods)
	MigrateParamsV100ToV101(t, ctx, *k, &keeperData, true, "missing minter with sequence id 2")
}

func TestMigrationWrongMintingStartTime(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	timeLinearMinter := createV101TimeLinearMinter(sdk.NewInt(10000))
	periodicReductionMinter := createV101TimePeriodicReductionMinter(4, 100000, sdk.MustNewDecFromStr("0.5"), sdk.NewInt(10000))
	startTime := time.Now()
	endTime1 := startTime.Add(time.Hour)
	endTime2 := endTime1.Add(time.Hour)
	V101MintingPeriods := []*v101.MintingPeriod{
		createV100MinterPeriod(1, &endTime1, "TIME_LINEAR_MINTER", nil, timeLinearMinter),
		createV100MinterPeriod(2, &endTime2, "PERIODIC_REDUCTION_MINTER", periodicReductionMinter, nil),
		createV100MinterPeriod(3, nil, "NO_MINTING", nil, nil),
	}
	setV101MinterConfig(t, ctx, &keeperData, endTime2, V101MintingPeriods)
	MigrateParamsV100ToV101(t, ctx, *k, &keeperData, true, "first minter end must be bigger than minter start")
}

func TestMigrationWrongMinterType(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	timeLinearMinter := createV101TimeLinearMinter(sdk.NewInt(10000))
	startTime := time.Now()
	endTime1 := startTime.Add(time.Hour)
	V101MintingPeriods := []*v101.MintingPeriod{
		createV100MinterPeriod(1, &endTime1, "WRONG_MINTER_TYPE", nil, timeLinearMinter),
	}
	setV101MinterConfig(t, ctx, &keeperData, startTime, V101MintingPeriods)
	MigrateParamsV100ToV101(t, ctx, *k, &keeperData, true, "wrong minting period type")
}

func TestMigrationWrongExponentialStepMinting(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	periodicReductionMinter := createV101TimePeriodicReductionMinter(4, 0, sdk.MustNewDecFromStr("0.5"), sdk.NewInt(10000))
	startTime := time.Now()
	endTime1 := startTime.Add(time.Hour)
	V101MintingPeriods := []*v101.MintingPeriod{
		createV100MinterPeriod(1, &endTime1, "PERIODIC_REDUCTION_MINTER", periodicReductionMinter, nil),
		createV100MinterPeriod(2, nil, "NO_MINTING", nil, nil),
	}
	setV101MinterConfig(t, ctx, &keeperData, startTime, V101MintingPeriods)
	MigrateParamsV100ToV101(t, ctx, *k, &keeperData, true, "minter with id 1 validation error: ExponentialStepMintingType error: stepDuration must be bigger than 0")
}

func TestMigrationWrongLinearMinting(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	timeLinearMinter := createV101TimeLinearMinter(sdk.NewInt(-10000))
	startTime := time.Now()
	endTime1 := startTime.Add(time.Hour)
	V101MintingPeriods := []*v101.MintingPeriod{
		createV100MinterPeriod(1, &endTime1, "TIME_LINEAR_MINTER", nil, timeLinearMinter),
		createV100MinterPeriod(3, nil, "NO_MINTING", nil, nil),
	}

	setV101MinterConfig(t, ctx, &keeperData, startTime, V101MintingPeriods)
	MigrateParamsV100ToV101(t, ctx, *k, &keeperData, true, "minter with id 1 validation error: LinearMintingType error: amount cannot be less than 0")
}

func setV101MinterConfig(t *testing.T, ctx sdk.Context, keeperData *common.AdditionalKeeperData, startTime time.Time, mintingPeriods []*v101.MintingPeriod) {
	minter := v101.Minter{
		Start:   startTime,
		Periods: mintingPeriods,
	}
	store := newStore(ctx, keeperData)
	bz, err := codec.NewLegacyAmino().MarshalJSON(minter)
	require.NoError(t, err)
	store.Set(v101.KeyMinter, bz)
}

func newStore(ctx sdk.Context, testUtil *common.AdditionalKeeperData) prefix.Store {
	return prefix.NewStore(ctx.KVStore(testUtil.StoreKey), append([]byte((testUtil.Subspace.Name())), '/'))
}

func getV101MinterConfig(ctx sdk.Context, keeperData *common.AdditionalKeeperData) (oldMinterConfig v101.Minter) {
	oldMinterConfigRaw := keeperData.Subspace.GetRaw(ctx, v101.KeyMinter)
	if err := codec.NewLegacyAmino().UnmarshalJSON(oldMinterConfigRaw, &oldMinterConfig); err != nil {
		panic(err)
	}
	return
}

func MigrateParamsV100ToV101(
	t *testing.T,
	ctx sdk.Context,
	keeper keeper.Keeper,
	keeperData *common.AdditionalKeeperData,
	expectError bool, errorMessage string,
) {
	oldMinterConfig := getV101MinterConfig(ctx, keeperData)

	err := v110.MigrateParams(ctx, &keeperData.Subspace)
	if expectError {
		require.EqualError(t, err, errorMessage)
		return
	}
	require.NoError(t, err)

	newParams := keeper.GetParams(ctx)
	newMinterConfig := newParams.MinterConfig

	require.EqualValues(t, len(newMinterConfig.Minters), len(oldMinterConfig.Periods))
	newMinters := newMinterConfig.Minters
	for i, oldMinterPeriod := range oldMinterConfig.Periods {
		require.EqualValues(t, newMinters[i].SequenceId, oldMinterPeriod.Position)
		require.EqualValues(t, newMinters[i].EndTime, oldMinterPeriod.PeriodEnd)

		switch oldMinterPeriod.Type {
		case "TIME_LINEAR_MINTER":
			require.EqualValues(t, newMinters[i].Type, types.LinearMintingType)
			break
		case "PERIODIC_REDUCTION_MINTER":
			require.EqualValues(t, newMinters[i].Type, types.ExponentialStepMintingType)
			break
		case "NO_MINTING":
			require.EqualValues(t, newMinters[i].Type, types.NoMintingType)
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
	periodicReductionMinter *v101.PeriodicReductionMinter,
	timeLinearMinter *v101.TimeLinearMinter,
) *v101.MintingPeriod {
	return &v101.MintingPeriod{
		Position:                position,
		PeriodicReductionMinter: periodicReductionMinter,
		TimeLinearMinter:        timeLinearMinter,
		PeriodEnd:               endTime,
		Type:                    minterType,
	}
}

func createV101TimeLinearMinter(
	amount sdk.Int,
) *v101.TimeLinearMinter {
	return &v101.TimeLinearMinter{
		Amount: amount,
	}
}

func createV101TimePeriodicReductionMinter(
	reductionPeriodLength int32,
	mintPeriod int32,
	reductionFactor sdk.Dec,
	mintAmount sdk.Int,
) *v101.PeriodicReductionMinter {
	return &v101.PeriodicReductionMinter{
		ReductionPeriodLength: reductionPeriodLength,
		ReductionFactor:       reductionFactor,
		MintAmount:            mintAmount,
		MintPeriod:            mintPeriod,
	}
}
