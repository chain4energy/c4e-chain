package v101_test

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	v100cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v100"
	v101cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v101"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/stretchr/testify/require"
	"strconv"

	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationSubDistributorsCorrectOrder(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)

	subDistributorSourceMain := createOldSubDistributor(types.MODULE_ACCOUNT, types.MAIN, types.BASE_ACCOUNT, "CUSTOM_ID")
	subDistributorSourceMain.Sources = subDistributorSourceMain.Sources[:1]

	subDistributorShareMain := createOldSubDistributor(types.BASE_ACCOUNT, types.MODULE_ACCOUNT, types.MAIN, "CUSTOM_ID")
	subDistributorShareMain.Destination.Share = subDistributorShareMain.Destination.Share[:1]

	oldSubDistributors := []v100cfedistributor.SubDistributor{
		createOldSubDistributor(types.MODULE_ACCOUNT, types.BASE_ACCOUNT, types.INTERNAL_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, types.BASE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.MAIN, types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		subDistributorShareMain,
		subDistributorSourceMain,
	}
	setOldSubdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV100ToV101(t, ctx, testUtil, false)
}

func TestMigrationSubDistributorsWrongOrder(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v100cfedistributor.SubDistributor{
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.MODULE_ACCOUNT, "CUSTOM_ID"),
	}

	setOldSubdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV100ToV101(t, ctx, testUtil, true)
}

func TestMigrationSubDistributorsDuplicates(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v100cfedistributor.SubDistributor{
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.BASE_ACCOUNT, "CUSTOM_ID"),
	}
	setOldSubdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV100ToV101(t, ctx, testUtil, true)
}

func TestMigrationSubDistributorsWrongAccType(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v100cfedistributor.SubDistributor{
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, "WRONG_ACCOUNT_TYPE", types.MODULE_ACCOUNT, "CUSTOM_ID"),
	}

	setOldSubdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV100ToV101(t, ctx, testUtil, true)
}

func setOldSubdistributors(t *testing.T, ctx sdk.Context, testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils, subdistributors []v100cfedistributor.SubDistributor) {
	store := newStore(ctx, testUtil)
	bz, err := codec.NewLegacyAmino().MarshalJSON(subdistributors)
	require.NoError(t, err)
	store.Set(types.KeySubDistributors, bz)
}

func newStore(ctx sdk.Context, testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils) prefix.Store {
	return prefix.NewStore(ctx.KVStore(testUtil.StoreKey), append([]byte((testUtil.Subspace.Name())), '/'))
}

func MigrateParamsV100ToV101(
	t *testing.T,
	ctx sdk.Context,
	testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils,
	wantError bool,
) {
	var oldSubDistributors []v100cfedistributor.SubDistributor
	store := newStore(ctx, testUtil)
	distributors := store.Get(types.KeySubDistributors)
	err := codec.NewLegacyAmino().UnmarshalJSON(distributors, &oldSubDistributors)
	require.NoError(t, err)

	err = v101cfedistributor.MigrateParams(ctx, testUtil.StoreKey, &testUtil.Subspace)
	if wantError {
		require.Error(t, err)
		return
	}
	require.NoError(t, err)

	newParams := testUtil.GetC4eDistributorKeeper().GetParams(ctx)
	newSubDistributors := newParams.SubDistributors

	require.EqualValues(t, len(newSubDistributors), len(oldSubDistributors))
	for i, oldSubDistributor := range oldSubDistributors {
		require.EqualValues(t, newSubDistributors[i].Name, oldSubDistributor.Name)
		require.EqualValues(t, newSubDistributors[i].Destinations.BurnShare, oldSubDistributor.Destination.BurnShare.Percent.Quo(sdk.NewDec(100)))
		require.EqualValues(t, newSubDistributors[i].Destinations.PrimaryShare.Id, oldSubDistributor.Destination.Account.Id)
		require.EqualValues(t, newSubDistributors[i].Destinations.PrimaryShare.Type, oldSubDistributor.Destination.Account.Type)

		require.EqualValues(t, len(newSubDistributors[i].Destinations.Shares), len(oldSubDistributor.Destination.Share))
		for j, oldShare := range oldSubDistributor.Destination.Share {
			require.EqualValues(t, newSubDistributors[i].Destinations.Shares[j].Share, oldShare.Percent.Quo(sdk.NewDec(100)))
			require.EqualValues(t, newSubDistributors[i].Destinations.Shares[j].Name, oldShare.Name)
			require.EqualValues(t, newSubDistributors[i].Destinations.Shares[j].Destination.Id, oldShare.Account.Id)
			require.EqualValues(t, newSubDistributors[i].Destinations.Shares[j].Destination.Type, oldShare.Account.Type)
		}

		require.EqualValues(t, len(newSubDistributors[i].Sources), len(oldSubDistributor.Sources))
		for j, oldSource := range oldSubDistributor.Sources {
			require.EqualValues(t, newSubDistributors[i].Sources[j].Id, oldSource.Id)
			require.EqualValues(t, newSubDistributors[i].Sources[j].Type, oldSource.Type)
		}
	}
}

func createOldSubDistributor(
	destinationType string,
	sourceType string,
	destinationShareType string,
	id string,
) v100cfedistributor.SubDistributor {
	var sources []*v100cfedistributor.Account
	mainAcc := v100cfedistributor.Account{
		Id:   id,
		Type: sourceType,
	}
	sources = append(sources, &mainAcc)
	for i := 0; i < 5; i++ {
		randomAccount := v100cfedistributor.Account{
			Id:   id + "custom_siffix_" + strconv.Itoa(i),
			Type: sourceType,
		}
		sources = append(sources, &randomAccount)
	}

	var shares []*v100cfedistributor.Share

	for i := 0; i < 5; i++ {
		share := v100cfedistributor.Share{
			Name: helpers.RandStringOfLength(10),
			Account: v100cfedistributor.Account{
				Id:   id + "custom_siffix_" + strconv.Itoa(i),
				Type: destinationShareType,
			},
			Percent: sdk.NewDec(5),
		}
		shares = append(shares, &share)
	}

	return v100cfedistributor.SubDistributor{
		Name: helpers.RandStringOfLength(10),
		Destination: v100cfedistributor.Destination{
			Account: v100cfedistributor.Account{
				Id:   id,
				Type: destinationType,
			},
			BurnShare: &v100cfedistributor.BurnShare{
				Percent: sdk.NewDec(25),
			},
			Share: shares,
		},
		Sources: sources,
	}
}
