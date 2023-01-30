package keeper_test

import (
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"testing"
)

func TestCorrectInitialClaim(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)

	airdropEntries, airdropCoinsSum := createTestAirdropEntries(acountsAddresses, 100000000)
	createCampaignMissionAndStart(testHelper, acountsAddresses[0].String())
	testHelper.C4eAirdropUtils.AddCoinsToAirdropEntrisCreator(acountsAddresses[0], airdropCoinsSum)
	testHelper.C4eAirdropUtils.AddAirdropEntries(acountsAddresses[0], 0, airdropEntries)
	testHelper.C4eAirdropUtils.ClaimInitial(acountsAddresses[1], 0, 800000000)
}
