package keeper_test

import (
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"testing"
	"time"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestClaimInitial(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	start := testHelper.Context.BlockTime()

	end := testHelper.Context.BlockTime().Add(1000)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour

	campaigns := []types.Campaign{
		{
			Id:            0,
			Owner:         acountsAddresses[0].String(),
			Enabled:       true,
			StartTime:     &start,
			EndTime:       &end,
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
			Description:   "test-campaign",
		},
	}
	weight := sdk.MustNewDecFromStr("0")
	weight2 := sdk.MustNewDecFromStr("0.2")
	missions := []types.Mission{
		{
			CampaignId:  0,
			Id:          0,
			Description: "test-mission",
			Weight:      &weight,
			MissionType: types.MissionInitialClaim,
		},
		{
			CampaignId:  0,
			Id:          1,
			Description: "test-mission",
			Weight:      &weight2,
			MissionType: types.MissionVote,
		},
	}
	genesisState := types.GenesisState{Missions: missions, Campaigns: campaigns}
	testHelper.C4eAirdropUtils.InitGenesis(genesisState)

	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, prepareAidropEntries(acountsAddresses[1].String()))

	testHelper.C4eAirdropUtils.ClaimInitial(0, acountsAddresses[1], 800000000)

}

func TestClaimInitialCampaignNotFound(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	start := testHelper.Context.BlockTime()
	end := testHelper.Context.BlockTime().Add(1000)
	campaigns := []types.Campaign{
		{
			Id:            0,
			Owner:         acountsAddresses[0].String(),
			Enabled:       true,
			StartTime:     &start,
			EndTime:       &end,
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
			Description:   "test-campaign",
		},
	}
	weight := sdk.MustNewDecFromStr("0")
	weight2 := sdk.MustNewDecFromStr("0.2")
	missions := []types.Mission{
		{
			CampaignId:  0,
			Id:          0,
			Description: "test-mission",
			Weight:      &weight,
			MissionType: types.MissionInitialClaim,
		},
		{
			CampaignId:  0,
			Id:          1,
			Description: "test-mission",
			Weight:      &weight2,
			MissionType: types.MissionVote,
		},
	}
	genesisState := types.GenesisState{Missions: missions, Campaigns: campaigns}
	testHelper.C4eAirdropUtils.InitGenesis(genesisState)

	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, prepareAidropEntries(acountsAddresses[1].String()))

	testHelper.C4eAirdropUtils.ClaimInitialError(2, acountsAddresses[0], "campaign not found: campaign id: 2 : not found")

}

func TestClaimInitialCampaignClaimError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	start := testHelper.Context.BlockTime()
	end := testHelper.Context.BlockTime().Add(1000)
	campaigns := []types.Campaign{
		{
			Id:            0,
			Owner:         acountsAddresses[0].String(),
			Enabled:       true,
			StartTime:     &start,
			EndTime:       &end,
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
			Description:   "test-campaign",
		},
	}
	weight := sdk.MustNewDecFromStr("0")
	weight2 := sdk.MustNewDecFromStr("0.2")
	missions := []types.Mission{
		{
			CampaignId:  0,
			Id:          0,
			Description: "test-mission",
			Weight:      &weight,
			MissionType: types.MissionInitialClaim,
		},
		{
			CampaignId:  0,
			Id:          1,
			Description: "test-mission",
			Weight:      &weight2,
			MissionType: types.MissionVote,
		},
	}
	genesisState := types.GenesisState{Missions: missions, Campaigns: campaigns}
	testHelper.C4eAirdropUtils.InitGenesis(genesisState)

	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, prepareAidropEntries(acountsAddresses[1].String()))
	userAirdropEntries := testHelper.C4eAirdropUtils.GetUserAirdropEntries(acountsAddresses[1].String())
	userAirdropEntries.GetAirdropEntries()[0].ClaimedMissions = []uint64{3}
	testHelper.C4eAirdropUtils.SetUserAirdropEntries(userAirdropEntries)

	testHelper.C4eAirdropUtils.ClaimInitialError(0, acountsAddresses[1], fmt.Sprintf("mission already claimed: address %s, campaignId: 1, missionId: 3: mission already claimed", acountsAddresses[0]))
}

func TestClaimInitialTwoCampaigns(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	start := testHelper.Context.BlockTime()
	end := testHelper.Context.BlockTime().Add(1000)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	campaigns := []types.Campaign{
		{
			Id:            0,
			Enabled:       true,
			StartTime:     &start,
			EndTime:       &end,
			Owner:         acountsAddresses[0].String(),
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
			Description:   "test-campaign",
		},
		{
			Id:            1,
			Enabled:       true,
			StartTime:     &start,
			Owner:         acountsAddresses[0].String(),
			EndTime:       &end,
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
			Description:   "test-campaign-1",
		},
	}

	zeroWeight := sdk.MustNewDecFromStr("0")
	weight1 := sdk.MustNewDecFromStr("0.2")
	weight2 := sdk.MustNewDecFromStr("0.3")

	missions := []types.Mission{
		{
			CampaignId:  0,
			Id:          0,
			Description: "test-mission",
			Weight:      &zeroWeight,
			MissionType: types.MissionInitialClaim,
		},
		{
			CampaignId:  0,
			Id:          1,
			Description: "test-mission",
			Weight:      &weight1,
			MissionType: types.MissionVote,
		},
		{
			CampaignId:  1,
			Id:          0,
			Description: "test-mission",
			Weight:      &zeroWeight,
			MissionType: types.MissionInitialClaim,
		},
		{
			CampaignId:  1,
			Id:          1,
			Description: "test-mission",
			Weight:      &weight2,
			MissionType: types.MissionVote,
		},
	}
	genesisState := types.GenesisState{Missions: missions, Campaigns: campaigns}
	testHelper.C4eAirdropUtils.InitGenesis(genesisState)

	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, prepareAidropEntries(acountsAddresses[1].String()))
	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 1, prepareAidropEntries(acountsAddresses[1].String()))

	testHelper.C4eAirdropUtils.ClaimInitial(0, acountsAddresses[1], 800000000)
	testHelper.C4eAirdropUtils.ClaimInitial(1, acountsAddresses[1], 700000000)
}
