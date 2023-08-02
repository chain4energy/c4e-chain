package cfetokenization

import (
	"github.com/chain4energy/c4e-chain/x/cfetokenization/keeper"
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the certificateType
	for _, elem := range genState.CertificateTypeList {
		k.SetCertificateType(ctx, elem)
	}

	// Set certificateType count
	k.SetCertificateTypeCount(ctx, genState.CertificateTypeCount)
	// Set all the userDevices
	for _, elem := range genState.UserDevicesList {
		k.SetUserDevices(ctx, elem)
	}

	// Set all the userCertificates
	for _, elem := range genState.UserCertificatesList {
		k.SetUserCertificates(ctx, elem)
	}
	for _, elem := range genState.DevicesList {
		k.SetDevice(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CertificateTypeList = k.GetAllCertificateType(ctx)
	genesis.CertificateTypeCount = k.GetCertificateTypeCount(ctx)
	genesis.UserDevicesList = k.GetAllUserDevices(ctx)
	genesis.DevicesList = k.GetAllDevices(ctx)
	genesis.UserCertificatesList = k.GetAllUserCertificates(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
