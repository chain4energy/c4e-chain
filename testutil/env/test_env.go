package env

import (
	"context"
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

const AuthorityModuleAddress = "c4e10d07y265gmmuvt4z0w9aw880jnsr700je62g0d"
