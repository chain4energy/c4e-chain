package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestVesting(t *testing.T) {
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, 0)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := types.AccountVestings{}
	accountVestings.Address = addr
	vesting := types.Vesting{
		Id:                1,
		VestingType:       "test",
		VestingStartBlock: 1000,
		LockEndBlock:      10000,
		VestingEndBlock:   110000,
		Vested:            sdk.NewInt(1000000),
		// Claimable: 0,
		// LastFreeingBlock: 0,
		FreeCoinsBlockPeriod: 10,
		// FreeCoinsPerPeriod: 0,
		DelegationAllowed: true,
		Withdrawn:         sdk.ZeroInt(),
	}

	vestingsArray := []*types.Vesting{&vesting}
	accountVestings.Vestings = vestingsArray

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	require.EqualValues(t, 1, len(response.Vestings))
	require.EqualValues(t, "test", response.Vestings[0].VestingType)
	require.EqualValues(t, 1000, response.Vestings[0].VestingStartHeight)
	require.EqualValues(t, 1000, response.Vestings[0].VestingStartHeight)
	require.EqualValues(t, 10000, response.Vestings[0].LockEndHeight)
	require.EqualValues(t, 110000, response.Vestings[0].VestingEndHeight)
	require.EqualValues(t, "0", response.Vestings[0].Withdrawable)
	require.EqualValues(t, "uc4e", response.Vestings[0].Vested.Denom)
	require.EqualValues(t, sdk.NewInt(1000000), response.Vestings[0].Vested.Amount)

	require.EqualValues(t, "1000000", response.Vestings[0].CurrentVestedAmount)
	require.EqualValues(t, true, response.Vestings[0].DelegationAllowed)
	require.EqualValues(t, "", response.DelegableAddress)

}

func TestVestingWithDelegableAddress(t *testing.T) {
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, 0)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := types.AccountVestings{}
	accountVestings.Address = addr
	vesting := types.Vesting{
		Id:                1,
		VestingType:       "test",
		VestingStartBlock: 1000,
		LockEndBlock:      10000,
		VestingEndBlock:   110000,
		Vested:            sdk.NewInt(1000000),
		// Claimable: 0,
		// LastFreeingBlock: 0,
		FreeCoinsBlockPeriod: 10,
		// FreeCoinsPerPeriod: 0,
		DelegationAllowed: true,
		Withdrawn:         sdk.ZeroInt(),
	}

	vestingsArray := []*types.Vesting{&vesting}
	accountVestings.Vestings = vestingsArray
	accountVestings.DelegableAddress = "del addr"

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	require.EqualValues(t, 1, len(response.Vestings))
	require.EqualValues(t, "test", response.Vestings[0].VestingType)
	require.EqualValues(t, 1000, response.Vestings[0].VestingStartHeight)
	require.EqualValues(t, 1000, response.Vestings[0].VestingStartHeight)
	require.EqualValues(t, 10000, response.Vestings[0].LockEndHeight)
	require.EqualValues(t, 110000, response.Vestings[0].VestingEndHeight)
	require.EqualValues(t, "0", response.Vestings[0].Withdrawable)
	require.EqualValues(t, "uc4e", response.Vestings[0].Vested.Denom)
	require.EqualValues(t, sdk.NewInt(1000000), response.Vestings[0].Vested.Amount)

	require.EqualValues(t, "1000000", response.Vestings[0].CurrentVestedAmount)
	require.EqualValues(t, true, response.Vestings[0].DelegationAllowed)
	require.EqualValues(t, accountVestings.DelegableAddress, response.DelegableAddress)

}

func TestVestingSomeToWithdraw(t *testing.T) {
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, 10100)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := types.AccountVestings{}
	accountVestings.Address = addr
	vesting := types.Vesting{
		Id:                1,
		VestingType:       "test",
		VestingStartBlock: 1000,
		LockEndBlock:      10000,
		VestingEndBlock:   110000,
		Vested:            sdk.NewInt(1000000),
		// Claimable: 0,
		// LastFreeingBlock: 0,
		FreeCoinsBlockPeriod: 10,
		// FreeCoinsPerPeriod: 0,
		DelegationAllowed: true,
		Withdrawn:         sdk.ZeroInt(),
	}

	vestingsArray := []*types.Vesting{&vesting}
	accountVestings.Vestings = vestingsArray

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	require.EqualValues(t, 1, len(response.Vestings))
	require.EqualValues(t, "test", response.Vestings[0].VestingType)

	require.EqualValues(t, 1000, response.Vestings[0].VestingStartHeight)
	require.EqualValues(t, 1000, response.Vestings[0].VestingStartHeight)
	require.EqualValues(t, 10000, response.Vestings[0].LockEndHeight)
	require.EqualValues(t, 110000, response.Vestings[0].VestingEndHeight)
	require.EqualValues(t, "1000", response.Vestings[0].Withdrawable)
	require.EqualValues(t, "uc4e", response.Vestings[0].Vested.Denom)
	require.EqualValues(t, sdk.NewInt(1000000), response.Vestings[0].Vested.Amount)

	require.EqualValues(t, "1000000", response.Vestings[0].CurrentVestedAmount)
	require.EqualValues(t, true, response.Vestings[0].DelegationAllowed)

}

func TestVestingSomeToWithdrawAndSomeWithdrawn(t *testing.T) {
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, 10100)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := types.AccountVestings{}
	accountVestings.Address = addr
	vesting := types.Vesting{
		Id:                1,
		VestingType:       "test",
		VestingStartBlock: 1000,
		LockEndBlock:      10000,
		VestingEndBlock:   110000,
		Vested:            sdk.NewInt(1000000),
		// Claimable: 0,
		// LastFreeingBlock: 0,
		FreeCoinsBlockPeriod: 10,
		// FreeCoinsPerPeriod: 0,
		DelegationAllowed: true,
		Withdrawn:         sdk.NewInt(500),
	}

	vestingsArray := []*types.Vesting{&vesting}
	accountVestings.Vestings = vestingsArray

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	require.EqualValues(t, 1, len(response.Vestings))
	require.EqualValues(t, "test", response.Vestings[0].VestingType)
	require.EqualValues(t, 1000, response.Vestings[0].VestingStartHeight)
	require.EqualValues(t, 1000, response.Vestings[0].VestingStartHeight)
	require.EqualValues(t, 10000, response.Vestings[0].LockEndHeight)
	require.EqualValues(t, 110000, response.Vestings[0].VestingEndHeight)
	require.EqualValues(t, "500", response.Vestings[0].Withdrawable)
	require.EqualValues(t, "uc4e", response.Vestings[0].Vested.Denom)
	require.EqualValues(t, sdk.NewInt(1000000), response.Vestings[0].Vested.Amount)

	require.EqualValues(t, "999500", response.Vestings[0].CurrentVestedAmount)
	require.EqualValues(t, true, response.Vestings[0].DelegationAllowed)

}

func TestVestingManyVestings(t *testing.T) {
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, 0)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := types.AccountVestings{}
	accountVestings.Address = addr
	vesting1 := types.Vesting{
		Id:                1,
		VestingType:       "test1",
		VestingStartBlock: 1000,
		LockEndBlock:      10000,
		VestingEndBlock:   110000,
		Vested:            sdk.NewInt(1000000),
		// Claimable: 0,
		// LastFreeingBlock: 0,
		FreeCoinsBlockPeriod: 10,
		// FreeCoinsPerPeriod: 0,
		DelegationAllowed: true,
		Withdrawn:         sdk.ZeroInt(),
	}
	vesting2 := types.Vesting{
		Id:                2,
		VestingType:       "test2",
		VestingStartBlock: 1000,
		LockEndBlock:      10000,
		VestingEndBlock:   110000,
		Vested:            sdk.NewInt(10000000),
		// Claimable: 0,
		// LastFreeingBlock: 0,
		FreeCoinsBlockPeriod: 10,
		// FreeCoinsPerPeriod: 0,
		DelegationAllowed: true,
		Withdrawn:         sdk.ZeroInt(),
	}
	vesting3 := types.Vesting{
		Id:                3,
		VestingType:       "test3",
		VestingStartBlock: 1000,
		LockEndBlock:      10000,
		VestingEndBlock:   110000,
		Vested:            sdk.NewInt(100000000),
		// Claimable: 0,
		// LastFreeingBlock: 0,
		FreeCoinsBlockPeriod: 10,
		// FreeCoinsPerPeriod: 0,
		DelegationAllowed: true,
		Withdrawn:         sdk.ZeroInt(),
	}

	vestingsArray := []*types.Vesting{&vesting1, &vesting2, &vesting3}
	accountVestings.Vestings = vestingsArray

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	require.EqualValues(t, 3, len(response.Vestings))

	i := 0
	for _, vestingInfo := range response.Vestings {
		if vestingInfo.VestingType == "test1" {
			require.EqualValues(t, sdk.NewInt(1000000), vestingInfo.Vested.Amount)
			require.EqualValues(t, 1, vestingInfo.Id)

			i++
		} else if vestingInfo.VestingType == "test2" {
			require.EqualValues(t, sdk.NewInt(10000000), vestingInfo.Vested.Amount)
			require.EqualValues(t, 2, vestingInfo.Id)

			i++
		} else if vestingInfo.VestingType == "test3" {
			require.EqualValues(t, sdk.NewInt(100000000), vestingInfo.Vested.Amount)
			require.EqualValues(t, 3, vestingInfo.Id)

			i++
		}
	}
	require.EqualValues(t, 3, i)

}
