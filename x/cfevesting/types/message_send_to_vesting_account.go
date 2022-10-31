package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendToVestingAccount = "send_to_vesting_account"

var _ sdk.Msg = &MsgSendToVestingAccount{}

func NewMsgSendToVestingAccount(fromAddress string, toAddress string, vestingPoolName string, amount sdk.Int, restartVesting bool) *MsgSendToVestingAccount {
	return &MsgSendToVestingAccount{
		FromAddress:    fromAddress,
		ToAddress:      toAddress,
		VestingPoolName:      vestingPoolName,
		Amount:         amount,
		RestartVesting: restartVesting,
	}
}

func (msg *MsgSendToVestingAccount) Route() string {
	return RouterKey
}

func (msg *MsgSendToVestingAccount) Type() string {
	return TypeMsgSendToVestingAccount
}

func (msg *MsgSendToVestingAccount) GetSigners() []sdk.AccAddress {
	fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{fromAddress}
}

func (msg *MsgSendToVestingAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendToVestingAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid fromAddress address (%s)", err)
	}
	return nil
}
