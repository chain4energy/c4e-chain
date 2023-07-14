package v200_test

import (
	v200 "github.com/chain4energy/c4e-chain/app/upgrades/v200"
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var (
	oldEarlyBirdRoundType = cfevestingtypes.VestingType{
		Name:          "Early-bird round",
		Free:          sdk.MustNewDecFromStr("0.10"),
		LockupPeriod:  (365 + 91) * 24 * time.Hour,
		VestingPeriod: 365 * 24 * time.Hour,
	}

	oldPublicRoundType = cfevestingtypes.VestingType{
		Name:          "Public round",
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 274 * 24 * time.Hour,
	}

	newEarlyBirdRoundVestingType = cfevestingtypes.VestingType{
		Name:          "Early-bird round",
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  0,
		VestingPeriod: 274 * 24 * time.Hour,
	}

	newPublicRoundVestingType = cfevestingtypes.VestingType{
		Name:          "Public round",
		Free:          sdk.MustNewDecFromStr("0.20"),
		LockupPeriod:  0,
		VestingPeriod: 183 * 24 * time.Hour,
	}

	fairdropVestingType = cfevestingtypes.VestingType{
		Name:          "Fairdrop",
		Free:          sdk.MustNewDecFromStr("0.01"),
		LockupPeriod:  183 * 24 * time.Hour,
		VestingPeriod: 91 * 24 * time.Hour,
	}

	moondropVestingType = cfevestingtypes.VestingType{
		Name:          "Moondrop",
		Free:          sdk.ZeroDec(),
		LockupPeriod:  730 * 24 * time.Hour,
		VestingPeriod: 730 * 24 * time.Hour,
	}
)

func TestUpdateVestingTypes(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)

	vestingTypesBefore := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.NotNil(t, vestingTypesBefore)
	require.Equal(t, 2, len(vestingTypesBefore.VestingTypes))

	err := v200.ModifyAndAddVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vestingTypesAfter := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 4, len(vestingTypesAfter.VestingTypes))
	expectedTypes := []*cfevestingtypes.VestingType{&newEarlyBirdRoundVestingType, &newPublicRoundVestingType, &fairdropVestingType, &moondropVestingType}
	require.ElementsMatch(t, expectedTypes, vestingTypesAfter.VestingTypes)
}

func TestUpdateVestingTypesVestingTypesNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	vestingTypesBefore := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.NotNil(t, vestingTypesBefore)
	require.Equal(t, 0, len(vestingTypesBefore.VestingTypes))

	err := v200.ModifyAndAddVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vestingTypesAfter := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 2, len(vestingTypesAfter.VestingTypes))
	expectedTypes := []*cfevestingtypes.VestingType{&fairdropVestingType, &moondropVestingType}
	require.ElementsMatch(t, expectedTypes, vestingTypesAfter.VestingTypes)
}

func addVestingTypes(testHelper *testapp.TestHelper) {
	vestingTypes := cfevestingtypes.VestingTypes{
		VestingTypes: []*cfevestingtypes.VestingType{&oldEarlyBirdRoundType, &oldPublicRoundType},
	}
	testHelper.App.CfevestingKeeper.SetVestingTypes(testHelper.Context, vestingTypes)
}
