package v120

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v110"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// getAllOldAccountVestingPoolsAndDelete returns all old version AccountVestingPools and deletes them from the KVStore
func getAllOldAccountVestingPoolsAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (oldAccPools []v110.AccountVestingPools, err error) {
	prefixStore := prefix.NewStore(store, v110.AccountVestingPoolsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v110.AccountVestingPools
		err := cdc.Unmarshal(iterator.Value(), &val)
		if err != nil {
			return nil, err
		}
		oldAccPools = append(oldAccPools, val)
		prefixStore.Delete(iterator.Key())
	}
	return
}

func setNewAccountVestingPools(store sdk.KVStore, cdc codec.BinaryCodec, oldAccPools []v110.AccountVestingPools) error {
	prefixStore := prefix.NewStore(store, types.AccountVestingPoolsKeyPrefix)
	for _, oldAccPool := range oldAccPools {
		oldPools := oldAccPool.VestingPools
		var newPools []*types.VestingPool
		for _, oldPool := range oldPools {
			newPool := types.VestingPool(*oldPool)
			newPools = append(newPools, &newPool)
		}

		newAccPool := types.AccountVestingPools{
			Owner:        oldAccPool.Address,
			VestingPools: newPools,
		}
		av, err := cdc.Marshal(&newAccPool)
		if err != nil {
			return err
		}
		prefixStore.Set([]byte(newAccPool.Owner), av)
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

// MigrateStore performs in-place store migrations from v1.1.0 to v1.2.0
// The migration includes:
// - Vesting pools structure changed.
// - Vesting types changed to map
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	return migrateVestingPools(store, cdc)
}
