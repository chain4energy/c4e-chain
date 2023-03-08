package v3

import (
	"github.com/chain4energy/c4e-chain/x/cfeminter/exported"
	v2 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v2"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateParams(ctx sdk.Context, storeKey storetypes.StoreKey, legacySubspace exported.Subspace, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	var oldParams types.LegacyParams
	legacySubspace.GetParamSet(ctx, &oldParams)

	var newParams types.Params
	newParams.MintDenom = oldParams.MintDenom
	newParams.StartTime = oldParams.MinterConfig.StartTime
	for _, oldMinter := range oldParams.MinterConfig.Minters {
		var newMinter types.Minter
		newMinter.SequenceId = oldMinter.SequenceId
		newMinter.EndTime = oldMinter.EndTime
		newParams.Minters = append(newParams.Minters, &newMinter)
		var config *codectypes.Any
		if oldMinter.Type == v2.ExponentialStepMintingType {
			config, _ = codectypes.NewAnyWithValue(oldMinter.ExponentialStepMinting)
		} else if oldMinter.Type == v2.LinearMintingType {
			config, _ = codectypes.NewAnyWithValue(oldMinter.LinearMinting)
		} else {
			config, _ = codectypes.NewAnyWithValue(&types.NoMinting{})
		}
		newMinter.Config = config
	}

	if err := newParams.Validate(); err != nil {
		return err
	}

	bz, err := cdc.Marshal(&newParams)
	if err != nil {
		return err
	}

	store.Set(types.ParamsKey, bz)

	return nil
}
