package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) States(goCtx context.Context, req *types.QueryStatesRequest) (*types.QueryStatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	states := k.GetALlStates(ctx)
	coinsOnDistriubtorAccount := k.GetAccountCoinsForModuleAccount(ctx, types.CollectorName)
	return &types.QueryStatesResponse{States: states, CoinsOnDistributorAccount: coinsOnDistriubtorAccount}, nil
}
