package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cfeev module sentinel errors
var (
	ErrOfferForChargerAlreadyExists       = sdkerrors.Register(ModuleName, 2, "energy transfer offer for this charger already exists")
	ErrBusyCharger                        = sdkerrors.Register(ModuleName, 3, "charger is busy")
	ErrEnergyTransferNotFound             = sdkerrors.Register(ModuleName, 4, "energy transfer not found")
	ErrEnergyTransferOfferNotFound        = sdkerrors.Register(ModuleName, 5, "energy transfer offer not found")
	ErrEnergyTransferOfferCannotBeRemoved = sdkerrors.Register(ModuleName, 6, "energy transfer offer cannot be removed")
	ErrCoinTransferFailed                 = sdkerrors.Register(ModuleName, 7, "coin transfer failed")
	ErrWrongEnergyTransferStatus          = sdkerrors.Register(ModuleName, 8, "energy transfer wrong status")
	ErrEnergyTransferCannotBeRemoved      = sdkerrors.Register(ModuleName, 9, "energy transfer cannot be removed")
)
