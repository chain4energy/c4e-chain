package v140_test

import (
	"cosmossdk.io/math"
	v140 "github.com/chain4energy/c4e-chain/app/upgrades/v140"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	CommunityPoolBeforeAmount = sdk.MustNewDecFromStr("100000000083683.601351991745061083")
)

func TestUpdateStrategicReserveShortTermPool_AccountVestingPoolsNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	err := v140.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v140.StrategicReservceShortTermPoolAccount)
	require.False(t, found)
	require.Equal(t, 0, len(accountVestingPools.VestingPools))

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveShortTermPool_Success(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v140.StrategicReservceShortTermPoolAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            v140.StrategicReservceShortTermPool,
				InitiallyLocked: math.NewInt(40_000_000_000_000),
				VestingType:     "Strategic reserve short term round",
			},
		},
	}
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(40_000_000_000_000), cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v140.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v140.StrategicReservceShortTermPoolAccount)
	require.True(t, found)
	require.Equal(t, 1, len(accountVestingPools.VestingPools))
	require.Equal(t, math.NewInt(20_000_000_000_000), accountVestingPools.VestingPools[0].InitiallyLocked)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveAccount_AccountNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	err := v140.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	strategicReserveAccountAddress, _ := sdk.AccAddressFromBech32(v140.StrategicReserveAccount)
	strategicReserveAccount := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, strategicReserveAccountAddress)
	require.Nil(t, strategicReserveAccount)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveAccount_Success(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	strategicReserveAccountAddress, _ := sdk.AccAddressFromBech32(v140.StrategicReserveAccount)
	liquidityPoolOwnerAccountAddress, _ := sdk.AccAddressFromBech32(v140.LiquidityPoolOwner)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(math.NewInt(10000000000000), liquidityPoolOwnerAccountAddress)

	addStrategicReserveVestingAccount(t, testHelper)
	err := v140.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	strategicReserveAccount := testHelper.App.AccountKeeper.GetAccount(testHelper.Context, strategicReserveAccountAddress).(*vestingtypes.ContinuousVestingAccount)
	require.NotNil(t, strategicReserveAccount)
	require.Equal(t, math.NewInt(49999990000000), strategicReserveAccount.OriginalVesting.AmountOf(testenv.DefaultTestDenom))

	liquidityPoolOwnerBalance := testHelper.BankUtils.GetAccountAllBalances(liquidityPoolOwnerAccountAddress)
	require.Equal(t, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(20000000000000))), liquidityPoolOwnerBalance)

	strategicReserveAccountBalance := testHelper.BankUtils.GetAccountAllBalances(strategicReserveAccountAddress)
	require.Equal(t, math.NewInt(49999990000000), strategicReserveAccountBalance.AmountOf(testenv.DefaultTestDenom))

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func addStrategicReserveVestingAccount(t *testing.T, testHelper *testapp.TestHelper) {
	err := testHelper.AuthUtils.CreateVestingAccount(v140.StrategicReserveAccount, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(79_999_990_000_000))), time.Unix(1727222400, 0), time.Unix(1821830400, 0))
	require.NoError(t, err)
}

func TestUpdateCommunityPool_Success(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	feePool := distrtypes.FeePool{
		CommunityPool: sdk.NewDecCoins(sdk.NewDecCoinFromDec(testenv.DefaultTestDenom, CommunityPoolBeforeAmount)),
	}
	testHelper.DistributionUtils.DistrKeeper.SetFeePool(testHelper.Context, feePool)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(CommunityPoolBeforeAmount.TruncateInt(), distrtypes.ModuleName)

	err := v140.UpdateCommunityPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	feePool = testHelper.DistributionUtils.DistrKeeper.GetFeePool(testHelper.Context)
	require.Equal(t, sdk.NewDecCoinsFromCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(40000000000000))), feePool.CommunityPool)

	communityPoolBalance := testHelper.DistributionUtils.DistrKeeper.GetFeePoolCommunityCoins(testHelper.Context)
	require.Equal(t, sdk.NewDecCoins(sdk.NewDecCoin(testenv.DefaultTestDenom, math.NewInt(40_000_000_000_000))), communityPoolBalance)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
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
