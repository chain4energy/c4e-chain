package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TokenParamsAll(c context.Context, req *types.QueryAllTokenParamsRequest) (*types.QueryAllTokenParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokenParamss []types.TokenParams
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tokenParamsStore := prefix.NewStore(store, types.KeyPrefix(types.TokenParamsKeyPrefix))

	pageRes, err := query.Paginate(tokenParamsStore, req.Pagination, func(key []byte, value []byte) error {
		var tokenParams types.TokenParams
		if err := k.cdc.Unmarshal(value, &tokenParams); err != nil {
			return err
		}

		tokenParamss = append(tokenParamss, tokenParams)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokenParamsResponse{TokenParams: tokenParamss, Pagination: pageRes}, nil
}

func (k Keeper) TokenParams(c context.Context, req *types.QueryGetTokenParamsRequest) (*types.QueryGetTokenParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTokenParams(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTokenParamsResponse{TokenParams: val}, nil
}
