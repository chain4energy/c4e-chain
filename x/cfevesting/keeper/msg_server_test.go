package keeper_test

import (
	"strconv"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"

	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

	"testing"

	"github.com/stretchr/testify/require"
)

const (
	vPool1 = "v-pool-1"
	vPool2 = "v-pool-2"
)

type VestingTestHelper struct {
	t *testing.T
	testhelper *testapp.TestHelper 
}

func NewVestingTestHelper(t *testing.T, testhelper *testapp.TestHelper) *VestingTestHelper{
	return &VestingTestHelper{t: t, testhelper: testhelper}
}

func (h *VestingTestHelper) SetupAccountsVestings(ctx sdk.Context, address string, numberOfVestings int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) types.AccountVestings {
	return h.SetupAccountsVestingsWithModification(ctx, func(*types.VestingPool) { /*do not modify*/ }, address, numberOfVestings, vestingAmount, withdrawnAmount)
}

func (h *VestingTestHelper) SetupAccountsVestingsWithModification(ctx sdk.Context, modifyVesting func(*types.VestingPool), address string, numberOfVestings int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) types.AccountVestings {
	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(numberOfVestings, 1, 1)
	accountVestings.Address = address

	for _, vesting := range accountVestings.VestingPools {
		vesting.Vested = vestingAmount
		vesting.Withdrawn = withdrawnAmount
		vesting.LastModificationVested = vestingAmount
		vesting.LastModificationWithdrawn = withdrawnAmount
		modifyVesting(vesting)
	}
	h.testhelper.App.CfevestingKeeper.SetAccountVestings(ctx, accountVestings)
	return accountVestings
}

func (h *VestingTestHelper) CreateVestingPool(ctx sdk.Context, address sdk.AccAddress, accountVestingsExistsBefore bool, accountVestingsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType types.VestingType, amountToVest sdk.Int, accAmountBefore sdk.Int, moduleAmountBefore sdk.Int,
	accAmountAfter sdk.Int, moduleAmountAfter sdk.Int) {

	_, accFound := h.testhelper.App.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(h.t, accountVestingsExistsBefore, accFound)

	h.testhelper.BankUtils.VerifyAccountDefultDenomBalance(ctx, address, accAmountBefore)
	h.testhelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, types.ModuleName, moduleAmountBefore)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(h.testhelper.App.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgCreateVestingPool{Creator: address.String(), Name: vestingPoolName,
		Amount: amountToVest, Duration: lockupDuration, VestingType: vestingType.Name}
	_, error := msgServer.CreateVestingPool(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, error)

	_, accFound = h.testhelper.App.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(h.t, accountVestingsExistsAfter, accFound)

	h.testhelper.BankUtils.VerifyAccountDefultDenomBalance(ctx, address, accAmountAfter)
	h.testhelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, types.ModuleName, moduleAmountAfter)
}

func newInts64Array(n int, v int64) []int64 {
	s := make([]int64, n)
	for i := range s {
		s[i] = v
	}
	return s
}

func newTimeArray(n int, v time.Time) []time.Time {
	s := make([]time.Time, n)
	for i := range s {
		s[i] = v
	}
	return s
}

func (h *VestingTestHelper) VerifyAccountVestingPools(ctx sdk.Context, address sdk.AccAddress,
	vestingNames []string, durations []time.Duration, vestingTypes []types.VestingType, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int) {

	h.VerifyAccountVestingsWithModification(ctx, address, 1, vestingNames, durations, vestingTypes, newTimeArray(len(vestingTypes), ctx.BlockTime()), vestedAmounts, withdrawnAmounts,
		newInts64Array(len(vestingTypes), 0), newTimeArray(len(vestingTypes), ctx.BlockTime()), vestedAmounts, withdrawnAmounts)
}

func (h *VestingTestHelper) VerifyAccountVestingsWithModification(ctx sdk.Context, address sdk.AccAddress,
	amountOfAllAccVestings int, vestingNames []string, durations []time.Duration, vestingTypes []types.VestingType, startsTimes []time.Time, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int,
	sentAmounts []int64, modificationsTimes []time.Time, modificationsVested []sdk.Int, modificationsWithdrawn []sdk.Int) {
	allAccVestings := h.testhelper.App.CfevestingKeeper.GetAllAccountVestings(ctx)

	accVestings, accFound := h.testhelper.App.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(h.t, true, accFound)

	require.EqualValues(h.t, amountOfAllAccVestings, len(allAccVestings))
	require.EqualValues(h.t, len(vestingTypes), len(accVestings.VestingPools))

	require.EqualValues(h.t, address.String(), accVestings.Address)

	for i, vesting := range accVestings.VestingPools {
		found := false
		if vesting.Id == int32(i+1) {
			require.EqualValues(h.t, i+1, vesting.Id)
			require.EqualValues(h.t, vestingNames[i], vesting.Name)
			require.EqualValues(h.t, vestingTypes[i].Name, vesting.VestingType)
			require.EqualValues(h.t, true, startsTimes[i].Equal(vesting.LockStart))
			require.EqualValues(h.t, true, ctx.BlockTime().Add(durations[i]).Equal(vesting.LockEnd))
			require.EqualValues(h.t, vestedAmounts[i], vesting.Vested)
			require.EqualValues(h.t, withdrawnAmounts[i], vesting.Withdrawn)

			require.EqualValues(h.t, sdk.NewInt(sentAmounts[i]), vesting.Sent)
			require.EqualValues(h.t, true, modificationsTimes[i].Equal(vesting.LastModification))
			require.EqualValues(h.t, modificationsVested[i], vesting.LastModificationVested)
			require.EqualValues(h.t, modificationsWithdrawn[i], vesting.LastModificationWithdrawn)
			found = true

		}
		require.True(h.t, found, "not found vesting id: "+strconv.Itoa(i+1))

	}

}

func (h *VestingTestHelper) SetupVestingTypes(ctx sdk.Context, numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) types.VestingTypes {
	return h.SetupVestingTypesWithModification(ctx, func(*types.VestingType) { /* do not modify */ }, numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *VestingTestHelper) SetupVestingTypesWithModification(ctx sdk.Context, modifyVestingType func(*types.VestingType), numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) types.VestingTypes {
	vestingTypesArray := testutils.Generate10BasedVestingTypes(numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
	for _, vestingType := range vestingTypesArray {
		modifyVestingType(vestingType)
	}
	vestingTypes := types.VestingTypes{VestingTypes: vestingTypesArray}
	h.testhelper.App.CfevestingKeeper.SetVestingTypes(ctx, vestingTypes)
	return vestingTypes
}

func (h *VestingTestHelper) WithdrawAllAvailable(ctx sdk.Context, address sdk.AccAddress, accountBalanceBefore sdk.Int, moduleBalanceBefore sdk.Int,
	accountBalanceAfter sdk.Int, moduleBalanceAfter sdk.Int) {

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(h.testhelper.App.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	h.testhelper.BankUtils.VerifyAccountDefultDenomBalance(ctx, address, accountBalanceBefore)
	h.testhelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, types.ModuleName, moduleBalanceBefore)

	msg := types.MsgWithdrawAllAvailable{Creator: address.String()}
	_, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)
	h.testhelper.BankUtils.VerifyAccountDefultDenomBalance(ctx, address, accountBalanceAfter)
	h.testhelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, types.ModuleName, moduleBalanceAfter)
}

func (h *VestingTestHelper) CompareStoredAcountVestings(ctx sdk.Context, address sdk.AccAddress, accVestings types.AccountVestings) {
	storedAccVestings, accFound := h.testhelper.App.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(h.t, true, accFound)

	testutils.AssertAccountVestings(h.t, accVestings, storedAccVestings)
}

func verifyVestingResponse(t *testing.T, response *types.QueryVestingPoolsResponse, accVestings types.AccountVestings, current time.Time, delegationAllowed bool) {
	require.EqualValues(t, len(accVestings.VestingPools), len(response.VestingPools))

	for _, vesting := range accVestings.VestingPools {
		found := false
		for _, vestingInfo := range response.VestingPools {
			if vesting.Id == vestingInfo.Id {
				require.EqualValues(t, vesting.VestingType, vestingInfo.VestingType)
				require.EqualValues(t, vesting.Name, vestingInfo.Name)
				require.EqualValues(t, testutils.GetExpectedWithdrawableForVesting(*vesting, current).String(), response.VestingPools[0].Withdrawable)
				require.EqualValues(t, true, vesting.LockStart.Equal(vestingInfo.LockStart))
				require.EqualValues(t, true, vesting.LockEnd.Equal(vestingInfo.LockEnd))
				require.EqualValues(t, commontestutils.DefaultTestDenom, response.VestingPools[0].Vested.Denom)
				require.EqualValues(t, vesting.Vested, response.VestingPools[0].Vested.Amount)
				require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn).String(), response.VestingPools[0].CurrentVestedAmount)
				require.EqualValues(t, vesting.Sent.String(), response.VestingPools[0].SentAmount)

				found = true
			}
		}
		require.True(t, found, "not found vesting nfo with Id: "+strconv.FormatInt(int64(vesting.Id), 10))
	}
}
