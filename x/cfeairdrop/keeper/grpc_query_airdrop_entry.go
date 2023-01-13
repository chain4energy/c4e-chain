package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AirdropEntryAll(c context.Context, req *types.QueryAllAirdropEntryRequest) (*types.QueryAllAirdropEntryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var airdropEntrys []types.AirdropEntry
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	airdropEntryStore := prefix.NewStore(store, types.KeyPrefix(types.AirdropEntryKey))

	pageRes, err := query.Paginate(airdropEntryStore, req.Pagination, func(key []byte, value []byte) error {
		var airdropEntry types.AirdropEntry
		if err := k.cdc.Unmarshal(value, &airdropEntry); err != nil {
			return err
		}

		airdropEntrys = append(airdropEntrys, airdropEntry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAirdropEntryResponse{AirdropEntry: airdropEntrys, Pagination: pageRes}, nil
}

func (k Keeper) AirdropEntry(c context.Context, req *types.QueryGetAirdropEntryRequest) (*types.QueryGetAirdropEntryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	airdropEntry, found := k.GetAirdropEntry(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAirdropEntryResponse{AirdropEntry: airdropEntry}, nil
}
