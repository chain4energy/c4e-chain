package keeper

import (
	"github.com/chain4energy/c4e-chain/x/energybank/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTokenParams set a specific tokenParams in the store from its index
func (k Keeper) SetTokenParams(ctx sdk.Context, tokenParams types.TokenParams) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenParamsKeyPrefix))
	b := k.cdc.MustMarshal(&tokenParams)
	store.Set(types.TokenParamsKey(
		tokenParams.Index,
	), b)
}

// GetTokenParams returns a tokenParams from its index
func (k Keeper) GetTokenParams(
	ctx sdk.Context,
	index string,

) (val types.TokenParams, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenParamsKeyPrefix))

	b := store.Get(types.TokenParamsKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTokenParams removes a tokenParams from the store
func (k Keeper) RemoveTokenParams(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenParamsKeyPrefix))
	store.Delete(types.TokenParamsKey(
		index,
	))
}

// GetAllTokenParams returns all tokenParams
func (k Keeper) GetAllTokenParams(ctx sdk.Context) (list []types.TokenParams) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenParamsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokenParams
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
