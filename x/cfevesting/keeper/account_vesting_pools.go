package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

func (k Keeper) GetAccountVestingPool(ctx sdk.Context, accountAddress string, name string) (vestingPool types.VestingPool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountVestingPoolsKeyPrefix)
	var accountVestingPools types.AccountVestingPools
	b := store.Get([]byte(accountAddress))
	if b == nil {
		found = false
		return
	}

	k.cdc.MustUnmarshal(b, &accountVestingPools)
	for _, vestPool := range accountVestingPools.VestingPools {
		if vestPool.Name == name {
			return vestingPool, true
		}
	}

	return
}

func (k Keeper) SetAccountVestingPools(ctx sdk.Context, accountVestingPools types.AccountVestingPools) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountVestingPoolsKeyPrefix)
	av := k.cdc.MustMarshal(&accountVestingPools)
	store.Set([]byte(accountVestingPools.Owner), av)
}

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

func (k Keeper) GetAllAccountVestingPools(ctx sdk.Context) (list types.AccountVestingPoolsList) {
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
