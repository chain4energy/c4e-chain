package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/v2/testutil/keeper"
	testutils "github.com/chain4energy/c4e-chain/v2/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestVestingTypesQueryEmpty(t *testing.T) {
	k, ctx := testkeeper.CfevestingKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)

	response, err := k.VestingType(wctx, &types.QueryVestingTypeRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryVestingTypeResponse{VestingTypes: []types.GenesisVestingType(nil)}, response)
}

func TestVestingTypesQueryNotEmpty(t *testing.T) {
	k, ctx := testkeeper.CfevestingKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	vestingTypes := types.VestingTypes{}
	vestingTypesArray := testutils.GenerateVestingTypes(10, 1)
	vestingTypes.VestingTypes = vestingTypesArray

	k.SetVestingTypes(ctx, vestingTypes)
	response, err := k.VestingType(wctx, &types.QueryVestingTypeRequest{})
	require.NoError(t, err)
	require.ElementsMatch(t, types.ConvertVestingTypesToGenesisVestingTypes(&vestingTypes), response.VestingTypes)
}
