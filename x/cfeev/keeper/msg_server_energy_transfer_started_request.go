package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnergyTransferStartedRequest(goCtx context.Context, msg *types.MsgEnergyTransferStartedRequest) (*types.MsgEnergyTransferStartedRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	energyTransfer, found := k.GetEnergyTransfer(ctx, msg.EnergyTransferId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrEnergyTransferNotFound, "energy transfer not found")
	}

	if energyTransfer.Status != types.TransferStatus_REQUESTED {
		return nil, sdkerrors.Wrap(types.ErrWrongEnergyTransferStatus, energyTransfer.Status.String())
	}

	// REQUESTED ==> ONGOING
	energyTransfer.Status = types.TransferStatus_ONGOING

	// update energyTransfer instance in the KVStore
	k.SetEnergyTransfer(ctx, energyTransfer)

	return &types.MsgEnergyTransferStartedRequestResponse{}, nil
}
