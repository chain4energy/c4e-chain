package v140

import (
	"github.com/chain4energy/c4e-chain/app/upgrades"
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
		if err = UpdateCommunityPool(ctx, appKeepers); err != nil {
			return vmResult, err
		}
		if err = ModifyVestingTypes(ctx, appKeepers); err != nil {
			return vmResult, err
		}
		if err = UpdateStrategicReserveShortTermPool(ctx, appKeepers); err != nil {
			return vmResult, err
		}
		return vmResult, UpdateStrategicReserveAccount(ctx, appKeepers)
	}
}
