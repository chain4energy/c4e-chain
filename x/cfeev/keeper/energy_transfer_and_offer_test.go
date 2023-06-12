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
)

func setupTestMsgServer(t testing.TB) (types.MsgServer, context.Context, keeper.Keeper) {
	k, ctx, _ := keepertest.CfeevKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx), *k
}

func TestEnergyTransferCancel(t *testing.T) {
	testHelper := app.SetupTestApp(t)

	offer, transfer := testHelper.C4eEvUtils.CreateExampleTestEVObjects()
	bankutils := testHelper.BankUtils

	bankutils.AddDefaultDenomCoinsToAccount(sdk.NewInt(5000), sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()))
	bankutils.VerifyAccountBalanceByDenom(sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()), "uc4e", sdk.NewInt(5000))

	fmt.Println("EV Driver acc address: " + transfer.GetDriverAccountAddress())

	newOfferId := testHelper.C4eEvUtils.PublishAndVerifyEnergyTransferOffer(testHelper.GetContext(), offer)
	energyTransferId := testHelper.C4eEvUtils.StartEnergyTransfer(testHelper.GetContext(), transfer, newOfferId)

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_BUSY)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_REQUESTED)

	// Energy transfer can be canceled when it has REQUESTED status so before the actual transfer begins
	// A similar approach to the SIP
	// Cancel energy transfer
	msgCancelTransfer := &types.MsgCancelEnergyTransfer{
		Creator:          transfer.OwnerAccountAddress,
		EnergyTransferId: energyTransferId,
		ChargerId:        offer.GetChargerId(),
		ErrorInfo:        "Test_cancel",
	}

	msgServer := keeper.NewMsgServerImpl(testHelper.App.CfeevKeeper)
	_, err := msgServer.CancelEnergyTransfer(testHelper.WrappedContext, msgCancelTransfer)
	if err != nil {
		panic(err)
	}
	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_ACTIVE)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_CANCELLED)
}

func TestEnergyTransferStartedFull(t *testing.T) {
	testHelper := app.SetupTestApp(t)

	offer, transfer := testHelper.C4eEvUtils.CreateExampleTestEVObjects()
	bankutils := testHelper.BankUtils

	bankutils.AddDefaultDenomCoinsToAccount(sdk.NewInt(5000), sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()))
	bankutils.VerifyAccountBalanceByDenom(sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()), "uc4e", sdk.NewInt(5000))

	newOfferId := testHelper.C4eEvUtils.PublishAndVerifyEnergyTransferOffer(testHelper.GetContext(), offer)
	energyTransferId := testHelper.C4eEvUtils.StartEnergyTransfer(testHelper.GetContext(), transfer, newOfferId)

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_BUSY)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_REQUESTED)

	// confirm that energy transfer has been started
	testHelper.C4eEvUtils.EnergyTransferStarted(testHelper.GetContext(), energyTransferId)

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_BUSY)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_ONGOING)

	// when energy transfer is completed - full charging
	testHelper.C4eEvUtils.EnergyTransferCompleted(testHelper.GetContext(), energyTransferId, transfer.GetEnergyToTransfer())

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_ACTIVE)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_PAID)

	bankutils.VerifyAccountDefaultDenomBalance(testHelper.GetContext(), sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()), sdk.NewInt(3768))
	bankutils.VerifyAccountDefaultDenomBalance(testHelper.GetContext(), sdk.MustAccAddressFromBech32(transfer.GetOwnerAccountAddress()), sdk.NewInt(1232))
}

func TestEnergyTransferStartedPartial(t *testing.T) {
	testHelper := app.SetupTestApp(t)

	offer, transfer := testHelper.C4eEvUtils.CreateExampleTestEVObjects()
	bankutils := testHelper.BankUtils

	bankutils.AddDefaultDenomCoinsToAccount(sdk.NewInt(5000), sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()))
	bankutils.VerifyAccountBalanceByDenom(sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()), "uc4e", sdk.NewInt(5000))

	newOfferId := testHelper.C4eEvUtils.PublishAndVerifyEnergyTransferOffer(testHelper.GetContext(), offer)
	energyTransferId := testHelper.C4eEvUtils.StartEnergyTransfer(testHelper.GetContext(), transfer, newOfferId)

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_BUSY)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_REQUESTED)

	// confirm that energy transfer has been started
	testHelper.C4eEvUtils.EnergyTransferStarted(testHelper.GetContext(), energyTransferId)

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_BUSY)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_ONGOING)

	// when energy transfer is completed - partial charging
	testHelper.C4eEvUtils.EnergyTransferCompleted(testHelper.GetContext(), energyTransferId, 10)

	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId, types.ChargerStatus_ACTIVE)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_PAID)

	bankutils.VerifyAccountDefaultDenomBalance(testHelper.GetContext(), sdk.MustAccAddressFromBech32(transfer.GetDriverAccountAddress()), sdk.NewInt(4440))
	bankutils.VerifyAccountDefaultDenomBalance(testHelper.GetContext(), sdk.MustAccAddressFromBech32(transfer.GetOwnerAccountAddress()), sdk.NewInt(560))
}
