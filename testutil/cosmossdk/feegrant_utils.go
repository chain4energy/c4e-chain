package cosmossdk

import (
	testenv "github.com/chain4energy/c4e-chain/v2/testutil/env"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	"github.com/stretchr/testify/require"
)

type FeegrantUtils struct {
	t              require.TestingT
	FeegrantKeeper *feegrantkeeper.Keeper
}

func NewFeegrantUtils(t require.TestingT, feegrantKeeper *feegrantkeeper.Keeper) FeegrantUtils {
	return FeegrantUtils{t: t, FeegrantKeeper: feegrantKeeper}
}

type ContextFeegrantUtils struct {
	FeegrantUtils
	testContext testenv.TestContext
}

func NewContextFeegrantUtils(t require.TestingT, testContext testenv.TestContext, feegrantKeeper *feegrantkeeper.Keeper) *ContextFeegrantUtils {
	feegrantUtils := NewFeegrantUtils(t, feegrantKeeper)
	return &ContextFeegrantUtils{FeegrantUtils: feegrantUtils, testContext: testContext}
}
