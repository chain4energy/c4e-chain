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

func (k Keeper) UsersEntries(c context.Context, req *types.QueryUsersEntriesRequest) (*types.QueryUsersEntriesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userEntry []types.UserEntry
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	claimRecordStore := prefix.NewStore(store, types.KeyPrefix(types.UsersEntriesKeyPrefix))

	pageRes, err := query.Paginate(claimRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var userAirdropEntry types.UserEntry
		if err := k.cdc.Unmarshal(value, &userAirdropEntry); err != nil {
			return err
		}

		userEntry = append(userEntry, userAirdropEntry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUsersEntriesResponse{UsersEntries: userEntry, Pagination: pageRes}, nil
}

func (k Keeper) UserEntry(c context.Context, req *types.QueryUserEntryRequest) (*types.QueryUserEntryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUserEntry(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryUserEntryResponse{UserEntry: val}, nil
}

func (k Keeper) AirdropDistrubitions(c context.Context, req *types.QueryAirdropDistrubitionsRequest) (*types.QueryAirdropDistrubitionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAirdropDistrubitions(
		ctx,
		req.CampaignId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryAirdropDistrubitionsResponse{AirdropCoins: val.AirdropCoins}, nil
}

func (k Keeper) AirdropClaimsLeft(c context.Context, req *types.QueryAirdropClaimsLeftRequest) (*types.QueryAirdropClaimsLeftResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAirdropClaimsLeft(
		ctx,
		req.CampaignId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryAirdropClaimsLeftResponse{AirdropCoins: val.AirdropCoins}, nil
}
