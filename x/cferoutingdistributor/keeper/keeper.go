package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"

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

////SetRoutingDistributor set the routing distributor
//func (k Keeper) SetRoutingDistributor(ctx sdk.Context, routingDistributor types.RoutingDistributor) {
//	store := ctx.KVStore(k.storeKey)
//	b := k.cdc.MustMarshal(&routingDistributor)
//	store.Set(types.RoutingDistributorKey, b)
//}

//func (k Keeper) SetRemainsMap(ctx sdk.Context, remainsMap map[string]types.Remains) {
//	store := ctx.KVStore(k.storeKey)
//	b := k.cdc.MustMarshal(&remainsMap)
//	store.Set(types.RemainsMapKey, b)
//}
//
//// GetRoutingDistributorr get the routing distributor
//func (k Keeper) GetRemainsMap(ctx sdk.Context) (remainsMap map[string]types.Remains) {
//	store := ctx.KVStore(k.storeKey)
//	b := store.Get(types.RemainsMapKey)
//	if b == nil {
//		panic("stored remains map should not have been nil")
//	}
//
//	k.cdc.MustUnmarshal(b, &remainsMap)
//	return
//}

// get the vesting types
func (k Keeper) GetRemains(ctx sdk.Context, accountAddress string) (remains types.Remains, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RemainsKeyPrefix)

	b := store.Get([]byte(accountAddress))
	if b == nil {
		found = false
		return
	}
	found = true
	k.cdc.MustUnmarshal(b, &remains)
	return
}

// GetAllAccountVestings returns all AccountVestings
func (k Keeper) GetAllRemains(ctx sdk.Context) (list []types.Remains) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RemainsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Remains
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// set the vesting types
func (k Keeper) SetRemains(ctx sdk.Context, remains types.Remains) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RemainsKeyPrefix)
	av := k.cdc.MustMarshal(&remains)
	store.Set([]byte(remains.Account.Address), av)
}

//// GetRoutingDistributorr get the routing distributor
//func (k Keeper) GetRoutingDistributorr(ctx sdk.Context) (routingDistributor types.RoutingDistributor) {
//	store := ctx.KVStore(k.storeKey)
//	b := store.Get(types.RoutingDistributorKey)
//	if b == nil {
//		panic("stored routing distributor should not have been nil")
//	}
//
//	k.cdc.MustUnmarshal(b, &routingDistributor)
//	return
//}

func (k Keeper) SendCoinsFromModuleToModule(ctx sdk.Context, coins sdk.Coins, moduleFrom string, moduleTo string) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, moduleFrom, moduleTo, coins)
}

func (k Keeper) SendCoinsFromModuleAccount(ctx sdk.Context, coins sdk.Coins, moduleFrom string, account sdk.AccAddress) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleFrom, account, coins)
}

func (k Keeper) SendCoinsToModuleAccount(ctx sdk.Context, coins sdk.Coins, account sdk.AccAddress, moduleTo string) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, account, moduleTo, coins)
}

func (k Keeper) BurnCoinsForSpecifiedModuleAccount(ctx sdk.Context, coins sdk.Coins, moduleAccountName string) error {
	return k.bankKeeper.BurnCoins(ctx, moduleAccountName, coins)

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

