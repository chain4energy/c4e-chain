package e2e

import (
	"os"
	"testing"

	v120 "github.com/chain4energy/c4e-chain/app/upgrades/v120"
	"github.com/stretchr/testify/suite"
)

type MigrationSetupSuite struct {
	BaseSetupSuite
}

func TestMigrationSuite(t *testing.T) {
	suite.Run(t, new(MigrationSetupSuite))
}

func (s *MigrationSetupSuite) SetupSuite() {
	bytes, err := os.ReadFile("./resources/mainnet-vestings-migration-app-state.json")
	if err != nil {
		panic(err)
	}
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, false, bytes)
}

func (s *MigrationSetupSuite) TestVestingsMigration() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	vestingTypes := node.QueryVestingTypes()
	s.Equal(7, len(vestingTypes))

	vestingPools := node.QueryVestingPools(v120.ValidatorsVestingPoolOwner)
	s.Equal(6, len(vestingPools))

}
