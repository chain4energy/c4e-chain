package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveTransfer = "remove_transfer"

var _ sdk.Msg = &MsgRemoveTransfer{}

func NewMsgRemoveTransfer(owner string, id uint64) *MsgRemoveTransfer {
	return &MsgRemoveTransfer{
		Owner: owner,
		Id:    id,
	}
}

func (msg *MsgRemoveTransfer) Route() string {
	return RouterKey
}

func (msg *MsgRemoveTransfer) Type() string {
	return TypeMsgRemoveTransfer
}

func (msg *MsgRemoveTransfer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
