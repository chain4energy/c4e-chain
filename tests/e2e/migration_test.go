package e2e

import (
	"cosmossdk.io/math"
	v200 "github.com/chain4energy/c4e-chain/app/upgrades/v200"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
	"time"
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
	s.Equal(6, len(campaigns))
	s.Equal(campaigns, createMainetCampaigns())

	userEntries := node.QueryUserEntries()
	s.Equal(107938, len(userEntries))

	vestingTypes := node.QueryVestingTypes()
	s.Equal(6, len(vestingTypes))
	s.ElementsMatch(createMainnetVestingTypes(), vestingTypes)
	balances, err := node.QueryBalances(v200.AirdropModuleAccountAddress)
	s.NoError(err)
	s.True(balances.IsEqual(sdk.NewCoins()))

	moondropAccount := node.QueryAccount(v200.MoondropVestingAccount)
	s.NotNil(moondropAccount)
	moondropVestingPools := node.QueryVestingPoolsInfo(v200.MoondropVestingAccount)
	s.Equal(1, len(moondropVestingPools))
	s.Equal(moondropVestingPools[0].VestingType, "Moondrop")

	moondropCampaign := node.QueryCampaign("0")
	s.NotNil(moondropCampaign)

	s.Equal(moondropVestingPools[0].Reservations[0].Amount, moondropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
	s.Equal(moondropVestingPools[0].CurrentlyLocked, "8899990000000")
	s.Equal(moondropVestingPools[0].InitiallyLocked.Amount, sdk.NewInt(8899990000000))
	s.Equal(moondropVestingPools[0].SentAmount, math.ZeroInt().String())
	s.Equal(moondropCampaign.CampaignCurrentAmount, moondropCampaign.CampaignTotalAmount)

	stakedropCampaign := node.QueryCampaign("1")
	s.NotNil(stakedropCampaign)
	s.Equal(stakedropCampaign.CampaignCurrentAmount, stakedropCampaign.CampaignTotalAmount)

	santadropCampaign := node.QueryCampaign("2")
	s.NotNil(santadropCampaign)
	s.Equal(santadropCampaign.CampaignCurrentAmount, santadropCampaign.CampaignTotalAmount)

	greendropCampaign := node.QueryCampaign("3")
	s.NotNil(greendropCampaign)
	s.Equal(greendropCampaign.CampaignCurrentAmount, greendropCampaign.CampaignTotalAmount)

	zealaydropCampaign := node.QueryCampaign("4")
	s.NotNil(zealaydropCampaign)
	s.Equal(zealaydropCampaign.CampaignCurrentAmount, zealaydropCampaign.CampaignTotalAmount)

	amadropCampaign := node.QueryCampaign("5")
	s.NotNil(amadropCampaign)
	s.Equal(amadropCampaign.CampaignCurrentAmount, amadropCampaign.CampaignTotalAmount)

	fairdropVestingPools := node.QueryVestingPoolsInfo(v200.NewAirdropVestingPoolOwner)
	for _, vestingPoolInfo := range fairdropVestingPools {
		if vestingPoolInfo.Name == "Fairdrop" {
			s.Equal(vestingPoolInfo.VestingType, "Fairdrop")
			s.Equal(vestingPoolInfo.InitiallyLocked.Amount, math.NewInt(20000000000000))
			s.Equal(vestingPoolInfo.SentAmount, math.ZeroInt().String())
			s.Equal(vestingPoolInfo.Reservations[0].Amount, stakedropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[1].Amount, santadropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[2].Amount, greendropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[3].Amount, zealaydropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[4].Amount, amadropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
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
		Free:              sdk.MustNewDecFromStr("0.01"),
	}

	moondropVestingType := cfevestingtypes.GenesisVestingType{
		Name:              "Moondrop",
		Free:              sdk.ZeroDec(),
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
	return []cfevestingtypes.GenesisVestingType{fairdropVestingType, moondropVestingType, earlyBirdRoundVestingType, publicRoundVestingType, advisorsPool, testVestingPool}
}

var (
	airdropStartTime      = time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)
	airdropEndTime        = time.Date(2031, 1, 1, 0, 0, 0, 0, time.UTC)
	airdropLockupPeriod   = 183 * 24 * time.Hour
	airdropVestingPeriod  = 91 * 24 * time.Hour
	moondropLockupPeriod  = 730 * 24 * time.Hour
	moondropVestingPeriod = 730 * 24 * time.Hour
)

func createMainetCampaigns() []cfeclaimtypes.Campaign {
	moondropCampaign := cfeclaimtypes.Campaign{
		Id:                     0,
		Owner:                  "c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8",
		Name:                   "Moon Drop",
		Description:            "",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  true,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.ZeroDec(),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           moondropLockupPeriod,
		VestingPeriod:          moondropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(1813750000000))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(1813750000000))),
		VestingPoolName:        "Moondrop",
	}

	stakedropCampaign := cfeclaimtypes.Campaign{
		Id:                     1,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "Stake Drop",
		Description:            "Stake Drop is the airdrop aimed to spread knowledge about the C4E ecosystem among the Cosmos $ATOM stakers community. The airdrop snapshot has been taken on September 28th, 2022 at 9:30 PM UTC (during the ATOM 2.0 roadmap announcement at the Cosmoverse Conference.",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.MustNewDecFromStr("0.01"),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           airdropLockupPeriod,
		VestingPeriod:          airdropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(8999999989680))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(8999999989680))),
		VestingPoolName:        "Fairdrop",
	}

	santadropCampaign := cfeclaimtypes.Campaign{
		Id:                     2,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "Santa Drop",
		Description:            "Santa Drop prize pool for was 10.000 C4E Tokens, with 10 lucky winners getting 1000 tokens per each. The participants had to complete the tasks to get a chance to be among lucky winners.",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.MustNewDecFromStr("0.01"),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           airdropLockupPeriod,
		VestingPeriod:          airdropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(10000000000))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(10000000000))),
		VestingPoolName:        "Fairdrop",
	}

	greendropCampaign := cfeclaimtypes.Campaign{
		Id:                     3,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "Green Drop",
		Description:            "It was the first airdrop competition aimed at spreading knowledge about the C4E ecosystem. The Prize Pool was 1.000.000 C4E tokens and what is best — all the participants who completed the tasks are eligible for the c4e tokens from it!",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.MustNewDecFromStr("0.01"),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           airdropLockupPeriod,
		VestingPeriod:          airdropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(996647490000))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(996647490000))),
		VestingPoolName:        "Fairdrop",
	}

	zealydropCampaign := cfeclaimtypes.Campaign{
		Id:                     4,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "Incentived Testnet I",
		Description:            "Incentivized Testnet Zealy campaign, is innovative approach designed to foster engagement and bolster network security. Community members are rewarded for participating in testnet and marketing tasks, receiving C4E tokens as a result of their contributions.",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.MustNewDecFromStr("0.01"),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           airdropLockupPeriod,
		VestingPeriod:          airdropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(99695340916))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(99695340916))),
		VestingPoolName:        "Fairdrop",
	}

	amadropCampaign := cfeclaimtypes.Campaign{
		Id:                     5,
		Owner:                  "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
		Name:                   "AMA Drop",
		Description:            "Have you been active at our AMA sessions and won C4E prizes? This Drop belongs to you.",
		CampaignType:           cfeclaimtypes.VestingPoolCampaign,
		RemovableClaimRecords:  false,
		FeegrantAmount:         math.ZeroInt(),
		InitialClaimFreeAmount: math.ZeroInt(),
		Free:                   sdk.MustNewDecFromStr("0.01"),
		Enabled:                false,
		StartTime:              airdropStartTime,
		EndTime:                airdropEndTime,
		LockupPeriod:           airdropLockupPeriod,
		VestingPeriod:          airdropVestingPeriod,
		CampaignCurrentAmount:  sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(2900000000))),
		CampaignTotalAmount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(2900000000))),
		VestingPoolName:        "Fairdrop",
	}
	return []cfeclaimtypes.Campaign{moondropCampaign, stakedropCampaign, santadropCampaign, greendropCampaign, zealydropCampaign, amadropCampaign}
}

type MainnetMigrationChainingSetupSuite struct {
	BaseSetupSuite
}

func TestMainnetMigrationChainingSuite(t *testing.T) {
	suite.Run(t, new(MainnetMigrationChainingSetupSuite))
}

func (s *MainnetMigrationChainingSetupSuite) SetupSuite() {
	//bytes, err := os.ReadFile("./resources/mainnet-v1.1.0-migration-app-state.json")
	//if err != nil {
	//	panic(err)
	//}
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, true, false, nil)
}

func (s *MainnetMigrationChainingSetupSuite) TestMainnetVestingsMigrationWhenChainingMigrations() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	campaigns := node.QueryCampaigns()
	s.Equal(6, len(campaigns))

	userEntries := node.QueryUserEntries()
	s.Equal(107938, len(userEntries))

	vestingTypes := node.QueryVestingTypes()
	s.Equal(6, len(vestingTypes))
	s.ElementsMatch(createMainnetVestingTypes(), vestingTypes)
	balances, err := node.QueryBalances(v200.AirdropModuleAccountAddress)
	s.NoError(err)
	s.True(balances.IsEqual(sdk.NewCoins()))

	moondropAccount := node.QueryAccount(v200.MoondropVestingAccount)
	s.NotNil(moondropAccount)
	moondropVestingPools := node.QueryVestingPoolsInfo(v200.MoondropVestingAccount)
	s.Equal(1, len(moondropVestingPools))
	s.Equal(moondropVestingPools[0].VestingType, "Moondrop")

	moondropCampaign := node.QueryCampaign("0")
	s.NotNil(moondropCampaign)

	s.Equal(moondropVestingPools[0].Reservations[0].Amount, moondropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
	s.Equal(moondropVestingPools[0].CurrentlyLocked, "8899990000000")
	s.Equal(moondropVestingPools[0].InitiallyLocked.Amount, sdk.NewInt(8899990000000))
	s.Equal(moondropVestingPools[0].SentAmount, math.ZeroInt().String())
	s.Equal(moondropCampaign.CampaignCurrentAmount, moondropCampaign.CampaignTotalAmount)

	stakedropCampaign := node.QueryCampaign("1")
	s.NotNil(stakedropCampaign)
	s.Equal(stakedropCampaign.CampaignCurrentAmount, stakedropCampaign.CampaignTotalAmount)

	santadropCampaign := node.QueryCampaign("2")
	s.NotNil(santadropCampaign)
	s.Equal(santadropCampaign.CampaignCurrentAmount, santadropCampaign.CampaignTotalAmount)

	greendropCampaign := node.QueryCampaign("3")
	s.NotNil(greendropCampaign)
	s.Equal(greendropCampaign.CampaignCurrentAmount, greendropCampaign.CampaignTotalAmount)

	zealaydropCampaign := node.QueryCampaign("4")
	s.NotNil(zealaydropCampaign)
	s.Equal(zealaydropCampaign.CampaignCurrentAmount, zealaydropCampaign.CampaignTotalAmount)

	amadropCampaign := node.QueryCampaign("5")
	s.NotNil(amadropCampaign)
	s.Equal(amadropCampaign.CampaignCurrentAmount, amadropCampaign.CampaignTotalAmount)

	fairdropVestingPools := node.QueryVestingPoolsInfo(v200.NewAirdropVestingPoolOwner)
	for _, vestingPoolInfo := range fairdropVestingPools {
		if vestingPoolInfo.Name == "Fairdrop" {
			s.Equal(vestingPoolInfo.VestingType, "Fairdrop")
			s.Equal(vestingPoolInfo.InitiallyLocked.Amount, math.NewInt(20000000000000))
			s.Equal(vestingPoolInfo.SentAmount, math.ZeroInt().String())
			s.Equal(vestingPoolInfo.Reservations[0].Amount, stakedropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[1].Amount, santadropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[2].Amount, greendropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[3].Amount, zealaydropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
			s.Equal(vestingPoolInfo.Reservations[4].Amount, amadropCampaign.CampaignCurrentAmount.AmountOf(testenv.DefaultTestDenom))
		}
	}
}
