package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgCreateAirdropCampaign = "create_airdrop_campaign"

var _ sdk.Msg = &MsgCreateAirdropCampaign{}

func NewMsgCreateAirdropCampaign(owner string, name string, description string, startTime int64,
	endTime int64, lockupPeriod time.Duration, vestingPeriod time.Duration) *MsgCreateAirdropCampaign {
	return &MsgCreateAirdropCampaign{
		Owner:         owner,
		Name:          name,
		Description:   description,
		StartTime:     startTime,
		EndTime:       endTime,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
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
	return nil
}
