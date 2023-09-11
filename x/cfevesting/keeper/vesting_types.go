package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GetAllVestingTypes returns all VestingTypes
func (k Keeper) GetAllVestingTypes(ctx sdk.Context) (vestingTypes types.VestingTypes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VestingTypesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	var list []*types.VestingType
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VestingType
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, &val)
	}
	vestingTypes.VestingTypes = list
	return
}

// get the vesting type by name
func (k Keeper) MustGetVestingType(ctx sdk.Context, name string) (vestingType *types.VestingType, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VestingTypesKeyPrefix)
	var vestType types.VestingType
	b := store.Get([]byte(name))
	if b == nil {
		return nil, fmt.Errorf("vesting type not found: " + name)
	}
	k.cdc.MustUnmarshal(b, &vestType)
	return &vestType, nil
}

// get the vesting type by name
func (k Keeper) MustGetVestingTypeForVestingPool(ctx sdk.Context, address string,
	vestingPoolName string) (*types.VestingType, error) {
	_, vestingPool, found := k.GetAccountVestingPool(ctx, address, vestingPoolName)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "vesting pool %s not found for address %s", vestingPoolName, address)
	}
	return k.MustGetVestingType(ctx, vestingPool.VestingType)
}

// set the vesting types
func (k Keeper) SetVestingTypes(ctx sdk.Context, vestingTypes types.VestingTypes) {
	for _, vt := range vestingTypes.VestingTypes {
		k.SetVestingType(ctx, *vt)
	}
}

// set the vesting type
func (k Keeper) SetVestingType(ctx sdk.Context, vestingType types.VestingType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VestingTypesKeyPrefix)
	av := k.cdc.MustMarshal(&vestingType)
	store.Set([]byte(vestingType.Name), av)
}

// get the vesting type by name
func (k Keeper) RemoveVestingType(ctx sdk.Context, name string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VestingTypesKeyPrefix)
	store.Delete([]byte(name))
}
