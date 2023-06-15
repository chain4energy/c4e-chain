package types

import "github.com/chain4energy/c4e-chain/types/util"

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

func MissionKey(campaignId uint64, missionId uint64) []byte {
	return append(util.GetUint64Key(campaignId), util.GetUint64Key(missionId)...)
}
