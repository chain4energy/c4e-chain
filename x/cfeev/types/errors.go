package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/cfeev module sentinel errors
var (
	ErrBusyCharger               = errors.Register(ModuleName, 1, "charger is busy")
	ErrWrongEnergyTransferStatus = errors.Register(ModuleName, 2, "energy transfer wrong status")
)
