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

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v131.StrategicReservcePoolOwnerAccount)
	require.False(t, found)
	require.Equal(t, 0, len(accountVestingPools.VestingPools))

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveShortTermPool_InitiallyLockedNegative(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v131.StrategicReservcePoolOwnerAccount,
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

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v131.StrategicReservcePoolOwnerAccount)
	require.True(t, found)
	require.Equal(t, 1, len(accountVestingPools.VestingPools))
	require.Equal(t, accountVestingPools.VestingPools[0].InitiallyLocked, math.NewInt(1_000))

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveShortTermPool_strategicReserveShortTermPoolNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v131.StrategicReservcePoolOwnerAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            v131.EarlyBirdRoundPoolName,
				InitiallyLocked: math.NewInt(8000000000000),
				VestingType:     "Early-bird round",
				Sent:            math.NewInt(75001000000),
			},
			{
				Name:            v131.PublicRoundPoolName,
				InitiallyLocked: math.NewInt(9000000000000),
				VestingType:     "Public round",
			},
		},
	}
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(16924999000000), cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v131.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v131.StrategicReservcePoolOwnerAccount)
	require.True(t, found)
	require.Equal(t, 2, len(accountVestingPools.VestingPools))
	require.Equal(t, accountVestingPools.VestingPools[0].InitiallyLocked, math.NewInt(8000000000000))
	require.Equal(t, accountVestingPools.VestingPools[1].InitiallyLocked, math.NewInt(9000000000000))

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveShortTermPool_earlyBirdRoundNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v131.StrategicReservcePoolOwnerAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            v131.StrategicReservceShortTermPool,
				InitiallyLocked: math.NewInt(40_000_000_000_000),
				VestingType:     "Strategic reserve short term round",
			},
			{
				Name:            v131.PublicRoundPoolName,
				InitiallyLocked: math.NewInt(9000000000000),
				VestingType:     "Public round",
			},
		},
	}
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(49000000000000), cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v131.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v131.StrategicReservcePoolOwnerAccount)
	require.True(t, found)
	require.Equal(t, 2, len(accountVestingPools.VestingPools))
	require.Equal(t, accountVestingPools.VestingPools[0].InitiallyLocked, math.NewInt(40_000_000_000_000))
	require.Equal(t, accountVestingPools.VestingPools[1].InitiallyLocked, math.NewInt(9000000000000))

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveShortTermPool_publicRoundNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v131.StrategicReservcePoolOwnerAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            v131.StrategicReservceShortTermPool,
				InitiallyLocked: math.NewInt(40_000_000_000_000),
				VestingType:     "Strategic reserve short term round",
			},
			{
				Name:            v131.EarlyBirdRoundPoolName,
				InitiallyLocked: math.NewInt(8000000000000),
				VestingType:     "Early-bird round",
				Sent:            math.NewInt(75001000000),
			},
		},
	}
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(47924999000000), cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v131.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v131.StrategicReservcePoolOwnerAccount)
	require.True(t, found)
	require.Equal(t, 2, len(accountVestingPools.VestingPools))
	require.Equal(t, accountVestingPools.VestingPools[0].InitiallyLocked, math.NewInt(40_000_000_000_000))
	require.Equal(t, accountVestingPools.VestingPools[1].InitiallyLocked, math.NewInt(8000000000000))

	testHelper.ValidateGenesisAndInvariants()
}

func TestUpdateStrategicReserveShortTermPool_BurnCoinsError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v131.StrategicReservcePoolOwnerAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            v131.StrategicReservceShortTermPool,
				InitiallyLocked: math.NewInt(40_000_000_000_000),
				VestingType:     "Strategic reserve short term round",
			},
			{
				Name:            v131.EarlyBirdRoundPoolName,
				InitiallyLocked: math.NewInt(8000000000000),
				VestingType:     "Early-bird round",
				Sent:            math.NewInt(75001000000),
			},
			{
				Name:            v131.PublicRoundPoolName,
				InitiallyLocked: math.NewInt(9000000000000),
				VestingType:     "Public round",
			},
		},
	}

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1), cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v131.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.EqualError(t, err, "spendable balance 1uc4e is smaller than 20000000000000uc4e: insufficient funds")
}

func TestUpdateStrategicReserveShortTermPool_Success(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: v131.StrategicReservcePoolOwnerAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            v131.StrategicReservceShortTermPool,
				InitiallyLocked: math.NewInt(40_000_000_000_000),
				VestingType:     "Strategic reserve short term round",
			},
			{
				Name:            v131.EarlyBirdRoundPoolName,
				InitiallyLocked: math.NewInt(8000000000000),
				VestingType:     "Early-bird round",
				Sent:            math.NewInt(75001000000),
			},
			{
				Name:            v131.PublicRoundPoolName,
				InitiallyLocked: math.NewInt(9000000000000),
				VestingType:     "Public round",
			},
		},
	}
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(56924999000000), cfevestingtypes.ModuleName)

	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetAccountVestingPools(testHelper.Context, accountVestingPools)

	err := v131.UpdateStrategicReserveShortTermPool(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v131.StrategicReservcePoolOwnerAccount)
	require.True(t, found)
	require.Equal(t, 3, len(accountVestingPools.VestingPools))
	require.Equal(t, math.NewInt(20_000_000_000_000), accountVestingPools.VestingPools[0].InitiallyLocked)
	require.Equal(t, math.NewInt(75001000000), accountVestingPools.VestingPools[1].InitiallyLocked)
	require.Equal(t, math.ZeroInt(), accountVestingPools.VestingPools[2].InitiallyLocked)

	strategicReservePoolOwnerAccountAddress, _ := sdk.AccAddressFromBech32(v131.StrategicReservcePoolOwnerAccount)
	testHelper.BankUtils.VerifyAccountDefaultDenomBalance(testHelper.Context, strategicReservePoolOwnerAccountAddress, math.NewInt(16924999000000))

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
		sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(35_000_000_000_000))),
		time.Unix(1727222400, 0), time.Unix(1821830400, 0))
	require.NoError(t, err)

	err = v131.UpdateStrategicReserveAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	strategicReserveAccountAddress, err := sdk.AccAddressFromBech32(v131.StrategicReserveAccount)
	require.NoError(t, err)

	testHelper.BankUtils.VerifyAccountDefaultDenomBalance(testHelper.Context, strategicReserveAccountAddress, math.NewInt(800000000000))
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
	require.Equal(t, math.NewInt(45799990000000), strategicReserveAccount.OriginalVesting.AmountOf(testenv.DefaultTestDenom))

	liquidityPoolOwnerBalance := testHelper.BankUtils.GetAccountAllBalances(liquidityPoolOwnerAccountAddress)
	require.Equal(t, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(20000000000000))), liquidityPoolOwnerBalance)

	strategicReserveAccountBalance := testHelper.BankUtils.GetAccountAllBalances(strategicReserveAccountAddress)
	require.Equal(t, math.NewInt(45799990000000), strategicReserveAccountBalance.AmountOf(testenv.DefaultTestDenom))

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
	newFeePoolAmount := sdk.NewDecCoins(sdk.NewDecCoinFromDec(testenv.DefaultTestDenom, sdk.MustNewDecFromStr("40000000083683.601351991745061083")))
	require.Equal(t, newFeePoolAmount, feePool.CommunityPool)

	communityPoolBalance := testHelper.DistributionUtils.DistrKeeper.GetFeePoolCommunityCoins(testHelper.Context)
	require.Equal(t, newFeePoolAmount, communityPoolBalance)

	testHelper.ValidateGenesisAndInvariants()
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
		VestingTypes: []*cfevestingtypes.VestingType{
			{
				Name:          "Strategic reserve short term round",
				Free:          sdk.MustNewDecFromStr("0.10"),
				LockupPeriod:  222 * 24 * time.Hour,
				VestingPeriod: 365 * 24 * time.Hour,
			},
			{
				Name:          "Early-bird round",
				Free:          sdk.MustNewDecFromStr("0.15"),
				LockupPeriod:  0,
				VestingPeriod: 274 * 24 * time.Hour,
			},
			{
				Name:          "VC round",
				Free:          sdk.MustNewDecFromStr("0.05"),
				LockupPeriod:  548 * 24 * time.Hour,
				VestingPeriod: 548 * 24 * time.Hour,
			},
			{
				Name:          "Public round",
				Free:          sdk.MustNewDecFromStr("0.2"),
				LockupPeriod:  0,
				VestingPeriod: 183 * 24 * time.Hour,
			},
		},
	}
	testHelper.C4eVestingUtils.GetC4eVestingKeeper().SetVestingTypes(testHelper.Context, vestingTypes)
}
