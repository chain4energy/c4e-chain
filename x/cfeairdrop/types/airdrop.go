package types

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MessageId uint64

func (cr *UserAirdropEntries) GetAidropEntry(camapaignId uint64) *AirdropEntry {
	for _, airdropEntryState := range cr.AirdropEntries {
		if airdropEntryState.CampaignId == camapaignId {
			return airdropEntryState
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
