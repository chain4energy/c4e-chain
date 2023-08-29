package keeper

import (
	"context"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
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
	userDevices, found := k.GetUserDevices(ctx, msg.Owner)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	}
	found = false
	for _, device := range userDevices.Devices {
		if device.DeviceAddress == msg.DeviceAddress {
			found = true
			break
		}
	}
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "device not found")
	}
	device, found := k.GetDevice(ctx, msg.DeviceAddress)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "device not found")
	}

	powerSum := uint64(0)

	var measurements []*types.Measurement

	for _, measurementId := range msg.Measurements {
		measurement, err := device.GetMeasurement(measurementId)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, err.Error())
		}
		if measurement.UsedForCertificate {
			return nil, sdkerrors.Wrap(c4eerrors.ErrParam, "measurement already used for certificate")
		}
		powerSum += measurement.ReversePower
		measurement.UsedForCertificate = true
		measurements = append(measurements, measurement)
	}
	device.UsedEnergyProduced += powerSum
	userCertificates.Certificates = append(userCertificates.Certificates, &types.Certificate{
		Id:                 uint64(len(userCertificates.Certificates)),
		CertyficateTypeId:  msg.CertyficateTypeId,
		Power:              powerSum,
		DeviceAddress:      msg.DeviceAddress,
		AllowedAuthorities: msg.AllowedAuthorities,
		Measurements:       measurements,
		Authority:          "",
		CertificateStatus:  types.CertificateStatus_INVALID,
		ValidUntil:         nil,
	})

	k.SetUserCertificates(ctx, userCertificates)
	k.SetDevice(ctx, device)

	return &types.MsgCreateUserCertificatesResponse{}, nil
}

func (k msgServer) AuthorizeCertificate(goCtx context.Context, msg *types.MsgAuthorizeCertificate) (*types.MsgAuthorizeCertificateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userCertificates, found := k.GetUserCertificates(ctx, msg.UserAddress)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	}
	certificate, err := userCertificates.GetUserCertificate(msg.CertificateId)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "certificate not found")
	}
	if certificate.CertificateStatus != types.CertificateStatus_INVALID {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "certificate already authorized or burned")
	}
	if !certificate.ValidateAuthorizer(msg.Authorizer) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not authorized")
	}
	certificate.CertificateStatus = types.CertificateStatus_VALID
	certificate.Authority = msg.Authorizer
	certificate.ValidUntil = msg.ValidUntil

	k.SetUserCertificates(ctx, userCertificates)

	return &types.MsgAuthorizeCertificateResponse{}, nil
}

func (k msgServer) AddCertificateToMarketplace(goCtx context.Context, msg *types.MsgAddCertificateToMarketplace) (*types.MsgAddCertificateToMarketplaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userCertificates, found := k.GetUserCertificates(ctx, msg.Owner)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	}
	certificate, err := userCertificates.GetUserCertificate(msg.CertificateId)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "certificate not found")
	}
	if certificate.CertificateStatus != types.CertificateStatus_VALID {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "certificate not authorized, burned or on marketplace")
	}
	certificate.CertificateStatus = types.CertificateStatus_ON_MARKETPLACE

	certificateOffer := types.CertificateOffer{
		CertificateId: msg.CertificateId,
		Owner:         msg.Owner,
		Buyer:         "",
		Price:         msg.Price,
		Authorizer:    certificate.Authority,
		Power:         certificate.Power,
		ValidUntil:    certificate.ValidUntil,
		Measurements:  certificate.Measurements,
	}
	k.AppendMarketplaceCertificate(ctx, certificateOffer)
	k.SetUserCertificates(ctx, userCertificates)

	return &types.MsgAddCertificateToMarketplaceResponse{}, nil
}

func (k msgServer) BurnCertificate(goCtx context.Context, msg *types.MsgBurnCertificate) (*types.MsgBurnCertificateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userCertificates, found := k.GetUserCertificates(ctx, msg.Owner)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	}
	certificate, err := userCertificates.GetUserCertificate(msg.CertificateId)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "certificate not found")
	}
	if certificate.CertificateStatus != types.CertificateStatus_VALID {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "certificate already authorized or burned")
	}
	certificates := k.GetMarketplaceCertificates(ctx)
	for _, cert := range certificates {
		if cert.CertificateId == msg.CertificateId && cert.Owner == msg.Owner {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "certificate already in marketplace")
		}
	}
	certificate.CertificateStatus = types.CertificateStatus_BURNED

	k.SetUserCertificates(ctx, userCertificates)

	userDevices, found := k.GetUserDevices(ctx, msg.Owner)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	}
	found = false
	for _, device := range userDevices.Devices {
		if device.DeviceAddress == msg.DeviceAddress {
			found = true
			break
		}
	}
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "device not found")
	}
	device, found := k.GetDevice(ctx, msg.DeviceAddress)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "device not found")
	}

	powerLeft := device.EnergyConsumedSum - device.FulfilledEnergyConsumed
	if powerLeft <= certificate.Power {
		device.FulfilledEnergyConsumed = device.EnergyConsumedSum
	} else {
		device.FulfilledEnergyConsumed += certificate.Power
	}
	for _, certificateMeasurement := range certificate.Measurements {
		reversePowerLeft := certificateMeasurement.ReversePower
		for _, deviceMeasurement := range device.Measurements {
			fulfiledActivePowerSum := deviceMeasurement.GetFulfilledActivePowerSum()
			if deviceMeasurement.GetFulfilledActivePowerSum() != deviceMeasurement.ActivePower {
				diff := deviceMeasurement.Timestamp.Sub(certificateMeasurement.Timestamp)
				if diff < 0 {
					diff = -diff
				}
				if diff <= k.ActionTimeWindow(ctx) {
					powerLeft = deviceMeasurement.ActivePower - fulfiledActivePowerSum
					if powerLeft <= reversePowerLeft {
						deviceMeasurement.FulfilledActivePower = append(deviceMeasurement.FulfilledActivePower, &types.FulfilledActivePower{
							CertificateId: certificate.Id,
							Amount:        powerLeft,
						})
						reversePowerLeft -= powerLeft
					} else {
						deviceMeasurement.FulfilledActivePower = append(deviceMeasurement.FulfilledActivePower, &types.FulfilledActivePower{
							CertificateId: certificate.Id,
							Amount:        reversePowerLeft,
						})
						reversePowerLeft = 0
						break
					}
				}
			}
		}
	}

	k.SetDevice(ctx, device)
	return &types.MsgBurnCertificateResponse{}, nil
}

func (k msgServer) BuyCertificate(goCtx context.Context, msg *types.MsgBuyCertificate) (*types.MsgBuyCertificateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	marketplaceCertificate, found := k.GetMarketplaceCertificate(ctx, msg.MarketplaceCertificateId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "certificate not found")
	}
	if marketplaceCertificate.Buyer != "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "certificate already bought")
	}

	buyerAccAddr, err := sdk.AccAddressFromBech32(msg.Buyer)
	buyerSpendable := k.bankKeeper.SpendableCoins(ctx, buyerAccAddr)
	if buyerSpendable.IsAllLT(marketplaceCertificate.Price) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not enough coins")
	}
	ownerAccAddr, err := sdk.AccAddressFromBech32(marketplaceCertificate.Owner)
	params := k.GetParams(ctx)
	amountOfPrice := marketplaceCertificate.Price.AmountOf("uc4e")
	marketplacePrice := params.MarketplaceFee.MulInt(amountOfPrice).TruncateInt()
	authorityPrice := params.AuthorityFee.MulInt(amountOfPrice).TruncateInt()
	ownerPrice := amountOfPrice.Sub(marketplacePrice).Sub(authorityPrice)

	err = k.bankKeeper.SendCoins(ctx, buyerAccAddr, ownerAccAddr, sdk.NewCoins(sdk.NewCoin("uc4e", ownerPrice)))
	if err != nil {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrAmount, "send coins error (%v)", err)
	}

	marketplaceAccAddress, err := sdk.AccAddressFromBech32(params.MarketplaceOwnerAddress)
	err = k.bankKeeper.SendCoins(ctx, buyerAccAddr, marketplaceAccAddress, sdk.NewCoins(sdk.NewCoin("uc4e", marketplacePrice)))
	if err != nil {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrAmount, "send coins error (%v)", err)
	}

	authorityAccAddress, err := sdk.AccAddressFromBech32(marketplaceCertificate.Authorizer)
	err = k.bankKeeper.SendCoins(ctx, buyerAccAddr, authorityAccAddress, sdk.NewCoins(sdk.NewCoin("uc4e", authorityPrice)))
	if err != nil {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrAmount, "send coins error (%v)", err)
	}

	userCertificates, found := k.GetUserCertificates(ctx, marketplaceCertificate.Owner)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "user not found")
	}

	var certCopy types.Certificate
	for i, cert := range userCertificates.Certificates {
		if cert.Id == marketplaceCertificate.CertificateId {
			certCopy = *cert
			userCertificates.Certificates = remove(userCertificates.Certificates, i)
		}
	}
	certCopy.CertificateStatus = types.CertificateStatus_VALID

	buyerCertificates, found := k.GetUserCertificates(ctx, msg.Buyer)
	if !found {
		buyerCertificates = types.UserCertificates{
			Owner: msg.Buyer,
		}
	}
	certCopy.Id = uint64(len(buyerCertificates.Certificates))
	buyerCertificates.Certificates = append(buyerCertificates.Certificates, &certCopy)
	k.SetUserCertificates(ctx, buyerCertificates)
	k.SetUserCertificates(ctx, userCertificates)
	marketplaceCertificate.Buyer = msg.Buyer
	k.SetMarketplaceCertificate(ctx, marketplaceCertificate)

	return &types.MsgBuyCertificateResponse{}, nil
}

func remove(slice []*types.Certificate, s int) []*types.Certificate {
	return append(slice[:s], slice[s+1:]...)
}
