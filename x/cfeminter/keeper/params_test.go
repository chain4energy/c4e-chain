package keeper_test

import (
	"testing"
	"time"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testminter "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/stretchr/testify/require"
)

func TestGetDefaultParams(t *testing.T) {
	k, ctx := testkeeper.CfeminterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	getParams := k.GetParams(ctx)
	require.EqualValues(t, params.MintDenom, getParams.MintDenom)
	testminter.CompareMinterParams(t, params, getParams)
}

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CfeminterKeeper(t)
	params := types.DefaultParams()
	params.MintDenom = "dfda"
	params.StartTime = time.Now().Add(time.Hour)
	params.Minters = createLinearMintings(time.Now())
	k.SetParams(ctx, params)

	getParams := k.GetParams(ctx)
	require.EqualValues(t, params.MintDenom, getParams.MintDenom)
	testminter.CompareMinterParams(t, params, getParams)
}

func TestSetParamsNoDenom(t *testing.T) {
	k, ctx := testkeeper.CfeminterKeeper(t)
	params := types.DefaultParams()
	params.MintDenom = ""
	params.StartTime = time.Now().Add(time.Hour)
	params.Minters = createLinearMintings(time.Now())
	require.PanicsWithValue(t, "value from ParamSetPair is invalid: denom cannot be empty", func() { k.SetParams(ctx, params) })
}

func TestSetParamsWrongMinterEndTime(t *testing.T) {
	k, ctx := testkeeper.CfeminterKeeper(t)
	params := types.DefaultParams()
	timeNow := time.Now()
	params.Minters = createLinearMintings(time.Now())
	params.Minters[0].EndTime = &timeNow
	params.MintDenom = "abc"
	params.StartTime = timeNow.Add(time.Hour)
	require.PanicsWithValue(t, "value from ParamSetPair is invalid: first minter end must be bigger than minter start", func() { k.SetParams(ctx, params) })
}
