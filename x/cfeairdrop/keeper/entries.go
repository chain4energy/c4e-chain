package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetClaimRecordXX set a specific claimRecordXX in the store from its index
func (k Keeper) SetUserEntry(ctx sdk.Context, userEntry types.UserEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserEntryKeyPrefix))
	b := k.cdc.MustMarshal(&userEntry)
	store.Set(types.UserEntryKey(
		userEntry.Address,
	), b)
}

// GetClaimRecordXX returns a claimRecordXX from its index
func (k Keeper) GetUserEntry(
	ctx sdk.Context,
	address string,
) (val types.UserEntry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserEntryKeyPrefix))

	b := store.Get(types.UserEntryKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllClaimRecordXX returns all claimRecordXX
func (k Keeper) GetUsersEntries(ctx sdk.Context) (list []types.UserEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserEntryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserEntry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
