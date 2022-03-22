package keeper_test

import (
	"strconv"
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/internal/testutils"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestVesting(t *testing.T) {
	height := int64(0)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.DelegableAddress = ""

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestings, height)
}

func TestVestingWithDelegableAddress(t *testing.T) {
	height := int64(0)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestings, height)

}

func TestVestingSomeToWithdraw(t *testing.T) {
	height := int64(10100)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, height)

}

func TestVestingSomeToWithdrawAndSomeWithdrawn(t *testing.T) {
	height := int64(10100)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.Vestings[0].Withdrawn = sdk.NewInt(500)
	accountVestings.Vestings[0].LastModificationWithdrawn = sdk.NewInt(500)

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestings, height)

}

func TestVestingSentAfterLockEndReceivingSide(t *testing.T) {
	height := int64(10100)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.Vestings[0].VestingStartBlock = accountVestings.Vestings[0].LockEndBlock
	accountVestings.Vestings[0].LastModificationBlock = accountVestings.Vestings[0].LockEndBlock

	accountVestings.Vestings[0].LockEndBlock -= 100

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, height)

}

func TestVestingSentAfterLockEndSendingSide(t *testing.T) {
	height := int64(10100)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	accountVestings.Vestings[0].LastModificationBlock = accountVestings.Vestings[0].LockEndBlock
	accountVestings.Vestings[0].Sent = sdk.NewInt(100000)
	accountVestings.Vestings[0].LastModificationVested = accountVestings.Vestings[0].LastModificationVested.Sub(sdk.NewInt(100000))

	accountVestings.Vestings[0].LockEndBlock -= 100

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, height)

}

func TestVestingSentAfterLockEndSendingSideAndWithdrawn(t *testing.T) {
	height := int64(10100)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	accountVestings.Vestings[0].LastModificationBlock = accountVestings.Vestings[0].LockEndBlock
	accountVestings.Vestings[0].Sent = sdk.NewInt(100000)
	accountVestings.Vestings[0].LastModificationVested = accountVestings.Vestings[0].LastModificationVested.Sub(sdk.NewInt(100000))
	accountVestings.Vestings[0].LastModificationWithdrawn = sdk.NewInt(400)

	accountVestings.Vestings[0].LockEndBlock -= 100

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, height)

}

func TestVestingManyVestings(t *testing.T) {
	height := int64(0)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(3, 1, 1)
	accountVestings.Address = addr

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, height)

}

func verifyVestingResponse(t *testing.T, response *types.QueryVestingResponse, accVestings types.AccountVestings, currentHeight int64) {
	require.EqualValues(t, len(accVestings.Vestings), len(response.Vestings))
	require.EqualValues(t, accVestings.DelegableAddress, response.DelegableAddress)

	for _, vesting := range accVestings.Vestings {
		found := false
		for _, vestingInfo := range response.Vestings {
			if vesting.Id == vestingInfo.Id {
				require.EqualValues(t, vesting.VestingType, vestingInfo.VestingType)
				require.EqualValues(t, testutils.GetExpectedWithdrawableForVesting(*vesting, currentHeight).String(), response.Vestings[0].Withdrawable)
				require.EqualValues(t, vesting.VestingStartBlock, vestingInfo.VestingStartHeight)
				require.EqualValues(t, vesting.LockEndBlock, vestingInfo.LockEndHeight)
				require.EqualValues(t, vesting.VestingEndBlock, vestingInfo.VestingEndHeight)
				require.EqualValues(t, "uc4e", response.Vestings[0].Vested.Denom)
				require.EqualValues(t, vesting.Vested, response.Vestings[0].Vested.Amount)
				require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn).String(), response.Vestings[0].CurrentVestedAmount)
				require.EqualValues(t, true, response.Vestings[0].DelegationAllowed)
				require.EqualValues(t, vesting.Sent.String(), response.Vestings[0].SentAmount)

				found = true
			}
		}
		require.True(t, found, "not found vesting nfo with Id: "+strconv.FormatInt(int64(vesting.Id), 10))
	}
}
