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
)

func TestUndelegate(t *testing.T) {
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"
	const delagableAddr = "cosmos1dfugyfm087qa3jrdglkeaew0wkn59jk8mgw6x6"
	const validatorAddr = "cosmosvaloper14k4pzckkre6uxxyd2lnhnpp8sngys9m6hl6ml7"
	addHelperModuleAccountPerms()
	accAddr, _ := sdk.AccAddressFromBech32(addr)
	delegableAccAddr, _ := sdk.AccAddressFromBech32(delagableAddr)
	valAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		require.Fail(t, err.Error())
	}
	const vt1 = "test1"
	const initBlock = 0
	const vested = 1000000
	accountVestings, vesting1 := createAccountVestings(addr, vt1, vested, 0)
	accountVestings.DelegableAddress = delagableAddr
	vesting1.DelegationAllowed = true

	app, ctx := setupApp(initBlock)

	PKs := testapp.CreateTestPubKeys(1)

	bank := app.BankKeeper
	staking := app.StakingKeeper
	dist := app.DistrKeeper
	k := app.CfevestingKeeper

	stakeParams := staking.GetParams(ctx)
	stakeParams.BondDenom = "uc4e"
	staking.SetParams(ctx, stakeParams)
	// adding coins to validotor
	addCoinsToAccount(vested, helperModuleAccount, ctx, bank, valAddr.Bytes())

	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(0, 1), sdk.NewDecWithPrec(0, 1), sdk.NewDec(0))
	delCoin := sdk.NewCoin(stakeParams.BondDenom, sdk.NewIntFromUint64(vested/2))
	createValidator(t, ctx, staking, valAddr, PKs[0], delCoin, commission)

	// adds coind to delegable account - means that coins in vesting for accAddr
	denom := addCoinsToAccount(vested, helperModuleAccount, ctx, bank, delegableAccAddr)
	// adds some coins to distibutor account - to allow test to process
	addCoinsToModuleByName(100000000, distrtypes.ModuleName, helperModuleAccount, ctx, bank)

	if len(staking.GetAllValidators(ctx)) == 0 {
		require.Fail(t, "no validators")
	}

	k.SetAccountVestings(ctx, accountVestings)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	coin := sdk.NewCoin(denom, sdk.NewIntFromUint64(vested/2))

	msg := types.MsgDelegate{DelegatorAddress: addr, ValidatorAddress: validatorAddr, Amount: coin}
	_, error := msgServer.Delegate(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	delegations := staking.GetAllDelegatorDelegations(ctx, delegableAccAddr)
	require.EqualValues(t, 1, len(delegations))
	delegation := delegations[0]
	require.EqualValues(t, sdk.NewDec(vested/2), delegation.Shares)

	validatorRewards := uint64(10000)
	valCons := sdk.NewDecCoin(denom, sdk.NewIntFromUint64(validatorRewards))
	val := app.StakingKeeper.Validator(ctx, valAddr)

	stakingmodule.EndBlocker(ctx, app.StakingKeeper)
	ctx = ctx.WithBlockHeight(initBlock + 1)
	// allocate reward to validator
	dist.AllocateTokensToValidator(ctx, val, sdk.NewDecCoins(valCons))
	msgServerCtx = sdk.WrapSDKContext(ctx)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, sdk.ZeroInt())
	verifyAccountBalance(t, bank, ctx, delegableAccAddr, denom, sdk.NewInt(vested/2))

	coin = sdk.NewCoin(denom, sdk.NewIntFromUint64(vested/2))
	msgUn := types.MsgUndelegate{DelegatorAddress: addr, ValidatorAddress: validatorAddr, Amount: coin}
	_, error = msgServer.Undelegate(msgServerCtx, &msgUn)
	require.EqualValues(t, nil, error)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, sdk.NewIntFromUint64(validatorRewards/2))
	verifyAccountBalance(t, bank, ctx, delegableAccAddr, denom, sdk.NewInt(vested/2))

}
