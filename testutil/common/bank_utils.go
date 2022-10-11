package common

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"
)

const helperModuleAccount = "helperTestAcc"

func AddHelperModuleAccountPermissions(maccPerms map[string][]string) map[string][]string {
	maccPerms[helperModuleAccount] = []string{authtypes.Burner, authtypes.Minter}
	return maccPerms
}

func AddHelperModuleAccountAddr(moduleAccAddrs map[string]bool) map[string]bool {
	moduleAccAddrs[authtypes.NewModuleAddress(helperModuleAccount).String()] = true
	return moduleAccAddrs
}

type BankUtils struct {
	t *testing.T
	helperAccountKeeper *authkeeper.AccountKeeper
	helperBankKeeper    bankkeeper.Keeper
}

func NewBankUtils(t *testing.T, ctx sdk.Context, helperAccountKeeper *authkeeper.AccountKeeper, helperBankKeeper bankkeeper.Keeper) *BankUtils {
	helperAccountKeeper.GetModuleAccount(ctx, helperModuleAccount)
	return &BankUtils{t: t, helperAccountKeeper: helperAccountKeeper, helperBankKeeper: helperBankKeeper}
}

func (bu *BankUtils) AddCoinsToAccount(ctx sdk.Context, coinsToMint sdk.Coin, toAddr sdk.AccAddress) {
	mintedCoins := sdk.NewCoins(coinsToMint)
	bu.helperBankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	bu.helperBankKeeper.SendCoinsFromModuleToAccount(ctx, helperModuleAccount, toAddr, mintedCoins)
}

func (bu *BankUtils) AddDefaultDenomCoinsToAccount(ctx sdk.Context, amount sdk.Int, toAddr sdk.AccAddress) (denom string) {
	coinsToMint := sdk.NewCoin(DefaultTestDenom, amount)
	bu.AddCoinsToAccount(ctx, coinsToMint, toAddr)
	return DefaultTestDenom
}

func (bu *BankUtils) AddCoinsToModule(ctx sdk.Context, coinsToMint sdk.Coin, moduleName string) {
	mintedCoins := sdk.NewCoins(coinsToMint)
	bu.helperBankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	bu.helperBankKeeper.SendCoinsFromModuleToModule(ctx, helperModuleAccount, moduleName, mintedCoins)
}

func (bu *BankUtils) AddDefaultDenomCoinsToModule(ctx sdk.Context, amount sdk.Int, moduleName string) (denom string) {
	coinsToMint := sdk.NewCoin(DefaultTestDenom, amount)
	bu.AddCoinsToModule(ctx, coinsToMint, moduleName)
	return DefaultTestDenom
}

func (bu *BankUtils) VerifyModuleAccountBalanceByDenom(ctx sdk.Context, accName string, denom string, expectedAmount sdk.Int) {
	moduleAccAddr := bu.helperAccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
	moduleBalance := bu.helperBankKeeper.GetBalance(ctx, moduleAccAddr, denom)
	require.Truef(bu.t, expectedAmount.Equal(moduleBalance.Amount), "expectedAmount %s <> module balance %s", expectedAmount, moduleBalance.Amount)
}

func (bu *BankUtils) VerifyModuleAccountDefultDenomBalance(ctx sdk.Context, accName string, expectedAmount sdk.Int) {
	bu.VerifyModuleAccountBalanceByDenom(ctx, accName, DefaultTestDenom, expectedAmount)
}

func (bu *BankUtils) VerifyAccountBalanceByDenom(ctx sdk.Context, addr sdk.AccAddress, denom string, expectedAmount sdk.Int) {
	balance := bu.helperBankKeeper.GetBalance(ctx, addr, denom)
	require.Truef(bu.t, expectedAmount.Equal(balance.Amount), "expectedAmount %s <> account balance %s", expectedAmount, balance.Amount)
}

func (bu *BankUtils) VerifyAccountDefultDenomBalance(ctx sdk.Context, addr sdk.AccAddress, expectedAmount sdk.Int) {
	bu.VerifyAccountBalanceByDenom(ctx, addr, DefaultTestDenom, expectedAmount)
}







// func VerifyModuleAccountBalanceByName(accName string, ctx sdk.Context, app *app.App, t *testing.T, expectedAmount sdk.Int) {
// 	VerifyModuleAccountDenomBalanceByName(accName, ctx, app, t, DefaultTestDenom, expectedAmount)
// }

// func VerifyModuleAccountDenomBalanceByName(accName string, ctx sdk.Context, app *app.App, t *testing.T, denom string, expectedAmount sdk.Int) {
// 	moduleAccAddr := app.AccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
// 	moduleBalance := app.BankKeeper.GetBalance(ctx, moduleAccAddr, denom)
// 	require.EqualValues(t, expectedAmount.String(), moduleBalance.Amount.String())
// }

// func AddCoinsToAccount(vested uint64, ctx sdk.Context, app *app.App, toAddr sdk.AccAddress) string {
// 	return AddCoinsToAccountInt(sdk.NewIntFromUint64(vested), ctx, app, toAddr)
// }

// func AddCoinsToAccountInt(amount sdk.Int, ctx sdk.Context, app *app.App, toAddr sdk.AccAddress) string {

// 	mintedCoin := sdk.NewCoin(DefaultTestDenom, amount)
// 	mintedCoins := sdk.NewCoins(mintedCoin)
// 	app.BankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
// 	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, helperModuleAccount, toAddr, mintedCoins)
// 	return DefaultTestDenom
// }

// func AddCoinsToModuleByName(vested uint64, modulaName string, ctx sdk.Context, app *app.App) string {
// 	mintedCoin := sdk.NewCoin(DefaultTestDenom, sdk.NewIntFromUint64(vested))
// 	mintedCoins := sdk.NewCoins(mintedCoin)
// 	app.BankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
// 	app.BankKeeper.SendCoinsFromModuleToModule(ctx, helperModuleAccount, modulaName, mintedCoins)
// 	return DefaultTestDenom
// }
