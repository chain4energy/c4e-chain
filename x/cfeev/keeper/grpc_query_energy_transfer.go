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

func (k Keeper) AllEnergyTransfers(c context.Context, req *types.QueryAllEnergyTransfersRequest) (*types.QueryAllEnergyTransfersResponse, error) {
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

	return &types.QueryAllEnergyTransfersResponse{EnergyTransfer: energyTransfers, Pagination: pageRes}, nil
}

func (k Keeper) EnergyTransfer(c context.Context, req *types.QueryEnergyTransferRequest) (*types.QueryEnergyTransferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	energyTransfer, found := k.GetEnergyTransfer(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryEnergyTransferResponse{EnergyTransfer: energyTransfer}, nil
}

func (k Keeper) EnergyTransfers(goCtx context.Context, req *types.QueryEnergyTransfersRequest) (*types.QueryEnergyTransfersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "wrong owner address (%s)", err.Error())
	}

	store := ctx.KVStore(k.storeKey)
	energyTransferStore := prefix.NewStore(store, types.EnergyTransferKey)
	var energyTransfers []types.EnergyTransfer
	pageRes, err := query.Paginate(energyTransferStore, req.Pagination, func(key []byte, value []byte) error {
		var energyTransfer types.EnergyTransfer
		if err := k.cdc.Unmarshal(value, &energyTransfer); err != nil {
			return err
		}

		if energyTransfer.GetOwnerAccountAddress() == req.GetOwner() && energyTransfer.Status == types.TransferStatus_PAID {
			energyTransfers = append(energyTransfers, energyTransfer)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryEnergyTransfersResponse{EnergyTransfer: energyTransfers, Pagination: pageRes}, nil
}
