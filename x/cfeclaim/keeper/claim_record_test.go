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

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimRecordEntries)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddManyUsersEntries(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries1, amountSum := createTestClaimRecordEntries(acountsAddresses[0:5], 100000000)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	claimRecordEntries2, amountSum := createTestClaimRecordEntries(acountsAddresses[5:10], 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimRecordEntries1)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimRecordEntries2)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddManyUsersEntriesVestingPoolCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	ownerAddress := acountsAddresses[0]
	claimRecordEntries1, amountSum1 := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses[0:5], 30)
	claimRecordEntries2, amountSum2 := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses[5:10], 30)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, amountSum1.AmountOf(testHelper.C4eVestingUtils.GetVestingDenom()).
		Add(amountSum2.AmountOf(testHelper.C4eVestingUtils.GetVestingDenom())), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries1)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries2)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsBalanceToSmall(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum.Sub(sdk.NewCoin(amountSum[1].Denom, math.NewInt(2))))
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimRecordEntries, "owner balance is too small (1000000045uc4e,1000000043uc4e2 < 1000000045uc4e,1000000045uc4e2): insufficient funds")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsEmptyAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	claimRecordEntries[0].Amount = sdk.NewCoins()
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimRecordEntries, "claim record entry index 0: claim record amount must be positive: wrong param value")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsZeroAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	claimRecordEntries[0].Amount[0].Amount = math.ZeroInt()
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimRecordEntries, "claim record entry index 0: wrong claim record entry amount (coin 0uc4e amount is not positive): wrong param value")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsEmptyAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	claimRecordEntries[0].UserEntryAddress = ""
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimRecordEntries, "claim record entry index 0: claim record entry user entry address parsing error (empty address string is not allowed): wrong param value")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsWrongOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[1], 0, claimRecordEntries, fmt.Sprintf("address %s is not owner of campaign with id %d: tx intended signer does not match the given signer", acountsAddresses[1], 0))
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsCampaignDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 1, claimRecordEntries, "campaign with id 1 not found: not found")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsCampaignNotEnabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimRecordEntries)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsCampaignIsOver(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimRecordEntries, fmt.Sprintf("campaign with id 0 campaign is over (end time - %s < %s): wrong param value", campaign.EndTime, blockTime))
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsclaimRecordExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 1113321)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimRecordEntries)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimRecordEntries, fmt.Sprintf("claim record entry index 0: campaignId 0 already exists for address: %s: entity already exists", claimRecordEntries[0].UserEntryAddress))
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsInitialClaimAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 12412512)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = math.NewInt(50000000)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimRecordEntries)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsCorrectFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	feegrantAmount := math.NewInt(2500000)
	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 512451234)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = feegrantAmount
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	feegrantCoins := sdk.NewCoins(sdk.NewCoin(testHelper.StakingUtils.GetStakingDenom(), campaign.FeegrantAmount.MulRaw(int64(len(acountsAddresses)))))
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], feegrantCoins)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimRecordEntries)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsWrongSrcAccountBalanceFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	feegrantAmount := math.NewInt(2500000)
	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 1234123)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = feegrantAmount
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.AddMission(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(acountsAddresses[0].String(), 0, nil, nil)

	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(acountsAddresses[0], 0, claimRecordEntries, "owner balance is too small (12341275uc4e,12341275uc4e2 < 37341275uc4e,12341275uc4e2): insufficient funds")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsCampaignVestingPoolCampaignNotStarted(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	ownerAddress := acountsAddresses[0]
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	claimRecordEntries, amountSum := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestAddClaimRecordsCampaignVestingInPoolAmountTooSmall(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	ownerAddress := acountsAddresses[0]
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	claimRecordEntries, amountSum := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 3000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	mission := prepareTestMission()

	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(ownerAddress, 0, claimRecordEntries, fmt.Sprintf("%s is smaller than %s: insufficient funds", sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000)), amountSum))
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecord(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 12354123)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimRecordEntries[2].UserEntryAddress, claimRecordEntries[2].Amount)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecordEnabledError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 1231451)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	testHelper.C4eClaimUtils.DeleteClaimRecordError(ownerAddress, 0, claimRecordEntries[1].UserEntryAddress, "campaign must have RemovableClaimRecords flag set to true to be able to delete its entries: invalid type")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecordInititalMissionClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 998765)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.RemovableClaimRecords = true
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	var amountDiff sdk.Coins
	for _, coin := range claimRecordEntries[1].Amount {
		amountDiff = amountDiff.Add(sdk.NewCoin(coin.Denom, sdk.NewInt(199753)))
	}
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimRecordEntries[1].UserEntryAddress, amountDiff)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecordEverythingClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 123451)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.RemovableClaimRecords = true
	mission := prepareTestMission()
	mission.MissionType = types.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	var amountDiff sdk.Coins
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimRecordEntries[1].UserEntryAddress, amountDiff)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecordAndInititalClaimInititalMissionClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.RemovableClaimRecords = true
	campaign.InitialClaimFreeAmount = math.NewInt(15)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	var amountDiff sdk.Coins
	for _, coin := range claimRecordEntries[1].Amount {
		amountDiff = amountDiff.Add(sdk.NewCoin(coin.Denom, sdk.NewInt(6)))
	}

	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimRecordEntries[1].UserEntryAddress, amountDiff)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecordAndInititalClaimAndFeegrantInititalMissionClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.RemovableClaimRecords = true
	campaign.InitialClaimFreeAmount = math.NewInt(15)
	campaign.FeegrantAmount = math.NewInt(15)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	feegrantCoin := sdk.NewCoin(testHelper.StakingUtils.GetStakingDenom(), campaign.FeegrantAmount.MulRaw(int64(len(acountsAddresses))))
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum.Add(feegrantCoin))
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	var amountDiff sdk.Coins
	for _, coin := range claimRecordEntries[1].Amount {
		amountDiff = amountDiff.Add(sdk.NewCoin(coin.Denom, sdk.NewInt(6)))
	}
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimRecordEntries[1].UserEntryAddress, amountDiff)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecordTwoMissionsClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.RemovableClaimRecords = true
	mission := prepareTestMission()
	mission.MissionType = types.MissionClaim
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	var amountDiff sdk.Coins
	for _, coin := range claimRecordEntries[1].Amount {
		amountDiff = amountDiff.Add(sdk.NewCoin(coin.Denom, sdk.NewInt(6)))
	}
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimRecordEntries[1].UserEntryAddress, amountDiff)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestVestingPoolCampaignDeleteClaimRecordInititalMissionClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	campaign.RemovableClaimRecords = true
	mission := prepareTestMission()
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	var amountDiff sdk.Coins
	for _, coin := range claimRecordEntries[1].Amount {
		amountDiff = amountDiff.Add(sdk.NewCoin(coin.Denom, sdk.NewInt(6)))
	}
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimRecordEntries[1].UserEntryAddress, amountDiff)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestVestingPoolCampaignManyDenoms(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(ownerAddress, 0, claimRecordEntries, "claim record entry index 0: for vesting pool campaigns, the claim record entry"+
		" must have only one coin with the denomination currently used by the cfevesting module (uc4e): wrong param value")
}

func TestVestingPoolCampaignWrongDenom(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	claimRecordEntries[0].Amount = sdk.NewCoins(sdk.NewCoin("wrongDenom", sdk.NewInt(100)))
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecordsError(ownerAddress, 0, claimRecordEntries, "claim record entry index 0: for vesting pool campaigns, the claim record entry"+
		" must have only one coin with the denomination currently used by the cfevesting module (uc4e): wrong param value")
}

func TestVestingPoolCampaignDeleteClaimRecordTwoMissionsClaimed(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	campaign.RemovableClaimRecords = true
	mission := prepareTestMission()
	mission.MissionType = types.MissionClaim
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.ClaimInitial(acountsAddresses[1], 0)
	testHelper.C4eClaimUtils.ClaimMission(0, 1, acountsAddresses[1])
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimRecordEntries[1].UserEntryAddress, claimRecordEntries[1].Amount.Sub(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(25))))
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestVestingPoolCampaignDeleteClaimWithFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
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
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	testHelper.C4eClaimUtils.AddMission(ownerAddress.String(), 0, mission)
	feegrantCoin := sdk.NewCoin(testHelper.StakingUtils.GetStakingDenom(), campaign.FeegrantAmount.MulRaw(int64(len(acountsAddresses))))
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum.Add(feegrantCoin))
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress.String(), 0, nil, nil)
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimRecordEntries[1].UserEntryAddress, claimRecordEntries[1].Amount)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecordVestingPoolCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createVestingPoolCampaignTestClaimRecordEntries(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, amountSum.AmountOf(testHelper.C4eVestingUtils.GetVestingDenom()), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.DeleteClaimRecord(ownerAddress, 0, claimRecordEntries[2].UserEntryAddress, claimRecordEntries[2].Amount)
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecordUserEntryNotExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(12, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses[:10], 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.DeleteClaimRecordError(ownerAddress, 0, acountsAddresses[10].String(), fmt.Sprintf("userEntry %s doesn't exist: not found", acountsAddresses[10]))
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecordCampaignNotExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.DeleteClaimRecordError(ownerAddress, 1, claimRecordEntries[2].UserEntryAddress, "campaign with id 1 not found: not found")
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func TestDeleteClaimRecordWrongOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimRecordEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 30)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimRecordEntries)
	testHelper.C4eClaimUtils.DeleteClaimRecordError(acountsAddresses[1], 0, claimRecordEntries[2].UserEntryAddress, fmt.Sprintf("address %s is not owner of campaign with id %d: tx intended signer does not match the given signer", acountsAddresses[1], 0))
	testHelper.C4eClaimUtils.ValidateGenesisAndInvariants()
}

func createTestClaimRecordEntries(addresses []sdk.AccAddress, startAmount int) (claimRecordEntries []*types.ClaimRecordEntry, amountSum sdk.Coins) {
	for i, addr := range addresses {
		coinsAmount := math.NewInt(int64(startAmount + i))
		newclaimRecord := types.ClaimRecordEntry{
			UserEntryAddress: addr.String(),
			Amount: sdk.NewCoins(
				sdk.NewCoin(testenv.DefaultTestDenom, coinsAmount),
				sdk.NewCoin(testenv.DefaultTestDenom2, coinsAmount),
			),
		}
		amountSum = amountSum.Add(newclaimRecord.Amount...)
		claimRecordEntries = append(claimRecordEntries, &newclaimRecord)
	}
	return
}

func createVestingPoolCampaignTestClaimRecordEntries(addresses []sdk.AccAddress, startAmount int) (claimRecordEntries []*types.ClaimRecordEntry, amountSum sdk.Coins) {
	for i, addr := range addresses {
		coinsAmount := math.NewInt(int64(startAmount + i))
		newclaimRecord := types.ClaimRecordEntry{
			UserEntryAddress: addr.String(),
			Amount: sdk.NewCoins(
				sdk.NewCoin(testenv.DefaultTestDenom, coinsAmount),
			),
		}
		amountSum = amountSum.Add(newclaimRecord.Amount...)
		claimRecordEntries = append(claimRecordEntries, &newclaimRecord)
	}
	return
}

func createCampaignMissionAndEnable(testHelper *testapp.TestHelper, ownerAddress string) {
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress, campaign)
	testHelper.C4eClaimUtils.AddMission(ownerAddress, 0, mission)
	testHelper.C4eClaimUtils.EnableCampaign(ownerAddress, 0, nil, nil)
}

func createNUsersEntries(keeper *keeper.Keeper, ctx sdk.Context, numberOfUsersEntries int, numberOfClaimEntreis int, addClaimAddress bool, addCompletedMissions bool) []types.UserEntry {
	userEntry := make([]types.UserEntry, numberOfUsersEntries)
	for i := range userEntry {
		userEntry[i].Address = strconv.Itoa(i)
		claimRecords := make([]types.ClaimRecord, numberOfClaimEntreis)
		for j := range claimRecords {
			claimRecords[i].Address = strconv.Itoa(1000000 + i)
			claimRecords[j].CampaignId = uint64(2000000 + i)
			claimRecords[j].Amount = sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(int64(3000000+i))))
			if addCompletedMissions {
				claimRecords[j].CompletedMissions = []uint64{uint64(4000000 + i), uint64(5000000 + i), uint64(6000000 + i)}
			}
		}
		keeper.SetUserEntry(ctx, userEntry[i])
	}
	return userEntry
}
