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
	k, ctx, _ := testkeeper.CfeminterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	getParams := k.GetParams(ctx)
	require.EqualValues(t, params.MintDenom, getParams.MintDenom)
	testminter.CompareMinterConfigs(t, params.MinterConfig, getParams.MinterConfig)
}

func TestGetParams(t *testing.T) {
	k, ctx, _ := testkeeper.CfeminterKeeper(t)
	params := types.DefaultParams()
	params.MintDenom = "dfda"
	params.MinterConfig = types.MinterConfig{
		StartTime: time.Now().Add(time.Hour),
		Minters:   createLinearMintings(time.Now()),
	}
	k.SetParams(ctx, params)

	getParams := k.GetParams(ctx)
	require.EqualValues(t, params.MintDenom, getParams.MintDenom)
	testminter.CompareMinterConfigs(t, params.MinterConfig, getParams.MinterConfig)
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
	params := types.DefaultParams()
	minters := createLinearMintings(time.Now())
	timeNow := time.Now()
	minters[0].EndTime = &timeNow
	params.MinterConfig = types.MinterConfig{
		StartTime: time.Now().Add(time.Hour),
		Minters:   minters,
	}
	err := k.SetParams(ctx, params)
	require.Error(t, err, "first minter end must be bigger than minter start")
}
