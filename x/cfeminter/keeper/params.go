package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	minter := k.Minter(ctx)
	startTime := k.StartTime(ctx)
	return types.NewParams(k.MintDenom(ctx), startTime, minter)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MintDenom returns the denom param
func (k Keeper) MintDenom(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyMintDenom, &res)
	return
}

// MintDenom returns the denom param
func (k Keeper) Minter(ctx sdk.Context) (res []*types.Minter) {
	k.paramstore.Get(ctx, types.KeyMinters, &res)
	return
}

// MintDenom returns the denom param
func (k Keeper) StartTime(ctx sdk.Context) (res time.Time) {
	k.paramstore.Get(ctx, types.KeyStartTime, &res)
	return
}
