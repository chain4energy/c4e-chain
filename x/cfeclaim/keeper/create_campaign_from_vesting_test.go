package keeper_test

import (
	"cosmossdk.io/math"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
	"time"
)

func TestAddClaimRecordsFromWhitelistedAccountAllCoinsLocked(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	vestingCoins := sdk.Coins{{Amount: math.NewInt(1000), Denom: testenv.DefaultTestDenom}}
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		vestingCoins,
		startTime,
		endTime,
		accBalance,
	)

	testHelper.C4eClaimUtils.AddClaimRecordsFromWhitelistedVestingAccount(accAddr2, math.NewInt(1000), math.NewInt(1000))
}

func TestAddClaimRecordsFromWhitelistedAccountAllCoinsUnlocked(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	vestingCoins := sdk.Coins{{Amount: math.NewInt(1000), Denom: testenv.DefaultTestDenom}}
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		vestingCoins,
		startTime,
		endTime,
		accBalance,
	)

	testHelper.SetContextBlockTime(endTime)
	testHelper.C4eClaimUtils.AddClaimRecordsFromWhitelistedVestingAccount(accAddr2, math.NewInt(500), math.NewInt(1000))
}

func TestAddClaimRecordsFromWhitelistedAccountHalfOfCoinsUnlocked(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	vestingCoins := sdk.Coins{{Amount: math.NewInt(1000), Denom: testenv.DefaultTestDenom}}
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		vestingCoins,
		startTime,
		endTime,
		accBalance,
	)

	testHelper.SetContextBlockTime(startTime.Add(time.Hour * 5))
	testHelper.C4eClaimUtils.AddClaimRecordsFromWhitelistedVestingAccount(accAddr2, math.NewInt(500), math.NewInt(500))
}

func TestAddClaimRecordsFromWhitelistedAccountHalfOfCoinsUnlockedBiggerAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	vestingCoins := sdk.Coins{{Amount: math.NewInt(1000), Denom: testenv.DefaultTestDenom}}
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		vestingCoins,
		startTime,
		endTime,
		accBalance,
	)

	testHelper.SetContextBlockTime(startTime.Add(time.Hour * 5))
	testHelper.C4eClaimUtils.AddClaimRecordsFromWhitelistedVestingAccount(accAddr2, math.NewInt(777), math.NewInt(777))
}

func TestAddClaimRecordsFromWhitelistedAccountHalfOfCoinsUnlockedAdditionalBalance(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(10, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := math.NewInt(1000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	vestingCoins := sdk.Coins{{Amount: math.NewInt(1000), Denom: testenv.DefaultTestDenom}}
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		vestingCoins,
		startTime,
		endTime,
		accBalance,
	)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(accAddr2, math.NewInt(100))
	testHelper.SetContextBlockTime(startTime.Add(time.Hour * 5))
	testHelper.C4eClaimUtils.AddClaimRecordsFromWhitelistedVestingAccount(accAddr2, math.NewInt(777), math.NewInt(777))
}

func TestAddClaimRecordsFromWhitelistedAccountHalfOfCoinsUnlockedDelegate(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorsAddresses := testcosmos.CreateAccounts(10, 1)
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, math.NewInt(100))
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := math.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	vestingCoins := sdk.Coins{{Amount: math.NewInt(10000), Denom: testenv.DefaultTestDenom}}
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		vestingCoins,
		startTime,
		endTime,
		accBalance,
	)
	testHelper.StakingUtils.MessageDelegate(2, 0, validatorsAddresses[0], accAddr2, math.NewInt(3000))
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(accAddr2, math.NewInt(100))

	testHelper.SetContextBlockTime(startTime.Add(time.Hour * 5))
	testHelper.C4eClaimUtils.AddClaimRecordsFromWhitelistedVestingAccount(accAddr2, math.NewInt(6000), math.NewInt(6100))
}

func TestAddClaimRecordsFromWhitelistedAccountHalfOfCoinsUnlockedDelegateAndUndelegate(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, validatorsAddresses := testcosmos.CreateAccounts(10, 1)
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, math.NewInt(100))
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accBalance := math.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	vestingCoins := sdk.Coins{{Amount: math.NewInt(10000), Denom: testenv.DefaultTestDenom}}
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 10)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		vestingCoins,
		startTime,
		endTime,
		accBalance,
	)
	testHelper.StakingUtils.MessageDelegate(2, 0, validatorsAddresses[0], accAddr2, math.NewInt(3000))
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(accAddr2, math.NewInt(100))

	testHelper.SetContextBlockTime(startTime.Add(time.Hour * 5))
	testHelper.StakingUtils.MessageUndelegate(3, 0, validatorsAddresses[0], accAddr2, math.NewInt(1000))
	testHelper.IncrementContextBlockHeight()
	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})
	testHelper.C4eClaimUtils.AddClaimRecordsFromWhitelistedVestingAccount(accAddr2, math.NewInt(6000), math.NewInt(6100))
}
