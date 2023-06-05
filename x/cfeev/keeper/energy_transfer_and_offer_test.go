package keeper_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/app"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupTestMsgServer(t testing.TB) (types.MsgServer, context.Context, keeper.Keeper) {
	k, ctx, _ := keepertest.CfeevKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx), *k
}

func TestEnergyTransferCancel(t *testing.T) {
	testHelper := app.SetupTestApp(t)

	offer, transfer := testHelper.C4eEvUtils.CreateExampleTestEVObjects()
	bankutils := testHelper.BankUtils

	// Populate the EV driver account with some coins
	// coin := sdk.NewCoin("uc4e", sdk.NewInt(5000))
	// bankutils.AddCoinsToModule(coin, types.ModuleName)
	// fmt.Println(bankutils.GetModuleAccountBalanceByDenom(testHelper.GetContext(), types.ModuleName, "uc4e"))

	bankutils.AddDefaultDenomCoinsToAccount(sdk.NewInt(5000), sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()))
	bankutils.VerifyAccountBalanceByDenom(sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()), "uc4e", sdk.NewInt(5000))

	fmt.Println("Driver: " + transfer.GetDriverAccountAddress())

	newOfferId := testHelper.C4eEvUtils.PublishAndVerifyEnergyTransferOffer(testHelper.GetContext(), offer)
	energyTransferId := testHelper.C4eEvUtils.StartEnergyTransferRequest(testHelper.GetContext(), transfer, newOfferId)

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_BUSY)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_REQUESTED)

	// Energy transfer can be canceled when it has REQUESTED status so before the actual transfer begins
	// A similar approach to the SIP
	// Cancel energy transfer
	msgCancelTransfer := &types.MsgCancelEnergyTransferRequest{
		Creator:          transfer.OwnerAccountAddress,
		EnergyTransferId: energyTransferId,
		ChargerId:        offer.GetChargerId(),
		ErrorInfo:        "Test_cancel",
	}

	msgServer := keeper.NewMsgServerImpl(testHelper.App.CfeevKeeper)
	_, err := msgServer.CancelEnergyTransferRequest(testHelper.WrappedContext, msgCancelTransfer)
	if err != nil {
		panic(err)
	}
	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_ACTIVE)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_CANCELLED)
}

func TestEnergyTransferStarted(t *testing.T) {
	testHelper := app.SetupTestApp(t)

	offer, transfer := testHelper.C4eEvUtils.CreateExampleTestEVObjects()
	bankutils := testHelper.BankUtils

	bankutils.AddDefaultDenomCoinsToAccount(sdk.NewInt(5000), sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()))
	bankutils.VerifyAccountBalanceByDenom(sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()), "uc4e", sdk.NewInt(5000))

	fmt.Println("Driver: " + transfer.GetDriverAccountAddress())

	newOfferId := testHelper.C4eEvUtils.PublishAndVerifyEnergyTransferOffer(testHelper.GetContext(), offer)
	energyTransferId := testHelper.C4eEvUtils.StartEnergyTransferRequest(testHelper.GetContext(), transfer, newOfferId)

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_BUSY)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_REQUESTED)

	// confirm that energy transfer has been started
	msgServer := keeper.NewMsgServerImpl(testHelper.App.CfeevKeeper)
	msgCancelTransferRequest := &types.MsgEnergyTransferStartedRequest{EnergyTransferId: energyTransferId}
	_, err := msgServer.EnergyTransferStartedRequest(testHelper.WrappedContext, msgCancelTransferRequest)
	require.NoError(t, err)

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_BUSY)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_ONGOING)

	testHelper.C4eEvUtils.EnergyTransferCompletedRequest(testHelper.GetContext(), energyTransferId, 22)

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_ACTIVE)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_PAID)
}
