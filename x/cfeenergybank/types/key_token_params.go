package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TokenParamsKeyPrefix is the prefix to retrieve all TokenParams
	TokenParamsKeyPrefix = "TokenParams/value/"
)

// TokenParamsKey returns the store key to retrieve a TokenParams from the index fields
func TokenParamsKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
