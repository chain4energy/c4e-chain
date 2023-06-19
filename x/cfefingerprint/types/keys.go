package types

var PayloadLinkKey = []byte{0x00}

const (
	// ModuleName defines the module name
	ModuleName = "cfefingerprint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cfefingerprint"
)

func GetStringKey(id string) []byte {
	return []byte(id)
}
