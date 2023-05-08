package types

var (
	ParamsKey                        = []byte{0x00}
	VestingTypesKeyPrefix            = []byte{0x01}
	AccountVestingPoolsKeyPrefix     = []byte{0x02}
	VestingPoolReservationsKeyPrefix = []byte{0x03}
)

const (
	// ModuleName defines the module name
	ModuleName = "cfevesting"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cfevesting"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	VestingAccountTraceKey      = "VestingAccountTrace-value-"
	VestingAccountTraceCountKey = "VestingAccountTrace-count-"
)

type VestingPoolReservations struct {
	VestingPoolReservations []VestingPoolReservation
}
