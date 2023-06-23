package cfeevutils

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
)

func (h *ContextC4eEvUtils) PublishEnergyTransferOffer(creator string, offer types.EnergyTransferOffer) uint64 {
	return h.C4eEvUtils.PublishEnergyTransferOffer(h.testContext.GetContext(), creator, offer)
}

func (h *ContextC4eEvUtils) PublishEnergyTransferOfferError(creator string, offer types.EnergyTransferOffer, errorMessage string) {
	h.C4eEvUtils.PublishEnergyTransferOfferError(h.testContext.GetContext(), creator, offer, errorMessage)
}

func (h *ContextC4eEvUtils) VerifyEnergyTransferOfferStatus(id uint64, status types.ChargerStatus) {
	h.C4eEvUtils.VerifyEnergyTransferOfferStatus(h.testContext.GetContext(), id, status)
}

func (h *ContextC4eEvUtils) VerifyEnergyTransferstatus(id uint64, status types.TransferStatus) {
	h.C4eEvUtils.VerifyEnergyTransferstatus(h.testContext.GetContext(), id, status)
}

func (h *ContextC4eEvUtils) EnergyTransferCompleted(energyTransferId uint64, usedServiceUnits uint64, offerId uint64,
	expectedDriverBalance math.Int, expectedChargerOwnerBalance math.Int) {
	h.C4eEvUtils.EnergyTransferCompleted(h.testContext.GetContext(), energyTransferId, usedServiceUnits, offerId, expectedDriverBalance, expectedChargerOwnerBalance)
}

func (h *ContextC4eEvUtils) EnergyTransferCompletedError(energyTransferId uint64, usedServiceUnits uint64, offerId uint64, errorMessage string) {
	h.C4eEvUtils.EnergyTransferCompletedError(h.testContext.GetContext(), energyTransferId, usedServiceUnits, offerId, errorMessage)
}

func (h *ContextC4eEvUtils) EnergyTransferStarted(energyTransferId uint64, offerId uint64) {
	h.C4eEvUtils.EnergyTransferStarted(h.testContext.GetContext(), energyTransferId, offerId)
}

func (h *ContextC4eEvUtils) EnergyTransferStartedError(energyTransferId uint64, offerId uint64, errorMessage string) {
	h.C4eEvUtils.EnergyTransferStartedError(h.testContext.GetContext(), energyTransferId, offerId, errorMessage)
}

func (h *ContextC4eEvUtils) VerifyGetEnergyTransferOffer(id uint64) {
	h.C4eEvUtils.VerifyGetEnergyTransferOffer(h.testContext.GetContext(), id)
}

func (h *ContextC4eEvUtils) VerifyGetEnergyTransfer(id uint64) {
	h.C4eEvUtils.VerifyGetEnergyTransfer(h.testContext.GetContext(), id)
}

func (h *ContextC4eEvUtils) StartEnergyTransfer(driver string, transfer types.EnergyTransfer, offerId uint64) uint64 {
	return h.C4eEvUtils.StartEnergyTransfer(h.testContext.GetContext(), driver, transfer, offerId)
}

func (h *ContextC4eEvUtils) StartEnergyTransferError(driver string, transfer types.EnergyTransfer, offerId uint64, errorMessage string) {
	h.C4eEvUtils.StartEnergyTransferError(h.testContext.GetContext(), driver, transfer, offerId, errorMessage)
}

func (h *ContextC4eEvUtils) RemoveEnergyTransferOffer(owner string, offerId uint64) {
	h.C4eEvUtils.RemoveEnergyTransferOffer(h.testContext.GetContext(), owner, offerId)
}

func (h *ContextC4eEvUtils) RemoveEnergyTransferOfferError(owner string, offerId uint64, errorMessage string) {
	h.C4eEvUtils.RemoveEnergyTransferOfferError(h.testContext.GetContext(), owner, offerId, errorMessage)
}

func (h *ContextC4eEvUtils) RemoveEnergyTransfer(owner string, transferId uint64) {
	h.C4eEvUtils.RemoveEnergyTransfer(h.testContext.GetContext(), owner, transferId)
}

func (h *ContextC4eEvUtils) RemoveEnergyTransferError(owner string, transferId uint64, errorMessage string) {
	h.C4eEvUtils.RemoveEnergyTransferError(h.testContext.GetContext(), owner, transferId, errorMessage)
}

func (h *ContextC4eEvUtils) CancelEnergyTransfer(transferId uint64, offerId uint64) {
	h.C4eEvUtils.CancelEnergyTransfer(h.testContext.GetContext(), transferId, offerId)
}

func (h *ContextC4eEvUtils) CancelEnergyTransferError(transferId uint64, offerId uint64, errorMessage string) {
	h.C4eEvUtils.CancelEnergyTransferError(h.testContext.GetContext(), transferId, offerId, errorMessage)
}
