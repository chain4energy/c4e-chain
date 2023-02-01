package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UsersEntriesKeyPrefix is the prefix to retrieve all UserEntry
	UsersEntriesKeyPrefix = "UserEntry/value/"
)

// UsersEntriesKey returns the store key to retrieve a UserEntry from the index fields
func UsersEntriesKey(
	address string,
) []byte {
	var key []byte

	indexBytes := []byte(address)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
