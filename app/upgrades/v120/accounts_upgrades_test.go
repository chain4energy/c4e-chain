package v120_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	v120 "github.com/chain4energy/c4e-chain/app/upgrades/v120"
	"github.com/chain4energy/c4e-chain/testutil/testapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/stretchr/testify/require"
)

var ()

var (
	endTimeUnix   = int64(1758758400)
	startTimeUnix = int64(1695686400)

	newEndTimeUnix   = endTimeUnix + 365*24*3600
	newStartTimeUnix = startTimeUnix + 366*24*3600
)

func TestModifyVestingAccounts(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	endTime := time.Unix(endTimeUnix, 0)
	startTime := time.Unix(startTimeUnix, 0)

	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(v120.Account1, math.NewInt(8899990000000), startTime, endTime)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(v120.Account2, math.NewInt(6574990000000), startTime, endTime)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(v120.Account3, math.NewInt(6574990000000), startTime, endTime)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(v120.Account4, math.NewInt(2949990000000), startTime, endTime)

	acc1 := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, sdk.MustAccAddressFromBech32(v120.Account1)).(*vestingtypes.ContinuousVestingAccount)
	acc2 := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, sdk.MustAccAddressFromBech32(v120.Account2)).(*vestingtypes.ContinuousVestingAccount)
	acc3 := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, sdk.MustAccAddressFromBech32(v120.Account3)).(*vestingtypes.ContinuousVestingAccount)
	acc4 := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, sdk.MustAccAddressFromBech32(v120.Account4)).(*vestingtypes.ContinuousVestingAccount)

	require.NoError(t, v120.ModifyVestingAccountsState(testHelper.Context, testHelper.App))

	acc1 = updateAcc(acc1)
	acc2 = updateAcc(acc2)
	acc3 = updateAcc(acc3)
	acc4 = updateAcc(acc4)

	acc1After := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, sdk.MustAccAddressFromBech32(v120.Account1)).(*vestingtypes.ContinuousVestingAccount)
	acc2After := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, sdk.MustAccAddressFromBech32(v120.Account2)).(*vestingtypes.ContinuousVestingAccount)
	acc3After := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, sdk.MustAccAddressFromBech32(v120.Account3)).(*vestingtypes.ContinuousVestingAccount)
	acc4After := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, sdk.MustAccAddressFromBech32(v120.Account4)).(*vestingtypes.ContinuousVestingAccount)

	require.EqualValues(t, acc1, acc1After)
	require.EqualValues(t, acc2, acc2After)
	require.EqualValues(t, acc3, acc3After)
	require.EqualValues(t, acc4, acc4After)

}

func updateAcc(account *vestingtypes.ContinuousVestingAccount) *vestingtypes.ContinuousVestingAccount {
	account.StartTime = newStartTimeUnix
	account.EndTime = newEndTimeUnix
	return account
}
