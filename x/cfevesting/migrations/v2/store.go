package v2

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v1"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// getAllOldAccountVestingPoolsAndDelete returns all old version AccountVestingPools and deletes them from the KVStore
func getAllOldAccountVestingPoolsAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (oldAccPools []v1.AccountVestingPools, err error) {
	prefixStore := prefix.NewStore(store, v1.AccountVestingPoolsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v1.AccountVestingPools
		err := cdc.Unmarshal(iterator.Value(), &val)
		if err != nil {
			return nil, err
		}
		oldAccPools = append(oldAccPools, val)
		prefixStore.Delete(iterator.Key())
	}
	return
}

func setNewAccountVestingPools(store sdk.KVStore, cdc codec.BinaryCodec, oldAccPools []v1.AccountVestingPools) error {
	prefixStore := prefix.NewStore(store, AccountVestingPoolsKeyPrefix)
	for _, oldAccPool := range oldAccPools {
		oldPools := oldAccPool.VestingPools
		var newPools []*VestingPool
		for _, oldPool := range oldPools {
			sent := oldPool.LastModificationWithdrawn.Add(oldPool.Vested).Sub(oldPool.Withdrawn).Sub(oldPool.LastModificationVested)

			newPool := VestingPool{
				Name:            oldPool.Name,
				VestingType:     oldPool.VestingType,
				LockStart:       oldPool.LockStart,
				LockEnd:         oldPool.LockEnd,
				InitiallyLocked: oldPool.Vested,
				Withdrawn:       oldPool.Withdrawn,
				Sent:            sent,
			}
			newPools = append(newPools, &newPool)
		}

		newAccPool := AccountVestingPools{
			Address:      oldAccPool.Address,
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

func getOldVestingTypesAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (vestingTypes v1.VestingTypes, err error) {
	b := store.Get(v1.VestingTypesKey)
	if b == nil {
		return vestingTypes, nil
	}

	err = cdc.Unmarshal(b, &vestingTypes)
	if err != nil {
		return vestingTypes, err
	}
	store.Delete(v1.VestingTypesKey)
	return
}

func setNewVestingTypes(store sdk.KVStore, cdc codec.BinaryCodec, vestingTypes v1.VestingTypes) error {
	for _, vt := range vestingTypes.VestingTypes {
		newVestingType := VestingType{
			Name:          vt.Name,
			VestingPeriod: vt.VestingPeriod,
			LockupPeriod:  vt.LockupPeriod,
			Free:          sdk.ZeroDec(),
		}
		if vt.Name == "Validators" {
			newVestingType.Free = sdk.MustNewDecFromStr("0.05")
		}
		err := setNewVestingType(store, cdc, newVestingType)
		if err != nil {
			return err
		}
	}
	return nil
}

// set the vesting type
func setNewVestingType(store sdk.KVStore, cdc codec.BinaryCodec, newVestingType VestingType) error {
	pStore := prefix.NewStore(store, VestingTypesKey)
	av, err := cdc.Marshal(&newVestingType)
	if err != nil {
		return err
	}
	pStore.Set([]byte(newVestingType.Name), av)
	return nil
}

func migrateVestingTypes(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldAccountVestingPools, err := getOldVestingTypesAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewVestingTypes(store, cdc, oldAccountVestingPools)
}

func migrateVestingPools(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldAccountVestingPools, err := getAllOldAccountVestingPoolsAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewAccountVestingPools(store, cdc, oldAccountVestingPools)
}

// MigrateStore performs in-place store migrations from v1.0.1 to v1.1.0
// The migration includes:
// - Vesting pools structure changed.
// - Vesting types changed to map
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	err := migrateVestingTypes(store, cdc)
	if err != nil {
		return err
	}
	return migrateVestingPools(store, cdc)
}
