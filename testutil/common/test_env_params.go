package common

import (
	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"time"
)

var TestEnvTime = time.Now()

const DefaultTestDenom = "uc4e"
const DefaultDistributionDestination = cfedistributortypes.GreenEnergyBoosterCollector
