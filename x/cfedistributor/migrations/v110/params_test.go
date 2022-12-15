package v110_test

import (
	"github.com/chain4energy/c4e-chain/testutil/module/cfedistributor"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v101"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v110"
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

	oldSubDistributors := []v101.SubDistributor{
		createOldSubDistributor(types.MODULE_ACCOUNT, types.BASE_ACCOUNT, types.INTERNAL_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, types.BASE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.MAIN, types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		subDistributorShareMain,
		subDistributorSourceMain,
	}
	setV101Subdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV101ToV110(t, ctx, testUtil, "")
}

func TestMigrationSubDistributorsWrongOrder(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v101.SubDistributor{
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.MAIN, types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
	}

	setV101Subdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV101ToV110(t, ctx, testUtil, "wrong order of subdistributors, after each occurrence of a subdistributor with the destination of internal or main account type there must be exactly one occurrence of a subdistributor with the source of internal account type, account id: MAIN")
}

func TestMigrationSubDistributorsDuplicates(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v101.SubDistributor{
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.BASE_ACCOUNT, "CUSTOM_ID"),
	}
	setV101Subdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV101ToV110(t, ctx, testUtil, "same MAIN account cannot occur twice within one subdistributor, subdistributor name: "+oldSubDistributors[2].Name)
}

func TestMigrationSubDistributorsWrongAccType(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v101.SubDistributor{
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, "WRONG_ACCOUNT_TYPE", types.MODULE_ACCOUNT, "CUSTOM_ID"),
	}

	setV101Subdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV101ToV110(t, ctx, testUtil, "account \"CUSTOM_ID\" is of the wrong type: WRONG_ACCOUNT_TYPE")
}

func TestMigrationSubDistributorsWrongModuleAccount(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v101.SubDistributor{
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.MODULE_ACCOUNT, "WRONG_CUSTOM_ID"),
	}

	setV101Subdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV101ToV110(t, ctx, testUtil, "module account \"WRONG_CUSTOM_ID_custom_siffix_0\" doesn't exist in maccPerms")
}

func setV101Subdistributors(t *testing.T, ctx sdk.Context, testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils, subdistributors []v101.SubDistributor) {
	store := newStore(ctx, testUtil)
	bz, err := codec.NewLegacyAmino().MarshalJSON(subdistributors)
	require.NoError(t, err)
	store.Set(types.KeySubDistributors, bz)
}

func newStore(ctx sdk.Context, testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils) prefix.Store {
	return prefix.NewStore(ctx.KVStore(testUtil.StoreKey), append([]byte((testUtil.Subspace.Name())), '/'))
}

func MigrateParamsV101ToV110(
	t *testing.T,
	ctx sdk.Context,
	testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils,
	errorMessage string,
) {
	types.SetMaccPerms(cfedistributor.GetTestMaccPerms())
	var oldSubDistributors []v101.SubDistributor
	store := newStore(ctx, testUtil)
	distributors := store.Get(types.KeySubDistributors)
	err := codec.NewLegacyAmino().UnmarshalJSON(distributors, &oldSubDistributors)
	require.NoError(t, err)

	err = v110.MigrateParams(ctx, &testUtil.Subspace)
	if len(errorMessage) > 0 {
		require.Equal(t, err.Error(), errorMessage)
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
) v101.SubDistributor {
	var sources []*v101.Account
	mainAcc := v101.Account{
		Id:   id,
		Type: sourceType,
	}
	sources = append(sources, &mainAcc)
	for i := 0; i < 5; i++ {
		randomAccount := v101.Account{
			Id:   id + "_custom_siffix_" + strconv.Itoa(i),
			Type: sourceType,
		}
		sources = append(sources, &randomAccount)
	}

	var shares []*v101.Share

	for i := 0; i < 5; i++ {
		share := v101.Share{
			Name: helpers.RandStringOfLength(10),
			Account: v101.Account{
				Id:   id + "_custom_siffix_" + strconv.Itoa(i),
				Type: destinationShareType,
			},
			Percent: sdk.NewDec(5),
		}
		shares = append(shares, &share)
	}

	return v101.SubDistributor{
		Name: helpers.RandStringOfLength(10),
		Destination: v101.Destination{
			Account: v101.Account{
				Id:   id,
				Type: destinationType,
			},
			BurnShare: &v101.BurnShare{
				Percent: sdk.NewDec(25),
			},
			Share: shares,
		},
		Sources: sources,
	}
}
