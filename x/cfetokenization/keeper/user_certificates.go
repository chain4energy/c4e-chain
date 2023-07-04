package keeper

import (
	"encoding/binary"

	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetUserCertificatesCount get the total number of userCertificates
func (k Keeper) GetUserCertificatesCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.UserCertificatesCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetUserCertificatesCount set the total number of userCertificates
func (k Keeper) SetUserCertificatesCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.UserCertificatesCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendUserCertificates appends a userCertificates in the store with a new id and update the count
func (k Keeper) AppendUserCertificates(
	ctx sdk.Context,
	userCertificates types.UserCertificates,
) uint64 {
	// Create the userCertificates
	count := k.GetUserCertificatesCount(ctx)

	// Set the ID of the appended value
	userCertificates.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCertificatesKey))
	appendedValue := k.cdc.MustMarshal(&userCertificates)
	store.Set(GetUserCertificatesIDBytes(userCertificates.Id), appendedValue)

	// Update userCertificates count
	k.SetUserCertificatesCount(ctx, count+1)

	return count
}

// SetUserCertificates set a specific userCertificates in the store
func (k Keeper) SetUserCertificates(ctx sdk.Context, userCertificates types.UserCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCertificatesKey))
	b := k.cdc.MustMarshal(&userCertificates)
	store.Set(GetUserCertificatesIDBytes(userCertificates.Id), b)
}

// GetUserCertificates returns a userCertificates from its id
func (k Keeper) GetUserCertificates(ctx sdk.Context, id uint64) (val types.UserCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCertificatesKey))
	b := store.Get(GetUserCertificatesIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserCertificates removes a userCertificates from the store
func (k Keeper) RemoveUserCertificates(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCertificatesKey))
	store.Delete(GetUserCertificatesIDBytes(id))
}

// GetAllUserCertificates returns all userCertificates
func (k Keeper) GetAllUserCertificates(ctx sdk.Context) (list []types.UserCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCertificatesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetUserCertificatesIDBytes returns the byte representation of the ID
func GetUserCertificatesIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetUserCertificatesIDFromBytes returns ID in uint64 format from a byte array
func GetUserCertificatesIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
