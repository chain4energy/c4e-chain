package v200_test

import (
	"cosmossdk.io/math"
	"fmt"
	v200 "github.com/chain4energy/c4e-chain/app/upgrades/v200"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	teamdropPool = cfevestingtypes.VestingPool{
		Name:            "Teamdrop",
		VestingType:     "Teamdrop",
		LockStart:       time.Date(2024, 9, 26, 2, 00, 00, 00, time.UTC),
		LockEnd:         time.Date(2026, 9, 25, 2, 00, 00, 00, time.UTC),
		InitiallyLocked: math.NewInt(8899990000000),
		Withdrawn:       sdk.ZeroInt(),
		Sent:            sdk.ZeroInt(),
		GenesisPool:     true,
		Reservations:    nil,
	}

	fairdropPool = cfevestingtypes.VestingPool{
		Name:            "Fairdrop",
		VestingType:     "Fairdrop",
		LockStart:       time.Date(2023, 6, 1, 23, 59, 59, 0, time.UTC),
		LockEnd:         time.Date(2026, 6, 1, 23, 59, 59, 0, time.UTC),
		InitiallyLocked: math.NewInt(20000000000000),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
		Reservations:    nil,
	}

	validatorsLockStart, _ = time.Parse("2006-01-02T15:04:05.000Z", "2022-09-22T14:00:00.000Z")

	newEarlyBirdRoundPool = cfevestingtypes.VestingPool{
		Name:            "Early-bird round pool",
		VestingType:     newEarlyBirdRoundVestingType.Name,
		InitiallyLocked: math.NewInt(8000000000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(2, 3, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
	}

	newPublicRoundPool = cfevestingtypes.VestingPool{
		Name:            "Public round pool",
		VestingType:     newEarlyBirdRoundVestingType.Name,
		InitiallyLocked: math.NewInt(9000000000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(1, 6, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
	}
)

// TODO w testach tez uzywac skopiowanych typow i metod keepera
func TestMigrateAirdropModuleAccount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000) 
	addVestingTypes(testHelper)
	addVestingPools(testHelper)
	airdropModuleAccAddress := addAirdropModuleAccount(testHelper)

	_, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v200.AirdropModuleAccountAddress)
	require.False(t, found)
	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v200.NewAirdropVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, 2, len(accountVestingPools.VestingPools))
	err := v200.MigrateAirdropModuleAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v200.NewAirdropVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, 3, len(accountVestingPools.VestingPools))
	expectedTypes := []*cfevestingtypes.VestingPool{&fairdropPool, &newEarlyBirdRoundPool, &newPublicRoundPool}
	require.ElementsMatch(t, expectedTypes, accountVestingPools.VestingPools)
	airdropModuleBalance := testHelper.BankUtils.GetAccountAllBalances(airdropModuleAccAddress)
	require.Equal(t, airdropModuleBalance, sdk.NewCoins())
}

func TestMigrateAirdropModuleAccountDoesnTExist(t *testing.T) { // TODO literowka DoesnT - zrobic male t
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addVestingPools(testHelper)

	_, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v200.AirdropModuleAccountAddress)
	require.False(t, found)
	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v200.NewAirdropVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, 2, len(accountVestingPools.VestingPools))
	err := v200.MigrateAirdropModuleAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v200.NewAirdropVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, 2, len(accountVestingPools.VestingPools))
}

func TestMigrateTeamdropAccount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addTeamdropVestingAccount(testHelper)
	accountVestingPools, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v200.TeamdropVestingAccount)
	require.False(t, found)
	err := v200.MigrateTeamdropVestingAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	accountVestingPools, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v200.TeamdropVestingAccount)
	require.True(t, found)
	require.Equal(t, 1, len(accountVestingPools.VestingPools))
	expectedTypes := []*cfevestingtypes.VestingPool{&teamdropPool}
	require.ElementsMatch(t, expectedTypes, accountVestingPools.VestingPools)
	teamdropAccAddress, _ := sdk.AccAddressFromBech32(v200.TeamdropVestingAccount)
	airdropModuleBalance := testHelper.BankUtils.GetAccountAllBalances(teamdropAccAddress)
	require.Equal(t, airdropModuleBalance, sdk.NewCoins())
}

func TestMigrateTeamdropAccountAccountNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	_, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v200.TeamdropVestingAccount)
	require.False(t, found)
	err := v200.MigrateTeamdropVestingAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	_, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v200.TeamdropVestingAccount)
	require.False(t, found)
}

func addTeamdropVestingAccount(testHelper *testapp.TestHelper) {
	err := testHelper.AuthUtils.CreateVestingAccount(v200.TeamdropVestingAccount, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(8899990000000))), testHelper.Context.BlockTime(), testHelper.Context.BlockTime().Add(time.Hour))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func addAirdropModuleAccount(testHelper *testapp.TestHelper) sdk.AccAddress {
	pubkey := secp256k1.GenPrivKey().PubKey()
	addr := sdk.AccAddress(pubkey.Address())
	baseFairdropAccount := authtypes.NewBaseAccount(addr, pubkey, 0, 0)
	fairdropAccount := authtypes.NewModuleAccount(baseFairdropAccount, "fairdrop")
	testHelper.App.AccountKeeper.SetAccount(testHelper.Context, fairdropAccount)
	airdropModuleAccAddress, _ := sdk.AccAddressFromBech32(v200.AirdropModuleAccountAddress)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(math.NewInt(20000000000000), airdropModuleAccAddress)
	return airdropModuleAccAddress
}

func addVestingPools(testHelper *testapp.TestHelper) {
	vpools := cfevestingtypes.AccountVestingPools{
		Owner:        v200.NewAirdropVestingPoolOwner,
		VestingPools: []*cfevestingtypes.VestingPool{&newEarlyBirdRoundPool, &newPublicRoundPool},
	}
	testHelper.App.CfevestingKeeper.SetAccountVestingPools(testHelper.Context, vpools)
}
