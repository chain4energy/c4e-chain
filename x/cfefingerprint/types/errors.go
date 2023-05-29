package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/cfefingerprint module sentinel errors
var (
	ErrAlreadyExists = errors.Register(ModuleName, 3, "entity already exists")
)
