package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetVestingPoolReservation(ctx sdk.Context, vestingPoolName string) (vestingPoolReservation types.VestingPoolReservation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VestingPoolReservationsKeyPrefix)
	b := store.Get([]byte(vestingPoolName))
	if b == nil {
		found = false
		return
	}
	found = true
	k.cdc.MustUnmarshal(b, &vestingPoolReservation)
	return
}

func (k Keeper) SetVestingPoolReservation(ctx sdk.Context, vestingPoolName string, vestingPoolReservation types.VestingPoolReservation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VestingPoolReservationsKeyPrefix)
	av := k.cdc.MustMarshal(&vestingPoolReservation)
	store.Set([]byte(vestingPoolName), av)
}
