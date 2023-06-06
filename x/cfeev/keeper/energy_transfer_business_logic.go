package keeper

import (
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	k.Logger(ctx).Debug("start energy transfer - driver: %s, charger ID: %s", driver, chargerId)

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

	// get energy transfer offer object by offer id
	offer, found := k.GetEnergyTransferOffer(ctx, energyTransferObj.EnergyTransferOfferId)

	if !found {
		return 0, status.Error(codes.NotFound, "energy transfer offer not found")
	}

	if offer.GetChargerStatus() == types.ChargerStatus_ACTIVE {
		offer.ChargerStatus = types.ChargerStatus_BUSY
	} else {
		return 0, sdkerrors.Wrap(types.ErrBusyCharger, "busy charger")
	}

	// check if the offered tariff has not been changed
	if !(offer.Tariff == energyTransferObj.OfferedTariff) {
		return 0, status.Error(codes.InvalidArgument, "wrong tariff")
	}

	k.Logger(ctx).Debug("start energy transfer - send collateral, driver: %s", energyTransferObj.DriverAccountAddress)
	fmt.Println(energyTransferObj.DriverAccountAddress)

	// send tokens to the escrow account
	coinsToTransfer := collateral.Amount.String() + collateral.Denom
	err := k.sendCollateralToEscrowAccount(ctx, energyTransferObj.DriverAccountAddress, coinsToTransfer)
	if err != nil {
		return 0, err
	}

	energyTransferId := k.AppendEnergyTransfer(ctx, energyTransferObj)

	// update the offer in the store
	k.SetEnergyTransferOffer(ctx, offer)

	// send notification event to connector, the event will emitted only if there is no previous errors
	event := &types.EnergyTransferCreatedEvent{
		EnergyTransferId:       energyTransferId,
		ChargerId:              chargerId,
		EnergyAmountToTransfer: energyTransferObj.GetEnergyToTransfer(),
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("new EnergyTransferCreated emit event error", "event", event, "error", err.Error())
	}

	offerEvent, err := k.EmitChangeOfferStatusEvent(ctx, energyTransferObj.GetEnergyTransferOfferId(), types.ChargerStatus_ACTIVE.String(), types.ChargerStatus_BUSY.String())
	if err != nil {
		k.Logger(ctx).Error("new EmitChangeOfferStatusEvent error", "event", offerEvent, "error", err.Error())
	}

	return energyTransferId, nil
}

// sends the collateral from the driver's account to a module account
func (k Keeper) sendCollateralToEscrowAccount(ctx sdk.Context, driverAccountAddress string, collateral string) error {
	driver, err := sdk.AccAddressFromBech32(driverAccountAddress)
	if err != nil {
		panic(err)
	}
	collateralCoins, err := sdk.ParseCoinsNormalized(collateral)
	if err != nil {
		panic(err)
	}
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, driver, types.ModuleName, collateralCoins)
	if sdkError != nil {
		return sdkError
	}

	return nil
}

func (k Keeper) EnergyTransferStarted(ctx sdk.Context, energyTransferId uint64) error {

	energyTransfer, found := k.GetEnergyTransfer(ctx, energyTransferId)
	if !found {
		return sdkerrors.Wrap(types.ErrEnergyTransferNotFound, "energy transfer not found")
	}

	if energyTransfer.Status != types.TransferStatus_REQUESTED {
		return sdkerrors.Wrap(types.ErrWrongEnergyTransferStatus, energyTransfer.Status.String())
	}

	// REQUESTED ==> ONGOING
	energyTransfer.Status = types.TransferStatus_ONGOING

	// update energyTransfer instance in the KVStore
	k.SetEnergyTransfer(ctx, energyTransfer)

	event := &types.BeginEnergyTransfer{
		EnergyTransferId:      energyTransferId,
		EnergyTransferOfferId: energyTransfer.GetEnergyTransferOfferId(),
	}
	err := ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("new begin energy transfer emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) EnergyTransferCompleted(ctx sdk.Context, energyTransferId uint64, usedServiceUnits int32) error {
	energyTransferObj, found := k.GetEnergyTransfer(ctx, energyTransferId)
	if !found {
		return sdkerrors.Wrap(types.ErrEnergyTransferNotFound, "energy transfer not found")
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
		err = k.sendTokensToTargetAccount(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer)
		energyTransferObj.Status = types.TransferStatus_PAID

	} else if energyTransferObj.EnergyToTransfer > usedServiceUnits {
		// calculate used tokens
		usedTokens := energyTransferObj.OfferedTariff * usedServiceUnits
		amount := sdk.NewInt(int64(usedTokens))
		coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, amount))
		err = k.sendTokensToTargetAccount(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer)

		// calculate unused tokens
		unusedTokens := energyTransferObj.Collateral - uint64(usedTokens)
		amount = sdk.NewInt(int64(unusedTokens))
		coinsToTransfer = sdk.NewCoins(sdk.NewCoin(denom, amount))
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
			amount := sdk.NewInt(int64(energyTransferObj.GetCollateral()))
			coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, amount))
			err = k.sendTokensToTargetAccount(ctx, energyTransferObj.OwnerAccountAddress, coinsToTransfer)
			energyTransferObj.Status = types.TransferStatus_PAID
		}
		// TODO:
	}

	if err != nil {
		return sdkerrors.Wrap(types.ErrCoinTransferFailed, "coin transfer failed")
	}

	// get energy transfer offer object by offer id
	offer, found := k.GetEnergyTransferOffer(ctx, energyTransferObj.EnergyTransferOfferId)
	if !found {
		return sdkerrors.Wrap(types.ErrEnergyTransferOfferNotFound, "energy transfer offer not found")
	}
	offer.ChargerStatus = types.ChargerStatus_ACTIVE

	// update both entities
	k.SetEnergyTransferOffer(ctx, offer)
	k.SetEnergyTransfer(ctx, energyTransferObj)

	event, err := k.EmitChangeOfferStatusEvent(ctx, energyTransferObj.GetEnergyTransferOfferId(), types.ChargerStatus_BUSY.String(), types.ChargerStatus_ACTIVE.String())
	if err != nil {
		k.Logger(ctx).Error("new EmitChangeOfferStatusEvent error", "event", event, "error", err.Error())
	}

	completeEvent := &types.CompleteEnergyTransfer{
		EnergyTransferId:      energyTransferObj.Id,
		EnergyTransferOfferId: energyTransferObj.GetEnergyTransferOfferId(),
		EnergyTransferred:     usedServiceUnits,
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("new complete energy transfer emit event error", "event", completeEvent, "error", err.Error())
	}

	return nil
}

func (k Keeper) sendTokensToTargetAccount(ctx sdk.Context, targetAccountAddress string, collateralCoins sdk.Coins) error {
	target, err := sdk.AccAddressFromBech32(targetAccountAddress)
	if err != nil {
		panic(err)
	}
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, target, collateralCoins)

	if sdkError != nil {
		return sdkError
	}

	return nil
}

func (k Keeper) CancelEnergyTransfer(ctx sdk.Context, energyTransferId uint64) error {
	k.Logger(ctx).Debug("CancelEnergyTransfer - ID=%d", energyTransferId)
	energyTransferObj, found := k.GetEnergyTransfer(ctx, energyTransferId)
	if !found {
		return sdkerrors.Wrap(types.ErrEnergyTransferNotFound, "energy transfer not found")
	}

	// if energyTransferObj.GetStatus() != types.TransferStatus_REQUESTED {
	// 	return nil, sdkerrors.Wrap(types.ErrWrongEnergyTransferStatus, energyTransferObj.GetStatus().String())
	// }

	energyTransferObj.Status = types.TransferStatus_CANCELLED

	// get energy transfer offer object by offer id
	offer, found := k.GetEnergyTransferOffer(ctx, energyTransferObj.EnergyTransferOfferId)
	if !found {
		return sdkerrors.Wrap(types.ErrEnergyTransferOfferNotFound, "energy transfer offer not found")
	}
	offer.ChargerStatus = types.ChargerStatus_ACTIVE

	// send the collateral back to the EV driver's account
	denom := k.Denom(ctx)
	amount := sdk.NewInt(int64(energyTransferObj.GetCollateral()))
	coinsToTransfer := sdk.NewCoins(sdk.NewCoin(denom, amount))
	err := k.sendTokensToTargetAccount(ctx, energyTransferObj.GetDriverAccountAddress(), coinsToTransfer)

	if err != nil {
		return sdkerrors.Wrap(types.ErrCoinTransferFailed, "coin transfer failed")
	}

	// update both entities
	k.SetEnergyTransferOffer(ctx, offer)
	k.SetEnergyTransfer(ctx, energyTransferObj)

	event, err := k.EmitChangeOfferStatusEvent(ctx, energyTransferObj.GetEnergyTransferOfferId(), types.ChargerStatus_BUSY.String(), types.ChargerStatus_ACTIVE.String())
	if err != nil {
		k.Logger(ctx).Error("new EmitChangeOfferStatusEvent error", "event", event, "error", err.Error())
	}

	cancelEvent, err := k.EmitCancelEnergyTransferEvent(ctx, energyTransferId, energyTransferObj.ChargerId, "")
	if err != nil {
		k.Logger(ctx).Error("new cancel energy transfer error", "event", cancelEvent, "error", err.Error())
	}

	return nil
}

func (k Keeper) EmitChangeOfferStatusEvent(ctx sdk.Context, energyTransferOfferId uint64, oldStatus, newStatus string) (*types.ChangeOfferStatus, error) {
	event := &types.ChangeOfferStatus{
		EnergyTransferOfferId: energyTransferOfferId,
		OldStatus:             oldStatus,
		NewStatus:             newStatus,
	}
	err := ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (k Keeper) EmitCancelEnergyTransferEvent(ctx sdk.Context, energyTransferId uint64, chargerId string, cancelReason string) (*types.CancelEnergyTransfer, error) {
	event := &types.CancelEnergyTransfer{
		EnergyTransferId: energyTransferId,
		ChargerId:        chargerId,
		CancelReason:     cancelReason,
	}
	err := ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		return event, err
	}
	return event, nil

}
