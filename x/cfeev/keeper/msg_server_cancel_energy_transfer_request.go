package keeper

import (
	"context"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CancelEnergyTransferRequest(goCtx context.Context, msg *types.MsgCancelEnergyTransferRequest) (*types.MsgCancelEnergyTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	energyTransferObj, found := k.GetEnergyTransfer(ctx, msg.EnergyTransferId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrEnergyTransferNotFound, "energy transfer not found")
	}

	// if energyTransferObj.GetStatus() != types.TransferStatus_REQUESTED {
	// 	return nil, sdkerrors.Wrap(types.ErrWrongEnergyTransferStatus, energyTransferObj.GetStatus().String())
	// }

	energyTransferObj.Status = types.TransferStatus_CANCELLED

	// get energy transfer offer object by offer id
	offer, found := k.GetEnergyTransferOffer(ctx, energyTransferObj.EnergyTransferOfferId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrEnergyTransferOfferNotFound, "energy transfer offer not found")
	}
	offer.ChargerStatus = types.ChargerStatus_ACTIVE

	// send the collateral back to the EV driver's account
	coinsToTransfer := strconv.FormatInt(int64(energyTransferObj.GetCollateral()), 10) + "uc4e"
	err := k.sendTokensToTargetAccount(ctx, energyTransferObj.GetDriverAccountAddress(), coinsToTransfer)

	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCoinTransferFailed, "coin transfer failed")
	}

	// update both entities
	k.SetEnergyTransferOffer(ctx, offer)
	k.SetEnergyTransfer(ctx, energyTransferObj)

	return &types.MsgCancelEnergyTransferRequestResponse{}, nil
}
