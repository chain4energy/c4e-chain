package v130_test

import (
	"cosmossdk.io/math"
	v130 "github.com/chain4energy/c4e-chain/app/upgrades/v130"
	"github.com/chain4energy/c4e-chain/app/upgrades/v130/claim"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	airdropStartTime      = time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)
	airdropEndTime        = time.Date(2031, 1, 1, 0, 0, 0, 0, time.UTC)
	airdropLockupPeriod   = 183 * 24 * time.Hour
	airdropVestingPeriod  = 91 * 24 * time.Hour
	moondropLockupPeriod  = 730 * 24 * time.Hour
	moondropVestingPeriod = 730 * 24 * time.Hour
)

var (
	moondropCampaign = cfeclaimtypes.Campaign{
		Id:                     0,
		Owner:                  "c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8",
		Name:                   "Moon Drop",
		Description:            "",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  true,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.ZeroDec(),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           moondropLockupPeriod,
		VestingPeriod:          moondropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(7280002000000))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(7280002000000))),
		VestingPoolName:        "Moondrop",
	}

	stakedropCampaign = cfeclaimtypes.Campaign{
		Id:                     1,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "Stake Drop",
		Description:            "Stake Drop is the airdrop aimed to spread knowledge about the C4E ecosystem among the Cosmos $ATOM stakers community. The airdrop snapshot has been taken on September 28th, 2022 at 9:30 PM UTC (during the ATOM 2.0 roadmap announcement at the Cosmoverse Conference.",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.MustNewDecFromStr("0.01"),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           airdropLockupPeriod,
		VestingPeriod:          airdropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(8999999989680))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(8999999989680))),
		VestingPoolName:        "Fairdrop",
	}

	santadropCampaign = cfeclaimtypes.Campaign{
		Id:                     2,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "Santa Drop",
		Description:            "Santa Drop prize pool for was 10.000 C4E Tokens, with 10 lucky winners getting 1000 tokens per each. The participants had to complete the tasks to get a chance to be among lucky winners.",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.MustNewDecFromStr("0.01"),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           airdropLockupPeriod,
		VestingPeriod:          airdropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(10000000000))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(10000000000))),
		VestingPoolName:        "Fairdrop",
	}

	greendropCampaign = cfeclaimtypes.Campaign{
		Id:                     3,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "Green Drop",
		Description:            "It was the first airdrop competition aimed at spreading knowledge about the C4E ecosystem. The Prize Pool was 1.000.000 C4E tokens and what is best — all the participants who completed the tasks are eligible for the c4e tokens from it!",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.MustNewDecFromStr("0.01"),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           airdropLockupPeriod,
		VestingPeriod:          airdropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(996647490000))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(996647490000))),
		VestingPoolName:        "Fairdrop",
	}

	zealydropCampaign = cfeclaimtypes.Campaign{
		Id:                     4,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "Incentived Testnet I",
		Description:            "Incentivized Testnet Zealy campaign, is innovative approach designed to foster engagement and bolster network security. Community members are rewarded for participating in testnet and marketing tasks, receiving C4E tokens as a result of their contributions.",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.MustNewDecFromStr("0.01"),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           airdropLockupPeriod,
		VestingPeriod:          airdropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(392677840916))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(392677840916))),
		VestingPoolName:        "Fairdrop",
	}

	amadropCampaign = cfeclaimtypes.Campaign{
		Id:                     5,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "AMA Drop",
		Description:            "Have you been active at our AMA sessions and won C4E prizes? This Drop belongs to you.",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.MustNewDecFromStr("0.01"),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           airdropLockupPeriod,
		VestingPeriod:          airdropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(2900000000))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(2900000000))),
		VestingPoolName:        "Fairdrop",
	}
)

func TestSetupCampaigns(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addVestingPools(testHelper)
	addMoondropVestingAccount(t, testHelper)
	_ = addAirdropModuleAccount(testHelper)

	_ = addAirdropModuleAccount(testHelper)
	campaigns := testHelper.C4eClaimUtils.GetCampaigns()
	require.Nil(t, campaigns)
	err := v130.ModifyAndAddVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = v130.MigrateMoondropVestingAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = v130.MigrateAirdropModuleAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = claim.SetupAirdrops(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	campaigns = testHelper.C4eClaimUtils.GetCampaigns()
	userEntires := testHelper.C4eClaimUtils.GetAllUsersEntries()
	require.Equal(t, 6, len(campaigns))
	require.Equal(t, 108359, len(userEntires))

	require.Equal(t, moondropCampaign, campaigns[0])
	require.Equal(t, stakedropCampaign, campaigns[1])
	require.Equal(t, santadropCampaign, campaigns[2])
	require.Equal(t, greendropCampaign, campaigns[3])
	require.Equal(t, zealydropCampaign, campaigns[4])
	require.Equal(t, amadropCampaign, campaigns[5])
}

func TestSetupCampaignsNoMoondropVestingAccount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addVestingPools(testHelper)
	addMoondropVestingAccount(t, testHelper)
	addAirdropModuleAccount(testHelper)
	campaigns := testHelper.C4eClaimUtils.GetCampaigns()
	require.Nil(t, campaigns)
	err := v130.ModifyAndAddVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = v130.MigrateAirdropModuleAccount(testHelper.Context, testHelper.App)
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
	addMoondropVestingAccount(t, testHelper)
	addAirdropModuleAccount(testHelper)

	campaigns := testHelper.App.CfeclaimKeeper.GetAllCampaigns(testHelper.Context)
	require.Nil(t, campaigns)
	err := v130.ModifyAndAddVestingTypes(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = v130.MigrateMoondropVestingAccount(testHelper.Context, testHelper.App)
	require.NoError(t, err)
	err = claim.SetupAirdrops(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	campaigns = testHelper.C4eClaimUtils.GetCampaigns()
	require.Nil(t, campaigns)
}
