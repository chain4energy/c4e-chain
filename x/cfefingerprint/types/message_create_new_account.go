package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateNewAccount = "create_new_account"

// This message type is temporary and mainly used for testing purposes. It will be removed in the future.
var _ sdk.Msg = &MsgCreateNewAccount{}

// Temporary message - to remove in the future
func NewMsgCreateNewAccount(creator string, accAddressString string) *MsgCreateNewAccount {
	return &MsgCreateNewAccount{
		Creator:          creator,
		AccAddressString: accAddressString,
	}
}

func (msg *MsgCreateNewAccount) Route() string {
	return RouterKey
}

func (msg *MsgCreateNewAccount) Type() string {
	return TypeMsgCreateNewAccount
}

func (msg *MsgCreateNewAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNewAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNewAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
