package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/stretchr/testify/require"
)

func TestGetAccountVestingPools(t *testing.T) {
	k, ctx := testkeeper.CfevestingKeeper(t)

	accountVestingPools := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(2, 10, 1, 1)

	k.SetAccountVestingPools(ctx, *accountVestingPools[0])

	accVestSored, _ := k.GetAccountVestingPools(ctx, accountVestingPools[0].Owner)
	testutils.AssertAccountVestingPools(t, *accountVestingPools[0], accVestSored)

	testutils.AssertAccountVestingPools(t, *accountVestingPools[0], k.DeleteAccountVestingPools(ctx, accountVestingPools[0].Owner))

	_, foundVest := k.GetAccountVestingPools(ctx, accountVestingPools[0].Owner)
	require.False(t, foundVest, "Should not be found")

	k.SetAccountVestingPools(ctx, *accountVestingPools[0])

	accVestSored, _ = k.GetAccountVestingPools(ctx, accountVestingPools[0].Owner)

	testutils.AssertAccountVestingPools(t, *accountVestingPools[0], accVestSored)

	k.SetAccountVestingPools(ctx, *accountVestingPools[1])

	accVestSored, _ = k.GetAccountVestingPools(ctx, accountVestingPools[1].Owner)
	testutils.AssertAccountVestingPools(t, *accountVestingPools[1], accVestSored)

	allVestingPools := k.GetAllAccountVestingPools(ctx)
	require.EqualValues(t, 2, len(allVestingPools))

	found := false
	for _, accVestExp := range allVestingPools {
		if accountVestingPools[0].Owner == accVestExp.Owner {
			testutils.AssertAccountVestingPools(t, *accountVestingPools[0], accVestExp)
			found = true
		}
	}
	require.True(t, found, "not found: "+accountVestingPools[0].Owner)

	found = false
	for _, accVestExp := range allVestingPools {
		if accountVestingPools[1].Owner == accVestExp.Owner {
			testutils.AssertAccountVestingPools(t, *accountVestingPools[1], accVestExp)
			found = true
		}
	}
	require.True(t, found, "not found: "+accountVestingPools[1].Owner)

	testutils.AssertAccountVestingPoolsArrays(t, accountVestingPools, testutils.ToAccountVestingPoolsPointersArray(allVestingPools))

}
