package keeper_test

import (
	"cosmossdk.io/math"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func TestSplitVesting(t *testing.T) {
	duration := 1000 * time.Hour

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 1)
	srcAccAddr := acountsAddresses[0]
	dstAccAddr := acountsAddresses[1]

	startTime := testenv.TestEnvTime

	type Source int

	const (
		Normal Source = iota
		GenesisAccount
		AccountFromGenesisPool
		AccountFromGenesisAccount
	)

	for _, tc := range []struct {
		desc                 string
		initialVestingAmount sdk.Coins
		amountToSend         sdk.Coins
		blockTime            time.Time
		vAccStartTime        time.Time
		vestingDuration      time.Duration
		Source               Source
	}{
		{desc: "before vesting start - one denom", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "vesting start - one denom", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime,
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "after vesting start - one denom", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime.Add(duration / 2),
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "before vesting start - many denoms but one denom to split", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999), sdk.NewInt(300000)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "vesting start - many denoms but one denom to split", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999), sdk.NewInt(300000)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime,
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "after vesting start - many denoms but one denom to split", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999), sdk.NewInt(300000)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime.Add(duration / 2),
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "before vesting start - many denoms", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999), sdk.NewInt(300000), sdk.NewInt(700000)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300), sdk.NewInt(25)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "vesting start - many denoms", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999), sdk.NewInt(300000), sdk.NewInt(700000)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300), sdk.NewInt(25)}), blockTime: startTime,
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "after vesting start - many denoms", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999), sdk.NewInt(300000), sdk.NewInt(700000)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300), sdk.NewInt(25)}), blockTime: startTime.Add(duration / 2),
			vAccStartTime: startTime, vestingDuration: duration},
		{desc: "src acc from vesting pool", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, Source: GenesisAccount},
		{desc: "src acc from vesting pool", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, Source: AccountFromGenesisPool},
		{desc: "src acc from vesting pool", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, Source: AccountFromGenesisAccount},
		// TODO cases fo VestingAccountTrace
	} {
		t.Run(tc.desc, func(t *testing.T) {
			testHelper := testapp.SetupTestAppWithHeight(t, 1000)
			testHelper.SetContextBlockTime(tc.blockTime)
			require.NoError(t, testHelper.AuthUtils.CreateVestingAccount(srcAccAddr.String(), tc.initialVestingAmount, tc.vAccStartTime, tc.vAccStartTime.Add(tc.vestingDuration)))
			if tc.Source != Normal {
				testHelper.App.CfevestingKeeper.AppendVestingAccountTrace(testHelper.Context,
					types.VestingAccountTrace{
						Address:            srcAccAddr.String(),
						Genesis:            tc.Source == GenesisAccount,
						FromGenesisPool:    tc.Source == AccountFromGenesisPool,
						FromGenesisAccount: tc.Source == AccountFromGenesisAccount,
					},
				)
			}

			msgServer := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper)

			lockedBefore := testHelper.BankUtils.GetAccountLockedCoins(srcAccAddr)

			balancesBefore := testHelper.BankUtils.GetAccountAllBalances(srcAccAddr)
			_, err := msgServer.SplitVesting(testHelper.WrappedContext, types.NewMsgSplitVesting(srcAccAddr.String(), dstAccAddr.String(), tc.amountToSend))
			require.NoError(t, err)

			testHelper.BankUtils.VerifyLockedCoins(srcAccAddr, lockedBefore.Sub(tc.amountToSend...), true)
			testHelper.BankUtils.VerifyAccountBalances(srcAccAddr, balancesBefore.Sub(tc.amountToSend...), true)

			testHelper.AuthUtils.VerifyIsContinuousVestingAccount(dstAccAddr)

			testHelper.BankUtils.VerifyAccountBalances(dstAccAddr, tc.amountToSend, true)

			newAccStartTime := tc.vAccStartTime
			if tc.blockTime.After(newAccStartTime) {
				newAccStartTime = tc.blockTime
			}
			testHelper.AuthUtils.VerifyVestingAccount(dstAccAddr, tc.amountToSend, newAccStartTime, tc.vAccStartTime.Add(duration))
			trace, found := testHelper.App.CfevestingKeeper.GetVestingAccountTrace(testHelper.Context, dstAccAddr.String())
			require.Equal(t, tc.Source != Normal, found)
			if tc.Source != Normal {
				require.EqualValues(t, types.VestingAccountTrace{
					Id:                 1,
					Address:            dstAccAddr.String(),
					Genesis:            false,
					FromGenesisPool:    tc.Source == AccountFromGenesisPool,
					FromGenesisAccount: tc.Source == GenesisAccount || tc.Source == AccountFromGenesisAccount,
				}, trace)
			}
		})
	}
}

func TestSplitVestingError(t *testing.T) {
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
		amountToSend                      sdk.Coins
		blockTime                         time.Time
		vAccStartTime                     time.Time
		vestingDuration                   time.Duration
		errorMessage                      string
		disableSend                       bool
		createToAddressAccountBeforeSplit bool
	}{
		{desc: "wrong src addr", srcAddr: "invalid", dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting: from acc address error: decoding bech32 failed: invalid bech32 string length 7: failed to parse"},
		{desc: "wrong dst addr", srcAddr: srcAccAddr.String(), dstAddr: "invalid", initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(300)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting: to acc address error: decoding bech32 failed: invalid bech32 string length 7: failed to parse"},
		{desc: "wrong amount - single zero", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{sdk.ZeroInt()}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting - invalid amount (coin 0denom1 amount is not positive): wrong param value"},
		{desc: "wrong amount - single nil", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{{}}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting - amount cannot be nil: wrong param value"},
		{desc: "wrong amount - single less than zero", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(8999999999999999999)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(-1)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting - invalid amount (coin -1denom1 amount is not positive): wrong param value"},
		{desc: "wrong amount - zero among many", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(3), sdk.ZeroInt(), sdk.NewInt(5)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting - invalid amount (coin denom2 amount is not positive): wrong param value"},
		{desc: "wrong amount - nil among many", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(3), {}, sdk.NewInt(5)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting - amount cannot be nil: wrong param value"},
		{desc: "wrong amount - less than zero among many", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(3), sdk.NewInt(-1), sdk.NewInt(5)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting - invalid amount (coin denom2 amount is not positive): wrong param value"},
		{desc: "send disabled", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(3), sdk.NewInt(1), sdk.NewInt(5)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting: send is disabled: wrong param value", disableSend: true},
		{desc: "destination not allowed to received funds", srcAddr: srcAccAddr.String(), dstAddr: blockedAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(3), sdk.NewInt(1), sdk.NewInt(5)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting: " + blockedAddr.String() + " is not allowed to receive funds: unauthorized"},
		{desc: "destination account already exists", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), amountToSend: createDenomCoins([]math.Int{sdk.NewInt(3), sdk.NewInt(1), sdk.NewInt(5)}), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting: split vesting coins - account address: " + dstAccAddr.String() + ": entity already exists", createToAddressAccountBeforeSplit: true},
		{desc: "no coins to unlock", srcAddr: srcAccAddr.String(), dstAddr: dstAccAddr.String(), initialVestingAmount: createDenomCoins([]math.Int{sdk.NewInt(12), sdk.NewInt(12), sdk.NewInt(12)}), amountToSend: sdk.NewCoins(), blockTime: startTime.Add(-duration),
			vAccStartTime: startTime, vestingDuration: duration, errorMessage: "split vesting: split vesting coins - no coins to split : wrong param value"},
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
			_, err := msgServer.SplitVesting(testHelper.WrappedContext, types.NewMsgSplitVesting(tc.srcAddr, tc.dstAddr, tc.amountToSend))
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
