package cosmossdk

import (
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"testing"

	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
)

type GovUtils struct {
	t         *testing.T
	GovKeeper *govkeeper.Keeper
}

func NewGovUtils(t *testing.T, govKeeper *govkeeper.Keeper) GovUtils {
	return GovUtils{t: t, GovKeeper: govKeeper}
}

type ContextGovUtils struct {
	GovUtils
	testContext testenv.TestContext
}

func NewContextGovUtils(t *testing.T, testContext testenv.TestContext, govKeeper *govkeeper.Keeper) *ContextGovUtils {
	govUtils := NewGovUtils(t, govKeeper)
	return &ContextGovUtils{GovUtils: govUtils, testContext: testContext}
}
