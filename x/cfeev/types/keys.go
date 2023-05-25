package types

const (
	// ModuleName defines the module name
	ModuleName = "cfeev"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cfeev"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	EnergyTransferOfferKey      = "EnergyTransferOffer/value/"
	EnergyTransferOfferCountKey = "EnergyTransferOffer/count/"
)

const (
	EnergyTransferKey      = "EnergyTransfer/value/"
	EnergyTransferCountKey = "EnergyTransfer/count/"
)
