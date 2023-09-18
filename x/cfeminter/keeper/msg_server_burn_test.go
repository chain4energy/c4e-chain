package keeper_test

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/v2/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/v2/testutil/cosmossdk"
	"github.com/chain4energy/c4e-chain/v2/testutil/sample"
	"testing"
)

func TestBurnAllCoins(t *testing.T) {
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]
	coins := sample.PrepareDifferentDenomCoins(10, math.NewInt(10000))
	testHelper.BankUtils.AddCoinsToAccount(coins, accAddr)

	testHelper.C4eMinterUtils.MessageBurn(accAddr.String(), coins)
}

func TestBurnAllCoinsDiff(t *testing.T) {
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]
	coins := sample.PrepareDifferentDenomCoins(6, math.NewInt(10000))
	testHelper.BankUtils.AddCoinsToAccount(coins[:3], accAddr)
	testHelper.C4eMinterUtils.MessageBurnError(accAddr.String(), coins, "balance is too small (10000uc4e0,10000uc4e1,10000uc4e2 < 10000uc4e0,10000uc4e1,10000uc4e2,10000uc4e3,10000uc4e4,10000uc4e5): insufficient funds")
}

func TestBurnHalfCoins(t *testing.T) {
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]
	coins := sample.PrepareDifferentDenomCoins(10, math.NewInt(10000))
	testHelper.BankUtils.AddCoinsToAccount(coins, accAddr)

	testHelper.C4eMinterUtils.MessageBurn(accAddr.String(), coins[:5])
}

func TestBurnOneCoinsHalfAmount(t *testing.T) {
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]
	coins := sample.PrepareDifferentDenomCoins(1, math.NewInt(10000))
	testHelper.BankUtils.AddCoinsToAccount(coins, accAddr)

	coins[0].Amount = math.NewInt(5000)
	testHelper.C4eMinterUtils.MessageBurn(accAddr.String(), coins)
}

func TestBurnOneCoins(t *testing.T) {
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]
	coins := sample.PrepareDifferentDenomCoins(1, math.NewInt(10000))
	testHelper.BankUtils.AddCoinsToAccount(coins, accAddr)

	testHelper.C4eMinterUtils.MessageBurn(accAddr.String(), coins)
}

func TestBurnAmountTooBig(t *testing.T) {
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]
	coins := sample.PrepareDifferentDenomCoins(3, math.NewInt(10000))
	testHelper.BankUtils.AddCoinsToAccount(coins, accAddr)
	coins[2].Amount = math.NewInt(100000000)
	testHelper.C4eMinterUtils.MessageBurnError(accAddr.String(), coins, "balance is too small (10000uc4e0,10000uc4e1,10000uc4e2 < 10000uc4e0,10000uc4e1,100000000uc4e2): insufficient funds")
}

func TestBurnAmountNegative(t *testing.T) {
	testHelper := app.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]
	coins := sample.PrepareDifferentDenomCoins(2, math.NewInt(10000))
	testHelper.BankUtils.AddCoinsToAccount(coins, accAddr)
	coins[1].Amount = math.NewInt(-10000)
	testHelper.C4eMinterUtils.MessageBurnError(accAddr.String(), coins, "amount is not positive: wrong param value")
}
