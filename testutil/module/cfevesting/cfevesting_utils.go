package cfevesting

import (
	"context"
	"strconv"
	"time"

	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/testutil/nullify"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chain4energy/c4e-chain/x/cfevesting"
	cfevestingmodulekeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"

	"github.com/stretchr/testify/require"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

type C4eVestingKeeperUtils struct {
	t                      require.TestingT
	helperCfevestingKeeper *cfevestingmodulekeeper.Keeper
}

func NewC4eVestingKeeperUtils(t require.TestingT, helperCfevestingKeeper *cfevestingmodulekeeper.Keeper) C4eVestingKeeperUtils {
	return C4eVestingKeeperUtils{t: t, helperCfevestingKeeper: helperCfevestingKeeper}
}

func (h *C4eVestingKeeperUtils) GetC4eVestingKeeper() *cfevestingmodulekeeper.Keeper {
	return h.helperCfevestingKeeper
}

func (h *C4eVestingKeeperUtils) SetupAccountVestingPools(ctx sdk.Context, address string, numberOfVestingPools int, vestingAmount math.Int, withdrawnAmount math.Int) cfevestingtypes.AccountVestingPools {
	return h.SetupAccountVestingPoolsWithModification(ctx, func(*cfevestingtypes.VestingPool) { /*do not modify*/ }, address, numberOfVestingPools, vestingAmount, withdrawnAmount)
}

func (h *C4eVestingKeeperUtils) SetupAccountVestingPoolsWithModification(ctx sdk.Context, modifyVesting func(*cfevestingtypes.VestingPool), address string, numberOfVestingPools int, vestingAmount math.Int, withdrawnAmount math.Int) cfevestingtypes.AccountVestingPools {
	accountVestingPools := GenerateOneAccountVestingPoolsWithAddressWith10BasedVestingPools(numberOfVestingPools, 1, 1)
	accountVestingPools.Owner = address

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
	testcosmos.CheckInvariant(h.t, ctx, invariant, failed, message)
}

func (h *C4eVestingKeeperUtils) CheckVestingPoolConsistentDataInvariant(ctx sdk.Context, failed bool, message string) {
	invariant := cfevestingmodulekeeper.VestingPoolConsistentDataInvariant(*h.helperCfevestingKeeper)
	testcosmos.CheckInvariant(h.t, ctx, invariant, failed, message)
}

func (h *C4eVestingKeeperUtils) CheckModuleAccountInvariant(ctx sdk.Context, failed bool, message string) {
	invariant := cfevestingmodulekeeper.ModuleAccountInvariant(*h.helperCfevestingKeeper)
	testcosmos.CheckInvariant(h.t, ctx, invariant, failed, message)
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
	bankUtils           *testcosmos.BankUtils
	authUtils           *testcosmos.AuthUtils
}

func NewC4eVestingUtils(t require.TestingT, helperCfevestingKeeper *cfevestingmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	helperBankKeeper *bankkeeper.Keeper,
	helperStakingKeeper *stakingkeeper.Keeper, bankUtils *testcosmos.BankUtils,
	authUtils *testcosmos.AuthUtils) C4eVestingUtils {
	return C4eVestingUtils{C4eVestingKeeperUtils: NewC4eVestingKeeperUtils(t, helperCfevestingKeeper), helperAccountKeeper: helperAccountKeeper,
		helperBankKeeper: helperBankKeeper, helperStakingKeeper: helperStakingKeeper, bankUtils: bankUtils, authUtils: authUtils}
}

func (h *C4eVestingUtils) MessageCreateVestingPool(ctx sdk.Context, address sdk.AccAddress, accountVestingPoolsExistsBefore bool, accountVestingPoolsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest math.Int, accAmountBefore math.Int, moduleAmountBefore math.Int,
	accAmountAfter math.Int, moduleAmountAfter math.Int) {
	h.MessageCreateVestingPoolWithGenesisParam(ctx, address, accountVestingPoolsExistsBefore, accountVestingPoolsExistsAfter, vestingPoolName, lockupDuration, vestingType,
		amountToVest, accAmountBefore, moduleAmountBefore, accAmountAfter, moduleAmountAfter, false)
}

func (h *C4eVestingUtils) MessageCreateGenesisVestingPool(ctx sdk.Context, address sdk.AccAddress, accountVestingPoolsExistsBefore bool, accountVestingPoolsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest math.Int, accAmountBefore math.Int, moduleAmountBefore math.Int,
	accAmountAfter math.Int, moduleAmountAfter math.Int) {
	h.MessageCreateVestingPoolWithGenesisParam(ctx, address, accountVestingPoolsExistsBefore, accountVestingPoolsExistsAfter, vestingPoolName, lockupDuration, vestingType,
		amountToVest, accAmountBefore, moduleAmountBefore, accAmountAfter, moduleAmountAfter, true)

}

func (h *C4eVestingUtils) SendToRepeatedContinuousVestingAccount(ctx sdk.Context, toAddress sdk.AccAddress,
	amount math.Int, free sdk.Dec, startTime int64, endTime int64) {
	coins := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amount))
	moduleBalance := h.bankUtils.GetModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName)
	accBalance := h.bankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)

	previousOriginalVesting := sdk.NewCoins()
	var previousPeriods []cfevestingtypes.ContinuousVestingPeriod
	if accountBefore != nil {
		if claimAccount, ok := accountBefore.(*cfevestingtypes.PeriodicContinuousVestingAccount); ok {
			previousOriginalVesting = previousOriginalVesting.Add(claimAccount.OriginalVesting...)
			previousPeriods = claimAccount.VestingPeriods
		}
	}
	_, err := h.helperCfevestingKeeper.SendToPeriodicContinuousVestingAccountFromModule(ctx, cfevestingtypes.ModuleName,
		toAddress.String(),
		coins,
		free,
		startTime,
		endTime,
	)
	require.NoError(h.t, err)

	h.bankUtils.VerifyAccountDefaultDenomBalance(ctx, toAddress, accBalance.Add(amount))
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleBalance.Sub(amount))

	claimAccount, ok := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	require.True(h.t, ok)
	newPeriods := append(previousPeriods, cfevestingtypes.ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: coins})
	h.VerifyRepeatedContinuousVestingAccount(ctx, toAddress, previousOriginalVesting.Add(coins...), startTime, endTime, newPeriods)
	require.NoError(h.t, claimAccount.Validate())
}

func (h *C4eVestingUtils) SendToRepeatedContinuousVestingAccountError(ctx sdk.Context, toAddress sdk.AccAddress,
	amount math.Int, free sdk.Dec, startTime int64, endTime int64, createAccount bool, errorMessage string) {
	coins := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amount))
	moduleBalance := h.bankUtils.GetModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName)
	accBalance := h.bankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	wasAccount := false
	if accountBefore != nil {
		_, wasAccount = accountBefore.(*cfevestingtypes.PeriodicContinuousVestingAccount)
	}
	_, err := h.helperCfevestingKeeper.SendToPeriodicContinuousVestingAccountFromModule(ctx, cfevestingtypes.ModuleName,
		toAddress.String(),
		coins,
		free,
		startTime,
		endTime,
	)
	require.EqualError(h.t, err, errorMessage)

	h.bankUtils.VerifyAccountDefaultDenomBalance(ctx, toAddress, accBalance)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleBalance)

	accountAfter := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	_, isAccount := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	_, ok := accountBefore.(*cfevestingtypes.PeriodicContinuousVestingAccount)
	if ok {
		require.EqualValues(h.t, true, isAccount)
		h.VerifyRepeatedContinuousVestingAccount(ctx, toAddress, sdk.NewCoins(), startTime, endTime, []cfevestingtypes.ContinuousVestingPeriod{})
	} else {
		require.EqualValues(h.t, wasAccount, isAccount)
		require.EqualValues(h.t, accountBefore, accountAfter)
	}

}

func (h *C4eVestingUtils) VerifyRepeatedContinuousVestingAccount(ctx sdk.Context, address sdk.AccAddress,
	expectedOriginalVesting sdk.Coins, expectedStartTime int64, expectedEndTime int64, expectedPeriods []cfevestingtypes.ContinuousVestingPeriod) {

	claimAccount, ok := h.helperAccountKeeper.GetAccount(ctx, address).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	require.True(h.t, ok)

	require.EqualValues(h.t, len(expectedPeriods), len(claimAccount.VestingPeriods))
	require.EqualValues(h.t, expectedStartTime, claimAccount.StartTime)
	require.EqualValues(h.t, expectedEndTime, claimAccount.EndTime)
	require.True(h.t, expectedOriginalVesting.IsEqual(claimAccount.OriginalVesting))
	for i := 0; i < len(expectedPeriods); i++ {
		require.EqualValues(h.t, expectedPeriods[i].StartTime, claimAccount.VestingPeriods[i].StartTime)
		require.EqualValues(h.t, expectedPeriods[i].EndTime, claimAccount.VestingPeriods[i].EndTime)
		require.EqualValues(h.t, expectedPeriods[i].Amount, claimAccount.VestingPeriods[i].Amount)
	}
	require.NoError(h.t, claimAccount.Validate())
}

func (h *C4eVestingUtils) MessageCreateVestingPoolWithGenesisParam(ctx sdk.Context, address sdk.AccAddress, accountVestingPoolsExistsBefore bool, accountVestingPoolsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest math.Int, accAmountBefore math.Int, moduleAmountBefore math.Int,
	accAmountAfter math.Int, moduleAmountAfter math.Int, isGenesisPool bool) {
	_, accFound := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, address.String())
	require.EqualValues(h.t, accountVestingPoolsExistsBefore, accFound)

	h.bankUtils.VerifyAccountDefaultDenomBalance(ctx, address, accAmountBefore)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleAmountBefore)

	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgCreateVestingPool{Owner: address.String(), Name: vestingPoolName,
		Amount: amountToVest, Duration: lockupDuration, VestingType: vestingType.Name}
	_, err := msgServer.CreateVestingPool(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)

	accVestingPools, accFound := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, address.String())
	require.EqualValues(h.t, accountVestingPoolsExistsAfter, accFound)

	if accFound {
		var vestingPool *cfevestingtypes.VestingPool = nil
		for _, vest := range accVestingPools.VestingPools {
			if vest.Name == vestingPoolName {
				if isGenesisPool {
					vest.GenesisPool = true
				}
				vestingPool = vest

			}
		}
		if isGenesisPool {
			h.helperCfevestingKeeper.SetAccountVestingPools(ctx, accVestingPools)
			accVestingPools, accFound = h.helperCfevestingKeeper.GetAccountVestingPools(ctx, address.String())
			require.EqualValues(h.t, accountVestingPoolsExistsAfter, accFound)
			for _, vest := range accVestingPools.VestingPools {
				if vest.Name == vestingPoolName {
					vestingPool = vest
				}
			}
			require.True(h.t, vestingPool.GenesisPool)

		}
		require.NotNil(h.t, vestingPool)
		h.VerifyVestingPool(ctx, vestingPool, vestingPoolName, vestingType.Name, ctx.BlockTime(), ctx.BlockTime().Add(lockupDuration),
			amountToVest, sdk.ZeroInt(), sdk.ZeroInt())
	}
	h.bankUtils.VerifyAccountDefaultDenomBalance(ctx, address, accAmountAfter)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleAmountAfter)

	if isGenesisPool && accountVestingPoolsExistsAfter {
		accVestingPools, accFound = h.helperCfevestingKeeper.GetAccountVestingPools(ctx, address.String())
		require.EqualValues(h.t, accountVestingPoolsExistsAfter, accFound)

	}
}

func (h *C4eVestingUtils) MessageCreateVestingPoolError(ctx sdk.Context, address sdk.AccAddress,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest math.Int,
	errorMessage string) {

	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgCreateVestingPool{Owner: address.String(), Name: vestingPoolName,
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
	vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, vestedAmounts []math.Int, withdrawnAmounts []math.Int, startAndModificationTime ...time.Time) {

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
	amountOfAllAccVestingPools int, vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, startsTimes []time.Time, vestedAmounts []math.Int, withdrawnAmounts []math.Int,
	sentAmounts []int64, modificationsTimes []time.Time, modificationsVested []math.Int, modificationsWithdrawn []math.Int) {

	allAccVestingPools := h.helperCfevestingKeeper.GetAllAccountVestingPools(ctx)

	accVestingPools, accFound := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, address.String())
	require.EqualValues(h.t, true, accFound)

	require.EqualValues(h.t, amountOfAllAccVestingPools, len(allAccVestingPools))
	require.EqualValues(h.t, len(vestingTypes), len(accVestingPools.VestingPools))

	require.EqualValues(h.t, address.String(), accVestingPools.Owner)

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
	expectedVestingType string, expectedLockStart time.Time, expectedLockEnd time.Time, expectedInitiallyLocked math.Int,
	expectedWithdrawn math.Int, expectedSent math.Int) {
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

func (h *C4eVestingUtils) MessageWithdrawAllAvailable(ctx sdk.Context, address sdk.AccAddress, accountBalanceBefore math.Int, moduleBalanceBefore math.Int,
	expectedWithdrawn math.Int) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	h.bankUtils.VerifyAccountDefaultDenomBalance(ctx, address, accountBalanceBefore)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleBalanceBefore)

	msg := cfevestingtypes.MsgWithdrawAllAvailable{Owner: address.String()}
	resp, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)
	require.True(h.t, expectedWithdrawn.Equal(resp.Withdrawn.Amount))
	require.EqualValues(h.t, h.helperCfevestingKeeper.GetParams(ctx).Denom, resp.Withdrawn.Denom)
	h.bankUtils.VerifyAccountDefaultDenomBalance(ctx, address, accountBalanceBefore.Add(expectedWithdrawn))
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleBalanceBefore.Sub(expectedWithdrawn))
}

func (h *C4eVestingUtils) MessageWithdrawAllAvailableError(ctx sdk.Context, address string, errorMessage string) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgWithdrawAllAvailable{Owner: address}
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
	require.EqualValues(h.t, expected.Params, got.GetParams())
	require.ElementsMatch(h.t, expected.VestingTypes, (*got).VestingTypes)
	require.EqualValues(h.t, len(expected.AccountVestingPools), len((*got).AccountVestingPools))
	AssertAccountVestingPoolsArrays(h.t, expected.AccountVestingPools, (*got).AccountVestingPools)
	require.EqualValues(h.t, expected.VestingAccountTraceCount, (*got).VestingAccountTraceCount)
	require.ElementsMatch(h.t, expected.VestingAccountTraces, (*got).VestingAccountTraces)

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
	testcosmos.ValidateManyInvariants(m.t, ctx, invariants)
}

func (h *C4eVestingUtils) QueryVestingsSummary(wctx context.Context, expectedResponse cfevestingtypes.QueryVestingsSummaryResponse) {
	resp, err := h.helperCfevestingKeeper.VestingsSummary(wctx, &cfevestingtypes.QueryVestingsSummaryRequest{})
	require.NoError(h.t, err)
	require.Equal(h.t, expectedResponse, *resp)
}

func (h *C4eVestingUtils) QueryGenesisVestingsSummary(wctx context.Context, expectedResponse cfevestingtypes.QueryGenesisVestingsSummaryResponse) {
	resp, err := h.helperCfevestingKeeper.GenesisVestingsSummary(wctx, &cfevestingtypes.QueryGenesisVestingsSummaryRequest{})
	require.NoError(h.t, err)
	require.Equal(h.t, expectedResponse, *resp)
}

func (h *C4eVestingUtils) SetVestingTypes(ctx sdk.Context, vestingTypes cfevestingtypes.VestingTypes) {
	h.helperCfevestingKeeper.SetVestingTypes(ctx, vestingTypes)
}

func (h *C4eVestingUtils) MessageSendToVestingAccount(ctx sdk.Context, fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolName string, amount math.Int, restartVesting bool, expectedLocked math.Int) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	vestingAccountCount := h.helperCfevestingKeeper.GetVestingAccountTraceCount(ctx)
	accAmountBefore := h.bankUtils.GetAccountDefultDenomBalance(ctx, vestingAccAddress)
	moduleAmountBefore := h.bankUtils.GetModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName)
	accountBefore := h.helperAccountKeeper.GetAccount(ctx, vestingAccAddress)

	newVestingAccountTraceId := vestingAccountCount
	if accountBefore != nil {
		newVestingAccountTraceId = vestingAccountCount - 1
	} else {
		vestingAccountCount++
	}
	vestingPools, found := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, fromAddress.String())
	require.Equal(h.t, true, found)
	foundVPool, _ := GetVestingPoolByName(vestingPools.VestingPools, vestingPoolName)
	require.NotNilf(h.t, foundVPool, "vesting pool no found. Name: %s", vestingPoolName)
	sentBefore := foundVPool.Sent
	msg := cfevestingtypes.MsgSendToVestingAccount{Owner: fromAddress.String(), ToAddress: vestingAccAddress.String(),
		VestingPoolName: vestingPoolName, Amount: amount, RestartVesting: restartVesting}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(h.t, nil, err)

	h.bankUtils.VerifyAccountDefaultDenomBalance(ctx, vestingAccAddress, accAmountBefore.Add(amount))
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfevestingtypes.ModuleName, moduleAmountBefore.Sub(amount))
	require.Equal(h.t, vestingAccountCount, h.helperCfevestingKeeper.GetVestingAccountTraceCount(ctx))

	vaccFromList, found := h.helperCfevestingKeeper.GetVestingAccountTraceById(ctx, newVestingAccountTraceId)
	require.Equal(h.t, true, found)
	expectedVestingAccountTrace := cfevestingtypes.VestingAccountTrace{
		Id:                 newVestingAccountTraceId,
		Address:            vestingAccAddress.String(),
		Genesis:            false,
		FromGenesisPool:    foundVPool.GenesisPool,
		FromGenesisAccount: false,
		PeriodsToTrace:     []uint64{0},
	}
	require.EqualValues(h.t, expectedVestingAccountTrace, vaccFromList)

	vestingPools, found = h.helperCfevestingKeeper.GetAccountVestingPools(ctx, fromAddress.String())
	require.Equal(h.t, true, found)

	foundVPool, _ = GetVestingPoolByName(vestingPools.VestingPools, vestingPoolName)
	require.NotNilf(h.t, foundVPool, "vesting pool no found. Name: %d", vestingPoolName)
	require.Equal(h.t, sentBefore.Add(amount), foundVPool.Sent)

	vestingType, err := h.helperCfevestingKeeper.GetVestingType(ctx, foundVPool.VestingType)
	require.NoError(h.t, err, "GetVestingType error")

	denom := h.helperCfevestingKeeper.Denom(ctx)

	lockedAmount := sdk.NewCoins(sdk.NewCoin(denom, expectedLocked))
	if restartVesting {
		h.authUtils.VerifyVestingAccount(ctx, vestingAccAddress, lockedAmount, ctx.BlockTime().Add(vestingType.LockupPeriod), ctx.BlockTime().Add(vestingType.LockupPeriod).Add(vestingType.VestingPeriod))
	} else {
		h.authUtils.VerifyVestingAccount(ctx, vestingAccAddress, lockedAmount, foundVPool.LockStart, foundVPool.LockEnd)
	}
}

func (h *C4eVestingUtils) MessageSendToVestingAccountError(ctx sdk.Context, fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolName string, amount math.Int, restartVesting bool, errorMessage string) {
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	vestingAccountCount := h.helperCfevestingKeeper.GetVestingAccountTraceCount(ctx)

	msg := cfevestingtypes.MsgSendToVestingAccount{Owner: fromAddress.String(), ToAddress: vestingAccAddress.String(),
		VestingPoolName: vestingPoolName, Amount: amount, RestartVesting: restartVesting}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualError(h.t, err, errorMessage)
	require.Equal(h.t, uint64(vestingAccountCount), h.helperCfevestingKeeper.GetVestingAccountTraceCount(ctx))
}

func (h *C4eVestingUtils) UnlockUnbondedContinuousVestingAccountCoins(ctx sdk.Context, ownerAddress sdk.AccAddress, amountsToUnlock sdk.Coins, expectedAccountBalances sdk.Coins, expectedLockedBalancesBefore sdk.Coins) {
	h.bankUtils.VerifyAccountBalances(ctx, ownerAddress, expectedAccountBalances, true)
	locked := (*h.helperBankKeeper).LockedCoins(ctx, ownerAddress)
	require.Truef(h.t, expectedLockedBalancesBefore.IsEqual(locked), "expectedLockedBalancesBefore %s <> locked %s", expectedLockedBalancesBefore, locked)

	_, err := h.helperCfevestingKeeper.UnlockUnbondedContinuousVestingAccountCoins(ctx, ownerAddress, amountsToUnlock)
	require.NoError(h.t, err, err)

	h.bankUtils.VerifyAccountBalances(ctx, ownerAddress, expectedAccountBalances, true)
	locked = (*h.helperBankKeeper).LockedCoins(ctx, ownerAddress)
	require.Truef(h.t, expectedLockedBalancesBefore.Sub(amountsToUnlock...).IsEqual(locked), "expectedLockedBalances %s <> locked %s", expectedLockedBalancesBefore.Sub(amountsToUnlock...), locked)
}

func (h *C4eVestingUtils) UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(ctx sdk.Context, ownerAddress sdk.AccAddress, amountToUnlock math.Int, expectedAccountBalance math.Int, expectedLockedBalanceBefore math.Int) {
	h.UnlockUnbondedContinuousVestingAccountCoins(ctx, ownerAddress,
		sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amountToUnlock)),
		sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, expectedAccountBalance)),
		sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, expectedLockedBalanceBefore)))
}

func (h *C4eVestingUtils) UnlockUnbondedContinuousVestingAccountCoinsError(ctx sdk.Context, ownerAddress sdk.AccAddress, amountsToUnlock sdk.Coins, expectedAccountBalances sdk.Coins, expectedLockedBalancesBefore sdk.Coins, expectedError string) {
	h.bankUtils.VerifyAccountBalances(ctx, ownerAddress, expectedAccountBalances, true)
	locked := (*h.helperBankKeeper).LockedCoins(ctx, ownerAddress)
	require.Truef(h.t, expectedLockedBalancesBefore.IsEqual(locked), "expectedLockedBalancesBefore %s <> locked %s", expectedLockedBalancesBefore, locked)

	_, err := h.helperCfevestingKeeper.UnlockUnbondedContinuousVestingAccountCoins(ctx, ownerAddress, amountsToUnlock)
	require.Error(h.t, err)
	require.EqualError(h.t, err, expectedError)

}

func (h *C4eVestingUtils) UnlockUnbondedDefaultDenomContinuousVestingAccountCoinsError(ctx sdk.Context, ownerAddress sdk.AccAddress, amountToUnlock math.Int, expectedAccountBalance math.Int, expectedLockedBalanceBefore math.Int, expectedError string) {
	h.UnlockUnbondedContinuousVestingAccountCoinsError(ctx, ownerAddress,
		sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amountToUnlock)),
		sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, expectedAccountBalance)),
		sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, expectedLockedBalanceBefore)), expectedError)
}

type ContextC4eVestingUtils struct {
	C4eVestingUtils
	testContext testenv.TestContext
}

func NewContextC4eVestingUtils(t require.TestingT, testContext testenv.TestContext, helperCfevestingKeeper *cfevestingmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	helperBankKeeper *bankkeeper.Keeper,
	helperStakingKeeper *stakingkeeper.Keeper, bankUtils *testcosmos.BankUtils,
	authUtils *testcosmos.AuthUtils) *ContextC4eVestingUtils {
	c4eVestingUtils := NewC4eVestingUtils(t, helperCfevestingKeeper, helperAccountKeeper, helperBankKeeper, helperStakingKeeper, bankUtils, authUtils)
	return &ContextC4eVestingUtils{C4eVestingUtils: c4eVestingUtils, testContext: testContext}
}

func (h *ContextC4eVestingUtils) SetupAccountVestingPools(address string, numberOfVestingPools int, vestingAmount math.Int, withdrawnAmount math.Int) cfevestingtypes.AccountVestingPools {
	return h.C4eVestingUtils.SetupAccountVestingPools(h.testContext.GetContext(), address, numberOfVestingPools, vestingAmount, withdrawnAmount)
}

func (h *ContextC4eVestingUtils) SetupAccountVestingsPoolsWithModification(modifyVesting func(*cfevestingtypes.VestingPool), address string, numberOfVestingPools int, vestingAmount math.Int, withdrawnAmount math.Int) cfevestingtypes.AccountVestingPools {
	return h.C4eVestingUtils.SetupAccountVestingPoolsWithModification(h.testContext.GetContext(), modifyVesting, address, numberOfVestingPools, vestingAmount, withdrawnAmount)
}

func (h *ContextC4eVestingUtils) MessageCreateVestingPool(address sdk.AccAddress, accountVestingPoolsExistsBefore bool, accountVestingPoolsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest math.Int, accAmountBefore math.Int, moduleAmountBefore math.Int,
	accAmountAfter math.Int, moduleAmountAfter math.Int) {
	h.C4eVestingUtils.MessageCreateVestingPool(h.testContext.GetContext(), address, accountVestingPoolsExistsBefore, accountVestingPoolsExistsAfter, vestingPoolName, lockupDuration,
		vestingType, amountToVest, accAmountBefore, moduleAmountBefore, accAmountAfter, moduleAmountAfter)
}

func (h *ContextC4eVestingUtils) AddTestVestingPool(address sdk.AccAddress, vestingPoolName string, vested math.Int, lockupPeriodInHours int64, vestingPeriodInHours int64) {
	accInitBalance := sdk.NewInt(10000)
	h.bankUtils.AddDefaultDenomCoinsToAccount(h.testContext.GetContext(), accInitBalance, address)

	vestingType := cfevestingtypes.VestingType{
		Name:          "test-vesting-type",
		LockupPeriod:  CreateDurationFromNumOfHours(lockupPeriodInHours),
		VestingPeriod: CreateDurationFromNumOfHours(vestingPeriodInHours),
		Free:          sdk.MustNewDecFromStr("0.05"),
	}
	h.helperCfevestingKeeper.SetVestingType(h.testContext.GetContext(), vestingType)

	h.C4eVestingUtils.MessageCreateVestingPool(h.testContext.GetContext(), address, false, true, vestingPoolName, 1000,
		vestingType, vested, accInitBalance, math.ZeroInt(), accInitBalance.Sub(vested), vested)
}

func (h *ContextC4eVestingUtils) MessageCreateGenesisVestingPool(address sdk.AccAddress, accountVestingPoolsExistsBefore bool, accountVestingPoolsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest math.Int, accAmountBefore math.Int, moduleAmountBefore math.Int,
	accAmountAfter math.Int, moduleAmountAfter math.Int) {
	h.C4eVestingUtils.MessageCreateGenesisVestingPool(h.testContext.GetContext(), address, accountVestingPoolsExistsBefore, accountVestingPoolsExistsAfter, vestingPoolName, lockupDuration,
		vestingType, amountToVest, accAmountBefore, moduleAmountBefore, accAmountAfter, moduleAmountAfter)
}

func (h *ContextC4eVestingUtils) VerifyAccountVestingPools(address sdk.AccAddress,
	vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, vestedAmounts []math.Int, withdrawnAmounts []math.Int, startAndModificationTime ...time.Time) {
	h.C4eVestingUtils.VerifyAccountVestingPools(h.testContext.GetContext(), address, vestingNames, durations, vestingTypes, vestedAmounts, withdrawnAmounts, startAndModificationTime...)
}

func (h *ContextC4eVestingUtils) VerifyAccountVestingPoolsWithModification(address sdk.AccAddress,
	amountOfAllAccVestingPools int, vestingNames []string, durations []time.Duration, vestingTypes []cfevestingtypes.VestingType, startsTimes []time.Time, vestedAmounts []math.Int, withdrawnAmounts []math.Int,
	sentAmounts []int64, modificationsTimes []time.Time, modificationsVested []math.Int, modificationsWithdrawn []math.Int) {
	h.C4eVestingUtils.VerifyAccountVestingPoolsWithModification(h.testContext.GetContext(), address, amountOfAllAccVestingPools, vestingNames, durations, vestingTypes, startsTimes, vestedAmounts,
		withdrawnAmounts, sentAmounts, modificationsTimes, modificationsVested, modificationsWithdrawn)
}

func (h *ContextC4eVestingUtils) SetupVestingTypes(numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) cfevestingtypes.VestingTypes {
	return h.C4eVestingUtils.SetupVestingTypes(h.testContext.GetContext(), numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *ContextC4eVestingUtils) SetupVestingTypesWithModification(modifyVestingType func(*cfevestingtypes.VestingType), numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) cfevestingtypes.VestingTypes {
	return h.C4eVestingUtils.SetupVestingTypesWithModification(h.testContext.GetContext(), modifyVestingType, numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func (h *ContextC4eVestingUtils) MessageWithdrawAllAvailable(address sdk.AccAddress, accountBalanceBefore math.Int, moduleBalanceBefore math.Int,
	expectedWithdrawn math.Int) {
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

func (h *ContextC4eVestingUtils) QueryGenesisVestings(expectedResponse cfevestingtypes.QueryGenesisVestingsSummaryResponse) {
	h.C4eVestingUtils.QueryGenesisVestingsSummary(h.testContext.GetWrappedContext(), expectedResponse)
}

func (h *ContextC4eVestingUtils) MessageCreateVestingPoolError(address sdk.AccAddress,
	vestingPoolName string, lockupDuration time.Duration, vestingType cfevestingtypes.VestingType, amountToVest math.Int,
	errorMessage string) {
	h.C4eVestingUtils.MessageCreateVestingPoolError(h.testContext.GetContext(), address, vestingPoolName, lockupDuration,
		vestingType, amountToVest, errorMessage)
}

func (h *ContextC4eVestingUtils) SetVestingTypes(vestingTypes cfevestingtypes.VestingTypes) {
	h.C4eVestingUtils.SetVestingTypes(h.testContext.GetContext(), vestingTypes)
}

func (h *ContextC4eVestingUtils) MessageSendToVestingAccount(fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolName string, amount math.Int, restartVesting bool, expectedLocked math.Int) {
	h.C4eVestingUtils.MessageSendToVestingAccount(h.testContext.GetContext(), fromAddress, vestingAccAddress, vestingPoolName, amount, restartVesting, expectedLocked)
}

func (h *ContextC4eVestingUtils) MessageSendToVestingAccountError(fromAddress sdk.AccAddress, vestingAccAddress sdk.AccAddress, vestingPoolName string, amount math.Int, restartVesting bool, errorMessage string) {
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
	amountBefore math.Int,
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
	amountBefore math.Int,
) {
	_, accFound := h.helperCfevestingKeeper.GetAccountVestingPools(ctx, toAddress.String())
	require.EqualValues(h.t, false, accFound)

	h.bankUtils.VerifyAccountDefaultDenomBalance(ctx, fromAddress, amountBefore)
	vestingAccountCountBefore := h.helperCfevestingKeeper.GetVestingAccountTraceCount(ctx)
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

	vestingAccountCountAfter := h.helperCfevestingKeeper.GetVestingAccountTraceCount(ctx)
	require.EqualValues(h.t, vestingAccountCountBefore, vestingAccountCountAfter)

	h.bankUtils.VerifyAccountDefaultDenomBalance(ctx, fromAddress, amountBefore.Sub(coins.AmountOf(testenv.DefaultTestDenom)))
	h.authUtils.VerifyVestingAccount(ctx, toAddress, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, coins.AmountOf(testenv.DefaultTestDenom))), startTime, endTime)
	_, found := h.helperCfevestingKeeper.GetVestingAccountTraceById(ctx, vestingAccountCountBefore)
	require.False(h.t, found)
}

func (h *ContextC4eVestingUtils) MessageCreateVestingAccountError(
	fromAddress sdk.AccAddress,
	toAddress sdk.AccAddress,
	amount sdk.Coins,
	startTime time.Time,
	endTime time.Time,
	amountBefore math.Int,
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
	amountBefore math.Int,
	errorMessage string,
) {
	vestingAccountCountBefore := h.helperCfevestingKeeper.GetVestingAccountTraceCount(ctx)
	msgServer, msgServerCtx := cfevestingmodulekeeper.NewMsgServerImpl(*h.helperCfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := cfevestingtypes.MsgCreateVestingAccount{
		FromAddress: fromAddress.String(),
		ToAddress:   toAddress.String(),
		Amount:      amount,
		StartTime:   startTime.Unix(),
		EndTime:     endTime.Unix(),
	}
	vestingAccountCountAfter := h.helperCfevestingKeeper.GetVestingAccountTraceCount(ctx)
	_, err := msgServer.CreateVestingAccount(msgServerCtx, &msg)
	require.EqualValues(h.t, errorMessage, err.Error())

	h.bankUtils.VerifyAccountDefaultDenomBalance(ctx, fromAddress, amountBefore)
	require.EqualValues(h.t, vestingAccountCountBefore, vestingAccountCountAfter)
	_, found := h.helperCfevestingKeeper.GetVestingAccountTraceById(ctx, vestingAccountCountBefore)
	require.Equal(h.t, false, found)
}

func (h *ContextC4eVestingUtils) UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(ownerAddress sdk.AccAddress, amountToUnlock math.Int, expectedAccountBalance math.Int, expectedLockedBalanceBefore math.Int) {
	h.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(h.testContext.GetContext(), ownerAddress, amountToUnlock, expectedAccountBalance, expectedLockedBalanceBefore)
}

func (h *ContextC4eVestingUtils) UnlockUnbondedDefaultDenomContinuousVestingAccountCoinsError(ownerAddress sdk.AccAddress, amountToUnlock math.Int, expectedAccountBalance math.Int, expectedLockedBalanceBefore math.Int, expectedError string) {
	h.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoinsError(h.testContext.GetContext(), ownerAddress, amountToUnlock, expectedAccountBalance, expectedLockedBalanceBefore, expectedError)
}

func (h *ContextC4eVestingUtils) UnlockUnbondedContinuousVestingAccountCoins(ownerAddress sdk.AccAddress, amountsToUnlock sdk.Coins, expectedAccountBalances sdk.Coins, expectedLockedBalancesBefore sdk.Coins) {
	h.C4eVestingUtils.UnlockUnbondedContinuousVestingAccountCoins(h.testContext.GetContext(), ownerAddress, amountsToUnlock, expectedAccountBalances, expectedLockedBalancesBefore)
}

func (h *ContextC4eVestingUtils) UnlockUnbondedContinuousVestingAccountCoinsError(ownerAddress sdk.AccAddress, amountsToUnlock sdk.Coins, expectedAccountBalances sdk.Coins, expectedLockedBalancesBefore sdk.Coins, expectedError string) {
	h.C4eVestingUtils.UnlockUnbondedContinuousVestingAccountCoinsError(h.testContext.GetContext(), ownerAddress, amountsToUnlock, expectedAccountBalances, expectedLockedBalancesBefore, expectedError)

}
