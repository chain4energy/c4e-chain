package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

const TypeMsgCreateVestingAccount = "create_vesting_account"

var _ sdk.Msg = &MsgCreateVestingAccount{}

func NewMsgCreateVestingAccount(fromAddress string, toAddress string, amount sdk.Coins, startTime int64, endTime int64) *MsgCreateVestingAccount {
	return &MsgCreateVestingAccount{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Amount:      amount,
		StartTime:   startTime,
		EndTime:     endTime,
	}
}

func (msg *MsgCreateVestingAccount) Route() string {
	return RouterKey
}

func (msg *MsgCreateVestingAccount) Type() string {
	return TypeMsgCreateVestingAccount
}

func (msg *MsgCreateVestingAccount) GetSigners() []sdk.AccAddress {
	fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{fromAddress}
}

func (msg *MsgCreateVestingAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVestingAccount) ValidateBasic() error {
	_, _, err := ValidateCreateVestingAccount(msg.FromAddress, msg.ToAddress, msg.Amount, msg.StartTime, msg.EndTime)
	return err
}

func ValidateCreateVestingAccount(fromAddress string, toAddress string, amount sdk.Coins, startTime int64,
	endTime int64) (fromAccAddress sdk.AccAddress, toAccAddress sdk.AccAddress, err error) {
	if amount == nil {
		return nil, nil, errors.Wrap(ErrParam, "create vesting account - coin amount cannot be nil")
	}
	if amount.IsAnyNil() {
		return nil, nil, errors.Wrap(ErrParam, "create vesting account - coin amount cannot be nil")
	}
	if amount.IsAnyNegative() {
		return nil, nil, errors.Wrap(ErrParam, "create vesting account - negative coin amount")
	}
	if startTime > endTime {
		return nil, nil, errors.Wrapf(ErrParam, "create vesting account - start time is after end time error (%s > %s)", time.Unix(startTime, 0).String(), time.Unix(endTime, 0).String())
	}
	fromAccAddress, err = sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		return nil, nil, errors.Wrap(ErrParsing, errors.Wrapf(err, "create vesting account - from-address parsing error: %s", fromAddress).Error())
	}
	toAccAddress, err = sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return nil, nil, errors.Wrap(ErrParsing, errors.Wrapf(err, "create vesting account - to-address parsing error: %s", toAddress).Error())
	}

	return fromAccAddress, toAccAddress, nil
}
