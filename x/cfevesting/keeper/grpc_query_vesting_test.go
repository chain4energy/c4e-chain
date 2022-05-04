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

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.Vestings[0].TransferAllowed = true

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestings, ctx.BlockTime(), true)
}

func TestVestingWithDelegableAddress(t *testing.T) {
	height := int64(0)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestings, ctx.BlockTime(), true)

}

func TestVestingSomeToWithdraw(t *testing.T) {
	height := int64(10100)
	time := testutils.TestEnvTime.Add(testutils.CreateDurationFromNumOfHours(10100))
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeightAndTime(t, height, time)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, ctx.BlockTime(), true)

}

func TestVestingSomeToWithdrawAndSomeWithdrawn(t *testing.T) {
	height := int64(10100)
	time := testutils.TestEnvTime.Add(testutils.CreateDurationFromNumOfHours(10100))
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeightAndTime(t, height, time)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.Vestings[0].Withdrawn = sdk.NewInt(500)
	accountVestings.Vestings[0].LastModificationWithdrawn = sdk.NewInt(500)

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestings, time, true)

}

func TestVestingSentAfterLockEndReceivingSide(t *testing.T) {
	height := int64(10100)
	time := testutils.TestEnvTime.Add(testutils.CreateDurationFromNumOfHours(10100))

	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeightAndTime(t, height, time)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.Vestings[0].LockStart = accountVestings.Vestings[0].LockEnd
	accountVestings.Vestings[0].LastModification = accountVestings.Vestings[0].LockEnd

	accountVestings.Vestings[0].LockEnd = accountVestings.Vestings[0].LockEnd.Add(testutils.CreateDurationFromNumOfHours(-100))

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, time, true)

}

func TestVestingSentAfterLockEndSendingSide(t *testing.T) {
	height := int64(10100)
	time := testutils.TestEnvTime.Add(testutils.CreateDurationFromNumOfHours(10100))

	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeightAndTime(t, height, time)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	accountVestings.Vestings[0].LastModification = accountVestings.Vestings[0].LockEnd
	accountVestings.Vestings[0].Sent = sdk.NewInt(100000)
	accountVestings.Vestings[0].LastModificationVested = accountVestings.Vestings[0].LastModificationVested.Sub(sdk.NewInt(100000))

	accountVestings.Vestings[0].LockEnd = accountVestings.Vestings[0].LockEnd.Add(testutils.CreateDurationFromNumOfHours(-100))

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, time, true)

}

func TestVestingSentAfterLockEndSendingSideAndWithdrawn(t *testing.T) {
	height := int64(10100)
	time := testutils.TestEnvTime.Add(testutils.CreateDurationFromNumOfHours(10100))

	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeightAndTime(t, height, time)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	accountVestings.Vestings[0].LastModification = accountVestings.Vestings[0].LockEnd
	accountVestings.Vestings[0].Sent = sdk.NewInt(100000)
	accountVestings.Vestings[0].LastModificationVested = accountVestings.Vestings[0].LastModificationVested.Sub(sdk.NewInt(100000))
	accountVestings.Vestings[0].LastModificationWithdrawn = sdk.NewInt(400)

	accountVestings.Vestings[0].LockEnd = accountVestings.Vestings[0].LockEnd.Add(testutils.CreateDurationFromNumOfHours(-100))

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, time, true)

}

func TestVestingManyVestings(t *testing.T) {
	height := int64(0)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(3, 1, 1)
	accountVestings.Address = addr

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, ctx.BlockTime(), true)

}
