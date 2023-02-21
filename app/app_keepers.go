package app

import (
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

)

var _ cfeupgradetypes.AppKeepers = (*App)(nil)

func (app *App) GetAccountKeeper() *authkeeper.AccountKeeper {
	return &app.AccountKeeper
}

func (app *App) GetBankKeeper() *bankkeeper.Keeper {
	return &app.BankKeeper
}

func (app *App) GetC4eVestingKeeper() *cfevestingkeeper.Keeper {
	return &app.CfevestingKeeper
}