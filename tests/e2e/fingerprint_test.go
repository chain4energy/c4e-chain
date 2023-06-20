package e2e

import (
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
	s.BaseSetupSuite.SetupSuite(false, false, false)
}

func (s *FingerprintSetupSuite) TestCreatePayloadLink() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := utils.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	node.BankSendBaseBalanceFromNode(creatorAddress)
	randPayloadHash := utils.RandStringOfLength(32)
	node.CreatePayloadLink(randPayloadHash, creatorWalletName)
}
