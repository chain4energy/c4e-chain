package app

import (
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfeclaimkeeper "github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
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

func (app *App) GetDistributionKeeper() *distrkeeper.Keeper {
	return &app.DistrKeeper
}

func (app *App) GetParamKeeper() *paramskeeper.Keeper {
	return &app.ParamsKeeper
}

func (app *App) GetC4eClaimKeeper() *cfeclaimkeeper.Keeper {
	return &app.CfeclaimKeeper
}
