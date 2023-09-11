package keeper

import (
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
)

var _ types.QueryServer = Keeper{}
