package keeper

import (
	"encoding/binary"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetTokensHistoryCount get the total number of tokensHistory
func (k Keeper) GetTokensHistoryCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TokensHistoryCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTokensHistoryCount set the total number of tokensHistory
func (k Keeper) SetTokensHistoryCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TokensHistoryCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTokensHistory appends a tokensHistory in the store with a new id and update the count
func (k Keeper) AppendTokensHistory(
	ctx sdk.Context,
	tokensHistory types.TokensHistory,
) uint64 {
	// Create the tokensHistory
	count := k.GetTokensHistoryCount(ctx)

	// Set the ID of the appended value
	tokensHistory.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensHistoryKey))
	appendedValue := k.cdc.MustMarshal(&tokensHistory)
	store.Set(GetTokensHistoryIDBytes(tokensHistory.Id), appendedValue)

	// Update tokensHistory count
	k.SetTokensHistoryCount(ctx, count+1)

	return count
}

// SetTokensHistory set a specific tokensHistory in the store
func (k Keeper) SetTokensHistory(ctx sdk.Context, tokensHistory types.TokensHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensHistoryKey))
	b := k.cdc.MustMarshal(&tokensHistory)
	store.Set(GetTokensHistoryIDBytes(tokensHistory.Id), b)
}

// GetTokensHistory returns a tokensHistory from its id
func (k Keeper) GetTokensHistory(ctx sdk.Context, id uint64) (val types.TokensHistory, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensHistoryKey))
	b := store.Get(GetTokensHistoryIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTokensHistory removes a tokensHistory from the store
func (k Keeper) RemoveTokensHistory(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensHistoryKey))
	store.Delete(GetTokensHistoryIDBytes(id))
}

// GetAllTokensHistory returns all tokensHistory
func (k Keeper) GetAllTokensHistory(ctx sdk.Context) (list []types.TokensHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensHistoryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokensHistory
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTokenHistoryUserAddress returns all tokenHistoryForUser
func (k Keeper) GetTokenHistoryUserAddress(ctx sdk.Context, userBlockchainAddress string) (list []types.TokensHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensHistoryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokensHistory
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.UserAddress == userBlockchainAddress {
			list = append(list, val)
		}
	}
	return
}

// GetTokensHistoryIDBytes returns the byte representation of the ID
func GetTokensHistoryIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTokensHistoryIDFromBytes returns ID in uint64 format from a byte array
func GetTokensHistoryIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
