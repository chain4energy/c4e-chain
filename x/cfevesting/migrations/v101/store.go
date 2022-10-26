package v101

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v100cfevesting "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v100"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
)

// getAllOldAccountVestingPoolsAndDelete returns all old version AccountVestingPools and deletes them from the KVStore
func getAllOldAccountVestingPoolsAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (list []v100cfevesting.AccountVestingPools, err error) {
	prefixStore := prefix.NewStore(store, types.AccountVestingPoolsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v100cfevesting.AccountVestingPools
		err := cdc.Unmarshal(iterator.Value(), &val)
		if err != nil {
			return nil, err
		}
		list = append(list, val)
		prefixStore.Delete(iterator.Key())
	}
	return
}

func setNewAccountVestingPools(store sdk.KVStore, cdc codec.BinaryCodec, oldAccPools []v100cfevesting.AccountVestingPools) error {
	prefixStore := prefix.NewStore(store, types.AccountVestingPoolsKeyPrefix)
	for _, oldAccPool := range oldAccPools {
		oldPools := oldAccPool.VestingPools
		newPools := []*types.VestingPool{}
		for _, oldPool := range oldPools {
			newPool := types.VestingPool{
				Id: oldPool.Id,
				Name: oldPool.Name + "New",
				VestingType: oldPool.VestingType,
				LockStart: oldPool.LockStart,
				LockEnd: oldPool.LockEnd,
				Vested: oldPool.Vested,
				Withdrawn: oldPool.Withdrawn,
				Sent: oldPool.Sent,
				LastModification: oldPool.LastModification,
				LastModificationVested: oldPool.LastModificationVested,
				LastModificationWithdrawn: oldPool.LastModificationWithdrawn,
			}
			newPools = append(newPools, &newPool)
		}

		newAccPool := types.AccountVestingPools{
			Address: oldAccPool.Address,
			VestingPools: newPools,

		}
		av, err := cdc.Marshal(&newAccPool)
		if err != nil {
			return err
		}
		prefixStore.Set([]byte(newAccPool.Address), av)
	}
	return nil
}

func migrateVestingPools(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldAccountVestingPools, err := getAllOldAccountVestingPoolsAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewAccountVestingPools(store, cdc, oldAccountVestingPools)
}

// MigrateStore performs in-place store migrations from v1.0.0 to v1.0.1. The
// migration includes:
//
// - Vesting pools structure changed.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	return migrateVestingPools(store, cdc)
}

