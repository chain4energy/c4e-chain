package keeper

import (
	"strconv"

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
func (k Keeper) GetMinterStateHistory(ctx sdk.Context, SequenceId int32) (state types.MinterState, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MinterStateHistoryKeyPrefix)

	b := store.Get([]byte(strconv.FormatInt(int64(SequenceId), 10)))
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
	store.Set([]byte(strconv.FormatInt(int64(state.Position), 10)), av)
}

// GetAllMinterStateHistory returns all historical minter states for ended Minters
func (k Keeper) GetAllMinterStateHistory(ctx sdk.Context) (list []types.MinterState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MinterStateHistoryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MinterState
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ConvertMinterStateHistory(ctx sdk.Context) (list []*types.MinterState) {
	history := make([]*types.MinterState, 0)
	stateHistory := k.GetAllMinterStateHistory(ctx)
	if len(stateHistory) > 0 {

		for i := 0; i < len(stateHistory); i++ {
			history = append(history, &stateHistory[i])

		}
	}
	return history
}
