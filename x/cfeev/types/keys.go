package types

var (
	ParamsKey                   = []byte{0x00}
	EnergyTransferOfferKey      = []byte{0x01}
	EnergyTransferOfferCountKey = []byte{0x02}
	EnergyTransferKey           = []byte{0x03}
	EnergyTransferCountKey      = []byte{0x04}
)

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
