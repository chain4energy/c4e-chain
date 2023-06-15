package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListOwnEnergyTransfer(goCtx context.Context, req *types.QueryListOwnEnergyTransferRequest) (*types.QueryListOwnEnergyTransferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var energyTransfers []types.EnergyTransfer
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	energyTransferStore := prefix.NewStore(store, types.EnergyTransferKey)

	pageRes, err := query.Paginate(energyTransferStore, req.Pagination, func(key []byte, value []byte) error {
		var energyTransfer types.EnergyTransfer
		if err := k.cdc.Unmarshal(value, &energyTransfer); err != nil {
			return err
		}

		if energyTransfer.GetDriverAccountAddress() == req.GetDriverAccAddress() && energyTransfer.Status == types.TransferStatus(req.GetTransferStatus()) {
			energyTransfers = append(energyTransfers, energyTransfer)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListOwnEnergyTransferResponse{EnergyTransfer: energyTransfers, Pagination: pageRes}, nil
}
