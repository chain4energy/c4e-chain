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
	bytes, err := os.ReadFile("./resources/mainnet-migration-app-state.json")
	if err != nil {
		panic(err)
	}
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, false, bytes)
}

func (s *MainnetMigrationSetupSuite) TestMainnetVestingsMigration() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	campaigns := node.QueryCampaigns()
	s.Equal(4, len(campaigns))

	userEntries := node.QueryUserEntries()
	s.Equal(107404, len(userEntries))
	// TODO: add verifications and more options (probably when the final version of the migration will be set)
}
