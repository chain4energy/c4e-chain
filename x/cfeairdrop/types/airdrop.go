package types

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MessageId uint64

const (
    INITIAL MessageId = iota
    DELEGATION
    VOTE
	end
)

func IsMessageId(id uint64) bool {
	return id < uint64(end) 
}

func (cr *ClaimRecord) GetCampaignRecord(camapaignId uint64) *CampaignRecord {
	for _, campaignRecord := range cr.CampaignRecords {
		if campaignRecord.CampaignId == camapaignId {
			return campaignRecord
		}
	}
	return nil
}

func (c *Campaign) IsEnabled(blockTime time.Time) error {
	if !c.Enabled {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaignId %d", c.CampaignId)
	}
	if blockTime.Before(c.StartTime) || (c.EndTime != nil && blockTime.After(*c.EndTime)) {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaignId %d disabled due time", c.CampaignId)
	}
	return nil
}
