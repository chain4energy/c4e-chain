package types

import (
	errortypes "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgCreateAirdropCampaign = "create_airdrop_campaign"

var _ sdk.Msg = &MsgCreateAirdropCampaign{}

func NewMsgCreateAirdropCampaign(owner string, name string, description string, allowFeegrant bool, initialClaimFreeAmount sdk.Int, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration) *MsgCreateAirdropCampaign {
	return &MsgCreateAirdropCampaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		AllowFeegrant:          allowFeegrant,
		InitialClaimFreeAmount: initialClaimFreeAmount,
		StartTime:              startTime,
		EndTime:                endTime,
		LockupPeriod:           lockupPeriod,
		VestingPeriod:          vestingPeriod,
	}
}

func (msg *MsgCreateAirdropCampaign) Route() string {
	return RouterKey
}

func (msg *MsgCreateAirdropCampaign) Type() string {
	return TypeMsgCreateAirdropCampaign
}

func (msg *MsgCreateAirdropCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAirdropCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAirdropCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Name == "" {
		return sdkerrors.Wrap(errortypes.ErrParam, "add mission to airdrop campaign empty name")
	}
	if msg.Description == "" {
		return sdkerrors.Wrap(errortypes.ErrParam, "add mission to airdrop campaign empty description")
	}
	//if msg.StartTime.Before(ctx.BlockTime()) {
	//	k.Logger(ctx).Error("create airdrop campaign start time in the past", "startTime", startTime)
	//	return sdkerrors.Wrapf(errortypes.ErrParam, "create airdrop campaign - start time in the past error  (%s < %s)", startTime, ctx.BlockTime())
	//}
	if msg.StartTime.After(msg.EndTime) {
		//k.Logger(ctx).Error("create airdrop campaign start time is after end time", "startTime", startTime, "endTime", endTime)
		return sdkerrors.Wrapf(errortypes.ErrParam, "create airdrop campaign - start time is after end time error (%s > %s)", msg.StartTime, msg.EndTime)
	}
	return nil
}
