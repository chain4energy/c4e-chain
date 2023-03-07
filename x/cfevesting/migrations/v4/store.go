package v4

import (
	"github.com/chain4energy/c4e-chain/x/cfedistributor/exported"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MigrateStore migrates the x/bank module state from the consensus version 2 to
// version 3. Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/bank
// module state.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, legacySubspace exported.Subspace, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	var currParams types.Params
	legacySubspace.GetParamSet(ctx, &currParams)

	if err := currParams.Validate(); err != nil {
		return err
	}

	bz, err := cdc.Marshal(&currParams)
	if err != nil {
		return err
	}

	store.Set(ParamsKey, bz)

	return nil
}
