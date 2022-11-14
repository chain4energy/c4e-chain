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

func (k Keeper) MissionAll(c context.Context, req *types.QueryAllMissionRequest) (*types.QueryAllMissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var missions []types.Mission
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	missionStore := prefix.NewStore(store, types.KeyPrefix(types.MissionKeyPrefix))

	pageRes, err := query.Paginate(missionStore, req.Pagination, func(key []byte, value []byte) error {
		var mission types.Mission
		if err := k.cdc.Unmarshal(value, &mission); err != nil {
			return err
		}

		missions = append(missions, mission)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMissionResponse{Mission: missions, Pagination: pageRes}, nil
}

func (k Keeper) Mission(c context.Context, req *types.QueryGetMissionRequest) (*types.QueryGetMissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMission(
		ctx,
		req.CampaignId,
		req.MissionId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMissionResponse{Mission: val}, nil
}
