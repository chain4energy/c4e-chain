package util

import "encoding/binary"

func GetUint64Key(campaignId uint64) []byte {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, campaignId)

	return key
}
