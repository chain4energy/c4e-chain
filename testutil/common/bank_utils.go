package common

import (
	"testing"

	"github.com/chain4energy/c4e-chain/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

const helperModuleAccount = "helperTestAcc"

func AddHelperModuleAccountPerms() {
	perms := []string{authtypes.Minter}
	app.AddMaccPerms(helperModuleAccount, perms)
}

func VerifyModuleAccountBalanceByName(accName string, ctx sdk.Context, app *app.App, t *testing.T, expectedAmount sdk.Int) {
	VerifyModuleAccountDenomBalanceByName(accName, ctx, app, t, Denom, expectedAmount)
}

func VerifyModuleAccountDenomBalanceByName(accName string, ctx sdk.Context, app *app.App, t *testing.T, denom string, expectedAmount sdk.Int) {
	moduleAccAddr := app.AccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
	moduleBalance := app.BankKeeper.GetBalance(ctx, moduleAccAddr, denom)
	require.EqualValues(t, expectedAmount, moduleBalance.Amount)
}

func AddCoinsToAccount(vested uint64, ctx sdk.Context, app *app.App, toAddr sdk.AccAddress) string {

	mintedCoin := sdk.NewCoin(Denom, sdk.NewIntFromUint64(vested))
	mintedCoins := sdk.NewCoins(mintedCoin)
	app.BankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, helperModuleAccount, toAddr, mintedCoins)
	return Denom
}

func AddCoinsToModuleByName(vested uint64, modulaName string, ctx sdk.Context, app *app.App) string {
	mintedCoin := sdk.NewCoin(Denom, sdk.NewIntFromUint64(vested))
	mintedCoins := sdk.NewCoins(mintedCoin)
	app.BankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	app.BankKeeper.SendCoinsFromModuleToModule(ctx, helperModuleAccount, modulaName, mintedCoins)
	return Denom
}
