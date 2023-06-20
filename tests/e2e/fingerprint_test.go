package e2e

import (
	"cosmossdk.io/math"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FingerprintSetupSuite struct {
	BaseSetupSuite
}

func TestFingerprintSuite(t *testing.T) {
	suite.Run(t, new(FingerprintSetupSuite))
}

func (s *FingerprintSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(true, false, false)
}

func (s *FingerprintSetupSuite) TestCreatePayloadLink() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := utils.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)

	randPayloadHash := utils.RandStringOfLength(32)
	node.CreatePayloadLink(randPayloadHash, creatorWalletName)

	balanceAfter, err := node.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceBeforeAmount.Sub(vestingAmount), balanceAfter.AmountOf(appparams.MicroC4eUnit))

	vestingPools := node.QueryVestingPoolsInfo(creatorAddress)
	s.Equal(1, len(vestingPools))

	sendToVestingAccAmount := vestingAmount.Quo(math.NewInt(2))
	node.SendToVestingAccount(creatorAddress, receiverAddress, randVestingPoolName, sendToVestingAccAmount.String(), "false")
	vestingPools = node.QueryVestingPoolsInfo(creatorAddress)
	s.Equal(sendToVestingAccAmount.String(), vestingPools[0].SentAmount)
}
