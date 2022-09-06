package keeper_test

import (
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"testing"

	"github.com/stretchr/testify/require"
)

const (
	vPool1 = "v-pool-1"
	vPool2 = "v-pool-2"
)

func verifyAccountBalance(t *testing.T, app *app.App, ctx sdk.Context, accAddr sdk.AccAddress, expectedAmount sdk.Int) {
	balance := app.BankKeeper.GetBalance(ctx, accAddr, commontestutils.Denom)
	require.EqualValues(t, expectedAmount, balance.Amount)
}

func verifyModuleAccount(t *testing.T, ctx sdk.Context, app *app.App, expected sdk.Int) {
	commontestutils.VerifyModuleAccountBalanceByName(types.ModuleName, ctx, app, t, expected)
}

func createValidator(t *testing.T, ctx sdk.Context, sk stakingkeeper.Keeper, addr sdk.ValAddress, pk cryptotypes.PubKey, coin sdk.Coin, commisions stakingtypes.CommissionRates) {
	msg, err := stakingtypes.NewMsgCreateValidator(addr, pk, coin, stakingtypes.Description{}, commisions, sdk.OneInt())
	msgSrvr := stakingkeeper.NewMsgServerImpl(sk)
	require.NoError(t, err)
	res, err := msgSrvr.CreateValidator(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, res)

}

func setupValidators(t *testing.T, ctx sdk.Context, app *app.App, validators []sdk.ValAddress, delegatePerValidator uint64) {
	PKs := commontestutils.CreateTestPubKeys(len(validators))
	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(0, 1), sdk.NewDecWithPrec(0, 1), sdk.NewDec(0))
	delCoin := sdk.NewCoin(commontestutils.Denom, sdk.NewIntFromUint64(delegatePerValidator))
	for i, valAddr := range validators {
		commontestutils.AddCoinsToAccount(delegatePerValidator, ctx, app, valAddr.Bytes())
		createValidator(t, ctx, app.StakingKeeper, valAddr, PKs[i], delCoin, commission)
	}
	require.EqualValues(t, len(validators), len(app.StakingKeeper.GetAllValidators(ctx)))
}

func setupStakingBondDenom(ctx sdk.Context, app *app.App) {
	stakeParams := app.StakingKeeper.GetParams(ctx)
	stakeParams.BondDenom = commontestutils.Denom
	app.StakingKeeper.SetParams(ctx, stakeParams)
}

func verifyUnbondingDelegations(t *testing.T, ctx sdk.Context, app *app.App, delegableAddress sdk.AccAddress,
	validators []sdk.ValAddress, unbondingAmount []int64) {

	unbondingDelegations := app.StakingKeeper.GetAllUnbondingDelegations(ctx, delegableAddress)
	require.EqualValues(t, len(validators), len(unbondingDelegations))

	for i, valAddr := range validators {
		found := false
		for _, delegation := range unbondingDelegations {
			if delegation.ValidatorAddress == valAddr.String() {
				require.EqualValues(t, 1, len(unbondingDelegations[0].Entries))
				require.EqualValues(t, sdk.NewInt(unbondingAmount[i]), unbondingDelegations[0].Entries[0].Balance)
				require.EqualValues(t, sdk.NewInt(unbondingAmount[i]), unbondingDelegations[0].Entries[0].InitialBalance)
				found = true
			}
		}
		require.True(t, found, "delegation not found. Validator Address: "+valAddr.String())
	}

}

func setupAccountsVestings(ctx sdk.Context, app *app.App, address string, numberOfVestings int, vestingAmount uint64, withdrawnAmount uint64) types.AccountVestings {
	return setupAccountsVestingsWithModification(ctx, app, func(*types.VestingPool) { /*do not modify*/ }, address, numberOfVestings, vestingAmount, withdrawnAmount)
}

func setupAccountsVestingsWithModification(ctx sdk.Context, app *app.App, modifyVesting func(*types.VestingPool), address string, numberOfVestings int, vestingAmount uint64, withdrawnAmount uint64) types.AccountVestings {
	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(numberOfVestings, 1, 1)
	accountVestings.Address = address

	for _, vesting := range accountVestings.VestingPools {
		vesting.Vested = sdk.NewIntFromUint64(vestingAmount)
		vesting.Withdrawn = sdk.NewIntFromUint64(withdrawnAmount)
		vesting.LastModificationVested = sdk.NewIntFromUint64(vestingAmount)
		vesting.LastModificationWithdrawn = sdk.NewIntFromUint64(withdrawnAmount)
		modifyVesting(vesting)
	}
	app.CfevestingKeeper.SetAccountVestings(ctx, accountVestings)
	return accountVestings
}

func allocateRewardsToValidator(ctx sdk.Context, app *app.App, validatorRewards uint64, valAddr sdk.ValAddress) {
	valCons := sdk.NewDecCoin(commontestutils.Denom, sdk.NewIntFromUint64(validatorRewards))
	val := app.StakingKeeper.Validator(ctx, valAddr)
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, sdk.NewDecCoins(valCons))
}

func verifyQueryRewards(t *testing.T, ctx sdk.Context, app *app.App, delegableAddr sdk.AccAddress, valAddr sdk.ValAddress, hasRewards bool, rewards uint64) {
	msgServerCtx := sdk.WrapSDKContext(ctx)
	query := distrtypes.QueryDelegationRewardsRequest{DelegatorAddress: delegableAddr.String(), ValidatorAddress: valAddr.String()}
	resp, _ := app.DistrKeeper.DelegationRewards(msgServerCtx, &query)
	if hasRewards {
		require.EqualValues(t, 1, len(resp.Rewards))
		require.EqualValues(t, sdk.NewDecFromInt(sdk.NewIntFromUint64(rewards)), resp.Rewards[0].Amount)
	} else {
		require.EqualValues(t, 0, len(resp.Rewards))
	}

}

func createVestingPool(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress, accountVestingsExistsBefore bool, accountVestingsExistsAfter bool,
	vestingPoolName string, lockupDuration time.Duration, vestingType types.VestingType, amountToVest int64, accAmountBefore int64, moduleAmountBefore int64,
	accAmountAfter int64, moduleAmountAfter int64) {

	_, accFound := app.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(t, accountVestingsExistsBefore, accFound)

	verifyAccountBalance(t, app, ctx, address, sdk.NewInt(accAmountBefore))
	moduleAccAddr := app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress()
	verifyAccountBalance(t, app, ctx, moduleAccAddr, sdk.NewInt(moduleAmountBefore))

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgCreateVestingPool{Creator: address.String(), Name: vestingPoolName,
		Amount: sdk.NewInt(amountToVest), Duration: lockupDuration, VestingType: vestingType.Name}
	_, error := msgServer.CreateVestingPool(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	_, accFound = app.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(t, accountVestingsExistsAfter, accFound)

	verifyAccountBalance(t, app, ctx, address, sdk.NewInt(accAmountAfter))
	verifyAccountBalance(t, app, ctx, moduleAccAddr, sdk.NewInt(moduleAmountAfter))
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

func verifyAccountVestingPools(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress,
	vestingNames []string, durations []time.Duration, vestingTypes []types.VestingType, vestedAmounts []int64, withdrawnAmounts []int64) {

	verifyAccountVestingsWithModification(t, ctx, app, address, 1, vestingNames, durations, vestingTypes, newTimeArray(len(vestingTypes), ctx.BlockTime()), vestedAmounts, withdrawnAmounts,
		newInts64Array(len(vestingTypes), 0), newTimeArray(len(vestingTypes), ctx.BlockTime()), vestedAmounts, withdrawnAmounts)
}

func verifyAccountVestingsWithModification(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress,
	amountOfAllAccVestings int, vestingNames []string, durations []time.Duration, vestingTypes []types.VestingType, startsTimes []time.Time, vestedAmounts []int64, withdrawnAmounts []int64,
	sentAmounts []int64, modificationsTimes []time.Time, modificationsVested []int64, modificationsWithdrawn []int64) {
	allAccVestings := app.CfevestingKeeper.GetAllAccountVestings(ctx)

	accVestings, accFound := app.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, amountOfAllAccVestings, len(allAccVestings))
	require.EqualValues(t, len(vestingTypes), len(accVestings.VestingPools))

	require.EqualValues(t, address.String(), accVestings.Address)

	for i, vesting := range accVestings.VestingPools {
		found := false
		if vesting.Id == int32(i+1) {
			require.EqualValues(t, i+1, vesting.Id)
			require.EqualValues(t, vestingNames[i], vesting.Name)
			require.EqualValues(t, vestingTypes[i].Name, vesting.VestingType)
			require.EqualValues(t, true, startsTimes[i].Equal(vesting.LockStart))
			require.EqualValues(t, true, ctx.BlockTime().Add(durations[i]).Equal(vesting.LockEnd))
			require.EqualValues(t, sdk.NewInt(vestedAmounts[i]), vesting.Vested)
			require.EqualValues(t, sdk.NewInt(withdrawnAmounts[i]), vesting.Withdrawn)

			require.EqualValues(t, sdk.NewInt(sentAmounts[i]), vesting.Sent)
			require.EqualValues(t, true, modificationsTimes[i].Equal(vesting.LastModification))
			require.EqualValues(t, sdk.NewInt(modificationsVested[i]), vesting.LastModificationVested)
			require.EqualValues(t, sdk.NewInt(modificationsWithdrawn[i]), vesting.LastModificationWithdrawn)
			found = true

		}
		require.True(t, found, "not found vesting id: "+strconv.Itoa(i+1))

	}

}

func setupVestingTypes(ctx sdk.Context, app *app.App, numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) types.VestingTypes {
	return setupVestingTypesWithModification(ctx, app, func(*types.VestingType) { /* do not modify */ }, numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
}

func setupVestingTypesWithModification(ctx sdk.Context, app *app.App, modifyVestingType func(*types.VestingType), numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) types.VestingTypes {
	vestingTypesArray := testutils.Generate10BasedVestingTypes(numberOfVestingTypes, amountOf10BasedVestingTypes, startId)
	for _, vestingType := range vestingTypesArray {
		modifyVestingType(vestingType)
	}
	vestingTypes := types.VestingTypes{VestingTypes: vestingTypesArray}
	app.CfevestingKeeper.SetVestingTypes(ctx, vestingTypes)
	return vestingTypes
}

func withdrawAllAvailable(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress, accountBalanceBefore int64, moduleBalanceBefore int64,
	accountBalanceAfter int64, moduleBalanceAfter int64) {

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	verifyAccountBalance(t, app, ctx, address, sdk.NewInt(accountBalanceBefore))
	verifyModuleAccount(t, ctx, app, sdk.NewInt(moduleBalanceBefore))
	msg := types.MsgWithdrawAllAvailable{Creator: address.String()}
	_, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)
	verifyAccountBalance(t, app, ctx, address, sdk.NewInt(accountBalanceAfter))
	verifyModuleAccount(t, ctx, app, sdk.NewInt(moduleBalanceAfter))
}

func withdrawAllAvailableDelegable(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress, delegableAddress sdk.AccAddress, accountBalanceBefore int64, delegableBalanceBefore int64,
	moduleBalanceBefore int64, accountBalanceAfter int64, delegableBalanceAfter int64, moduleBalanceAfter int64) {

	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableBalanceBefore))
	withdrawAllAvailable(t, ctx, app, address, accountBalanceBefore, moduleBalanceBefore, accountBalanceAfter, moduleBalanceAfter)
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableBalanceAfter))
}

func compareStoredAcountVestings(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress, accVestings types.AccountVestings) {
	storedAccVestings, accFound := app.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(t, true, accFound)

	testutils.AssertAccountVestings(t, accVestings, storedAccVestings)
}

func addCoinsToCfevestingModule(vested uint64, ctx sdk.Context, app *app.App) string {
	return commontestutils.AddCoinsToModuleByName(vested, types.ModuleName, ctx, app)
}

func getVestings(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress) *types.QueryVestingPoolsResponse {
	msgServerCtx := sdk.WrapSDKContext(ctx)
	vestingData, error := app.CfevestingKeeper.VestingPools(msgServerCtx, &types.QueryVestingPoolsRequest{Address: address.String()})
	require.EqualValues(t, nil, error)
	return vestingData
}

func verifyVestingResponseWithStoredAccountVestings(t *testing.T, ctx sdk.Context, app *app.App, response *types.QueryVestingPoolsResponse, address sdk.AccAddress, current time.Time, delegationAllowed bool) {
	accVests, found := app.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(t, true, found)
	verifyVestingResponse(t, response, accVests, current, delegationAllowed)
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
				require.EqualValues(t, "uc4e", response.VestingPools[0].Vested.Denom)
				require.EqualValues(t, vesting.Vested, response.VestingPools[0].Vested.Amount)
				require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn).String(), response.VestingPools[0].CurrentVestedAmount)
				require.EqualValues(t, vesting.Sent.String(), response.VestingPools[0].SentAmount)

				found = true
			}
		}
		require.True(t, found, "not found vesting nfo with Id: "+strconv.FormatInt(int64(vesting.Id), 10))
	}
}
