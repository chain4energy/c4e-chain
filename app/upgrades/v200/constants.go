package v200

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/chain4energy/c4e-chain/v2/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
)

const UpgradeName = "v2.0.0"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			consensusparamtypes.ModuleName,
			crisistypes.ModuleName,
			wasmtypes.StoreKey,
		},
		Deleted: []string{
			"cfesignature",
		},
	},
}
