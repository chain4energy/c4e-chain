package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateMinters = "update_minters"

var _ sdk.Msg = &MsgUpdateMinters{}

func (msg *MsgUpdateMinters) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMinters) Type() string {
	return TypeMsgUpdateMinters
}

func (msg *MsgUpdateMinters) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMinters) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMinters) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

const TypeMsgUpdateMintDenom = "update_mint_denom"

var _ sdk.Msg = &MsgUpdateMintDenom{}

func (msg *MsgUpdateMintDenom) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMintDenom) Type() string {
	return TypeMsgUpdateMintDenom
}

func (msg *MsgUpdateMintDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMintDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMintDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgUpdateStartTime = "update_start_time"

var _ sdk.Msg = &MsgUpdateStartTime{}

func (msg *MsgUpdateStartTime) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStartTime) Type() string {
	return TypeMsgUpdateStartTime
}

func (msg *MsgUpdateStartTime) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStartTime) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStartTime) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
