package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func TestCampaignCurrentAmountSumCheckInvariantEmptyClaimsLeft(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	testHelper.C4eClaimUtils.CheckNonNegativeCoinStateInvariant(testHelper.Context, false,
		"cfeclaim: campaigns current amount sum invariant\ncampaigns list is empty\n")
}

func TestCampaignCurrentAmountSumCheckInvariantCorrect(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	testHelper.C4eClaimUtils.CheckNonNegativeCoinStateInvariant(testHelper.Context, false,
		"cfeclaim: campaigns current amount sum invariant\nclaim claims left sum is equal to cfeclaim module account balance\n")
}

func TestCampaignCurrentAmountSumCheckInvariantError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	claimEntries, amountSum := createTestClaimRecordEntries(acountsAddresses, 100000000)
	createCampaignMissionAndEnable(testHelper, acountsAddresses[0].String())
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	coinsToAddToModule := sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(1000))
	amountSumCoins := sdk.NewCoin(testenv.DefaultTestDenom, amountSum)
	testHelper.BankUtils.BankUtils.AddCoinsToModule(testHelper.Context, coinsToAddToModule, cfeclaimtypes.ModuleName)
	testHelper.C4eClaimUtils.CheckNonNegativeCoinStateInvariant(testHelper.Context, true,
		fmt.Sprintf("cfeclaim: campaigns current amount sum invariant\ncampaigns current amount sum is not equal to cfeclaim module account balance (%s != %s)\n",
			amountSumCoins, amountSumCoins.Add(coinsToAddToModule)))
}
