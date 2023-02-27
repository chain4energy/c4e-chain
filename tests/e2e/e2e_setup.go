package e2e

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/chain"
	"github.com/stretchr/testify/suite"
	"os"
	"strconv"
	"time"
)

const (
	AverageBlockTime  = time.Second * 6
	debugLogEnv       = "C4E_E2E_DEBUG_LOG"
	forkHeightEnv     = "C4E_E2E_FORK_HEIGHT"
	skipCleanupEnv    = "C4E_E2E_SKIP_CLEANUP"
	upgradeVersionEnv = "C4E_E2E_UPGRADE_VERSION"
)

type BaseSetupSuite struct {
	suite.Suite
	configurer configurer.Configurer
	forkHeight int
}

func (s *BaseSetupSuite) SetupSuite(startUpgrade, startIBC bool) {
	s.T().Log("setting up e2e integration test suite...")
	var (
		err             error
		upgradeSettings configurer.UpgradeSettings
	)

	if startUpgrade {
		s.T().Log("start upgrade was true, starting upgrade setup")
		upgradeSettings.IsEnabled = startUpgrade
		if str := os.Getenv(upgradeVersionEnv); len(str) > 0 {
			upgradeSettings.Version = str
			s.T().Log(fmt.Sprintf("upgrade version set to %s", upgradeSettings.Version))
		}

	}

	if startIBC {
		s.T().Log("startIBC was true, starting IBC setup")
	}

	isDebugLogEnabled := false
	if str := os.Getenv(debugLogEnv); len(str) > 0 {
		isDebugLogEnabled, err = strconv.ParseBool(str)
		s.Require().NoError(err)
		if isDebugLogEnabled {
			s.T().Log("debug logging is enabled. container logs from running cli commands will be printed to stdout")
		}
	}

	s.configurer, err = configurer.StartDockerContainers(s.T(), startIBC, isDebugLogEnabled, upgradeSettings)
	s.Require().NoError(err)
}

func (s *BaseSetupSuite) TearDownSuite() {
	if str := os.Getenv(skipCleanupEnv); len(str) > 0 {
		skipCleanup, err := strconv.ParseBool(str)
		s.Require().NoError(err)

		if skipCleanup {
			s.T().Log("skipping e2e resources clean up...")
			return
		}
	}

	err := s.configurer.ClearResources()
	s.Require().NoError(err)
}

func (s *BaseSetupSuite) validateTotalSupply(node *chain.NodeConfig, denom string, gte bool, waitFor time.Duration) {
	totalSupplyBefore, err := node.QuerySupplyOf(denom)
	s.NoError(err)
	time.Sleep(time.Second * waitFor)
	totalSupplyAfter, err := node.QuerySupplyOf(denom)
	s.NoError(err)
	s.Equal(totalSupplyAfter.GT(totalSupplyBefore), gte)
}

func (s *BaseSetupSuite) validateTotalSupplyAfterPeriod(node *chain.NodeConfig, denom string, increment, sequenceId int) {
	for i := 0; i < sequenceId; i++ {
		totalSupplyBefore, err := node.QuerySupplyOf(denom)
		s.NoError(err)
		time.Sleep(AverageBlockTime)
		totalSupplyAfter, err := node.QuerySupplyOf(denom)
		fmt.Println(totalSupplyAfter.Int64())
		fmt.Println(totalSupplyBefore.Int64())
		s.NoError(err)
		s.Equal(totalSupplyAfter, totalSupplyBefore.AddRaw(int64(increment)))
	}
}

func (s *BaseSetupSuite) validateBalanceOfAccount(node *chain.NodeConfig, denom, accAddress string, gte bool, waitFor time.Duration) {
	totalSupplyBefore, err := node.QueryBalances(accAddress)
	s.NoError(err)
	time.Sleep(time.Second * waitFor)
	totalSupplyAfter, err := node.QueryBalances(accAddress)
	s.NoError(err)
	s.Equal(totalSupplyAfter.AmountOf(denom).GT(totalSupplyBefore.AmountOf(denom)), gte)
}
