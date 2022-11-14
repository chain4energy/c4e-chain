package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MissionKeyPrefix is the prefix to retrieve all Mission
	MissionKeyPrefix = "Mission/value/"
)

func MissionKey(campaignId uint64, MissionId uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, campaignId)

	bz = append(bz, []byte("/")...)
	bz2 := make([]byte, 8)
	binary.BigEndian.PutUint64(bz2, campaignId)
	bz = append(bz, bz2...)
	return bz
}
