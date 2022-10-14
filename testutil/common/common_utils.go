package common

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TestContext interface {
	GetContext() sdk.Context
	GetWrappedContext() context.Context
}


