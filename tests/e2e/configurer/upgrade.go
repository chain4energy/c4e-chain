package configurer

import (
	"encoding/json"
	"fmt"
	"github.com/chain4energy/c4e-chain/app/params"
	"os"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/chain"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/config"
	"github.com/chain4energy/c4e-chain/tests/e2e/containers"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
)

type UpgradeSettings struct {
	IsEnabled  bool
	Version    string
	ForkHeight int64 // non-zero height implies that this is a fork upgrade.
}

type UpgradeConfigurer struct {
	baseConfigurer
	upgradeVersion string
	forkHeight     int64 // forkHeight > 0 implies that this is a fork upgrade. Otherwise, proposal upgrade.
}

var _ Configurer = (*UpgradeConfigurer)(nil)

func NewUpgradeConfigurer(t *testing.T, chainConfigs []*chain.Config, setupTests setupFn, containerManager *containers.Manager, upgradeVersion string, forkHeight int64) Configurer {
	return &UpgradeConfigurer{
		baseConfigurer: baseConfigurer{
			chainConfigs:     chainConfigs,
			containerManager: containerManager,
			setupTests:       setupTests,
			syncUntilHeight:  forkHeight + defaultSyncUntilHeight,
			t:                t,
		},
		forkHeight:     forkHeight,
		upgradeVersion: upgradeVersion,
	}
}

func (uc *UpgradeConfigurer) ConfigureChains() error {
	for _, chainConfig := range uc.chainConfigs {
		if err := uc.ConfigureChain(chainConfig); err != nil {
			return err
		}
	}
	return nil
}

func (uc *UpgradeConfigurer) ConfigureChain(chainConfig *chain.Config) error {
	uc.t.Logf("starting upgrade e2e infrastructure for chain-id: %s", chainConfig.Id)
	tmpDir, err := os.MkdirTemp("", "chain4energy-e2e-testnet-")
	if err != nil {
		return err
	}

	validatorConfigBytes, err := json.Marshal(chainConfig.ValidatorInitConfigs)
	if err != nil {
		return err
	}

	forkHeight := uc.forkHeight
	if forkHeight > 0 {
		forkHeight = forkHeight - config.ForkHeightPreUpgradeOffset
	}

	chainInitResource, err := uc.containerManager.RunChainInitResource(chainConfig.Id, int(chainConfig.VotingPeriod), int(chainConfig.ExpeditedVotingPeriod), validatorConfigBytes, tmpDir, int(forkHeight))
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%v/%v-encode", tmpDir, chainConfig.Id)
	uc.t.Logf("serialized init file for chain-id %v: %v", chainConfig.Id, fileName)

	// loop through the reading and unmarshaling of the init file a total of maxRetries or until error is nil
	// without this, test attempts to unmarshal file before docker container is finished writing
	var initializedChain initialization.Chain
	for i := 0; i < config.MaxRetries; i++ {
		initializedChainBytes, _ := os.ReadFile(fileName)
		err = json.Unmarshal(initializedChainBytes, &initializedChain)
		if err == nil {
			break
		}

		if i == config.MaxRetries-1 {
			if err != nil {
				return err
			}
		}

		if i > 0 {
			time.Sleep(1 * time.Second)
		}
	}
	if err := uc.containerManager.PurgeResource(chainInitResource); err != nil {
		return err
	}
	uc.initializeChainConfigFromInitChain(&initializedChain, chainConfig)
	return nil
}

func (uc *UpgradeConfigurer) RunSetup() error {
	return uc.setupTests(uc)
}

func (uc *UpgradeConfigurer) RunUpgrade() error {
	if uc.forkHeight > 0 {
		return uc.runForkUpgrade()
	}
	return uc.runProposalUpgrade()
}

func (uc *UpgradeConfigurer) runProposalUpgrade() error {
	for _, chainConfig := range uc.chainConfigs {
		for validatorIdx, node := range chainConfig.NodeConfigs {
			if validatorIdx == 0 {
				currentHeight, err := node.QueryCurrentHeight()
				if err != nil {
					return err
				}
				chainConfig.UpgradePropHeight = currentHeight + int64(chainConfig.VotingPeriod) + int64(config.PropSubmitBlocks) + int64(config.PropBufferBlocks)
				node.SubmitUpgradeProposal(uc.upgradeVersion, chainConfig.UpgradePropHeight, sdk.NewCoin(params.CoinDenom, sdk.NewInt(config.InitialMinDeposit)))
				chainConfig.LatestProposalNumber += 1
				node.DepositProposal(chainConfig.LatestProposalNumber)
			}
			node.VoteYesProposal(initialization.ValidatorWalletName, chainConfig.LatestProposalNumber)
		}
	}

	// wait till all chains halt at upgrade height
	for _, chainConfig := range uc.chainConfigs {
		uc.t.Logf("waiting to reach upgrade height on chain %s", chainConfig.Id)
		chainConfig.WaitUntilHeight(chainConfig.UpgradePropHeight)
		uc.t.Logf("upgrade height reached on chain %s", chainConfig.Id)
	}

	// remove all containers so we can upgrade them to the new version
	for _, chainConfig := range uc.chainConfigs {
		for _, validatorConfig := range chainConfig.NodeConfigs {
			err := uc.containerManager.RemoveNodeResource(validatorConfig.Name)
			if err != nil {
				return err
			}
		}
	}

	// remove all containers so we can upgrade them to the new version
	for _, chainConfig := range uc.chainConfigs {
		if err := uc.upgradeContainers(chainConfig, chainConfig.UpgradePropHeight); err != nil {
			return err
		}
	}
	return nil
}

func (uc *UpgradeConfigurer) runForkUpgrade() error {
	for _, chainConfig := range uc.chainConfigs {
		uc.t.Logf("waiting to reach fork height on chain %s", chainConfig.Id)
		chainConfig.WaitUntilHeight(uc.forkHeight)
		uc.t.Logf("fork height reached on chain %s", chainConfig.Id)
	}
	return nil
}

func (uc *UpgradeConfigurer) upgradeContainers(chainConfig *chain.Config, propHeight int64) error {
	// upgrade containers to the locally compiled daemon
	uc.t.Logf("starting upgrade for chain-id: %s...", chainConfig.Id)
	uc.containerManager.C4eRepository = containers.CurrentBranchC4eRepository
	uc.containerManager.C4eTag = containers.CurrentBranchC4eTag

	for _, node := range chainConfig.NodeConfigs {
		if err := node.Run(); err != nil {
			return err
		}
	}

	uc.t.Logf("waiting to upgrade containers on chain %s", chainConfig.Id)
	chainConfig.WaitUntilHeight(propHeight)
	uc.t.Logf("upgrade successful on chain %s", chainConfig.Id)
	return nil
}