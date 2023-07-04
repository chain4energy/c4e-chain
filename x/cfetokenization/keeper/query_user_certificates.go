package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserCertificatesAll(goCtx context.Context, req *types.QueryAllUserCertificatesRequest) (*types.QueryAllUserCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userCertificatess []types.UserCertificates
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	userCertificatesStore := prefix.NewStore(store, types.KeyPrefix(types.UserCertificatesKey))

	pageRes, err := query.Paginate(userCertificatesStore, req.Pagination, func(key []byte, value []byte) error {
		var userCertificates types.UserCertificates
		if err := k.cdc.Unmarshal(value, &userCertificates); err != nil {
			return err
		}

		userCertificatess = append(userCertificatess, userCertificates)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserCertificatesResponse{UserCertificates: userCertificatess, Pagination: pageRes}, nil
}

func (k Keeper) UserCertificates(goCtx context.Context, req *types.QueryGetUserCertificatesRequest) (*types.QueryGetUserCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	userCertificates, found := k.GetUserCertificates(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetUserCertificatesResponse{UserCertificates: userCertificates}, nil
}
