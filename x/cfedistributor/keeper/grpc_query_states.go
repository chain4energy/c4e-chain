package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) States(goCtx context.Context, req *types.QueryStatesRequest) (*types.QueryStatesResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	states := k.GetAllStates(ctx)
	coinsOnDistriubtorAccount := k.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
	return &types.QueryStatesResponse{States: states, CoinsOnDistributorAccount: coinsOnDistriubtorAccount}, nil
}
