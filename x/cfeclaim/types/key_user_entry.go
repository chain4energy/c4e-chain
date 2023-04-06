package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserEntryKeyPrefix is the prefix to retrieve all UserEntry
	UserEntryKeyPrefix = "UserEntry/value/"
)

// UserEntryKey returns the store key to retrieve a UserEntry from the index fields
func UserEntryKey(
	address string,
) []byte {
	var key []byte

	indexBytes := []byte(address)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
