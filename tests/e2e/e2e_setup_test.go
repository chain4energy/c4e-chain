package e2e

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer"
	"github.com/stretchr/testify/suite"
	"os"
	"strconv"
)

const (
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
	os.Setenv(upgradeVersionEnv, "v1.0.1")
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

		if str := os.Getenv(forkHeightEnv); len(str) > 0 {
			upgradeSettings.ForkHeight, err = strconv.ParseInt(str, 0, 64)
			s.Require().NoError(err)
			s.T().Log(fmt.Sprintf("fork upgrade is enabled, %s was set to height %d", forkHeightEnv, upgradeSettings.ForkHeight))
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

	s.configurer, err = configurer.New(s.T(), startIBC, isDebugLogEnabled, upgradeSettings)
	s.Require().NoError(err)

	err = s.configurer.ConfigureChains()
	s.Require().NoError(err)

	err = s.configurer.RunSetup()
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
