package e2e

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type MainnetMigrationSetupSuite struct {
	BaseSetupSuite
}

func TestMainnetMigrationSuite(t *testing.T) {
	suite.Run(t, new(MainnetMigrationSetupSuite))
}

func (s *MainnetMigrationSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(true, false)
}

func (s *MainnetMigrationSetupSuite) TestMainnetVestingsMigration() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	campaigns := node.QueryCampaigns()
	s.Equal(4, len(campaigns))

	userEntries := node.QueryUserEntries()
	s.Equal(107404, len(userEntries))

}
