package e2e

//func (s *IntegrationTestSuite) TestIBCTokenTransfer() {
//	chainA := s.chainConfigs[0]
//	chainB := s.chainConfigs[1]
//	// compare coins of receiver pre and post IBC send
//	// diff should only be the amount sent
//	s.sendIBC(chainA, chainB, chainB.validators[0].validator.PublicAddress, initialization.OsmoToken)
//}

func (s *IntegrationTestSuite) TestSubmitTextProposal() {
	chainA := s.chainConfigs[0]
	// compare coins of receiver pre and post IBC send
	// diff should only be the amount sent
	s.submitTextProposal(chainA, "Proposal 1")
}
