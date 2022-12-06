package v110

import (
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v101"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
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

func setNewMinterStates(store sdk.KVStore, cdc codec.BinaryCodec, oldState v101.MinterState) error {
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
	return setNewMinterStates(store, cdc, oldMinterState)
}

// MigrateStore performs in-place store migrations from v1.0.0 to v1.0.1. The
// migration includes:
//
// - SubDistributor State rename CoinStates to Remains.
// - If burn is set to true state account must be nil
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	return migrateMinterState(store, cdc)
}
