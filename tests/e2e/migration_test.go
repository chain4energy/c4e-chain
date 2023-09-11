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
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, false, false, nil)
}

func (s *MainnetMigrationSetupSuite) TestMainnetVestingsMigration() {

}

type MainnetMigrationChainingSetupSuite struct {
	BaseSetupSuite
}

func TestMainnetMigrationChainingSuite(t *testing.T) {
	suite.Run(t, new(MainnetMigrationChainingSetupSuite))
}

func (s *MainnetMigrationChainingSetupSuite) SetupSuite() {
	//bytes, err := os.ReadFile("./resources/mainnet-v1.1.0-migration-app-state.json")
	//if err != nil {
	//	panic(err)
	//}
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, true, false, nil)
}

func (s *MainnetMigrationChainingSetupSuite) TestMainnetVestingsMigrationWhenChainingMigrations() {

}
