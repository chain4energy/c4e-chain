package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// get the vesting types
func (k Keeper) GetAccountVestingPools(ctx sdk.Context, accountAddress string) (accountVestingPools types.AccountVestingPools, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountVestingPoolsKeyPrefix)

	b := store.Get([]byte(accountAddress))
	if b == nil {
		found = false
		return
	}
	found = true
	k.cdc.MustUnmarshal(b, &accountVestingPools)
	return
}

// set the vesting types
func (k Keeper) SetAccountVestingPools(ctx sdk.Context, accountVestingPools types.AccountVestingPools) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountVestingPoolsKeyPrefix)
	av := k.cdc.MustMarshal(&accountVestingPools)
	store.Set([]byte(accountVestingPools.Address), av)
}

// get the vesting types
func (k Keeper) DeleteAccountVestingPools(ctx sdk.Context, accountAddress string) (accountVestingPools types.AccountVestingPools) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountVestingPoolsKeyPrefix)
	key := []byte(accountAddress)
	b := store.Get(key)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &accountVestingPools)
	store.Delete(key)
	return
}

// GetAllAccountVestingPools returns all AccountVestingPools
func (k Keeper) GetAllAccountVestingPools(ctx sdk.Context) (list []types.AccountVestingPools) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountVestingPoolsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccountVestingPools
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
