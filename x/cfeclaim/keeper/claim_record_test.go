package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

func TestUsersEntriesGet(t *testing.T) {
	k, ctx := keepertest.CfeclaimKeeper(t)
	items := createNUsersEntries(k, ctx, 10, 0, false, false)
	for _, item := range items {
		rst, found := k.GetUserEntry(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNUsersEntries(k, ctx, 10, 10, false, false)
	for _, item := range items {
		rst, found := k.GetUserEntry(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNUsersEntries(k, ctx, 10, 10, true, false)
	for _, item := range items {
		rst, found := k.GetUserEntry(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNUsersEntries(k, ctx, 10, 10, false, true)
	for _, item := range items {
		rst, found := k.GetUserEntry(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestUsersEntriesGetAll(t *testing.T) {
	k, ctx := keepertest.CfeclaimKeeper(t)
	items := createNUsersEntries(k, ctx, 10, 0, false, false)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllUsersEntries(ctx)),
	)

	items = createNUsersEntries(k, ctx, 10, 10, true, true)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllUsersEntries(ctx)),
	)
}

func TestAddClaimRecords(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
}

func TestAddManyUsersEntries(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries1, amountSum := createTestClaimRecords(acountsAddresses[0:5], 100000000)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	claimEntries2, amountSum := createTestClaimRecords(acountsAddresses[5:10], 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries1)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries2)
}

func TestAddManyUsersEntriesVestingPoolCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	ownerAddress := acountsAddresses[0]
	claimEntries1, amountSum1 := createTestClaimRecords(acountsAddresses[0:5], 30)
	claimEntries2, amountSum2 := createTestClaimRecords(acountsAddresses[5:10], 30)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, amountSum1.Add(amountSum2), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries1)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries2)
}

func TestAddClaimRecordsBalanceToSmall(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum.Sub(math.NewInt(2)))
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimEntries, "owner balance is too small (1000000043uc4e < 1000000045uc4e): insufficient funds")
}

func TestAddClaimRecordsEmptyAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	claimEntries[0].Amount = sdk.NewCoins()
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimEntries, "claim records index 0: claim record must has at least one coin and all amounts must be positive: wrong param value")
}

func TestAddClaimRecordsZeroAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	claimEntries[0].Amount[0].Amount = math.ZeroInt()
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimEntries, "claim records index 0: claim record must has at least one coin and all amounts must be positive: wrong param value")
}

func TestAddClaimRecordsEmptyAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	claimEntries[0].Address = ""
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimEntries, "claim records index 0: claim record empty address: wrong param value")
}

func TestAddClaimRecordsWrongOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[1], 0, claimEntries, "you are not the campaign owner: wrong transaction signer")
}

func TestAddClaimRecordsCampaignDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 1, claimEntries, "campaign with id 1 not found: entity does not exist")
}

func TestAddClaimRecordsCampaignNotEnabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
}

func TestAddClaimRecordsCampaignIsOver(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimEntries, fmt.Sprintf("campaign with id 0 campaign is over (end time - %s < %s): wrong param value", campaign.EndTime, blockTime))
}

func TestAddClaimRecordsclaimRecordExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimEntries, fmt.Sprintf("claim records index 0: campaignId 0 already exists for address: %s: entity already exists", claimEntries[0].Address))
}

func TestAddClaimRecordsInitialClaimAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = math.NewInt(50000000)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
}

func TestAddClaimRecordsCorrectFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	feegrantAmount := math.NewInt(2500000)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = feegrantAmount
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], feegrantAmount.MulRaw(int64(len(acountsAddresses))))
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
}

func TestAddClaimRecordsWrongSrcAccountBalanceFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	feegrantAmount := math.NewInt(2500000)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = feegrantAmount
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimEntries, "owner balance is too small (1000000045uc4e < 1025000045uc4e): insufficient funds")
}

func TestAddClaimRecordsCampaignVestingPoolCampaignNotStarted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	ownerAddress := acountsAddresses[0]
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
}

func TestAddClaimRecordsCampaignVestingInPoolAmountTooSmall(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	ownerAddress := acountsAddresses[0]
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 3000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(ownerAddress, 0, claimEntries, fmt.Sprintf("%s is smaller than %s: insufficient funds", sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000)), sdk.NewCoin(testenv.DefaultTestDenom, amountSum)))
}

func TestDeleteClaimRecord(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimEntries[2].Address, claimEntries[2].Amount)
}

func TestDeleteClaimRecordEnabledError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 25)
	testHelper.C4eClaimUtils.DeleteClaimRecordError(ownerAddress, 0, claimEntries[1].Address, "campaign must have RemovableClaimRecords flag set to true to be able to delete its entries: invalid type")
}

func TestDeleteClaimRecordInititalMissionClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.RemovableClaimRecords = true
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 25)
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimEntries[1].Address, claimEntries[1].Amount.Sub(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(25))))
}

func TestDeleteClaimRecordTwoMissionsClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.RemovableClaimRecords = true
	mission := prepareTestMission()
	mission.MissionType = types.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 19)
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimEntries[1].Address, claimEntries[1].Amount.Sub(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(25))))
}

func TestVestingPoolCampaignDeleteClaimRecordInititalMissionClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	campaign.RemovableClaimRecords = true
	mission := prepareTestMission()
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 25)
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimEntries[1].Address, claimEntries[1].Amount.Sub(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(25))))
}

func TestVestingPoolCampaignDeleteClaimRecordTwoMissionsClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	campaign.RemovableClaimRecords = true
	mission := prepareTestMission()
	mission.MissionType = types.MissionClaim
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0, 19)
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimEntries[1].Address, claimEntries[1].Amount.Sub(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(25))))
}

func TestVestingPoolCampaignDeleteClaimWithFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	campaign.RemovableClaimRecords = true
	campaign.FeegrantAmount = math.NewInt(10000)
	mission := prepareTestMission()
	mission.MissionType = types.MissionClaim
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum.Add(campaign.FeegrantAmount.MulRaw(int64(len(claimEntries)))))
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimEntries[1].Address, claimEntries[1].Amount)
}

func TestDeleteClaimRecordVestingPoolCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, amountSum, 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimEntries[2].Address, claimEntries[2].Amount)
}

func TestDeleteClaimRecordUserEntryNotExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(12, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses[:10], 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.DeleteClaimRecordError(ownerAddress, 0, acountsAddresses[10].String(), fmt.Sprintf("userEntry %s doesn't exist: entity does not exist", acountsAddresses[10]))
}

func TestDeleteClaimRecordCampaignNotExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.DeleteClaimRecordError(ownerAddress, 1, claimEntries[2].Address, "campaign with id 1 not found: entity does not exist")
}

func TestDeleteClaimRecordWrongOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	testHelper.C4eClaimUtils.DeleteClaimRecordError(acountsAddresses[1], 0, claimEntries[2].Address, "you are not the campaign owner: wrong transaction signer")
}

func createTestClaimRecords(addresses []sdk.AccAddress, startAmount int) (claimEntries []*types.ClaimRecord, amountSum math.Int) {
	amountSum = math.ZeroInt()
	for i, addr := range addresses {
		coinsAmount := math.NewInt(int64(startAmount + i))
		newclaimRecord := types.ClaimRecord{
			Address: addr.String(),
			Amount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, coinsAmount)),
		}
		amountSum = amountSum.Add(coinsAmount)
		claimEntries = append(claimEntries, &newclaimRecord)
	}
	return
}

func createCampaignMissionAndEnable(testHelper *testapp.TestHelper, ownerAddress string) {
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress, campaign)
	testHelper.C4eClaimUtils.AddMissionToCampaign(ownerAddress, 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress, 0, nil, nil)
}

func createNUsersEntries(keeper *keeper.Keeper, ctx sdk.Context, numberOfUsersEntries int, numberOfClaimEntreis int, addClaimAddress bool, addCompletedMissions bool) []types.UserEntry {
	userEntry := make([]types.UserEntry, numberOfUsersEntries)
	for i := range userEntry {
		userEntry[i].Address = strconv.Itoa(i)
		if addClaimAddress {
			userEntry[i].ClaimAddress = strconv.Itoa(1000000 + i)
		}
		claimRecordStates := make([]types.ClaimRecord, numberOfClaimEntreis)
		for j := range claimRecordStates {
			claimRecordStates[j].CampaignId = uint64(2000000 + i)
			claimRecordStates[j].Amount = sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(int64(3000000+i))))
			if addCompletedMissions {
				claimRecordStates[j].CompletedMissions = []uint64{uint64(4000000 + i), uint64(5000000 + i), uint64(6000000 + i)}
			}

		}
		keeper.SetUserEntry(ctx, userEntry[i])
	}
	return userEntry
}
