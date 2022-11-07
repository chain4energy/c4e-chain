package v101

import (
	"fmt"
	v100cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v100"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

//getAllV100AccountVestingPoolsAndDelete returns all old version AccountVestingPools and deletes them from the KVStore
func getAllV100SubDistributorStatesAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (list []v100cfedistributor.State, err error) {
	prefixStore := prefix.NewStore(store, v100cfedistributor.RemainsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v100cfedistributor.State
		err := cdc.Unmarshal(iterator.Value(), &val)
		if err != nil {
			return nil, err
		}
		list = append(list, val)
		prefixStore.Delete(iterator.Key())
	}
	return
}

func setNewSubdistributorStates(store sdk.KVStore, cdc codec.BinaryCodec, oldStates []v100cfedistributor.State) error {
	prefixStore := prefix.NewStore(store, types.StateKeyPrefix)
	for _, oldState := range oldStates {
		newAccount := types.Account{
			Id:   oldState.Account.Id,
			Type: oldState.Account.Id,
		}
		newState := types.State{
			Account: &newAccount,
			Burn:    oldState.Burn,
			Remains: oldState.CoinsStates,
		}

		av, err := cdc.Marshal(&newState)
		if err != nil {
			return err
		}
		prefixStore.Set([]byte(GetStateKey(newState)), av)
	}
	return nil
}

func GetStateKey(state types.State) string {
	if state.Account != nil && state.Account.Id != "" {
		return state.Account.Id
	} else {
		//its state burn
		return types.BurnStateKey
	}
}

func migrateSubdistributorStates(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldAccountVestingPools, err := getAllV100SubDistributorStatesAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewSubdistributorStates(store, cdc, oldAccountVestingPools)
}

func UpdateParams(ctx sdk.Context, paramStore *paramtypes.Subspace, cdc codec.BinaryCodec) error {
	var res []v100cfedistributor.SubDistributor
	rawDistributors := paramStore.GetRaw(ctx, types.KeySubDistributors)
	var Cdc = codec.NewLegacyAmino()
	err := cdc.Unmarshal(rawDistributors, &res)
	var newSubdistributors []types.SubDistributor
	//for _, oldSubdistributor := range res {
	//	var newShares []*types.DestinationShare
	//	for _, oldShare := range oldSubdistributor.Destination.Share {
	//		newShare := types.DestinationShare{
	//			Share: oldShare.Percent,
	//			Destination: types.Account{
	//				Id:   oldShare.Account.Id,
	//				Type: oldShare.Account.Type,
	//			},
	//			Name: oldShare.Name,
	//		}
	//		newShares = append(newShares, &newShare)
	//	}
	//
	//	newSubdistributor := types.SubDistributor{
	//		Name: oldSubdistributor.Name,
	//		Destinations: types.Destinations{
	//			Shares:    newShares,
	//			BurnShare: oldSubdistributor.Destination.BurnShare.Percent,
	//			PrimaryShare: types.Account{
	//				Id:   oldSubdistributor.Destination.Account.Id,
	//				Type: oldSubdistributor.Destination.Account.Type,
	//			},
	//		},
	//	}
	//	newSubdistributors = append(newSubdistributors, newSubdistributor)
	//}
	DefaultSubDistributors := []types.SubDistributor{
		{
			Name: "default_distributor",
			Destinations: types.Destinations{
				PrimaryShare: types.Account{
					Id:   types.ValidatorsRewardsCollector,
					Type: types.MODULE_ACCOUNT,
				},
				BurnShare: sdk.ZeroDec(),
			},
			Sources: []*types.Account{
				{
					Id:   "",
					Type: types.MAIN,
				},
			},
		},
	}
	newSubdistributors = DefaultSubDistributors
	paramStore.Set(ctx, types.KeySubDistributors, newSubdistributors)

	return nil
}

//MigrateStore performs in-place store migrations from v1.0.0 to v1.0.1. The
//migration includes:
//
//- Vesting pools structure changed.
//- Vesting types changed to map
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	fmt.Println("CFEDISTRIBUTOR MIGRATION START")
	store := ctx.KVStore(storeKey)
	return migrateSubdistributorStates(store, cdc)
}
