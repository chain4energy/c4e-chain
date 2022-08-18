package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.SubDistributors(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// SubDistributors returns the SubDistributors param
func (k Keeper) SubDistributors(ctx sdk.Context) (res []types.SubDistributor) {
	k.paramstore.Get(ctx, types.KeySubDistributors, &res)
	return
}
