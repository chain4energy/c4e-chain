package keeper_test

import (
	"testing"
	"github.com/chain4energy/c4e-chain/x/cfevesting/internal/testutils"
	"github.com/chain4energy/c4e-chain/app"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/stretchr/testify/require"
)

const helperModuleAccount = "helperTestAcc"
const denom = "uc4e"

func addHelperModuleAccountPerms() {
	perms := []string{authtypes.Minter}
	app.AddMaccPerms(helperModuleAccount, perms)
}

// TODO remove
// func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
// 	k, ctx := keepertest.CfevestingKeeper(t)
// 	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
// }

func addCoinsToAccount(vested uint64, ctx sdk.Context, app *app.App, toAddr sdk.AccAddress) string {
	denom := "uc4e"
	mintedCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(vested))
	mintedCoins := sdk.NewCoins(mintedCoin)
	app.BankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, helperModuleAccount, toAddr, mintedCoins)
	return denom
}

func createAccountVestings(addr string, vested uint64, withdrawn uint64) (types.AccountVestings, *types.Vesting) {
	accountVestings :=testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.Vestings[0].Vested = sdk.NewIntFromUint64(vested)
	accountVestings.Vestings[0].DelegationAllowed = false
	accountVestings.Vestings[0].Withdrawn = sdk.NewIntFromUint64(withdrawn)
	accountVestings.Vestings[0].LastModificationVested = sdk.NewIntFromUint64(vested)
	accountVestings.Vestings[0].LastModificationWithdrawn = sdk.NewIntFromUint64(withdrawn)
	return accountVestings, accountVestings.Vestings[0]
}

func addCoinsToModuleByName(vested uint64, modulaName string, ctx sdk.Context, app *app.App) string {
	denom := "uc4e"
	mintedCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(vested))
	mintedCoins := sdk.NewCoins(mintedCoin)
	app.BankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	app.BankKeeper.SendCoinsFromModuleToModule(ctx, helperModuleAccount, modulaName, mintedCoins)
	return denom
}

func verifyAccountBalance(t *testing.T, app *app.App, ctx sdk.Context, accAddr sdk.AccAddress, expectedAmount sdk.Int) {
	balance := app.BankKeeper.GetBalance(ctx, accAddr, denom)
	require.EqualValues(t, expectedAmount, balance.Amount)
}

func verifyModuleAccountByName(accName string, auth authkeeper.AccountKeeper, ctx sdk.Context, bank bankkeeper.Keeper, denom string, t *testing.T, expected sdk.Int) {
	moduleAccAddr := auth.GetModuleAccount(ctx, accName).GetAddress()
	moduleBalance := bank.GetBalance(ctx, moduleAccAddr, denom)
	require.EqualValues(t, expected, moduleBalance.Amount)
}

func verifyModuleAccount(auth authkeeper.AccountKeeper, ctx sdk.Context, bank bankkeeper.Keeper, denom string, t *testing.T, expected sdk.Int) {
	verifyModuleAccountByName(types.ModuleName, auth, ctx, bank, denom, t, expected)
}

func createValidator(t *testing.T, ctx sdk.Context, sk stakingkeeper.Keeper, addr sdk.ValAddress, pk cryptotypes.PubKey, coin sdk.Coin, commisions stakingtypes.CommissionRates) {
	msg, err := stakingtypes.NewMsgCreateValidator(addr, pk, coin, stakingtypes.Description{}, commisions, sdk.OneInt())
	msgSrvr := stakingkeeper.NewMsgServerImpl(sk)
	require.NoError(t, err)
	res, err := msgSrvr.CreateValidator(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, res)

}
