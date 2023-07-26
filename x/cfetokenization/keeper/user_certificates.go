package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetUserCertificates set a specific userCertificates in the store
func (k Keeper) SetUserCertificates(ctx sdk.Context, userCertificates types.UserCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCertificatesKey))
	b := k.cdc.MustMarshal(&userCertificates)
	store.Set([]byte(userCertificates.Owner), b)
}

// GetUserCertificates returns a userCertificates from its id
func (k Keeper) GetUserCertificates(ctx sdk.Context, owner string) (val types.UserCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCertificatesKey))
	b := store.Get([]byte(owner))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserCertificates removes a userCertificates from the store
func (k Keeper) RemoveUserCertificates(ctx sdk.Context, owner string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCertificatesKey))
	store.Delete([]byte(owner))
}

// GetAllUserCertificates returns all userCertificates
func (k Keeper) GetAllUserCertificates(ctx sdk.Context) (list []types.UserCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCertificatesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
