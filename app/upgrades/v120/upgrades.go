package v120

import (
	"github.com/chain4energy/c4e-chain/app/upgrades"
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfedistributormoduletypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfemintermoduletypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfesignaturetypes "github.com/chain4energy/c4e-chain/x/cfesignature/types"
	cfevestingmoduletypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	appKeepers cfeupgradetypes.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for _, subspace := range appKeepers.GetParamKeeper().GetSubspaces() {
			subspace := subspace

			var keyTable paramstypes.KeyTable
			switch subspace.Name() {
			case cfedistributormoduletypes.ModuleName:
				keyTable = cfedistributormoduletypes.ParamKeyTable() //nolint:staticcheck

			case cfemintermoduletypes.ModuleName:
				keyTable = cfemintermoduletypes.ParamKeyTable() //nolint:staticcheck

			case cfevestingmoduletypes.ModuleName:
				keyTable = cfevestingmoduletypes.ParamKeyTable() //nolint:staticcheck

			case cfesignaturetypes.ModuleName:
				keyTable = cfesignaturetypes.ParamKeyTable() //nolint:staticcheck
			}
			if !subspace.HasKeyTable() {
				subspace.WithKeyTable(keyTable)

			}
		}
		vmResult, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vmResult, err
		}
		UpdateVestingAccountTraces(ctx, appKeepers)
		if err := ModifyVestingPoolsState(ctx, appKeepers); err != nil {
			return vmResult, err
		}
		return vmResult, ModifyVestingAccountsState(ctx, appKeepers)
	}
}
