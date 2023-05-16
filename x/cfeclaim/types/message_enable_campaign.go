package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgEnableCampaign = "start_claim_campaign"

var _ sdk.Msg = &MsgEnableCampaign{}

func NewMsgEnableCampaign(owner string, campaignId uint64, startTime *time.Time, endTime *time.Time) *MsgEnableCampaign {
	return &MsgEnableCampaign{
		Owner:      owner,
		CampaignId: campaignId,
		StartTime:  startTime,
		EndTime:    endTime,
	}
}

func (msg *MsgEnableCampaign) Route() string {
	return RouterKey
}

func (msg *MsgEnableCampaign) Type() string {
	return TypeMsgEnableCampaign
}

func (msg *MsgEnableCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnableCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
