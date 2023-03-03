package keeper

import (
	"encoding/binary"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetParams sets the x/mint module parameters.
func (k Keeper) SetParams(ctx sdk.Context, p types.Params) error {
	if err := p.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&p)
	store.Set(types.ParamsKey, bz)
	return nil
}

// GetParams returns the current x/mint module parameters.
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return p
	}

	k.cdc.MustUnmarshal(bz, &p)
	return p
}

// SetParams sets the x/mint module parameters.
func (k Keeper) SetMinterConfig(ctx sdk.Context, minterConfig types.MinterConfigI, id uint64) error {
	bz, err := k.MarshalConfig(minterConfig)
	if err != nil {
		return err
	}

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store := ctx.KVStore(k.storeKey)
	store.Set(key, bz)

	return nil
}

// MintDenom returns the denom param
func (k Keeper) MintDenom(ctx sdk.Context) (res string) {
	return k.GetParams(ctx).MintDenom
}

func (k Keeper) MarshalConfig(minterConfig types.MinterConfigI) ([]byte, error) {
	return k.cdc.MarshalInterface(minterConfig)
}

func (k Keeper) UnmarshalConfig(bz []byte) (types.MinterConfigI, error) {
	var minterConfig types.MinterConfigI
	return minterConfig, k.cdc.UnmarshalInterface(bz, &minterConfig)
}
