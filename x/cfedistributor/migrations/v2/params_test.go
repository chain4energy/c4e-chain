package v2_test

import (
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	cfedistributortestutils "github.com/chain4energy/c4e-chain/testutil/module/cfedistributor"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	v1 "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v1"
	v2 "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v2"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestMigrationSubDistributorsCorrectOrder(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)

	subDistributorSourceMain := createOldSubDistributor(types.ModuleAccount, types.Main, types.BaseAccount, "CUSTOM_ID")
	subDistributorSourceMain.Sources = subDistributorSourceMain.Sources[:1]

	subDistributorShareMain := createOldSubDistributor(types.BaseAccount, types.ModuleAccount, types.Main, "CUSTOM_ID")
	subDistributorShareMain.Destination.Share = subDistributorShareMain.Destination.Share[:1]

	oldSubDistributors := []v1.SubDistributor{
		createOldSubDistributor(types.ModuleAccount, types.BaseAccount, types.InternalAccount, "CUSTOM_ID"),
		createOldSubDistributor(types.InternalAccount, types.ModuleAccount, types.BaseAccount, "CUSTOM_ID"),
		createOldSubDistributor(types.BaseAccount, types.InternalAccount, types.ModuleAccount, "CUSTOM_ID"),
		createOldSubDistributor(types.Main, types.InternalAccount, types.ModuleAccount, "CUSTOM_ID"),
		subDistributorShareMain,
		subDistributorSourceMain,
	}
	setV2Subdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV1ToV2(t, ctx, testUtil, false, "")
}

func TestMigrationSubDistributorsWrongOrder(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v1.SubDistributor{
		createOldSubDistributor(types.InternalAccount, types.BaseAccount, types.ModuleAccount, "CUSTOM_ID"),
		createOldSubDistributor(types.Main, types.InternalAccount, types.ModuleAccount, "CUSTOM_ID"),
	}

	setV2Subdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV1ToV2(t, ctx, testUtil, true, "wrong order of subdistributors, after each occurrence of a subdistributor with the destination of internal or main account type there must be exactly one occurrence of a subdistributor with the source of internal account type, account id: MAIN")
}

func TestMigrationSubDistributorsDuplicates(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v1.SubDistributor{
		createOldSubDistributor(types.InternalAccount, types.BaseAccount, types.ModuleAccount, "CUSTOM_ID"),
		createOldSubDistributor(types.BaseAccount, types.InternalAccount, types.ModuleAccount, "CUSTOM_ID"),
		createOldSubDistributor(types.BaseAccount, types.Main, types.BaseAccount, "CUSTOM_ID"),
	}
	setV2Subdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV1ToV2(t, ctx, testUtil, true, "same MAIN account cannot occur twice within one subdistributor, subdistributor name: "+oldSubDistributors[2].Name)
}

func TestMigrationSubDistributorsWrongAccType(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v1.SubDistributor{
		createOldSubDistributor(types.InternalAccount, types.BaseAccount, types.ModuleAccount, "CUSTOM_ID"),
		createOldSubDistributor(types.BaseAccount, types.Main, types.ModuleAccount, "CUSTOM_ID"),
		createOldSubDistributor(types.BaseAccount, "WRONG_ACCOUNT_TYPE", types.ModuleAccount, "CUSTOM_ID"),
	}

	setV2Subdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV1ToV2(t, ctx, testUtil, true, "subdistributor "+oldSubDistributors[2].Name+" source with id \"CUSTOM_ID\" validation error: account \"CUSTOM_ID\" is of the wrong type: WRONG_ACCOUNT_TYPE")
}

func TestMigrationSubDistributorsWrongModuleAccount(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v1.SubDistributor{
		createOldSubDistributor(types.BaseAccount, types.Main, types.ModuleAccount, "WRONG_CUSTOM_ID"),
	}

	setV2Subdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV1ToV2(t, ctx, testUtil, true, "subdistributor "+oldSubDistributors[0].Name+" destinations validation error: destination share "+oldSubDistributors[0].Destination.Share[0].Name+" destination account validation error: module account \"WRONG_CUSTOM_ID_custom_siffix_0\" doesn't exist in maccPerms")
}

func setV2Subdistributors(t *testing.T, ctx sdk.Context, testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils, subdistributors []v1.SubDistributor) {
	store := newStore(ctx, testUtil)
	bz, err := codec.NewLegacyAmino().MarshalJSON(subdistributors)
	require.NoError(t, err)
	store.Set(types.KeySubDistributors, bz)
}

func newStore(ctx sdk.Context, testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils) prefix.Store {
	return prefix.NewStore(ctx.KVStore(testUtil.StoreKey), append([]byte((testUtil.Subspace.Name())), '/'))
}

func MigrateParamsV1ToV2(
	t *testing.T,
	ctx sdk.Context,
	testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils,
	expectError bool, errorMessage string,
) {
	cfedistributortestutils.SetTestMaccPerms()
	var oldSubDistributors []v1.SubDistributor
	store := newStore(ctx, testUtil)
	distributors := store.Get(types.KeySubDistributors)
	err := codec.NewLegacyAmino().UnmarshalJSON(distributors, &oldSubDistributors)
	require.NoError(t, err)

	err = v2.MigrateParams(ctx, &testUtil.Subspace)
	if expectError {
		require.EqualError(t, err, errorMessage)
		return
	}
	require.NoError(t, err)
	var newSubDistributors []types.SubDistributor
	newSubDistributorsRaw := store.Get(types.KeySubDistributors)
	err = codec.NewLegacyAmino().UnmarshalJSON(newSubDistributorsRaw, &newSubDistributors)
	require.NoError(t, err)

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
) v1.SubDistributor {
	var sources []*v1.Account
	mainAcc := v1.Account{
		Id:   cfedistributortestutils.GetAccountTestId(id, "", sourceType),
		Type: sourceType,
	}
	sources = append(sources, &mainAcc)
	for i := 0; i < 5; i++ {
		randomAccount := v1.Account{
			Id:   cfedistributortestutils.GetAccountTestId(id+"_custom_siffix_"+strconv.Itoa(i), "", sourceType),
			Type: sourceType,
		}
		sources = append(sources, &randomAccount)
	}

	var shares []*v1.Share

	for i := 0; i < 5; i++ {
		share := v1.Share{
			Name: helpers.RandStringOfLength(10),
			Account: v1.Account{
				Id:   cfedistributortestutils.GetAccountTestId(id+"_custom_siffix_"+strconv.Itoa(i), "", destinationShareType),
				Type: destinationShareType,
			},
			Percent: sdk.NewDec(5),
		}
		shares = append(shares, &share)
	}

	return v1.SubDistributor{
		Name: helpers.RandStringOfLength(10),
		Destination: v1.Destination{
			Account: v1.Account{
				Id:   cfedistributortestutils.GetAccountTestId(id, "", destinationType),
				Type: destinationType,
			},
			BurnShare: &v1.BurnShare{
				Percent: sdk.NewDec(25),
			},
			Share: shares,
		},
		Sources: sources,
	}
}
