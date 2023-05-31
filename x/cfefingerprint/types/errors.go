package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/cfefingerprint module sentinel errors
var (
	ErrAlreadyExists = errors.Register(ModuleName, 1, "entity already exists")
	ErrParam         = errors.Register(ModuleName, 10, "wrong param value")
)
