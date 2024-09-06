package v141

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/chain4energy/c4e-chain/app/upgrades"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	appKeepers upgrades.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		if _, exists := vm[wasmtypes.ModuleName]; !exists {
			ctx.Logger().Info("wasm not exist")
			for _, subspace := range appKeepers.GetC4eParamsKeeper().GetSubspaces() {
				subspace := subspace

				var keyTable paramstypes.KeyTable
				switch subspace.Name() {
				case authtypes.ModuleName:
					keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
				case banktypes.ModuleName:
					keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
				case stakingtypes.ModuleName:
					keyTable = stakingtypes.ParamKeyTable()
				case minttypes.ModuleName:
					keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
				case distrtypes.ModuleName:
					keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
				case slashingtypes.ModuleName:
					keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
				case govtypes.ModuleName:
					keyTable = govv1.ParamKeyTable() //nolint:staticcheck
				case crisistypes.ModuleName:
					keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck
					// ibc types
				case ibctransfertypes.ModuleName:
					keyTable = ibctransfertypes.ParamKeyTable()
				case icahosttypes.SubModuleName:
					keyTable = icahosttypes.ParamKeyTable()
				case icacontrollertypes.SubModuleName:
					keyTable = icacontrollertypes.ParamKeyTable()
					// wasm
				case wasmtypes.ModuleName:
					keyTable = wasmtypes.ParamKeyTable() //nolint:staticcheck
				default:
					continue
				}

				if !subspace.HasKeyTable() {
					subspace.WithKeyTable(keyTable)
				}
			}

			baseAppLegacySS := appKeepers.GetC4eParamsKeeper().Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())
			baseapp.MigrateParams(ctx, baseAppLegacySS, appKeepers.GetC4eConsensusParamsKeeper())
			vmResult, err := mm.RunMigrations(ctx, configurator, vm)
			if err != nil {
				return vmResult, err
			}

			// Set permission for uploading new Wasm code with specific addresses
			wasmParams := wasmtypes.Params{
				CodeUploadAccess: wasmtypes.AccessConfig{
					Permission: wasmtypes.AccessTypeAnyOfAddresses,
					Addresses: []string{
						"c4e1r2ennr6ywv567lks3q5gujt4def726fe3t2tus",
						"c4e1e0ddzmhw2ze2glszkgjk6tfvcfzv68cmnfrnaq",
						"c4e19473sdmlkkvcdh6z3tqedtqsdqj4jjv7htsuaa",
						"c4e1psaq0n2lzh84lzgh39kghuy0n256xltlcmea52",
						"c4e1jr0ft7p2fgqxjrqxsakz9re0ae5499uz2cfmra",
						"c4e1tw2crl23vluafhcvydhpnkejwth70y9knpsht7",
						"c4e19x0fmrjnhqgze4c0c0st5jdpqgu3t4a2zn9t8r",
						"c4e183f5fu67gagckux336kmjw75qw7dha5ycn5f6r",
						"c4e10qfgech3v82uztzl20tl7uldsq8nk9gl92mm55",
						"c4e16cwpandmj9np4huguzs32g0htm58p0cp9df8gj",
						"c4e10ep2sxpf2kj6j26w7f4uuafedkuf9sf9xqq3sl",
						"c4e16n7yweagu3fxfzvay6cz035hddda7z3ntdxq3l",
					},
				},
				InstantiateDefaultPermission: wasmtypes.AccessTypeEverybody,
			}

			err = appKeepers.GetWasmKeeper().SetParams(ctx, wasmParams)
			if err != nil {
				return vmResult, err
			}

			return vmResult, nil

		} else {
			return mm.RunMigrations(ctx, configurator, vm)
		}
	}
}
