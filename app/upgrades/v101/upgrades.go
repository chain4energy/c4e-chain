package v2

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/app/upgrades"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ica "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts"
	icacontrollertypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	//keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// set regen module consensus version
		//fromVM[ecocredit.ModuleName] = 2
		//fromVM[data.ModuleName] = 1

		// save oldIcaVersion, so we can skip icahost.InitModule in longer term tests.
		oldIcaVersion := fromVM[icatypes.ModuleName]

		// Add Interchain Accounts host module
		// set the ICS27 consensus version so InitGenesis is not run
		fromVM[icatypes.ModuleName] = mm.Modules[icatypes.ModuleName].ConsensusVersion()

		// create ICS27 Controller submodule params, controller module not enabled.
		controllerParams := icacontrollertypes.Params{ControllerEnabled: false}

		// create ICS27 Host submodule params, host module not enabled.
		hostParams := icahosttypes.Params{
			HostEnabled:   false,
			AllowMessages: []string{},
		}

		mod, found := mm.Modules[icatypes.ModuleName]
		if !found {
			panic(fmt.Sprintf("module %s is not in the module manager", icatypes.ModuleName))
		}

		icaMod, ok := mod.(ica.AppModule)
		if !ok {
			panic(fmt.Sprintf("expected module %s to be type %T, got %T", icatypes.ModuleName, ica.AppModule{}, mod))
		}

		// skip InitModule in upgrade tests after the upgrade has gone through.
		if oldIcaVersion != fromVM[icatypes.ModuleName] {
			icaMod.InitModule(ctx, controllerParams, hostParams)
		}

		// transfer module consensus version has been bumped to 2
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
