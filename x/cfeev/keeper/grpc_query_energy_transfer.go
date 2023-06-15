package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EnergyTransferAll(c context.Context, req *types.QueryAllEnergyTransferRequest) (*types.QueryAllEnergyTransferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var energyTransfers []types.EnergyTransfer
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	energyTransferStore := prefix.NewStore(store, types.EnergyTransferKey)

	pageRes, err := query.Paginate(energyTransferStore, req.Pagination, func(key []byte, value []byte) error {
		var energyTransfer types.EnergyTransfer
		if err := k.cdc.Unmarshal(value, &energyTransfer); err != nil {
			return err
		}

		energyTransfers = append(energyTransfers, energyTransfer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEnergyTransferResponse{EnergyTransfer: energyTransfers, Pagination: pageRes}, nil
}

func (k Keeper) EnergyTransfer(c context.Context, req *types.QueryGetEnergyTransferRequest) (*types.QueryGetEnergyTransferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	energyTransfer, found := k.GetEnergyTransfer(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetEnergyTransferResponse{EnergyTransfer: energyTransfer}, nil
}
