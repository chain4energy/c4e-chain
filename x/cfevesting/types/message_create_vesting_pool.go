package types

import (
	"cosmossdk.io/math"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateVestingPool = "create_vesting_pool"

var _ sdk.Msg = &MsgCreateVestingPool{}

func NewMsgCreateVestingPool(creator string, name string, amount math.Int, duration time.Duration, vestingType string) *MsgCreateVestingPool {
	return &MsgCreateVestingPool{
		Creator:     creator,
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
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVestingPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVestingPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
