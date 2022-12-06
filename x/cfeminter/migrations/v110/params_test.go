package v110_test

import (
	"github.com/chain4energy/c4e-chain/testutil/common"
	"github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v101"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v110"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/stretchr/testify/require"
	"time"

	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationSubDistributorsCorrectOrder(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	timeLinearMinter := createOldTimeLinearMinter(sdk.NewInt(10000))
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)
	V101MintingPeriods := []*v101.MintingPeriod{
		{
			Position:         1,
			Type:             "TIME_LINEAR_MINTER",
			TimeLinearMinter: &timeLinearMinter,
			PeriodEnd:        &endTime,
		},
	}
	setV101MinterConfig(t, ctx, &keeperData, startTime, V101MintingPeriods)
	MigrateParamsV100ToV101(t, ctx, *k, &keeperData, "")
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
	errorMessage string,
) {
	oldMinterConfig := getV101MinterConfig(ctx, keeperData)

	err := v110.MigrateParams(ctx, keeperData.StoreKey, &keeperData.Subspace)
	if len(errorMessage) > 0 {
		require.Equal(t, err.Error(), errorMessage)
		return
	}
	require.NoError(t, err)

	newParams := keeper.GetParams(ctx)
	newMinterConfig := newParams.MinterConfig

	require.EqualValues(t, len(newMinterConfig.Minters), len(oldMinterConfig.Periods))
	//for i, oldSubDistributor := range oldSubDistributors {
	//	require.EqualValues(t, newSubDistributors[i].Name, oldSubDistributor.Name)
	//	require.EqualValues(t, newSubDistributors[i].Destinations.BurnShare, oldSubDistributor.Destination.BurnShare.Percent.Quo(sdk.NewDec(100)))
	//	require.EqualValues(t, newSubDistributors[i].Destinations.PrimaryShare.Id, oldSubDistributor.Destination.Account.Id)
	//	require.EqualValues(t, newSubDistributors[i].Destinations.PrimaryShare.Type, oldSubDistributor.Destination.Account.Type)
	//
	//	require.EqualValues(t, len(newSubDistributors[i].Destinations.Shares), len(oldSubDistributor.Destination.Share))
	//	for j, oldShare := range oldSubDistributor.Destination.Share {
	//		require.EqualValues(t, newSubDistributors[i].Destinations.Shares[j].Share, oldShare.Percent.Quo(sdk.NewDec(100)))
	//		require.EqualValues(t, newSubDistributors[i].Destinations.Shares[j].Name, oldShare.Name)
	//		require.EqualValues(t, newSubDistributors[i].Destinations.Shares[j].Destination.Id, oldShare.Account.Id)
	//		require.EqualValues(t, newSubDistributors[i].Destinations.Shares[j].Destination.Type, oldShare.Account.Type)
	//	}
	//
	//	require.EqualValues(t, len(newSubDistributors[i].Sources), len(oldSubDistributor.Sources))
	//	for j, oldSource := range oldSubDistributor.Sources {
	//		require.EqualValues(t, newSubDistributors[i].Sources[j].Id, oldSource.Id)
	//		require.EqualValues(t, newSubDistributors[i].Sources[j].Type, oldSource.Type)
	//	}
	//}
}

func createOldMinterPeriod(
	endTime *time.Time,
	periodicReductionMinter *v101.PeriodicReductionMinter,
	timeLinearMinter *v101.TimeLinearMinter,
	position int32,
	minterType string,
) v101.MintingPeriod {
	return v101.MintingPeriod{
		Position:                position,
		PeriodicReductionMinter: periodicReductionMinter,
		TimeLinearMinter:        timeLinearMinter,
		PeriodEnd:               endTime,
		Type:                    minterType,
	}
}

func createOldTimeLinearMinter(
	amount sdk.Int,
) v101.TimeLinearMinter {
	return v101.TimeLinearMinter{
		Amount: amount,
	}
}

func createOldTimePeriodicReductionMinter(
	reductionPeriodLength int32,
	mintPeriod int32,
	reductionFactor sdk.Dec,
	mintAmount sdk.Int,
) v101.PeriodicReductionMinter {
	return v101.PeriodicReductionMinter{
		ReductionPeriodLength: reductionPeriodLength,
		ReductionFactor:       reductionFactor,
		MintAmount:            mintAmount,
		MintPeriod:            mintPeriod,
	}
}
