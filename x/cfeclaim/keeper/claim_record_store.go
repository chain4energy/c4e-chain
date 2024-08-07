package keeper

import (
	"cosmossdk.io/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetUserEntry set a specific claimRecordXX in the store from its index
func (k Keeper) SetUserEntry(ctx sdk.Context, userEntry types.UserEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UserEntryKeyPrefix)
	b := k.cdc.MustMarshal(&userEntry)
	store.Set([]byte(
		userEntry.Address,
	), b)
}

// GetUserEntry returns a claimRecordXX from its index
func (k Keeper) GetUserEntry(
	ctx sdk.Context,
	address string,
) (val types.UserEntry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UserEntryKeyPrefix)

	b := store.Get([]byte(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllUsersEntries returns all UserEntries
func (k Keeper) GetAllUsersEntries(ctx sdk.Context) (list []types.UserEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UserEntryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserEntry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) MustGetUserEntry(ctx sdk.Context, userAddress string) (types.UserEntry, error) {
	userEntry, found := k.GetUserEntry(
		ctx,
		userAddress,
	)
	if !found {
		return types.UserEntry{}, errors.Wrapf(sdkerrors.ErrNotFound, "userEntry %s doesn't exist", userAddress)
	}

	return userEntry, nil
}
