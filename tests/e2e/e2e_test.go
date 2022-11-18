package e2e

import (
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	cfedistributormoduletypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"io"
	"os"
	"time"
)

// TestSuperfluidVoting tests that superfluid voting is functioning as expected.
// It does so by doing the following:
// - creating a pool
// - attempting to submit a proposal to enable superfluid voting in that pool
// - voting yes on the proposal from the validator wallet
// - voting no on the proposal from the delegator wallet
// - ensuring that delegator's wallet overwrites the validator's vote

// Copy a file from A to B with io.Copy
func copyFile(a, b string) error {
	source, err := os.Open(a)
	if err != nil {
		return err
	}
	defer source.Close()
	destination, err := os.Create(b)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}
	return nil
}

//func (s *IntegrationTestSuite) TestStateSync() {
//	if s.skipStateSync {
//		s.T().Skip()
//	}
//
//	chainA := s.configurer.GetChainConfig(0)
//	runningNode, err := chainA.GetDefaultNode()
//	s.Require().NoError(err)
//
//	persistentPeers := chainA.GetPersistentPeers()
//
//	stateSyncHostPort := fmt.Sprintf("%s:26657", runningNode.Name)
//	stateSyncRPCServers := []string{stateSyncHostPort, stateSyncHostPort}
//
//	// get trust height and trust hash.
//	trustHeight, err := runningNode.QueryCurrentHeight()
//	s.Require().NoError(err)
//
//	trustHash, err := runningNode.QueryHashFromBlock(trustHeight)
//	s.Require().NoError(err)
//
//	stateSynchingNodeConfig := &initialization.NodeConfig{
//		Name:               "state-sync",
//		Pruning:            "default",
//		PruningKeepRecent:  "0",
//		PruningInterval:    "0",
//		SnapshotInterval:   1500,
//		SnapshotKeepRecent: 2,
//	}
//
//	tempDir, err := os.MkdirTemp("", "osmosis-e2e-statesync-")
//	s.Require().NoError(err)
//
//	// configure genesis and config files for the state-synchin node.
//	nodeInit, err := initialization.InitSingleNode(
//		chainA.Id,
//		tempDir,
//		filepath.Join(runningNode.ConfigDir, "config", "genesis.json"),
//		stateSynchingNodeConfig,
//		time.Duration(chainA.VotingPeriod),
//		// time.Duration(chainA.ExpeditedVotingPeriod),
//		trustHeight,
//		trustHash,
//		stateSyncRPCServers,
//		persistentPeers,
//	)
//	s.Require().NoError(err)
//
//	stateSynchingNode := chainA.CreateNode(nodeInit)
//
//	// ensure that the running node has snapshots at a height > trustHeight.
//	hasSnapshotsAvailable := func(syncInfo coretypes.SyncInfo) bool {
//		snapshotHeight := runningNode.SnapshotInterval
//		if uint64(syncInfo.LatestBlockHeight) < snapshotHeight {
//			s.T().Logf("snapshot height is not reached yet, current (%d), need (%d)", syncInfo.LatestBlockHeight, snapshotHeight)
//			return false
//		}
//
//		snapshots, err := runningNode.QueryListSnapshots()
//		s.Require().NoError(err)
//
//		for _, snapshot := range snapshots {
//			if snapshot.Height > uint64(trustHeight) {
//				s.T().Log("found state sync snapshot after trust height")
//				return true
//			}
//		}
//		s.T().Log("state sync snashot after trust height is not found")
//		return false
//	}
//	runningNode.WaitUntil(hasSnapshotsAvailable)
//
//	// start the state synchin node.
//	err = stateSynchingNode.Run()
//	s.Require().NoError(err)
//
//	// ensure that the state synching node cathes up to the running node.
//	s.Require().Eventually(func() bool {
//		stateSyncNodeHeight, err := stateSynchingNode.QueryCurrentHeight()
//		s.Require().NoError(err)
//		runningNodeHeight, err := runningNode.QueryCurrentHeight()
//		s.Require().NoError(err)
//		return stateSyncNodeHeight == runningNodeHeight
//	},
//		3*time.Minute,
//		500*time.Millisecond,
//	)
//
//	// stop the state synching node.
//	err = chainA.RemoveNode(stateSynchingNode.Name)
//	s.Require().NoError(err)
//}

func (s *IntegrationTestSuite) TestCfedistributorParamsProposal() {
	chainA := s.configurer.GetChainConfig(0)

	node, err := chainA.GetDefaultNode()
	proposal, _ := os.ReadFile("./scripts/update-subdistributors.json")
	//var result map[string]interface{}
	//_ = json.Unmarshal(proposal, result)
	//newSubdistributors := []cfedistributormoduletypes.SubDistributor{
	//	{
	//		Name: "New subdistributor",
	//		Sources: []*cfedistributormoduletypes.Account{
	//
	//		},
	//		Destinations: cfedistributormoduletypes.Destinations{
	//
	//		},
	//	},
	//}
	//proposal := paramsutils.ParamChangeProposalJSON{
	//	Title:       "Param Change",
	//	Description: "Changing the rate limit contract param",
	//	Changes: paramsutils.ParamChangesJSON{
	//		paramsutils.ParamChangeJSON{
	//			Subspace: cfedistributormoduletypes.ModuleName,
	//			Key:      "SubDistributors",
	//			Value:    ,
	//		},
	//	},
	//	Deposit: "625000000uosmo",
	//}
	s.NoError(err)

	node.SubmitParamChangeProposal(string(proposal), initialization.ValidatorWalletName)
	chainA.LatestProposalNumber += 1

	for _, n := range chainA.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainA.LatestProposalNumber)
	}

	// The value is returned as a string, so we have to unmarshal twice
	type Params struct {
		Key      string `json:"key"`
		Subspace string `json:"subspace"`
		Value    string `json:"value"`
	}

	s.Eventually(
		func() bool {
			var params Params
			node.QueryParams(cfedistributormoduletypes.ModuleName, "SubDistributors", &params)
			return params.Value != ""
		},
		1*time.Minute,
		10*time.Millisecond,
		"Osmosis node failed to retrieve params",
	)

}
