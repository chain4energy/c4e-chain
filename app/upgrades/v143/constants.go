package v143

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/chain4energy/c4e-chain/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v7/types"
)

const UpgradeName = "v1.4.3"
const UpgradeNameTn = "v1.4.3-tn"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			consensusparamtypes.ModuleName,
			crisistypes.ModuleName,
			wasmtypes.StoreKey,
			ibchookstypes.StoreKey,
		},
		Deleted: []string{
			"cfesignature",
		},
	},
}

var UpgradeTn = upgrades.Upgrade{
	UpgradeName:          UpgradeNameTn,
	CreateUpgradeHandler: CreateUpgradeHandlerTn,
	StoreUpgrades:        store.StoreUpgrades{},
}
