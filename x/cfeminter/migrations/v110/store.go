package v110

import (
	"encoding/binary"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v101"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func getV101MinterStateAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (minterState v101.MinterState, err error) {
	b := store.Get(v101.MinterStateKey)
	if b == nil {
		return minterState, nil
	}

	err = cdc.Unmarshal(b, &minterState)
	if err != nil {
		return minterState, err
	}
	store.Delete(v101.MinterStateKey)
	return
}

func setNewMinterState(store sdk.KVStore, cdc codec.BinaryCodec, oldState v101.MinterState) error {
	newState := types.MinterState{
		SequenceId:                  uint32(oldState.Position),
		AmountMinted:                oldState.AmountMinted,
		LastMintBlockTime:           oldState.LastMintBlockTime,
		RemainderFromPreviousPeriod: oldState.RemainderFromPreviousPeriod,
		RemainderToMint:             oldState.RemainderToMint,
	}
	err := newState.Validate()
	if err != nil {
		return err
	}
	av, err := cdc.Marshal(&newState)
	if err != nil {
		return err
	}
	store.Set(v101.MinterStateKey, av)
	return nil
}

func migrateMinterState(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldMinterState, err := getV101MinterStateAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewMinterState(store, cdc, oldMinterState)
}

func getV101MinterStateHistoryAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (minterStateList []v101.MinterState, err error) {
	prefixStore := prefix.NewStore(store, v101.MinterStateHistoryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v101.MinterState
		cdc.MustUnmarshal(iterator.Value(), &val)
		minterStateList = append(minterStateList, val)
		prefixStore.Delete(iterator.Key())
	}

	return
}

func setNewMinterStateHistory(store sdk.KVStore, cdc codec.BinaryCodec, minterStateList []v101.MinterState) error {
	prefixStore := prefix.NewStore(store, types.MinterStateHistoryKeyPrefix)
	for _, V101MinterState := range minterStateList {
		newState := types.MinterState{
			SequenceId:                  uint32(V101MinterState.Position),
			AmountMinted:                V101MinterState.AmountMinted,
			LastMintBlockTime:           V101MinterState.LastMintBlockTime,
			RemainderFromPreviousPeriod: V101MinterState.RemainderFromPreviousPeriod,
			RemainderToMint:             V101MinterState.RemainderToMint,
		}
		err := newState.Validate()
		if err != nil {
			return err
		}
		av, err := cdc.Marshal(&newState)
		if err != nil {
			return err
		}
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, newState.SequenceId)
		prefixStore.Set(bs, av)
	}

	return nil
}

func migrateMinterStateHistory(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldMinterState, err := getV101MinterStateHistoryAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewMinterStateHistory(store, cdc, oldMinterState)
}

// MigrateStore performs in-place store migrations from v1.0.1 to v1.1.0
// The migration includes:
// - SubDistributor State rename CoinStates to Remains.
// - If burn is set to true state account must be nil
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	if err := migrateMinterStateHistory(store, cdc); err != nil {
		return err
	}
	return migrateMinterState(store, cdc)
}
