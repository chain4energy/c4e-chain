package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bankKeeper,
	}
}

// get the minter
func (k Keeper) GetHalvingMinter(ctx sdk.Context) (minter types.HalvingMinter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.HalvingMinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &minter)
	return
}

// set the minter
func (k Keeper) SetHalvingMinter(ctx sdk.Context, minter types.HalvingMinter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.HalvingMinterKey, b)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)

	//k.bankKeeper.
}

func (k Keeper) SendCoinsToCommonAccount(ctx sdk.Context, coins sdk.Coins) error {
	k.Logger(ctx).Info("SendCoinsToCommonAccount")
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ModuleName, coins)

}
