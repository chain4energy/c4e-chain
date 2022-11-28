package e2e

import (
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/chain"
	"testing"
	"time"
)

func TestRunChainWithOptions(t *testing.T) {
	var upgradeSettings configurer.UpgradeSettings

	//upgradeSettings.Version = "v1.0.1"
	//upgradeSettings.IsEnabled = true

	_, err := configurer.StartDockerContainers(t, false, true, upgradeSettings)
	if err != nil {
		return
	}
}

func (s *BaseSetupSuite) validateTotalSupply(node *chain.NodeConfig, denom string, gte bool, waitFor time.Duration) {
	totalSupplyBefore, err := node.QuerySupplyOf(denom)
	s.NoError(err)
	time.Sleep(time.Second * waitFor)
	totalSupplyAfter, err := node.QuerySupplyOf(denom)
	s.NoError(err)
	s.Equal(totalSupplyAfter.GT(totalSupplyBefore), gte)
}

func (s *BaseSetupSuite) validateBalanceOfAccount(node *chain.NodeConfig, denom, accAddress string, gte bool, waitFor time.Duration) {
	totalSupplyBefore, err := node.QueryBalances(accAddress)
	s.NoError(err)
	time.Sleep(time.Second * waitFor)
	totalSupplyAfter, err := node.QueryBalances(accAddress)
	s.NoError(err)
	s.Equal(totalSupplyAfter.AmountOf(denom).GT(totalSupplyBefore.AmountOf(denom)), gte)
}
