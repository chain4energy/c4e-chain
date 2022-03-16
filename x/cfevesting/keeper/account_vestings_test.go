package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
)

func TestGetAccountVestings(t *testing.T) {
	k, ctx := testkeeper.CfevestingKeeper(t)
	accountVestings := types.AccountVestings{}
	accountVestings.Address = "someAddr1"
	vesting1 := types.Vesting{
		VestingType:          "test1",
		VestingStartBlock:    2324,
		LockEndBlock:         42423,
		VestingEndBlock:      4243,
		Vested:               14243,
		Claimable:            24243,
		LastFreeingBlock:     34243,
		FreeCoinsBlockPeriod: 44243,
		FreeCoinsPerPeriod:   54243,
		DelegationAllowed:    true,
		Withdrawn:            0,
	}
	vesting2 := types.Vesting{
		VestingType:          "test2",
		VestingStartBlock:    92324,
		LockEndBlock:         942423,
		VestingEndBlock:      94243,
		Vested:               914243,
		Claimable:            924243,
		LastFreeingBlock:     934243,
		FreeCoinsBlockPeriod: 944243,
		FreeCoinsPerPeriod:   954243,
		DelegationAllowed:    false,
		Withdrawn:            0,
	}
	vestingsArray := []*types.Vesting{&vesting1, &vesting2}
	accountVestings.Vestings = vestingsArray

	k.SetAccountVestings(ctx, accountVestings)
	accVestSored, _ := k.GetAccountVestings(ctx, accountVestings.Address)
	require.EqualValues(t, accountVestings, accVestSored)
	require.EqualValues(t, vesting1, *accVestSored.Vestings[0])
	require.EqualValues(t, vesting2, *accVestSored.Vestings[1])

	require.EqualValues(t, accountVestings, k.DeleteAccountVestings(ctx, accountVestings.Address))

	_, foundVest := k.GetAccountVestings(ctx, accountVestings.Address)
	require.False(t, foundVest, "Should not be found")
	// require.PanicsWithValue(t, "stored minter should not have been nil", func() { k.GetAccountVestings(ctx, accountVestings.Address) }, "Code did not panic or wrong panic value")

	k.SetAccountVestings(ctx, accountVestings)

	accVestSored, _ = k.GetAccountVestings(ctx, accountVestings.Address)

	require.EqualValues(t, accountVestings, accVestSored)
	require.EqualValues(t, vesting1, *accVestSored.Vestings[0])
	require.EqualValues(t, vesting2, *accVestSored.Vestings[1])

	accountVestings2 := types.AccountVestings{}
	accountVestings2.Address = "someAddr2"

	vesting21 := types.Vesting{
		VestingType:          "test3",
		VestingStartBlock:    2324,
		LockEndBlock:         42423,
		VestingEndBlock:      4243,
		Vested:               14243,
		Claimable:            24243,
		LastFreeingBlock:     34243,
		FreeCoinsBlockPeriod: 44243,
		FreeCoinsPerPeriod:   54243,
		DelegationAllowed:    true,
		Withdrawn:            0,
	}
	vesting22 := types.Vesting{
		VestingType:          "test4",
		VestingStartBlock:    92324,
		LockEndBlock:         942423,
		VestingEndBlock:      94243,
		Vested:               914243,
		Claimable:            924243,
		LastFreeingBlock:     934243,
		FreeCoinsBlockPeriod: 944243,
		FreeCoinsPerPeriod:   954243,
		DelegationAllowed:    false,
		Withdrawn:            0,
	}

	// vesting21 := types.Vesting{"test3", 2324, 42423, 4243, 14243, 24243, 34243, 44243, 54243, true, 0}
	// vesting22 := types.Vesting{"test4", 92324, 942423, 94243, 914243, 924243, 934243, 944243, 954243, false, 0}
	vestingsArray2 := []*types.Vesting{&vesting21, &vesting22}
	accountVestings2.Vestings = vestingsArray2

	k.SetAccountVestings(ctx, accountVestings2)

	accVestSored, _ = k.GetAccountVestings(ctx, accountVestings2.Address)
	require.EqualValues(t, accountVestings2, accVestSored)
	require.EqualValues(t, vesting21, *accVestSored.Vestings[0])
	require.EqualValues(t, vesting22, *accVestSored.Vestings[1])

	allVestings := k.GetAllAccountVestings(ctx)
	require.EqualValues(t, 2, len(allVestings))

	found := false
	for i, accVestExp := range allVestings {
		fmt.Println("accVestExp: " + strconv.Itoa(i) + " - " + accVestExp.Address)
		if accountVestings.Address == accVestExp.Address {
			require.EqualValues(t, accountVestings, accVestExp)
			found = true
		}
	}
	require.True(t, found, "not found: "+accountVestings.Address)

	found = false
	for i, accVestExp := range allVestings {
		fmt.Println("accVestExp: " + strconv.Itoa(i) + " - " + accVestExp.Address)
		if accountVestings2.Address == accVestExp.Address {
			require.EqualValues(t, accountVestings2, accVestExp)
			found = true
		}
	}
	require.True(t, found, "not found: "+accountVestings2.Address)
}
