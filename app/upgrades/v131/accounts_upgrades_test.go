package v131_test

import (
	"cosmossdk.io/math"
	v131 "github.com/chain4energy/c4e-chain/app/upgrades/v131"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUpdateStrategicReserveShortTermPool_AccountVestingPoolsNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	err := v131.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v131.StrategicReservceShortTermPoolAccount)
	require.False(t, found)
	require.Equal(t, 0, len(accountVestingPools.VestingPools))

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveShortTermPool_InitiallyLockedNegative(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v131.StrategicReservceShortTermPoolAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            v131.StrategicReservceShortTermPool,
				InitiallyLocked: math.NewInt(1_000),
				VestingType:     "Strategic reserve short term round",
			},
		},
	}
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1_000), cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v131.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v131.StrategicReservceShortTermPoolAccount)
	require.True(t, found)
	require.Equal(t, 1, len(accountVestingPools.VestingPools))
	require.Equal(t, accountVestingPools.VestingPools[0].InitiallyLocked, math.NewInt(1_000))

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveShortTermPool_strategicReserveShortTermPoolNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v131.StrategicReservceShortTermPoolAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            "Test pool",
				InitiallyLocked: math.NewInt(1_000),
				VestingType:     "Strategic reserve short term round",
			},
		},
	}
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1_000), cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v131.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v131.StrategicReservceShortTermPoolAccount)
	require.True(t, found)
	require.Equal(t, 1, len(accountVestingPools.VestingPools))
	require.Equal(t, accountVestingPools.VestingPools[0].InitiallyLocked, math.NewInt(1_000))

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveShortTermPool_BurnCoinsError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v131.StrategicReservceShortTermPoolAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            v131.StrategicReservceShortTermPool,
				InitiallyLocked: math.NewInt(40_000_000_000_000),
				VestingType:     "Strategic reserve short term round",
			},
		},
	}

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1), cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v131.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.EqualError(t, err, "1uc4e is smaller than 20000000000000uc4e: insufficient funds")
}

func TestUpdateStrategicReserveShortTermPool_Success(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v131.StrategicReservceShortTermPoolAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            v131.StrategicReservceShortTermPool,
				InitiallyLocked: math.NewInt(40_000_000_000_000),
				VestingType:     "Strategic reserve short term round",
			},
		},
	}
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(40_000_000_000_000), cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v131.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v131.StrategicReservceShortTermPoolAccount)
	require.True(t, found)
	require.Equal(t, 1, len(accountVestingPools.VestingPools))
	require.Equal(t, math.NewInt(20_000_000_000_000), accountVestingPools.VestingPools[0].InitiallyLocked)

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveAccount_AccountNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	err := v131.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	strategicReserveAccountAddress, _ := sdk.AccAddressFromBech32(v131.StrategicReserveAccount)
	strategicReserveAccount := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, strategicReserveAccountAddress)
	require.Nil(t, strategicReserveAccount)

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveAccount_NotContinuousVestingAccount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	err := testHelper.AuthUtils.CreateBaseAccount(testHelper.Context, v131.StrategicReserveAccount, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(79_999_990_000_000))))
	require.NoError(t, err)

	err = v131.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	strategicReserveAccountAddress, err := sdk.AccAddressFromBech32(v131.StrategicReserveAccount)
	require.NoError(t, err)

	account := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, strategicReserveAccountAddress)
	_, ok := account.(*types.BaseAccount)
	require.True(t, ok)
	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveAccount_NegativeOriginalVesting(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	err := testHelper.AuthUtils.CreateVestingAccount(v131.StrategicReserveAccount,
		sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(200_000))),
		time.Unix(1727222400, 0), time.Unix(1821830400, 0))
	require.NoError(t, err)

	err = v131.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	strategicReserveAccountAddress, err := sdk.AccAddressFromBech32(v131.StrategicReserveAccount)
	require.NoError(t, err)

	testHelper.BankUtils.VerifyAccountDefaultDenomBalance(testHelper.Context, strategicReserveAccountAddress, math.NewInt(200_000))
	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveAccount_SendCoinsError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	err := testHelper.AuthUtils.CreateVestingAccount(v131.StrategicReserveAccount,
		sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(31_000_000_000_000))),
		time.Unix(1727222400, 0), time.Unix(1821830400, 0))
	require.NoError(t, err)

	err = v131.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	strategicReserveAccountAddress, err := sdk.AccAddressFromBech32(v131.StrategicReserveAccount)
	require.NoError(t, err)

	testHelper.BankUtils.VerifyAccountDefaultDenomBalance(testHelper.Context, strategicReserveAccountAddress, math.NewInt(31_000_000_000_000))
	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveAccount_Success(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	strategicReserveAccountAddress, _ := sdk.AccAddressFromBech32(v131.StrategicReserveAccount)
	liquidityPoolOwnerAccountAddress, _ := sdk.AccAddressFromBech32(v131.LiquidityPoolOwner)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(math.NewInt(10000000000000), liquidityPoolOwnerAccountAddress)

	addStrategicReserveVestingAccount(t, testHelper)
	err := v131.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	strategicReserveAccount := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, strategicReserveAccountAddress).(*vestingtypes.ContinuousVestingAccount)
	require.NotNil(t, strategicReserveAccount)
	require.Equal(t, math.NewInt(50000000000000), strategicReserveAccount.OriginalVesting.AmountOf(testenv.DefaultTestDenom))

	liquidityPoolOwnerBalance := testHelper.BankUtils.GetAccountAllBalances(liquidityPoolOwnerAccountAddress)
	require.Equal(t, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(20000000000000))), liquidityPoolOwnerBalance)

	strategicReserveAccountBalance := testHelper.BankUtils.GetAccountAllBalances(strategicReserveAccountAddress)
	require.Equal(t, math.NewInt(50000000000000), strategicReserveAccountBalance.AmountOf(testenv.DefaultTestDenom))

	testHelper.ValidateGenesisAndInvariants()
}

func addStrategicReserveVestingAccount(t *testing.T, testHelper *testapp.TestHelper) {
	err := testHelper.AuthUtils.CreateVestingAccount(v131.StrategicReserveAccount, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(79_999_990_000_000))), time.Unix(1727222400, 0), time.Unix(1821830400, 0))
	require.NoError(t, err)
}

func TestUpdateCommunityPool_SuccessLatestValue(t *testing.T) {
	CommunityPoolBeforeAmount := sdk.MustNewDecFromStr("100000000083683.601351991745061083")
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	feePool := distrtypes.FeePool{
		CommunityPool: sdk.NewDecCoins(sdk.NewDecCoinFromDec(testenv.DefaultTestDenom, CommunityPoolBeforeAmount)),
	}
	testHelper.DistributionUtils.DistrKeeper.SetFeePool(testHelper.Context, feePool)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(CommunityPoolBeforeAmount.TruncateInt(), distrtypes.ModuleName)

	err := v131.UpdateCommunityPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	feePool = testHelper.DistributionUtils.DistrKeeper.GetFeePool(testHelper.Context)
	require.Equal(t, sdk.NewDecCoinsFromCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(40000000000000))), feePool.CommunityPool)

	communityPoolBalance := testHelper.DistributionUtils.DistrKeeper.GetFeePoolCommunityCoins(testHelper.Context)
	require.Equal(t, sdk.NewDecCoins(sdk.NewDecCoin(testenv.DefaultTestDenom, math.NewInt(40_000_000_000_000))), communityPoolBalance)

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateCommunityPool_SuccessDifferentValues(t *testing.T) {
	for _, tc := range []struct {
		amount sdk.Dec
	}{
		{amount: sdk.MustNewDecFromStr("100000000083683.1")},
		{amount: sdk.MustNewDecFromStr("100000000083683.2")},
		{amount: sdk.MustNewDecFromStr("100000000083683.3")},
		{amount: sdk.MustNewDecFromStr("100000000083683.4")},
		{amount: sdk.MustNewDecFromStr("100000000083683.5")},
		{amount: sdk.MustNewDecFromStr("100000000083683.6")},
		{amount: sdk.MustNewDecFromStr("100000000083683.7")},
		{amount: sdk.MustNewDecFromStr("100000000083683.8")},
		{amount: sdk.MustNewDecFromStr("100000000083683.9")},
		{amount: sdk.MustNewDecFromStr("100000000083683.1")},
		{amount: sdk.MustNewDecFromStr("100000000073683.13")},
		{amount: sdk.MustNewDecFromStr("100000000063683.15")},
		{amount: sdk.MustNewDecFromStr("100000000053683.17")},
		{amount: sdk.MustNewDecFromStr("100000000043683.19")},
		{amount: sdk.MustNewDecFromStr("100000000033683.191")},
	} {
		t.Run(tc.amount.String(), func(t *testing.T) {
			testHelper := testapp.SetupTestAppWithHeight(t, 1000)
			feePool := distrtypes.FeePool{
				CommunityPool: sdk.NewDecCoins(sdk.NewDecCoinFromDec(testenv.DefaultTestDenom, tc.amount)),
			}
			testHelper.DistributionUtils.DistrKeeper.SetFeePool(testHelper.Context, feePool)
			testHelper.BankUtils.AddDefaultDenomCoinsToModule(tc.amount.TruncateInt(), distrtypes.ModuleName)

			err := v131.UpdateCommunityPool(testHelper.Context, testHelper.App)
			require.NoError(t, err)

			feePool = testHelper.DistributionUtils.DistrKeeper.GetFeePool(testHelper.Context)
			require.Equal(t, sdk.NewDecCoinsFromCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(40000000000000))), feePool.CommunityPool)

			communityPoolBalance := testHelper.DistributionUtils.DistrKeeper.GetFeePoolCommunityCoins(testHelper.Context)
			require.Equal(t, sdk.NewDecCoins(sdk.NewDecCoin(testenv.DefaultTestDenom, math.NewInt(40_000_000_000_000))), communityPoolBalance)

			testHelper.ValidateGenesisAndInvariants()
		})
	}
}

func TestUpdateCommunityPool_AmountLower(t *testing.T) {
	CommunityPoolBeforeAmount := sdk.NewDec(39_000_000_000_000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	feePool := distrtypes.FeePool{
		CommunityPool: sdk.NewDecCoins(sdk.NewDecCoinFromDec(testenv.DefaultTestDenom, CommunityPoolBeforeAmount)),
	}
	testHelper.DistributionUtils.DistrKeeper.SetFeePool(testHelper.Context, feePool)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(CommunityPoolBeforeAmount.TruncateInt(), distrtypes.ModuleName)

	err := v131.UpdateCommunityPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	feePool = testHelper.DistributionUtils.DistrKeeper.GetFeePool(testHelper.Context)
	require.Equal(t, sdk.NewDecCoinsFromCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(39_000_000_000_000))), feePool.CommunityPool)

	communityPoolBalance := testHelper.DistributionUtils.DistrKeeper.GetFeePoolCommunityCoins(testHelper.Context)
	require.Equal(t, sdk.NewDecCoins(sdk.NewDecCoin(testenv.DefaultTestDenom, math.NewInt(39_000_000_000_000))), communityPoolBalance)

	testHelper.ValidateGenesisAndInvariants()
}

func addVestingTypes(testHelper *testapp.TestHelper) {
	vestingTypes := cfevestingtypes.VestingTypes{
		VestingTypes: []*cfevestingtypes.VestingType{{
			Name:          "Strategic reserve short term round",
			Free:          sdk.MustNewDecFromStr("0.10"),
			LockupPeriod:  222 * 24 * time.Hour,
			VestingPeriod: 365 * 24 * time.Hour,
		}},
	}
	testHelper.App.CfevestingKeeper.SetVestingTypes(testHelper.Context, vestingTypes)
}
