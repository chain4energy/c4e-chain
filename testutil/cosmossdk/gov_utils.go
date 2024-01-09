package cosmossdk

import (
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	"github.com/stretchr/testify/require"
)

type GovUtils struct {
	t         require.TestingT
	GovKeeper *govkeeper.Keeper
}

func NewGovUtils(t require.TestingT, govKeeper *govkeeper.Keeper) GovUtils {
	return GovUtils{t: t, GovKeeper: govKeeper}
}

type ContextGovUtils struct {
	GovUtils
	testContext testenv.TestContext
}

func NewContextGovUtils(t require.TestingT, testContext testenv.TestContext, govKeeper *govkeeper.Keeper) *ContextGovUtils {
	govUtils := NewGovUtils(t, govKeeper)
	return &ContextGovUtils{GovUtils: govUtils, testContext: testContext}
}
