package upgrades

import (
	cfedistributorv2migration "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v2"
	cfedistributormoduletypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfeminterv2migration "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v2"
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
			keyTable = cfedistributorv2migration.ParamKeyTable() //nolint:staticcheck

		case cfemintermoduletypes.ModuleName:
			keyTable = cfeminterv2migration.ParamKeyTable() //nolint:staticcheck

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
