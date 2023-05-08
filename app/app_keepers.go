package app

import (
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfeairdropkeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
)

var _ cfeupgradetypes.AppKeepers = (*App)(nil)

func (app *App) GetC4eAirdropKeeper() *cfeairdropkeeper.Keeper {
	return &app.CfeairdropKeeper
}

func (app *App) GetAccountKeeper() *authkeeper.AccountKeeper {
	return &app.AccountKeeper
}

func (app *App) GetBankKeeper() *bankkeeper.Keeper {
	return &app.BankKeeper
}

func (app *App) GetC4eVestingKeeper() *cfevestingkeeper.Keeper {
	return &app.CfevestingKeeper
}

func (app *App) GetParamKeeper() *paramskeeper.Keeper {
	return &app.ParamsKeeper
}
