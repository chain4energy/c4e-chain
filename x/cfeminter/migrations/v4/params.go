package v4

import (
	v3 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v3"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var ParamsKey = []byte{0x00}

func MigrateParams(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	var oldParams v3.Params
	bz := store.Get(ParamsKey)

	cdc.MustUnmarshal(bz, &oldParams)
	for _, oldMinter := range oldParams.Minters {
		if oldMinter.Config.TypeUrl == "/chain4energy.c4echain.cfeminter.ExponentialStepMinting" {
			oldMinter.Config.TypeUrl = "/c4echain.cfeminter.ExponentialStepMinting"
		} else if oldMinter.Config.TypeUrl == "/chain4energy.c4echain.cfeminter.LinearMinting" {
			oldMinter.Config.TypeUrl = "/c4echain.cfeminter.LinearMinting"
		} else if oldMinter.Config.TypeUrl == "/chain4energy.c4echain.cfeminter.NoMinting" {
			oldMinter.Config.TypeUrl = "/c4echain.cfeminter.NoMinting"
		}
	}

	bz, err := cdc.Marshal(&oldParams)
	if err != nil {
		return err
	}

	store.Set(ParamsKey, bz)

	return nil
}
