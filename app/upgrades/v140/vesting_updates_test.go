package v140_test

import (
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v140 "github.com/chain4energy/c4e-chain/app/upgrades/v140"
)

const (
	VcRoundTypeName        = "VC round"
	ValidatorRoundTypeName = "Valdiator round"
)

func TestModifyVestingTypesNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	vestingTypesBefore := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.NotNil(t, vestingTypesBefore)
	require.Equal(t, 0, len(vestingTypesBefore.VestingTypes))

	err := v140.ModifyVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vestingTypesAfter := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 0, len(vestingTypesAfter.VestingTypes))
}

func TestModifyVestingTypesVestingTypeExists(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	vcRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          VcRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  548 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, vcRoundTypeBefore)

	validatorRoundTypeBefore := cfevestingtypes.VestingType{
		Name:          ValidatorRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	testHelper.App.CfevestingKeeper.SetVestingType(testHelper.Context, validatorRoundTypeBefore)

	err := v140.ModifyVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vcRoundTypeAfter, err := testHelper.C4eVestingUtils.GetC4eVestingKeeper().MustGetVestingType(testHelper.Context, VcRoundTypeName)
	require.Nil(t, err)
	expectedVcRoundType := &cfevestingtypes.VestingType{
		Name:          VcRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.08"),
		LockupPeriod:  122 * 24 * time.Hour,
		VestingPeriod: 305 * 24 * time.Hour,
	}
	require.EqualValues(t, expectedVcRoundType, vcRoundTypeAfter)

	validatorRoundTypeAfter, err := testHelper.C4eVestingUtils.GetC4eVestingKeeper().MustGetVestingType(testHelper.Context, ValidatorRoundTypeName)
	require.Nil(t, err)
	expectedValidatorRoundType := &cfevestingtypes.VestingType{
		Name:          ValidatorRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.08"),
		LockupPeriod:  122 * 24 * time.Hour,
		VestingPeriod: 305 * 24 * time.Hour,
	}
	require.EqualValues(t, expectedValidatorRoundType, validatorRoundTypeAfter)
}
