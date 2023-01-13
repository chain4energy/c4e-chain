package keeper

import (
	errortypes "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

func (k Keeper) CreateAidropCampaign(ctx sdk.Context, owner string, name string, description string, startTime int64,
	endTime int64, lockupPeriod time.Duration, vestingPeriod time.Duration) error {
	k.Logger(ctx).Debug("create aidrop campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)
	if startTime <= ctx.BlockTime().Unix() {
		k.Logger(ctx).Error("create airdrop campaign start time in the past", "startTime", startTime)
		return sdkerrors.Wrapf(errortypes.ErrParam, "create airdrop campaign - start time in the past error  (%s < %s)", time.Unix(startTime, 0).String(), ctx.BlockTime())

	}
	if startTime > endTime {
		k.Logger(ctx).Error("create airdrop campaign start time is after end time", "startTime", startTime, "endTime", endTime)
		return sdkerrors.Wrapf(errortypes.ErrParam, "create airdrop campaign - start time is after end time error (%s > %s)", time.Unix(startTime, 0).String(), time.Unix(endTime, 0).String())
	}
	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Error("create vesting account owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(errortypes.ErrParsing, sdkerrors.Wrapf(err, "create vesting account - owner parsing error: %s", owner).Error())
	}

	campaign := types.NewAirdropCampaign(owner, name, description, time.Unix(startTime, 0), time.Unix(endTime, 0), lockupPeriod, vestingPeriod)
	k.AppendNewCampaign(ctx, *campaign)
	return nil
}
