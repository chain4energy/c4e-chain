package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/stretchr/testify/require"
)

func TestGetAccountVestingPools(t *testing.T) {
	k, ctx := testkeeper.CfevestingKeeper(t)

	accountVestingPools := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(2, 10, 1, 1)

	k.SetAccountVestingPools(ctx, *accountVestingPools[0])

	accVestSored, _ := k.GetAccountVestingPools(ctx, accountVestingPools[0].Address)
	testutils.AssertAccountVestingPools(t, *accountVestingPools[0], accVestSored)

	testutils.AssertAccountVestingPools(t, *accountVestingPools[0], k.DeleteAccountVestingPools(ctx, accountVestingPools[0].Address))

	_, foundVest := k.GetAccountVestingPools(ctx, accountVestingPools[0].Address)
	require.False(t, foundVest, "Should not be found")

	k.SetAccountVestingPools(ctx, *accountVestingPools[0])

	accVestSored, _ = k.GetAccountVestingPools(ctx, accountVestingPools[0].Address)

	testutils.AssertAccountVestingPools(t, *accountVestingPools[0], accVestSored)

	k.SetAccountVestingPools(ctx, *accountVestingPools[1])

	accVestSored, _ = k.GetAccountVestingPools(ctx, accountVestingPools[1].Address)
	testutils.AssertAccountVestingPools(t, *accountVestingPools[1], accVestSored)

	allVestingPools := k.GetAllAccountVestingPools(ctx)
	require.EqualValues(t, 2, len(allVestingPools))

	found := false
	for i, accVestExp := range allVestingPools {
		fmt.Println("accVestExp: " + strconv.Itoa(i) + " - " + accVestExp.Address)
		if accountVestingPools[0].Address == accVestExp.Address {
			testutils.AssertAccountVestingPools(t, *accountVestingPools[0], accVestExp)
			found = true
		}
	}
	require.True(t, found, "not found: "+accountVestingPools[0].Address)

	found = false
	for i, accVestExp := range allVestingPools {
		fmt.Println("accVestExp: " + strconv.Itoa(i) + " - " + accVestExp.Address)
		if accountVestingPools[1].Address == accVestExp.Address {
			testutils.AssertAccountVestingPools(t, *accountVestingPools[1], accVestExp)
			found = true
		}
	}
	require.True(t, found, "not found: "+accountVestingPools[1].Address)

	testutils.AssertAccountVestingPoolsArrays(t, accountVestingPools, testutils.ToAccountVestingPoolsPointersArray(allVestingPools))

}
