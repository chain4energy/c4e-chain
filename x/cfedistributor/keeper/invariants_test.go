package keeper_test

import (
	"testing"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestNonNegativeCoinStateInvariantCorrect(t *testing.T) {
	testUtil, _, ctx := testkeeper.CfedistributorKeeperTestUtil(t)

	state := types.State{Account: &types.Account{Id: "test", Type: types.INTERNAL_ACCOUNT}, Burn: false, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(1324)}}}
	testUtil.SetState(ctx, state)
	state = types.State{Burn: true, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(0)}}}
	testUtil.SetState(ctx, state)

	testUtil.CheckNonNegativeCoinStateInvariant(ctx, false,
		"cfedistributor: nonnegative coin state invariant\n\tno negative coin states\n")
}

func TestNonNegativeCoinStateInvariantNegativeAccountSate(t *testing.T) {
	testUtil, _, ctx := testkeeper.CfedistributorKeeperTestUtil(t)

	state := types.State{Account: &types.Account{Id: "test", Type: types.INTERNAL_ACCOUNT}, Burn: false, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(-1)}}}
	testUtil.SetState(ctx, state)
	state = types.State{Burn: true, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(0)}}}
	testUtil.SetState(ctx, state)

	testUtil.CheckNonNegativeCoinStateInvariant(ctx, true,
		"cfedistributor: nonnegative coin state invariant\n\tnegative coin state -1.000000000000000000uc4e in state INTERNAL_ACCOUNT-test\n")
}

func TestNonNegativeCoinStateInvariantNegativeBurnSate(t *testing.T) {
	testUtil, _, ctx := testkeeper.CfedistributorKeeperTestUtil(t)

	state := types.State{Account: &types.Account{Id: "test", Type: types.INTERNAL_ACCOUNT}, Burn: false, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(324)}}}
	testUtil.SetState(ctx, state)
	state = types.State{Account: &types.Account{Id: "", Type: ""}, Burn: true, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(-1)}}}
	testUtil.SetState(ctx, state)

	testUtil.CheckNonNegativeCoinStateInvariant(ctx, true,
		"cfedistributor: nonnegative coin state invariant\n\tnegative coin state -1.000000000000000000uc4e in state BURN\n")
}

func TestStateSumBalanceCheckInvariantCorrect(t *testing.T) {
	testHelper := testapp.SetupTestApp(t)

	state := types.State{Account: &types.Account{Id: "test", Type: types.INTERNAL_ACCOUNT}, Burn: false, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(1324)}}}
	testHelper.C4eDistributorUtils.SetState(state)
	state = types.State{Account: &types.Account{Id: "test2", Type: types.INTERNAL_ACCOUNT}, Burn: false, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(200)}}}
	testHelper.C4eDistributorUtils.SetState(state)
	state = types.State{Burn: true, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(0)}}}
	testHelper.C4eDistributorUtils.SetState(state)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1524), types.DistributorMainAccount)

	testHelper.C4eDistributorUtils.CheckStateSumBalanceCheckInvariant(false,
		"cfedistributor: state sum balance check invariant\n\tsum of states coins: 1524uc4e\n\tdistributor account balance: 1524uc4e\n")
}

func TestStateSumBalanceCheckInvariantSumNotInt(t *testing.T) {
	testHelper := testapp.SetupTestApp(t)

	state := types.State{Account: &types.Account{Id: "test", Type: types.INTERNAL_ACCOUNT}, Burn: false, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.MustNewDecFromStr("12.132")}}}
	testHelper.C4eDistributorUtils.SetState(state)
	state = types.State{Account: &types.Account{Id: "test2", Type: types.INTERNAL_ACCOUNT}, Burn: false, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(200)}}}
	testHelper.C4eDistributorUtils.SetState(state)
	state = types.State{Burn: true, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(0)}}}
	testHelper.C4eDistributorUtils.SetState(state)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1524), types.DistributorMainAccount)

	testHelper.C4eDistributorUtils.CheckStateSumBalanceCheckInvariant(true,
		"cfedistributor: state sum balance check invariant\n\tthe sum of the states should be integer: sum: 212.132000000000000000uc4e\n")
}

func TestStateSumBalanceCheckInvariantSumDiffersFromModuleAccountBalance(t *testing.T) {
	testHelper := testapp.SetupTestApp(t)

	state := types.State{Account: &types.Account{Id: "test", Type: types.INTERNAL_ACCOUNT}, Burn: false, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(1324)}}}
	testHelper.C4eDistributorUtils.SetState(state)
	state = types.State{Account: &types.Account{Id: "test2", Type: types.INTERNAL_ACCOUNT}, Burn: false, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(200)}}}
	testHelper.C4eDistributorUtils.SetState(state)
	state = types.State{Burn: true, Remains: sdk.DecCoins{sdk.DecCoin{Denom: commontestutils.DefaultTestDenom, Amount: sdk.NewDec(0)}}}
	testHelper.C4eDistributorUtils.SetState(state)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1523), types.DistributorMainAccount)

	testHelper.C4eDistributorUtils.CheckStateSumBalanceCheckInvariant(true,
		"cfedistributor: state sum balance check invariant\n\tsum of states coins: 1524uc4e\n\tdistributor account balance: 1523uc4e\n")
}
