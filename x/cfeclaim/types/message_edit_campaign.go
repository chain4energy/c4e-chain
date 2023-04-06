package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgEditCampaign = "edit_claim_campaign"

var _ sdk.Msg = &MsgEditCampaign{}

func NewMsgEditCampaign(owner string, campaignId uint64, name string, description string, feegrantAmount *sdk.Int, initialClaimFreeAmount *sdk.Int, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration) *MsgEditCampaign {
	return &MsgEditCampaign{
		Owner:                  owner,
		CampaignId:             campaignId,
		Name:                   name,
		Description:            description,
		FeegrantAmount:         feegrantAmount,
		InitialClaimFreeAmount: initialClaimFreeAmount,
		StartTime:              startTime,
		EndTime:                endTime,
		LockupPeriod:           lockupPeriod,
		VestingPeriod:          vestingPeriod,
	}
}

func (msg *MsgEditCampaign) Route() string {
	return RouterKey
}

func (msg *MsgEditCampaign) Type() string {
	return TypeMsgEditCampaign
}

func (msg *MsgEditCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEditCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
