package app

import (
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

var _ cfeupgradetypes.AppKeepers = (*App)(nil)

func (app *App) GetAccountKeeper() *authkeeper.AccountKeeper {
	return &app.AccountKeeper
}

func (app *App) GetBankKeeper() *bankkeeper.Keeper {
	return &app.BankKeeper
}
