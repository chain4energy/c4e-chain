package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams gets the auth module's parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	return types.NewParams(k.RoutingDistributor(ctx))
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) RoutingDistributor(ctx sdk.Context) (routingDistributor []types.SubDistributor) {
	k.paramstore.Get(ctx, types.KeyRoutingDistributor, &routingDistributor)
	return
}
