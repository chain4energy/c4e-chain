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
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaignId %d", c.Id)
	}
	if blockTime.Before(*c.StartTime) {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaignId %d not started: time %s < startTime %s", c.Id, blockTime, c.StartTime)
	}
	if c.EndTime != nil && blockTime.After(*c.EndTime) {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaignId %d ended: time %s > endTime %s", c.Id, blockTime, c.EndTime)
	}
	return nil
}

func NewAirdropCampaign(owner string, name string, description string, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration) *Campaign {

	return &Campaign{
		Id:            0,
		Owner:         owner,
		Name:          name,
		Description:   description,
		Enabled:       false,
		StartTime:     &startTime,
		EndTime:       &endTime,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
	}
}
