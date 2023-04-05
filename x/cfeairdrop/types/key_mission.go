package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MissionKeyPrefix is the prefix to retrieve all Mission
	MissionKeyPrefix      = "Mission/value/"
	MissionCountKeyPrefix = "Mission/count/"
)

func MissionKey(campaignId uint64, missionId uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, campaignId)

	bz = append(bz, []byte("/")...)
	bz2 := make([]byte, 8)
	binary.BigEndian.PutUint64(bz2, missionId)
	bz = append(bz, bz2...)
	return bz
}

func MissionCountKey(campaignId uint64) []byte {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, campaignId)

	key = append(key, []byte("/")...)

	return key
}
