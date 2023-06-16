package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StartEnergyTransfer(
	ctx sdk.Context,
	owner string,
	driver string,
	energyTransferOfferId uint64,
	chargerId string,
	offeredTariff uint64,
	energyToTransfer uint64,
	collateral math.Int,
) (*uint64, error) {
	k.Logger(ctx).Debug("start energy transfer", "owner", owner, "driver", driver, "energyTransferOfferId",
		energyTransferOfferId, "chargerId", chargerId, "offeredTariff", offeredTariff, "energyToTransfer", energyToTransfer, "collateral", collateral)

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
		return nil, errors.Wrap(c4eerrors.ErrParam, "wrong tariff")
	}

	k.Logger(ctx).Debug("start energy transfer send collateral", "driver", driver)

	driverAccAddress, err := sdk.AccAddressFromBech32(driver)
	if err != nil {
		return nil, err
	}
	coinsToSend := sdk.NewCoins(sdk.NewCoin(k.Denom(ctx), collateral))
	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, driverAccAddress, types.ModuleName, coinsToSend); err != nil {
		return nil, err
	}

	var energyTransferObj = types.EnergyTransfer{
		OwnerAccountAddress:   owner,
		DriverAccountAddress:  driver,
		EnergyTransferOfferId: energyTransferOfferId,
		ChargerId:             chargerId,
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
		ChargerId:              chargerId,
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
		return errors.Wrapf(types.ErrWrongEnergyTransferStatus,
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
	energyTransferObj, err := k.MustGetEnergyTransfer(ctx, energyTransferId)
	if err != nil {
		return err
	}

	if energyTransferObj.EnergyToTransfer == usedServiceUnits {
		if err = k.sendEntireCallateralToCPOwner(ctx, energyTransferObj); err != nil {
			return err
		}
	} else if energyTransferObj.EnergyToTransfer > usedServiceUnits {
		if err = k.sendUnusedCallateral(ctx, energyTransferObj, usedServiceUnits); err != nil {
			return err
		}
	} else if usedServiceUnits > energyTransferObj.EnergyToTransfer {
		// TODO: In some cases, the charger may charge more watt-hours than intended, please handle these cases accordingly.
		// Proposal - if the number of watt-hours that has been sent exceeds the number set earlier by less than 4, do not return
		// an error, otherwise inform the user about it
		if (usedServiceUnits - energyTransferObj.EnergyToTransfer) < types.SAFE_AMOUNT_TO_EXCEED_BY_CHARGER {
			if err = k.sendEntireCallateralToCPOwner(ctx, energyTransferObj); err != nil {
				return err
			}
		}
		// TODO: for now we don't inform the user about the exceeded amount of energy, but we should do it in the future
		if err = k.sendEntireCallateralToCPOwner(ctx, energyTransferObj); err != nil {
			return err
		}
		k.Logger(ctx).Error("used service units exeed energy to transfer", "energyToTransfer",
			energyTransferObj.EnergyToTransfer, "usedServiceUnits", usedServiceUnits)
	}

	energyTransferObj.Status = types.TransferStatus_PAID
	energyTransferObj.PaidDate = ctx.BlockTime()
	energyTransferObj.EnergyTransferred = usedServiceUnits
	k.SetEnergyTransfer(ctx, *energyTransferObj)

	offer, err := k.MustGetEnergyTransferOffer(ctx, energyTransferObj.EnergyTransferOfferId)
	if err != nil {
		return err
	}
	offer.ChargerStatus = types.ChargerStatus_ACTIVE
	k.SetEnergyTransferOffer(ctx, *offer)

	k.EmitChangeOfferStatusEvent(ctx, energyTransferObj.GetEnergyTransferOfferId(), types.ChargerStatus_BUSY, types.ChargerStatus_ACTIVE)

	event := &types.EventCompleteEnergyTransfer{
		EnergyTransferId:      energyTransferId,
		EnergyTransferOfferId: energyTransferObj.GetEnergyTransferOfferId(),
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
	return k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer)
}

func (k Keeper) sendUnusedCallateral(ctx sdk.Context, energyTransferObj *types.EnergyTransfer, usedServiceUnits uint64) error {
	denom := k.Denom(ctx)
	usedTokens := math.NewIntFromUint64(energyTransferObj.OfferedTariff * usedServiceUnits)
	coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, usedTokens))
	if err := k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer); err != nil {
		return err
	}

	unusedTokens := energyTransferObj.Collateral.Sub(usedTokens)
	coinsToTransfer = sdk.NewCoins(sdk.NewCoin(denom, unusedTokens))
	return k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.DriverAccountAddress, coinsToTransfer)
}

func (k Keeper) CancelEnergyTransfer(ctx sdk.Context, energyTransferId uint64) error {
	k.Logger(ctx).Debug("cancel energy transfer", "energyTransferId", energyTransferId)
	energyTransferObj, err := k.MustGetEnergyTransfer(ctx, energyTransferId)
	if err != nil {
		return err
	}

	// TODO: there were some problems with this (sometimes status was ongoing), leave commented out for now
	// if energyTransferObj.GetStatus() != types.TransferStatus_REQUESTED {
	// 	return nil, errors.Wrap(types.ErrWrongEnergyTransferStatus, energyTransferObj.GetStatus().String())
	// }

	energyTransferObj.Status = types.TransferStatus_CANCELLED

	offer, err := k.MustGetEnergyTransferOffer(ctx, energyTransferObj.EnergyTransferOfferId)
	if err != nil {
		return err
	}
	offer.ChargerStatus = types.ChargerStatus_ACTIVE

	// send the collateral back to the EV driver's account
	denom := k.Denom(ctx)
	coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, energyTransferObj.GetCollateral()))
	if err = k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.GetDriverAccountAddress(), coinsToTransfer); err != nil {
		return err
	}

	k.SetEnergyTransferOffer(ctx, *offer)
	k.SetEnergyTransfer(ctx, *energyTransferObj)

	k.EmitChangeOfferStatusEvent(ctx, energyTransferObj.GetEnergyTransferOfferId(), types.ChargerStatus_BUSY, types.ChargerStatus_ACTIVE)
	event := &types.EventCancelEnergyTransfer{
		EnergyTransferId: energyTransferId,
		ChargerId:        energyTransferObj.ChargerId,
		CancelReason:     "", // it's empty placeholder for now
	}
	if err = ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("cancel energy transfer emit event error", "event", event, "error", err.Error())
	}

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
