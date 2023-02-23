package v120_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	v120 "github.com/chain4energy/c4e-chain/app/upgrades/v120"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var (
	advisorsType = cfevestingtypes.VestingType{
		Name:          "Advisors",
		Free:          sdk.ZeroDec(),
		LockupPeriod:  365 * 24 * time.Hour,
		VestingPeriod: 730 * 24 * time.Hour,
	}

	oldValidatorsType = cfevestingtypes.VestingType{
		Name:          "Validators",
		Free:          sdk.ZeroDec(),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}

	newValidatorRoundType = cfevestingtypes.VestingType{
		Name:          "Validator round",
		Free:          sdk.ZeroDec(),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}

	newVcRoundType = cfevestingtypes.VestingType{
		Name:          "VC round",
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  548 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}

	newEarlyBirdRoundType = cfevestingtypes.VestingType{
		Name:          "Early-bird round",
		Free:          sdk.MustNewDecFromStr("0.10"),
		LockupPeriod:  (365 + 91) * 24 * time.Hour,
		VestingPeriod: 365 * 24 * time.Hour,
	}

	newPublicRoundType = cfevestingtypes.VestingType{
		Name:          "Public round",
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 274 * 24 * time.Hour,
	}

	newStrategicReserveShortTermRoundType = cfevestingtypes.VestingType{
		Name:          "Strategic reserve short term round",
		Free:          sdk.MustNewDecFromStr("0.20"),
		LockupPeriod:  365 * 24 * time.Hour,
		VestingPeriod: 365 * 24 * time.Hour,
	}
)

var (
	validatorsLockStart, _ = time.Parse("2006-01-02T15:04:05.000Z", "2022-09-22T14:00:00.000Z")
	validatorsLockEnd, _   = time.Parse("2006-01-02T15:04:05.000Z", "2024-12-26T00:00:00.000Z")

	advisorsLockStart, _ = time.Parse("2006-01-02T15:04:05.000Z", "2022-09-22T14:00:00.000Z")
	advisorsLockEnd, _   = time.Parse("2006-01-02T15:04:05.000Z", "2025-09-25T00:00:00.000Z")

	advisorsPool = cfevestingtypes.VestingPool{
		Name:            "Advisors pool",
		VestingType:     advisorsType.Name,
		InitiallyLocked: math.NewInt(12087500000000),
		LockStart:       advisorsLockStart,
		LockEnd:         advisorsLockEnd,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.NewInt(500000000000),
	}

	oldValidatorsPool = cfevestingtypes.VestingPool{
		Name:            "Validators pool",
		VestingType:     oldValidatorsType.Name,
		InitiallyLocked: math.NewInt(80498690000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockEnd,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.NewInt(95000000000),
	}

	newValidatorsRoundPool = cfevestingtypes.VestingPool{
		Name:            "Validator round pool",
		VestingType:     newValidatorRoundType.Name,
		InitiallyLocked: math.NewInt(8498690000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockEnd,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.NewInt(95000000000),
	}

	newVcRoundPool = cfevestingtypes.VestingPool{
		Name:            "VC round pool",
		VestingType:     newVcRoundType.Name,
		InitiallyLocked: math.NewInt(15000000000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(3, 0, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
	}

	newEarlyBirdRoundPool = cfevestingtypes.VestingPool{
		Name:            "Early-bird round pool",
		VestingType:     newEarlyBirdRoundType.Name,
		InitiallyLocked: math.NewInt(8000000000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(2, 3, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
	}

	newPublicRoundPool = cfevestingtypes.VestingPool{
		Name:            "Public round pool",
		VestingType:     newPublicRoundType.Name,
		InitiallyLocked: math.NewInt(9000000000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(1, 6, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
	}

	newStrategicRoundPool = cfevestingtypes.VestingPool{
		Name:            "Strategic reserve short term round pool",
		VestingType:     newStrategicReserveShortTermRoundType.Name,
		InitiallyLocked: math.NewInt(40000000000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(2, 0, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
	}
)

func TestSplitVestingPools(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addVestingPools(testHelper)

	avps, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	sumBefore := math.ZeroInt()
	for _, vp := range avps.VestingPools {
		sumBefore = sumBefore.Add(vp.GetCurrentlyLocked())
	}

	err := v120.ModifyVestingPoolsState(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vts := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 6, len(vts.VestingTypes))
	expectedTypes := []*cfevestingtypes.VestingType{&advisorsType, &newValidatorRoundType, &newVcRoundType,
		&newEarlyBirdRoundType, &newPublicRoundType, &newStrategicReserveShortTermRoundType}
	require.ElementsMatch(t, expectedTypes, vts.VestingTypes)

	avps, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	sumAfter := math.ZeroInt()
	for _, vp := range avps.VestingPools {
		sumAfter = sumAfter.Add(vp.GetCurrentlyLocked())
	}
	require.Equal(t, sumBefore, sumAfter)

	require.True(t, found)
	require.Equal(t, v120.ValidatorsVestingPoolOwner, avps.Owner)
	require.Equal(t, 6, len(avps.VestingPools))

	expectedPools := []*cfevestingtypes.VestingPool{&advisorsPool, &newValidatorsRoundPool, &newVcRoundPool,
		&newEarlyBirdRoundPool, &newPublicRoundPool, &newStrategicRoundPool}
	require.ElementsMatch(t, expectedPools, avps.VestingPools)

}

func TestSplitVestingPoolsNoVestingPools(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)

	_, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.False(t, found)

	err := v120.ModifyVestingPoolsState(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vts := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 2, len(vts.VestingTypes))
	expectedTypes := []*cfevestingtypes.VestingType{&advisorsType, &oldValidatorsType}
	require.ElementsMatch(t, expectedTypes, vts.VestingTypes)

	_, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.False(t, found)

}

func TestSplitVestingPoolsNoValidatorsVestingPool(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addAdvisorsVestingPools(testHelper)

	avps, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, 1, len(avps.VestingPools))

	err := v120.ModifyVestingPoolsState(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vts := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 2, len(vts.VestingTypes))
	expectedTypes := []*cfevestingtypes.VestingType{&advisorsType, &oldValidatorsType}
	require.ElementsMatch(t, expectedTypes, vts.VestingTypes)

	avps, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, v120.ValidatorsVestingPoolOwner, avps.Owner)
	require.Equal(t, 1, len(avps.VestingPools))

	expectedPools := []*cfevestingtypes.VestingPool{&advisorsPool}
	require.ElementsMatch(t, expectedPools, avps.VestingPools)

}

func TestSplitVestingPoolsNoEnoughValidatorsVestingPoolToSplit(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	vestinngPoolNotEnough := addVestingPoolsNotEnoughCoins(testHelper)

	avps, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, 2, len(avps.VestingPools))

	err := v120.ModifyVestingPoolsState(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vts := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 2, len(vts.VestingTypes))
	expectedTypes := []*cfevestingtypes.VestingType{&advisorsType, &oldValidatorsType}
	require.ElementsMatch(t, expectedTypes, vts.VestingTypes)

	avps, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, v120.ValidatorsVestingPoolOwner, avps.Owner)
	require.Equal(t, 2, len(avps.VestingPools))

	expectedPools := []*cfevestingtypes.VestingPool{&advisorsPool, &vestinngPoolNotEnough}
	require.ElementsMatch(t, expectedPools, avps.VestingPools)

}

func TestSplitVestingPoolsNoVestingType(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addAdvisorsVestingTypes(testHelper)
	addVestingPools(testHelper)

	avps, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, 2, len(avps.VestingPools))

	err := v120.ModifyVestingPoolsState(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vts := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 1, len(vts.VestingTypes))
	expectedTypes := []*cfevestingtypes.VestingType{&advisorsType}
	require.ElementsMatch(t, expectedTypes, vts.VestingTypes)

	avps, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, v120.ValidatorsVestingPoolOwner, avps.Owner)
	require.Equal(t, 2, len(avps.VestingPools))

	expectedPools := []*cfevestingtypes.VestingPool{&advisorsPool, &oldValidatorsPool}
	require.ElementsMatch(t, expectedPools, avps.VestingPools)

}

func addAdvisorsVestingTypes(testHelper *testapp.TestHelper) {
	vestingTypes := cfevestingtypes.VestingTypes{
		VestingTypes: []*cfevestingtypes.VestingType{&advisorsType},
	}
	testHelper.App.CfevestingKeeper.SetVestingTypes(testHelper.Context, vestingTypes)
}

func addVestingTypes(testHelper *testapp.TestHelper) {
	vestingTypes := cfevestingtypes.VestingTypes{
		VestingTypes: []*cfevestingtypes.VestingType{&advisorsType, &oldValidatorsType},
	}
	testHelper.App.CfevestingKeeper.SetVestingTypes(testHelper.Context, vestingTypes)
}

func addVestingPools(testHelper *testapp.TestHelper) {
	vpools := cfevestingtypes.AccountVestingPools{
		Owner:        v120.ValidatorsVestingPoolOwner,
		VestingPools: []*cfevestingtypes.VestingPool{&advisorsPool, &oldValidatorsPool},
	}
	testHelper.App.CfevestingKeeper.SetAccountVestingPools(testHelper.Context, vpools)
}

func addVestingPoolsNotEnoughCoins(testHelper *testapp.TestHelper) cfevestingtypes.VestingPool {
	oldValidatorsPoolNotEnough := oldValidatorsPool
	oldValidatorsPoolNotEnough.InitiallyLocked = oldValidatorsPool.Sent.Add(newVcRoundPool.InitiallyLocked).
		Add(newEarlyBirdRoundPool.InitiallyLocked).Add(newPublicRoundPool.InitiallyLocked).
		Add(newStrategicRoundPool.InitiallyLocked).SubRaw(1)
	vpools := cfevestingtypes.AccountVestingPools{
		Owner:        v120.ValidatorsVestingPoolOwner,
		VestingPools: []*cfevestingtypes.VestingPool{&advisorsPool, &oldValidatorsPoolNotEnough},
	}
	testHelper.App.CfevestingKeeper.SetAccountVestingPools(testHelper.Context, vpools)
	return oldValidatorsPoolNotEnough
}

func addAdvisorsVestingPools(testHelper *testapp.TestHelper) {
	vpools := cfevestingtypes.AccountVestingPools{
		Owner:        v120.ValidatorsVestingPoolOwner,
		VestingPools: []*cfevestingtypes.VestingPool{&advisorsPool},
	}
	testHelper.App.CfevestingKeeper.SetAccountVestingPools(testHelper.Context, vpools)
}
