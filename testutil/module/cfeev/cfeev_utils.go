package cfeevutils

import (
	"cosmossdk.io/math"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	cfeevmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type C4eEvKeeperUtils struct {
	t                 require.TestingT
	helperCfeevKeeper *cfeevmodulekeeper.Keeper
}

func NewC4eEvKeeperUtils(t require.TestingT, helperCfeevKeeper *cfeevmodulekeeper.Keeper) C4eEvKeeperUtils {
	return C4eEvKeeperUtils{t: t, helperCfeevKeeper: helperCfeevKeeper}
}

type C4eEvUtils struct {
	C4eEvKeeperUtils
	BankUtils *testcosmos.BankUtils
}

func NewC4eEvUtils(t require.TestingT,
	helpeCfeevmodulekeeper cfeevmodulekeeper.Keeper,
	bankUtils *testcosmos.BankUtils) C4eEvUtils {
	return C4eEvUtils{C4eEvKeeperUtils: NewC4eEvKeeperUtils(t, &helpeCfeevmodulekeeper), BankUtils: bankUtils}
}

type ContextC4eEvUtils struct {
	C4eEvUtils
	testContext testenv.TestContext
}

func NewContextC4eClaimUtils(t require.TestingT,
	testContext testenv.TestContext,
	helpeCfeevmodulekeeper *cfeevmodulekeeper.Keeper,
	bankUtils *testcosmos.BankUtils) *ContextC4eEvUtils {

	c4eEvUtils := NewC4eEvUtils(t, *helpeCfeevmodulekeeper, bankUtils)
	return &ContextC4eEvUtils{C4eEvUtils: c4eEvUtils, testContext: testContext}
}

func (h *C4eEvUtils) VerifyGetEnergyTransferOffer(ctx sdk.Context, id uint64) {
	_, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransferOffer(ctx, id)
	require.True(h.t, found)
}

func (h *C4eEvUtils) VerifyGetEnergyTransfer(ctx sdk.Context, id uint64) {
	_, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, id)
	require.True(h.t, found)
}

func (h *C4eEvUtils) VerifyEnergyTransferStatus(ctx sdk.Context, id uint64, status types.TransferStatus) {
	et, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, id)
	require.True(h.t, found)
	require.Equal(h.t, status, et.Status)
}

func (h *C4eEvUtils) VerifyEnergyTransferOfferStatus(ctx sdk.Context, id uint64, status types.ChargerStatus) {
	eto, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransferOffer(ctx, id)
	require.True(h.t, found)
	require.Equal(h.t, status, eto.GetChargerStatus())
}

func (h *C4eEvUtils) CreateExampleTestEVObjects() (types.EnergyTransferOffer, types.EnergyTransfer) {

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
		Collateral:            math.NewInt(1232),
	}

	return energyTransferOffer, energyTransfer
}

func (h *C4eEvUtils) PublishAndVerifyEnergyTransferOffer(ctx sdk.Context, offer types.EnergyTransferOffer) uint64 {

	msgPublishEnergyTransferOffer := &types.MsgPublishEnergyTransferOffer{
		Creator:   offer.Owner,
		ChargerId: offer.ChargerId,
		Tariff:    offer.Tariff,
		Location:  offer.Location,
		Name:      offer.Name,
		PlugType:  offer.PlugType,
	}

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)

	newOfferId, err := msgServer.PublishEnergyTransferOffer(ctx, msgPublishEnergyTransferOffer)
	require.NoError(h.t, err)

	h.VerifyGetEnergyTransferOffer(ctx, newOfferId.GetId())

	return newOfferId.GetId()
}

func (h *C4eEvUtils) StartEnergyTransfer(ctx sdk.Context, transfer types.EnergyTransfer, newOfferId uint64) uint64 {
	msgStartTransfer := &types.MsgStartEnergyTransfer{
		Creator:               transfer.DriverAccountAddress,
		EnergyTransferOfferId: newOfferId,
		ChargerId:             transfer.ChargerId,
		OwnerAccountAddress:   transfer.OwnerAccountAddress,
		OfferedTariff:         transfer.OfferedTariff,
		EnergyToTransfer:      transfer.EnergyToTransfer,
		Collateral:            &transfer.Collateral,
	}

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)

	startTransferResponse, err := msgServer.StartEnergyTransfer(ctx, msgStartTransfer)
	require.NoError(h.t, err)

	return startTransferResponse.GetId()
}

func (h *C4eEvUtils) EnergyTransferCompleted(ctx sdk.Context, energyTransferId uint64, usedServiceUnits uint64) {
	msg := &types.MsgEnergyTransferCompleted{EnergyTransferId: energyTransferId, UsedServiceUnits: usedServiceUnits}

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	_, err := msgServer.EnergyTransferCompleted(ctx, msg)
	require.NoError(h.t, err)
}

func (h *C4eEvUtils) EnergyTransferStarted(ctx sdk.Context, energyTransferId uint64) {
	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	msgConfirmTransferRequestStarted := &types.MsgEnergyTransferStarted{EnergyTransferId: energyTransferId}
	_, err := msgServer.EnergyTransferStarted(ctx, msgConfirmTransferRequestStarted)
	require.NoError(h.t, err)
}
