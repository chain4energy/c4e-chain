package keeper

import (
	"context"
	"math"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnergyTransferCompletedRequest(goCtx context.Context, msg *types.MsgEnergyTransferCompletedRequest) (*types.MsgEnergyTransferCompletedRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	energyTransferObj, found := k.GetEnergyTransfer(ctx, msg.EnergyTransferId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrEnergyTransferNotFound, "energy transfer not found")
	}

	var err error

	date, err := ctx.BlockTime().UTC().MarshalText()
	if err != nil {
		energyTransferObj.PaidDate = "Error"
	} else {
		energyTransferObj.PaidDate = string(date)
	}

	usedServiceUnits := msg.GetUsedServiceUnits()
	energyTransferObj.EnergyTransferred = usedServiceUnits

	if energyTransferObj.EnergyToTransfer == usedServiceUnits {
		// send entire callateral to CP owner's account
		coinsToTransfer := strconv.FormatInt(int64(energyTransferObj.GetCollateral()), 10) + "uc4e"
		err = k.sendTokensToTargetAccount(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer)
		energyTransferObj.Status = types.TransferStatus_PAID

	} else if energyTransferObj.EnergyToTransfer > usedServiceUnits {
		// calculate used tokens
		usedTokens := energyTransferObj.OfferedTariff * usedServiceUnits
		coinsToTransfer := strconv.FormatInt(int64(usedTokens), 10) + "uc4e"
		err = k.sendTokensToTargetAccount(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer)

		// calculate unused tokens
		unusedTokens := energyTransferObj.Collateral - uint64(usedTokens)
		coinsToTransfer = strconv.FormatInt(int64(unusedTokens), 10) + "uc4e"
		err = k.sendTokensToTargetAccount(ctx, energyTransferObj.DriverAccountAddress, coinsToTransfer)

		// set status
		energyTransferObj.Status = types.TransferStatus_PAID
		if err != nil {
			// TODO:
		}
		// TODO: handle exceeded limit value

	} else if usedServiceUnits > energyTransferObj.EnergyToTransfer {
		if (usedServiceUnits - energyTransferObj.EnergyToTransfer) < 4 {
			// send entire callateral to CP owner's account
			coinsToTransfer := strconv.FormatInt(int64(energyTransferObj.GetCollateral()), 10) + "uc4e"
			err = k.sendTokensToTargetAccount(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer)
			energyTransferObj.Status = types.TransferStatus_PAID
		}
		// TODO:
	}

	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCoinTransferFailed, "coin transfer failed")
	}

	// get energy transfer offer object by offer id
	offer, found := k.GetEnergyTransferOffer(ctx, energyTransferObj.EnergyTransferOfferId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrEnergyTransferOfferNotFound, "energy transfer offer not found")
	}
	offer.ChargerStatus = types.ChargerStatus_ACTIVE

	// update both entities
	k.SetEnergyTransferOffer(ctx, offer)
	k.SetEnergyTransfer(ctx, energyTransferObj)

	// TODO: Handling the response
	return &types.MsgEnergyTransferCompletedRequestResponse{}, nil
}

func (k msgServer) sendTokensToTargetAccount(ctx sdk.Context, targetAccountAddress string, collateral string) error {
	target, err := sdk.AccAddressFromBech32(targetAccountAddress)
	if err != nil {
		panic(err)
	}
	collateralCoins, err := sdk.ParseCoinsNormalized(collateral)
	if err != nil {
		panic(err)
	}
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, target, collateralCoins)

	if sdkError != nil {
		return sdkError
	}

	return nil
}

func compareWithTolerane(a, b float64) bool {
	tolerance := 0.001
	if diff := math.Abs(a - b); diff < tolerance {
		return true
	} else {
		return false
	}
}
