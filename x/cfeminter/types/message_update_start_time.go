package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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
