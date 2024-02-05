package v140_test

import (
	"cosmossdk.io/math"
	v140 "github.com/chain4energy/c4e-chain/app/upgrades/v140"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	cfedistributormoduletypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUpdateStrategicReserveShortTermPool_AccountVestingPoolsNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	err := v140.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v140.StrategicReservceShortTermPoolAccount)
	require.True(t, found)
	require.Equal(t, 0, len(accountVestingPools.VestingPools))
}

func TestUpdateStrategicReserveShortTermPool_Success(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v140.StrategicReservceShortTermPoolAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            v140.StrategicReservceShortTermPool,
				InitiallyLocked: math.NewInt(15000000000000),
			},
		},
	}
	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v140.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v140.StrategicReservceShortTermPoolAccount)
	require.True(t, found)
	require.Equal(t, 1, len(accountVestingPools.VestingPools))
	require.Equal(t, math.NewInt(5000000000000), accountVestingPools.VestingPools[0].InitiallyLocked)
}

func TestUpdateStrategicReserveAccount_AccountNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	err := v140.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	strategicReserveAccountAddress, _ := sdk.AccAddressFromBech32(v140.StrategicReservceAccount)
	strategicReserveAccount := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, strategicReserveAccountAddress)
	require.Nil(t, strategicReserveAccount)
}

func TestUpdateStrategicReserveAccount_AccountNotContinuousVestingAccount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	strategicReserveAccountAddress, _ := sdk.AccAddressFromBech32(v140.StrategicReservceAccount)
	strategicReserveAccount := &authtypes.BaseAccount{
		Address:       v140.StrategicReservceAccount,
		AccountNumber: 0,
		Sequence:      0,
	}
	testHelper.App.AccountKeeper.SetAccount(testHelper.Context, strategicReserveAccount)

	err := v140.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	reserveAccount := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, strategicReserveAccountAddress)
	require.NotNil(t, reserveAccount)
}

func TestUpdateStrategicReserveAccount_Success(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	strategicReserveAccountAddress, _ := sdk.AccAddressFromBech32(v140.StrategicReservceAccount)
	liquidityPoolOwnerAccountAddress, _ := sdk.AccAddressFromBech32(v140.LiquidityPoolOwner)
	strategicReserveAccount := &vestingtypes.ContinuousVestingAccount{
		BaseVestingAccount: &vestingtypes.BaseVestingAccount{
			BaseAccount: &authtypes.BaseAccount{
				Address:       v140.StrategicReservceAccount,
				AccountNumber: 0,
				Sequence:      0,
			},
			OriginalVesting:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(40000000000000))),
			DelegatedFree:    nil,
			DelegatedVesting: nil,
			EndTime:          testHelper.Context.BlockTime().Add(time.Hour).Unix(),
		},
	}
	testHelper.App.AccountKeeper.SetAccount(testHelper.Context, strategicReserveAccount)

	err := v140.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	strategicReserveAccount = testHelper.App.AccountKeeper.GetAccount(testHelper.Context, strategicReserveAccountAddress).(*vestingtypes.ContinuousVestingAccount)
	require.NotNil(t, strategicReserveAccount)
	require.Equal(t, math.NewInt(30000000000000), strategicReserveAccount.OriginalVesting.AmountOf(testenv.DefaultTestDenom))

	liquidityPoolOwnerBalance := testHelper.BankUtils.GetAccountAllBalances(liquidityPoolOwnerAccountAddress)
	require.Equal(t, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000000000000))), liquidityPoolOwnerBalance)

	strategicReserveAccountBalance := testHelper.BankUtils.GetAccountAllBalances(strategicReserveAccountAddress)
	require.Equal(t, sdk.NewCoins(), strategicReserveAccountBalance)
}

func TestUpdateCommunityPool_Success(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	feePool := cfedistributormoduletypes.FeePool{
		CommunityPool: sdk.NewDecCoinsFromCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(100000000000000))),
	}
	testHelper.DistributionUtils.DistrKeeper.SetFeePool(testHelper.Context, feePool)

	err := v140.UpdateCommunityPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	feePool = testHelper.DistributionUtils.DistrKeeper.GetFeePool(testHelper.Context)
	require.Equal(t, sdk.NewDecCoinsFromCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(40000000000000))), feePool.CommunityPool)

	communityPoolBalance := testHelper.DistributionUtils.DistrKeeper.GetFeePoolCommunityCoins(testHelper.Context)
	require.Equal(t, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000000000000))), communityPoolBalance)
}
