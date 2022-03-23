package keeper_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	testapp "github.com/chain4energy/c4e-chain/app"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingmodule "github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"github.com/chain4energy/c4e-chain/x/cfevesting/internal/testutils"

)

func TestRedelegate(t *testing.T) {
	addHelperModuleAccountPerms()
	// const initBlock = 0
	const vested = 1000000
	app, ctx := setupApp(0)
	// k := app.CfevestingKeeper
	setupStakingBondDenom(ctx, app)

	acountsAddresses, validatorsAddresses := testutils.CreateAccounts(2, 2)

	setupValidators(t, ctx, app, validatorsAddresses, vested/2)

	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	// adds coind to delegable account - means that coins in vesting for accAddr
	addCoinsToAccount(vested, ctx, app, delegableAccAddr)
	// adds some coins to distibutor account - to allow test to process
	addCoinsToModuleByName(100000000, distrtypes.ModuleName, ctx, app)

	valAddr := validatorsAddresses[0]
	valAddr2 := validatorsAddresses[1]

	setupAccountsVestings(ctx, app, accAddr.String(), delegableAccAddr.String(), vested)

	delegate(t, ctx, app, accAddr, delegableAccAddr, valAddr, vested/2, vested, vested/2)
	verifyDelegations(t, ctx, app, delegableAccAddr,  []sdk.ValAddress{valAddr}, []int64{vested/2})

	stakingmodule.EndBlocker(ctx, app.StakingKeeper)
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)

	validatorRewards := uint64(10000)
	allocateRewardsToValidator(ctx, app, validatorRewards, valAddr)

	redelegate(t, ctx, app, accAddr, delegableAccAddr, valAddr, valAddr2, vested/2, 0, vested/2, validatorRewards/2, vested/2)
	verifyDelegations(t, ctx, app, delegableAccAddr,  []sdk.ValAddress{valAddr2}, []int64{vested/2})

}

func redelegate(t *testing.T, ctx sdk.Context, app *testapp.App, delegatorAddress sdk.AccAddress, delegableAddress sdk.AccAddress, 
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

func setupValidators(t *testing.T, ctx sdk.Context, app *testapp.App,  validators []sdk.ValAddress, delegatePerValidator uint64) {
	PKs := testutils.CreateTestPubKeys(len(validators))
	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(0, 1), sdk.NewDecWithPrec(0, 1), sdk.NewDec(0))
	delCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(delegatePerValidator))
	for i, valAddr := range validators {
		addCoinsToAccount(delegatePerValidator, ctx, app, valAddr.Bytes())
		createValidator(t, ctx, app.StakingKeeper, valAddr, PKs[i], delCoin, commission)
	}
	require.EqualValues(t, len(validators), len(app.StakingKeeper.GetAllValidators(ctx)))
}

func setupStakingBondDenom(ctx sdk.Context, app *testapp.App) {
	stakeParams := app.StakingKeeper.GetParams(ctx)
	stakeParams.BondDenom = denom
	app.StakingKeeper.SetParams(ctx, stakeParams)
}

func delegate(t *testing.T, ctx sdk.Context, app *testapp.App, delegatorAddress sdk.AccAddress, delegableAddress sdk.AccAddress,
				validatorAddress sdk.ValAddress, delegateAmount uint64, delegableAccountAmountBefore int64, delegableAccountAmountAfter int64) {
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAccountAmountBefore))

	coin := sdk.NewCoin(denom, sdk.NewIntFromUint64(delegateAmount))
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgDelegate{DelegatorAddress: delegatorAddress.String(), ValidatorAddress: validatorAddress.String(), Amount: coin}
	_, err := msgServer.Delegate(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAccountAmountAfter))
}

func verifyDelegations(t *testing.T, ctx sdk.Context, app *testapp.App, delegableAddress sdk.AccAddress,
		validators []sdk.ValAddress, delegated []int64, ) {
	delegations := app.StakingKeeper.GetAllDelegatorDelegations(ctx, delegableAddress)
	require.EqualValues(t, len(validators), len(delegations))
	for i, valAddr  := range validators {
		found := false
		for _, delegation := range delegations{
			if delegation.ValidatorAddress == valAddr.String() {
				require.EqualValues(t, sdk.NewDec(delegated[i]), delegation.Shares)
				found = true
			}
		}
		require.True(t, found, "delegation not found. Validator Address: "+valAddr.String())
	}

}

func setupAccountsVestings(ctx sdk.Context, app *testapp.App, address string, delegableAddress string, vesingAmount uint64) {
	accountVestings, vesting1 := createAccountVestings(address, vesingAmount, 0)
	accountVestings.DelegableAddress = delegableAddress
	vesting1.DelegationAllowed = true
	app.CfevestingKeeper.SetAccountVestings(ctx, accountVestings)
}

func allocateRewardsToValidator(ctx sdk.Context, app *testapp.App, validatorRewards uint64, valAddr sdk.ValAddress) {
	valCons := sdk.NewDecCoin(denom, sdk.NewIntFromUint64(validatorRewards))
	val := app.StakingKeeper.Validator(ctx, valAddr)
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, sdk.NewDecCoins(valCons))
}
