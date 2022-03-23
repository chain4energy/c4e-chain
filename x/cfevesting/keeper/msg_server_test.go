package keeper_test

import (
	"github.com/chain4energy/c4e-chain/app"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"testing"

	"github.com/stretchr/testify/require"
)

const helperModuleAccount = "helperTestAcc"
const denom = "uc4e"

func addHelperModuleAccountPerms() {
	perms := []string{authtypes.Minter}
	app.AddMaccPerms(helperModuleAccount, perms)
}

func addCoinsToAccount(vested uint64, ctx sdk.Context, app *app.App, toAddr sdk.AccAddress) string {
	denom := "uc4e"
	mintedCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(vested))
	mintedCoins := sdk.NewCoins(mintedCoin)
	app.BankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, helperModuleAccount, toAddr, mintedCoins)
	return denom
}

func createAccountVestings(addr string, vested uint64, withdrawn uint64) (types.AccountVestings, *types.Vesting) {
	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
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

func redelegate(t *testing.T, ctx sdk.Context, app *app.App, delegatorAddress sdk.AccAddress, delegableAddress sdk.AccAddress,
	validatorSrcAddress sdk.ValAddress, validatorDstAddress sdk.ValAddress, redelegateAmount uint64, delegatorAmountBefore uint64, delegableAmountBefore uint64, delegatorAmountAfter uint64, delegableAmountAfter uint64) {

	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewIntFromUint64(delegatorAmountBefore))
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewIntFromUint64(delegableAmountBefore))

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)
	coin := sdk.NewCoin(denom, sdk.NewIntFromUint64(redelegateAmount))
	msgRe := types.MsgBeginRedelegate{
		DelegatorAddress:    delegatorAddress.String(),
		ValidatorSrcAddress: validatorSrcAddress.String(),
		ValidatorDstAddress: validatorDstAddress.String(),
		Amount:              coin,
	}
	_, err := msgServer.BeginRedelegate(msgServerCtx, &msgRe)
	require.EqualValues(t, nil, err)

	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewIntFromUint64(delegatorAmountAfter))

	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewIntFromUint64(delegableAmountAfter))
}

func setupValidators(t *testing.T, ctx sdk.Context, app *app.App, validators []sdk.ValAddress, delegatePerValidator uint64) {
	PKs := commontestutils.CreateTestPubKeys(len(validators))
	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(0, 1), sdk.NewDecWithPrec(0, 1), sdk.NewDec(0))
	delCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(delegatePerValidator))
	for i, valAddr := range validators {
		addCoinsToAccount(delegatePerValidator, ctx, app, valAddr.Bytes())
		createValidator(t, ctx, app.StakingKeeper, valAddr, PKs[i], delCoin, commission)
	}
	require.EqualValues(t, len(validators), len(app.StakingKeeper.GetAllValidators(ctx)))
}

func setupStakingBondDenom(ctx sdk.Context, app *app.App) {
	stakeParams := app.StakingKeeper.GetParams(ctx)
	stakeParams.BondDenom = denom
	app.StakingKeeper.SetParams(ctx, stakeParams)
}

func delegate(t *testing.T, ctx sdk.Context, app *app.App, delegatorAddress sdk.AccAddress, delegableAddress sdk.AccAddress,
	validatorAddress sdk.ValAddress, delegateAmount uint64, delegatorAccountAmountBefore int64, delegableAccountAmountBefore int64, delegatorAccountAmountAfter int64, delegableAccountAmountAfter int64) {
	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewInt(delegatorAccountAmountBefore))
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAccountAmountBefore))

	coin := sdk.NewCoin(denom, sdk.NewIntFromUint64(delegateAmount))
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgDelegate{DelegatorAddress: delegatorAddress.String(), ValidatorAddress: validatorAddress.String(), Amount: coin}
	_, err := msgServer.Delegate(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)
	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewInt(delegatorAccountAmountAfter))
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAccountAmountAfter))
}

func verifyDelegations(t *testing.T, ctx sdk.Context, app *app.App, delegableAddress sdk.AccAddress,
	validators []sdk.ValAddress, delegated []int64) {
	delegations := app.StakingKeeper.GetAllDelegatorDelegations(ctx, delegableAddress)
	require.EqualValues(t, len(validators), len(delegations))
	for i, valAddr := range validators {
		found := false
		for _, delegation := range delegations {
			if delegation.ValidatorAddress == valAddr.String() {
				require.EqualValues(t, sdk.NewDec(delegated[i]), delegation.Shares)
				found = true
			}
		}
		require.True(t, found, "delegation not found. Validator Address: "+valAddr.String())
	}

}

func setupAccountsVestings(ctx sdk.Context, app *app.App, address string, delegableAddress string, vesingAmount uint64, delegationAllowed bool) {
	accountVestings, vesting1 := createAccountVestings(address, vesingAmount, 0)
	accountVestings.DelegableAddress = delegableAddress
	vesting1.DelegationAllowed = delegationAllowed
	app.CfevestingKeeper.SetAccountVestings(ctx, accountVestings)
}

func allocateRewardsToValidator(ctx sdk.Context, app *app.App, validatorRewards uint64, valAddr sdk.ValAddress) {
	valCons := sdk.NewDecCoin(denom, sdk.NewIntFromUint64(validatorRewards))
	val := app.StakingKeeper.Validator(ctx, valAddr)
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, sdk.NewDecCoins(valCons))
}

func verifyQueryRewards(t *testing.T, ctx sdk.Context, app *app.App, delegableAddr sdk.AccAddress, valAddr sdk.ValAddress, hasRewards bool, rewards uint64) {
	msgServerCtx := sdk.WrapSDKContext(ctx)
	query := distrtypes.QueryDelegationRewardsRequest{DelegatorAddress: delegableAddr.String(), ValidatorAddress: valAddr.String()}
	resp, _ := app.DistrKeeper.DelegationRewards(msgServerCtx, &query)
	if hasRewards {
		require.EqualValues(t, 1, len(resp.Rewards))
		require.EqualValues(t, sdk.NewDecFromInt(sdk.NewIntFromUint64(rewards)), resp.Rewards[0].Amount)
	} else {
		require.EqualValues(t, 0, len(resp.Rewards))
	}
	
}