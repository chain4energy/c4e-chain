package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CampaignKeyPrefix is the prefix to retrieve all CampaignO
	CampaignKeyPrefix            = "Campaign/value/"
	CampaignTotalAmountKeyPrefix = "CampaignTotalAmount/value/"
	CampaignAmountLeftPrefix     = "CampaignAmountLeft/value/"
	CampaignCountKey             = "AidropCampaignCountKey-count-"
)

// CampaignKey returns the store key to retrieve a CampaignO from the index fields
func CampaignKey(
	id uint64,
) []byte {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, id)

	key = append(key, []byte("/")...)

	return key
}

// CampaignTotalAmountKey returns the store key to retrieve a CampaignO from the index fields
func CampaignTotalAmountKey(
	campaignId uint64,
) []byte {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, campaignId)

	key = append(key, []byte("/")...)

	return key
}

// CampaignAmountLeftKey returns the store key to retrieve a CampaignO from the index fields
func CampaignAmountLeftKey(
	campaignId uint64,
) []byte {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, campaignId)

	key = append(key, []byte("/")...)

	return key
}
