package types

import (
	errortypes "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgCreateCampaign = "create_airdrop_campaign"

var _ sdk.Msg = &MsgCreateCampaign{}

func NewMsgCreateCampaign(owner string, name string, description string, argFeegrantAmount *sdk.Int, initialClaimFreeAmount *sdk.Int, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration) *MsgCreateCampaign {
	return &MsgCreateCampaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		FeegrantAmount:         argFeegrantAmount,
		InitialClaimFreeAmount: initialClaimFreeAmount,
		StartTime:              startTime,
		EndTime:                endTime,
		LockupPeriod:           lockupPeriod,
		VestingPeriod:          vestingPeriod,
	}
}

func (msg *MsgCreateCampaign) Route() string {
	return RouterKey
}

func (msg *MsgCreateCampaign) Type() string {
	return TypeMsgCreateCampaign
}

func (msg *MsgCreateCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCampaign) ValidateBasic() error {
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
	if msg.StartTime.Before(time.Now()) {
		return sdkerrors.Wrapf(errortypes.ErrParam, "create airdrop campaign - start time in the past error  (%s < %s)", msg.StartTime, time.Now())
	}
	if msg.StartTime.After(*msg.EndTime) {
		return sdkerrors.Wrapf(errortypes.ErrParam, "create airdrop campaign - start time is after end time error (%s > %s)", msg.StartTime, msg.EndTime)
	}
	return nil
}
