package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/v2/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) States(goCtx context.Context, req *types.QueryStatesRequest) (*types.QueryStatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	states := k.GetAllStates(ctx)
	coinsOnDistriubtorAccount := k.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
	return &types.QueryStatesResponse{States: states, CoinsOnDistributorAccount: coinsOnDistriubtorAccount}, nil
}
