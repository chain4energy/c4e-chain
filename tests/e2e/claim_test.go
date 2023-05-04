package e2e

import (
	"cosmossdk.io/math"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"math/rand"
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
	s.BaseSetupSuite.SetupSuite(false, false)
}

func (s *ClaimSetupSuite) TestDefaultCampaign() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	receiverWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	receiverAddress := node.CreateWallet(receiverWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, sdk.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	balanceBeforeAmount := balanceBefore.AmountOf(appparams.CoinDenom)
	vestingAmount := balanceBeforeAmount.Quo(sdk.NewInt(4))

	startTime := time.Now()
	endTime := startTime.Add(time.Duration(helpers.RandIntBetween(r, 1000000, 10000000)))
	lockupPeriod := time.Duration(helpers.RandIntBetween(r, 1000000, 10000000))
	vestingPeriod := time.Duration(helpers.RandIntBetween(r, 1000000, 10000000))
	randVestingPoolName := helpers.RandStringOfLength(10)
	randVestingDescription := helpers.RandStringOfLength(10)
	feegrantAmount := math.ZeroInt().String()
	inititalClaimFreeAmount := math.ZeroInt().String()
	campaignType := types.DefaultCampaign
	vestingPoolName := ""
	node.CreateCampaign(randVestingPoolName, randVestingDescription, campaignType.String(), feegrantAmount, inititalClaimFreeAmount,
		startTime.String(), endTime.String(), lockupPeriod.String(), vestingPeriod.String(), vestingPoolName, creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBeforeAmount.Sub(vestingAmount), balanceAfter.AmountOf(appparams.CoinDenom))

	vestingPools := node.QueryVestingPools(creatorAddress)
	s.Equal(1, len(vestingPools))

	sendToVestingAccAmount := vestingAmount.Quo(sdk.NewInt(2))
	node.SendToVestingAccount(creatorAddress, receiverAddress, randVestingPoolName, sendToVestingAccAmount.String(), "false")
	vestingPools = node.QueryVestingPools(creatorAddress)
	s.Equal(sendToVestingAccAmount.String(), vestingPools[0].SentAmount)
}
