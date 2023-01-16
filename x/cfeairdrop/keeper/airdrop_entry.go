package keeper

import (
	"encoding/binary"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAirdropEntryCount get the total number of airdropEntry
func (k Keeper) GetAirdropEntryCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AirdropEntryCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAirdropEntryCount set the total number of airdropEntry
func (k Keeper) SetAirdropEntryCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AirdropEntryCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetAirdropEntry returns a airdropEntry from its id
func (k Keeper) GetAirdropEntry(ctx sdk.Context, id uint64) (val types.AirdropEntry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropEntryKey))
	b := store.Get(GetAirdropEntryIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAirdropEntry removes a airdropEntry from the store
func (k Keeper) RemoveAirdropEntry(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropEntryKey))
	store.Delete(GetAirdropEntryIDBytes(id))
}

// GetAllAirdropEntry returns all airdropEntry
func (k Keeper) GetAllAirdropEntry(ctx sdk.Context) (list []types.AirdropEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropEntryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AirdropEntry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAirdropEntryIDBytes returns the byte representation of the ID
func GetAirdropEntryIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAirdropEntryIDFromBytes returns ID in uint64 format from a byte array
func GetAirdropEntryIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
