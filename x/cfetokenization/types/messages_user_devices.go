package types

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const (
	TypeMsgAssingDeviceToUser = "assing_device_to_user"
	TypeMsgAcceptDevice       = "accept_device"
)

var _ sdk.Msg = &MsgAssignDeviceToUser{}

func NewMsgAssignDeviceToUser(deviceAddress, userAddress string) *MsgAssignDeviceToUser {
	return &MsgAssignDeviceToUser{
		UserAddress:   userAddress,
		DeviceAddress: deviceAddress,
	}
}

func (msg *MsgAssignDeviceToUser) Route() string {
	return RouterKey
}

func (msg *MsgAssignDeviceToUser) Type() string {
	return TypeMsgAssingDeviceToUser
}

func (msg *MsgAssignDeviceToUser) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.DeviceAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgAssignDeviceToUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAssignDeviceToUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.DeviceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid device address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgAcceptDevice{}

func NewMsgAcceptDevice(userAddress string, deviceAddress string, deviceName string) *MsgAcceptDevice {
	return &MsgAcceptDevice{
		UserAddress:   userAddress,
		DeviceAddress: deviceAddress,
		DeviceName:    deviceName,
	}
}

func (msg *MsgAcceptDevice) Route() string {
	return RouterKey
}

func (msg *MsgAcceptDevice) Type() string {
	return TypeMsgAcceptDevice
}

func (msg *MsgAcceptDevice) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgAcceptDevice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptDevice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.DeviceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid device address (%s)", err)
	}

	if msg.DeviceName == "" {
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "empty device name")
	}
	return nil
}

var _ sdk.Msg = &MsgAddMeasurement{}

func NewMsgAddMeasurement(deviceAddress string, timestamp *time.Time, activePower uint64, reversePower uint64, metadata string) *MsgAddMeasurement {
	return &MsgAddMeasurement{
		DeviceAddress: deviceAddress,
		Timestamp:     timestamp,
		ActivePower:   activePower,
		ReversePower:  reversePower,
		Metadata:      metadata,
	}
}

func (msg *MsgAddMeasurement) Route() string {
	return RouterKey
}

func (msg *MsgAddMeasurement) Type() string {
	return TypeMsgAcceptDevice
}

func (msg *MsgAddMeasurement) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.DeviceAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgAddMeasurement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddMeasurement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.DeviceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.DeviceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid device address (%s)", err)
	}

	if msg.Timestamp == nil {
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "empty timestamp")
	}
	return nil
}
