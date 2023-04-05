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
	claimRecordStore := prefix.NewStore(store, types.KeyPrefix(types.UserEntryKeyPrefix))

	pageRes, err := query.Paginate(claimRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var userclaimRecord types.UserEntry
		if err := k.cdc.Unmarshal(value, &userclaimRecord); err != nil {
			return err
		}

		userEntry = append(userEntry, userclaimRecord)
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

func (k Keeper) CampaignTotalAmount(c context.Context, req *types.QueryCampaignTotalAmountRequest) (*types.QueryCampaignTotalAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCampaignTotalAmount(
		ctx,
		req.CampaignId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryCampaignTotalAmountResponse{Amount: val.Amount}, nil
}

func (k Keeper) CampaignAmountLeft(c context.Context, req *types.QueryCampaignAmountLeftRequest) (*types.QueryCampaignAmountLeftResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCampaignAmountLeft(
		ctx,
		req.CampaignId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryCampaignAmountLeftResponse{Amount: val.Amount}, nil
}
