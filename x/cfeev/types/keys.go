package types

var (
	ParamsKey                    = []byte{0x00}
	VestingTypesKeyPrefix        = []byte{0x01}
	AccountVestingPoolsKeyPrefix = []byte{0x02}
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
