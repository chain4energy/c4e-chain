package keeper

import (
	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		authority     string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	authority string,
) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		authority:     authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

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
