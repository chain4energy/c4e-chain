package v3

import (
	"github.com/chain4energy/c4e-chain/types/subspace"
	v2 "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v2"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var ParamsKey = []byte{0x00}

// MigrateParams migrates the x/cfedistributor module state from the consensus version 2 to
// version 3. Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/cfedistributor
// module state.
func MigrateParams(ctx sdk.Context, storeKey storetypes.StoreKey, legacySubspace subspace.Subspace, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	var currParams v2.Params
	legacySubspace.GetParamSet(ctx, &currParams)

	bz, err := cdc.Marshal(&currParams)
	if err != nil {
		return err
	}

	store.Set(ParamsKey, bz)

	return nil
}
