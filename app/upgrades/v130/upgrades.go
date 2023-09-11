package v130

import (
	"github.com/chain4energy/c4e-chain/v2/app/upgrades"
	"github.com/chain4energy/c4e-chain/v2/app/upgrades/v130/claim"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	appKeepers upgrades.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		vmResult, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vmResult, err
		}

		if err = ModifyAndAddVestingTypes(ctx, appKeepers); err != nil {
			return vmResult, err
		}
		if err = MigrateAirdropModuleAccount(ctx, appKeepers); err != nil {
			return vmResult, err
		}
		if err = MigrateMoondropVestingAccount(ctx, appKeepers); err != nil {
			return vmResult, err
		}
		return vmResult, claim.SetupAirdrops(ctx, appKeepers)
	}
}
