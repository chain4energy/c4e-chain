package keeper

import (
	"context"
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AssignDeviceToUser(goCtx context.Context, msg *types.MsgAssignDeviceToUser) (*types.MsgAssignDeviceToUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.GetPendingDevice(ctx, msg.DeviceAddress); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "device already assigned")
	}
	var pendingDevice = types.PendingDevice{
		DeviceAddress: msg.DeviceAddress,
		UserAddress:   msg.UserAddress,
	}

	k.SetPendingDevice(ctx, pendingDevice)

	return &types.MsgAssignDeviceToUserResponse{}, nil
}

func (k msgServer) AcceptDevice(goCtx context.Context, msg *types.MsgAcceptDevice) (*types.MsgAcceptDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pendingDevice, found := k.GetPendingDevice(ctx, msg.DeviceAddress)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.DeviceAddress))
	}
	if pendingDevice.UserAddress != msg.UserAddress {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect user")
	}

	userDevices, found := k.GetUserDevices(ctx, msg.UserAddress)
	if !found {
		userDevices = types.UserDevices{
			Owner: msg.UserAddress,
		}
	}
	userDevices.Devices = append(userDevices.Devices, &types.UserDevice{
		DeviceAddress: pendingDevice.DeviceAddress,
		Name:          msg.DeviceName,
		Location:      msg.DeviceLocation,
	})

	k.SetUserDevices(ctx, userDevices)

	return &types.MsgAcceptDeviceResponse{}, nil
}

func (k msgServer) AddMeasurement(goCtx context.Context, msg *types.MsgAddMeasurement) (*types.MsgAddMeasurementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	device, found := k.GetDevice(ctx, msg.DeviceAddress)
	if !found {
		device = types.Device{
			DeviceAddress: msg.DeviceAddress,
		}
	}
	device.Measurements = append(device.Measurements, &types.Measurement{
		Id:                 uint64(len(device.Measurements)),
		Timestamp:          *msg.Timestamp,
		ActivePower:        msg.ActivePower,
		UsedForCertificate: false,
		ReversePower:       msg.ReversePower,
		Metadata:           msg.Metadata,
	})
	device.ActivePowerSum += msg.ActivePower
	device.ReversePowerSum += msg.ReversePower

	k.SetDevice(ctx, device)

	return &types.MsgAddMeasurementResponse{}, nil
}
