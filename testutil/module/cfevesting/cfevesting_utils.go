package cfevesting

import (
	"context"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/testutil/nullify"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting"
	cfevestingmodulekeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"

	"github.com/stretchr/testify/require"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

type C4eVestingKeeperUtils struct {
	t                      *testing.T
	helperCfevestingKeeper *cfevestingmodulekeeper.Keeper
}

func NewC4eVestingKeeperUtils(t *testing.T, helperCfevestingKeeper *cfevestingmodulekeeper.Keeper) C4eVestingKeeperUtils {
	return C4eVestingKeeperUtils{t: t, helperCfevestingKeeper: helperCfevestingKeeper}
}

func (h *C4eVestingKeeperUtils) SetupAccountsVestings(ctx sdk.Context, address string, numberOfVestings int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) cfevestingtypes.AccountVestings {
	return h.SetupAccountsVestingsWithModification(ctx, func(*cfevestingtypes.VestingPool) { /*do not modify*/ }, address, numberOfVestings, vestingAmount, withdrawnAmount)
}

func (h *C4eVestingKeeperUtils) SetupAccountsVestingsWithModification(ctx sdk.Context, modifyVesting func(*cfevestingtypes.VestingPool), address string, numberOfVestings int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) cfevestingtypes.AccountVestings {
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

func (h *C4eVestingKeeperUtils) CheckNonNegativeVestingPoolAmountsInvariant(ctx sdk.Context, failed bool, message string) {
	invariant := cfevestingmodulekeeper.NonNegativeVestingPoolAmountsInvariant(*h.helperCfevestingKeeper)
	commontestutils.CheckInvariant(h.t, ctx, invariant, failed, message)
}

func (h *C4eVestingKeeperUtils) CheckVestingPoolConsistentDataInvariant(ctx sdk.Context, failed bool, message string) {
	invariant := cfevestingmodulekeeper.VestingPoolConsistentDataInvariant(*h.helperCfevestingKeeper)
	commontestutils.CheckInvariant(h.t, ctx, invariant, failed, message)
}

func (h *C4eVestingKeeperUtils) CheckModuleAccountInvariant(ctx sdk.Context, failed bool, message string) {
	invariant := cfevestingmodulekeeper.ModuleAccountInvariant(*h.helperCfevestingKeeper)
	commontestutils.CheckInvariant(h.t, ctx, invariant, failed, message)
}

type C4eVestingUtils struct {
	C4eVestingKeeperUtils
	helperAccountKeeper *authkeeper.AccountKeeper
	helperBankKeeper    *bankkeeper.Keeper
	helperStakingKeeper *stakingkeeper.Keeper
	bankUtils           *commontestutils.BankUtils
	authUtils           *commontestutils.AuthUtils
}

func NewC4eVestingUtils(t *testing.T, helperCfevestingKeeper *cfevestingmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	helperBankKeeper *bankkeeper.Keeper,
	helperStakingKeeper *stakingkeeper.Keeper, bankUtils *commontestutils.BankUtils,
	authUtils *commontestutils.AuthUtils) C4eVestingUtils {
	return C4eVestingUtils{C4eVestingKeeperUtils: NewC4eVestingKeeperUtils(t, helperCfevestingKeeper), helperAccountKeeper: helperAccountKeeper,
		helperBankKeeper: helperBankKeeper, helperStakingKeeper: helperStakingKeeper, bankUtils: bankUtils, authUtils: authUtils}
}

func (h *C4eVestingUtils) MessageCreateVestingPool(ctx sdk.Context, address sdk.AccAddress, accountVestingsExistsBefore bool, accountVestingsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest sdk.Int, accAmountBefore sdk.Int, moduleAmountBefore sdk.Int,
	accAmountAfter sdk.Int, moduleAmountAfter sdk.Int, vestingPollsIds ...int) {
	_, accFound := h.helperCfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(h.t, accountVestingsExistsBefore, accFound)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accAmountBefore)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleAmountBefore)

	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgCreateVestingPool{Creator: address.String(), Name: vestingPoolName,
		Amount: amountToVest, Duration: lockupDuration, VestingType: vestingType.Name}
	_, error := msgServer.CreateVestingPool(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, error)

	accVestingPools, accFound := h.helperCfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(h.t, accountVestingsExistsAfter, accFound)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accAmountAfter)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleAmountAfter)

	if vestingPollsIds != nil {
		require.EqualValues(h.t, len(vestingPollsIds), len(accVestingPools.VestingPools))
		for i, v := range vestingPollsIds {
			vestingPool := accVestingPools.VestingPools[i]
			require.EqualValues(h.t, v, vestingPool.Id)
		}

	}
}

func (h *C4eVestingUtils) MessageCreateVestingPoolError(ctx sdk.Context, address sdk.AccAddress,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest sdk.Int,
	errorMessage string) {

	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgCreateVestingPool{Creator: address.String(), Name: vestingPoolName,
		Amount: amountToVest, Duration: lockupDuration, VestingType: vestingType.Name}
	_, err := msgServer.CreateVestingPool(msgServerCtx, &msg)

	require.EqualError(h.t, err, errorMessage)

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
	vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int, startAndModificationTime ...time.Time) {

	// startAndModificationTime allows to handle 3 cases - time from context, one time for all
	var times []time.Time
	if len(startAndModificationTime) == 0 {
		times = newTimeArray(len(vestingTypes), ctx.BlockTime())
	} else if len(startAndModificationTime) == 1 {
		times = newTimeArray(len(vestingTypes), startAndModificationTime[0])
	} else {
		times = startAndModificationTime
	}

	h.VerifyAccountVestingPoolsWithModification(ctx, address, 1, vestingNames, durations, vestingTypes, times, vestedAmounts, withdrawnAmounts,
		newInts64Array(len(vestingTypes), 0), times, vestedAmounts, withdrawnAmounts)
}

func (h *C4eVestingUtils) VerifyAccountVestingPoolsWithModification(ctx sdk.Context, address sdk.AccAddress,
	amountOfAllAccVestings int, vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, startsTimes []time.Time, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int,
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
			require.EqualValues(h.t, true, startsTimes[i].Add(durations[i]).Equal(vesting.LockEnd))
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

func (h *C4eVestingUtils) SetupVestingTypes(ctx sdk.Context, numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) cfevestingtypes.VestingTypes {
	return h.SetupVestingTypesWithModification(ctx, func(*cfevestingtypes.VestingType) { /* do not modify */ }, numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *C4eVestingUtils) SetupVestingTypesWithModification(ctx sdk.Context, modifyVestingType func(*cfevestingtypes.VestingType), numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) cfevestingtypes.VestingTypes {
	vestingTypesArray := Generate10BasedVestingTypes(numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
	for _, vestingType := range vestingTypesArray {
		modifyVestingType(vestingType)
	}
	vestingTypes := cfevestingtypes.VestingTypes{VestingTypes: vestingTypesArray}
	h.helperCfevestingKeeper.SetVestingTypes(ctx, vestingTypes)
	return vestingTypes
}

func (h *C4eVestingUtils) MessageWithdrawAllAvailable(ctx sdk.Context, address sdk.AccAddress, accountBalanceBefore sdk.Int, moduleBalanceBefore sdk.Int,
	accountBalanceAfter sdk.Int, moduleBalanceAfter sdk.Int) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accountBalanceBefore)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleBalanceBefore)

	msg := cfevestingtypes.MsgWithdrawAllAvailable{Creator: address.String()}
	_, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)
	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accountBalanceAfter)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleBalanceAfter)
}

func (h *C4eVestingUtils) MessageWithdrawAllAvailableError(ctx sdk.Context, address string, errorMessage string) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgWithdrawAllAvailable{Creator: address}
	_, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)

	require.EqualError(h.t, err, errorMessage)
}

func (h *C4eVestingUtils) CompareStoredAcountVestings(ctx sdk.Context, address sdk.AccAddress, accVestings cfevestingtypes.AccountVestings) {
	storedAccVestings, accFound := h.helperCfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(h.t, true, accFound)

	AssertAccountVestings(h.t, accVestings, storedAccVestings)
}

func (h *C4eVestingUtils) InitGenesis(ctx sdk.Context, genState cfevestingtypes.GenesisState) {
	cfevesting.InitGenesis(ctx, *h.helperCfevestingKeeper, genState, h.helperAccountKeeper, *h.helperBankKeeper, h.helperStakingKeeper)
}

func (h *C4eVestingUtils) InitGenesisError(ctx sdk.Context, genState cfevestingtypes.GenesisState, errorMessage string) {
	require.PanicsWithError(h.t, errorMessage,
		func() {
			cfevesting.InitGenesis(ctx, *h.helperCfevestingKeeper, genState, h.helperAccountKeeper, *h.helperBankKeeper, h.helperStakingKeeper)
		}, "")
}

func (h *C4eVestingUtils) ExportGenesis(ctx sdk.Context, expected cfevestingtypes.GenesisState) {
	got := cfevesting.ExportGenesis(ctx, *h.helperCfevestingKeeper)
	require.NotNil(h.t, got)
	// require.EqualValues(h.t, expected, *got)
	require.EqualValues(h.t, expected.Params, got.GetParams())
	require.EqualValues(h.t, expected.VestingTypes, (*got).VestingTypes)
	require.EqualValues(h.t, len(expected.AccountVestingsList.Vestings), len((*got).AccountVestingsList.Vestings))
	AssertAccountVestingsArrays(h.t, expected.AccountVestingsList.Vestings, (*got).AccountVestingsList.Vestings)
	require.EqualValues(h.t, expected.VestingAccountCount, (*got).VestingAccountCount)
	require.ElementsMatch(h.t, expected.VestingAccountList, (*got).VestingAccountList)

	nullify.Fill(&expected)
	nullify.Fill(got)
}

func (h *C4eVestingUtils) QueryVestings(wctx context.Context, expectedResponse cfevestingtypes.QueryVestingsResponse) {
	resp, err := h.helperCfevestingKeeper.Vestings(wctx, &cfevestingtypes.QueryVestingsRequest{})
	require.NoError(h.t, err)
	require.Equal(h.t, expectedResponse, *resp)
}

func (h *C4eVestingUtils) SetVestingTypes(ctx sdk.Context, vestingTypes cfevestingtypes.VestingTypes) {
	h.helperCfevestingKeeper.SetVestingTypes(ctx, vestingTypes)
}

func (h *C4eVestingUtils) MessageSendToVestingAccount(ctx sdk.Context, fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolId int32, amount sdk.Int, restartVesting bool) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	vestingAccountCount := h.helperCfevestingKeeper.GetVestingAccountCount(ctx)

	msg := cfevestingtypes.MsgSendToVestingAccount{FromAddress: fromAddress.String(), ToAddress: vestingAccAddress.String(),
		VestingId: vestingPoolId, Amount: amount, RestartVesting: restartVesting}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, vestingAccAddress, sdk.NewInt(100))

	require.Equal(h.t, uint64(vestingAccountCount+1), h.helperCfevestingKeeper.GetVestingAccountCount(ctx))
	vaccFromList, found := h.helperCfevestingKeeper.GetVestingAccount(ctx, uint64(vestingAccountCount))
	require.Equal(h.t, true, found)
	require.Equal(h.t, vestingAccAddress.String(), vaccFromList.Address)

	vestingPools, found := h.helperCfevestingKeeper.GetAccountVestings(ctx, fromAddress.String())
	require.Equal(h.t, true, found)

	var foundVPool *cfevestingtypes.VestingPool = nil

	for _, vPool := range vestingPools.VestingPools {
		if vPool.Id == vestingPoolId {
			foundVPool = vPool
			break
		}
	}
	require.NotNilf(h.t, foundVPool, "vesting pool no found. Id: %d", vestingPoolId)

	vestingType, err := h.helperCfevestingKeeper.GetVestingType(ctx, foundVPool.VestingType)
	require.NoError(h.t, err, "GetVestingType error")

	denom := h.helperCfevestingKeeper.Denom(ctx)
	h.authUtils.VerifyVestingAccount(ctx, vestingAccAddress, denom, sdk.NewInt(100), ctx.BlockTime().Add(vestingType.LockupPeriod), ctx.BlockTime().Add(vestingType.LockupPeriod).Add(vestingType.VestingPeriod))
}

func (h *C4eVestingUtils) MessageSendToVestingAccountError(ctx sdk.Context, fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolId int32, amount sdk.Int, restartVesting bool, errorMessage string) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	vestingAccountCount := h.helperCfevestingKeeper.GetVestingAccountCount(ctx)

	msg := cfevestingtypes.MsgSendToVestingAccount{FromAddress: fromAddress.String(), ToAddress: vestingAccAddress.String(),
		VestingId: vestingPoolId, Amount: amount, RestartVesting: restartVesting}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualError(h.t, err, errorMessage)
	require.Equal(h.t, uint64(vestingAccountCount), h.helperCfevestingKeeper.GetVestingAccountCount(ctx))
}

type ContextC4eVestingUtils struct {
	C4eVestingUtils
	testContext commontestutils.TestContext
}

func NewContextC4eVestingUtils(t *testing.T, testContext commontestutils.TestContext, helperCfevestingKeeper *cfevestingmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	helperBankKeeper *bankkeeper.Keeper,
	helperStakingKeeper *stakingkeeper.Keeper, bankUtils *commontestutils.BankUtils,
	authUtils *commontestutils.AuthUtils) *ContextC4eVestingUtils {
	c4eVestingUtils := NewC4eVestingUtils(t, helperCfevestingKeeper, helperAccountKeeper, helperBankKeeper, helperStakingKeeper, bankUtils, authUtils)
	return &ContextC4eVestingUtils{C4eVestingUtils: c4eVestingUtils, testContext: testContext}
}

func (h *ContextC4eVestingUtils) SetupAccountsVestings(address string, numberOfVestings int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) cfevestingtypes.AccountVestings {
	return h.C4eVestingUtils.SetupAccountsVestings(h.testContext.GetContext(), address, numberOfVestings, vestingAmount, withdrawnAmount)
}

func (h *ContextC4eVestingUtils) SetupAccountsVestingsWithModification(modifyVesting func(*cfevestingtypes.VestingPool), address string, numberOfVestings int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) cfevestingtypes.AccountVestings {
	return h.C4eVestingUtils.SetupAccountsVestingsWithModification(h.testContext.GetContext(), modifyVesting, address, numberOfVestings, vestingAmount, withdrawnAmount)
}

func (h *ContextC4eVestingUtils) MessageCreateVestingPool(address sdk.AccAddress, accountVestingsExistsBefore bool, accountVestingsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest sdk.Int, accAmountBefore sdk.Int, moduleAmountBefore sdk.Int,
	accAmountAfter sdk.Int, moduleAmountAfter sdk.Int, vestingPollsIds ...int) {
	h.C4eVestingUtils.MessageCreateVestingPool(h.testContext.GetContext(), address, accountVestingsExistsBefore, accountVestingsExistsAfter, vestingPoolName, lockupDuration,
		vestingType, amountToVest, accAmountBefore, moduleAmountBefore, accAmountAfter, moduleAmountAfter, vestingPollsIds...)
}

func (h *ContextC4eVestingUtils) VerifyAccountVestingPools(address sdk.AccAddress,
	vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int, startAndModificationTime ...time.Time) {
	h.C4eVestingUtils.VerifyAccountVestingPools(h.testContext.GetContext(), address, vestingNames, durations, vestingTypes, vestedAmounts, withdrawnAmounts, startAndModificationTime...)
}

func (h *ContextC4eVestingUtils) VerifyAccountVestingPoolsWithModification(address sdk.AccAddress,
	amountOfAllAccVestings int, vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, startsTimes []time.Time, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int,
	sentAmounts []int64, modificationsTimes []time.Time, modificationsVested []sdk.Int, modificationsWithdrawn []sdk.Int) {
	h.C4eVestingUtils.VerifyAccountVestingPoolsWithModification(h.testContext.GetContext(), address, amountOfAllAccVestings, vestingNames, durations, vestingTypes, startsTimes, vestedAmounts,
		withdrawnAmounts, sentAmounts, modificationsTimes, modificationsVested, modificationsWithdrawn)
}

func (h *ContextC4eVestingUtils) SetupVestingTypes(numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) cfevestingtypes.VestingTypes {
	return h.C4eVestingUtils.SetupVestingTypes(h.testContext.GetContext(), numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *ContextC4eVestingUtils) SetupVestingTypesWithModification(modifyVestingType func(*cfevestingtypes.VestingType), numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) cfevestingtypes.VestingTypes {
	return h.C4eVestingUtils.SetupVestingTypesWithModification(h.testContext.GetContext(), modifyVestingType, numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *ContextC4eVestingUtils) MessageWithdrawAllAvailable(address sdk.AccAddress, accountBalanceBefore sdk.Int, moduleBalanceBefore sdk.Int,
	accountBalanceAfter sdk.Int, moduleBalanceAfter sdk.Int) {
	h.C4eVestingUtils.MessageWithdrawAllAvailable(h.testContext.GetContext(), address, accountBalanceBefore, moduleBalanceBefore, accountBalanceAfter, moduleBalanceAfter)
}

func (h *ContextC4eVestingUtils) CompareStoredAcountVestings(address sdk.AccAddress, accVestings cfevestingtypes.AccountVestings) {
	h.C4eVestingUtils.CompareStoredAcountVestings(h.testContext.GetContext(), address, accVestings)
}

func (h *ContextC4eVestingUtils) InitGenesis(genState cfevestingtypes.GenesisState) {
	h.C4eVestingUtils.InitGenesis(h.testContext.GetContext(), genState)
}

func (h *ContextC4eVestingUtils) QueryVestings(expectedResponse cfevestingtypes.QueryVestingsResponse) {
	h.C4eVestingUtils.QueryVestings(h.testContext.GetWrappedContext(), expectedResponse)
}

func (h *ContextC4eVestingUtils) MessageCreateVestingPoolError(address sdk.AccAddress,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest sdk.Int,
	errorMessage string) {
	h.C4eVestingUtils.MessageCreateVestingPoolError(h.testContext.GetContext(), address, vestingPoolName, lockupDuration,
		vestingType, amountToVest, errorMessage)
}

func (h *ContextC4eVestingUtils) SetVestingTypes(vestingTypes cfevestingtypes.VestingTypes) {
	h.C4eVestingUtils.SetVestingTypes(h.testContext.GetContext(), vestingTypes)
}

func (h *ContextC4eVestingUtils) MessageSendToVestingAccount(fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolId int32, amount sdk.Int, restartVesting bool) {
	h.C4eVestingUtils.MessageSendToVestingAccount(h.testContext.GetContext(), fromAddress, vestingAccAddress, vestingPoolId, amount, restartVesting)
}

func (h *ContextC4eVestingUtils) MessageSendToVestingAccountError(fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolId int32, amount sdk.Int, restartVesting bool, errorMessage string) {
	h.C4eVestingUtils.MessageSendToVestingAccountError(h.testContext.GetContext(), fromAddress, vestingAccAddress, vestingPoolId, amount, restartVesting, errorMessage)
}

func (h *ContextC4eVestingUtils) MessageWithdrawAllAvailableError(address string, errorMessage string) {
	h.C4eVestingUtils.MessageWithdrawAllAvailableError(h.testContext.GetContext(), address, errorMessage)
}

func (h *ContextC4eVestingUtils) ExportGenesis(expected cfevestingtypes.GenesisState) {
	h.C4eVestingUtils.ExportGenesis(h.testContext.GetContext(), expected)
}

func (h *ContextC4eVestingUtils) InitGenesisError(genState cfevestingtypes.GenesisState, errorMessage string) {
	h.C4eVestingUtils.InitGenesisError(h.testContext.GetContext(), genState, errorMessage)
}

func (h *ContextC4eVestingUtils) CheckModuleAccountInvariant(failed bool, message string) {
	h.C4eVestingUtils.CheckModuleAccountInvariant(h.testContext.GetContext(), failed, message)
}
