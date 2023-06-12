package types

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgSendToVestingAccount = "send_to_vesting_account"

var _ sdk.Msg = &MsgSendToVestingAccount{}

func NewMsgSendToVestingAccount(owner string, toAddress string, vestingPoolName string, amount math.Int, restartVesting bool) *MsgSendToVestingAccount {
	return &MsgSendToVestingAccount{
		Owner:           owner,
		ToAddress:       toAddress,
		VestingPoolName: vestingPoolName,
		Amount:          amount,
		RestartVesting:  restartVesting,
	}
}

func (msg *MsgSendToVestingAccount) Route() string {
	return RouterKey
}

func (msg *MsgSendToVestingAccount) Type() string {
	return TypeMsgSendToVestingAccount
}

func (msg *MsgSendToVestingAccount) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgSendToVestingAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendToVestingAccount) ValidateBasic() error {
	_, _, err := ValidateSendToVestingAccount(msg.Owner, msg.ToAddress, msg.VestingPoolName, msg.Amount)
	return err
}

func ValidateSendToVestingAccount(owner string, toAddr string, vestingPoolName string, amount math.Int) (ownerAccAddress sdk.AccAddress, toAccAddress sdk.AccAddress, error error) {
	if vestingPoolName == "" {
		return nil, nil, errors.Wrap(ErrParam, "send to new vesting account - empty name")
	}
	if amount.IsNil() {
		return nil, nil, errors.Wrap(ErrAmount, "send to new vesting account - amount cannot be nil")
	}
	if !amount.IsPositive() {
		return nil, nil, errors.Wrap(ErrAmount, "send to new vesting account - amount is <= 0")
	}
	if owner == toAddr {
		return nil, nil, errors.Wrapf(ErrIdenticalAccountsAddresses, "send to new vesting account - identical from address (%s) and to address (%s)", owner, toAddr)
	}
	ownerAccAddress, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return nil, nil, errors.Wrap(ErrParsing, errors.Wrap(err, "send to new vesting account - owner acc address error").Error())
	}
	toAccAddress, err = sdk.AccAddressFromBech32(toAddr)
	if err != nil {
		return nil, nil, errors.Wrap(ErrParsing, errors.Wrap(err, "send to new vesting account - to acc address error").Error())
	}

	return ownerAccAddress, toAccAddress, nil
}
