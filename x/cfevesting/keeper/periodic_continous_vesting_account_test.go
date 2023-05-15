package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfevestingmodulekeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"testing"
	"time"
)

func TestCreateAccount(t *testing.T) {
	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	moduleAmount := math.NewInt(10000)
	amount := math.NewInt(1000)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.SendToRepeatedContinuousVestingAccount(testHelper.Context, acountsAddresses[0],
		amount,
		sdk.ZeroDec(),
		startTimeUnix,
		endTimeUnix,
	)

	testHelper.BankUtils.VerifyAccountDefultDenomLocked(testHelper.Context, acountsAddresses[0], amount)
	testHelper.SetContextBlockTime(testenv.TestEnvTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(testHelper.Context, acountsAddresses[0], amount.QuoRaw(2))
	testHelper.SetContextBlockTime(endTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(testHelper.Context, acountsAddresses[0], math.ZeroInt())

	testHelper.SetContextBlockTime(startTime)
	testHelper.C4eVestingUtils.SendToRepeatedContinuousVestingAccount(testHelper.Context, acountsAddresses[0],
		amount,
		sdk.ZeroDec(),
		startTimeUnix,
		endTimeUnix,
	)

	testHelper.BankUtils.VerifyAccountDefultDenomLocked(testHelper.Context, acountsAddresses[0], amount.MulRaw(2))
	testHelper.SetContextBlockTime(testenv.TestEnvTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(testHelper.Context, acountsAddresses[0], amount)
	testHelper.SetContextBlockTime(endTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(testHelper.Context, acountsAddresses[0], math.ZeroInt())

	testHelper.SetContextBlockTime(startTime)
	testHelper.C4eVestingUtils.SendToRepeatedContinuousVestingAccount(testHelper.Context, acountsAddresses[0],
		amount,
		sdk.ZeroDec(),
		startTimeUnix,
		endTimeUnix,
	)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(testHelper.Context, acountsAddresses[0], amount.MulRaw(3))
	testHelper.SetContextBlockTime(testenv.TestEnvTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(testHelper.Context, acountsAddresses[0], amount.QuoRaw(2).MulRaw(3))
	testHelper.SetContextBlockTime(endTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(testHelper.Context, acountsAddresses[0], math.ZeroInt())
}

func TestCreateAccountSendDisabled(t *testing.T) {
	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	moduleAmount := math.NewInt(10000)
	amount := math.NewInt(1000)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, cfevestingtypes.ModuleName)
	testHelper.BankUtils.DisableSend()

	testHelper.C4eVestingUtils.SendToRepeatedContinuousVestingAccountError(testHelper.Context, acountsAddresses[0],
		amount,
		sdk.ZeroDec(),
		startTimeUnix,
		endTimeUnix, true, fmt.Sprintf("%s transfers are currently disabled: send transactions are disabled", testenv.DefaultTestDenom),
	)
}

func TestCreateAccountBlockedAddress(t *testing.T) {
	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	blockedAccounts := testHelper.App.ModuleAccountAddrs()
	blockedAccounts[acountsAddresses[0].String()] = true
	testHelper.App.BankKeeper = bankkeeper.NewBaseKeeper(
		testHelper.App.AppCodec(), testHelper.App.GetKey(banktypes.StoreKey), testHelper.App.AccountKeeper, testHelper.App.GetSubspace(banktypes.ModuleName), blockedAccounts,
	)

	testHelper.App.CfevestingKeeper = *cfevestingmodulekeeper.NewKeeper(
		testHelper.App.AppCodec(),
		testHelper.App.GetKey(cfevestingtypes.StoreKey),
		testHelper.App.GetKey(cfevestingtypes.MemStoreKey),
		testHelper.App.GetSubspace(cfevestingtypes.ModuleName),
		testHelper.App.BankKeeper,
		testHelper.App.StakingKeeper,
		testHelper.App.AccountKeeper,
		testHelper.App.DistrKeeper,
		testHelper.App.GovKeeper,
		appparams.GetAuthority(),
	)
	moduleAmount := math.NewInt(10000)
	amount := math.NewInt(1000)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, cfevestingtypes.ModuleName)
	testHelper.C4eVestingUtils.SendToRepeatedContinuousVestingAccountError(testHelper.Context, acountsAddresses[0],
		amount,
		sdk.ZeroDec(),
		startTimeUnix,
		endTimeUnix, true,
		fmt.Sprintf("account address: %s is not allowed to receive funds error: unauthorized", acountsAddresses[0]),
	)
}

func TestCreateAccountWrongAccountType(t *testing.T) {
	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	account := testHelper.App.AccountKeeper.NewAccountWithAddress(testHelper.Context, acountsAddresses[0])
	baseAccount, _ := account.(*authtypes.BaseAccount)
	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount, sdk.NewCoins(), time.Now().Add(time.Hour).Unix())
	testHelper.App.AccountKeeper.SetAccount(testHelper.Context, baseVestingAccount)
	moduleAmount := math.NewInt(10000)
	amount := math.NewInt(100)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, cfevestingtypes.ModuleName)
	testHelper.C4eVestingUtils.SendToRepeatedContinuousVestingAccountError(testHelper.Context, acountsAddresses[0],
		amount,
		sdk.ZeroDec(),
		startTimeUnix,
		endTimeUnix, false, "account already exists and is not of PeriodicContinuousVestingAccount nor BaseAccount type, got: *types.BaseVestingAccount: invalid account type",
	)
}

func TestCreateAccountSendError(t *testing.T) {
	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	amount := math.NewInt(10000000000)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(amount, cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.SendToRepeatedContinuousVestingAccountError(testHelper.Context, acountsAddresses[0],
		amount.AddRaw(1),
		sdk.ZeroDec(),
		startTimeUnix,
		endTimeUnix, true, "10000000000uc4e is smaller than 10000000001uc4e: insufficient funds",
	)
}
