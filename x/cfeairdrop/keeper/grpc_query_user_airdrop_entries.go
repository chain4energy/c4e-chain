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

func (k Keeper) UsersAirdropEntries(c context.Context, req *types.QueryUsersAirdropEntriesRequest) (*types.QueryUsersAirdropEntriesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userAirdropEntries []types.UserAirdropEntries
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	claimRecordStore := prefix.NewStore(store, types.KeyPrefix(types.UserAirdropEntriesKeyPrefix))

	pageRes, err := query.Paginate(claimRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var userAirdropEntry types.UserAirdropEntries
		if err := k.cdc.Unmarshal(value, &userAirdropEntry); err != nil {
			return err
		}

		userAirdropEntries = append(userAirdropEntries, userAirdropEntry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUsersAirdropEntriesResponse{UsersAirdropEntries: userAirdropEntries, Pagination: pageRes}, nil
}

func (k Keeper) UserAirdropEntries(c context.Context, req *types.QueryUserAirdropEntriesRequest) (*types.QueryUserAirdropEntriesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUserAirdropEntries(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryUserAirdropEntriesResponse{UserAirdropEntries: val}, nil
}
