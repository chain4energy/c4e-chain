package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateDenomParam = "update_denom_param"

var _ sdk.Msg = &MsgUpdateDenomParam{}

func NewMsgUpdateDenomParam(creator string) *MsgUpdateDenomParam {
	return &MsgUpdateDenomParam{
		Authority: creator,
	}
}

func (msg *MsgUpdateDenomParam) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDenomParam) Type() string {
	return TypeMsgUpdateDenomParam
}

func (msg *MsgUpdateDenomParam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDenomParam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDenomParam) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
