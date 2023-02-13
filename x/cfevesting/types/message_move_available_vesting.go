package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMoveAvailableVesting = "move_available_vesting"

var _ sdk.Msg = &MsgMoveAvailableVesting{}

func NewMsgMoveAvailableVesting(fromAddress string, toAddress string) *MsgMoveAvailableVesting {
	return &MsgMoveAvailableVesting{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
	}
}

func (msg *MsgMoveAvailableVesting) Route() string {
	return RouterKey
}

func (msg *MsgMoveAvailableVesting) Type() string {
	return TypeMsgMoveAvailableVesting
}

func (msg *MsgMoveAvailableVesting) GetSigners() []sdk.AccAddress {
	fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{fromAddress}
}

func (msg *MsgMoveAvailableVesting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMoveAvailableVesting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid fromAddress address (%s)", err)
	}
	return nil
}
