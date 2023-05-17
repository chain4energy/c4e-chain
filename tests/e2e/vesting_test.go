package e2e

import (
	"cosmossdk.io/math"
	"fmt"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/chain"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/stretchr/testify/suite"
	"strconv"
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
	s.BaseSetupSuite.SetupSuite(true, false)
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

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	balanceBeforeAmount := balanceBefore.AmountOf(appparams.CoinDenom)
	vestingAmount := balanceBeforeAmount.Quo(math.NewInt(4))
	randVestingPoolName := helpers.RandStringOfLength(5)
	node.CreateVestingPool(randVestingPoolName, vestingAmount.String(), (10 * time.Minute).String(), vestingTypes[0].Name, creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBeforeAmount.Sub(vestingAmount), balanceAfter.AmountOf(appparams.CoinDenom))

	vestingPools := node.QueryVestingPools(creatorAddress)
	s.Equal(1, len(vestingPools))

	sendToVestingAccAmount := vestingAmount.Quo(math.NewInt(2))
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

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	balanceBeforeAmount := balanceBefore.AmountOf(appparams.CoinDenom)
	vestingAmount := balanceBeforeAmount.Quo(math.NewInt(4))
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

func (s *VestingSetupSuite) TestCreateVestingPool() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	vestingTypes := node.QueryVestingTypes()
	s.Greater(len(vestingTypes), 0)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	balanceBeforeAmount := balanceBefore.AmountOf(appparams.CoinDenom)
	vestingAmount := balanceBeforeAmount.Quo(math.NewInt(4))
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

func (s *VestingSetupSuite) TestCreateVestingAccount() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	newVestingWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	newVestingAccountAddress := node.CreateWallet(newVestingWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	amountToSend := balanceBefore.AmountOf(appparams.CoinDenom).Quo(math.NewInt(4))
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)

	coinToSend := sdk.NewCoin(appparams.CoinDenom, amountToSend)
	node.CreateVestingAccount(newVestingAccountAddress, coinToSend.String(),
		strconv.FormatInt(startTime.Unix(), 10), strconv.FormatInt(endTime.Unix(), 10), creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBefore.AmountOf(appparams.CoinDenom).Sub(amountToSend), balanceAfter.AmountOf(appparams.CoinDenom))
	s.ValidateVestingAccount(node, newVestingAccountAddress, coinToSend)
}

func (s *VestingSetupSuite) ValidateVestingAccount(node *chain.NodeConfig, address string, initiallyLocked sdk.Coin) {
	acc := node.QueryAccount(address)
	vestingAccount, ok := acc.(*vestingtypes.ContinuousVestingAccount)
	s.True(ok)
	s.Equal(vestingAccount.OriginalVesting, sdk.NewCoins(initiallyLocked))
}

func (s *VestingSetupSuite) TestSplitVesting() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	newVestingWalletName := helpers.RandStringOfLength(10)
	splitVestingWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	newVestingAccountAddress := node.CreateWallet(newVestingWalletName)
	splitVestingAccountAddress := node.CreateWallet(splitVestingWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	amountToSend := balanceBefore.AmountOf(appparams.CoinDenom).Quo(math.NewInt(4))
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)
	coinToSend := sdk.NewCoin(appparams.CoinDenom, amountToSend)
	node.CreateVestingAccount(newVestingAccountAddress, coinToSend.String(),
		strconv.FormatInt(startTime.Unix(), 10), strconv.FormatInt(endTime.Unix(), 10), creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBefore.AmountOf(appparams.CoinDenom).Sub(amountToSend), balanceAfter.AmountOf(appparams.CoinDenom))
	s.ValidateVestingAccount(node, newVestingAccountAddress, coinToSend)

	coinToSplit := sdk.NewCoin(appparams.CoinDenom, amountToSend.QuoRaw(2))
	node.SplitVesting(splitVestingAccountAddress, coinToSplit.String(), newVestingAccountAddress)
	s.ValidateVestingAccount(node, splitVestingAccountAddress, coinToSplit)
}

func (s *VestingSetupSuite) TestMoveAvailableVesting() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	newVestingWalletName := helpers.RandStringOfLength(10)
	splitVestingWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	newVestingAccountAddress := node.CreateWallet(newVestingWalletName)
	splitVestingAccountAddress := node.CreateWallet(splitVestingWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	amountToSend := balanceBefore.AmountOf(appparams.CoinDenom).Quo(math.NewInt(4))
	startTime := time.Now().Add(time.Minute)
	endTime := startTime.Add(time.Hour)
	coinToSend := sdk.NewCoin(appparams.CoinDenom, amountToSend)
	node.CreateVestingAccount(newVestingAccountAddress, coinToSend.String(),
		strconv.FormatInt(startTime.Unix(), 10), strconv.FormatInt(endTime.Unix(), 10), creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBefore.AmountOf(appparams.CoinDenom).Sub(amountToSend), balanceAfter.AmountOf(appparams.CoinDenom))
	s.ValidateVestingAccount(node, newVestingAccountAddress, coinToSend)

	node.MoveAvailableVesting(splitVestingAccountAddress, newVestingAccountAddress)
	s.ValidateVestingAccount(node, splitVestingAccountAddress, coinToSend)
}

func (s *VestingSetupSuite) TestMoveAvailableVestingByDenoms() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	newVestingWalletName := helpers.RandStringOfLength(10)
	splitVestingWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	newVestingAccountAddress := node.CreateWallet(newVestingWalletName)
	splitVestingAccountAddress := node.CreateWallet(splitVestingWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	amountToSend := balanceBefore.AmountOf(appparams.CoinDenom).Quo(math.NewInt(4))
	startTime := time.Now().Add(time.Minute)
	endTime := startTime.Add(time.Hour)
	coinToSend := sdk.NewCoin(appparams.CoinDenom, amountToSend)
	node.CreateVestingAccount(newVestingAccountAddress, coinToSend.String(),
		strconv.FormatInt(startTime.Unix(), 10), strconv.FormatInt(endTime.Unix(), 10), creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBefore.AmountOf(appparams.CoinDenom).Sub(amountToSend), balanceAfter.AmountOf(appparams.CoinDenom))
	s.ValidateVestingAccount(node, newVestingAccountAddress, coinToSend)

	node.MoveAvailableVestingByDenoms(splitVestingAccountAddress, appparams.CoinDenom, newVestingAccountAddress)
	s.ValidateVestingAccount(node, splitVestingAccountAddress, coinToSend)
}

func (s *VestingSetupSuite) TestMoveAvailableVestingNoCoinsToMove() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	newVestingWalletName := helpers.RandStringOfLength(10)
	splitVestingWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	newVestingAccountAddress := node.CreateWallet(newVestingWalletName)
	splitVestingAccountAddress := node.CreateWallet(splitVestingWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	amountToSend := balanceBefore.AmountOf(appparams.CoinDenom).Quo(math.NewInt(4))
	startTime := time.Now()
	endTime := startTime.Add(time.Second)
	coinToSend := sdk.NewCoin(appparams.CoinDenom, amountToSend)
	node.CreateVestingAccount(newVestingAccountAddress, coinToSend.String(),
		strconv.FormatInt(startTime.Unix(), 10), strconv.FormatInt(endTime.Unix(), 10), creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBefore.AmountOf(appparams.CoinDenom).Sub(amountToSend), balanceAfter.AmountOf(appparams.CoinDenom))
	s.ValidateVestingAccount(node, newVestingAccountAddress, coinToSend)

	node.MoveAvailableVestingError(splitVestingAccountAddress, newVestingAccountAddress, "move available vesting: split vesting coins - no coins to split : wrong param value")
	node.QueryAccountNotFound(splitVestingAccountAddress)
}

func (s *VestingSetupSuite) TestMoveAvailableVestingByDenomNoCoinsToMove() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	newVestingWalletName := helpers.RandStringOfLength(10)
	splitVestingWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	newVestingAccountAddress := node.CreateWallet(newVestingWalletName)
	splitVestingAccountAddress := node.CreateWallet(splitVestingWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	amountToSend := balanceBefore.AmountOf(appparams.CoinDenom).Quo(math.NewInt(4))
	startTime := time.Now()
	endTime := startTime.Add(time.Second)
	coinToSend := sdk.NewCoin(appparams.CoinDenom, amountToSend)
	node.CreateVestingAccount(newVestingAccountAddress, coinToSend.String(),
		strconv.FormatInt(startTime.Unix(), 10), strconv.FormatInt(endTime.Unix(), 10), creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBefore.AmountOf(appparams.CoinDenom).Sub(amountToSend), balanceAfter.AmountOf(appparams.CoinDenom))
	s.ValidateVestingAccount(node, newVestingAccountAddress, coinToSend)

	node.MoveAvailableVestingByDenomsError(splitVestingAccountAddress, appparams.CoinDenom, newVestingAccountAddress, "move available vesting by denoms: split vesting coins - no coins to split : wrong param value")
	node.QueryAccountNotFound(splitVestingAccountAddress)
}

func (s *VestingSetupSuite) TestSplitVestingWrongAmount() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	newVestingWalletName := helpers.RandStringOfLength(10)
	splitVestingWalletName := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	newVestingAccountAddress := node.CreateWallet(newVestingWalletName)
	splitVestingAccountAddress := node.CreateWallet(splitVestingWalletName)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	amountToSend := balanceBefore.AmountOf(appparams.CoinDenom).Quo(math.NewInt(4))
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)
	coinToSend := sdk.NewCoin(appparams.CoinDenom, amountToSend)
	node.CreateVestingAccount(newVestingAccountAddress, coinToSend.String(),
		strconv.FormatInt(startTime.Unix(), 10), strconv.FormatInt(endTime.Unix(), 10), creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBefore.AmountOf(appparams.CoinDenom).Sub(amountToSend), balanceAfter.AmountOf(appparams.CoinDenom))
	s.ValidateVestingAccount(node, newVestingAccountAddress, coinToSend)

	coinToSplit := sdk.NewCoin(appparams.CoinDenom, amountToSend.AddRaw(1000000))
	node.SplitVestingError(splitVestingAccountAddress, coinToSplit.String(), newVestingAccountAddress,
		fmt.Sprintf("split vesting: split vesting coins: account %s: not enough to unlock.", newVestingAccountAddress))
	node.QueryAccountNotFound(splitVestingAccountAddress)
}

func (s *VestingSetupSuite) TestDoubleSplitVesting() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	newVestingWalletName := helpers.RandStringOfLength(10)
	splitVestingWalletName1 := helpers.RandStringOfLength(10)
	splitVestingWalletName2 := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	newVestingAccountAddress := node.CreateWallet(newVestingWalletName)
	splitVestingAccountAddress1 := node.CreateWallet(splitVestingWalletName1)
	splitVestingAccountAddress2 := node.CreateWallet(splitVestingWalletName2)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	amountToSend := balanceBefore.AmountOf(appparams.CoinDenom).Quo(math.NewInt(4))
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)
	coinToSend := sdk.NewCoin(appparams.CoinDenom, amountToSend)
	node.CreateVestingAccount(newVestingAccountAddress, coinToSend.String(),
		strconv.FormatInt(startTime.Unix(), 10), strconv.FormatInt(endTime.Unix(), 10), creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBefore.AmountOf(appparams.CoinDenom).Sub(amountToSend), balanceAfter.AmountOf(appparams.CoinDenom))
	s.ValidateVestingAccount(node, newVestingAccountAddress, coinToSend)

	coinToSplit := sdk.NewCoin(appparams.CoinDenom, amountToSend.QuoRaw(5))
	node.SplitVesting(splitVestingAccountAddress1, coinToSplit.String(), newVestingAccountAddress)
	s.ValidateVestingAccount(node, splitVestingAccountAddress1, coinToSplit)
	node.SplitVesting(splitVestingAccountAddress2, coinToSplit.String(), newVestingAccountAddress)
	s.ValidateVestingAccount(node, splitVestingAccountAddress2, coinToSplit)
}

func (s *VestingSetupSuite) TestDoubleMoveAvailableVesting() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := helpers.RandStringOfLength(10)
	newVestingWalletName := helpers.RandStringOfLength(10)
	splitVestingWalletName1 := helpers.RandStringOfLength(10)
	splitVestingWalletName2 := helpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	newVestingAccountAddress := node.CreateWallet(newVestingWalletName)
	splitVestingAccountAddress1 := node.CreateWallet(splitVestingWalletName1)
	splitVestingAccountAddress2 := node.CreateWallet(splitVestingWalletName2)

	node.BankSend(sdk.NewCoin(appparams.CoinDenom, math.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	s.NoError(err)

	amountToSend := balanceBefore.AmountOf(appparams.CoinDenom).Quo(math.NewInt(4))
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)
	coinToSend := sdk.NewCoin(appparams.CoinDenom, amountToSend)
	node.CreateVestingAccount(newVestingAccountAddress, coinToSend.String(),
		strconv.FormatInt(startTime.Unix(), 10), strconv.FormatInt(endTime.Unix(), 10), creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBefore.AmountOf(appparams.CoinDenom).Sub(amountToSend), balanceAfter.AmountOf(appparams.CoinDenom))
	s.ValidateVestingAccount(node, newVestingAccountAddress, coinToSend)

	node.MoveAvailableVesting(splitVestingAccountAddress1, newVestingAccountAddress)
	node.MoveAvailableVestingError(splitVestingAccountAddress2, newVestingAccountAddress, "move available vesting: split vesting coins - no coins to split : wrong param value")
	node.QueryAccountNotFound(splitVestingAccountAddress2)
}
