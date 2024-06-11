package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAllStates returns all States
func (k Keeper) GetAllStates(ctx sdk.Context) (list []types.State) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StateKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.State
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

// GetState return state by key
func (k Keeper) GetState(ctx sdk.Context, stateKey string) (remains types.State, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StateKeyPrefix)

	b := store.Get([]byte(stateKey))
	if b == nil {
		found = false
		return
	}
	found = true
	k.cdc.MustUnmarshal(b, &remains)
	return
}

func (k Keeper) GetBurnState(ctx sdk.Context) (remains types.State, found bool) {
	return k.GetState(ctx, types.BurnStateKey)
}

// SetState Set the state
func (k Keeper) SetState(ctx sdk.Context, state types.State) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StateKeyPrefix)
	av := k.cdc.MustMarshal(&state)
	store.Set([]byte(state.GetStateKey()), av)
}
