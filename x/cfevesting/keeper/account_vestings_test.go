package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/stretchr/testify/require"
)

func TestGetAccountVestings(t *testing.T) {
	k, ctx := testkeeper.CfevestingKeeper(t)

	accountVestings := testutils.GenerateAccountVestingsWithRandomVestings(2, 10, 1, 1)

	k.SetAccountVestings(ctx, *accountVestings[0])

	accVestSored, _ := k.GetAccountVestings(ctx, accountVestings[0].Address)
	testutils.AssertAccountVestings(t, *accountVestings[0], accVestSored)

	testutils.AssertAccountVestings(t, *accountVestings[0], k.DeleteAccountVestings(ctx, accountVestings[0].Address))

	_, foundVest := k.GetAccountVestings(ctx, accountVestings[0].Address)
	require.False(t, foundVest, "Should not be found")

	k.SetAccountVestings(ctx, *accountVestings[0])

	accVestSored, _ = k.GetAccountVestings(ctx, accountVestings[0].Address)

	testutils.AssertAccountVestings(t, *accountVestings[0], accVestSored)

	k.SetAccountVestings(ctx, *accountVestings[1])

	accVestSored, _ = k.GetAccountVestings(ctx, accountVestings[1].Address)
	testutils.AssertAccountVestings(t, *accountVestings[1], accVestSored)

	allVestings := k.GetAllAccountVestings(ctx)
	require.EqualValues(t, 2, len(allVestings))

	found := false
	for i, accVestExp := range allVestings {
		fmt.Println("accVestExp: " + strconv.Itoa(i) + " - " + accVestExp.Address)
		if accountVestings[0].Address == accVestExp.Address {
			testutils.AssertAccountVestings(t, *accountVestings[0], accVestExp)
			found = true
		}
	}
	require.True(t, found, "not found: "+accountVestings[0].Address)

	found = false
	for i, accVestExp := range allVestings {
		fmt.Println("accVestExp: " + strconv.Itoa(i) + " - " + accVestExp.Address)
		if accountVestings[1].Address == accVestExp.Address {
			testutils.AssertAccountVestings(t, *accountVestings[1], accVestExp)
			found = true
		}
	}
	require.True(t, found, "not found: "+accountVestings[1].Address)

	testutils.AssertAccountVestingsArrays(t, accountVestings, testutils.ToAccountVestingsPointersArray(allVestings))

}
