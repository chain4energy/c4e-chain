package keeper

import (
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// get the vesting types
func (k Keeper) GetVestingTypes(ctx sdk.Context) (vestingTypes types.VestingTypes) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.VestingTypesKey)
	if b == nil {
		return vestingTypes
	}

	k.cdc.MustUnmarshal(b, &vestingTypes)
	return
}

// get the vesting type by name
func (k Keeper) GetVestingType(ctx sdk.Context, name string) (vestingType types.VestingType, err error) {
	vestingTypes := k.GetVestingTypes(ctx)
	// TODO optimize storing vesting types
	for _, vestingTypeLocal := range vestingTypes.GetVestingTypes() {
		if vestingTypeLocal.GetName() == name {
			return *vestingTypeLocal, nil
		}
	}
	return vestingType, fmt.Errorf("vesting type not found: " + name)
}

// set the vesting types
func (k Keeper) SetVestingTypes(ctx sdk.Context, vestingTypes types.VestingTypes) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&vestingTypes)
	store.Set(types.VestingTypesKey, b)
}
