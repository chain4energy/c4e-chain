package keeper_test

import (
	"github.com/chain4energy/c4e-chain/testutil/app"
	"testing"
	"time"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testminter "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/stretchr/testify/require"
)

func TestGetDefaultParams(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	k := testHelper.C4eMinterUtils.GetC4eMinterKeeper()
	params := types.DefaultParams()
	err := k.SetParams(testHelper.Context, params)
	require.NoError(t, err)

	getParams := k.GetParams(testHelper.Context)
	testminter.CompareCfeminterParams(t, params, getParams)
}

func TestSetParamsNoDenom(t *testing.T) {
	k, ctx, _ := testkeeper.CfeminterKeeper(t)
	params := types.DefaultParams()
	params.MintDenom = ""
	err := k.SetParams(ctx, params)
	require.Error(t, err, "denom cannot be empty")
}

func TestSetParamsWrongMinterEndTime(t *testing.T) {
	k, ctx, _ := testkeeper.CfeminterKeeper(t)
	minters := createLinearMintings(time.Now())
	timeNow := time.Now()
	minters[0].EndTime = &timeNow
	params := types.Params{
		MintDenom: "dfda",
		StartTime: time.Now().Add(time.Hour),
		Minters:   minters,
	}

	err := k.SetParams(ctx, params)
	require.Error(t, err, "first minter end must be bigger than minter start")
}
