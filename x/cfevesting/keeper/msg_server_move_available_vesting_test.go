package keeper_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
)

func TestMoveAvailableVesting(t *testing.T) {
	duration := 1000 * time.Hour

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 1)
	srcAccAddr := acountsAddresses[0]
	dstAccAddr := acountsAddresses[1]

	startTime := testenv.TestEnvTime

	for _, tc := range []struct {
		desc                 string
		initialVestingAmount sdk.Coins
		blockTime            time.Time
		vAccStartTime        time.Time
		vestingDuration      time.Duration
	}{
		{desc: "before vesting start - one denom", initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(8999999999999999999)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "vesting start - one denom", initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(8999999999999999999)}), blockTime: startTime,
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "after vesting start - one denom", initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(8999999999999999999)}), blockTime: startTime.Add(duration / 2),
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "before vesting start - many denoms", initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(8999999999999999999), sdk.NewInt(300000), sdk.NewInt(700000)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "vesting start - many denoms", initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(8999999999999999999), sdk.NewInt(300000), sdk.NewInt(700000)}), blockTime: startTime,
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "after vesting start - many denoms", initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(8999999999999999999), sdk.NewInt(300000), sdk.NewInt(700000)}), blockTime: startTime.Add(duration / 2),
			vAccStartTime: startTime, vestingDuration: duration},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			testHelper := testapp.SetupTestAppWithHeight(t, 1000)
			testHelper.SetContextBlockTime(tc.blockTime)
			require.NoError(t, testHelper.AuthUtils.CreateVestingAccount(srcAccAddr.String(), tc.initialVestingAmount, tc.vAccStartTime, tc.vAccStartTime.Add(tc.vestingDuration)))

			msgServer := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper)

			lockedBefore := testHelper.BankUtils.GetAccountLockedCoins(srcAccAddr)

			balancesBefore := testHelper.BankUtils.GetAccountAllBalances(srcAccAddr)
			_, err := msgServer.MoveAvailableVesting(testHelper.WrappedContext, types.NewMsgMoveAvailableVesting(srcAccAddr.String(), dstAccAddr.String()))
			require.NoError(t, err)

			testHelper.BankUtils.VerifyLockedCoins(srcAccAddr, sdk.NewCoins(), true)
			testHelper.BankUtils.VerifyAccountBalances(srcAccAddr, balancesBefore.Sub(lockedBefore...), true)

			testHelper.AuthUtils.VerifyIsContinuousVestingAccount(dstAccAddr)

			testHelper.BankUtils.VerifyAccountBalances(dstAccAddr, lockedBefore, true)

			newAccStartTime := tc.vAccStartTime
			if tc.blockTime.After(newAccStartTime) {
				newAccStartTime = tc.blockTime
			}
			testHelper.AuthUtils.VerifyVestingAccount(dstAccAddr, lockedBefore, newAccStartTime, tc.vAccStartTime.Add(duration))

		})
	}
}

func TestMoveAvailableVestingError(t *testing.T) {
	duration := 1000 * time.Hour

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 1)
	srcAccAddr := acountsAddresses[0]
	dstAccAddr := acountsAddresses[1]
	blockedAddr := authtypes.NewModuleAddress(types.ModuleName)
	startTime := testenv.TestEnvTime

	for _, tc := range []struct {
		desc                              string
		srcAddr                           string
		dstAddr                           string
		initialVestingAmount              sdk.Coins
		blockTime                         time.Time
		vAccStartTime                     time.Time
		vestingDuration                   time.Duration
		errorMessage                      string
		disableSend                       bool
		createToAddressAccountBeforeSplit bool
	}{
		{desc: "wrong src addr", srcAddr: "invalid", dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(8999999999999999999)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "move available vesting - error parsing from address: invalid: decoding bech32 failed: invalid bech32 string length 7"},
		{desc: "wrong dst addr", srcAddr: srcAccAddr.String(), dstAddr: "invalid", initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(8999999999999999999)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "move available vesting: split vesting coins - error parsing to address: invalid: decoding bech32 failed: invalid bech32 string length 7"},
		{desc: "send disabled", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "move available vesting: denom1 transfers are currently disabled: send transactions are disabled", disableSend: true},
		{desc: "destination not allowed to received funds", srcAddr: srcAccAddr.String(), dstAddr: blockedAddr.String(), initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "move available vesting: " + blockedAddr.String() + " is not allowed to receive funds: unauthorized"},
		{desc: "destination account already exists", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "move available vesting: split vesting coins - account address: " + dstAccAddr.String() + ": entity already exists", createToAddressAccountBeforeSplit: true},
		{desc: "nothing to share", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]sdk.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), blockTime: startTime.Add(duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "move available vesting: split vesting coins - no coins to split : wrong param value"},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			testHelper := testapp.SetupTestAppWithHeight(t, 1000)
			testHelper.SetContextBlockTime(tc.blockTime)
			require.NoError(t, testHelper.AuthUtils.CreateVestingAccount(srcAccAddr.String(), tc.initialVestingAmount, tc.vAccStartTime, tc.vAccStartTime.Add(tc.vestingDuration)))

			msgServer := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper)

			lockedBefore := testHelper.BankUtils.GetAccountLockedCoins(srcAccAddr)

			if tc.disableSend {
				testHelper.BankUtils.DisableSend()
			}

			if tc.createToAddressAccountBeforeSplit {
				testHelper.AuthUtils.CreateDefaultDenomBaseAccount(tc.dstAddr, sdk.NewInt(1))
			}

			balancesBefore := testHelper.BankUtils.GetAccountAllBalances(srcAccAddr)
			_, err := msgServer.MoveAvailableVesting(testHelper.WrappedContext, types.NewMsgMoveAvailableVesting(tc.srcAddr, tc.dstAddr))
			require.Error(t, err)
			require.EqualError(t, err, tc.errorMessage)

			testHelper.BankUtils.VerifyLockedCoins(srcAccAddr, lockedBefore, true)
			testHelper.BankUtils.VerifyAccountBalances(srcAccAddr, balancesBefore, true)

			if !tc.createToAddressAccountBeforeSplit {
				testHelper.AuthUtils.VerifyAccountDoesNotExist(dstAccAddr)
			}

		})
	}
}
