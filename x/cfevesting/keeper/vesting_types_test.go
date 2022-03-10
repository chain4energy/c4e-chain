package keeper_test

import (
	"fmt"
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
)

func TestGetVestingTypes(t *testing.T) {
	k, ctx := testkeeper.CfevestingKeeper(t)
	vestingTypes := types.VestingTypes{}
	vestingType1 := types.VestingType{"test1", 2324, 42423, 4243, true}
	vestingType2 := types.VestingType{"test2", 1111, 112233, 445566, false}

	vestingTypesArray := []*types.VestingType{&vestingType1, &vestingType2}
	vestingTypes.VestingTypes = vestingTypesArray

	k.SetVestingTypes(ctx, vestingTypes)

	require.EqualValues(t, vestingTypes, k.GetVestingTypes(ctx))
	require.EqualValues(t, vestingType1, *k.GetVestingTypes(ctx).VestingTypes[0])
	require.EqualValues(t, vestingType2, *k.GetVestingTypes(ctx).VestingTypes[1])

}

func TestGetVestingTypeByName(t *testing.T) {
	k, ctx := testkeeper.CfevestingKeeper(t)
	vestingTypes := types.VestingTypes{}
	vestingType1 := types.VestingType{"test1", 2324, 42423, 4243, true}
	vestingType2 := types.VestingType{"test2", 1111, 112233, 445566, false}

	vestingTypesArray := []*types.VestingType{&vestingType1, &vestingType2}
	vestingTypes.VestingTypes = vestingTypesArray

	k.SetVestingTypes(ctx, vestingTypes)

	vt, err := k.GetVestingType(ctx, "test1")
	require.EqualValues(t, nil, err)
	require.EqualValues(t, vestingType1, vt)

	vt, err = k.GetVestingType(ctx, "test2")
	require.EqualValues(t, nil, err)
	require.EqualValues(t, vestingType2, vt)

	vt, err = k.GetVestingType(ctx, "not_exist")
	require.EqualValues(t, fmt.Errorf("vesting type not found: not_exist"), err)
	require.EqualValues(t, types.VestingType{}, vt)
}
