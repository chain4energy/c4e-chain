package keeper

import (
	"encoding/binary"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
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

// get the vesting types
func (k Keeper) GetMinterStateHistory(ctx sdk.Context, sequenceId uint32) (state types.MinterState, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MinterStateHistoryKeyPrefix)
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, sequenceId)
	b := store.Get(bs)
	if b == nil {
		found = false
		return
	}
	found = true
	k.cdc.MustUnmarshal(b, &state)
	return
}

// set the vesting types
func (k Keeper) SetMinterStateHistory(ctx sdk.Context, state types.MinterState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MinterStateHistoryKeyPrefix)
	av := k.cdc.MustMarshal(&state)
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, state.SequenceId)
	store.Set(bs, av)
}

// GetAllMinterStateHistory returns all historical minter states for ended Minter
func (k Keeper) GetAllMinterStateHistory(ctx sdk.Context) (list []*types.MinterState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MinterStateHistoryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MinterState
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, &val)
	}

	return
}
