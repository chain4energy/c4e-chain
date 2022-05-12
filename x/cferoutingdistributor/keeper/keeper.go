package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		paramstore    paramtypes.Subspace
		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

//SetRoutingDistributor set the routing distributor
func (k Keeper) SetRoutingDistributor(ctx sdk.Context, routingDistributor types.RoutingDistributor) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&routingDistributor)
	store.Set(types.RoutingDistributorKey, b)
}

// GetRoutingDistributorr get the routing distributor
func (k Keeper) GetRoutingDistributorr(ctx sdk.Context) (routingDistributor types.RoutingDistributor) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.RoutingDistributorKey)
	if b == nil {
		panic("stored routing distributor should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &routingDistributor)
	return
}

func (k Keeper) SendCoinsFromModuleToModule(ctx sdk.Context, coins sdk.Coins, moduleFrom string, moduleTo string) error {
	k.Logger(ctx).Info("Send coins from module: " + moduleFrom + " to module: " + moduleTo)
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, moduleFrom, moduleTo, coins)
}

func (k Keeper) SendCoinsFromModuleAccount(ctx sdk.Context, coins sdk.Coins, moduleFrom string, account sdk.AccAddress) error {
	k.Logger(ctx).Info("Send coins from module: " + moduleFrom + " to account: " + account.String())
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleFrom, account, coins)
}

func (k Keeper) GetAccountCoins(ctx sdk.Context, account sdk.AccAddress) sdk.Coins {
	return k.bankKeeper.GetAllBalances(ctx, account)
}

func (k Keeper) GetAccountAddressModuleAccount(ctx sdk.Context, accountName string) sdk.AccAddress {
	return k.accountKeeper.GetModuleAccount(ctx, accountName).GetAddress()
}

func (k Keeper) GetAccountCoinsForModuleAccount(ctx sdk.Context, accountName string) sdk.Coins {
	accAddress := k.GetAccountAddressModuleAccount(ctx, accountName)
	return k.GetAccountCoins(ctx, accAddress)
}
