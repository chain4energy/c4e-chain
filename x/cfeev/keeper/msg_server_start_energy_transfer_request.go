package keeper

import (
	"context"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) StartEnergyTransferRequest(goCtx context.Context, msg *types.MsgStartEnergyTransferRequest) (*types.MsgStartEnergyTransferRequestResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// handling tariffFromMsg string from message
	tariffFromMsg, err := strconv.ParseInt(msg.GetOfferedTariff(), 10, 32)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request, incorrectly defined tariff")
	}
	offeredTariff := int32(tariffFromMsg)

	// handle dynamic collateral
	dynamicCollateral := msg.GetCollateral()
	collateralAmount := dynamicCollateral.Amount

	var energyTransferObj = types.EnergyTransfer{
		OwnerAccountAddress:   msg.OwnerAccountAddress,
		DriverAccountAddress:  msg.Creator,
		EnergyTransferOfferId: msg.EnergyTransferOfferId,
		ChargerId:             msg.ChargerId,
		Status:                types.TransferStatus_REQUESTED,
		OfferedTariff:         offeredTariff,
		EnergyToTransfer:      msg.GetEnergyToTransfer(),
		Collateral:            collateralAmount.Uint64(),
	}

	// get energy transfer offer object by offer id
	offer, found := k.GetEnergyTransferOffer(ctx, energyTransferObj.EnergyTransferOfferId)

	if !found {
		return nil, status.Error(codes.NotFound, "energy transfer offer not found")
	}

	if offer.GetChargerStatus() == types.ChargerStatus_ACTIVE {
		offer.ChargerStatus = types.ChargerStatus_BUSY
	} else {
		return nil, sdkerrors.Wrap(types.ErrBusyCharger, "busy charger")
	}

	// check if the offered tariff has not been changed
	if !(offer.Tariff == energyTransferObj.OfferedTariff) {
		return nil, status.Error(codes.InvalidArgument, "wrong tariff")
	}

	// send tokens to the escrow account
	coinsToTransfer := msg.GetCollateral().Amount.String() + msg.GetCollateral().Denom
	err = k.sendCollateralToEscrowAccount(ctx, energyTransferObj.DriverAccountAddress, coinsToTransfer)
	if err != nil {
		return nil, status.Error(codes.Aborted, "sending funds to the escrow aborted")
	}

	energyTransferId := k.AppendEnergyTransfer(ctx, energyTransferObj)
	_ = energyTransferId

	// update the offer in the store
	k.SetEnergyTransferOffer(ctx, offer)

	// send notification event to connector, the event will emitted only if there is no previous errors
	event := &types.EnergyTransferCreatedEvent{
		EnergyTransferId:       energyTransferId,
		ChargerId:              msg.ChargerId,
		EnergyAmountToTransfer: energyTransferObj.GetEnergyToTransfer(),
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("new EnergyTransferCreated emit event error", "event", event, "error", err.Error())
	}

	return &types.MsgStartEnergyTransferRequestResponse{Id: energyTransferId}, nil
}

// sends the collateral from the driver's account to a module account
func (k msgServer) sendCollateralToEscrowAccount(ctx sdk.Context, driverAccountAddress string, collateral string) error {
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
