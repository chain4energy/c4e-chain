package v200

import (
	"github.com/chain4energy/c4e-chain/app/upgrades"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const UpgradeName = "v2.0.0" // TODO raczej 1.3.0 ale do ustalenia jeszcze

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			cfeclaimtypes.ModuleName,
		},
	},
}
