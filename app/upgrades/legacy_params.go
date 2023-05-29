package upgrades

import (
	cfedistributormoduletypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfemintermoduletypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfesignaturetypes "github.com/chain4energy/c4e-chain/x/cfesignature/types"
	cfevestingv3migration "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v3"
	cfevestingmoduletypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func RegisterLegacyParamsKeyTables(appKeepers AppKeepers) {
	for _, subspace := range appKeepers.GetParamKeeper().GetSubspaces() {
		subspace := subspace

		var keyTable paramstypes.KeyTable
		switch subspace.Name() {
		case cfedistributormoduletypes.ModuleName:
			keyTable = cfedistributormoduletypes.ParamKeyTable() //nolint:staticcheck

		case cfemintermoduletypes.ModuleName:
			keyTable = cfemintermoduletypes.ParamKeyTable() //nolint:staticcheck

		case cfevestingmoduletypes.ModuleName:
			keyTable = cfevestingv3migration.ParamKeyTable() //nolint:staticcheck

		case cfesignaturetypes.ModuleName:
			keyTable = cfesignaturetypes.ParamKeyTable() //nolint:staticcheck
		}
		if !subspace.HasKeyTable() {
			subspace.WithKeyTable(keyTable)
		}
	}
}
