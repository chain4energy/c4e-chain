package e2e

import (
	"encoding/json"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	"os"
)

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

func (s *IntegrationTestSuite) TestDistributorParamsChange() {
	//proposal := paramsutils.ParamChangeProposalJSON{
	//	Title:       "CfeDistributor module params change",
	//	Description: "Change cfedistributor params",
	//	Changes: paramsutils.ParamChangesJSON{
	//		paramsutils.ParamChangeJSON{
	//			Subspace: cfedistributormoduletypes.ModuleName,
	//			Key:      string(cfedistributormoduletypes.KeySubDistributors),
	//			Value:    []byte(fmt.Sprintf(`"%s"`, contracts[0])),
	//		},
	//	},
	//	Deposit: "625000000uosmo",
	//}
	proposal, _ := os.ReadFile("./scripts/update-subdistributors.json")
	proposalJson, err := json.Marshal(proposal)
	SubmitParamChangeProposal(string(proposalJson), initialization.ValidatorWalletName)
	s.NoError(err)
	chainA := s.chainConfigs[0]
	// compare coins of receiver pre and post IBC send
	// diff should only be the amount sent
	s.submitTextProposal(chainA, "Proposal 1")
}
