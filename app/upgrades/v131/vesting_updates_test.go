package v131_test

import (
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v131 "github.com/chain4energy/c4e-chain/app/upgrades/v131"
)

func TestModifyVestingTypesVCRoundNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	validatorRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Validator round",
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, validatorRoundTypeBefore)

	publicRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Public round",
		Free:          sdk.MustNewDecFromStr("0.2"),
		LockupPeriod:  0,
		VestingPeriod: 183 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, publicRoundTypeBefore)

	earlyBirdRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Early-bird round",
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  0,
		VestingPeriod: 274 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, earlyBirdRoundTypeBefore)

	vestingTypesBefore := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.NotNil(t, vestingTypesBefore)
	require.Equal(t, 3, len(vestingTypesBefore.VestingTypes))

	err := v131.ModifyVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vestingTypesAfter := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 3, len(vestingTypesAfter.VestingTypes))

	for _, vtBefore := range vestingTypesBefore.VestingTypes {
		vtAfter, found := findVestingTypeByName(vestingTypesAfter.VestingTypes, vtBefore.Name)
		require.True(t, found)
		require.EqualValues(t, vtBefore, vtAfter)
	}

	testHelper.ValidateGenesisAndInvariants()
}

func TestModifyVestingTypesPublicRoundNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	vcRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "VC round",
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  548 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, vcRoundTypeBefore)

	validatorRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Validator round",
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, validatorRoundTypeBefore)

	earlyBirdRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Early-bird round",
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  0,
		VestingPeriod: 274 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, earlyBirdRoundTypeBefore)

	vestingTypesBefore := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.NotNil(t, vestingTypesBefore)
	require.Equal(t, 3, len(vestingTypesBefore.VestingTypes))

	err := v131.ModifyVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vestingTypesAfter := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 3, len(vestingTypesAfter.VestingTypes))

	for _, vtBefore := range vestingTypesBefore.VestingTypes {
		vtAfter, found := findVestingTypeByName(vestingTypesAfter.VestingTypes, vtBefore.Name)
		require.True(t, found)
		require.EqualValues(t, vtBefore, vtAfter)
	}

	testHelper.ValidateGenesisAndInvariants()
}

func TestModifyVestingTypesEarlyBirdRoundNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	vcRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "VC round",
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  548 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, vcRoundTypeBefore)

	validatorRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Validator round",
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, validatorRoundTypeBefore)

	publicRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Public round",
		Free:          sdk.MustNewDecFromStr("0.2"),
		LockupPeriod:  0,
		VestingPeriod: 183 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, publicRoundTypeBefore)

	vestingTypesBefore := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.NotNil(t, vestingTypesBefore)
	require.Equal(t, 3, len(vestingTypesBefore.VestingTypes))

	err := v131.ModifyVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vestingTypesAfter := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 3, len(vestingTypesAfter.VestingTypes))

	for _, vtBefore := range vestingTypesBefore.VestingTypes {
		vtAfter, found := findVestingTypeByName(vestingTypesAfter.VestingTypes, vtBefore.Name)
		require.True(t, found)
		require.EqualValues(t, vtBefore, vtAfter)
	}

	testHelper.ValidateGenesisAndInvariants()
}

func TestModifyVestingTypesValidatorRoundNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	vcRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "VC round",
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  548 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, vcRoundTypeBefore)

	publicRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Public round",
		Free:          sdk.MustNewDecFromStr("0.2"),
		LockupPeriod:  0,
		VestingPeriod: 183 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, publicRoundTypeBefore)

	earlyBirdRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Early-bird round",
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  0,
		VestingPeriod: 274 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, earlyBirdRoundTypeBefore)

	vestingTypesBefore := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.NotNil(t, vestingTypesBefore)
	require.Equal(t, 3, len(vestingTypesBefore.VestingTypes))

	err := v131.ModifyVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vestingTypesAfter := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 3, len(vestingTypesAfter.VestingTypes))

	for _, vtBefore := range vestingTypesBefore.VestingTypes {
		vtAfter, found := findVestingTypeByName(vestingTypesAfter.VestingTypes, vtBefore.Name)
		require.True(t, found)
		require.EqualValues(t, vtBefore, vtAfter)
	}

	testHelper.ValidateGenesisAndInvariants()
}

func findVestingTypeByName(vestingTypes []*cfevestingtypes.VestingType, name string) (*cfevestingtypes.VestingType, bool) {
	for _, vt := range vestingTypes {
		if vt.Name == name {
			return vt, true
		}
	}
	return &cfevestingtypes.VestingType{}, false
}

func TestModifyVestingTypesVestingTypeExists(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	vcRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "VC round",
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  548 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, vcRoundTypeBefore)

	validatorRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Validator round",
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, validatorRoundTypeBefore)

	publicRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Public round",
		Free:          sdk.MustNewDecFromStr("0.2"),
		LockupPeriod:  0,
		VestingPeriod: 183 * 24 * time.Hour,
	}

	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, publicRoundTypeBefore)

	earyBirdRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          "Early-bird round",
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  0,
		VestingPeriod: 274 * 24 * time.Hour,
	}

	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, earyBirdRoundTypeBefore)

	err := v131.ModifyVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vcRoundTypeAfter, err := testHelper.C4eVestingUtils.GetC4eVestingKeeper().MustGetVestingType(testHelper.Context, "VC round")
	require.Nil(t, err)
	expectedVcRoundType := &cfevestingtypes.VestingType{
		Name:          "VC round",
		Free:          sdk.MustNewDecFromStr("0.08"),
		LockupPeriod:  122 * 24 * time.Hour,
		VestingPeriod: 305 * 24 * time.Hour,
	}
	require.EqualValues(t, expectedVcRoundType, vcRoundTypeAfter)

	validatorRoundTypeAfter, err := testHelper.C4eVestingUtils.GetC4eVestingKeeper().MustGetVestingType(testHelper.Context, "Validator round")
	require.Nil(t, err)
	expectedValidatorRoundType := &cfevestingtypes.VestingType{
		Name:          "Validator round",
		Free:          sdk.MustNewDecFromStr("0.08"),
		LockupPeriod:  122 * 24 * time.Hour,
		VestingPeriod: 305 * 24 * time.Hour,
	}
	require.EqualValues(t, expectedValidatorRoundType, validatorRoundTypeAfter)

	publicRoundTypeFAfter, err := testHelper.C4eVestingUtils.GetC4eVestingKeeper().MustGetVestingType(testHelper.Context, "Public round")
	require.Nil(t, err)
	expectedPublicRoundType := &cfevestingtypes.VestingType{
		Name:          "Public round",
		Free:          sdk.MustNewDecFromStr("0.2"),
		LockupPeriod:  30 * 24 * time.Hour,
		VestingPeriod: 152 * 24 * time.Hour,
	}
	require.EqualValues(t, expectedPublicRoundType, publicRoundTypeFAfter)

	earlyBirdRoundTypeAfter, err := testHelper.C4eVestingUtils.GetC4eVestingKeeper().MustGetVestingType(testHelper.Context, "Early-bird round")
	require.Nil(t, err)
	expectedEarlyBirdRoundType := &cfevestingtypes.VestingType{
		Name:          "Early-bird round",
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  61 * 24 * time.Hour,
		VestingPeriod: 213 * 24 * time.Hour,
	}
	require.EqualValues(t, expectedEarlyBirdRoundType, earlyBirdRoundTypeAfter)

	testHelper.ValidateGenesisAndInvariants()
}
