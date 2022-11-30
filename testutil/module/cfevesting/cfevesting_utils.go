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

func (h *C4eVestingKeeperUtils) GetC4eVestingKeeper() *cfevestingmodulekeeper.Keeper {
	return h.helperCfevestingKeeper
}

func (h *C4eVestingKeeperUtils) SetupAccountVestingPools(ctx sdk.Context, address string, numberOfVestingPools int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) cfevestingtypes.AccountVestingPools {
	return h.SetupAccountVestingPoolsWithModification(ctx, func(*cfevestingtypes.VestingPool) { /*do not modify*/ }, address, numberOfVestingPools, vestingAmount, withdrawnAmount)
}

func (h *C4eVestingKeeperUtils) SetupAccountVestingPoolsWithModification(ctx sdk.Context, modifyVesting func(*cfevestingtypes.VestingPool), address string, numberOfVestingPools int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) cfevestingtypes.AccountVestingPools {
	accountVestingPools := GenerateOneAccountVestingPoolsWithAddressWith10BasedVestingPools(numberOfVestingPools, 1, 1)
	accountVestingPools.Address = address

	for _, vesting := range accountVestingPools.VestingPools {
		vesting.InitiallyLocked = vestingAmount
		vesting.Withdrawn = withdrawnAmount
		modifyVesting(vesting)
	}
	h.helperCfevestingKeeper.SetAccountVestingPools(ctx, accountVestingPools)
	return accountVestingPools
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

func (h *C4eVestingUtils) SetupVestingTypesForAccountsVestingPools(ctx sdk.Context) {
	accountVestingPools := h.helperCfevestingKeeper.GetAllAccountVestingPools(ctx)
	vestingTypes := cfevestingtypes.VestingTypes{VestingTypes: GenerateVestingTypesForAccountVestingPools(accountVestingPools)}
	h.helperCfevestingKeeper.SetVestingTypes(ctx, vestingTypes)
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

func (h *C4eVestingUtils) MessageCreateVestingPool(ctx sdk.Context, address sdk.AccAddress, accountVestingPoolsExistsBefore bool, accountVestingPoolsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest sdk.Int, accAmountBefore sdk.Int, moduleAmountBefore sdk.Int,
	accAmountAfter sdk.Int, moduleAmountAfter sdk.Int) {
	_, accFound := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, address.String())
	require.EqualValues(h.t, accountVestingPoolsExistsBefore, accFound)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accAmountBefore)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleAmountBefore)

	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgCreateVestingPool{Creator: address.String(), Name: vestingPoolName,
		Amount: amountToVest, Duration: lockupDuration, VestingType: vestingType.Name}
	_, err := msgServer.CreateVestingPool(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)

	accVestingPools, accFound := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, address.String())
	require.EqualValues(h.t, accountVestingPoolsExistsAfter, accFound)

	if accFound {
		var vestingPool *cfevestingtypes.VestingPool = nil
		for _, vest := range accVestingPools.VestingPools {
			if vest.Name == vestingPoolName {
				vestingPool = vest
			}
		}
		require.NotNil(h.t, vestingPool)
		h.VerifyVestingPool(ctx, vestingPool, vestingPoolName, vestingType.Name, ctx.BlockTime(), ctx.BlockTime().Add(lockupDuration),
			amountToVest, sdk.ZeroInt(), sdk.ZeroInt())
	}
	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accAmountAfter)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleAmountAfter)
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
	amountOfAllAccVestingPools int, vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, startsTimes []time.Time, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int,
	sentAmounts []int64, modificationsTimes []time.Time, modificationsVested []sdk.Int, modificationsWithdrawn []sdk.Int) {

	allAccVestingPools := h.helperCfevestingKeeper.GetAllAccountVestingPools(ctx)

	accVestingPools, accFound := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, address.String())
	require.EqualValues(h.t, true, accFound)

	require.EqualValues(h.t, amountOfAllAccVestingPools, len(allAccVestingPools))
	require.EqualValues(h.t, len(vestingTypes), len(accVestingPools.VestingPools))

	require.EqualValues(h.t, address.String(), accVestingPools.Address)

	for i, vesting := range accVestingPools.VestingPools {
		found := false
		// if vesting.Id == int32(i+1) {
		h.VerifyVestingPool(ctx, vesting, vestingNames[i], vestingTypes[i].Name, startsTimes[i], startsTimes[i].Add(durations[i]),
			vestedAmounts[i], withdrawnAmounts[i], sdk.NewInt(sentAmounts[i]))
		found = true

		// }
		require.True(h.t, found, "not found vesting id: "+strconv.Itoa(i+1))

	}

}

func (h *C4eVestingUtils) VerifyVestingPool(ctx sdk.Context, vp *cfevestingtypes.VestingPool, expectedName string,
	expectedVestingType string, expectedLockStart time.Time, expectedLockEnd time.Time, expectedInitiallyLocked sdk.Int,
	expectedWithdrawn sdk.Int, expectedSent sdk.Int) {
	require.EqualValues(h.t, expectedName, vp.Name)
	require.EqualValues(h.t, expectedVestingType, vp.VestingType)
	require.EqualValues(h.t, true, expectedLockStart.Equal(vp.LockStart))
	require.EqualValues(h.t, true, expectedLockEnd.Equal(vp.LockEnd))
	require.EqualValues(h.t, expectedInitiallyLocked, vp.InitiallyLocked)
	require.EqualValues(h.t, expectedWithdrawn, vp.Withdrawn)
	require.EqualValues(h.t, expectedSent, vp.Sent)
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
	expectedWithdrawn sdk.Int) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accountBalanceBefore)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleBalanceBefore)

	msg := cfevestingtypes.MsgWithdrawAllAvailable{Creator: address.String()}
	resp, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)
	require.True(h.t, expectedWithdrawn.Equal(resp.Withdrawn.Amount))
	require.EqualValues(h.t, h.helperCfevestingKeeper.GetParams(ctx).Denom, resp.Withdrawn.Denom)
	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, address, accountBalanceBefore.Add(expectedWithdrawn))
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleBalanceBefore.Sub(expectedWithdrawn))
}

func (h *C4eVestingUtils) MessageWithdrawAllAvailableError(ctx sdk.Context, address string, errorMessage string) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgWithdrawAllAvailable{Creator: address}
	_, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)

	require.EqualError(h.t, err, errorMessage)
}

func (h *C4eVestingUtils) CompareStoredAcountVestingPools(ctx sdk.Context, address sdk.AccAddress, accVestingPools cfevestingtypes.AccountVestingPools) {
	storedAccVestingPools, accFound := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, address.String())
	require.EqualValues(h.t, true, accFound)

	AssertAccountVestingPools(h.t, accVestingPools, storedAccVestingPools)
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
	require.EqualValues(h.t, expected.Params, got.GetParams())
	require.EqualValues(h.t, expected.Params, got.GetParams())
	require.EqualValues(h.t, expected.Params, got.GetParams())
	require.EqualValues(h.t, expected.Params, got.GetParams())
	require.EqualValues(h.t, expected.Params, got.GetParams())

	require.ElementsMatch(h.t, expected.VestingTypes, (*got).VestingTypes)
	require.EqualValues(h.t, len(expected.AccountVestingPools), len((*got).AccountVestingPools))
	AssertAccountVestingPoolsArrays(h.t, expected.AccountVestingPools, (*got).AccountVestingPools)
	require.EqualValues(h.t, expected.VestingAccountCount, (*got).VestingAccountCount)
	require.ElementsMatch(h.t, expected.VestingAccountList, (*got).VestingAccountList)

	nullify.Fill(&expected)
	nullify.Fill(got)
}

func (m *C4eVestingUtils) ExportGenesisAndValidate(ctx sdk.Context) {
	exportedGenesis := cfevesting.ExportGenesis(ctx, *m.helperCfevestingKeeper)
	err := exportedGenesis.Validate()
	require.NoError(m.t, err)
	err = cfevesting.ValidateAccountsOnGenesis(
		ctx,
		*m.helperCfevestingKeeper,
		*exportedGenesis,
		*m.helperAccountKeeper,
		*m.helperBankKeeper,
		*m.helperStakingKeeper)
	require.NoError(m.t, err)
}

func (m *C4eVestingUtils) ValidateInvariants(ctx sdk.Context) {
	invariants := []sdk.Invariant{
		cfevestingmodulekeeper.ModuleAccountInvariant(*m.helperCfevestingKeeper),
		cfevestingmodulekeeper.VestingPoolConsistentDataInvariant(*m.helperCfevestingKeeper),
		cfevestingmodulekeeper.NonNegativeVestingPoolAmountsInvariant(*m.helperCfevestingKeeper),
	}
	commontestutils.ValidateManyInvariants(m.t, ctx, invariants)
}

func (h *C4eVestingUtils) QueryVestingsSummary(wctx context.Context, expectedResponse cfevestingtypes.QueryVestingsSummaryResponse) {
	resp, err := h.helperCfevestingKeeper.VestingsSummary(wctx, &cfevestingtypes.QueryVestingsSummaryRequest{})
	require.NoError(h.t, err)
	require.Equal(h.t, expectedResponse, *resp)
}

func (h *C4eVestingUtils) SetVestingTypes(ctx sdk.Context, vestingTypes cfevestingtypes.VestingTypes) {
	h.helperCfevestingKeeper.SetVestingTypes(ctx, vestingTypes)
}

func (h *C4eVestingUtils) MessageSendToVestingAccount(ctx sdk.Context, fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolName string, amount sdk.Int, restartVesting bool) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	vestingAccountCount := h.helperCfevestingKeeper.GetVestingAccountCount(ctx)
	accAmountBefore := h.bankUtils.GetAccountDefultDenomBalance(ctx, vestingAccAddress)
	moduleAmountBefore := h.bankUtils.GetModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName)

	vestingPools, found := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, fromAddress.String())
	require.Equal(h.t, true, found)
	foundVPool, _ := GetVestingPoolByName(vestingPools.VestingPools, vestingPoolName)
	require.NotNilf(h.t, foundVPool, "vesting pool no found. Name: %s", vestingPoolName)
	sentBefore := foundVPool.Sent
	msg := cfevestingtypes.MsgSendToVestingAccount{FromAddress: fromAddress.String(), ToAddress: vestingAccAddress.String(),
		VestingPoolName: vestingPoolName, Amount: amount, RestartVesting: restartVesting}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, vestingAccAddress, accAmountBefore.Add(amount))
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleAmountBefore.Sub(amount))

	require.Equal(h.t, uint64(vestingAccountCount+1), h.helperCfevestingKeeper.GetVestingAccountCount(ctx))
	vaccFromList, found := h.helperCfevestingKeeper.GetVestingAccount(ctx, uint64(vestingAccountCount))
	require.Equal(h.t, true, found)
	require.Equal(h.t, vestingAccAddress.String(), vaccFromList.Address)

	vestingPools, found = h.helperCfevestingKeeper.GetAccountVestingPools(ctx, fromAddress.String())
	require.Equal(h.t, true, found)

	foundVPool, _ = GetVestingPoolByName(vestingPools.VestingPools, vestingPoolName)
	require.NotNilf(h.t, foundVPool, "vesting pool no found. Name: %d", vestingPoolName)
	require.Equal(h.t, sentBefore.Add(amount), foundVPool.Sent)

	vestingType, err := h.helperCfevestingKeeper.GetVestingType(ctx, foundVPool.VestingType)
	require.NoError(h.t, err, "GetVestingType error")

	denom := h.helperCfevestingKeeper.Denom(ctx)
	decAmount := amount.ToDec()
	lockedVesting := decAmount.Sub(decAmount.Mul(vestingType.InitialBonus)).TruncateInt()
	if restartVesting {
		h.authUtils.VerifyVestingAccount(ctx, vestingAccAddress, denom, lockedVesting, ctx.BlockTime().Add(vestingType.LockupPeriod), ctx.BlockTime().Add(vestingType.LockupPeriod).Add(vestingType.VestingPeriod))
	} else {
		h.authUtils.VerifyVestingAccount(ctx, vestingAccAddress, denom, lockedVesting, foundVPool.LockStart, foundVPool.LockEnd)
	}
}

func (h *C4eVestingUtils) MessageSendToVestingAccountError(ctx sdk.Context, fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolName string, amount sdk.Int, restartVesting bool, errorMessage string) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	vestingAccountCount := h.helperCfevestingKeeper.GetVestingAccountCount(ctx)

	msg := cfevestingtypes.MsgSendToVestingAccount{FromAddress: fromAddress.String(), ToAddress: vestingAccAddress.String(),
		VestingPoolName: vestingPoolName, Amount: amount, RestartVesting: restartVesting}
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

func (h *ContextC4eVestingUtils) SetupAccountVestingPools(address string, numberOfVestingPools int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) cfevestingtypes.AccountVestingPools {
	return h.C4eVestingUtils.SetupAccountVestingPools(h.testContext.GetContext(), address, numberOfVestingPools, vestingAmount, withdrawnAmount)
}

func (h *ContextC4eVestingUtils) SetupAccountVestingsPoolsWithModification(modifyVesting func(*cfevestingtypes.VestingPool), address string, numberOfVestingPools int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) cfevestingtypes.AccountVestingPools {
	return h.C4eVestingUtils.SetupAccountVestingPoolsWithModification(h.testContext.GetContext(), modifyVesting, address, numberOfVestingPools, vestingAmount, withdrawnAmount)
}

func (h *ContextC4eVestingUtils) MessageCreateVestingPool(address sdk.AccAddress, accountVestingPoolsExistsBefore bool, accountVestingPoolsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest sdk.Int, accAmountBefore sdk.Int, moduleAmountBefore sdk.Int,
	accAmountAfter sdk.Int, moduleAmountAfter sdk.Int) {
	h.C4eVestingUtils.MessageCreateVestingPool(h.testContext.GetContext(), address, accountVestingPoolsExistsBefore, accountVestingPoolsExistsAfter, vestingPoolName, lockupDuration,
		vestingType, amountToVest, accAmountBefore, moduleAmountBefore, accAmountAfter, moduleAmountAfter)
}

func (h *ContextC4eVestingUtils) VerifyAccountVestingPools(address sdk.AccAddress,
	vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int, startAndModificationTime ...time.Time) {
	h.C4eVestingUtils.VerifyAccountVestingPools(h.testContext.GetContext(), address, vestingNames, durations, vestingTypes, vestedAmounts, withdrawnAmounts, startAndModificationTime...)
}

func (h *ContextC4eVestingUtils) VerifyAccountVestingPoolsWithModification(address sdk.AccAddress,
	amountOfAllAccVestingPools int, vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, startsTimes []time.Time, vestedAmounts []sdk.Int, withdrawnAmounts []sdk.Int,
	sentAmounts []int64, modificationsTimes []time.Time, modificationsVested []sdk.Int, modificationsWithdrawn []sdk.Int) {
	h.C4eVestingUtils.VerifyAccountVestingPoolsWithModification(h.testContext.GetContext(), address, amountOfAllAccVestingPools, vestingNames, durations, vestingTypes, startsTimes, vestedAmounts,
		withdrawnAmounts, sentAmounts, modificationsTimes, modificationsVested, modificationsWithdrawn)
}

func (h *ContextC4eVestingUtils) SetupVestingTypes(numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) cfevestingtypes.VestingTypes {
	return h.C4eVestingUtils.SetupVestingTypes(h.testContext.GetContext(), numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *ContextC4eVestingUtils) SetupVestingTypesWithModification(modifyVestingType func(*cfevestingtypes.VestingType), numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) cfevestingtypes.VestingTypes {
	return h.C4eVestingUtils.SetupVestingTypesWithModification(h.testContext.GetContext(), modifyVestingType, numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *ContextC4eVestingUtils) MessageWithdrawAllAvailable(address sdk.AccAddress, accountBalanceBefore sdk.Int, moduleBalanceBefore sdk.Int,
	expectedWithdrawn sdk.Int) {
	h.C4eVestingUtils.MessageWithdrawAllAvailable(h.testContext.GetContext(), address, accountBalanceBefore, moduleBalanceBefore, expectedWithdrawn)
}

func (h *ContextC4eVestingUtils) CompareStoredAcountVestingPools(address sdk.AccAddress, accVestingPools cfevestingtypes.AccountVestingPools) {
	h.C4eVestingUtils.CompareStoredAcountVestingPools(h.testContext.GetContext(), address, accVestingPools)
}

func (h *ContextC4eVestingUtils) InitGenesis(genState cfevestingtypes.GenesisState) {
	h.C4eVestingUtils.InitGenesis(h.testContext.GetContext(), genState)
}

func (h *ContextC4eVestingUtils) QueryVestings(expectedResponse cfevestingtypes.QueryVestingsSummaryResponse) {
	h.C4eVestingUtils.QueryVestingsSummary(h.testContext.GetWrappedContext(), expectedResponse)
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

func (h *ContextC4eVestingUtils) MessageSendToVestingAccount(fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolName string, amount sdk.Int, restartVesting bool) {
	h.C4eVestingUtils.MessageSendToVestingAccount(h.testContext.GetContext(), fromAddress, vestingAccAddress, vestingPoolName, amount, restartVesting)
}

func (h *ContextC4eVestingUtils) MessageSendToVestingAccountError(fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolName string, amount sdk.Int, restartVesting bool, errorMessage string) {
	h.C4eVestingUtils.MessageSendToVestingAccountError(h.testContext.GetContext(), fromAddress, vestingAccAddress, vestingPoolName, amount, restartVesting, errorMessage)
}

func (h *ContextC4eVestingUtils) MessageWithdrawAllAvailableError(address string, errorMessage string) {
	h.C4eVestingUtils.MessageWithdrawAllAvailableError(h.testContext.GetContext(), address, errorMessage)
}

func (h *ContextC4eVestingUtils) ExportGenesis(expected cfevestingtypes.GenesisState) {
	h.C4eVestingUtils.ExportGenesis(h.testContext.GetContext(), expected)
}

func (m *ContextC4eVestingUtils) ValidateGenesisAndInvariants() {
	m.C4eVestingUtils.ExportGenesisAndValidate(m.testContext.GetContext())
	m.C4eVestingUtils.ValidateInvariants(m.testContext.GetContext())
}

func (h *ContextC4eVestingUtils) InitGenesisError(genState cfevestingtypes.GenesisState, errorMessage string) {
	h.C4eVestingUtils.InitGenesisError(h.testContext.GetContext(), genState, errorMessage)
}

func (h *ContextC4eVestingUtils) CheckModuleAccountInvariant(failed bool, message string) {
	h.C4eVestingUtils.CheckModuleAccountInvariant(h.testContext.GetContext(), failed, message)
}

func (h *ContextC4eVestingUtils) SetupVestingTypesForAccountsVestingPools() {
	h.C4eVestingUtils.SetupVestingTypesForAccountsVestingPools(h.testContext.GetContext())
}

func (h *ContextC4eVestingUtils) MessageCreateVestingAccount(
	fromAddress sdk.AccAddress,
	toAddress sdk.AccAddress,
	coins sdk.Coins,
	startTime time.Time,
	endTime time.Time,
	amountBefore sdk.Int,
) {
	h.C4eVestingUtils.MessageCreateVestingAccount(
		h.testContext.GetContext(),
		fromAddress,
		toAddress,
		coins,
		startTime,
		endTime,
		amountBefore,
	)
}

func (h *C4eVestingUtils) MessageCreateVestingAccount(
	ctx sdk.Context,
	fromAddress sdk.AccAddress,
	toAddress sdk.AccAddress,
	coins sdk.Coins,
	startTime time.Time,
	endTime time.Time,
	amountBefore sdk.Int,
) {
	_, accFound := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, toAddress.String())
	require.EqualValues(h.t, false, accFound)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, fromAddress, amountBefore)
	vestingAccountCountBefore := h.helperCfevestingKeeper.GetVestingAccountCount(ctx)
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgCreateVestingAccount{
		FromAddress: fromAddress.String(),
		ToAddress:   toAddress.String(),
		Amount:      coins,
		StartTime:   startTime.Unix(),
		EndTime:     endTime.Unix(),
	}
	_, err := msgServer.CreateVestingAccount(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)

	vestingAccountCountAfter := h.helperCfevestingKeeper.GetVestingAccountCount(ctx)
	require.EqualValues(h.t, vestingAccountCountBefore+1, vestingAccountCountAfter)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, fromAddress, amountBefore.Sub(coins.AmountOf(commontestutils.DefaultTestDenom)))
	h.authUtils.VerifyVestingAccount(ctx, toAddress, commontestutils.DefaultTestDenom, coins.AmountOf(commontestutils.DefaultTestDenom), startTime, endTime)
	accFromList, found := h.helperCfevestingKeeper.GetVestingAccount(ctx, vestingAccountCountBefore)
	require.Equal(h.t, true, found)
	require.Equal(h.t, toAddress.String(), accFromList.Address)
}

func (h *ContextC4eVestingUtils) MessageCreateVestingAccountError(
	fromAddress sdk.AccAddress,
	toAddress sdk.AccAddress,
	amount sdk.Coins,
	startTime time.Time,
	endTime time.Time,
	amountBefore sdk.Int,
	errorMessage string,
) {
	h.C4eVestingUtils.MessageCreateVestingAccountError(
		h.testContext.GetContext(),
		fromAddress,
		toAddress,
		amount,
		startTime,
		endTime,
		amountBefore,
		errorMessage,
	)
}

func (h *C4eVestingUtils) MessageCreateVestingAccountError(
	ctx sdk.Context,
	fromAddress sdk.AccAddress,
	toAddress sdk.AccAddress,
	amount sdk.Coins,
	startTime time.Time,
	endTime time.Time,
	amountBefore sdk.Int,
	errorMessage string,
) {
	vestingAccountCountBefore := h.helperCfevestingKeeper.GetVestingAccountCount(ctx)
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgCreateVestingAccount{
		FromAddress: fromAddress.String(),
		ToAddress:   toAddress.String(),
		Amount:      amount,
		StartTime:   startTime.Unix(),
		EndTime:     endTime.Unix(),
	}
	vestingAccountCountAfter := h.helperCfevestingKeeper.GetVestingAccountCount(ctx)
	_, err := msgServer.CreateVestingAccount(msgServerCtx, &msg)
	require.EqualValues(h.t, errorMessage, err.Error())

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, fromAddress, amountBefore)
	require.EqualValues(h.t, vestingAccountCountBefore, vestingAccountCountAfter)
	_, found := h.helperCfevestingKeeper.GetVestingAccount(ctx, vestingAccountCountBefore)
	require.Equal(h.t, false, found)
}
