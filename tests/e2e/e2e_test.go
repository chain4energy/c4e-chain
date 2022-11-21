package e2e

import (
	"encoding/json"
	"fmt"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/config"
	appparams "github.com/chain4energy/c4e-chain/tests/e2e/encoding/params"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfemintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramsutils "github.com/cosmos/cosmos-sdk/x/params/client/utils"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"io"
	"os"
	"path/filepath"
	"time"
)

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

func (s *IntegrationTestSuite) TestStateSync() {
	if s.skipStateSync {
		s.T().Skip()
	}

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

	// start the state synchin node.
	err = stateSynchingNode.Run()
	s.Require().NoError(err)

	// ensure that the state synching node cathes up to the running node.
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

	// stop the state synching node.
	err = chainA.RemoveNode(stateSynchingNode.Name)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestCfedistributorParamsProposal() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()

	newSubDistributors := []cfedistributortypes.SubDistributor{
		{
			Name: "New subdistributor",
			Sources: []*cfedistributortypes.Account{
				{
					Id:   cfedistributortypes.GreenEnergyBoosterCollector,
					Type: cfedistributortypes.MAIN,
				},
			},
			Destinations: cfedistributortypes.Destinations{
				PrimaryShare: cfedistributortypes.Account{
					Id:   cfedistributortypes.ValidatorsRewardsCollector,
					Type: cfedistributortypes.MODULE_ACCOUNT,
				},
				BurnShare: sdk.ZeroDec(),
			},
		},
	}

	newSubDistributorsJSON, err := json.Marshal(newSubDistributors)
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "CfeDistributor module params change",
		Description: "Change cfedistributor params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfedistributortypes.ModuleName,
				Key:      string(cfedistributortypes.KeySubDistributors),
				Value:    newSubDistributorsJSON,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	s.NoError(err)
	node.SubmitParamChangeProposal(string(proposalJSON), initialization.ValidatorWalletName)
	chainA.LatestProposalNumber += 1

	for _, n := range chainA.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainA.LatestProposalNumber)
	}

	type Params struct {
		Key      string `json:"key"`
		Subspace string `json:"subspace"`
		Value    string `json:"value"`
	}

	s.Eventually(
		func() bool {
			var params Params
			node.QueryParams(cfedistributortypes.ModuleName, string(cfedistributortypes.KeySubDistributors), &params)
			return params.Value == string(newSubDistributorsJSON)
		},
		1*time.Minute,
		10*time.Millisecond,
		"C4e node failed to retrieve params",
	)
}

func (s *IntegrationTestSuite) TestCfeminterParamsProposal() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()

	periodEnd := time.Now().Add(10 * time.Hour)
	newMinter := cfemintertypes.Minter{
		Start: time.Now(),
		Periods: []*cfemintertypes.MintingPeriod{
			{
				Position:                1,
				Type:                    cfemintertypes.TIME_LINEAR_MINTER,
				PeriodEnd:               &periodEnd,
				PeriodicReductionMinter: nil,
				TimeLinearMinter: &cfemintertypes.TimeLinearMinter{
					Amount: sdk.NewInt(1000000000),
				},
			},
		},
	}
	newMinterJSON, err := json.Marshal(newMinter)
	s.NoError(err)

	newMintDenomJSON, err := json.Marshal("uc4e")
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "Cfeminter module params change",
		Description: "Change cfeminter params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMinter),
				Value:    newMinterJSON,
			},
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMintDenom),
				Value:    newMintDenomJSON,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	s.NoError(err)
	node.SubmitParamChangeProposal(string(proposalJSON), initialization.ValidatorWalletName)
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
			node.QueryParams(cfemintertypes.ModuleName, string(cfemintertypes.KeyMinter), &params)
			return params.Value == string(newMinterJSON)
		},
		1*time.Minute,
		10*time.Millisecond,
		"C4e node failed to retrieve params",
	)

	s.Eventually(
		func() bool {
			var params Params
			node.QueryParams(cfemintertypes.ModuleName, string(cfemintertypes.KeyMintDenom), &params)
			return params.Value == string(newMintDenomJSON)
		},
		1*time.Minute,
		10*time.Millisecond,
		"C4e node failed to retrieve params",
	)
}

func (s *IntegrationTestSuite) TestCreateVestingPool() {
	const (
		walletName  = "user-1"
		baseBalance = 10000000
	)
	chainA := s.configurer.GetChainConfig(0)
	chainANode, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorAddress := chainANode.CreateWallet(walletName)
	chainANode.BankSend(sdk.NewCoin(appparams.CoinDenom, sdk.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)

	balance, err := chainANode.QueryBalances(creatorAddress)
	balanceAmount := balance.AmountOf(appparams.CoinDenom)
	fmt.Println("Balance before: " + balance.String())
	vestingAmount := balanceAmount.Quo(sdk.NewInt(4))
	fmt.Println("Vesting amount: " + vestingAmount.String())
	s.NoError(err)

	chainANode.CreateVestingPool(
		helpers.RandStringOfLength(5),
		vestingAmount.String(),
		(10 * time.Minute).String(),
		"Short vesting with lockup",
		walletName)
	fmt.Println("Created vesting pool: " + vestingAmount.String())
	newBalance, err := chainANode.QueryBalances(creatorAddress)
	s.NoError(err)
	s.Equal(balanceAmount.Sub(vestingAmount), newBalance)
	vestingPools := chainANode.QueryVestingPools(creatorAddress)
	s.Equal(1, len(vestingPools))
	// setup wallets and send gamm tokens to these wallets on chainA

}
