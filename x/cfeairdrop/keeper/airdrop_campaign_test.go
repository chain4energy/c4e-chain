package keeper_test

import (
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"
)

func TestCreateCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
}

func TestCreateCampaignEmptyName(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.Name = ""
	testHelper.C4eAirdropUtils.CreateAirdropCampaignError(acountsAddresses[0].String(), campaign, "create airdrop campaign - empty campaign name error: wrong param value")
}

func TestCreateCampaignEmptyDescription(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.Description = ""
	testHelper.C4eAirdropUtils.CreateAirdropCampaignError(acountsAddresses[0].String(), campaign, "create airdrop campaign - empty campaign description error: wrong param value")
}

func TestCreateCampaignStartTimeAfterEndTime(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	startTimeAfterEndTime := campaign.EndTime.Add(time.Hour)
	campaign.StartTime = startTimeAfterEndTime
	testHelper.C4eAirdropUtils.CreateAirdropCampaignError(acountsAddresses[0].String(), campaign, fmt.Sprintf("create airdrop campaign - start time is after end time error (%s > %s): wrong param value", campaign.StartTime, campaign.EndTime))
}

func TestCreateCampaignStartTimeInThePast(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	startTimeInThePast := campaign.StartTime.Add(-time.Hour)
	campaign.StartTime = startTimeInThePast
	testHelper.C4eAirdropUtils.CreateAirdropCampaignError(acountsAddresses[0].String(), campaign, fmt.Sprintf("create airdrop campaign - start time in the past error (%s < %s): wrong param value", campaign.StartTime, testHelper.Context.BlockTime()))
}

func TestCreateCampaignNegativeInitialClaimAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = sdk.NewInt(-100)
	testHelper.C4eAirdropUtils.CreateAirdropCampaignError(acountsAddresses[0].String(), campaign, fmt.Sprintf("create airdrop campaign - initial claim free amount (%s) cannot be negative: wrong param value", campaign.InitialClaimFreeAmount))
}

func TestCreateCampaignNegativeFeegrantAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = sdk.NewInt(-100)
	testHelper.C4eAirdropUtils.CreateAirdropCampaignError(acountsAddresses[0].String(), campaign, fmt.Sprintf("create airdrop campaign - feegrant amount (%s) cannot be negative: wrong param value", campaign.FeegrantAmount))
}

func prepareTestCampaigns(ctx sdk.Context) []types.Campaign {
	start := ctx.BlockTime()
	end := ctx.BlockTime().Add(time.Second * 10)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	return []types.Campaign{
		{
			Id:            0,
			Name:          "Name",
			Description:   "test-campaign",
			Enabled:       true,
			StartTime:     start,
			EndTime:       end,
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
		},
	}
}

func prepareTestCampaign(ctx sdk.Context) types.Campaign {
	start := ctx.BlockTime()
	end := ctx.BlockTime().Add(time.Second * 10)
	lockupPeriod := time.Hour
	vestingPeriod := 3 * time.Hour
	return types.Campaign{
		Id:            0,
		Name:          "Name",
		Description:   "test-campaign",
		Enabled:       true,
		StartTime:     start,
		EndTime:       end,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
	}
}
