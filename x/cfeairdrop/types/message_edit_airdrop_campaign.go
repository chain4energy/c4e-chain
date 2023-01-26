package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgEditAirdropCampaign = "edit_airdrop_campaign"

var _ sdk.Msg = &MsgEditAirdropCampaign{}

func NewMsgEditAirdropCampaign(owner string, campaignId uint64, name string, description string, feegrantAmount *sdk.Int, initialClaimFreeAmount *sdk.Int, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration) *MsgEditAirdropCampaign {
	return &MsgEditAirdropCampaign{
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

func (msg *MsgEditAirdropCampaign) Route() string {
	return RouterKey
}

func (msg *MsgEditAirdropCampaign) Type() string {
	return TypeMsgEditAirdropCampaign
}

func (msg *MsgEditAirdropCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEditAirdropCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditAirdropCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
