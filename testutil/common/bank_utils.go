package common

import (
	"testing"

	"github.com/chain4energy/c4e-chain/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func VerifyModuleAccountBalanceByName(accName string, ctx sdk.Context, app *app.App, t *testing.T, expectedAmount sdk.Int) {
	VerifyModuleAccountDenomBalanceByName(accName, ctx, app, t, Denom, expectedAmount)
}

func VerifyModuleAccountDenomBalanceByName(accName string, ctx sdk.Context, app *app.App, t *testing.T, denom string, expectedAmount sdk.Int) {
	moduleAccAddr := app.AccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
	moduleBalance := app.BankKeeper.GetBalance(ctx, moduleAccAddr, denom)
	require.EqualValues(t, expectedAmount, moduleBalance.Amount)
}

