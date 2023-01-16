package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ClaimRecordKeyPrefix is the prefix to retrieve all UserAirdropEntries
	ClaimRecordKeyPrefix = "UserAirdropEntries/value/"
)

// UserAirdropEntriesKey returns the store key to retrieve a UserAirdropEntries from the index fields
func UserAirdropEntriesKey(
	address string,
) []byte {
	var key []byte

	indexBytes := []byte(address)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
