package e2e

import (
	appparams "github.com/chain4energy/c4e-chain/tests/e2e/encoding/params"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type VestingSetupSuite struct {
	BaseSetupSuite
}

func TestIntegrationVestingSuite(t *testing.T) {
	suite.Run(t, new(VestingSetupSuite))
}

func (s *VestingSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(true, false)
}

func (s *VestingSetupSuite) TestSendToVestingAccount() {
	const (
		creatorWalletName  = "user-1"
		receiverWalletName = "user-2"
		baseBalance        = 10000000
	)
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorAddress := node.CreateWallet(creatorWalletName)
	vestingTypes := node.QueryVestingTypes()

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, sdk.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balance, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	balanceAmount := balance.AmountOf(appparams.CoinDenom)
	vestingAmount := balanceAmount.Quo(sdk.NewInt(4))
	randVestingPoolName := helpers.RandStringOfLength(5)
	node.CreateVestingPool(randVestingPoolName, vestingAmount.String(), (10 * time.Minute).String(), vestingTypes[0].Name, creatorWalletName)

	newBalance, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceAmount.Sub(vestingAmount), newBalance.AmountOf(appparams.CoinDenom))

	vestingPools := node.QueryVestingPools(creatorAddress)
	s.Equal(1, len(vestingPools))
	receiverAddress := node.CreateWallet(receiverWalletName)
	sendToVestingAccAmount := vestingAmount.Quo(sdk.NewInt(2))
	node.SendToVestingAccount(creatorAddress, receiverAddress, randVestingPoolName, sendToVestingAccAmount.String(), "false")
	vestingPools = node.QueryVestingPools(creatorAddress)
	s.Equal(sendToVestingAccAmount.String(), vestingPools[0].SentAmount)
}

func (s *VestingSetupSuite) TestWithdrawAllAvaliable() {
	const (
		creatorWalletName  = "user-3"
		receiverWalletName = "user-4"
		baseBalance        = 10000000
	)
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorAddress := node.CreateWallet(creatorWalletName)
	vestingTypes := node.QueryVestingTypes()

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, sdk.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balance, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	balanceAmount := balance.AmountOf(appparams.CoinDenom)
	vestingAmount := balanceAmount.Quo(sdk.NewInt(4))
	randVestingPoolName := helpers.RandStringOfLength(5)
	node.CreateVestingPool(randVestingPoolName, vestingAmount.String(), (10 * time.Second).String(), vestingTypes[0].Name, creatorWalletName)
	vestingPools := node.QueryVestingPools(creatorAddress)
	s.Equal(vestingPools[0].Withdrawable, "0")
	s.Equal(vestingPools[0].CurrentlyLocked, vestingAmount.String())
	s.Eventually(
		func() bool {
			vestingPools := node.QueryVestingPools(creatorAddress)
			if vestingAmount.String() == vestingPools[0].Withdrawable {
				node.WithdrawAllAvaliable(creatorAddress)
				vestingPools = node.QueryVestingPools(creatorAddress)
				return s.True(vestingPools[0].Withdrawable == "0")
			}
			return false
		},
		time.Minute,
		10*time.Second,
		"C4e node failed to validate params",
	)
	newBalance, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceAmount, newBalance.AmountOf(appparams.CoinDenom))
}
