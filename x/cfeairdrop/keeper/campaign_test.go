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
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
}

func TestCreateManyAirdropCampaigns(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaigns := prepareNTestCampaigns(testHelper.Context, 10)
	for _, campaign := range campaigns {
		testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	}
}

func TestCreateCampaignEmptyName(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.Name = ""
	testHelper.C4eAirdropUtils.CreateCampaignError(acountsAddresses[0].String(), campaign, "create airdrop campaign - empty campaign name error: wrong param value")
}

func TestCreateCampaignEmptyDescription(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.Description = ""
	testHelper.C4eAirdropUtils.CreateCampaignError(acountsAddresses[0].String(), campaign, "create airdrop campaign - empty campaign description error: wrong param value")
}

func TestCreateCampaignStartTimeAfterEndTime(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	startTimeAfterEndTime := campaign.EndTime.Add(time.Hour)
	campaign.StartTime = startTimeAfterEndTime
	testHelper.C4eAirdropUtils.CreateCampaignError(acountsAddresses[0].String(), campaign, fmt.Sprintf("create airdrop campaign - start time is after end time error (%s > %s): wrong param value", campaign.StartTime, campaign.EndTime))
}

func TestCreateCampaignStartTimeInThePast(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	startTimeInThePast := campaign.StartTime.Add(-time.Hour)
	campaign.StartTime = startTimeInThePast
	testHelper.C4eAirdropUtils.CreateCampaignError(acountsAddresses[0].String(), campaign, fmt.Sprintf("create airdrop campaign - start time in the past error (%s < %s): wrong param value", campaign.StartTime, testHelper.Context.BlockTime()))
}

func TestCreateCampaignNegativeInitialClaimAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = sdk.NewInt(-100)
	testHelper.C4eAirdropUtils.CreateCampaignError(acountsAddresses[0].String(), campaign, fmt.Sprintf("create airdrop campaign - initial claim free amount (%s) cannot be negative: wrong param value", campaign.InitialClaimFreeAmount))
}

func TestCreateCampaignNegativeFeegrantAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = sdk.NewInt(-100)
	testHelper.C4eAirdropUtils.CreateCampaignError(acountsAddresses[0].String(), campaign, fmt.Sprintf("create airdrop campaign - feegrant amount (%s) cannot be negative: wrong param value", campaign.FeegrantAmount))
}

func TestCreateCampaignAndStart(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
}

func TestCreateManyCampaignsAndStart(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaigns := prepareNTestCampaigns(testHelper.Context, 10)
	for i, campaign := range campaigns {
		testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
		testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), uint64(i))
	}
}

func TestCreateCampaignAndStartTimeAfterTimeNowError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	blockTime := campaign.StartTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.StartCampaignError(acountsAddresses[0].String(), 0, fmt.Sprintf("start airdrop campaign - campaign with id 0 start time in the past error (%s < %s): wrong param value", campaign.StartTime, blockTime))
}

func TestCreateCampaignAndStartOwnerNotValidError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaignError(acountsAddresses[1].String(), 0, "start airdrop campaign you are not the owner of this campaign: tx intended signer does not match the given signer")
}

func TestCreateCampaignCampaignDoesntExistError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaignError(acountsAddresses[0].String(), 1, "start airdrop campaign campaign with id 1 not found: entity does not exist")
}

func TestCreateCampaignCampaignEnabledError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	testHelper.C4eAirdropUtils.StartCampaignError(acountsAddresses[0].String(), 0, "start airdrop campaign campaign with id 0 has already started: entity already exists")
}

func TestCreateCampaignCloseCampaignCloseActionBurn(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CampaignCloseBurn)
}

func TestCreateCampaignCloseCampaignCloseActionSendToOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CampaignCloseSendToOwner)
}

func TestCreateCampaignCloseCampaignCloseActionSendToCommunityPool(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CampaignCloseSendToCommunityPool)
}

func TestCreateCampaignCloseCampaignCloseActionBurnAndFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = sdk.NewInt(1000)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum.Add(campaign.FeegrantAmount.MulRaw(int64(len(airdropEntries)))))
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CampaignCloseBurn)
}

func TestCreateCampaignCloseCampaignCloseActionSendToOwnerAndFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = sdk.NewInt(1000)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum.Add(campaign.FeegrantAmount.MulRaw(int64(len(airdropEntries)))))
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CampaignCloseSendToOwner)
}

func TestCreateCampaignCloseCampaignCloseActionSendToCommunityPoolAndFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = sdk.NewInt(1000)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum.Add(campaign.FeegrantAmount.MulRaw(int64(len(airdropEntries)))))
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CampaignCloseSendToCommunityPool)
}

func TestCreateCampaignCloseCampaignWrongCloseAction(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseCampaignError(acountsAddresses[0].String(), 0, types.CampaignCloseActionUnspecified, "wrong campaign close action type: invalid type")
}

func TestCreateManyCampaignsAndClose(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaigns := prepareNTestCampaigns(testHelper.Context, 10)
	contextNow := testHelper.Context.BlockTime()
	for i, campaign := range campaigns {
		testHelper.SetContextBlockTime(contextNow)
		testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
		testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), uint64(i))
		blockTime := campaign.EndTime.Add(time.Minute)
		testHelper.SetContextBlockTime(blockTime)
		testHelper.C4eAirdropUtils.CloseCampaign(acountsAddresses[0].String(), uint64(i), types.CampaignCloseBurn)
	}
}

func TestCreateCampaignCloseCampaignCampaignNotStartedError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseCampaignError(acountsAddresses[0].String(), 0, types.CampaignCloseBurn, fmt.Sprintf("close airdrop campaign - campaign with id %d is already closed or have not started yet error: campaign is disabled", 0))
}

func TestCreateCampaignCloseCampaignCampaignNotOverYetError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	testHelper.C4eAirdropUtils.CloseCampaignError(acountsAddresses[0].String(), 0, types.CampaignCloseAction_CLOSE_ACTION_UNSPECIFIED, fmt.Sprintf("close airdrop campaign - campaign with id %d campaign is not over yet (endtime - %s < %s): wrong param value", 0, campaign.EndTime, testHelper.Context.BlockTime()))
}

func TestCreateCampaignCloseCampaignCampaigDoesntExistError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseCampaignError(acountsAddresses[0].String(), 1, types.CampaignCloseAction_CLOSE_ACTION_UNSPECIFIED, "close airdrop campaign - campaign with id 1 not found error: entity does not exist")
}

func TestCreateCampaignCloseCampaignYouAreNotTheOwnerErrror(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.CloseCampaignError(acountsAddresses[1].String(), 0, types.CampaignCloseAction_CLOSE_ACTION_UNSPECIFIED, "close airdrop campaign - you are not the owner error: tx intended signer does not match the given signer")
}

func prepareNTestCampaigns(ctx sdk.Context, n int) []types.Campaign {
	campaigns := make([]types.Campaign, n)
	for i := range campaigns {
		campaigns[i] = prepareTestCampaign(ctx)
	}
	return campaigns
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
