package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendVesting = "send_vesting"

var _ sdk.Msg = &MsgSendVesting{}

func NewMsgSendVesting(fromAddress string, toAddress string, vestingId int32, amount sdk.Int, restartVesting bool) *MsgSendVesting {
	return &MsgSendVesting{
		FromAddress:    fromAddress,
		ToAddress:      toAddress,
		VestingId:      vestingId,
		Amount:         amount,
		RestartVesting: restartVesting,
	}
}

func (msg *MsgSendVesting) Route() string {
	return RouterKey
}

func (msg *MsgSendVesting) Type() string {
	return TypeMsgSendVesting
}

func (msg *MsgSendVesting) GetSigners() []sdk.AccAddress {
	fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{fromAddress}
}

func (msg *MsgSendVesting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendVesting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid fromAddress address (%s)", err)
	}
	return nil
}
