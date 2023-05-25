package keeper_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"cosmossdk.io/math"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

/*
These tests are constructed as a mix of unit tests and integration tests
as described at https://docs.cosmos.network/main/building-modules/testing#integration-tests.

Thus integration tests within single module.
These tests interact with the tested module via the defined Msg and Query services.
*/
func createExampleTestObjects() (types.EnergyTransferOffer, types.EnergyTransfer) {

	location := types.Location{Latitude: "34.4", Longitude: "5.2"}

	var energyTransferOffer = types.EnergyTransferOffer{
		Owner:         "c4e1k4quu6r2jl0afrn7m5h6ta7707r4lm4cxktxq5",
		ChargerId:     "EVGC011221122GK0122",
		ChargerStatus: types.ChargerStatus_ACTIVE,
		Location:      &location,
		Tariff:        56,
		Name:          "Test charging point",
		PlugType:      types.PlugType_Type2,
	}

	var energyTransfer = types.EnergyTransfer{
		OwnerAccountAddress:   "c4e1k4quu6r2jl0afrn7m5h6ta7707r4lm4cxktxq5",
		DriverAccountAddress:  "c4e1z793mr75a8v7604w55rvzwfjakz4ec9vwk9z86",
		EnergyTransferOfferId: 0,
		ChargerId:             "EVGC011221122GK0122",
		Status:                types.TransferStatus_REQUESTED,
		OfferedTariff:         56,
		EnergyToTransfer:      22,
		Collateral:            1232,
	}

	return energyTransferOffer, energyTransfer
}

func setupTestMsgServer(t testing.TB) (types.MsgServer, context.Context, keeper.Keeper) {
	k, ctx := keepertest.CfeevKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx), *k
}

func TestEnergyTransferCancel(t *testing.T) {
	msgServer, goCtx, k := setupTestMsgServer(t)
	offer, transfer := createExampleTestObjects()
	ctx := sdk.UnwrapSDKContext(goCtx)

	bankKeeper := k.GetBankKeeper()

	// Populate the EV driver account with some coins
	coins := sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(5000)))
	driver, err := sdk.AccAddressFromBech32(transfer.DriverAccountAddress)
	err = bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, driver, coins)

	if err != nil {
		panic(err)
	}

	msg := &types.MsgPublishEnergyTransferOffer{
		Creator:   offer.Owner,
		ChargerId: offer.ChargerId,
		Tariff:    offer.Tariff,
		Location:  offer.Location,
		Name:      offer.Name,
		PlugType:  offer.PlugType,
	}

	result, err := msgServer.PublishEnergyTransferOffer(goCtx, msg)

	if err != nil {
		panic(err)
	}

	if result == nil {
		panic(fmt.Errorf("unexpected nil result"))
	}

	foundOffer, found := k.GetEnergyTransferOffer(ctx, result.Id)
	require.True(t, found)

	collateral := sdk.Coin{Denom: "uc4e", Amount: math.NewInt(int64(transfer.Collateral))}

	msgStartTransfer := &types.MsgStartEnergyTransferRequest{
		Creator:               transfer.DriverAccountAddress,
		EnergyTransferOfferId: foundOffer.Id,
		ChargerId:             transfer.ChargerId,
		OwnerAccountAddress:   transfer.OwnerAccountAddress,
		OfferedTariff:         strconv.Itoa(int(transfer.OfferedTariff)),
		EnergyToTransfer:      transfer.EnergyToTransfer,
		Collateral:            &collateral,
	}

	startTransferResponse, err := msgServer.StartEnergyTransferRequest(goCtx, msgStartTransfer)

	if err != nil {
		panic(err)
	}

	if startTransferResponse == nil {
		panic(fmt.Errorf("unexpected nil result"))
	}

	energyTransferId := startTransferResponse.GetId()

	foundTransfer, found := k.GetEnergyTransfer(ctx, energyTransferId)
	require.True(t, found)

	foundOffer, found = k.GetEnergyTransferOffer(ctx, foundTransfer.EnergyTransferOfferId)
	require.True(t, found)

	require.Equal(t, types.ChargerStatus_BUSY, foundOffer.ChargerStatus)
	require.Equal(t, types.TransferStatus_REQUESTED, foundTransfer.Status)

	// Energy transfer can be canceled when it has REQUESTED status so before the actual transfer begins
	// A similar approach to the SIP
	// Cancel energy transfer

	msgCancelTransfer := &types.MsgCancelEnergyTransferRequest{
		Creator:          transfer.OwnerAccountAddress,
		EnergyTransferId: foundTransfer.Id,
		ChargerId:        foundTransfer.ChargerId,
		ErrorInfo:        "Test_cancel",
	}

	_, err = msgServer.CancelEnergyTransferRequest(goCtx, msgCancelTransfer)

	if err != nil {
		panic(err)
	}

	foundTransfer, found = k.GetEnergyTransfer(ctx, energyTransferId)
	require.True(t, found)

	foundOffer, found = k.GetEnergyTransferOffer(ctx, foundTransfer.EnergyTransferOfferId)
	require.True(t, found)

	require.Equal(t, types.TransferStatus_CANCELLED, foundTransfer.Status)
	require.Equal(t, types.ChargerStatus_ACTIVE, foundOffer.ChargerStatus)
}
