package cosmossdk

import (
	"cosmossdk.io/math"

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
	t                   require.TestingT
	helperAccountKeeper *authkeeper.AccountKeeper
	helperBankKeeper    bankkeeper.Keeper
}

func (bu *BankUtils) VerifyAccountDefultDenomSpendableCoins(ctx sdk.Context, addr sdk.AccAddress, expectedAmount math.Int) {
	bu.VerifyAccountSpendableByDenom(ctx, addr, testenv.DefaultTestDenom, expectedAmount)
}

func (bu *BankUtils) VerifyAccountSpendableByDenom(ctx sdk.Context, addr sdk.AccAddress, denom string, expectedAmount math.Int) {
	locked := bu.helperBankKeeper.SpendableCoins(ctx, addr).AmountOf(denom)
	require.Truef(bu.t, expectedAmount.Equal(locked), "expectedAmount %s <> spendable amount %s", expectedAmount, locked)
}

func NewBankUtils(t require.TestingT, ctx sdk.Context, helperAccountKeeper *authkeeper.AccountKeeper, helperBankKeeper bankkeeper.Keeper) BankUtils {
	helperAccountKeeper.GetModuleAccount(ctx, helperModuleAccount)
	return BankUtils{t: t, helperAccountKeeper: helperAccountKeeper, helperBankKeeper: helperBankKeeper}
}

func (bu *BankUtils) AddCoinsToAccount(ctx sdk.Context, coinsToMint sdk.Coins, toAddr sdk.AccAddress) {
	// mintedCoins := sdk.NewCoins(coinsToMint)
	bu.helperBankKeeper.MintCoins(ctx, helperModuleAccount, coinsToMint)
	bu.helperBankKeeper.SendCoinsFromModuleToAccount(ctx, helperModuleAccount, toAddr, coinsToMint)
}

func (bu *BankUtils) AddDefaultDenomCoinsToAccount(ctx sdk.Context, amount math.Int, toAddr sdk.AccAddress) (denom string) {
	coinsToMint := sdk.NewCoin(testenv.DefaultTestDenom, amount)
	mintedCoins := sdk.NewCoins(coinsToMint)
	bu.AddCoinsToAccount(ctx, mintedCoins, toAddr)
	return testenv.DefaultTestDenom
}

func (bu *BankUtils) AddCoinsToModule(ctx sdk.Context, coinsToMint sdk.Coin, moduleName string) {
	mintedCoins := sdk.NewCoins(coinsToMint)
	bu.helperBankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	bu.helperBankKeeper.SendCoinsFromModuleToModule(ctx, helperModuleAccount, moduleName, mintedCoins)
}

func (bu *BankUtils) AddDefaultDenomCoinsToModule(ctx sdk.Context, amount math.Int, moduleName string) (denom string) {
	coinsToMint := sdk.NewCoin(testenv.DefaultTestDenom, amount)
	bu.AddCoinsToModule(ctx, coinsToMint, moduleName)
	return testenv.DefaultTestDenom
}

func (bu *BankUtils) GetModuleAccountBalanceByDenom(ctx sdk.Context, accName string, denom string) math.Int {
	moduleAccAddr := bu.helperAccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
	return bu.helperBankKeeper.GetBalance(ctx, moduleAccAddr, denom).Amount
}

func (bu *BankUtils) GetModuleAccountAllBalances(ctx sdk.Context, accName string) sdk.Coins {
	moduleAccAddr := bu.helperAccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
	return bu.helperBankKeeper.GetAllBalances(ctx, moduleAccAddr)
}

func (bu *BankUtils) GetModuleAccountDefultDenomBalance(ctx sdk.Context, accName string) math.Int {
	return bu.GetModuleAccountBalanceByDenom(ctx, accName, testenv.DefaultTestDenom)
}

func (bu *BankUtils) GetAccountBalanceByDenom(ctx sdk.Context, addr sdk.AccAddress, denom string) math.Int {
	return bu.helperBankKeeper.GetBalance(ctx, addr, denom).Amount
}

func (bu *BankUtils) GetAccountDefultDenomBalance(ctx sdk.Context, addr sdk.AccAddress) math.Int {
	return bu.GetAccountBalanceByDenom(ctx, addr, testenv.DefaultTestDenom)
}

func (bu *BankUtils) GetAccountAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return bu.helperBankKeeper.GetAllBalances(ctx, addr)
}

func (bu *BankUtils) GetAccountLockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return bu.helperBankKeeper.LockedCoins(ctx, addr)
}

func (bu *BankUtils) VerifyModuleAccountBalanceByDenom(ctx sdk.Context, accName string, denom string, expectedAmount math.Int) {
	moduleAccAddr := bu.helperAccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
	moduleBalance := bu.helperBankKeeper.GetBalance(ctx, moduleAccAddr, denom)
	require.Truef(bu.t, expectedAmount.Equal(moduleBalance.Amount), "expectedAmount %s <> module balance %s", expectedAmount, moduleBalance.Amount)
}

func (bu *BankUtils) VerifyModuleAccountAllBalances(ctx sdk.Context, accName string, expectedAmount sdk.Coins) {
	moduleAccAddr := bu.helperAccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
	moduleBalance := bu.helperBankKeeper.GetAllBalances(ctx, moduleAccAddr)
	require.Truef(bu.t, expectedAmount.IsEqual(moduleBalance), "expectedAmount %s <> module balance %s", expectedAmount, moduleBalance)
}

func (bu *BankUtils) VerifyAccountDefultDenomLocked(ctx sdk.Context, addr sdk.AccAddress, expectedAmount math.Int) {
	bu.VerifyAccountLockedByDenom(ctx, addr, testenv.DefaultTestDenom, expectedAmount)
}

func (bu *ContextBankUtils) DisableDefaultSend() {
	bu.BankUtils.DisableDefaultSend(bu.testContext.GetContext())
}

func (bu *BankUtils) DisableDefaultSend(ctx sdk.Context) {
	params := bu.helperBankKeeper.GetParams(ctx)
	params.DefaultSendEnabled = false
	bu.helperBankKeeper.SetParams(ctx, params)
}

func (bu *BankUtils) VerifyAccountLockedByDenom(ctx sdk.Context, addr sdk.AccAddress, denom string, expectedAmount math.Int) {
	locked := bu.helperBankKeeper.LockedCoins(ctx, addr).AmountOf(denom)
	require.Truef(bu.t, expectedAmount.Equal(locked), "expectedAmount %s <> account locked %s", expectedAmount, locked)
}

func (bu *BankUtils) VerifyModuleAccountDefultDenomBalance(ctx sdk.Context, accName string, expectedAmount math.Int) {
	bu.VerifyModuleAccountBalanceByDenom(ctx, accName, testenv.DefaultTestDenom, expectedAmount)
}

func (bu *BankUtils) VerifyAccountBalanceByDenom(ctx sdk.Context, addr sdk.AccAddress, denom string, expectedAmount math.Int) {
	balance := bu.helperBankKeeper.GetBalance(ctx, addr, denom)
	require.Truef(bu.t, expectedAmount.Equal(balance.Amount), "expectedAmount %s <> account balance %s", expectedAmount, balance.Amount)
}

func (bu *BankUtils) VerifyAccountBalances(ctx sdk.Context, addr sdk.AccAddress, expectedBalances sdk.Coins, isAllBalances bool) {
	balances := bu.helperBankKeeper.GetAllBalances(ctx, addr)
	if isAllBalances {
		require.Truef(bu.t, expectedBalances.IsEqual(balances), "expectedBalances %s <> account balances %s", expectedBalances, balances)

	} else {
		for _, expectedBalance := range expectedBalances {
			require.Truef(bu.t, expectedBalance.Amount.Equal(balances.AmountOf(expectedBalance.Denom)), "expectedBalance %s <> account balance %s for denom %s", expectedBalance.Amount, balances.AmountOf(expectedBalance.Denom), expectedBalance.Denom)
		}
	}
}

func (bu *BankUtils) VerifyAccountAllBalances(ctx sdk.Context, addr sdk.AccAddress, expectedBalances sdk.Coins) {
	balances := bu.helperBankKeeper.GetAllBalances(ctx, addr)
	require.Truef(bu.t, expectedBalances.IsEqual(balances), "expectedBalances %s <> account balances %s", expectedBalances, balances)
}

func (bu *BankUtils) VerifyLockedCoins(ctx sdk.Context, addr sdk.AccAddress, expectedLockedCoins sdk.Coins, isAllLocked bool) {
	balances := bu.helperBankKeeper.LockedCoins(ctx, addr)
	if isAllLocked {
		require.Truef(bu.t, expectedLockedCoins.IsEqual(balances), "expectedLockedCoins %s <> account locked coins %s", expectedLockedCoins, balances)

	} else {
		for _, expectedLockedCoin := range expectedLockedCoins {
			require.Truef(bu.t, expectedLockedCoin.Amount.Equal(balances.AmountOf(expectedLockedCoin.Denom)), "expectedLockedCoin %s <> account locked coins %s for denom %s", expectedLockedCoin.Amount, balances.AmountOf(expectedLockedCoin.Denom), expectedLockedCoin.Denom)
		}
	}
}

func (bu *BankUtils) VerifyAccountDefaultDenomBalance(ctx sdk.Context, addr sdk.AccAddress, expectedAmount math.Int) {
	bu.VerifyAccountBalanceByDenom(ctx, addr, testenv.DefaultTestDenom, expectedAmount)
}

func (bu *BankUtils) VerifyTotalSupplyByDenom(ctx sdk.Context, denom string, expectedAmount math.Int) {
	supply := bu.helperBankKeeper.GetSupply(ctx, denom).Amount
	require.Truef(bu.t, expectedAmount.Equal(supply), "expectedAmount %s <> supply %s", expectedAmount, supply)
}

func (bu *BankUtils) VerifyDefultDenomTotalSupply(ctx sdk.Context, expectedAmount math.Int) {
	bu.VerifyTotalSupplyByDenom(ctx, testenv.DefaultTestDenom, expectedAmount)
}

func (bu *BankUtils) DisableSend(ctx sdk.Context) {
	params := bu.helperBankKeeper.GetParams(ctx)
	params.DefaultSendEnabled = false
	bu.helperBankKeeper.SetParams(ctx, params)
}

type ContextBankUtils struct {
	BankUtils
	testContext testenv.TestContext
}

func NewContextBankUtils(t require.TestingT, testContext testenv.TestContext, helperAccountKeeper *authkeeper.AccountKeeper, helperBankKeeper bankkeeper.Keeper) *ContextBankUtils {
	bankUtils := NewBankUtils(t, testContext.GetContext(), helperAccountKeeper, helperBankKeeper)
	return &ContextBankUtils{BankUtils: bankUtils, testContext: testContext}
}

func (bu *ContextBankUtils) AddCoinsToAccount(coinsToMint sdk.Coins, toAddr sdk.AccAddress) {
	bu.BankUtils.AddCoinsToAccount(bu.testContext.GetContext(), coinsToMint, toAddr)
}

func (bu *ContextBankUtils) AddDefaultDenomCoinsToAccount(amount math.Int, toAddr sdk.AccAddress) (denom string) {
	return bu.BankUtils.AddDefaultDenomCoinsToAccount(bu.testContext.GetContext(), amount, toAddr)
}

func (bu *ContextBankUtils) AddCoinsToModule(coinsToMint sdk.Coin, moduleName string) {
	bu.BankUtils.AddCoinsToModule(bu.testContext.GetContext(), coinsToMint, moduleName)

}

func (bu *ContextBankUtils) AddDefaultDenomCoinsToModule(amount math.Int, moduleName string) (denom string) {
	return bu.BankUtils.AddDefaultDenomCoinsToModule(bu.testContext.GetContext(), amount, moduleName)
}

func (bu *ContextBankUtils) VerifyModuleAccountBalanceByDenom(accName string, denom string, expectedAmount math.Int) {
	bu.BankUtils.VerifyModuleAccountBalanceByDenom(bu.testContext.GetContext(), accName, denom, expectedAmount)
}

func (bu *ContextBankUtils) VerifyModuleAccountDefultDenomBalance(accName string, expectedAmount math.Int) {
	bu.BankUtils.VerifyModuleAccountDefultDenomBalance(bu.testContext.GetContext(), accName, expectedAmount)
}

func (bu *ContextBankUtils) VerifyAccountBalanceByDenom(addr sdk.AccAddress, denom string, expectedAmount math.Int) {
	bu.BankUtils.VerifyAccountBalanceByDenom(bu.testContext.GetContext(), addr, denom, expectedAmount)
}

func (bu *ContextBankUtils) VerifyAccountDefultDenomBalance(addr sdk.AccAddress, expectedAmount math.Int) {
	bu.BankUtils.VerifyAccountDefaultDenomBalance(bu.testContext.GetContext(), addr, expectedAmount)
}

func (bu *ContextBankUtils) VerifyTotalSupplyByDenom(denom string, expectedAmount math.Int) {
	bu.BankUtils.VerifyTotalSupplyByDenom(bu.testContext.GetContext(), denom, expectedAmount)
}

func (bu *ContextBankUtils) VerifyDefultDenomTotalSupply(expectedAmount math.Int) {
	bu.BankUtils.VerifyDefultDenomTotalSupply(bu.testContext.GetContext(), expectedAmount)
}

func (bu *ContextBankUtils) GetModuleAccountDefultDenomBalance(accName string) math.Int {
	return bu.BankUtils.GetModuleAccountDefultDenomBalance(bu.testContext.GetContext(), accName)
}

func (bu *ContextBankUtils) GetAccountDefultDenomBalance(addr sdk.AccAddress) math.Int {
	return bu.BankUtils.GetAccountDefultDenomBalance(bu.testContext.GetContext(), addr)

}

func (bu *ContextBankUtils) VerifyAccountBalances(addr sdk.AccAddress, expectedBalances sdk.Coins, isAllBalances bool) {
	bu.BankUtils.VerifyAccountBalances(bu.testContext.GetContext(), addr, expectedBalances, isAllBalances)
}

func (bu *ContextBankUtils) VerifyLockedCoins(addr sdk.AccAddress, expectedLockedCoins sdk.Coins, isAllLocked bool) {
	bu.BankUtils.VerifyLockedCoins(bu.testContext.GetContext(), addr, expectedLockedCoins, isAllLocked)
}

func (bu *ContextBankUtils) GetAccountAllBalances(addr sdk.AccAddress) sdk.Coins {
	return bu.BankUtils.GetAccountAllBalances(bu.testContext.GetContext(), addr)
}

func (bu *ContextBankUtils) GetAccountLockedCoins(addr sdk.AccAddress) sdk.Coins {
	return bu.BankUtils.GetAccountLockedCoins(bu.testContext.GetContext(), addr)
}

func (bu *ContextBankUtils) DisableSend() {
	bu.BankUtils.DisableSend(bu.testContext.GetContext())

}
