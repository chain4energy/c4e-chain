package v3

import (
	"github.com/chain4energy/c4e-chain/v2/types/subspace"
	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/migrations/v2"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var ParamsKey = []byte{0x00}

// MigrateParams migrates the x/cfeminter module state from the consensus version 2 to
// version 3. Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/cfeminter
// module state.
// The migration also includes:
// - cfeminter module refactoring
// - delete type field from minterConfig
// - minter config is now of type Any rather than using 2 fields (LinearMinting and ExponentialStepMinting)
// - MinterConfig was deleted and minters and start-time was moved directly to cfeminter params
func MigrateParams(ctx sdk.Context, storeKey storetypes.StoreKey, legacySubspace subspace.Subspace, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	var oldParams v2.Params
	if !legacySubspace.HasKeyTable() {
		legacySubspace.WithKeyTable(v2.ParamKeyTable())
	}
	legacySubspace.GetParamSet(ctx, &oldParams)

	var newParams Params
	newParams.MintDenom = oldParams.MintDenom
	newParams.StartTime = oldParams.MinterConfig.StartTime
	for _, oldMinter := range oldParams.MinterConfig.Minters {
		var newMinter Minter
		newMinter.SequenceId = oldMinter.SequenceId
		newMinter.EndTime = oldMinter.EndTime
		newParams.Minters = append(newParams.Minters, &newMinter)
		var config *codectypes.Any
		var err error
		if oldMinter.Type == v2.ExponentialStepMintingType {
			config, err = codectypes.NewAnyWithValue(oldMinter.ExponentialStepMinting)
			config.TypeUrl = "/c4echain.cfeminter.ExponentialStepMinting"
			if err != nil {
				return err
			}
		} else if oldMinter.Type == v2.LinearMintingType {
			config, err = codectypes.NewAnyWithValue(oldMinter.LinearMinting)
			config.TypeUrl = "/c4echain.cfeminter.LinearMinting"
			if err != nil {
				return err
			}
		} else {
			config, err = codectypes.NewAnyWithValue(&NoMinting{})
			config.TypeUrl = "/c4echain.cfeminter.NoMinting"
			if err != nil {
				return err
			}
		}
		newMinter.Config = config
	}

	bz, err := cdc.Marshal(&newParams)
	if err != nil {
		return err
	}

	store.Set(ParamsKey, bz)

	return nil
}
