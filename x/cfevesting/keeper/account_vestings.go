package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// get the vesting types
func (k Keeper) GetAccountVestings(ctx sdk.Context, accountAddress string) (accountVestings types.AccountVestings, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountVestingsKeyPrefix)

	b := store.Get([]byte(accountAddress))
	if b == nil {
		found = false
		return
	}
	found = true
	k.cdc.MustUnmarshal(b, &accountVestings)
	return
}

// set the vesting types
func (k Keeper) SetAccountVestings(ctx sdk.Context, accountVestings types.AccountVestings) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountVestingsKeyPrefix)
	av := k.cdc.MustMarshal(&accountVestings)
	store.Set([]byte(accountVestings.Address), av)
}

// get the vesting types
func (k Keeper) DeleteAccountVestings(ctx sdk.Context, accountAddress string) (accountVestings types.AccountVestings) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountVestingsKeyPrefix)
	key := []byte(accountAddress)
	b := store.Get(key)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &accountVestings)
	store.Delete(key)
	return
}

// GetAllAccountVestings returns all AccountVestings
func (k Keeper) GetAllAccountVestings(ctx sdk.Context) (list []types.AccountVestings) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountVestingsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccountVestings
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

