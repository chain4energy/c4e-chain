package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetUserDevices set a specific userDevices in the store
func (k Keeper) SetUserDevices(ctx sdk.Context, userDevices types.UserDevices) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserDevicesKey))
	b := k.cdc.MustMarshal(&userDevices)
	store.Set([]byte(userDevices.Owner), b)
}

// GetUserDevices returns a userDevices from its id
func (k Keeper) GetUserDevices(ctx sdk.Context, owner string) (val types.UserDevices, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserDevicesKey))
	b := store.Get([]byte(owner))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserDevices removes a userDevices from the store
func (k Keeper) RemoveUserDevices(ctx sdk.Context, owner string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserDevicesKey))
	store.Delete([]byte(owner))
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

// SetUserDevices set a specific userDevices in the store
func (k Keeper) SetPendingDevice(ctx sdk.Context, pendingDevice types.PendingDevice) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingDeviceKey))
	b := k.cdc.MustMarshal(&pendingDevice)
	store.Set([]byte(pendingDevice.DeviceAddress), b)
}

// SetUserDevices set a specific userDevices in the store
func (k Keeper) GetPendingDevice(ctx sdk.Context, deviceAddress string) (val types.PendingDevice, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingDeviceKey))
	b := store.Get([]byte(deviceAddress))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetUserDevices set a specific userDevices in the store
func (k Keeper) SetDevice(ctx sdk.Context, device types.Device) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKey))
	b := k.cdc.MustMarshal(&device)
	store.Set([]byte(device.DeviceAddress), b)
}

// SetUserDevices set a specific userDevices in the store
func (k Keeper) GetDevice(ctx sdk.Context, deviceAddress string) (val types.Device, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKey))
	b := store.Get([]byte(deviceAddress))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetUserDevices set a specific userDevices in the store
func (k Keeper) GetAllDevices(ctx sdk.Context) (list []types.Device) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Device
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
