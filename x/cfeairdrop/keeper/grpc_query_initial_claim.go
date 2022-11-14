package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) InitialClaimAll(c context.Context, req *types.QueryAllInitialClaimRequest) (*types.QueryAllInitialClaimResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var initialClaims []types.InitialClaim
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	initialClaimStore := prefix.NewStore(store, types.KeyPrefix(types.InitialClaimKeyPrefix))

	pageRes, err := query.Paginate(initialClaimStore, req.Pagination, func(key []byte, value []byte) error {
		var initialClaim types.InitialClaim
		if err := k.cdc.Unmarshal(value, &initialClaim); err != nil {
			return err
		}

		initialClaims = append(initialClaims, initialClaim)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInitialClaimResponse{InitialClaim: initialClaims, Pagination: pageRes}, nil
}

func (k Keeper) InitialClaim(c context.Context, req *types.QueryGetInitialClaimRequest) (*types.QueryGetInitialClaimResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetInitialClaim(
		ctx,
		req.CampaignId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInitialClaimResponse{InitialClaim: val}, nil
}
