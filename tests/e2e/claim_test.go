package e2e

import (
	"cosmossdk.io/math"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/tests/e2e/util"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	cfeclaimcli "github.com/chain4energy/c4e-chain/x/cfeclaim/client/cli"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingmoduletypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type ClaimSetupSuite struct {
	BaseSetupSuite
}

func TestClaimSuite(t *testing.T) {
	suite.Run(t, new(ClaimSetupSuite))
}

func (s *ClaimSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(false, false, false)
}

func (s *ClaimSetupSuite) TestDefaultCampaign() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	receiverWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	receiverAddress := node.CreateWallet(receiverWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)

	s.NoError(err)

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	startTime := time.Now().Add(time.Second * 30)
	endTime := startTime.Add(helpers.RandDurationBetween(r, 40, 45))
	lockupPeriod := helpers.RandDurationBetween(r, 10000, 10000000)
	vestingPeriod := helpers.RandDurationBetween(r, 10000, 10000000)
	randName := helpers.RandStringOfLength(10)
	randDescription := helpers.RandStringOfLength(10)
	feegrantAmount := math.NewInt(2500000).String()
	inititalClaimFreeAmount := math.ZeroInt().String()
	campaignType := types.DefaultCampaign
	vestingPoolName := ""
	node.CreateCampaign(randName, randDescription, campaignType.String(), feegrantAmount, inititalClaimFreeAmount,
		startTime.Format(cfeclaimcli.TimeLayout), endTime.Format(cfeclaimcli.TimeLayout), lockupPeriod.String(), vestingPeriod.String(), vestingPoolName, creatorWalletName)

	campaigns := node.QueryCampaigns()
	campaignId := len(campaigns) - 1

	randMissionName := helpers.RandStringOfLength(10)
	randMissionDescription := helpers.RandStringOfLength(10)
	node.AddMission(strconv.Itoa(campaignId), randMissionName, randMissionDescription, types.MissionClaim.String(), "0.5", "", creatorWalletName)

	node.EnableCampaign(strconv.Itoa(campaignId), "", "", creatorWalletName)

	userEntires := []types.ClaimRecord{
		{
			Address: receiverAddress,
			Amount:  sdk.NewCoins(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance/4))),
		},
	}
	proposalJSON, err := util.NewClaimRecordsListJson(userEntires)
	s.NoError(err)
	node.AddClaimRecords(strconv.Itoa(campaignId), proposalJSON, creatorWalletName)

	node.ClaimInitialMission(strconv.Itoa(campaignId), "", receiverWalletName)
	node.ClaimMission(strconv.Itoa(campaignId), "1", receiverWalletName)

	node.CloseCampaign(strconv.Itoa(campaignId), creatorWalletName)
}

func (s *ClaimSetupSuite) TestDynamicCampaign() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	receiverWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	receiverAddress := node.CreateWallet(receiverWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)

	s.NoError(err)

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	startTime := time.Now().Add(time.Second * 30)
	endTime := startTime.Add(helpers.RandDurationBetween(r, 40, 45))
	lockupPeriod := helpers.RandDurationBetween(r, 10000, 10000000)
	vestingPeriod := helpers.RandDurationBetween(r, 10000, 10000000)
	randName := helpers.RandStringOfLength(10)
	randDescription := helpers.RandStringOfLength(10)
	feegrantAmount := math.NewInt(2500000).String()
	inititalClaimFreeAmount := math.ZeroInt().String()
	campaignType := types.DefaultCampaign
	vestingPoolName := ""
	node.CreateCampaign(randName, randDescription, campaignType.String(), feegrantAmount, inititalClaimFreeAmount,
		startTime.Format(cfeclaimcli.TimeLayout), endTime.Format(cfeclaimcli.TimeLayout), lockupPeriod.String(), vestingPeriod.String(), vestingPoolName, creatorWalletName)

	campaigns := node.QueryCampaigns()
	campaignId := len(campaigns) - 1

	randMissionName := helpers.RandStringOfLength(10)
	randMissionDescription := helpers.RandStringOfLength(10)
	node.AddMission(strconv.Itoa(campaignId), randMissionName, randMissionDescription, types.MissionClaim.String(), "0.5", "", creatorWalletName)

	node.EnableCampaign(strconv.Itoa(campaignId), "", "", creatorWalletName)

	userEntires := []types.ClaimRecord{
		{
			Address: receiverAddress,
			Amount:  sdk.NewCoins(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance/4))),
		},
	}
	proposalJSON, err := util.NewClaimRecordsListJson(userEntires)
	s.NoError(err)
	node.AddClaimRecords(strconv.Itoa(campaignId), proposalJSON, creatorWalletName)

	node.ClaimInitialMission(strconv.Itoa(campaignId), "", receiverWalletName)
	node.ClaimMission(strconv.Itoa(campaignId), "1", receiverWalletName)

	node.CloseCampaign(strconv.Itoa(campaignId), creatorWalletName)
}

func (s *ClaimSetupSuite) TestVestingPoolCampaign() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	receiverWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	receiverAddress := node.CreateWallet(receiverWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	balanceBeforeAmount := balanceBefore.AmountOf(appparams.CoinDenom)
	vestingAmount := balanceBeforeAmount.Quo(math.NewInt(4))
	randVestingPoolName := helpers.RandStringOfLength(5)
	vestingPoolDuration := 10 * time.Second
	vestingTypes := node.QueryVestingTypes()
	s.Greater(len(vestingTypes), 0)
	vestingType := vestingTypes[0]
	node.CreateVestingPool(randVestingPoolName, vestingAmount.String(), vestingPoolDuration.String(), vestingType.Name, creatorWalletName)

	startTime := time.Now().Add(time.Second * 30)
	endTime := startTime.Add(helpers.RandDurationBetween(r, 40, 45))
	randName := helpers.RandStringOfLength(10)
	randDescription := helpers.RandStringOfLength(10)
	feegrantAmount := math.NewInt(2500000).String()
	inititalClaimFreeAmount := math.ZeroInt().String()
	campaignType := types.VestingPoolCampaign
	lockupDuration, err := cfevestingmoduletypes.DurationFromUnits(cfevestingmoduletypes.PeriodUnit(vestingType.LockupPeriodUnit), vestingType.LockupPeriod)
	s.NoError(err)
	vestingDuration, err := cfevestingmoduletypes.DurationFromUnits(cfevestingmoduletypes.PeriodUnit(vestingType.VestingPeriodUnit), vestingType.VestingPeriod)
	s.NoError(err)
	node.CreateCampaign(randName, randDescription, campaignType.String(), feegrantAmount, inititalClaimFreeAmount,
		startTime.Format(cfeclaimcli.TimeLayout), endTime.Format(cfeclaimcli.TimeLayout), lockupDuration.String(), vestingDuration.String(), randVestingPoolName, creatorWalletName)

	campaigns := node.QueryCampaigns()
	campaignId := len(campaigns) - 1

	randMissionName := helpers.RandStringOfLength(10)
	randMissionDescription := helpers.RandStringOfLength(10)
	node.AddMission(strconv.Itoa(campaignId), randMissionName, randMissionDescription, types.MissionClaim.String(), "0.5", "", creatorWalletName)

	node.EnableCampaign(strconv.Itoa(campaignId), "", "", creatorWalletName)

	userEntires := []types.ClaimRecord{
		{
			Address: receiverAddress,
			Amount:  sdk.NewCoins(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance/4))),
		},
	}
	proposalJSON, err := util.NewClaimRecordsListJson(userEntires)
	s.NoError(err)
	node.AddClaimRecords(strconv.Itoa(campaignId), proposalJSON, creatorWalletName)

	node.ClaimInitialMission(strconv.Itoa(campaignId), "", receiverWalletName)
	node.ClaimMission(strconv.Itoa(campaignId), "1", receiverWalletName)

	node.CloseCampaign(strconv.Itoa(campaignId), creatorWalletName)
}

// TODO: add verifications and more options (probably when adding manual E2E tests)
