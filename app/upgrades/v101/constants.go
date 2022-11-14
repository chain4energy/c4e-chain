package v2

import (
	"github.com/chain4energy/c4e-chain/app/upgrades"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	//store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	icacontrollertypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/types"
)

const UpgradeName = "v1.0.1"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			group.ModuleName,
			icahosttypes.StoreKey,
			//ibcfeetypes.StoreKey,
			icacontrollertypes.StoreKey,
		},
	},
}
