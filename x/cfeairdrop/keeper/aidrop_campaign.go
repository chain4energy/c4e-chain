package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

func (k Keeper) CreateAidropCampaign(ctx sdk.Context, creator string, owner string, name string, campaignDuration time.Duration, lockupPeriod time.Duration, vestingPeriod time.Duration, description string) error {
	k.Logger(ctx).Debug("create aidrop campaign", "creator", creator, "owner", owner, "name", name,
		"campaignDuration", campaignDuration, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod, "description", description)

	return nil
}
