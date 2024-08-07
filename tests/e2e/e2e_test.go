package e2e

import (
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer"
	"testing"
)

func TestRunChainWithOptions(t *testing.T) {
	var upgradeSettings configurer.UpgradeSettings

	upgradeSettings.Version = "v1.3.1"
	upgradeSettings.IsEnabled = false

	_, err := configurer.StartDockerContainers(t, false, true, "", upgradeSettings)
	if err != nil {
		return
	}
}
