package types

var RemainsMapKey = []byte{0x03}
var RemainsKeyPrefix = []byte{0x04}

const (
	// ModuleName defines the module name
	ModuleName = "cferoutingdistributor"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cferoutingdistributor"

	CollectorName = ModuleName

	BurnStateKey = "burn_state_key"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
