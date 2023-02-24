package keeper

import (
	"encoding/binary"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetVestingAccountCount get the total number of vestingAccount
func (k Keeper) GetVestingAccountCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VestingAccountCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVestingAccountCount set the total number of vestingAccount
func (k Keeper) SetVestingAccountCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VestingAccountCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVestingAccount appends a vestingAccount in the store with a new id and update the count
func (k Keeper) AppendVestingAccount(
	ctx sdk.Context,
	vestingAccount types.VestingAccount,
) uint64 {
	// Create the vestingAccount
	count := k.GetVestingAccountCount(ctx)

	// Set the ID of the appended value
	vestingAccount.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountKey))
	appendedValue := k.cdc.MustMarshal(&vestingAccount)
	store.Set(GetVestingAccountIDBytes(vestingAccount.Id), appendedValue)

	// Update vestingAccount count
	k.SetVestingAccountCount(ctx, count+1)

	return count
}

// SetVestingAccount set a specific vestingAccount in the store
func (k Keeper) SetVestingAccount(ctx sdk.Context, vestingAccount types.VestingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountKey))
	b := k.cdc.MustMarshal(&vestingAccount)
	store.Set(GetVestingAccountIDBytes(vestingAccount.Id), b)
}

// GetVestingAccount returns a vestingAccount from its id
func (k Keeper) GetVestingAccount(ctx sdk.Context, id uint64) (val types.VestingAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountKey))
	b := store.Get(GetVestingAccountIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVestingAccount removes a vestingAccount from the store
func (k Keeper) RemoveVestingAccount(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountKey))
	store.Delete(GetVestingAccountIDBytes(id))
}

// GetAllVestingAccount returns all vestingAccount
func (k Keeper) GetAllVestingAccount(ctx sdk.Context) (list []types.VestingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()
	list = []types.VestingAccount{}
	for ; iterator.Valid(); iterator.Next() {
		var val types.VestingAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVestingAccountIDBytes returns the byte representation of the ID
func GetVestingAccountIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetVestingAccountIDFromBytes returns ID in uint64 format from a byte array
func GetVestingAccountIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) IsOnVestingAccountList(ctx sdk.Context, address string) bool {
	list := k.GetAllVestingAccount(ctx)
	for _, acc := range list {
		if address == acc.Address {
			return true
		}
	}
	return false
}
