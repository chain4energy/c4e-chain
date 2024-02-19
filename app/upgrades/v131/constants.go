package v131

import (
	"github.com/chain4energy/c4e-chain/app/upgrades"
)

const UpgradeName = "v1.3.1"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
}
