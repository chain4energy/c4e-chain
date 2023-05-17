package v200

import (
	"github.com/chain4energy/c4e-chain/app/upgrades"
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	"github.com/chain4energy/c4e-chain/app/upgrades/v200/claim"
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
		vmResult, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vmResult, err
		}

		if err = modifyAndAddVestingTypes(ctx, appKeepers); err != nil {
			return vmResult, err
		}
		if err = migrateAirdropModuleAccount(ctx, appKeepers); err != nil {
			return vmResult, err
		}
		if err = migrateVestingAccount(ctx, appKeepers); err != nil {
			return vmResult, err
		}
		return vmResult, claim.Creates(ctx, appKeepers)
	}
}
