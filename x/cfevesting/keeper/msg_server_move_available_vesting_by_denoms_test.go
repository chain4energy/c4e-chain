package keeper_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
)

func TestMoveVestingByDenoms(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	initialAmount := sdk.NewInt(8999999999999999999)
	duration := 1000 * time.Hour

	amountToSend := sdk.NewInt(300)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]

	startTime := testHelper.Context.BlockTime()

	require.NoError(t, testHelper.AuthUtils.CreateDefaultDenomVestingAccount(accAddr.String(), initialAmount, startTime, startTime.Add(duration)))

	msgServer := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper)

	coins := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amountToSend))
	msgServer.SplitVesting(testHelper.WrappedContext, types.NewMsgSplitVesting(accAddr.String(), acountsAddresses[1].String(), coins))

	ownerAccount := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, acountsAddresses[1])
	require.NotNil(t, ownerAccount)

	_, ok := ownerAccount.(*vestingtypes.ContinuousVestingAccount)
	require.True(t, ok)

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(acountsAddresses[1], amountToSend)

	// require.True(t, amountToSend.Equal(testHelper.App.BankKeeper.LockedCoins(testHelper.Context, acountsAddresses[1]).AmountOf(testenv.DefaultTestDenom)))
	// _ = vestingAcc

	testHelper.AuthUtils.VerifyVestingAccount(acountsAddresses[1], testenv.DefaultTestDenom, amountToSend, startTime, startTime.Add(duration))
}
