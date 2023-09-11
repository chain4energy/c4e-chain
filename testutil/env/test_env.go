package env

import (
	"context"
	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"time"

	cfedistributortypes "github.com/chain4energy/c4e-chain/v2/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TestContext interface {
	GetContext() sdk.Context
	GetWrappedContext() context.Context
}

var TestEnvTime = time.Now()

const DefaultTestDenom = "uc4e"
const DefaultTestDenom2 = "uc4e2"
const DefaultDistributionDestination = cfedistributortypes.GreenEnergyBoosterCollector

var NoMintingConfig, _ = codectypes.NewAnyWithValue(&types.NoMinting{})
