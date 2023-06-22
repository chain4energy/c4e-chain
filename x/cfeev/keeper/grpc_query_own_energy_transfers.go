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

func (k Keeper) OwnEnergyTransfers(goCtx context.Context, req *types.QueryOwnEnergyTransfersRequest) (*types.QueryOwnEnergyTransfersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var EnergyTransfers []types.EnergyTransfer
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	EnergyTransferstore := prefix.NewStore(store, types.EnergyTransferKey)
	_, err := sdk.AccAddressFromBech32(req.Driver)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "wrong driver address (%s)", err.Error())
	}
	pageRes, err := query.Paginate(EnergyTransferstore, req.Pagination, func(key []byte, value []byte) error {
		var energyTransfer types.EnergyTransfer
		if err := k.cdc.Unmarshal(value, &energyTransfer); err != nil {
			return err
		}

		if energyTransfer.GetDriverAccountAddress() == req.GetDriver() && energyTransfer.Status == types.TransferStatus(req.GetTransferStatus()) {
			EnergyTransfers = append(EnergyTransfers, energyTransfer)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryOwnEnergyTransfersResponse{EnergyTransfers: EnergyTransfers, Pagination: pageRes}, nil
}
