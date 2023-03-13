package types

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgCreateVestingPool = "create_vesting_pool"

var _ sdk.Msg = &MsgCreateVestingPool{}

func NewMsgCreateVestingPool(owner string, name string, amount math.Int, duration time.Duration, vestingType string) *MsgCreateVestingPool {
	return &MsgCreateVestingPool{
		Owner:       owner,
		Name:        name,
		Amount:      amount,
		Duration:    duration,
		VestingType: vestingType,
	}
}

func (msg *MsgCreateVestingPool) Route() string {
	return RouterKey
}

func (msg *MsgCreateVestingPool) Type() string {
	return TypeMsgCreateVestingPool
}

func (msg *MsgCreateVestingPool) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgCreateVestingPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVestingPool) ValidateBasic() error {
	_, err := ValidateCreateVestingPool(msg.Owner, msg.Name, msg.Amount, msg.Duration)
	if err != nil {
		return err
	}
	return nil
}

func ValidateCreateVestingPool(address string, vestingPoolName string, amount math.Int, duration time.Duration) (accAddress sdk.AccAddress, error error) {
	if vestingPoolName == "" {
		return nil, errors.Wrap(ErrParam, "add vesting pool empty name")
	}
	if amount.IsNil() {
		return nil, errors.Wrap(ErrAmount, "add vesting pool - amount cannot be nil")
	}
	if amount.IsNegative() {
		return nil, errors.Wrap(ErrAmount, "add vesting pool - amount is <= 0")
	}
	if duration <= 0 {
		return nil, errors.Wrap(ErrParam, "add vesting pool - duration is <= 0 or nil")
	}
	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return nil, errors.Wrap(ErrParsing, errors.Wrap(err, "add vesting pool - vesting acc address error").Error())
	}

	return accAddress, nil
}
