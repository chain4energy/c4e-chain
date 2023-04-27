package e2e

import (
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
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

func (s *ClaimSetupSuite) TestAirdropCampaign() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	receiverWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	receiverAddress := node.CreateWallet(receiverWalletName)

	vestingTypes := node.QueryVestingTypes()
	s.Greater(len(vestingTypes), 0)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, sdk.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	balanceBeforeAmount := balanceBefore.AmountOf(appparams.CoinDenom)
	vestingAmount := balanceBeforeAmount.Quo(sdk.NewInt(4))
	randVestingPoolName := helpers.RandStringOfLength(5)
	node.CreateVestingPool(randVestingPoolName, vestingAmount.String(), (10 * time.Minute).String(), vestingTypes[0].Name, creatorWalletName)

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
