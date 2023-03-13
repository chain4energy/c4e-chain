package e2e

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	"github.com/stretchr/testify/suite"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type IbcSetupSuite struct {
	BaseSetupSuite
}

func TestIbcSuite(t *testing.T) {
	suite.Run(t, new(IbcSetupSuite))
}

func (s *IbcSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(true, true)
}

func (s *IbcSetupSuite) TestIbcTokenTransfer() {
	chainA := s.configurer.GetChainConfig(0)
	chainB := s.configurer.GetChainConfig(1)
	chainA.SendIBC(chainB, chainB.NodeConfigs[0].PublicAddress, initialization.C4eToken)
	chainA.SendIBC(chainB, chainB.NodeConfigs[0].PublicAddress, initialization.C4eToken)
}

func (s *IbcSetupSuite) TestStateSync() {
	chainA := s.configurer.GetChainConfig(0)
	runningNode, err := chainA.GetDefaultNode()
	s.Require().NoError(err)

	persistentPeers := chainA.GetPersistentPeers()

	stateSyncHostPort := fmt.Sprintf("%s:26657", runningNode.Name)
	stateSyncRPCServers := []string{stateSyncHostPort, stateSyncHostPort}

	trustHeight, err := runningNode.QueryCurrentHeight()
	s.Require().NoError(err)

	trustHash, err := runningNode.QueryHashFromBlock(trustHeight)
	s.Require().NoError(err)

	stateSynchingNodeConfig := &initialization.NodeConfig{
		Name:               "state-sync",
		Pruning:            "default",
		PruningKeepRecent:  "0",
		PruningInterval:    "0",
		SnapshotInterval:   1500,
		SnapshotKeepRecent: 2,
	}

	tempDir, err := os.MkdirTemp("", "c4e-e2e-statesync-")
	s.Require().NoError(err)

	nodeInit, err := initialization.InitSingleNode(
		chainA.Id,
		tempDir,
		filepath.Join(runningNode.ConfigDir, "config", "genesis.json"),
		stateSynchingNodeConfig,
		trustHeight,
		trustHash,
		stateSyncRPCServers,
		persistentPeers,
	)
	s.Require().NoError(err)

	stateSynchingNode := chainA.CreateNode(nodeInit)

	hasSnapshotsAvailable := func(syncInfo coretypes.SyncInfo) bool {
		snapshotHeight := runningNode.SnapshotInterval
		if uint64(syncInfo.LatestBlockHeight) < snapshotHeight {
			s.T().Logf("snapshot height is not reached yet, current (%d), need (%d)", syncInfo.LatestBlockHeight, snapshotHeight)
			return false
		}

		snapshots, err := runningNode.QueryListSnapshots()
		s.Require().NoError(err)

		for _, snapshot := range snapshots {
			if snapshot.Height > uint64(trustHeight) {
				s.T().Log("found state sync snapshot after trust height")
				return true
			}
		}
		s.T().Log("state sync snashot after trust height is not found")
		return false
	}
	runningNode.WaitUntil(hasSnapshotsAvailable)

	err = stateSynchingNode.Run()
	s.Require().NoError(err)

	s.Require().Eventually(func() bool {
		stateSyncNodeHeight, err := stateSynchingNode.QueryCurrentHeight()
		s.Require().NoError(err)
		runningNodeHeight, err := runningNode.QueryCurrentHeight()
		s.Require().NoError(err)
		return stateSyncNodeHeight == runningNodeHeight
	},
		3*time.Minute,
		500*time.Millisecond,
	)

	err = chainA.RemoveNode(stateSynchingNode.Name)
	s.Require().NoError(err)
}
