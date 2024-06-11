package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CfevestingKeeper(t)
	params := types.DefaultParams()
	params.Denom = "testDenom"
	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))

}
