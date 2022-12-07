package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgCreateAirdropCampaign = "create_airdrop_campaign"

var _ sdk.Msg = &MsgCreateAirdropCampaign{}

func NewMsgCreateAirdropCampaign(creator string, owner string, name string, campaignDuration time.Duration, lockupPeriod time.Duration, vestingPeriod time.Duration, description string) *MsgCreateAirdropCampaign {
	return &MsgCreateAirdropCampaign{
		Creator:          creator,
		Owner:            owner,
		Name:             name,
		CampaignDuration: campaignDuration,
		LockupPeriod:     lockupPeriod,
		VestingPeriod:    vestingPeriod,
		Description:      description,
	}
}

func (msg *MsgCreateAirdropCampaign) Route() string {
	return RouterKey
}

func (msg *MsgCreateAirdropCampaign) Type() string {
	return TypeMsgCreateAirdropCampaign
}

func (msg *MsgCreateAirdropCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
