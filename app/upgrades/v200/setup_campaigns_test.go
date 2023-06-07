package v200_test

import (
	"cosmossdk.io/math"
	v200 "github.com/chain4energy/c4e-chain/app/upgrades/v200"
	"github.com/chain4energy/c4e-chain/app/upgrades/v200/claim"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	teamdropCampaign = cfeclaimtypes.Campaign{
		Id:                     0,
		Owner:                  "c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8",
		Name:                   "teamdrop",
		Description:            "teamdrop",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  true,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.NewInt(1000000),
		Free:                   sdk.ZeroDec(),
		Enabled:                false,
		StartTime:              time.Time{},
		EndTime:                time.Time{},
		LockupPeriod:           730 * 24 * time.Hour,
		VestingPeriod:          730 * 24 * time.Hour,
		VestingPoolName:        "Teamdrop",
	}

	stakedropCampaign = cfeclaimtypes.Campaign{
		Id:                     1,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "stakedrop",
		Description:            "stakedrop",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.NewInt(1000000),
		Free:                   sdk.ZeroDec(),
		Enabled:                false,
		StartTime:              time.Time{},
		EndTime:                time.Time{},
		LockupPeriod:           183 * 24 * time.Hour,
		VestingPeriod:          91 * 24 * time.Hour,
		VestingPoolName:        "Fairdrop",
	}

	santadropCampaign = cfeclaimtypes.Campaign{
		Id:                     2,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "santadrop",
		Description:            "santadrop",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.NewInt(1000000),
		Free:                   sdk.ZeroDec(),
		Enabled:                false,
		StartTime:              time.Time{},
		EndTime:                time.Time{},
		LockupPeriod:           183 * 24 * time.Hour,
		VestingPeriod:          91 * 24 * time.Hour,
		VestingPoolName:        "Fairdrop",
	}

	gleamdropCampaign = cfeclaimtypes.Campaign{
		Id:                     3,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "gleamdrop",
		Description:            "gleamdrop",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.NewInt(1000000),
		Free:                   sdk.ZeroDec(),
		Enabled:                false,
		StartTime:              time.Time{},
		EndTime:                time.Time{},
		LockupPeriod:           183 * 24 * time.Hour,
		VestingPeriod:          91 * 24 * time.Hour,
		VestingPoolName:        "Fairdrop",
	}
)

func TestSetupCampaigns(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addVestingPools(testHelper)
	addTeamdropVestingAccount(testHelper)
	_ = addAirdropModuleAccount(testHelper)

	_ = addAirdropModuleAccount(testHelper)
	campaigns := testHelper.C4eClaimUtils.GetCampaigns()
	require.Nil(t, campaigns)
	err := v200.ModifyAndAddVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = v200.MigrateTeamdropVestingAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = v200.MigrateAirdropModuleAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = claim.SetupAirdrops(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	campaigns = testHelper.C4eClaimUtils.GetCampaigns()
	userEntires := testHelper.C4eClaimUtils.GetAllUsersEntries()
	require.Equal(t, 4, len(campaigns))
	require.Equal(t, 107404, len(userEntires))

	// TODO: add elements match when all parameters of the campaign will be defined
}

func TestSetupCampaignsNoTeamdropVestingAccount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addVestingPools(testHelper)
	addTeamdropVestingAccount(testHelper)
	addAirdropModuleAccount(testHelper)
	campaigns := testHelper.C4eClaimUtils.GetCampaigns()
	require.Nil(t, campaigns)
	err := v200.ModifyAndAddVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = v200.MigrateAirdropModuleAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = claim.SetupAirdrops(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	campaigns = testHelper.C4eClaimUtils.GetCampaigns()
	require.Nil(t, campaigns)
}

func TestSetupCampaignsNoAirdropModuleAccount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addVestingPools(testHelper)
	addTeamdropVestingAccount(testHelper)
	addAirdropModuleAccount(testHelper)

	campaigns := testHelper.App.CfeclaimKeeper.GetAllCampaigns(testHelper.Context)
	require.Nil(t, campaigns)
	err := v200.ModifyAndAddVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = v200.MigrateTeamdropVestingAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = claim.SetupAirdrops(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	campaigns = testHelper.C4eClaimUtils.GetCampaigns()
	require.Nil(t, campaigns)
}
