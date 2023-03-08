package v2

import (
	"encoding/binary"
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v1"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func getOldMinterStateAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (oldMinterState v1.MinterState, err error) {
	b := store.Get(v1.MinterStateKey)
	if b == nil {
		return oldMinterState, fmt.Errorf("stored minter state should not have been nil")
	}

	err = cdc.Unmarshal(b, &oldMinterState)
	if err != nil {
		return oldMinterState, err
	}
	store.Delete(v1.MinterStateKey)
	return
}

func setNewMinterState(store sdk.KVStore, cdc codec.BinaryCodec, oldMinterState v1.MinterState) error {
	newMinterState := types.LegacyMinterState{
		SequenceId:                  uint32(oldMinterState.Position),
		AmountMinted:                oldMinterState.AmountMinted,
		LastMintBlockTime:           oldMinterState.LastMintBlockTime,
		RemainderFromPreviousPeriod: oldMinterState.RemainderFromPreviousPeriod,
		RemainderToMint:             oldMinterState.RemainderToMint,
	}
	err := newMinterState.Validate()
	if err != nil {
		return err
	}
	av, err := cdc.Marshal(&newMinterState)
	if err != nil {
		return err
	}
	store.Set(v1.MinterStateKey, av)
	return nil
}

func migrateMinterState(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldMinterState, err := getOldMinterStateAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewMinterState(store, cdc, oldMinterState)
}

func getOldMinterStateHistoryAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (oldMinterStateHistoryList []*v1.MinterState, err error) {
	prefixStore := prefix.NewStore(store, v1.MinterStateHistoryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v1.MinterState
		cdc.MustUnmarshal(iterator.Value(), &val)
		oldMinterStateHistoryList = append(oldMinterStateHistoryList, &val)
		prefixStore.Delete(iterator.Key())
	}

	return
}

func setNewMinterStateHistory(store sdk.KVStore, cdc codec.BinaryCodec, oldMinterStateHistory []*v1.MinterState) error {
	prefixStore := prefix.NewStore(store, types.MinterStateHistoryKeyPrefix)
	for _, oldMinterState := range oldMinterStateHistory {
		newMinterState := types.LegacyMinterState{
			SequenceId:                  uint32(oldMinterState.Position),
			AmountMinted:                oldMinterState.AmountMinted,
			LastMintBlockTime:           oldMinterState.LastMintBlockTime,
			RemainderFromPreviousPeriod: oldMinterState.RemainderFromPreviousPeriod,
			RemainderToMint:             oldMinterState.RemainderToMint,
		}
		err := newMinterState.Validate()
		if err != nil {
			return err
		}
		av, err := cdc.Marshal(&newMinterState)
		if err != nil {
			return err
		}
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, newMinterState.SequenceId)
		prefixStore.Set(bs, av)
	}

	return nil
}

func migrateMinterStateHistory(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldMinterState, err := getOldMinterStateHistoryAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewMinterStateHistory(store, cdc, oldMinterState)
}

// MigrateStore performs in-place store migrations from v1.0.1 to v1.1.0
// The migration includes:
// - MinterState change type of Position from int32 to uint32.
// - MinterState rename Position to SequenceId.
// - History of minter states in KvStore is now identified by SequenceId and its key is set in different way
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	if err := migrateMinterStateHistory(store, cdc); err != nil {
		return err
	}
	return migrateMinterState(store, cdc)
}
