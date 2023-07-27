package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateUserCertificate(goCtx context.Context, msg *types.MsgCreateUserCertificates) (*types.MsgCreateUserCertificatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userCertificates, found := k.GetUserCertificates(ctx, msg.Owner)
	if !found {
		userCertificates = types.UserCertificates{
			Owner: msg.Owner,
		}
	}
	_, found = k.GetDevice(ctx, msg.Owner)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	}
	device, found := k.GetDevice(ctx, msg.Owner)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "device not found")
	}
	if device.PowerSum-device.UsedPower < msg.Power {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not enough power")
	}

	device.UsedPower += msg.Power
	userCertificates.Certificates = append(userCertificates.Certificates, &types.Certificate{
		Id:                 uint64(len(userCertificates.Certificates)),
		CertyficateTypeId:  msg.CertyficateTypeId,
		Power:              msg.Power,
		DeviceAddress:      msg.DeviceAddress,
		AllowedAuthorities: msg.AllowedAuthorities,
		Authority:          "",
		CertificateStatus:  types.CertificateStatus_INVALID,
	})

	k.SetUserCertificates(ctx, userCertificates)
	k.SetDevice(ctx, device)

	return &types.MsgCreateUserCertificatesResponse{}, nil
}

func (k msgServer) AuthorizeCertificate(goCtx context.Context, msg *types.MsgAuthorizeCertificate) (*types.MsgAuthorizeCertificateResponse, error) {
	//ctx := sdk.UnwrapSDKContext(goCtx)
	//
	//userCertificates, found := k.GetUserCertificates(ctx, msg.Owner)
	//if !found {
	//	userCertificates = types.UserCertificates{
	//		Owner: msg.Owner,
	//	}
	//}
	//_, found = k.GetDevice(ctx, msg.Owner)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	//}
	//device, found := k.GetDevice(ctx, msg.Owner)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "device not found")
	//}
	//if device.PowerSum-device.UsedPower < msg.Power {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not enough power")
	//}
	//
	//device.UsedPower += msg.Power
	//userCertificates.Certificates = append(userCertificates.Certificates, &types.Certificate{
	//	Id:                 uint64(len(userCertificates.Certificates)),
	//	CertyficateTypeId:  msg.CertyficateTypeId,
	//	Power:              msg.Power,
	//	DeviceAddress:      msg.DeviceAddress,
	//	AllowedAuthorities: msg.AllowedAuthorities,
	//	Authority:          "",
	//	CertificateStatus:  types.CertificateStatus_INVALID,
	//})
	//
	//k.SetUserCertificates(ctx, userCertificates)
	//k.SetDevice(ctx, device)

	return &types.MsgAuthorizeCertificateResponse{}, nil
}

func (k msgServer) AddCertificateToMarketplace(goCtx context.Context, msg *types.MsgAddCertificateToMarketplace) (*types.MsgAddCertificateToMarketplaceResponse, error) {
	//ctx := sdk.UnwrapSDKContext(goCtx)
	//
	//userCertificates, found := k.GetUserCertificates(ctx, msg.Owner)
	//if !found {
	//	userCertificates = types.UserCertificates{
	//		Owner: msg.Owner,
	//	}
	//}
	//_, found = k.GetDevice(ctx, msg.Owner)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	//}
	//device, found := k.GetDevice(ctx, msg.Owner)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "device not found")
	//}
	//if device.PowerSum-device.UsedPower < msg.Power {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not enough power")
	//}
	//
	//device.UsedPower += msg.Power
	//userCertificates.Certificates = append(userCertificates.Certificates, &types.Certificate{
	//	Id:                 uint64(len(userCertificates.Certificates)),
	//	CertyficateTypeId:  msg.CertyficateTypeId,
	//	Power:              msg.Power,
	//	DeviceAddress:      msg.DeviceAddress,
	//	AllowedAuthorities: msg.AllowedAuthorities,
	//	Authority:          "",
	//	CertificateStatus:  types.CertificateStatus_INVALID,
	//})
	//
	//k.SetUserCertificates(ctx, userCertificates)
	//k.SetDevice(ctx, device)

	return &types.MsgAddCertificateToMarketplaceResponse{}, nil
}

func (k msgServer) BurnCertificate(goCtx context.Context, msg *types.MsgBurnCertificate) (*types.MsgBurnCertificateResponse, error) {
	//ctx := sdk.UnwrapSDKContext(goCtx)
	//
	//userCertificates, found := k.GetUserCertificates(ctx, msg.Owner)
	//if !found {
	//	userCertificates = types.UserCertificates{
	//		Owner: msg.Owner,
	//	}
	//}
	//_, found = k.GetDevice(ctx, msg.Owner)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	//}
	//device, found := k.GetDevice(ctx, msg.Owner)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "device not found")
	//}
	//if device.PowerSum-device.UsedPower < msg.Power {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not enough power")
	//}
	//
	//device.UsedPower += msg.Power
	//userCertificates.Certificates = append(userCertificates.Certificates, &types.Certificate{
	//	Id:                 uint64(len(userCertificates.Certificates)),
	//	CertyficateTypeId:  msg.CertyficateTypeId,
	//	Power:              msg.Power,
	//	DeviceAddress:      msg.DeviceAddress,
	//	AllowedAuthorities: msg.AllowedAuthorities,
	//	Authority:          "",
	//	CertificateStatus:  types.CertificateStatus_INVALID,
	//})
	//
	//k.SetUserCertificates(ctx, userCertificates)
	//k.SetDevice(ctx, device)

	return &types.MsgBurnCertificateResponse{}, nil
}

func (k msgServer) BuyCertificate(goCtx context.Context, msg *types.MsgBuyCertificate) (*types.MsgBuyCertificateResponse, error) {
	//ctx := sdk.UnwrapSDKContext(goCtx)
	//
	//userCertificates, found := k.GetUserCertificates(ctx, msg.Owner)
	//if !found {
	//	userCertificates = types.UserCertificates{
	//		Owner: msg.Owner,
	//	}
	//}
	//_, found = k.GetDevice(ctx, msg.Owner)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	//}
	//device, found := k.GetDevice(ctx, msg.Owner)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "device not found")
	//}
	//if device.PowerSum-device.UsedPower < msg.Power {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not enough power")
	//}
	//
	//device.UsedPower += msg.Power
	//userCertificates.Certificates = append(userCertificates.Certificates, &types.Certificate{
	//	Id:                 uint64(len(userCertificates.Certificates)),
	//	CertyficateTypeId:  msg.CertyficateTypeId,
	//	Power:              msg.Power,
	//	DeviceAddress:      msg.DeviceAddress,
	//	AllowedAuthorities: msg.AllowedAuthorities,
	//	Authority:          "",
	//	CertificateStatus:  types.CertificateStatus_INVALID,
	//})
	//
	//k.SetUserCertificates(ctx, userCertificates)
	//k.SetDevice(ctx, device)

	return &types.MsgBuyCertificateResponse{}, nil
}
