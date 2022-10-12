package cfevesting

import (
	"fmt"
	"strconv"
	"time"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"testing"

	cfevestingmodulekeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/stretchr/testify/require"
)

type C4eVestingUtils struct {
	t                      *testing.T
	helperCfevestingKeeper *cfevestingmodulekeeper.Keeper
	bankUtils              *commontestutils.BankUtils
}

func NewC4eVestingUtils(t *testing.T, helperCfevestingKeeper *cfevestingmodulekeeper.Keeper, bankUtils *commontestutils.BankUtils) C4eVestingUtils {
	return C4eVestingUtils{t: t, helperCfevestingKeeper: helperCfevestingKeeper, bankUtils: bankUtils}
}

func (h *C4eVestingUtils) SetupAccountsVestings(ctx sdk.Context, address string, numberOfVestings int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) types.AccountVestings {
	return h.SetupAccountsVestingsWithModification(ctx, func(*types.VestingPool) { /*do not modify*/ }, address, numberOfVestings, vestingAmount, withdrawnAmount)
}

func (h *C4eVestingUtils) SetupAccountsVestingsWithModification(ctx sdk.Context, modifyVesting func(*types.VestingPool), address string, numberOfVestings int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) types.AccountVestings {
	accountVestings := GenerateOneAccountVestingsWithAddressWith10BasedVestings(numberOfVestings, 1, 1)
	accountVestings.Address = address

	for _, vesting := range accountVestings.VestingPools {
		vesting.Vested = vestingAmount
		vesting.Withdrawn = withdrawnAmount
		vesting.LastModificationVested = vestingAmount
		vesting.LastModificationWithdrawn = withdrawnAmount
		modifyVesting(vesting)
	}
	h.helperCfevestingKeeper.SetAccountVestings(ctx, accountVestings)
	return accountVestings
}

func (h *C4eVestingUtils) CreateVestingPool(ctx sdk.Context, address sdk.AccAddress, accountVestingsExistsBefore bool, accountVestingsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType types.VestingType, amountToVest sdk.Int, accAmountBefore sdk.Int, moduleAmountBefore sdk.Int,
	accAmountAfter sdk.Int, moduleAmountAfter sdk.Int) {
	_, accFound := h.helperCfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(h.t, accountVestingsExistsBefore, accFound)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accAmountBefore)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, types.ModuleName, moduleAmountBefore)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgCreateVestingPool{Creator: address.String(), Name: vestingPoolName,
		Amount: amountToVest, Duration: lockupDuration, VestingType: vestingType.Name}
	_, error := msgServer.CreateVestingPool(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, error)

	_, accFound = h.helperCfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(h.t, accountVestingsExistsAfter, accFound)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accAmountAfter)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, types.ModuleName, moduleAmountAfter)
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

func (h *C4eVestingUtils) VerifyAccountVestingPools(ctx sdk.Context, address sdk.AccAddress,
	vestingNames []string, durations []time.Duration, vestingTypes []types.VestingType, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int) {
	blockTime := ctx.BlockTime()
	h.VerifyAccountVestingsWithModification(ctx, address, 1, vestingNames, durations, vestingTypes, newTimeArray(len(vestingTypes), blockTime), vestedAmounts, withdrawnAmounts,
		newInts64Array(len(vestingTypes), 0), newTimeArray(len(vestingTypes), blockTime), vestedAmounts, withdrawnAmounts)
}

func (h *C4eVestingUtils) VerifyAccountVestingsWithModification(ctx sdk.Context, address sdk.AccAddress,
	amountOfAllAccVestings int, vestingNames []string, durations []time.Duration, vestingTypes []types.VestingType, startsTimes []time.Time, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int,
	sentAmounts []int64, modificationsTimes []time.Time, modificationsVested []sdk.Int, modificationsWithdrawn []sdk.Int) {

	allAccVestings := h.helperCfevestingKeeper.GetAllAccountVestings(ctx)

	accVestings, accFound := h.helperCfevestingKeeper.GetAccountVestings(ctx, address.String())
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

func (h *C4eVestingUtils) SetupVestingTypes(ctx sdk.Context, numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) types.VestingTypes {
	return h.SetupVestingTypesWithModification(ctx, func(*types.VestingType) { /* do not modify */ }, numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *C4eVestingUtils) SetupVestingTypesWithModification(ctx sdk.Context, modifyVestingType func(*types.VestingType), numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) types.VestingTypes {
	vestingTypesArray := Generate10BasedVestingTypes(numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
	for _, vestingType := range vestingTypesArray {
		modifyVestingType(vestingType)
	}
	vestingTypes := types.VestingTypes{VestingTypes: vestingTypesArray}
	h.helperCfevestingKeeper.SetVestingTypes(ctx, vestingTypes)
	return vestingTypes
}

func (h *C4eVestingUtils) WithdrawAllAvailable(ctx sdk.Context, address sdk.AccAddress, accountBalanceBefore sdk.Int, moduleBalanceBefore sdk.Int,
	accountBalanceAfter sdk.Int, moduleBalanceAfter sdk.Int) {
	fmt.Printf("CCCCCCCCCC: %d\r\n", ctx.BlockHeight())
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accountBalanceBefore)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, types.ModuleName, moduleBalanceBefore)

	msg := types.MsgWithdrawAllAvailable{Creator: address.String()}
	_, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)
	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accountBalanceAfter)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, types.ModuleName, moduleBalanceAfter)
}

func (h *C4eVestingUtils) CompareStoredAcountVestings(ctx sdk.Context, address sdk.AccAddress, accVestings types.AccountVestings) {
	storedAccVestings, accFound := h.helperCfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(h.t, true, accFound)

	AssertAccountVestings(h.t, accVestings, storedAccVestings)
}

type ContextC4eVestingUtils struct {
	C4eVestingUtils
	testContext commontestutils.TestContext
}

func NewContextC4eVestingUtils(t *testing.T, testContext commontestutils.TestContext, helperCfevestingKeeper *cfevestingmodulekeeper.Keeper, bankUtils *commontestutils.BankUtils) *ContextC4eVestingUtils {
	c4eVestingUtilsUtils := NewC4eVestingUtils(t, helperCfevestingKeeper, bankUtils)
	return &ContextC4eVestingUtils{C4eVestingUtils: c4eVestingUtilsUtils, testContext: testContext}
}

func (h *ContextC4eVestingUtils) SetupAccountsVestings(address string, numberOfVestings int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) types.AccountVestings {
	return h.C4eVestingUtils.SetupAccountsVestings(h.testContext.GetContext(), address, numberOfVestings, vestingAmount, withdrawnAmount)
}

func (h *ContextC4eVestingUtils) SetupAccountsVestingsWithModification(modifyVesting func(*types.VestingPool), address string, numberOfVestings int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) types.AccountVestings {
	return h.C4eVestingUtils.SetupAccountsVestingsWithModification(h.testContext.GetContext(), modifyVesting, address, numberOfVestings, vestingAmount, withdrawnAmount)
}

func (h *ContextC4eVestingUtils) CreateVestingPool(address sdk.AccAddress, accountVestingsExistsBefore bool, accountVestingsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType types.VestingType, amountToVest sdk.Int, accAmountBefore sdk.Int, moduleAmountBefore sdk.Int,
	accAmountAfter sdk.Int, moduleAmountAfter sdk.Int) {
	h.C4eVestingUtils.CreateVestingPool(h.testContext.GetContext(), address, accountVestingsExistsBefore, accountVestingsExistsAfter, vestingPoolName, lockupDuration,
		vestingType, amountToVest, accAmountBefore, moduleAmountBefore, accAmountAfter, moduleAmountAfter)
}

func (h *ContextC4eVestingUtils) VerifyAccountVestingPools(address sdk.AccAddress,
	vestingNames []string, durations []time.Duration, vestingTypes []types.VestingType, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int) {
	h.C4eVestingUtils.VerifyAccountVestingPools(h.testContext.GetContext(), address, vestingNames, durations, vestingTypes, vestedAmounts, withdrawnAmounts)
}

func (h *ContextC4eVestingUtils) VerifyAccountVestingsWithModification(address sdk.AccAddress,
	amountOfAllAccVestings int, vestingNames []string, durations []time.Duration, vestingTypes []types.VestingType, startsTimes []time.Time, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int,
	sentAmounts []int64, modificationsTimes []time.Time, modificationsVested []sdk.Int, modificationsWithdrawn []sdk.Int) {
	h.C4eVestingUtils.VerifyAccountVestingsWithModification(h.testContext.GetContext(), address, amountOfAllAccVestings, vestingNames, durations, vestingTypes, startsTimes, vestedAmounts,
		withdrawnAmounts, sentAmounts, modificationsTimes, modificationsVested, modificationsWithdrawn)
}

func (h *ContextC4eVestingUtils) SetupVestingTypes(numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) types.VestingTypes {
	return h.C4eVestingUtils.SetupVestingTypes(h.testContext.GetContext(), numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *ContextC4eVestingUtils) SetupVestingTypesWithModification(modifyVestingType func(*types.VestingType), numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) types.VestingTypes {
	return h.C4eVestingUtils.SetupVestingTypesWithModification(h.testContext.GetContext(), modifyVestingType, numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *ContextC4eVestingUtils) WithdrawAllAvailable(address sdk.AccAddress, accountBalanceBefore sdk.Int, moduleBalanceBefore sdk.Int,
	accountBalanceAfter sdk.Int, moduleBalanceAfter sdk.Int) {
	h.C4eVestingUtils.WithdrawAllAvailable(h.testContext.GetContext(), address, accountBalanceBefore, moduleBalanceBefore, accountBalanceAfter, moduleBalanceAfter)
}

func (h *ContextC4eVestingUtils) CompareStoredAcountVestings(address sdk.AccAddress, accVestings types.AccountVestings) {
	h.C4eVestingUtils.CompareStoredAcountVestings(h.testContext.GetContext(), address, accVestings)
}
