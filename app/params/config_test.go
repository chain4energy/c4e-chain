package params_test

import (
	"github.com/chain4energy/c4e-chain/v2/testutil/app"
	"github.com/chain4energy/c4e-chain/v2/testutil/cosmossdk"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDenoms(t *testing.T) {
	_ = app.SetupTestAppWithHeight(t, 1000)

	c4eCoin := sdk.NewCoin("c4e", sdk.NewInt(1))
	uc4eCoin := sdk.NewCoin("uc4e", sdk.NewInt(1000000))
	parsedCoin, _ := sdk.ParseCoinNormalized(c4eCoin.String())

	require.True(t, parsedCoin.Equal(uc4eCoin))
}

func TestDenomsAbcd(t *testing.T) {
	_ = app.SetupTestAppWithHeight(t, 1000)

	c4eCoin := sdk.NewCoin("abcd", sdk.NewInt(1))
	uc4eCoin := sdk.NewCoin("uc4e", sdk.NewInt(1000000))
	parsedCoin, _ := sdk.ParseCoinNormalized(c4eCoin.String())

	require.False(t, parsedCoin.Equal(uc4eCoin))
}

func TestDenomsBank(t *testing.T) {
	testhelper := app.SetupTestAppWithHeight(t, 1000)

	c4eCoin := sdk.NewCoin("c4e", sdk.NewInt(1))
	accAddress := sdk.MustAccAddressFromBech32(cosmossdk.CreateRandomAccAddress())
	testhelper.BankUtils.AddCoinsToAccount(sdk.NewCoins(c4eCoin), accAddress)

	uc4eCoin := sdk.NewCoin("uc4e", sdk.NewInt(1000000))
	balances := testhelper.BankUtils.GetAccountAllBalances(accAddress)
	require.Panics(t, func() { balances.IsEqual(sdk.NewCoins(uc4eCoin)) })
}
