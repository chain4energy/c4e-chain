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

func (k Keeper) UserDevicesAll(goCtx context.Context, req *types.QueryAllUserDevicesRequest) (*types.QueryAllUserDevicesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userDevicess []types.UserDevices
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	userDevicesStore := prefix.NewStore(store, types.KeyPrefix(types.UserDevicesKey))

	pageRes, err := query.Paginate(userDevicesStore, req.Pagination, func(key []byte, value []byte) error {
		var userDevices types.UserDevices
		if err := k.cdc.Unmarshal(value, &userDevices); err != nil {
			return err
		}

		userDevicess = append(userDevicess, userDevices)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserDevicesResponse{UserDevices: userDevicess, Pagination: pageRes}, nil
}

func (k Keeper) UserDevices(goCtx context.Context, req *types.QueryGetUserDevicesRequest) (*types.QueryGetUserDevicesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	userDevices, found := k.GetUserDevices(ctx, req.Owner)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetUserDevicesResponse{UserDevices: userDevices}, nil
}
