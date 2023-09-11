package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/v2/testutil/keeper"
	"github.com/chain4energy/c4e-chain/v2/x/cfedistributor/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx, _ := testkeeper.CfedistributorKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
