package v2

import (
	"github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v1"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func getAllOldSubDistributorStatesAndDelete(store sdk.KVStore, cdc codec.BinaryCodec) (states []v1.State, err error) {
	prefixStore := prefix.NewStore(store, v1.RemainsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v1.State
		err := cdc.Unmarshal(iterator.Value(), &val)
		if err != nil {
			return nil, err
		}
		states = append(states, val)
		prefixStore.Delete(iterator.Key())
	}
	return
}

func setNewSubDistributorStates(store sdk.KVStore, cdc codec.BinaryCodec, oldStates []v1.State) error {
	prefixStore := prefix.NewStore(store, StateKeyPrefix)

	for _, oldState := range oldStates {
		var newAccount *Account
		if oldState.Burn {
			newAccount = nil
		} else {
			newAccount = &Account{
				Id:   oldState.Account.Id,
				Type: oldState.Account.Type,
			}
		}

		newState := State{
			Account: newAccount,
			Burn:    oldState.Burn,
			Remains: oldState.CoinsStates,
		}

		av, err := cdc.Marshal(&newState)
		if err != nil {
			return err
		}
		prefixStore.Set([]byte(newState.GetStateKey()), av)
	}
	return nil
}

func migrateSubDistributorStates(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldDistributorStates, err := getAllOldSubDistributorStatesAndDelete(store, cdc)
	if err != nil {
		return err
	}
	return setNewSubDistributorStates(store, cdc, oldDistributorStates)
}

// MigrateStore performs in-place store migrations from v1.0.1 to v1.1.0.
// The migration includes:
// - SubDistributor State rename CoinStates to Remains.
// - If burn is set to true state account must be nil
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	return migrateSubDistributorStates(store, cdc)
}
