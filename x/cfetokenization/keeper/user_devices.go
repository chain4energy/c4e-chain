package keeper

import (
	"encoding/binary"

	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetUserDevicesCount get the total number of userDevices
func (k Keeper) GetUserDevicesCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.UserDevicesCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetUserDevicesCount set the total number of userDevices
func (k Keeper) SetUserDevicesCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.UserDevicesCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendUserDevices appends a userDevices in the store with a new id and update the count
func (k Keeper) AppendUserDevices(
	ctx sdk.Context,
	userDevices types.UserDevices,
) uint64 {
	// Create the userDevices
	count := k.GetUserDevicesCount(ctx)

	// Set the ID of the appended value
	userDevices.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserDevicesKey))
	appendedValue := k.cdc.MustMarshal(&userDevices)
	store.Set(GetUserDevicesIDBytes(userDevices.Id), appendedValue)

	// Update userDevices count
	k.SetUserDevicesCount(ctx, count+1)

	return count
}

// SetUserDevices set a specific userDevices in the store
func (k Keeper) SetUserDevices(ctx sdk.Context, userDevices types.UserDevices) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserDevicesKey))
	b := k.cdc.MustMarshal(&userDevices)
	store.Set(GetUserDevicesIDBytes(userDevices.Id), b)
}

// GetUserDevices returns a userDevices from its id
func (k Keeper) GetUserDevices(ctx sdk.Context, id uint64) (val types.UserDevices, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserDevicesKey))
	b := store.Get(GetUserDevicesIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserDevices removes a userDevices from the store
func (k Keeper) RemoveUserDevices(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserDevicesKey))
	store.Delete(GetUserDevicesIDBytes(id))
}

// GetAllUserDevices returns all userDevices
func (k Keeper) GetAllUserDevices(ctx sdk.Context) (list []types.UserDevices) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserDevicesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserDevices
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetUserDevicesIDBytes returns the byte representation of the ID
func GetUserDevicesIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetUserDevicesIDFromBytes returns ID in uint64 format from a byte array
func GetUserDevicesIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
