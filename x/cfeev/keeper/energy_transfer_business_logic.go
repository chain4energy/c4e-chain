package keeper

import (
	"cosmossdk.io/errors"
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
	offeredTariff int32,
	energyToTransfer int32,
	collateral sdk.Coin,
) (uint64, error) {
	k.Logger(ctx).Debug("start energy transfer", "driver", driver, "chargerId", chargerId)

	var energyTransferObj = types.EnergyTransfer{
		OwnerAccountAddress:   owner,
		DriverAccountAddress:  driver,
		EnergyTransferOfferId: energyTransferOfferId,
		ChargerId:             chargerId,
		Status:                types.TransferStatus_REQUESTED,
		OfferedTariff:         offeredTariff,
		EnergyToTransfer:      energyToTransfer,
		Collateral:            collateral.Amount.Uint64(),
	}

	offer, err := k.MustGetEnergyTransferOffer(ctx, energyTransferObj.EnergyTransferOfferId)
	if err != nil {
		return 0, err
	}

	if offer.GetChargerStatus() == types.ChargerStatus_ACTIVE {
		offer.ChargerStatus = types.ChargerStatus_BUSY
	} else {
		return 0, errors.Wrap(types.ErrBusyCharger, "busy charger")
	}

	// check if the offered tariff has not been changed
	if offer.Tariff != energyTransferObj.OfferedTariff {
		return 0, errors.Wrap(c4eerrors.ErrParam, "wrong tariff")
	}

	k.Logger(ctx).Debug("start energy transfer send collateral", "driver", energyTransferObj.DriverAccountAddress)

	driverAccAddress, err := sdk.AccAddressFromBech32(energyTransferObj.DriverAccountAddress)
	if err != nil {
		panic(err)
	}
	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, driverAccAddress, types.ModuleName, sdk.NewCoins(collateral)); err != nil {
		return 0, err
	}

	energyTransferId := k.AppendEnergyTransfer(ctx, energyTransferObj)

	k.SetEnergyTransferOffer(ctx, *offer)
	k.EmitEnergyTransferCreatedEvent(ctx, energyTransferId, chargerId, energyTransferObj.GetEnergyToTransfer())
	k.EmitChangeOfferStatusEvent(ctx, energyTransferObj.GetEnergyTransferOfferId(), types.ChargerStatus_ACTIVE.String(), types.ChargerStatus_BUSY.String())

	return energyTransferId, nil
}

func (k Keeper) EnergyTransferStarted(ctx sdk.Context, energyTransferId uint64) error {
	energyTransfer, err := k.MustGetEnergyTransfer(ctx, energyTransferId)
	if err != nil {
		return err
	}

	if energyTransfer.Status != types.TransferStatus_REQUESTED {
		return errors.Wrap(types.ErrWrongEnergyTransferStatus, energyTransfer.Status.String())
	}

	// REQUESTED ==> ONGOING
	energyTransfer.Status = types.TransferStatus_ONGOING

	k.SetEnergyTransfer(ctx, *energyTransfer)
	k.EmitBeginEnergyTransferEvent(ctx, energyTransferId, energyTransfer.GetEnergyTransferOfferId())

	return nil
}

func (k Keeper) EnergyTransferCompleted(ctx sdk.Context, energyTransferId uint64, usedServiceUnits int32) error {
	energyTransferObj, found := k.GetEnergyTransfer(ctx, energyTransferId)
	if !found {
		return errors.Wrap(types.ErrEnergyTransferNotFound, "energy transfer not found")
	}

	var err error

	date, err := ctx.BlockTime().UTC().MarshalText()
	if err != nil {
		energyTransferObj.PaidDate = "Error"
	} else {
		energyTransferObj.PaidDate = string(date)
	}

	denom := k.Denom(ctx)
	energyTransferObj.EnergyTransferred = usedServiceUnits

	if energyTransferObj.EnergyToTransfer == usedServiceUnits {
		// send entire callateral to CP owner's account
		amount := sdk.NewInt(int64(energyTransferObj.GetCollateral()))
		coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, amount))
		if err = k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer); err != nil {
			return err
		}
		energyTransferObj.Status = types.TransferStatus_PAID

	} else if energyTransferObj.EnergyToTransfer > usedServiceUnits {
		// calculate used tokens
		usedTokens := energyTransferObj.OfferedTariff * usedServiceUnits
		amount := sdk.NewInt(int64(usedTokens))
		coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, amount))
		if err = k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer); err != nil {
			return err
		}

		// calculate unused tokens
		unusedTokens := energyTransferObj.Collateral - uint64(usedTokens)
		amount = sdk.NewInt(int64(unusedTokens))
		coinsToTransfer = sdk.NewCoins(sdk.NewCoin(denom, amount))
		if err = k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.DriverAccountAddress, coinsToTransfer); err != nil {
			return err
		}

		// set status
		energyTransferObj.Status = types.TransferStatus_PAID

		// TODO: handle exceeded limit value

	} else if usedServiceUnits > energyTransferObj.EnergyToTransfer {
		if (usedServiceUnits - energyTransferObj.EnergyToTransfer) < 4 {
			// send entire callateral to CP owner's account
			amount := sdk.NewInt(int64(energyTransferObj.GetCollateral()))
			coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, amount))
			if err = k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer); err != nil {
				return err
			}
			energyTransferObj.Status = types.TransferStatus_PAID
		}
		// TODO:
	}

	// get energy transfer offer object by offer id
	offer, err := k.MustGetEnergyTransferOffer(ctx, energyTransferObj.EnergyTransferOfferId)
	if err != nil {
		return err
	}
	offer.ChargerStatus = types.ChargerStatus_ACTIVE

	// update both entities
	k.SetEnergyTransferOffer(ctx, *offer)
	k.SetEnergyTransfer(ctx, energyTransferObj)

	k.EmitChangeOfferStatusEvent(ctx, energyTransferObj.GetEnergyTransferOfferId(), types.ChargerStatus_BUSY.String(), types.ChargerStatus_ACTIVE.String())
	k.EmitCompleteEnergyTransferEvent(ctx, energyTransferObj.Id, energyTransferObj.GetEnergyTransferOfferId(), usedServiceUnits)

	return nil
}

func (k Keeper) parseAddressAndSendTokensFromModule(ctx sdk.Context, targetAccountAddress string, collateralCoins sdk.Coins) error {
	target, err := sdk.AccAddressFromBech32(targetAccountAddress)
	if err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, target, collateralCoins)
}

func (k Keeper) CancelEnergyTransfer(ctx sdk.Context, energyTransferId uint64) error {
	k.Logger(ctx).Debug("CancelEnergyTransfer - ID=%d", energyTransferId)
	energyTransferObj, err := k.MustGetEnergyTransfer(ctx, energyTransferId)
	if err != nil {
		return err
	}

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
	amount := sdk.NewInt(int64(energyTransferObj.GetCollateral()))
	coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, amount))
	if err = k.parseAddressAndSendTokensFromModule(ctx, energyTransferObj.GetDriverAccountAddress(), coinsToTransfer); err != nil {
		return err
	}

	// update both entities
	k.SetEnergyTransferOffer(ctx, *offer)
	k.SetEnergyTransfer(ctx, *energyTransferObj)

	k.EmitChangeOfferStatusEvent(ctx, energyTransferObj.GetEnergyTransferOfferId(), types.ChargerStatus_BUSY.String(), types.ChargerStatus_ACTIVE.String())
	k.EmitCancelEnergyTransferEvent(ctx, energyTransferId, energyTransferObj.ChargerId, "")

	return nil
}

func (k Keeper) EmitChangeOfferStatusEvent(ctx sdk.Context, energyTransferOfferId uint64, oldStatus, newStatus string) {
	event := &types.ChangeOfferStatus{
		EnergyTransferOfferId: energyTransferOfferId,
		OldStatus:             oldStatus,
		NewStatus:             newStatus,
	}
	if err := ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("ChangeOfferStatus error", "event", event, "error", err.Error())
	}
}

func (k Keeper) EmitCancelEnergyTransferEvent(ctx sdk.Context, energyTransferId uint64, chargerId string, cancelReason string) {
	event := &types.CancelEnergyTransfer{
		EnergyTransferId: energyTransferId,
		ChargerId:        chargerId,
		CancelReason:     cancelReason,
	}
	if err := ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("CancelEnergyTransfer error", "event", event, "error", err.Error())
	}
}

func (k Keeper) EmitCompleteEnergyTransferEvent(ctx sdk.Context, energyTransferId uint64, energyTransferOfferId uint64, usedServiceUnits int32) {
	event := &types.CompleteEnergyTransfer{
		EnergyTransferId:      energyTransferId,
		EnergyTransferOfferId: energyTransferOfferId,
		EnergyTransferred:     usedServiceUnits,
	}
	if err := ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("CompleteEnergyTransfer error", "event", event, "error", err.Error())
	}
}

func (k Keeper) EmitEnergyTransferCreatedEvent(ctx sdk.Context, energyTransferId uint64, chargerId string, energyAmountToTransfer int32) {
	event := &types.EnergyTransferCreatedEvent{
		EnergyTransferId:       energyTransferId,
		ChargerId:              chargerId,
		EnergyAmountToTransfer: energyAmountToTransfer,
	}
	if err := ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("EnergyTransferCreatedEvent error", "event", event, "error", err.Error())
	}
}

func (k Keeper) EmitBeginEnergyTransferEvent(ctx sdk.Context, energyTransferId uint64, energyTransferOfferId uint64) {
	event := &types.BeginEnergyTransfer{
		EnergyTransferId:      energyTransferId,
		EnergyTransferOfferId: energyTransferOfferId,
	}
	if err := ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("BeginEnergyTransfer event error", "event", event, "error", err.Error())
	}
}
