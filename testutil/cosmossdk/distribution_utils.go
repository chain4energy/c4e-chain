package cosmossdk

import (
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/stretchr/testify/require"
)

type DistributionUtils struct {
	t           require.TestingT
	DistrKeeper *distrkeeper.Keeper
}

func NewDistributionUtils(t require.TestingT, distrKeeper *distrkeeper.Keeper) DistributionUtils {
	return DistributionUtils{t: t, DistrKeeper: distrKeeper}
}

type ContextDistributionUtils struct {
	DistributionUtils
	testContext testenv.TestContext
}

func NewContextDistributionUtils(t require.TestingT, testContext testenv.TestContext, distrKeeper *distrkeeper.Keeper) *ContextDistributionUtils {
	distributionUtils := NewDistributionUtils(t, distrKeeper)
	return &ContextDistributionUtils{DistributionUtils: distributionUtils, testContext: testContext}
}
