package app

import (
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfeairdropkeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

var _ cfeupgradetypes.AppKeepers = (*App)(nil)

func (app *App) GetAirdropKeeper() *cfeairdropkeeper.Keeper {
	return &app.CfeairdropKeeper
}

func (app *App) GetAccountKeeper() *authkeeper.AccountKeeper {
	return &app.AccountKeeper
}
