package e2e

import (
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
)

func (s *IntegrationTestSuite) TestIBCTokenTransfer() {
	chainA := s.chainConfigs[0]
	chainB := s.chainConfigs[1]
	// compare coins of receiver pre and post IBC send
	// diff should only be the amount sent
	s.sendIBC(chainA, chainB, chainB.validators[0].validator.PublicAddress, initialization.C4eToken)
}
