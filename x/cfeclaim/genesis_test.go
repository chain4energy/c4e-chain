package cfeclaim_test

import (
	"cosmossdk.io/math"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var startTime = time.Now().Add(time.Hour)
var validGenesisState = types.GenesisState{
	Campaigns: []types.Campaign{
		{
			Id:                     0,
			Owner:                  "c4e15ky9du8a2wlstz6fpx3p4mqpjyrm5cgq3kx2f7",
			Name:                   "Campaign 1",
			Description:            "Campaign 1 description",
			CampaignType:           2,
			FeegrantAmount:         math.NewInt(300),
			InitialClaimFreeAmount: math.NewInt(500),
			Free:                   sdk.ZeroDec(),
			Enabled:                false,
			StartTime:              time.Now(),
			EndTime:                time.Now().Add(time.Hour),
			LockupPeriod:           time.Hour * 10000,
			VestingPeriod:          time.Hour * 10000,
			VestingPoolName:        "Vesting1",
		},
		{
			Id:                     1,
			Owner:                  "c4e15ky9du8a2wlstz6fpx3p4mqpjyrm5cgpvqjl5v",
			Name:                   "Campaign 2",
			Description:            "Campaign 2 description",
			CampaignType:           1,
			FeegrantAmount:         math.NewInt(100),
			InitialClaimFreeAmount: math.NewInt(300),
			Enabled:                true,
			Free:                   sdk.ZeroDec(),
			StartTime:              time.Now(),
			EndTime:                time.Now().Add(time.Hour),
			LockupPeriod:           1234,
			VestingPeriod:          1000,
		},
	},
	Missions: []types.Mission{
		{
			Id:             0,
			CampaignId:     0,
			Name:           "Mission 1",
			Description:    "Mission 1 description",
			MissionType:    1,
			Weight:         sdk.MustNewDecFromStr("0"),
			ClaimStartDate: &startTime,
		},
		{
			Id:             1,
			CampaignId:     0,
			Name:           "Mission 1",
			Description:    "Mission 1 description",
			MissionType:    2,
			Weight:         sdk.MustNewDecFromStr("0.5"),
			ClaimStartDate: &startTime,
		},
		{
			Id:             0,
			CampaignId:     1,
			Name:           "Mission 2",
			Description:    "Mission 2 description",
			MissionType:    1,
			Weight:         sdk.MustNewDecFromStr("0"),
			ClaimStartDate: nil,
		},
		{
			Id:             1,
			CampaignId:     1,
			Name:           "Mission 2",
			Description:    "Mission 2 description",
			MissionType:    2,
			Weight:         sdk.MustNewDecFromStr("0.2"),
			ClaimStartDate: nil,
		},
		{
			Id:             2,
			CampaignId:     1,
			Name:           "Mission 2",
			Description:    "Mission 2 description",
			MissionType:    3,
			Weight:         sdk.MustNewDecFromStr("0.4"),
			ClaimStartDate: nil,
		},
	},
	UsersEntries: []types.UserEntry{
		{
			Address: "c4e1asgp8qrlznsjs7ww5f60lf64gx04s6nsrte4dv",
			ClaimRecords: []*types.ClaimRecord{
				{
					CampaignId:        0,
					Address:           "c4e1asgp8qrlznsjs7ww5f60lf64gx04s6nsrte4dv",
					Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(1234))),
					CompletedMissions: []uint64{0, 1},
					ClaimedMissions:   []uint64{0},
				},
				{
					CampaignId:        1,
					Address:           "c4e1asgp8qrlznsjs7ww5f60lf64gx04s6nsrte4dv",
					Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000))),
					CompletedMissions: []uint64{0},
					ClaimedMissions:   []uint64{0},
				},
			},
		},
		{
			Address: "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
			ClaimRecords: []*types.ClaimRecord{
				{
					CampaignId:        0,
					Address:           "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
					Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(1234))),
					CompletedMissions: []uint64{0, 1},
					ClaimedMissions:   []uint64{0},
				},
				{
					CampaignId:        1,
					Address:           "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
					Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000))),
					CompletedMissions: []uint64{0, 1},
					ClaimedMissions:   []uint64{0, 1},
				},
			},
		},
	},
}

func TestValidGenesis(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	accAddress, err := sdk.AccAddressFromBech32(validGenesisState.Campaigns[0].Owner)
	require.NoError(t, err)
	testHelper.C4eVestingUtils.AddTestVestingPool(accAddress, "Vesting1", math.NewInt(10000), 100, 100)

	testHelper.C4eClaimUtils.InitGenesis(validGenesisState)
	testHelper.C4eClaimUtils.ExportGenesis(validGenesisState)
}

func TestInvalidCampaignInGenesis(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	accAddress, err := sdk.AccAddressFromBech32(validGenesisState.Campaigns[0].Owner)
	require.NoError(t, err)
	testHelper.C4eVestingUtils.AddTestVestingPool(accAddress, "Vesting1", math.NewInt(10000), 100, 100)
	invalidGenesis := proto.Clone(&validGenesisState).(*types.GenesisState)
	invalidGenesis.Campaigns[0].VestingPoolName = ""
	testHelper.C4eClaimUtils.InitGenesisError(*invalidGenesis, "for VESTING_POOL type campaigns, the vesting pool name must be provided: wrong param value")
}

func TestInvalidMissionWeightInGenesis(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	accAddress, err := sdk.AccAddressFromBech32(validGenesisState.Campaigns[0].Owner)
	require.NoError(t, err)
	testHelper.C4eVestingUtils.AddTestVestingPool(accAddress, "Vesting1", math.NewInt(10000), 100, 100)
	invalidGenesis := proto.Clone(&validGenesisState).(*types.GenesisState)
	invalidGenesis.Missions[4].Weight = sdk.MustNewDecFromStr("0.9")
	testHelper.C4eClaimUtils.InitGenesisError(*invalidGenesis, "all campaign missions weight sum is >= 1 (1.100000000000000000 > 1) error: wrong param value")
}

func TestInvalidMissionInitialClaimInGenesis(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	accAddress, err := sdk.AccAddressFromBech32(validGenesisState.Campaigns[0].Owner)
	require.NoError(t, err)
	testHelper.C4eVestingUtils.AddTestVestingPool(accAddress, "Vesting1", math.NewInt(10000), 100, 100)
	invalidGenesis := proto.Clone(&validGenesisState).(*types.GenesisState)
	invalidGenesis.Missions[3].MissionType = types.MissionInitialClaim
	testHelper.C4eClaimUtils.InitGenesisError(*invalidGenesis, "there can be only one mission with InitialClaim type and must be first in the campaign: wrong param value")
}

func TestInvalidMissionCampaignDoestExistInGenesis(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	accAddress, err := sdk.AccAddressFromBech32(validGenesisState.Campaigns[0].Owner)
	require.NoError(t, err)
	testHelper.C4eVestingUtils.AddTestVestingPool(accAddress, "Vesting1", math.NewInt(10000), 100, 100)
	invalidGenesis := proto.Clone(&validGenesisState).(*types.GenesisState)
	invalidGenesis.Missions[3].CampaignId = 100
	testHelper.C4eClaimUtils.InitGenesisError(*invalidGenesis, "mission Mission 2: campaign with id 100 not found: not found")
}

func TestInvalidUserEntryInGenesis(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	accAddress, err := sdk.AccAddressFromBech32(validGenesisState.Campaigns[0].Owner)
	require.NoError(t, err)
	testHelper.C4eVestingUtils.AddTestVestingPool(accAddress, "Vesting1", math.NewInt(10000), 100, 100)
	invalidGenesis := proto.Clone(&validGenesisState).(*types.GenesisState)
	invalidGenesis.UsersEntries[0].Address = ""
	testHelper.C4eClaimUtils.InitGenesisError(*invalidGenesis, "userEntry index: 0: empty address string is not allowed")
}

func TestInvalidUserEntryFrongClaimedMissionInGenesis(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	accAddress, err := sdk.AccAddressFromBech32(validGenesisState.Campaigns[0].Owner)
	require.NoError(t, err)
	testHelper.C4eVestingUtils.AddTestVestingPool(accAddress, "Vesting1", math.NewInt(10000), 100, 100)
	invalidGenesis := proto.Clone(&validGenesisState).(*types.GenesisState)
	invalidGenesis.UsersEntries[0].ClaimRecords[0].ClaimedMissions = []uint64{10}
	testHelper.C4eClaimUtils.InitGenesisError(*invalidGenesis, "userEntry index: 0, claimRecord index: 0, claimed mission index: 0: mission not found - campaignId 0, missionId 10: not found")
}

func TestInvalidUserEntryFrongCompletedMissionInGenesis(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	accAddress, err := sdk.AccAddressFromBech32(validGenesisState.Campaigns[0].Owner)
	require.NoError(t, err)
	testHelper.C4eVestingUtils.AddTestVestingPool(accAddress, "Vesting1", math.NewInt(10000), 100, 100)
	invalidGenesis := proto.Clone(&validGenesisState).(*types.GenesisState)
	invalidGenesis.UsersEntries[0].ClaimRecords[0].CompletedMissions = []uint64{10}
	testHelper.C4eClaimUtils.InitGenesisError(*invalidGenesis, "userEntry index: 0, claimRecord index: 0, completed mission index: 0: mission not found - campaignId 0, missionId 10: not found")
}
