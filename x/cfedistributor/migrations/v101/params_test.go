package v101_test

import (
	subdistributortestutils "github.com/chain4energy/c4e-chain/testutil/module/cfedistributor/subdistributor"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	v100cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v100"
	v101cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v101"
	v100cfevesting "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v100"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"testing"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMigrationSubDistributors(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(5, 0)
	testUtil, _, ctx, paramsSubspace := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	SetupV100SubdistributorParams(testUtil, paramsSubspace, ctx)
	MigrateParamsV100ToV101(t, testUtil, ctx, paramsSubspace)
}

func SetupV100SubdistributorParams(testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils, paramsSubspace typesparams.Subspace, ctx sdk.Context) v100cfedistributor.SubDistributor {
	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, subdistributortestutils.PrepareBurningDistributor(subdistributortestutils.MainCollector))
	subdistributors = append(subdistributors, subdistributortestutils.PrepareInflationToPassAcoutSubDistr(subdistributortestutils.MainCollector))
	subdistributors = append(subdistributors, subdistributortestutils.PrepareInflationSubDistributor(subdistributortestutils.MainCollector, true))
	paramsSubspace.SetParamSet()
	return subdistributors
}

func SetV100AccountVestingPools(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, accountVestingPools v100cfevesting.AccountVestingPools) {
	store := prefix.NewStore(ctx.KVStore(storeKey), v100cfevesting.AccountVestingPoolsKeyPrefix)
	av := cdc.MustMarshal(&accountVestingPools)
	store.Set([]byte(accountVestingPools.Address), av)
}

func MigrateParamsV100ToV101(t *testing.T, testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils, ctx sdk.Context, paramsSubspace typesparams.Subspace) {
	oldParams := testUtil.GetC4eDistributorKeeper().GetParams(ctx)
	err := v101cfedistributor.MigrateParams(ctx, &paramsSubspace)
	newParams := testUtil.GetC4eDistributorKeeper().GetParams(ctx)
	newSubdistributors := newParams.SubDistributors
	require.NoError(t, err)

	require.EqualValues(t, len(oldParams.SubDistributors), len(newSubdistributors))
	for i, oldSubdistributor := range oldParams.SubDistributors {
		require.EqualValues(t, oldSubdistributor.Name, newSubdistributors[i].Name)
	}
}

func getAllv100DistributorStates(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (list []v100cfedistributor.State, err error) {

	prefixStore := prefix.NewStore(ctx.KVStore(storeKey), v100cfedistributor.RemainsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v100cfedistributor.State
		err := cdc.Unmarshal(iterator.Value(), &val)
		if err != nil {
			return nil, err
		}
		list = append(list, val)
	}
	return list, nil

}

func generateV100SubDistributor() v100cfedistributor.SubDistributor {
	destAccount := v100cfedistributor.Account{
		Id:   helpers.RandStringOfLength(10),
		Type: "MAIN",
	}
	burnShare := sdk.MustNewDecFromStr("0.51")
	destination := v100cfedistributor.Destination{
		Account: destAccount,
		Share:   nil,
		BurnShare: &v100cfedistributor.BurnShare{
			Percent: burnShare,
		},
	}

	distributor1 := v100cfedistributor.SubDistributor{
		Name: helpers.RandStringOfLength(10),
		Sources: []*v100cfedistributor.Account{
			{
				Id:   helpers.RandStringOfLength(10),
				Type: types.MODULE_ACCOUNT,
			},
		},
		Destination: destination,
	}

	return distributor1
}
