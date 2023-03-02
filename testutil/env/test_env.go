package env

import (
	"context"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"time"

	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TestContext interface {
	GetContext() sdk.Context
	GetWrappedContext() context.Context
}

var TestEnvTime = time.Now()

const DefaultTestDenom = "uc4e"
const DefaultDistributionDestination = cfedistributortypes.GreenEnergyBoosterCollector

func GetAuthority() string {
	return authtypes.NewModuleAddress(govtypes.ModuleName).String()
}
