package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
)

var _ types.QueryServer = Keeper{}
