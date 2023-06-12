package cfeevutils

import "github.com/chain4energy/c4e-chain/x/cfeev/types"

func (h *ContextC4eEvUtils) PublishAndVerifyEnergyTransferOffer(offer types.EnergyTransferOffer) uint64 {
	return h.C4eEvUtils.PublishAndVerifyEnergyTransferOffer(h.testContext.GetContext(), offer)
}

func (h *ContextC4eEvUtils) VerifyEnergyTransferOfferStatus(id uint64, status types.ChargerStatus) {
	h.C4eEvUtils.VerifyEnergyTransferOfferStatus(h.testContext.GetContext(), id, status)
}

func (h *ContextC4eEvUtils) VerifyEnergyTransferStatus(id uint64, status types.TransferStatus) {
	h.C4eEvUtils.VerifyEnergyTransferStatus(h.testContext.GetContext(), id, status)
}

func (h *ContextC4eEvUtils) EnergyTransferCompleted(energyTransferId uint64, usedServiceUnits int32) {
	h.C4eEvUtils.EnergyTransferCompleted(h.testContext.GetContext(), energyTransferId, usedServiceUnits)
}

func (h *ContextC4eEvUtils) EnergyTransferStarted(energyTransferId uint64) {
	h.C4eEvUtils.EnergyTransferStarted(h.testContext.GetContext(), energyTransferId)
}

func (h *ContextC4eEvUtils) VerifyGetEnergyTransferOffer(id uint64) {
	h.C4eEvUtils.VerifyGetEnergyTransferOffer(h.testContext.GetContext(), id)
}

func (h *ContextC4eEvUtils) VerifyGetEnergyTransfer(id uint64) {
	h.C4eEvUtils.VerifyGetEnergyTransfer(h.testContext.GetContext(), id)
}

func (h *ContextC4eEvUtils) StartEnergyTransfer(transfer types.EnergyTransfer, newOfferId uint64) uint64 {
	return h.C4eEvUtils.StartEnergyTransfer(h.testContext.GetContext(), transfer, newOfferId)
}
