package v130

import (
	"github.com/chain4energy/c4e-chain/v2/app/upgrades"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const UpgradeName = "v1.3.0"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{cfeclaimtypes.ModuleName},
	},
}
