package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetClaimRecordXX set a specific claimRecordXX in the store from its index
func (k Keeper) SetUserAirdropEntries(ctx sdk.Context, userAirdropEntries types.UserAirdropEntries) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserAirdropEntriesKeyPrefix))
	b := k.cdc.MustMarshal(&userAirdropEntries)
	store.Set(types.UserAirdropEntriesKey(
		userAirdropEntries.Address,
	), b)
}

// GetClaimRecordXX returns a claimRecordXX from its index
func (k Keeper) GetUserAirdropEntries(
	ctx sdk.Context,
	address string,
) (val types.UserAirdropEntries, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserAirdropEntriesKeyPrefix))

	b := store.Get(types.UserAirdropEntriesKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllClaimRecordXX returns all claimRecordXX
func (k Keeper) GetUsersAirdropEntries(ctx sdk.Context) (list []types.UserAirdropEntries) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserAirdropEntriesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserAirdropEntries
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
