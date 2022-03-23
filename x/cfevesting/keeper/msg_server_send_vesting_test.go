package keeper_test

import (
	"strconv"
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	testapp "github.com/chain4energy/c4e-chain/app"
	"github.com/chain4energy/c4e-chain/x/cfevesting/internal/testutils"

)

func TestSendVestingDelegationNotAllowedNoVestingRestart(t *testing.T) {
	sendVestingDelegation(t, false, false, noWithdrawableNoDelegatedEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestart(t *testing.T) {
	sendVestingDelegation(t, true, false, noWithdrawableNoDelegatedEnoughToSend)
}

func TestSendVestingDelegationNotAllowedVestingRestart(t *testing.T) {
	sendVestingDelegation(t, false, true, noWithdrawableNoDelegatedEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestart(t *testing.T) {
	sendVestingDelegation(t, true, true, noWithdrawableNoDelegatedEnoughToSend)
}

func TestSendVestingDelegationNotAllowedNoVestingRestartWithWithdrawable(t *testing.T) {
	sendVestingDelegation(t, false, false, withdrawableEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestartWithWithdrawable(t *testing.T) {
	sendVestingDelegation(t, true, false, withdrawableEnoughToSend)
}

func TestSendVestingDelegationNotAllowedVestingRestartWithWithdrawable(t *testing.T) {
	sendVestingDelegation(t, false, true, withdrawableEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestartWithWithdrawable(t *testing.T) {
	sendVestingDelegation(t, true, true, withdrawableEnoughToSend)
}

func TestSendVestingDelegationNotAllowedNoVestingRestartWithWithdrawableAndNotEnoughToSend(t *testing.T) {
	sendVestingDelegation(t, false, false, withdrawableNotEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestartWithWithdrawableAndNotEnoughToSend(t *testing.T) {
	sendVestingDelegation(t, true, false, withdrawableNotEnoughToSend)
}

func TestSendVestingDelegationNotAllowedVestingRestartWithWithdrawableAndNotEnoughToSend(t *testing.T) {
	sendVestingDelegation(t, false, true, withdrawableNotEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestartWithWithdrawableAndNotEnoughToSend(t *testing.T) {
	sendVestingDelegation(t, true, true, withdrawableNotEnoughToSend)
}

func TestSendVestingDelegationNotAllowedNoVestingRestartWithAllWithdrawable(t *testing.T) {
	sendVestingDelegation(t, false, false, allWithdrawableEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestartWithAllWithdrawable(t *testing.T) {
	sendVestingDelegation(t, true, false, allWithdrawableEnoughToSend)
}

func TestSendVestingDelegationNotAllowedVestingRestartWithAllWithdrawable(t *testing.T) {
	sendVestingDelegation(t, false, true, allWithdrawableEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestartWithAllWithdrawable(t *testing.T) {
	sendVestingDelegation(t, true, true, allWithdrawableEnoughToSend)
}


func TestSendVestingDelegationAllowedNoVestingRestartWithDelegatedEnoughToSend(t *testing.T) {
	sendVestingDelegation(t, true, false, delegatedEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestartWithDelegatedEnoughToSend(t *testing.T) {
	sendVestingDelegation(t, true, true, delegatedEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestartWithDelegatedNotEnoughToSend(t *testing.T) {
	sendVestingDelegation(t, true, false, delegatedNotEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestartWithDelegatedNotEnoughToSend(t *testing.T) {
	sendVestingDelegation(t, true, true, delegatedNotEnoughToSend)
}

const noWithdrawableNoDelegatedEnoughToSend = 0
const withdrawableEnoughToSend = 1
const withdrawableNotEnoughToSend = 2
const allWithdrawableEnoughToSend = 3
const delegatedEnoughToSend = 4
const delegatedNotEnoughToSend = 5

func sendVestingDelegation(t *testing.T, delegationAllowed bool, restartVesting bool, testType int) {
	const addrVestSrc = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"
	const delagableAddrVestSrc = "cosmos1dfugyfm087qa3jrdglkeaew0wkn59jk8mgw6x6"

	const addrVestDst = "cosmos1k056avwx8gx3jnte7k8plpxgk7ymsyegxpu64c"
	const delagableAddrVestDst = "cosmos10wnn4eqpppcxd3aq5lsfd7sj2wqhzjcqgh6s2m"


	const validatorAddr = "cosmosvaloper14k4pzckkre6uxxyd2lnhnpp8sngys9m6hl6ml7"

	accAddr, _ := sdk.AccAddressFromBech32(addrVestSrc)
	delegableAccAddr, _ := sdk.AccAddressFromBech32(delagableAddrVestSrc)

	valAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		require.Fail(t, err.Error())
	}

	const vt1 = "test1"
	const initBlock = 1000
	const vested = 1000
	const accInitBalance = 10000
	vestingTypes := types.VestingTypes{}
	vestingType1 := types.VestingType{
		Name:                 vt1,
		LockupPeriod:         1000,
		VestingPeriod:        5000,
		TokenReleasingPeriod: 10,
		DelegationsAllowed:   delegationAllowed,
	}
	vestingType2 := types.VestingType{
		Name:                 "test2",
		LockupPeriod:         1111,
		VestingPeriod:        112233,
		TokenReleasingPeriod: 445566,
		DelegationsAllowed:   false,
	}

	vestingTypesArray := []*types.VestingType{&vestingType1, &vestingType2}
	vestingTypes.VestingTypes = vestingTypesArray

	addHelperModuleAccountPerms()

	app := testapp.Setup(false)
	header := tmproto.Header{}
	header.Height = initBlock
	ctx := app.BaseApp.NewContext(false, header)

	bank := app.BankKeeper
	auth := app.AccountKeeper
	staking := app.StakingKeeper
	denom := "uc4e"


	if testType == delegatedEnoughToSend || testType == delegatedNotEnoughToSend {
		PKs := testutils.CreateTestPubKeys(1)
		stakeParams := staking.GetParams(ctx)
		stakeParams.BondDenom = "uc4e"
		staking.SetParams(ctx, stakeParams)
		// adding coins to validotor
		addCoinsToAccount(vested, ctx, app, valAddr.Bytes())

		commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(0, 1), sdk.NewDecWithPrec(0, 1), sdk.NewDec(0))
		delCoin := sdk.NewCoin(stakeParams.BondDenom, sdk.NewIntFromUint64(vested/2))
		createValidator(t, ctx, staking, valAddr, PKs[0], delCoin, commission)
	}

	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	k := app.CfevestingKeeper

	k.SetVestingTypes(ctx, vestingTypes)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgVest{Creator: addrVestSrc, Amount: sdk.NewInt(vested), VestingType: vt1}
	_, error := msgServer.Vest(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVesting, accFound := k.GetAccountVestings(ctx, addrVestDst)
	require.EqualValues(t, false, accFound)

	accVesting, accFound = k.GetAccountVestings(ctx, addrVestSrc)
	require.EqualValues(t, true, accFound)

	var expectedDelegableAcc string
	if delegationAllowed {
		expectedDelegableAcc = delagableAddrVestSrc
	} else {
		expectedDelegableAcc = ""
	}
	require.EqualValues(t, 1, len(accVesting.Vestings))
	require.EqualValues(t, expectedDelegableAcc, accVesting.DelegableAddress)

	require.EqualValues(t, addrVestSrc, accVesting.Address)
	vesting := accVesting.Vestings[0]
	require.EqualValues(t, 1, vesting.Id)
	require.EqualValues(t, vt1, vesting.VestingType)
	require.EqualValues(t, initBlock, vesting.VestingStartBlock)
	require.EqualValues(t, initBlock+vestingType1.LockupPeriod, vesting.LockEndBlock)
	require.EqualValues(t, initBlock+vestingType1.LockupPeriod+vestingType1.VestingPeriod, vesting.VestingEndBlock)
	require.EqualValues(t, sdk.NewInt(vested), vesting.Vested)
	require.EqualValues(t, vestingType1.TokenReleasingPeriod, vesting.FreeCoinsBlockPeriod)
	require.EqualValues(t, delegationAllowed, vesting.DelegationAllowed)
	require.EqualValues(t, sdk.NewInt(vested), vesting.LastModificationVested)
	require.EqualValues(t, initBlock, vesting.LastModificationBlock)
	require.EqualValues(t, sdk.ZeroInt(), vesting.LastModificationWithdrawn)
	require.EqualValues(t, sdk.ZeroInt(), vesting.Sent)

	balance := bank.GetBalance(ctx, accAddr, denom)
	require.EqualValues(t, sdk.NewIntFromUint64(accInitBalance-vested), balance.Amount)
	moduleAccAddr := auth.GetModuleAccount(ctx, types.ModuleName).GetAddress()
	moduleBalance := bank.GetBalance(ctx, moduleAccAddr, denom)
	if delegationAllowed {
		require.EqualValues(t, sdk.ZeroInt(), moduleBalance.Amount)
		delegableBalance := bank.GetBalance(ctx, delegableAccAddr, denom)
		require.EqualValues(t, sdk.NewIntFromUint64(vested), delegableBalance.Amount)
	} else {
		require.EqualValues(t, sdk.NewIntFromUint64(vested), moduleBalance.Amount)
	}

	// -----------
	
	addHeight := int64(500)
	if (testType == withdrawableEnoughToSend) {
		addHeight = int64(2000)
	} else if (testType == withdrawableNotEnoughToSend) {
		addHeight = int64(5600)
	} else if (testType == allWithdrawableEnoughToSend) {
		addHeight = int64(7000)
	}
	ctx = ctx.WithBlockHeight(initBlock + addHeight)
	msgServer, msgServerCtx = keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	vestingData, error := k.Vesting(msgServerCtx, &types.QueryVestingRequest{Address: addrVestSrc})
	require.EqualValues(t, nil, error)

	witdrawable := int64(0)
	if (testType == withdrawableEnoughToSend) {
		witdrawable = int64(200)
	} else if (testType == withdrawableNotEnoughToSend) {
		witdrawable = int64(920)
	} else if (testType == allWithdrawableEnoughToSend) {
		witdrawable = int64(1000)
	}

	sentAmount := int64(vested / 10)
	
	require.EqualValues(t, strconv.FormatInt(int64(witdrawable), 10), vestingData.Vestings[0].Withdrawable)

	delegateAmount := uint64(0)
	if testType == delegatedEnoughToSend || testType == delegatedNotEnoughToSend {
		delegateAmount = uint64(300)

		if testType == delegatedNotEnoughToSend {
			delegateAmount = uint64(950)
		}
		coin := sdk.NewCoin(denom, sdk.NewIntFromUint64(delegateAmount))
		msgDelegate := types.MsgDelegate{DelegatorAddress: addrVestSrc, ValidatorAddress: validatorAddr, Amount: coin}
		_, error := msgServer.Delegate(msgServerCtx, &msgDelegate)
		require.EqualValues(t, nil, error)
	}

	msgSendVesting := types.MsgSendVesting{FromAddress: addrVestSrc, ToAddress: addrVestDst, VestingId: 1, Amount: sdk.NewInt(sentAmount), RestartVesting: restartVesting}
	_, error = msgServer.SendVesting(msgServerCtx, &msgSendVesting)

	if testType == withdrawableNotEnoughToSend || testType == allWithdrawableEnoughToSend {
		require.NotEqualValues(t, nil, error)
		// require.EqualError()
		require.EqualError(t, error,
			"vesting available: " + strconv.FormatInt(vested-witdrawable, 10) + " is smaller than " +  strconv.FormatInt(sentAmount, 10) + ": insufficient funds")
	} else if testType == delegatedNotEnoughToSend {
		require.NotEqualValues(t, nil, error)
		// require.EqualError()
		require.EqualError(t, error,
			"vesting available: " + strconv.FormatUint(uint64(vested)-delegateAmount, 10) + 
			" is smaller than " +  strconv.FormatInt(sentAmount, 10) + " - probably delageted to validator.: insufficient funds")
	} else {
		require.EqualValues(t, nil, error)
	}
	wasSent := true
	if testType == withdrawableNotEnoughToSend || testType == allWithdrawableEnoughToSend || testType == delegatedNotEnoughToSend {
		wasSent = false
		sentAmount = int64(0)
	}

	accVesting, accFound = k.GetAccountVestings(ctx, addrVestSrc)
	require.EqualValues(t, true, accFound)
	require.EqualValues(t, 1, len(accVesting.Vestings))
	vesting = accVesting.Vestings[0]

	require.EqualValues(t, 1, vesting.Id)
	require.EqualValues(t, vt1, vesting.VestingType)
	require.EqualValues(t, initBlock, vesting.VestingStartBlock)
	require.EqualValues(t, initBlock+vestingType1.LockupPeriod, vesting.LockEndBlock)

	require.EqualValues(t, initBlock+vestingType1.LockupPeriod+vestingType1.VestingPeriod, vesting.VestingEndBlock)

	require.EqualValues(t, sdk.NewInt(vested), vesting.Vested)
	require.EqualValues(t, vestingType1.TokenReleasingPeriod, vesting.FreeCoinsBlockPeriod)
	require.EqualValues(t, delegationAllowed, vesting.DelegationAllowed)

	require.EqualValues(t, sdk.NewInt(witdrawable), vesting.Withdrawn)

	if wasSent {
		require.EqualValues(t, sdk.NewInt(vested-sentAmount-witdrawable), vesting.LastModificationVested)
		require.EqualValues(t, initBlock+addHeight, vesting.LastModificationBlock)
		require.EqualValues(t, sdk.ZeroInt(), vesting.LastModificationWithdrawn)
	} else {
		require.EqualValues(t, sdk.NewInt(vested), vesting.LastModificationVested)
		require.EqualValues(t, initBlock, vesting.LastModificationBlock)
		require.EqualValues(t, sdk.NewInt(witdrawable), vesting.LastModificationWithdrawn)
	}
	require.EqualValues(t, sdk.NewInt(sentAmount), vesting.Sent)

	balance = bank.GetBalance(ctx, accAddr, denom)
	require.EqualValues(t, sdk.NewIntFromUint64(uint64(accInitBalance-vested+witdrawable)), balance.Amount)
	moduleBalance = bank.GetBalance(ctx, moduleAccAddr, denom)
	if delegationAllowed {
		require.EqualValues(t, sdk.ZeroInt(), moduleBalance.Amount)
		delegableBalance := bank.GetBalance(ctx, delegableAccAddr, denom)
		require.EqualValues(t, sdk.NewIntFromUint64(uint64(vested-sentAmount-witdrawable)-delegateAmount), delegableBalance.Amount)
	} else {
		require.EqualValues(t, sdk.NewIntFromUint64(uint64(vested-witdrawable)), moduleBalance.Amount)
	}


	accVesting, accFound = k.GetAccountVestings(ctx, addrVestDst)
	require.EqualValues(t, wasSent, accFound)
	if !wasSent {
		return
	}
	require.EqualValues(t, 1, len(accVesting.Vestings))
	vesting = accVesting.Vestings[0]

	require.EqualValues(t, 1, vesting.Id)
	require.EqualValues(t, vt1, vesting.VestingType)
	require.EqualValues(t, ctx.BlockHeight(), vesting.VestingStartBlock)

	if restartVesting {
		require.EqualValues(t, ctx.BlockHeight()+vestingType1.LockupPeriod, vesting.LockEndBlock)
		require.EqualValues(t, ctx.BlockHeight()+vestingType1.LockupPeriod+vestingType1.VestingPeriod, vesting.VestingEndBlock)
	} else {
		require.EqualValues(t, initBlock+vestingType1.LockupPeriod, vesting.LockEndBlock)
		require.EqualValues(t, initBlock+vestingType1.LockupPeriod+vestingType1.VestingPeriod, vesting.VestingEndBlock)
	}

	require.EqualValues(t, sdk.NewInt(sentAmount), vesting.Vested)
	require.EqualValues(t, vestingType1.TokenReleasingPeriod, vesting.FreeCoinsBlockPeriod)
	require.EqualValues(t, delegationAllowed, vesting.DelegationAllowed)

	require.EqualValues(t, sdk.NewInt(sentAmount), vesting.LastModificationVested)
	require.EqualValues(t, ctx.BlockHeight(), vesting.LastModificationBlock)
	require.EqualValues(t, sdk.ZeroInt(), vesting.LastModificationWithdrawn)
	require.EqualValues(t, sdk.ZeroInt(), vesting.Sent)


	moduleBalance = bank.GetBalance(ctx, moduleAccAddr, denom)
	if delegationAllowed {
		require.EqualValues(t, delagableAddrVestDst, accVesting.DelegableAddress)
		require.EqualValues(t, sdk.ZeroInt(), moduleBalance.Amount)
		delegableAccAddrDst, _ := sdk.AccAddressFromBech32(delagableAddrVestDst)
		delegableBalance := bank.GetBalance(ctx, delegableAccAddrDst, denom)
		require.EqualValues(t, sdk.NewIntFromUint64(vested/10), delegableBalance.Amount)
	} else {
		require.EqualValues(t, sdk.NewIntFromUint64(uint64(vested-witdrawable)), moduleBalance.Amount)
	}
}
