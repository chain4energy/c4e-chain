package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)



// get the minter
func (k Keeper) GetMinterState(ctx sdk.Context) (minter types.MinterState) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinterStateKey)
	if b == nil {
		panic("stored minter state should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &minter)
	return
}

// set the minter
func (k Keeper) SetMinterState(ctx sdk.Context, minter types.MinterState) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.MinterStateKey, b)
}


