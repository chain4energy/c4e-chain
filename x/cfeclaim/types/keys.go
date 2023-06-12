package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "cfeclaim"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cfeclaim"
)

var (
	CampaignKeyPrefix      = []byte{0x00}
	CampaignCountKeyPrefix = []byte{0x01}
	MissionKeyPrefix       = []byte{0x02}
	MissionCountKeyPrefix  = []byte{0x03}
	UserEntryKeyPrefix     = []byte{0x04}
)

func UserEntryKey(
	address string,
) []byte {
	var key []byte

	indexBytes := []byte(address)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func MissionKey(campaignId uint64, missionId uint64) []byte {
	bz := GetUint64Key(campaignId)
	return append(bz, GetUint64Key(missionId)...)
}

func GetUint64Key(campaignId uint64) []byte {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, campaignId)

	key = append(key, []byte("/")...)

	return key
}
