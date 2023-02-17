package v120

import (
	"github.com/chain4energy/c4e-chain/app/upgrades"
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	appKeepers cfeupgradetypes.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		//todo Prepare migration
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
