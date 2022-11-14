package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// InitialClaimKeyPrefix is the prefix to retrieve all InitialClaim
	InitialClaimKeyPrefix = "InitialClaim/value/"
)
