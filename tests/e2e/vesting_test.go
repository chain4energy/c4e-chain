package e2e

import (
	appparams "github.com/chain4energy/c4e-chain/tests/e2e/encoding/params"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const (
	baseBalance = 10000000
)

type VestingSetupSuite struct {
	BaseSetupSuite
}

func TestVestingSuite(t *testing.T) {
	suite.Run(t, new(VestingSetupSuite))
}

func (s *VestingSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(false, false)
}

func (s *VestingSetupSuite) TestSendToVestingAccount() {
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

func (s *VestingSetupSuite) TestWithdrawAllAvailable() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	vestingTypes := node.QueryVestingTypes()
	s.Greater(len(vestingTypes), 0)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, sdk.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	balanceBeforeAmount := balanceBefore.AmountOf(appparams.CoinDenom)
	vestingAmount := balanceBeforeAmount.Quo(sdk.NewInt(4))
	randVestingPoolName := helpers.RandStringOfLength(5)
	vestingPoolDuration := 10 * time.Second
	node.CreateVestingPool(randVestingPoolName, vestingAmount.String(), vestingPoolDuration.String(), vestingTypes[0].Name, creatorWalletName)

	vestingPools := node.QueryVestingPools(creatorAddress)
	s.Equal(vestingPools[0].Withdrawable, "0")
	s.Equal(vestingPools[0].CurrentlyLocked, vestingAmount.String())

	s.Eventually(
		func() bool {
			vestingPools := node.QueryVestingPools(creatorAddress)
			if vestingAmount.String() == vestingPools[0].Withdrawable {
				node.WithdrawAllAvailable(creatorAddress)
				vestingPools = node.QueryVestingPools(creatorAddress)
				return s.True(vestingPools[0].Withdrawable == "0")
			}
			return false
		},
		time.Minute,
		vestingPoolDuration,
		"C4e node failed to withdraw all avaliable",
	)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBeforeAmount, balanceAfter.AmountOf(appparams.CoinDenom))
}
