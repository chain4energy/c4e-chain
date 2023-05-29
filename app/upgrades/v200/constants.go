package v200

import (
	"github.com/chain4energy/c4e-chain/app/upgrades"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	icacontrollertypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/types"
)

const UpgradeName = "v2.0.0"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{group.ModuleName, icahosttypes.StoreKey, icacontrollertypes.StoreKey, cfeclaimtypes.ModuleName},
		Deleted: []string{"monitoringp"},
	},
}
