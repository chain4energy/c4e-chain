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

func TestUserAirdropEntriesGet(t *testing.T) {
	k, ctx := keepertest.CfeairdropKeeper(t)
	items := createNUserAirdropEntries(k, ctx, 10, 0, false, false)
	for _, item := range items {
		rst, found := k.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNUserAirdropEntries(k, ctx, 10, 10, false, false)
	for _, item := range items {
		rst, found := k.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNUserAirdropEntries(k, ctx, 10, 10, true, false)
	for _, item := range items {
		rst, found := k.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}

	items = createNUserAirdropEntries(k, ctx, 10, 10, false, true)
	for _, item := range items {
		rst, found := k.GetUserAirdropEntries(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestUserAirdropEntriesGetAll(t *testing.T) {
	k, ctx := keepertest.CfeairdropKeeper(t)
	items := createNUserAirdropEntries(k, ctx, 10, 0, false, false)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetUsersAirdropEntries(ctx)),
	)

	items = createNUserAirdropEntries(k, ctx, 10, 10, true, true)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetUsersAirdropEntries(ctx)),
	)
}

func TestAddUserAirdropEntries(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, airdropEntries)
}

func TestAddManyUserAirdropEntries(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries1, airdropCoinsSum := createTestAirdropEntries(acountsAddresses[0:5], 100000000)
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	airdropEntries2, airdropCoinsSum := createTestAirdropEntries(acountsAddresses[5:10], 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, airdropEntries1)
	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, airdropEntries2)
}

func TestAddUserAirdropEntriesBalanceToSmall(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum.Sub(sdk.NewInt(2)))
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[0], 0, airdropEntries, fmt.Sprintf("%suc4e is smaller than %suc4e: insufficient funds", airdropCoinsSum.Sub(sdk.NewInt(2)), airdropCoinsSum))
}

func TestAddUserAirdropEntriesEmptyAirdropCoins(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	airdropEntries[0].AirdropCoins = sdk.NewCoins()
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - airdrop entry at index 0 airdrop entry must has at least one coin: wrong param value")
}

func TestAddUserAirdropEntriesZeroAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	airdropEntries[0].AirdropCoins[0].Amount = sdk.ZeroInt()
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - airdrop entry at index 0 amount is 0: wrong param value")
}

func TestAddUserAirdropEntriesEmptyAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	airdropEntries[0].Address = ""
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - airdrop entry empty address on index 0: wrong param value")
}

func TestAddUserAirdropEntriesWrongOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[1], 0, airdropEntries, "add campaign entries - you are not the owner of campaign with id 0: tx intended signer does not match the given signer")
}

func TestAddUserAirdropEntriesCampaignDoesntExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[0], 1, airdropEntries, "add campaign entries -  campaign with id 1 doesn't exist: entity does not exist")
}

func TestAddUserAirdropEntriesCampaignNotEnabled(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - campaign 0 is disabled: campaign is disabled")
}

func TestAddUserAirdropEntriesCampaignIsOver(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[0], 0, airdropEntries, fmt.Sprintf("add campaign entries - campaign 0 is disabled (end time %s < %s): campaign is disabled", campaign.EndTime, blockTime))
}

func TestAddUserAirdropEntriesAirdropEntryExist(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[0], 0, airdropEntries, fmt.Sprintf("campaignId 0 already exists for address: %s: entity already exists", airdropEntries[0].Address))
}

func TestAddUserAirdropEntriesInitialClaimAmountError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = sdk.NewInt(100000000000000000)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - airdrop entry at index 0 initial claim amount 80000000 < campaign initial claim free amount (100000000000000000): wrong param value")
}

func TestAddUserAirdropEntriesInitialClaimAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = sdk.NewInt(50000000)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, airdropEntries)
}

func TestAddUserAirdropEntriesCorrectFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	feegrantAmount := sdk.NewInt(2500000)
	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = feegrantAmount
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], feegrantAmount.MulRaw(int64(len(acountsAddresses))))
	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, airdropEntries)
}

func TestAddUserAirdropEntriesWrongSrcAccountBalanceFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	feegrantAmount := sdk.NewInt(2500000)
	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = feegrantAmount
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(acountsAddresses[0].String(), 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(acountsAddresses[0].String(), 0)

	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntriesError(acountsAddresses[0], 0, airdropEntries, "add campaign entries - owner balance is too small (1000000045uc4e < 1025000045uc4e): insufficient funds")
}

func createTestAirdropEntries(addresses []sdk.AccAddress, startAmount int) (airdropEntries []*types.AirdropEntry, airdropCoinsSum sdk.Int) {
	airdropCoinsSum = sdk.ZeroInt()
	for i, addr := range addresses {
		coinsAmount := sdk.NewInt(int64(startAmount + i))
		newAirdropEntry := types.AirdropEntry{
			Address:      addr.String(),
			AirdropCoins: sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, coinsAmount)),
		}
		airdropCoinsSum = airdropCoinsSum.Add(coinsAmount)
		airdropEntries = append(airdropEntries, &newAirdropEntry)
	}
	return
}

func createCampaignMissionAndStart(testHelper *testapp.TestHelper, ownerAddress string) {
	campaign := prepareTestCampaign(testHelper.Context)
	mission := prepareTestMission()
	testHelper.C4eAirdropUtils.CreateAirdropCampaign(ownerAddress, campaign)
	testHelper.C4eAirdropUtils.AddMissionToAirdropCampaign(ownerAddress, 0, mission)
	testHelper.C4eAirdropUtils.StartAirdropCampaign(ownerAddress, 0)
}

func createNUserAirdropEntries(keeper *keeper.Keeper, ctx sdk.Context, numberOfUserAirdropEntries int, numberOfAirdropEntreis int, addClaimAddress bool, addCompletedMissions bool) []types.UserAirdropEntries {
	userAirdropEntries := make([]types.UserAirdropEntries, numberOfUserAirdropEntries)
	for i := range userAirdropEntries {
		userAirdropEntries[i].Address = strconv.Itoa(i)
		if addClaimAddress {
			userAirdropEntries[i].ClaimAddress = strconv.Itoa(1000000 + i)
		}
		airdropEntryStates := make([]types.AirdropEntry, numberOfAirdropEntreis)
		for j := range airdropEntryStates {
			airdropEntryStates[j].CampaignId = uint64(2000000 + i)
			airdropEntryStates[j].AirdropCoins = sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(int64(3000000+i))))
			if addCompletedMissions {
				airdropEntryStates[j].CompletedMissions = []uint64{uint64(4000000 + i), uint64(5000000 + i), uint64(6000000 + i)}
			}

		}
		keeper.SetUserAirdropEntries(ctx, userAirdropEntries[i])
	}
	return userAirdropEntries
}
