package e2e

import (
	"cosmossdk.io/math"
	v200 "github.com/chain4energy/c4e-chain/app/upgrades/v200"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type MainnetMigrationSetupSuite struct {
	BaseSetupSuite
}

func TestMainnetMigrationSuite(t *testing.T) {
	suite.Run(t, new(MainnetMigrationSetupSuite))
}

func (s *MainnetMigrationSetupSuite) SetupSuite() {
	bytes, err := os.ReadFile("./resources/mainnet-migration-app-state.json")
	if err != nil {
		panic(err)
	}
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, false, false, bytes)
}

func (s *MainnetMigrationSetupSuite) TestMainnetVestingsMigration() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	campaigns := node.QueryCampaigns()
	s.Equal(4, len(campaigns))

	userEntries := node.QueryUserEntries()
	s.Equal(107404, len(userEntries))

	vestingTypes := node.QueryVestingTypes()
	s.Equal(6, len(vestingTypes))
	s.ElementsMatch(createMainnetVestingTypes(), vestingTypes)
	balances, err := node.QueryBalances(v200.AirdropModuleAccountAddress)
	s.NoError(err)
	s.True(balances.IsEqual(sdk.NewCoins()))

	teamdropAccount := node.QueryAccount(v200.TeamdropVestingAccount)
	s.NotNil(teamdropAccount)
	teamdropVestingPools := node.QueryVestingPoolsInfo(v200.TeamdropVestingAccount)
	s.Equal(1, len(teamdropVestingPools))
	s.Equal(teamdropVestingPools[0].VestingType, "Teamdrop")

	teamdropCampaign := node.QueryCampaign("0")
	s.NotNil(teamdropCampaign)

	s.Equal(teamdropVestingPools[0].Reservations[0].Amount, teamdropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
	s.Equal(teamdropVestingPools[0].CurrentlyLocked, "8899990000000")
	s.Equal(teamdropVestingPools[0].InitiallyLocked.Amount, sdk.NewInt(8899990000000))
	s.Equal(teamdropVestingPools[0].SentAmount, math.ZeroInt().String())
	s.Equal(teamdropCampaign.CampaignCurrentAmount, teamdropCampaign.CampaignTotalAmount)

	santadropCampaign := node.QueryCampaign("2")
	s.NotNil(santadropCampaign)
	s.Equal(santadropCampaign.CampaignCurrentAmount, santadropCampaign.CampaignTotalAmount)

	stakedropCampaign := node.QueryCampaign("1")
	s.NotNil(stakedropCampaign)
	s.Equal(stakedropCampaign.CampaignCurrentAmount, stakedropCampaign.CampaignTotalAmount)

	gleamdropCampaign := node.QueryCampaign("3")
	s.NotNil(gleamdropCampaign)
	s.Equal(gleamdropCampaign.CampaignCurrentAmount, gleamdropCampaign.CampaignTotalAmount)

	fairdropVestingPools := node.QueryVestingPoolsInfo(v200.NewAirdropVestingPoolOwner)
	for _, vestingPoolInfo := range fairdropVestingPools {
		if vestingPoolInfo.Name == "Fairdrop" {
			s.Equal(vestingPoolInfo.VestingType, "Fairdrop")
			s.Equal(vestingPoolInfo.InitiallyLocked.Amount, math.NewInt(20000000000000))
			s.Equal(vestingPoolInfo.SentAmount, math.ZeroInt().String())
			s.Equal(vestingPoolInfo.Reservations[0].Amount, stakedropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[1].Amount, santadropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[2].Amount, gleamdropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
		}
	}

}

func (s *MainnetMigrationSetupSuite) validateAllTokensReserved(vestingPoolInfo *cfevestingtypes.VestingPoolInfo, campaignTotalAmount sdk.Coins, campaignCurrentAmount sdk.Coins) {
	s.Equal(campaignTotalAmount, campaignCurrentAmount)
	s.Equal(vestingPoolInfo.CurrentlyLocked, campaignTotalAmount.AmountOf(testenv.DefaultTestDenom).String())
	s.Equal(vestingPoolInfo.InitiallyLocked.Amount, campaignTotalAmount.AmountOf(testenv.DefaultTestDenom))
	s.Equal(vestingPoolInfo.SentAmount, math.ZeroInt().String())
}

func createMainnetVestingTypes() []cfevestingtypes.GenesisVestingType {
	fairdropVestingType := cfevestingtypes.GenesisVestingType{
		Name:              "Fairdrop",
		LockupPeriod:      183,
		LockupPeriodUnit:  "day",
		VestingPeriod:     91,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.10"),
	}

	teamdropVestingType := cfevestingtypes.GenesisVestingType{
		Name:              "Teamdrop",
		Free:              sdk.MustNewDecFromStr("0.10"),
		LockupPeriod:      730,
		LockupPeriodUnit:  "day",
		VestingPeriod:     730,
		VestingPeriodUnit: "day",
	}
	earlyBirdRoundVestingType := cfevestingtypes.GenesisVestingType{
		Name:              "Early-bird round",
		Free:              sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:      0,
		LockupPeriodUnit:  "day",
		VestingPeriod:     274,
		VestingPeriodUnit: "day",
	}
	publicRoundVestingType := cfevestingtypes.GenesisVestingType{
		Name:              "Public round",
		Free:              sdk.MustNewDecFromStr("0.20"),
		LockupPeriod:      0,
		LockupPeriodUnit:  "day",
		VestingPeriod:     183,
		VestingPeriodUnit: "day",
	}
	advisorsPool := cfevestingtypes.GenesisVestingType{
		Name:              "Advisors",
		LockupPeriod:      365,
		LockupPeriodUnit:  "day",
		VestingPeriod:     730,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.000000000000000000"),
	}
	testVestingPool := cfevestingtypes.GenesisVestingType{
		Name:              "TestVestingPool",
		LockupPeriod:      30,
		LockupPeriodUnit:  "second",
		VestingPeriod:     30,
		VestingPeriodUnit: "second",
		Free:              sdk.MustNewDecFromStr("0.000000000000000000"),
	}
	return []cfevestingtypes.GenesisVestingType{fairdropVestingType, teamdropVestingType, earlyBirdRoundVestingType, publicRoundVestingType, advisorsPool, testVestingPool}
}

type MainnetMigrationChainingSetupSuite struct {
	BaseSetupSuite
}

func TestMainnetMigrationChainingSuite(t *testing.T) {
	suite.Run(t, new(MainnetMigrationChainingSetupSuite))
}

func (s *MainnetMigrationChainingSetupSuite) SetupSuite() {
	bytes, err := os.ReadFile("./resources/mainnet-v1.1.0-migration-app-state.json")
	if err != nil {
		panic(err)
	}
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, true, false, bytes)
}

func (s *MainnetMigrationChainingSetupSuite) TestMainnetVestingsMigrationWhenChainingMigrations() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	campaigns := node.QueryCampaigns()
	s.Equal(4, len(campaigns))

	userEntries := node.QueryUserEntries()
	s.Equal(107404, len(userEntries))

	vestingTypes := node.QueryVestingTypes()
	s.Equal(6, len(vestingTypes))
	s.ElementsMatch(createMainnetVestingTypes(), vestingTypes)
	balances, err := node.QueryBalances(v200.AirdropModuleAccountAddress)
	s.NoError(err)
	s.True(balances.IsEqual(sdk.NewCoins()))

	teamdropAccount := node.QueryAccount(v200.TeamdropVestingAccount)
	s.NotNil(teamdropAccount)
	teamdropVestingPools := node.QueryVestingPoolsInfo(v200.TeamdropVestingAccount)
	s.Equal(1, len(teamdropVestingPools))
	s.Equal(teamdropVestingPools[0].VestingType, "Teamdrop")

	teamdropCampaign := node.QueryCampaign("0")
	s.NotNil(teamdropCampaign)

	s.Equal(teamdropVestingPools[0].Reservations[0].Amount, teamdropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
	s.Equal(teamdropVestingPools[0].CurrentlyLocked, "8899990000000")
	s.Equal(teamdropVestingPools[0].InitiallyLocked.Amount, sdk.NewInt(8899990000000))
	s.Equal(teamdropVestingPools[0].SentAmount, math.ZeroInt().String())
	s.Equal(teamdropCampaign.CampaignCurrentAmount, teamdropCampaign.CampaignTotalAmount)

	santadropCampaign := node.QueryCampaign("2")
	s.NotNil(santadropCampaign)
	s.Equal(santadropCampaign.CampaignCurrentAmount, santadropCampaign.CampaignTotalAmount)

	stakedropCampaign := node.QueryCampaign("1")
	s.NotNil(stakedropCampaign)
	s.Equal(stakedropCampaign.CampaignCurrentAmount, stakedropCampaign.CampaignTotalAmount)

	gleamdropCampaign := node.QueryCampaign("3")
	s.NotNil(gleamdropCampaign)
	s.Equal(gleamdropCampaign.CampaignCurrentAmount, gleamdropCampaign.CampaignTotalAmount)

	fairdropVestingPools := node.QueryVestingPoolsInfo(v200.NewAirdropVestingPoolOwner)
	for _, vestingPoolInfo := range fairdropVestingPools {
		if vestingPoolInfo.Name == "Fairdrop" {
			s.Equal(vestingPoolInfo.VestingType, "Fairdrop")
			s.Equal(vestingPoolInfo.InitiallyLocked.Amount, math.NewInt(20000000000000))
			s.Equal(vestingPoolInfo.SentAmount, math.ZeroInt().String())
			s.Equal(vestingPoolInfo.Reservations[0].Amount, stakedropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[1].Amount, santadropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[2].Amount, gleamdropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
		}
	}

}
