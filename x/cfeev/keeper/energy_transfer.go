package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) StartEnergyTransfer(
	ctx sdk.Context,
	driver string,
	energyTransferOfferId uint64,
	offeredTariff uint64,
	energyToTransfer uint64,
) (*uint64, error) {
	k.Logger(ctx).Debug("start energy transfer", "driver", driver, "energyTransferOfferId",
		energyTransferOfferId, "offeredTariff", offeredTariff, "energyToTransfer", energyToTransfer)

	if err := types.ValidateStartEnergyTransfer(driver, offeredTariff, energyToTransfer); err != nil {
		return nil, err
	}

	offer, err := k.MustGetEnergyTransferOffer(ctx, energyTransferOfferId)
	if err != nil {
		return nil, err
	}

	if offer.GetChargerStatus() != types.ChargerStatus_ACTIVE {
		return nil, types.ErrBusyCharger
	}
	offer.ChargerStatus = types.ChargerStatus_BUSY

	// check if the offered tariff has not been changed
	if offer.Tariff != offeredTariff {
		return nil, errors.Wrapf(c4eerrors.ErrParam, "wrong tariff expected %d got %d", offer.Tariff, offeredTariff)
	}

	k.Logger(ctx).Debug("start energy transfer send collateral", "driver", driver)

	driverAccAddress, err := sdk.AccAddressFromBech32(driver)
	if err != nil {
		return nil, err
	}

	collateral := math.NewIntFromUint64(offeredTariff * energyToTransfer)
	coinsToSend := sdk.NewCoins(sdk.NewCoin(k.Denom(ctx), collateral))
	spendableCoins := k.bankKeeper.SpendableCoins(ctx, driverAccAddress)
	if !coinsToSend.IsAllLTE(spendableCoins) {
		return nil, errors.Wrapf(sdkerrors.ErrInsufficientFunds, "owner balance is too small (%s < %s)", spendableCoins, coinsToSend)
	}

	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, driverAccAddress, types.ModuleName, coinsToSend); err != nil {
		return nil, err
	}

	var energyTransferObj = types.EnergyTransfer{
		Owner:                 offer.GetOwner(),
		Driver:                driver,
		EnergyTransferOfferId: energyTransferOfferId,
		ChargerId:             offer.GetChargerId(),
		Status:                types.TransferStatus_REQUESTED,
		OfferedTariff:         offeredTariff,
		EnergyToTransfer:      energyToTransfer,
		Collateral:            collateral,
	}
	energyTransferId := k.AppendEnergyTransfer(ctx, energyTransferObj)

	k.SetEnergyTransferOffer(ctx, *offer)
	k.EmitChangeOfferStatusEvent(ctx, energyTransferObj.GetEnergyTransferOfferId(), types.ChargerStatus_ACTIVE, types.ChargerStatus_BUSY)
	event := &types.EventEnergyTransferCreated{
		EnergyTransferId:       energyTransferId,
		ChargerId:              offer.GetChargerId(),
		EnergyAmountToTransfer: energyTransferObj.GetEnergyToTransfer(),
	}
	if err = ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("energy transfer created emit event error", "event", event, "error", err.Error())
	}
	k.Logger(ctx).Debug("start energy transfer ret", "energyTransferId", energyTransferId)
	return &energyTransferId, nil
}

func (k Keeper) EnergyTransferStarted(ctx sdk.Context, energyTransferId uint64) error {
	k.Logger(ctx).Debug("energy transfer started", "energyTransferId", energyTransferId)

	energyTransfer, err := k.MustGetEnergyTransfer(ctx, energyTransferId)
	if err != nil {
		return err
	}

	if energyTransfer.Status != types.TransferStatus_REQUESTED {
		return errors.Wrapf(sdkerrors.ErrInvalidType,
			"energy transfer status must be %s not %s", types.TransferStatus_name[int32(types.TransferStatus_REQUESTED)], energyTransfer.Status.String())
	}

	// REQUESTED ==> ONGOING
	energyTransfer.Status = types.TransferStatus_ONGOING

	k.SetEnergyTransfer(ctx, *energyTransfer)

	event := &types.EventBeginEnergyTransfer{
		EnergyTransferId:      energyTransferId,
		EnergyTransferOfferId: energyTransfer.GetEnergyTransferOfferId(),
	}
	if err = ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("begin energy transfer emit event error", "event", event, "error", err.Error())
	}
	return nil
}

func (k Keeper) EnergyTransferCompleted(ctx sdk.Context, energyTransferId uint64, usedServiceUnits uint64) error {
	k.Logger(ctx).Debug("energy transfer completed", "energyTransferId", energyTransferId, "usedServiceUnits", usedServiceUnits)
	energyTransfer, err := k.MustGetEnergyTransfer(ctx, energyTransferId)
	if err != nil {
		return err
	}

	if energyTransfer.Status != types.TransferStatus_REQUESTED && energyTransfer.Status != types.TransferStatus_ONGOING {
		return errors.Wrapf(sdkerrors.ErrInvalidType,
			"energy transfer status must be %s or %s not %s",
			types.TransferStatus_name[int32(types.TransferStatus_REQUESTED)],
			types.TransferStatus_name[int32(types.TransferStatus_ONGOING)],
			energyTransfer.Status.String())
	}

	if energyTransfer.EnergyToTransfer == usedServiceUnits {
		if err = k.sendEntireCallateralToCPOwner(ctx, energyTransfer); err != nil {
			return err
		}
	} else if energyTransfer.EnergyToTransfer > usedServiceUnits {
		if err = k.sendUnusedCallateral(ctx, energyTransfer, usedServiceUnits); err != nil {
			return err
		}
	} else if usedServiceUnits > energyTransfer.EnergyToTransfer {
		// TODO: In some cases, the charger may charge more watt-hours than intended, please handle these cases accordingly.
		// Proposal - if the number of watt-hours that has been sent exceeds the number set earlier by less than 4, do not return
		// an error, otherwise inform the user about it
		if (usedServiceUnits - energyTransfer.EnergyToTransfer) < types.SafeAmountToExceedByCharger {
			if err = k.sendEntireCallateralToCPOwner(ctx, energyTransfer); err != nil {
				return err
			}
		} else {
			// TODO: for now we don't inform the user about the exceeded amount of energy, but we should do it in the future
			if err = k.sendEntireCallateralToCPOwner(ctx, energyTransfer); err != nil {
				return err
			}
			k.Logger(ctx).Error("used service units exeed energy to transfer", "energyToTransfer",
				energyTransfer.EnergyToTransfer, "usedServiceUnits", usedServiceUnits)
		}
	}

	energyTransfer.Status = types.TransferStatus_PAID
	energyTransfer.PaidDate = ctx.BlockTime()
	energyTransfer.EnergyTransferred = usedServiceUnits
	k.SetEnergyTransfer(ctx, *energyTransfer)

	offer, err := k.MustGetEnergyTransferOffer(ctx, energyTransfer.EnergyTransferOfferId)
	if err != nil {
		return err
	}
	offer.ChargerStatus = types.ChargerStatus_ACTIVE
	k.SetEnergyTransferOffer(ctx, *offer)

	k.EmitChangeOfferStatusEvent(ctx, energyTransfer.GetEnergyTransferOfferId(), types.ChargerStatus_BUSY, types.ChargerStatus_ACTIVE)

	event := &types.EventCompleteEnergyTransfer{
		EnergyTransferId:      energyTransferId,
		EnergyTransferOfferId: energyTransfer.GetEnergyTransferOfferId(),
		EnergyTransferred:     usedServiceUnits,
	}
	if err = ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("complete energy transfer emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) sendEntireCallateralToCPOwner(ctx sdk.Context, energyTransferObj *types.EnergyTransfer) error {
	denom := k.Denom(ctx)
	coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, energyTransferObj.GetCollateral()))
	return k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.GetOwner(), coinsToTransfer)
}

func (k Keeper) sendUnusedCallateral(ctx sdk.Context, energyTransferObj *types.EnergyTransfer, usedServiceUnits uint64) error {
	denom := k.Denom(ctx)
	usedTokens := math.NewIntFromUint64(energyTransferObj.OfferedTariff * usedServiceUnits)
	coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, usedTokens))
	if err := k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.GetOwner(), coinsToTransfer); err != nil {
		return err
	}

	unusedTokens := energyTransferObj.Collateral.Sub(usedTokens)
	coinsToTransfer = sdk.NewCoins(sdk.NewCoin(denom, unusedTokens))
	return k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.GetDriver(), coinsToTransfer)
}

func (k Keeper) CancelEnergyTransfer(ctx sdk.Context, energyTransferId uint64) error {
	k.Logger(ctx).Debug("cancel energy transfer", "energyTransferId", energyTransferId)
	energyTransfer, err := k.MustGetEnergyTransfer(ctx, energyTransferId)
	if err != nil {
		return err
	}

	// TODO: there were some problems with this (sometimes status was ongoing), leave commented out for now
	if energyTransfer.GetStatus() != types.TransferStatus_REQUESTED {
		return errors.Wrapf(sdkerrors.ErrInvalidType,
			"energy transfer status must be %s not %s",
			types.TransferStatus_name[int32(types.TransferStatus_REQUESTED)],
			energyTransfer.Status.String())
	}

	energyTransfer.Status = types.TransferStatus_CANCELLED

	offer, err := k.MustGetEnergyTransferOffer(ctx, energyTransfer.EnergyTransferOfferId)
	if err != nil {
		return err
	}
	offer.ChargerStatus = types.ChargerStatus_ACTIVE

	// send the collateral back to the EV driver's account
	denom := k.Denom(ctx)
	coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, energyTransfer.GetCollateral()))
	if err = k.parseAddressAndSendTokensFromModule(ctx, energyTransfer.GetDriver(), coinsToTransfer); err != nil {
		return err
	}

	k.SetEnergyTransferOffer(ctx, *offer)
	k.SetEnergyTransfer(ctx, *energyTransfer)

	k.EmitChangeOfferStatusEvent(ctx, energyTransfer.GetEnergyTransferOfferId(), types.ChargerStatus_BUSY, types.ChargerStatus_ACTIVE)
	event := &types.EventCancelEnergyTransfer{
		EnergyTransferId: energyTransferId,
		ChargerId:        energyTransfer.ChargerId,
		CancelReason:     "", // it's empty placeholder for now
	}
	if err = ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("cancel energy transfer emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) RemoveEnergyTransfer(ctx sdk.Context, id uint64) error {
	energyTransfer, err := k.MustGetEnergyTransfer(ctx, id)
	if err != nil {
		return err
	}

	if energyTransfer.Status != types.TransferStatus_PAID && energyTransfer.Status != types.TransferStatus_CANCELLED {
		return errors.Wrapf(sdkerrors.ErrInvalidType, "energy transfer status must be %s or %s not %s",
			types.TransferStatus_name[int32(types.TransferStatus_PAID)],
			types.TransferStatus_name[int32(types.TransferStatus_CANCELLED)],
			energyTransfer.Status.String())
	}

	k.DeleteEnergyTransfer(ctx, id)
	return nil
}

func (k Keeper) EmitChangeOfferStatusEvent(ctx sdk.Context, energyTransferOfferId uint64, oldStatus, newStatus types.ChargerStatus) {
	event := &types.EventChangeOfferStatus{
		EnergyTransferOfferId: energyTransferOfferId,
		OldStatus:             oldStatus,
		NewStatus:             newStatus,
	}
	if err := ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("change offer status emit event error", "event", event, "error", err.Error())
	}
}

func (k Keeper) parseAddressAndSendTokensFromModule(ctx sdk.Context, targetAccountAddress string, collateralCoins sdk.Coins) error {
	target, err := sdk.AccAddressFromBech32(targetAccountAddress)
	if err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, target, collateralCoins)
}
