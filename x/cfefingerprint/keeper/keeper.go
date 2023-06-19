package keeper

import (
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		proto      codec.JSONCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		authKeeper types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	proto codec.JSONCodec,
	storeKey,
	memKey storetypes.StoreKey,
	authKeeper types.AccountKeeper,
) *Keeper {
	return &Keeper{
		cdc:        cdc,
		proto:      proto,
		storeKey:   storeKey,
		memKey:     memKey,
		authKeeper: authKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
