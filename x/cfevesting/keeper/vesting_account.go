package keeper

import (
	"encoding/binary"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetVestingAccountTraceCount get the total number of vestingAccount
func (k Keeper) GetVestingAccountTraceCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VestingAccountTraceCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVestingAccountTraceCount set the total number of vestingAccount
func (k Keeper) SetVestingAccountTraceCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VestingAccountTraceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVestingAccountTrace appends a vestingAccount in the store with a new id and update the count
func (k Keeper) AppendVestingAccountTrace(
	ctx sdk.Context,
	vestingAccountTrace types.VestingAccountTrace,
) uint64 {
	// Create the vestingAccountTrace
	count := k.GetVestingAccountTraceCount(ctx)

	// Set the ID of the appended value
	vestingAccountTrace.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountTraceKey))
	appendedValue := k.cdc.MustMarshal(&vestingAccountTrace)
	store.Set([]byte(vestingAccountTrace.Address), appendedValue)

	// Update vestingAccount count
	k.SetVestingAccountTraceCount(ctx, count+1)

	return count
}

// SetVestingAccountTrace set a specific vestingAccount in the store
func (k Keeper) SetVestingAccountTrace(ctx sdk.Context, vestingAccountTrace types.VestingAccountTrace) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountTraceKey))
	b := k.cdc.MustMarshal(&vestingAccountTrace)
	store.Set([]byte(vestingAccountTrace.Address), b)
}

// GetVestingAccountTraceById returns a vestingAccount from its id
func (k Keeper) GetVestingAccountTraceById(ctx sdk.Context, id uint64) (val types.VestingAccountTrace, found bool) {
	list := k.GetAllVestingAccountTrace(ctx)
	for _, acc := range list {
		if id == acc.Id {
			return acc, true
		}
	}
	return val, false
}

// RemoveVestingAccountTrace removes a vestingAccount from the store
func (k Keeper) RemoveVestingAccountTrace(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountTraceKey))
	store.Delete([]byte(address))
}

// GetAllVestingAccountTrace returns all vestingAccount
func (k Keeper) GetAllVestingAccountTrace(ctx sdk.Context) (list []types.VestingAccountTrace) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountTraceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()
	list = []types.VestingAccountTrace{}
	for ; iterator.Valid(); iterator.Next() {
		var val types.VestingAccountTrace
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVestingAccountTraceIDBytes returns the byte representation of the ID
// func GetVestingAccountTraceAdressBytes(address string) []byte {
// 	bz := make([]byte, 8)
// 	binary.BigEndian.PutUint64(bz, id)
// 	return bz
// }

// GetVestingAccountTraceIDFromBytes returns ID in uint64 format from a byte array
// func GetVestingAccountTraceAddressFromBytes(bz []byte) string {
// 	return binary.BigEndian.Uint64(bz)
// }

// GetVestingAccountById returns a vestingAccount from its id
func (k Keeper) GetVestingAccountTrace(ctx sdk.Context, address string) (val types.VestingAccountTrace, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountTraceKey))
	b := store.Get([]byte(address))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
