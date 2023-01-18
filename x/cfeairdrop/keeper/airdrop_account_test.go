package keeper_test

import (
	"fmt"
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	cfeairdropmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	cfeairdropmoduletypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func TestCreateAirdropAccount(t *testing.T) {
	startTime := commontestutils.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := commontestutils.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	moduleAmount := sdk.NewInt(10000)
	amount := sdk.NewInt(1000)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)

	testHelper.C4eAirdropUtils.SendToAirdropAccount(acountsAddresses[0],
		amount,
		startTimeUnix,
		endTimeUnix, true,
	)

	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount)
	testHelper.SetContextBlockTime(commontestutils.TestEnvTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount.QuoRaw(2))
	testHelper.SetContextBlockTime(endTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], sdk.ZeroInt())

	testHelper.SetContextBlockTime(startTime)
	testHelper.C4eAirdropUtils.SendToAirdropAccount(acountsAddresses[0],
		amount,
		startTimeUnix,
		endTimeUnix, false,
	)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount.MulRaw(2))
	testHelper.SetContextBlockTime(commontestutils.TestEnvTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount)
	testHelper.SetContextBlockTime(endTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], sdk.ZeroInt())

	testHelper.SetContextBlockTime(startTime)
	testHelper.C4eAirdropUtils.SendToAirdropAccount(acountsAddresses[0],
		amount,
		startTimeUnix,
		endTimeUnix, false,
	)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount.MulRaw(3))
	testHelper.SetContextBlockTime(commontestutils.TestEnvTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount.QuoRaw(2).MulRaw(3))
	testHelper.SetContextBlockTime(endTime)
	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], sdk.ZeroInt())
}

func TestCreateAirdropAccountSendDisabled(t *testing.T) {
	startTime := commontestutils.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := commontestutils.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	moduleAmount := sdk.NewInt(10000)
	amount := sdk.NewInt(1000)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)
	testHelper.BankUtils.DisableDefaultSend()
	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
		amount,
		startTimeUnix,
		endTimeUnix, true, "send to airdrop account - send coins disabled: uc4e transfers are currently disabled: send transactions are disabled",
		false,
	)
}

func TestCreateAirdropAccountBlockedAddress(t *testing.T) {
	startTime := commontestutils.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := commontestutils.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	blockedAccounts := testHelper.App.ModuleAccountAddrs()
	blockedAccounts[acountsAddresses[0].String()] = true
	testHelper.App.BankKeeper = bankkeeper.NewBaseKeeper(
		testHelper.App.AppCodec(), testHelper.App.GetKey(banktypes.StoreKey), testHelper.App.AccountKeeper, testHelper.App.GetSubspace(banktypes.ModuleName), blockedAccounts,
	)
	testHelper.App.CfeairdropKeeper = *cfeairdropmodulekeeper.NewKeeper(
		testHelper.App.AppCodec(),
		testHelper.App.GetKey(cfeairdropmoduletypes.StoreKey),
		testHelper.App.GetKey(cfeairdropmoduletypes.MemStoreKey),
		testHelper.App.GetSubspace(cfeairdropmoduletypes.ModuleName),

		testHelper.App.AccountKeeper,
		testHelper.App.BankKeeper,
		testHelper.App.FeeGrantKeeper,
	)

	moduleAmount := sdk.NewInt(10000)
	amount := sdk.NewInt(1000)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)
	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
		amount,
		startTimeUnix,
		endTimeUnix, true,
		fmt.Sprintf("send to airdrop account - account address: %s is not allowed to receive funds error: unauthorized", acountsAddresses[0]),
		false,
	)
}

func TestCreateAirdropAccountNotExist(t *testing.T) {
	startTime := commontestutils.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := commontestutils.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	moduleAmount := sdk.NewInt(10000)
	amount := sdk.NewInt(1000)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)
	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
		amount,
		startTimeUnix,
		endTimeUnix, false, fmt.Sprintf("create airdrop account - account does not exist: %s: entity does not exist", acountsAddresses[0]),
		false,
	)
}

func TestCreateAirdropAccountWrongAccountType(t *testing.T) {
	startTime := commontestutils.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := commontestutils.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	baseAccount := testHelper.App.AccountKeeper.NewAccountWithAddress(testHelper.Context, acountsAddresses[0])
	testHelper.App.AccountKeeper.SetAccount(testHelper.Context, baseAccount)
	moduleAmount := sdk.NewInt(10000)
	amount := sdk.NewInt(1000)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)
	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
		amount,
		startTimeUnix,
		endTimeUnix, false, "send to airdrop account - expected RepeatedContinuousVestingAccount, got: *types.BaseAccount: invalid account type",
		false,
	)
	// TODO: verify
	//testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
	//	amount,
	//	startTimeUnix,
	//	endTimeUnix, true, "send to airdrop account - expected RepeatedContinuousVestingAccount, got: *types.BaseAccount: invalid account type",
	//	false,
	//)
}

func TestCreateAirdropAccountSendError(t *testing.T) {
	startTime := commontestutils.TestEnvTime.Add(-24 * 100 * time.Hour)
	endTime := commontestutils.TestEnvTime.Add(24 * 100 * time.Hour)
	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	// moduleAmount := sdk.NewInt(10000)
	amount := sdk.NewInt(1000)

	startTimeUnix := startTime.Unix()
	endTimeUnix := endTime.Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(amount, types.ModuleName)

	// testHelper.C4eAirdropUtils.SendToAirdropAccount(acountsAddresses[0],
	// 	amount.MulRaw(2),
	// 	startTimeUnix,
	// 	endTimeUnix, true,
	// )
	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
		amount.AddRaw(1),
		startTimeUnix,
		endTimeUnix, true, "send to airdrop account - send coins to airdrop account error (to: cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq, amount: 1001uc4e): 1000uc4e is smaller than 1001uc4e: insufficient funds: failed to send coins",
		true,
	)

	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
		amount.AddRaw(1),
		startTimeUnix,
		endTimeUnix, false, "send to airdrop account - send coins to airdrop account error (to: cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq, amount: 1001uc4e): 1000uc4e is smaller than 1001uc4e: insufficient funds: failed to send coins",
		false,
	)
}
