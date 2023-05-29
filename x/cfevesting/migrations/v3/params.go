package v3

import (
	"github.com/chain4energy/c4e-chain/types/subspace"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MigrateParams migrates the x/cfevesting module state from the consensus version 2 to
// version 3. Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/cfevesting
// module state.
func MigrateParams(ctx sdk.Context, storeKey storetypes.StoreKey, legacySubspace subspace.Subspace, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	var currParams Params
	legacySubspace.GetParamSet(ctx, &currParams)

	bz, err := cdc.Marshal(&currParams)
	if err != nil {
		return err
	}

	store.Set(types.ParamsKey, bz)

	return nil
}
