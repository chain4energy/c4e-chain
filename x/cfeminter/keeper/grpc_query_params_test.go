package keeper_test

import (
	"testing"
	"time"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testminter "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"

)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.CfeminterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	params.MintDenom = "dfda"
	params.Minter = createMinter(time.Now())

	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.EqualValues(t, params.MintDenom, response.Params.MintDenom)
	testminter.CompareMinters(t, params.Minter, response.Params.Minter)

}
