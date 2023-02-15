package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawAllAvailable = "withdraw_all_available"

var _ sdk.Msg = &MsgWithdrawAllAvailable{}

func NewMsgWithdrawAllAvailable(owner string) *MsgWithdrawAllAvailable {
	return &MsgWithdrawAllAvailable{
		Owner: owner,
	}
}

func (msg *MsgWithdrawAllAvailable) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawAllAvailable) Type() string {
	return TypeMsgWithdrawAllAvailable
}

func (msg *MsgWithdrawAllAvailable) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgWithdrawAllAvailable) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawAllAvailable) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
