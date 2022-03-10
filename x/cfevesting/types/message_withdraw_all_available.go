package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawAllAvailable = "withdraw_all_available"

var _ sdk.Msg = &MsgWithdrawAllAvailable{}

func NewMsgWithdrawAllAvailable(creator string) *MsgWithdrawAllAvailable {
	return &MsgWithdrawAllAvailable{
		Creator: creator,
	}
}

func (msg *MsgWithdrawAllAvailable) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawAllAvailable) Type() string {
	return TypeMsgWithdrawAllAvailable
}

func (msg *MsgWithdrawAllAvailable) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawAllAvailable) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawAllAvailable) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
