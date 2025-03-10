package configurer

import (
	"cosmossdk.io/math"
	"encoding/json"
	"fmt"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"os"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/chain"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/config"
	"github.com/chain4energy/c4e-chain/tests/e2e/containers"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
)

type UpgradeSettings struct {
	IsEnabled               bool
	Version                 string
	OldInitialAppStateBytes []byte
	MigrationChaining       bool
}

type UpgradeConfigurer struct {
	baseConfigurer
	upgradeVersion        string
	upgradeLegacyProposal bool
}

var _ Configurer = (*UpgradeConfigurer)(nil)

func NewUpgradeConfigurer(t *testing.T, chainConfigs []*chain.Config, setupTests setupFn, containerManager *containers.Manager, upgradeVersion string, upgradeLegacyProposal bool) Configurer {
	return &UpgradeConfigurer{
		baseConfigurer: baseConfigurer{
			chainConfigs:     chainConfigs,
			containerManager: containerManager,
			setupTests:       setupTests,
			syncUntilHeight:  defaultSyncUntilHeight,
			t:                t,
		},
		upgradeVersion:        upgradeVersion,
		upgradeLegacyProposal: upgradeLegacyProposal,
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

	chainInitResource, err := uc.containerManager.RunChainInitResource(chainConfig.Id, int(chainConfig.VotingPeriod), int(chainConfig.ExpeditedVotingPeriod), validatorConfigBytes, tmpDir, chainConfig.AppStateBytes)
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
				if uc.upgradeLegacyProposal { // TODO: fix after new upgrade
					node.SubmitLegacyUpgradeProposal(uc.upgradeVersion, chainConfig.UpgradePropHeight, sdk.NewCoin(appparams.MicroC4eUnit, math.NewInt(config.InitialMinDeposit)))
				} else {
					node.SubmitUpgradeProposal(uc.upgradeVersion, chainConfig.UpgradePropHeight, sdk.NewCoin(appparams.MicroC4eUnit, math.NewInt(config.InitialMinDeposit)))
				}
				chainConfig.LatestProposalNumber += 1
				node.DepositProposalLegacy(chainConfig.LatestProposalNumber)
			}
			node.VoteYesProposalLegacy(initialization.ValidatorWalletName, chainConfig.LatestProposalNumber)
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
	uc.t.Logf("upgrade height reached successful on chain %s", chainConfig.Id)
	chainConfig.WaitUntilHeight(propHeight + config.UpgradeBufferBlocks)
	uc.t.Logf("upgrade successful on chain %s", chainConfig.Id)
	return nil
}
