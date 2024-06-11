package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
)

var _ types.QueryServer = Keeper{}
