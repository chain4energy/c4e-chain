package keeper

import (
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"

)

// GetAllVestingTypes returns all VestingTypes
func (k Keeper) GetAllVestingTypes(ctx sdk.Context) (vestingTypes types.VestingTypes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VestingTypesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	list := []*types.VestingType{}
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
func (k Keeper) GetVestingType(ctx sdk.Context, name string) (vestingType types.VestingType, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VestingTypesKeyPrefix)

	b := store.Get([]byte(name))
	if b == nil {
		err = fmt.Errorf("vesting type not found: " + name)
		return
	}
	k.cdc.MustUnmarshal(b, &vestingType)
	return
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
