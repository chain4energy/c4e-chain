package e2e

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type MainnetMigrationSetupSuite struct {
	BaseSetupSuite
}

func TestMainnetMigrationSuite(t *testing.T) {
	suite.Run(t, new(MainnetMigrationSetupSuite))
}

func (s *MainnetMigrationSetupSuite) SetupSuite() {
	bytes, err := os.ReadFile("./resources/mainnet-v1.3.1-migration-app-state.json")
	if err != nil {
		panic(err)
	}
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, false, false, bytes)
}

func (s *MainnetMigrationSetupSuite) TestMainnetVestingsMigration() {
	chainA := s.configurer.GetChainConfig(0)
	_, err := chainA.GetDefaultNode()
	s.NoError(err)
}
