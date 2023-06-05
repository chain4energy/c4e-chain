package keeper

import (
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StartEnergyTransferRequest(
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
