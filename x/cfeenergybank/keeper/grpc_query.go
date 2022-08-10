package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
)

var _ types.QueryServer = Keeper{}
