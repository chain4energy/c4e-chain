package v3

import (
	"encoding/binary"
	"github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v2"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// getAllOldAccountVestingPoolsAndDelete returns all old version AccountVestingPools and deletes them from the KVStore
func getAllOldAccountVestingPoolsAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (oldAccPools []v2.AccountVestingPools, err error) {
	prefixStore := prefix.NewStore(store, v2.AccountVestingPoolsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v2.AccountVestingPools
		err := cdc.Unmarshal(iterator.Value(), &val)
		if err != nil {
			return nil, err
		}
		oldAccPools = append(oldAccPools, val)
		prefixStore.Delete(iterator.Key())
	}
	return
}

func setNewAccountVestingPools(store sdk.KVStore, cdc codec.BinaryCodec, oldAccPools []v2.AccountVestingPools) error {
	prefixStore := prefix.NewStore(store, types.AccountVestingPoolsKeyPrefix)
	for _, oldAccPool := range oldAccPools {
		oldPools := oldAccPool.VestingPools
		var newPools []*VestingPool
		for _, oldPool := range oldPools {
			newPool := VestingPool{
				Name:            oldPool.Name,
				VestingType:     oldPool.VestingType,
				LockStart:       oldPool.LockStart,
				LockEnd:         oldPool.LockEnd,
				InitiallyLocked: oldPool.InitiallyLocked,
				Withdrawn:       oldPool.Withdrawn,
				Sent:            oldPool.Sent,
			}
			newPools = append(newPools, &newPool)
		}

		newAccPool := AccountVestingPools{
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

func getOldVestingAccountTracesCountAndDelete(store sdk.KVStore) uint64 {
	prefixStore := prefix.NewStore(store, []byte{})
	byteKey := types.KeyPrefix(v2.VestingAccountCountKey)
	bz := prefixStore.Get(byteKey)
	prefixStore.Delete(byteKey)
	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}
	// Parse bytes
	return binary.BigEndian.Uint64(bz)

}

// getAllOldAccountVestingTracesAndDelete returns all old version AccountVestingTrace and deletes them from the KVStore
func getAllOldVestingAccountTracesAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (oldVestingAccounntTraces []v2.VestingAccount, err error) {
	prefixStore := prefix.NewStore(store, types.KeyPrefix(v2.VestingAccountKey))
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})

	defer iterator.Close()
	oldVestingAccounntTraces = []v2.VestingAccount{}
	for ; iterator.Valid(); iterator.Next() {
		var val v2.VestingAccount
		if err := cdc.Unmarshal(iterator.Value(), &val); err != nil {
			return nil, err
		}
		oldVestingAccounntTraces = append(oldVestingAccounntTraces, val)
	}

	return
}

func setNewVestingAccountTracesCount(store sdk.KVStore, count uint64) {
	prefixStore := prefix.NewStore(store, []byte{})
	byteKey := types.KeyPrefix(types.VestingAccountTraceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	prefixStore.Set(byteKey, bz)
}

func setNewVestingAccountAccountTraces(store sdk.KVStore, cdc codec.BinaryCodec, oldVestingAccounntTraces []v2.VestingAccount) error {
	prefixStore := prefix.NewStore(store, types.KeyPrefix(types.VestingAccountTraceKey))
	for _, oldVestingAccounntTrace := range oldVestingAccounntTraces {
		vestingAccountTrace := VestingAccountTrace{
			Id:                 oldVestingAccounntTrace.Id,
			Address:            oldVestingAccounntTrace.Address,
			Genesis:            false,
			FromGenesisPool:    false,
			FromGenesisAccount: false,
		}
		b, err := cdc.Marshal(&vestingAccountTrace)
		if err != nil {
			return err
		}
		prefixStore.Set([]byte(vestingAccountTrace.Address), b)
	}
	return nil
}

func migrateVestingAccountTrace(store sdk.KVStore, cdc codec.BinaryCodec) error {
	count := getOldVestingAccountTracesCountAndDelete(store)
	setNewVestingAccountTracesCount(store, count)

	oldAccountVestingTraces, err := getAllOldVestingAccountTracesAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewVestingAccountAccountTraces(store, cdc, oldAccountVestingTraces)
}

// MigrateStore performs in-place store migrations from v1.1.0 to v1.2.0
// The migration includes:
// - Vesting pools structure changed.
// - Vesting Acounts Traces changed to map
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	if err := migrateVestingAccountTrace(store, cdc); err != nil {
		return err
	}
	return migrateVestingPools(store, cdc)
}
