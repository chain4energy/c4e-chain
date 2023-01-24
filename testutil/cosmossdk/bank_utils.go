package cosmossdk

import (
	"testing"

	testenv "github.com/chain4energy/c4e-chain/testutil/env"
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
	t                   *testing.T
	helperAccountKeeper *authkeeper.AccountKeeper
	helperBankKeeper    bankkeeper.Keeper
}

func NewBankUtils(t *testing.T, ctx sdk.Context, helperAccountKeeper *authkeeper.AccountKeeper, helperBankKeeper bankkeeper.Keeper) BankUtils {
	helperAccountKeeper.GetModuleAccount(ctx, helperModuleAccount)
	return BankUtils{t: t, helperAccountKeeper: helperAccountKeeper, helperBankKeeper: helperBankKeeper}
}

func (bu *BankUtils) AddCoinsToAccount(ctx sdk.Context, coinsToMint sdk.Coin, toAddr sdk.AccAddress) {
	mintedCoins := sdk.NewCoins(coinsToMint)
	bu.helperBankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	bu.helperBankKeeper.SendCoinsFromModuleToAccount(ctx, helperModuleAccount, toAddr, mintedCoins)
}

func (bu *BankUtils) AddDefaultDenomCoinsToAccount(ctx sdk.Context, amount sdk.Int, toAddr sdk.AccAddress) (denom string) {
	coinsToMint := sdk.NewCoin(testenv.DefaultTestDenom, amount)
	bu.AddCoinsToAccount(ctx, coinsToMint, toAddr)
	return testenv.DefaultTestDenom
}

func (bu *BankUtils) AddCoinsToModule(ctx sdk.Context, coinsToMint sdk.Coin, moduleName string) {
	mintedCoins := sdk.NewCoins(coinsToMint)
	bu.helperBankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	bu.helperBankKeeper.SendCoinsFromModuleToModule(ctx, helperModuleAccount, moduleName, mintedCoins)
}

func (bu *BankUtils) AddDefaultDenomCoinsToModule(ctx sdk.Context, amount sdk.Int, moduleName string) (denom string) {
	coinsToMint := sdk.NewCoin(testenv.DefaultTestDenom, amount)
	bu.AddCoinsToModule(ctx, coinsToMint, moduleName)
	return testenv.DefaultTestDenom
}

func (bu *BankUtils) GetModuleAccountBalanceByDenom(ctx sdk.Context, accName string, denom string) sdk.Int {
	moduleAccAddr := bu.helperAccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
	return bu.helperBankKeeper.GetBalance(ctx, moduleAccAddr, denom).Amount
}

func (bu *BankUtils) GetModuleAccountDefultDenomBalance(ctx sdk.Context, accName string) sdk.Int {
	return bu.GetModuleAccountBalanceByDenom(ctx, accName, testenv.DefaultTestDenom)
}

func (bu *BankUtils) GetAccountBalanceByDenom(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Int {
	return bu.helperBankKeeper.GetBalance(ctx, addr, denom).Amount
}

func (bu *BankUtils) GetAccountDefultDenomBalance(ctx sdk.Context, addr sdk.AccAddress) sdk.Int {
	return bu.GetAccountBalanceByDenom(ctx, addr, testenv.DefaultTestDenom)

}

func (bu *BankUtils) VerifyModuleAccountBalanceByDenom(ctx sdk.Context, accName string, denom string, expectedAmount sdk.Int) {
	moduleAccAddr := bu.helperAccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
	moduleBalance := bu.helperBankKeeper.GetBalance(ctx, moduleAccAddr, denom)
	require.Truef(bu.t, expectedAmount.Equal(moduleBalance.Amount), "expectedAmount %s <> module balance %s", expectedAmount, moduleBalance.Amount)
}

func (bu *BankUtils) VerifyModuleAccountDefultDenomBalance(ctx sdk.Context, accName string, expectedAmount sdk.Int) {
	bu.VerifyModuleAccountBalanceByDenom(ctx, accName, testenv.DefaultTestDenom, expectedAmount)
}

func (bu *BankUtils) VerifyAccountBalanceByDenom(ctx sdk.Context, addr sdk.AccAddress, denom string, expectedAmount sdk.Int) {
	balance := bu.helperBankKeeper.GetBalance(ctx, addr, denom)
	require.Truef(bu.t, expectedAmount.Equal(balance.Amount), "expectedAmount %s <> account balance %s", expectedAmount, balance.Amount)
}

func (bu *BankUtils) VerifyAccountDefultDenomBalance(ctx sdk.Context, addr sdk.AccAddress, expectedAmount sdk.Int) {
	bu.VerifyAccountBalanceByDenom(ctx, addr, testenv.DefaultTestDenom, expectedAmount)
}

func (bu *BankUtils) VerifyTotalSupplyByDenom(ctx sdk.Context, denom string, expectedAmount sdk.Int) {
	supply := bu.helperBankKeeper.GetSupply(ctx, denom).Amount
	require.Truef(bu.t, expectedAmount.Equal(supply), "expectedAmount %s <> supply %s", expectedAmount, supply)
}

func (bu *BankUtils) VerifyDefultDenomTotalSupply(ctx sdk.Context, expectedAmount sdk.Int) {
	bu.VerifyTotalSupplyByDenom(ctx, testenv.DefaultTestDenom, expectedAmount)
}

type ContextBankUtils struct {
	BankUtils
	testContext testenv.TestContext
}

func NewContextBankUtils(t *testing.T, testContext testenv.TestContext, helperAccountKeeper *authkeeper.AccountKeeper, helperBankKeeper bankkeeper.Keeper) *ContextBankUtils {
	bankUtils := NewBankUtils(t, testContext.GetContext(), helperAccountKeeper, helperBankKeeper)
	return &ContextBankUtils{BankUtils: bankUtils, testContext: testContext}
}

func (bu *ContextBankUtils) AddCoinsToAccount(coinsToMint sdk.Coin, toAddr sdk.AccAddress) {
	bu.BankUtils.AddCoinsToAccount(bu.testContext.GetContext(), coinsToMint, toAddr)
}

func (bu *ContextBankUtils) AddDefaultDenomCoinsToAccount(amount sdk.Int, toAddr sdk.AccAddress) (denom string) {
	return bu.BankUtils.AddDefaultDenomCoinsToAccount(bu.testContext.GetContext(), amount, toAddr)
}

func (bu *ContextBankUtils) AddCoinsToModule(coinsToMint sdk.Coin, moduleName string) {
	bu.BankUtils.AddCoinsToModule(bu.testContext.GetContext(), coinsToMint, moduleName)

}

func (bu *ContextBankUtils) AddDefaultDenomCoinsToModule(amount sdk.Int, moduleName string) (denom string) {
	return bu.BankUtils.AddDefaultDenomCoinsToModule(bu.testContext.GetContext(), amount, moduleName)
}

func (bu *ContextBankUtils) VerifyModuleAccountBalanceByDenom(accName string, denom string, expectedAmount sdk.Int) {
	bu.BankUtils.VerifyModuleAccountBalanceByDenom(bu.testContext.GetContext(), accName, denom, expectedAmount)
}

func (bu *ContextBankUtils) VerifyModuleAccountDefultDenomBalance(accName string, expectedAmount sdk.Int) {
	bu.BankUtils.VerifyModuleAccountDefultDenomBalance(bu.testContext.GetContext(), accName, expectedAmount)
}

func (bu *ContextBankUtils) VerifyAccountBalanceByDenom(addr sdk.AccAddress, denom string, expectedAmount sdk.Int) {
	bu.BankUtils.VerifyAccountBalanceByDenom(bu.testContext.GetContext(), addr, denom, expectedAmount)
}

func (bu *ContextBankUtils) VerifyAccountDefultDenomBalance(addr sdk.AccAddress, expectedAmount sdk.Int) {
	bu.BankUtils.VerifyAccountDefultDenomBalance(bu.testContext.GetContext(), addr, expectedAmount)
}

func (bu *ContextBankUtils) VerifyTotalSupplyByDenom(denom string, expectedAmount sdk.Int) {
	bu.BankUtils.VerifyTotalSupplyByDenom(bu.testContext.GetContext(), denom, expectedAmount)
}

func (bu *ContextBankUtils) VerifyDefultDenomTotalSupply(expectedAmount sdk.Int) {
	bu.BankUtils.VerifyDefultDenomTotalSupply(bu.testContext.GetContext(), expectedAmount)
}

func (bu *ContextBankUtils) GetModuleAccountDefultDenomBalance(accName string) sdk.Int {
	return bu.BankUtils.GetModuleAccountDefultDenomBalance(bu.testContext.GetContext(), accName)
}

func (bu *ContextBankUtils) GetAccountDefultDenomBalance(addr sdk.AccAddress) sdk.Int {
	return bu.BankUtils.GetAccountDefultDenomBalance(bu.testContext.GetContext(), addr)

}
