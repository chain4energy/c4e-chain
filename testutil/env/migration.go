package env

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
)

type AdditionalKeeperData struct {
	Cdc      *codec.ProtoCodec
	StoreKey *storetypes.KVStoreKey
	Subspace typesparams.Subspace
}
