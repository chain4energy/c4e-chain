package v200

import (
	"github.com/chain4energy/c4e-chain/app/upgrades"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfeevmoduletypes "github.com/chain4energy/c4e-chain/x/cfeev/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const UpgradeName = "v2.0.0"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{cfeclaimtypes.ModuleName, cfeevmoduletypes.ModuleName},
	},
}
