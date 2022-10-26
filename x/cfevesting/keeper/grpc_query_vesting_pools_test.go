package keeper_test

import (
	"testing"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestVesting(t *testing.T) {
	height := int64(0)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestingPools := testutils.GenerateOneAccountVestingPoolsWithAddressWith10BasedVestingPools(1, 1, 1)
	accountVestingPools.Address = addr

	keeper.SetAccountVestingPools(ctx, accountVestingPools)

	response, err := keeper.VestingPools(wctx, &types.QueryVestingPoolsRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestingPools, ctx.BlockTime(), true)
}

func TestVestingSomeToWithdraw(t *testing.T) {
	height := int64(10100)
	time := commontestutils.TestEnvTime.Add(testutils.CreateDurationFromNumOfHours(10100))
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeightAndTime(t, height, time)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestingPools := testutils.GenerateOneAccountVestingPoolsWithAddressWith10BasedVestingPools(1, 1, 1)
	accountVestingPools.Address = addr

	keeper.SetAccountVestingPools(ctx, accountVestingPools)

	response, err := keeper.VestingPools(wctx, &types.QueryVestingPoolsRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestingPools, ctx.BlockTime(), true)

}

func TestVestingSomeToWithdrawAndSomeWithdrawn(t *testing.T) {
	height := int64(10100)
	time := commontestutils.TestEnvTime.Add(testutils.CreateDurationFromNumOfHours(10100))
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeightAndTime(t, height, time)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestingPools := testutils.GenerateOneAccountVestingPoolsWithAddressWith10BasedVestingPools(1, 1, 1)
	accountVestingPools.Address = addr
	accountVestingPools.VestingPools[0].Withdrawn = sdk.NewInt(500)
	// accountVestingPools.VestingPools[0].LastModificationWithdrawn = sdk.NewInt(500)

	keeper.SetAccountVestingPools(ctx, accountVestingPools)

	response, err := keeper.VestingPools(wctx, &types.QueryVestingPoolsRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestingPools, time, true)

}

func TestVestingSentAfterLockEndReceivingSide(t *testing.T) {
	height := int64(10100)
	time := commontestutils.TestEnvTime.Add(testutils.CreateDurationFromNumOfHours(10100))

	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeightAndTime(t, height, time)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestingPools := testutils.GenerateOneAccountVestingPoolsWithAddressWith10BasedVestingPools(1, 1, 1)
	accountVestingPools.Address = addr
	accountVestingPools.VestingPools[0].LockStart = accountVestingPools.VestingPools[0].LockEnd
	// accountVestingPools.VestingPools[0].LastModification = accountVestingPools.VestingPools[0].LockEnd

	accountVestingPools.VestingPools[0].LockEnd = accountVestingPools.VestingPools[0].LockEnd.Add(testutils.CreateDurationFromNumOfHours(-100))

	keeper.SetAccountVestingPools(ctx, accountVestingPools)

	response, err := keeper.VestingPools(wctx, &types.QueryVestingPoolsRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestingPools, time, true)

}

func TestVestingSentAfterLockEndSendingSide(t *testing.T) {
	height := int64(10100)
	time := commontestutils.TestEnvTime.Add(testutils.CreateDurationFromNumOfHours(10100))

	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeightAndTime(t, height, time)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestingPools := testutils.GenerateOneAccountVestingPoolsWithAddressWith10BasedVestingPools(1, 1, 1)
	accountVestingPools.Address = addr

	// accountVestingPools.VestingPools[0].LastModification = accountVestingPools.VestingPools[0].LockEnd
	accountVestingPools.VestingPools[0].Sent = sdk.NewInt(100000)
	// accountVestingPools.VestingPools[0].LastModificationVested = accountVestingPools.VestingPools[0].LastModificationVested.Sub(sdk.NewInt(100000))

	accountVestingPools.VestingPools[0].LockEnd = accountVestingPools.VestingPools[0].LockEnd.Add(testutils.CreateDurationFromNumOfHours(-100))

	keeper.SetAccountVestingPools(ctx, accountVestingPools)

	response, err := keeper.VestingPools(wctx, &types.QueryVestingPoolsRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestingPools, time, true)

}

func TestVestingSentAfterLockEndSendingSideAndWithdrawn(t *testing.T) {
	height := int64(10100)
	time := commontestutils.TestEnvTime.Add(testutils.CreateDurationFromNumOfHours(10100))

	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeightAndTime(t, height, time)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestingPools := testutils.GenerateOneAccountVestingPoolsWithAddressWith10BasedVestingPools(1, 1, 1)
	accountVestingPools.Address = addr

	// accountVestingPools.VestingPools[0].LastModification = accountVestingPools.VestingPools[0].LockEnd
	accountVestingPools.VestingPools[0].Sent = sdk.NewInt(100000)
	// accountVestingPools.VestingPools[0].LastModificationVested = accountVestingPools.VestingPools[0].LastModificationVested.Sub(sdk.NewInt(100000))
	// accountVestingPools.VestingPools[0].LastModificationWithdrawn = sdk.NewInt(400)

	accountVestingPools.VestingPools[0].LockEnd = accountVestingPools.VestingPools[0].LockEnd.Add(testutils.CreateDurationFromNumOfHours(-100))

	keeper.SetAccountVestingPools(ctx, accountVestingPools)

	response, err := keeper.VestingPools(wctx, &types.QueryVestingPoolsRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestingPools, time, true)

}

func TestVestingManyVestings(t *testing.T) {
	height := int64(0)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestingPools := testutils.GenerateOneAccountVestingPoolsWithAddressWith10BasedVestingPools(3, 1, 1)
	accountVestingPools.Address = addr

	keeper.SetAccountVestingPools(ctx, accountVestingPools)

	response, err := keeper.VestingPools(wctx, &types.QueryVestingPoolsRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestingPools, ctx.BlockTime(), true)

}
