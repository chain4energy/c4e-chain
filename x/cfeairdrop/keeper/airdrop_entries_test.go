package keeper_test

import (
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

func TestUsersEntriesGet(t *testing.T) {
	k, ctx := keepertest.CfeairdropKeeper(t)
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
	k, ctx := keepertest.CfeairdropKeeper(t)
	items := createNUsersEntries(k, ctx, 10, 0, false, false)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetUsersEntries(ctx)),
	)

	items = createNUsersEntries(k, ctx, 10, 10, true, true)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetUsersEntries(ctx)),
	)
}

func TestAddUsersEntries(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
}

func TestAddManyUsersEntries(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries1, amountSum := createTestClaimRecords(acountsAddresses[0:5], 100000000)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	airdropEntries2, amountSum := createTestClaimRecords(acountsAddresses[5:10], 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries1)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries2)
}

func TestAddUsersEntriesBalanceToSmall(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum.Sub(sdk.NewInt(2)))
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[0], 0, airdropEntries, fmt.Sprintf("add campaign entries - owner balance is too small (1000000043uc4e < 1000000045uc4e): insufficient funds"))
}

func TestAddUsersEntriesEmptyAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	airdropEntries[0].Amount = sdk.NewCoins()
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - airdrop entry at index 0 airdrop entry must has at least one coin: wrong param value")
}

func TestAddUsersEntriesZeroAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	airdropEntries[0].Amount[0].Amount = sdk.ZeroInt()
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - airdrop entry at index 0 amount is 0: wrong param value")
}

func TestAddUsersEntriesEmptyAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	airdropEntries[0].Address = ""
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - airdrop entry empty address on index 0: wrong param value")
}

func TestAddUsersEntriesWrongOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[1], 0, airdropEntries, "add campaign entries - you are not the owner of campaign with id 0: tx intended signer does not match the given signer")
}

func TestAddUsersEntriesCampaignDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[0], 1, airdropEntries, "add campaign entries -  campaign with id 1 doesn't exist: entity does not exist")
}

func TestAddUsersEntriesCampaignNotEnabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - campaign 0 is disabled: campaign is disabled")
}

func TestAddUsersEntriesCampaignIsOver(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[0], 0, airdropEntries, fmt.Sprintf("add campaign entries - campaign 0 is disabled (end time %s < %s): campaign is disabled", campaign.EndTime, blockTime))
}

func TestAddUsersEntriesclaimRecordExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[0], 0, airdropEntries, fmt.Sprintf("campaignId 0 already exists for address: %s: entity already exists", airdropEntries[0].Address))
}

func TestAddUsersEntriesInitialClaimAmountError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = sdk.NewInt(100000000000000000)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - airdrop entry at index 0 initial claim amount 80000000 < campaign initial claim free amount (100000000000000000): wrong param value")
}

func TestAddUsersEntriesInitialClaimAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = sdk.NewInt(50000000)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
}

func TestAddUsersEntriesCorrectFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	feegrantAmount := sdk.NewInt(2500000)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = feegrantAmount
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], feegrantAmount.MulRaw(int64(len(acountsAddresses))))
	testHelper.C4eAirdropUtils.AddClaimRecords(acountsAddresses[0], 0, airdropEntries)
}

func TestAddUsersEntriesWrongSrcAccountBalanceFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	feegrantAmount := sdk.NewInt(2500000)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = feegrantAmount
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eAirdropUtils.AddClaimRecordsError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - owner balance is too small (1000000045uc4e < 1025000045uc4e): insufficient funds")
}

func TestAddClaimRecordsFromWhitelistedAccount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := sdk.NewInt(100000000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(999990045)
	coins := sdk.Coins{{Amount: sendAmount, Denom: testenv.DefaultTestDenom}}
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(accAddr2.String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(accAddr2.String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(accAddr2.String(), 0)
	airdropEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	fmt.Println(amountSum)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(accAddr2, sdk.NewInt(10000))
	testHelper.C4eAirdropUtils.AddClaimRecords(accAddr2, 0, airdropEntries)

	fmt.Println(accAddr2.String())
}

func TestAddClaimRecordsFromWhitelistedAccountTimeInFuture(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := sdk.NewInt(100000000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(999990045)
	coins := sdk.Coins{{Amount: sendAmount, Denom: testenv.DefaultTestDenom}}
	startTime := testHelper.Context.BlockTime()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.EndTime = endTime
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(accAddr2.String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(accAddr2.String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(accAddr2.String(), 0)
	airdropEntries, _ := createTestClaimRecords(acountsAddresses, 100000000)
	blockTime := startTime.Add(time.Hour * 5)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(accAddr2, sdk.NewInt(10000))
	testHelper.C4eAirdropUtils.AddClaimRecords(accAddr2, 0, airdropEntries)

	fmt.Println(accAddr2.String())
}

func TestAddClaimRecordsFromWhitelistedAccountTimeInFutureBiggerAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := sdk.NewInt(100000000000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(1599990045)
	coins := sdk.Coins{{Amount: sendAmount, Denom: testenv.DefaultTestDenom}}
	startTime := testHelper.Context.BlockTime()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.EndTime = endTime
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(accAddr2.String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(accAddr2.String(), 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(accAddr2.String(), 0)
	airdropEntries, airdropCoinSum := createTestClaimRecords(acountsAddresses, 100000000)
	fmt.Println(airdropCoinSum)
	blockTime := startTime.Add(time.Hour * 5)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.AddCoinsToCampaignOwnerAcc(accAddr2, sdk.NewInt(10000))
	testHelper.C4eAirdropUtils.AddClaimRecords(accAddr2, 0, airdropEntries)

	fmt.Println(accAddr2.String())
}

func createTestClaimRecords(addresses []sdk.AccAddress, startAmount int) (airdropEntries []*types.ClaimRecord, amountSum sdk.Int) {
	amountSum = sdk.ZeroInt()
	for i, addr := range addresses {
		coinsAmount := sdk.NewInt(int64(startAmount + i))
		newclaimRecord := types.ClaimRecord{
			Address: addr.String(),
			Amount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, coinsAmount)),
		}
		amountSum = amountSum.Add(coinsAmount)
		airdropEntries = append(airdropEntries, &newclaimRecord)
	}
	return
}

func createCampaignMissionAndStart(testHelper *testapp.TestHelper, ownerAddress string) {
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateCampaign(ownerAddress, campaign)
	testHelper.C4eAirdropUtils.AddMissionToCampaign(ownerAddress, 0, mission)
	testHelper.C4eAirdropUtils.StartCampaign(ownerAddress, 0)
}

func createNUsersEntries(keeper *keeper.Keeper, ctx sdk.Context, numberOfUsersEntries int, numberOfAirdropEntreis int, addClaimAddress bool, addCompletedMissions bool) []types.UserEntry {
	userEntry := make([]types.UserEntry, numberOfUsersEntries)
	for i := range userEntry {
		userEntry[i].Address = strconv.Itoa(i)
		if addClaimAddress {
			userEntry[i].ClaimAddress = strconv.Itoa(1000000 + i)
		}
		claimRecordStates := make([]types.ClaimRecord, numberOfAirdropEntreis)
		for j := range claimRecordStates {
			claimRecordStates[j].CampaignId = uint64(2000000 + i)
			claimRecordStates[j].Amount = sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(int64(3000000+i))))
			if addCompletedMissions {
				claimRecordStates[j].CompletedMissions = []uint64{uint64(4000000 + i), uint64(5000000 + i), uint64(6000000 + i)}
			}

		}
		keeper.SetUserEntry(ctx, userEntry[i])
	}
	return userEntry
}
