package cosmossdk

import (
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"testing"

	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
)

type FeegrantUtils struct {
	t              *testing.T
	FeegrantKeeper *feegrantkeeper.Keeper
}

func NewFeegrantUtils(t *testing.T, feegrantKeeper *feegrantkeeper.Keeper) FeegrantUtils {
	return FeegrantUtils{t: t, FeegrantKeeper: feegrantKeeper}
}

type ContextFeegrantUtils struct {
	FeegrantUtils
	testContext testenv.TestContext
}

func NewContextFeegrantUtils(t *testing.T, testContext testenv.TestContext, feegrantKeeper *feegrantkeeper.Keeper) *ContextFeegrantUtils {
	feegrantUtils := NewFeegrantUtils(t, feegrantKeeper)
	return &ContextFeegrantUtils{FeegrantUtils: feegrantUtils, testContext: testContext}
}
