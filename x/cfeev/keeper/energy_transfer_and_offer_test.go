package keeper_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"cosmossdk.io/math"
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

	coins := sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(5000)))
	bankutils.AddCoinsToAccount(coins, sdk.AccAddress(transfer.GetDriverAccountAddress()))
	bankutils.VerifyAccountBalanceByDenom(sdk.AccAddress(transfer.GetDriverAccountAddress()), "uc4e", sdk.NewInt(5000))

	fmt.Println("Driver: " + transfer.GetDriverAccountAddress())

	msgPublishEnergyTransferOffer := &types.MsgPublishEnergyTransferOffer{
		Creator:   offer.Owner,
		ChargerId: offer.ChargerId,
		Tariff:    offer.Tariff,
		Location:  offer.Location,
		Name:      offer.Name,
		PlugType:  offer.PlugType,
	}

	msgServer := keeper.NewMsgServerImpl(testHelper.App.CfeevKeeper)

	newOfferId, err := msgServer.PublishEnergyTransferOffer(testHelper.WrappedContext, msgPublishEnergyTransferOffer)
	require.NoError(t, err)

	testHelper.C4eEvUtils.VerifyGetEnergyTransferOffer(testHelper.GetContext(), newOfferId.GetId())

	collateral := sdk.Coin{Denom: "uc4e", Amount: math.NewInt(int64(transfer.Collateral))}

	msgStartTransfer := &types.MsgStartEnergyTransferRequest{
		Creator:               transfer.DriverAccountAddress,
		EnergyTransferOfferId: newOfferId.GetId(),
		ChargerId:             transfer.ChargerId,
		OwnerAccountAddress:   transfer.OwnerAccountAddress,
		OfferedTariff:         strconv.Itoa(int(transfer.OfferedTariff)),
		EnergyToTransfer:      transfer.EnergyToTransfer,
		Collateral:            &collateral,
	}

	startTransferResponse, err := msgServer.StartEnergyTransferRequest(testHelper.WrappedContext, msgStartTransfer)

	if err != nil {
		panic(err)
	}
	if startTransferResponse == nil {
		panic(fmt.Errorf("unexpected nil result"))
	}

	energyTransferId := startTransferResponse.GetId()
	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId.GetId(), types.ChargerStatus_BUSY)
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

	_, err = msgServer.CancelEnergyTransferRequest(testHelper.WrappedContext, msgCancelTransfer)
	if err != nil {
		panic(err)
	}
	testHelper.C4eEvUtils.VerifyEnergyTransferOfferStatus(testHelper.GetContext(), newOfferId.GetId(), types.ChargerStatus_ACTIVE)
	testHelper.C4eEvUtils.VerifyEnergyTransferStatus(testHelper.GetContext(), energyTransferId, types.TransferStatus_CANCELLED)
}
