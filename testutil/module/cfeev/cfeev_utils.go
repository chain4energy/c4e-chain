package cfeevutils

import (
	"cosmossdk.io/math"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/testutil/sample"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	"github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	cfeevmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"math/rand"
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

func (h *C4eEvUtils) VerifyEnergyTransferOffer(ctx sdk.Context, id uint64, expectedOffer types.EnergyTransferOffer) {
	offer, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransferOffer(ctx, id)
	require.True(h.t, found)
	require.Equal(h.t, expectedOffer, offer)
}

func (h *C4eEvUtils) VerifyEnergyTransfer(ctx sdk.Context, id uint64, expectedTransfer types.EnergyTransfer) {
	energyTransfer, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, id)
	require.True(h.t, found)
	require.Equal(h.t, expectedTransfer, energyTransfer)
}

func (h *C4eEvUtils) VerifyGetEnergyTransfer(ctx sdk.Context, id uint64) {
	_, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, id)
	require.True(h.t, found)
}

func (h *C4eEvUtils) VerifyEnergyTransferstatus(ctx sdk.Context, id uint64, status types.TransferStatus) {
	et, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, id)
	require.True(h.t, found)
	require.Equal(h.t, status, et.Status)
}

func (h *C4eEvUtils) VerifyEnergyTransferOfferStatus(ctx sdk.Context, id uint64, status types.ChargerStatus) {
	energyTransferOffer, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransferOffer(ctx, id)
	require.True(h.t, found)
	require.Equal(h.t, status, energyTransferOffer.GetChargerStatus())
}

func (h *C4eEvUtils) CreateExampleTestEVObjects() (types.EnergyTransferOffer, types.EnergyTransfer) {
	r := utils.RandomSource()
	location := RandomLocation(r)

	var energyTransferOffer = types.EnergyTransferOffer{
		Owner:         sample.AccAddress(),
		ChargerId:     utils.RandStringOfLength(15),
		ChargerStatus: types.ChargerStatus_ACTIVE,
		Location:      location,
		Tariff:        56,
		Name:          utils.RandStringOfLength(10),
		PlugType:      types.PlugType_Type2,
	}

	var energyTransfer = types.EnergyTransfer{
		Owner:                 sample.AccAddress(),
		Driver:                sample.AccAddress(),
		EnergyTransferOfferId: 0,
		ChargerId:             utils.RandStringOfLength(10),
		Status:                types.TransferStatus_REQUESTED,
		OfferedTariff:         56,
		EnergyToTransfer:      22,
		Collateral:            math.NewInt(1232),
	}

	return energyTransferOffer, energyTransfer
}

func (h *C4eEvUtils) PublishEnergyTransferOffer(ctx sdk.Context, creator string, offer types.EnergyTransferOffer) uint64 {
	msgPublishEnergyTransferOffer := &types.MsgPublishEnergyTransferOffer{
		Creator:   creator,
		ChargerId: offer.ChargerId,
		Tariff:    offer.Tariff,
		Location:  offer.Location,
		Name:      offer.Name,
		PlugType:  offer.PlugType,
	}
	offer.Owner = creator
	offer.ChargerStatus = types.ChargerStatus_ACTIVE
	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	energyOffersCountBefore := h.helperCfeevKeeper.GetEnergyTransferOfferCount(ctx)
	res, err := msgServer.PublishEnergyTransferOffer(ctx, msgPublishEnergyTransferOffer)
	require.NoError(h.t, err)
	offerId := res.GetId()
	energyOffersCountAfter := h.helperCfeevKeeper.GetEnergyTransferOfferCount(ctx)
	require.Equal(h.t, energyOffersCountBefore+1, energyOffersCountAfter)
	require.Equal(h.t, energyOffersCountBefore, offerId)
	h.VerifyEnergyTransferOffer(ctx, offerId, offer)
	return offerId
}

func (h *C4eEvUtils) PublishEnergyTransferOfferError(ctx sdk.Context, creator string, offer types.EnergyTransferOffer, errorMessage string) {
	msgPublishEnergyTransferOffer := &types.MsgPublishEnergyTransferOffer{
		Creator:   creator,
		ChargerId: offer.ChargerId,
		Tariff:    offer.Tariff,
		Location:  offer.Location,
		Name:      offer.Name,
		PlugType:  offer.PlugType,
	}

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	energyOffersCountBefore := h.helperCfeevKeeper.GetEnergyTransferOfferCount(ctx)
	_, err := msgServer.PublishEnergyTransferOffer(ctx, msgPublishEnergyTransferOffer)
	require.EqualError(h.t, err, errorMessage)
	_, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransferOffer(ctx, energyOffersCountBefore+1)
	require.False(h.t, found)
	energyOffersCountAfter := h.helperCfeevKeeper.GetEnergyTransferOfferCount(ctx)
	require.Equal(h.t, energyOffersCountAfter, energyOffersCountBefore)
}

func (h *C4eEvUtils) StartEnergyTransfer(ctx sdk.Context, driver string, transfer types.EnergyTransfer, offerId uint64) uint64 {
	msgStartTransfer := &types.MsgStartEnergyTransfer{
		Creator:               driver,
		EnergyTransferOfferId: offerId,
		OfferedTariff:         transfer.OfferedTariff,
		EnergyToTransfer:      transfer.EnergyToTransfer,
	}
	offer, found := h.helperCfeevKeeper.GetEnergyTransferOffer(ctx, offerId)
	require.True(h.t, found)
	transfer.Owner = offer.Owner
	transfer.Driver = driver
	transfer.EnergyTransferOfferId = offerId
	transfer.EnergyTransferred = uint64(0)
	transfer.ChargerId = offer.ChargerId
	transfer.Status = types.TransferStatus_REQUESTED
	transfer.Collateral = math.NewIntFromUint64(transfer.EnergyToTransfer * transfer.OfferedTariff)
	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	energyTransferCountBefore := h.helperCfeevKeeper.GetEnergyTransferCount(ctx)
	startTransferResponse, err := msgServer.StartEnergyTransfer(ctx, msgStartTransfer)
	require.NoError(h.t, err)
	energyTransferId := startTransferResponse.GetId()
	energyTransferCountAfter := h.helperCfeevKeeper.GetEnergyTransferCount(ctx)
	h.VerifyEnergyTransferOfferStatus(ctx, offerId, types.ChargerStatus_BUSY)
	h.VerifyEnergyTransferstatus(ctx, energyTransferId, types.TransferStatus_REQUESTED)
	h.VerifyEnergyTransfer(ctx, energyTransferId, transfer)
	require.Equal(h.t, energyTransferCountBefore+1, energyTransferCountAfter)
	return energyTransferId
}

func (h *C4eEvUtils) StartEnergyTransferError(ctx sdk.Context, driver string, transfer types.EnergyTransfer, offerId uint64, errorMessage string) {
	msgStartTransfer := &types.MsgStartEnergyTransfer{
		Creator:               driver,
		EnergyTransferOfferId: offerId,
		OfferedTariff:         transfer.OfferedTariff,
		EnergyToTransfer:      transfer.EnergyToTransfer,
	}
	offerBefore, energyTransferOfferFoundBefore := h.helperCfeevKeeper.GetEnergyTransferOffer(ctx, offerId)

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	energyTransferCountBefore := h.helperCfeevKeeper.GetEnergyTransferCount(ctx)
	_, err := msgServer.StartEnergyTransfer(ctx, msgStartTransfer)
	require.EqualError(h.t, err, errorMessage)
	energyTransferCountAfter := h.helperCfeevKeeper.GetEnergyTransferCount(ctx)
	if energyTransferOfferFoundBefore {
		h.VerifyEnergyTransferOffer(ctx, offerId, offerBefore)
	}

	require.Equal(h.t, energyTransferCountBefore, energyTransferCountAfter)
}

func (h *C4eEvUtils) EnergyTransferCompleted(ctx sdk.Context, energyTransferId uint64, usedServiceUnits uint64, offerId uint64,
	expectedDriverBalance math.Int, expectedChargerOwnerBalance math.Int) {
	msg := &types.MsgEnergyTransferCompleted{EnergyTransferId: energyTransferId, UsedServiceUnits: usedServiceUnits}

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	_, err := msgServer.EnergyTransferCompleted(ctx, msg)
	require.NoError(h.t, err)
	h.VerifyEnergyTransferOfferStatus(ctx, offerId, types.ChargerStatus_ACTIVE)
	h.VerifyEnergyTransferstatus(ctx, energyTransferId, types.TransferStatus_PAID)
	transfer, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, energyTransferId)
	require.True(h.t, found)
	h.BankUtils.VerifyAccountDefaultDenomBalance(ctx, sdk.MustAccAddressFromBech32(transfer.GetDriver()), expectedDriverBalance)
	h.BankUtils.VerifyAccountDefaultDenomBalance(ctx, sdk.MustAccAddressFromBech32(transfer.GetOwner()), expectedChargerOwnerBalance)
}

func (h *C4eEvUtils) EnergyTransferCompletedError(ctx sdk.Context, energyTransferId uint64, usedServiceUnits uint64, offerId uint64, errorMessage string) {
	msg := &types.MsgEnergyTransferCompleted{EnergyTransferId: energyTransferId, UsedServiceUnits: usedServiceUnits}
	transferBefore, transferFoundBefore := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, energyTransferId)
	offerBefore, offerFoundBeofre := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransferOffer(ctx, offerId)
	var driverBalanceBefore math.Int
	var ownerBalanceBefore math.Int
	if transferFoundBefore {
		driverAccount := sdk.MustAccAddressFromBech32(transferBefore.GetDriver())
		ownerAccount := sdk.MustAccAddressFromBech32(transferBefore.GetOwner())
		driverBalanceBefore = h.BankUtils.GetAccountDefultDenomBalance(ctx, driverAccount)
		ownerBalanceBefore = h.BankUtils.GetAccountDefultDenomBalance(ctx, ownerAccount)
	}

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	_, err := msgServer.EnergyTransferCompleted(ctx, msg)
	require.EqualError(h.t, err, errorMessage)
	if offerFoundBeofre {
		h.VerifyEnergyTransferOffer(ctx, offerId, offerBefore)
	}
	if transferFoundBefore {
		h.VerifyEnergyTransfer(ctx, energyTransferId, transferBefore)
		h.BankUtils.VerifyAccountDefaultDenomBalance(ctx, sdk.MustAccAddressFromBech32(transferBefore.GetDriver()), driverBalanceBefore)
		h.BankUtils.VerifyAccountDefaultDenomBalance(ctx, sdk.MustAccAddressFromBech32(transferBefore.GetOwner()), ownerBalanceBefore)
	}
}

func (h *C4eEvUtils) EnergyTransferStarted(ctx sdk.Context, energyTransferId uint64, offerId uint64) {
	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	msgConfirmTransferRequestStarted := &types.MsgEnergyTransferStarted{EnergyTransferId: energyTransferId}
	_, err := msgServer.EnergyTransferStarted(ctx, msgConfirmTransferRequestStarted)
	h.VerifyEnergyTransferOfferStatus(ctx, offerId, types.ChargerStatus_BUSY)
	h.VerifyEnergyTransferstatus(ctx, energyTransferId, types.TransferStatus_ONGOING)
	require.NoError(h.t, err)
}

func (h *C4eEvUtils) EnergyTransferStartedError(ctx sdk.Context, energyTransferId uint64, offerId uint64, errorMessage string) {
	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	msgConfirmTransferRequestStarted := &types.MsgEnergyTransferStarted{EnergyTransferId: energyTransferId}
	offerBefore, energyTransferOfferFoundBefore := h.helperCfeevKeeper.GetEnergyTransferOffer(ctx, offerId)
	transferBefore, energyTransferFoundBefore := h.helperCfeevKeeper.GetEnergyTransfer(ctx, offerId)
	_, err := msgServer.EnergyTransferStarted(ctx, msgConfirmTransferRequestStarted)
	require.EqualError(h.t, err, errorMessage)

	if energyTransferOfferFoundBefore {
		h.VerifyEnergyTransferOffer(ctx, offerId, offerBefore)
	}
	if energyTransferFoundBefore {
		h.VerifyEnergyTransfer(ctx, offerId, transferBefore)
	}
}

func RandomLocation(r *rand.Rand) *types.Location {
	latitude := utils.RandomDecBetween(r, -90, 90)
	longitude := utils.RandomDecBetween(r, -180, 180)
	return &types.Location{
		Latitude:  &latitude,
		Longitude: &longitude,
	}
}

func (h *C4eEvUtils) RemoveEnergyTransferOffer(ctx sdk.Context, owner string, offerId uint64) {
	msgRemoveTransferOffer := &types.MsgRemoveEnergyOffer{
		Owner: owner,
		Id:    offerId,
	}

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	energyOffersCountBefore := h.helperCfeevKeeper.GetEnergyTransferOfferCount(ctx)
	_, err := msgServer.RemoveEnergyOffer(ctx, msgRemoveTransferOffer)
	require.NoError(h.t, err)
	_, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransferOffer(ctx, offerId)
	energyOffersCountAfter := h.helperCfeevKeeper.GetEnergyTransferOfferCount(ctx)
	require.False(h.t, found)
	require.Equal(h.t, energyOffersCountBefore, energyOffersCountAfter)
}

func (h *C4eEvUtils) RemoveEnergyTransferOfferError(ctx sdk.Context, owner string, offerId uint64, errorMessage string) {
	msgRemoveOffer := &types.MsgRemoveEnergyOffer{
		Owner: owner,
		Id:    offerId,
	}
	_, existsBefore := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransferOffer(ctx, offerId)
	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	energyOffersCountBefore := h.helperCfeevKeeper.GetEnergyTransferOfferCount(ctx)
	_, err := msgServer.RemoveEnergyOffer(ctx, msgRemoveOffer)
	require.EqualError(h.t, err, errorMessage)
	if existsBefore {
		_, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransferOffer(ctx, offerId)
		require.True(h.t, found)
	}

	energyOffersCountAfter := h.helperCfeevKeeper.GetEnergyTransferOfferCount(ctx)
	require.Equal(h.t, energyOffersCountBefore, energyOffersCountAfter)
}

func (h *C4eEvUtils) RemoveEnergyTransfer(ctx sdk.Context, owner string, transferId uint64) {
	msgRemoveTransfer := &types.MsgRemoveTransfer{
		Owner: owner,
		Id:    transferId,
	}

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	energyTransferCountBefore := h.helperCfeevKeeper.GetEnergyTransferCount(ctx)
	_, err := msgServer.RemoveEnergyTransfer(ctx, msgRemoveTransfer)
	require.NoError(h.t, err)
	_, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, transferId)
	energyTransferCountAfter := h.helperCfeevKeeper.GetEnergyTransferCount(ctx)
	require.False(h.t, found)
	require.Equal(h.t, energyTransferCountBefore, energyTransferCountAfter)
}

func (h *C4eEvUtils) RemoveEnergyTransferError(ctx sdk.Context, owner string, transferId uint64, errorMessage string) {
	msgRemoveTransfer := &types.MsgRemoveTransfer{
		Owner: owner,
		Id:    transferId,
	}
	_, existsBefore := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, transferId)
	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	energyTransferCountBefore := h.helperCfeevKeeper.GetEnergyTransferCount(ctx)
	_, err := msgServer.RemoveEnergyTransfer(ctx, msgRemoveTransfer)
	require.EqualError(h.t, err, errorMessage)
	if existsBefore {
		_, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, transferId)
		require.True(h.t, found)
	}

	energyTransferCountAfter := h.helperCfeevKeeper.GetEnergyTransferCount(ctx)
	require.Equal(h.t, energyTransferCountBefore, energyTransferCountAfter)
}

func (h *C4eEvUtils) CancelEnergyTransfer(ctx sdk.Context, energyTransferId uint64, offerId uint64) {
	transferBefore, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, energyTransferId)
	ownerAccount := sdk.MustAccAddressFromBech32(transferBefore.GetOwner())
	ownerBalanceBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, ownerAccount)

	msg := &types.MsgCancelEnergyTransfer{
		Creator:          "",
		EnergyTransferId: energyTransferId,
		ErrorInfo:        "",
		ErrorCode:        "",
	}

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	_, err := msgServer.CancelEnergyTransfer(ctx, msg)
	require.NoError(h.t, err)
	h.VerifyEnergyTransferOfferStatus(ctx, offerId, types.ChargerStatus_ACTIVE)
	h.VerifyEnergyTransferstatus(ctx, energyTransferId, types.TransferStatus_CANCELLED)
	transfer, found := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, energyTransferId)
	require.True(h.t, found)
	h.BankUtils.VerifyAccountDefaultDenomBalance(ctx, sdk.MustAccAddressFromBech32(transfer.GetDriver()), math.NewIntFromUint64(transferBefore.EnergyToTransfer*transferBefore.OfferedTariff))
	h.BankUtils.VerifyAccountDefaultDenomBalance(ctx, sdk.MustAccAddressFromBech32(transfer.GetOwner()), ownerBalanceBefore)

}

func (h *C4eEvUtils) CancelEnergyTransferError(ctx sdk.Context, energyTransferId uint64, offerId uint64, errorMessage string) {
	transferBefore, transferFoundBefore := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransfer(ctx, energyTransferId)
	offerBefore, offerFoundBeofre := h.C4eEvKeeperUtils.helperCfeevKeeper.GetEnergyTransferOffer(ctx, offerId)
	var driverBalanceBefore math.Int
	var ownerBalanceBefore math.Int
	if transferFoundBefore {
		driverAccount := sdk.MustAccAddressFromBech32(transferBefore.GetDriver())
		ownerAccount := sdk.MustAccAddressFromBech32(transferBefore.GetOwner())
		driverBalanceBefore = h.BankUtils.GetAccountDefultDenomBalance(ctx, driverAccount)
		ownerBalanceBefore = h.BankUtils.GetAccountDefultDenomBalance(ctx, ownerAccount)
	}

	msg := &types.MsgCancelEnergyTransfer{
		Creator:          "",
		EnergyTransferId: energyTransferId,
		ErrorInfo:        "",
		ErrorCode:        "",
	}

	msgServer := keeper.NewMsgServerImpl(*h.helperCfeevKeeper)
	_, err := msgServer.CancelEnergyTransfer(ctx, msg)
	require.EqualError(h.t, err, errorMessage)
	if offerFoundBeofre {
		h.VerifyEnergyTransferOffer(ctx, offerId, offerBefore)
	}
	if transferFoundBefore {
		h.VerifyEnergyTransfer(ctx, energyTransferId, transferBefore)
		h.BankUtils.VerifyAccountDefaultDenomBalance(ctx, sdk.MustAccAddressFromBech32(transferBefore.GetDriver()), driverBalanceBefore)
		h.BankUtils.VerifyAccountDefaultDenomBalance(ctx, sdk.MustAccAddressFromBech32(transferBefore.GetOwner()), ownerBalanceBefore)
	}

}
