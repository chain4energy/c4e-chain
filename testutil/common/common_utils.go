package common

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TestContext interface {
	GetContext() sdk.Context
}


