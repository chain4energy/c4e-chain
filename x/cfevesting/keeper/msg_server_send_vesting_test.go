package keeper_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"

)

func TestSendVestingDelegationNotAllowedNoVestingRestart(t *testing.T) {
	sendVestingTest(t, false, false, noWithdrawableNoDelegatedEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestart(t *testing.T) {
	sendVestingTest(t, true, false, noWithdrawableNoDelegatedEnoughToSend)
}

func TestSendVestingDelegationNotAllowedVestingRestart(t *testing.T) {
	sendVestingTest(t, false, true, noWithdrawableNoDelegatedEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestart(t *testing.T) {
	sendVestingTest(t, true, true, noWithdrawableNoDelegatedEnoughToSend)
}

func TestSendVestingDelegationNotAllowedNoVestingRestartWithWithdrawable(t *testing.T) {
	sendVestingTest(t, false, false, withdrawableEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestartWithWithdrawable(t *testing.T) {
	sendVestingTest(t, true, false, withdrawableEnoughToSend)
}

func TestSendVestingDelegationNotAllowedVestingRestartWithWithdrawable(t *testing.T) {
	sendVestingTest(t, false, true, withdrawableEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestartWithWithdrawable(t *testing.T) {
	sendVestingTest(t, true, true, withdrawableEnoughToSend)
}

func TestSendVestingDelegationNotAllowedNoVestingRestartWithWithdrawableAndNotEnoughToSend(t *testing.T) {
	sendVestingTest(t, false, false, withdrawableNotEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestartWithWithdrawableAndNotEnoughToSend(t *testing.T) {
	sendVestingTest(t, true, false, withdrawableNotEnoughToSend)
}

func TestSendVestingDelegationNotAllowedVestingRestartWithWithdrawableAndNotEnoughToSend(t *testing.T) {
	sendVestingTest(t, false, true, withdrawableNotEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestartWithWithdrawableAndNotEnoughToSend(t *testing.T) {
	sendVestingTest(t, true, true, withdrawableNotEnoughToSend)
}

func TestSendVestingDelegationNotAllowedNoVestingRestartWithAllWithdrawable(t *testing.T) {
	sendVestingTest(t, false, false, allWithdrawableEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestartWithAllWithdrawable(t *testing.T) {
	sendVestingTest(t, true, false, allWithdrawableEnoughToSend)
}

func TestSendVestingDelegationNotAllowedVestingRestartWithAllWithdrawable(t *testing.T) {
	sendVestingTest(t, false, true, allWithdrawableEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestartWithAllWithdrawable(t *testing.T) {
	sendVestingTest(t, true, true, allWithdrawableEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestartWithDelegatedEnoughToSend(t *testing.T) {
	sendVestingTest(t, true, false, delegatedEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestartWithDelegatedEnoughToSend(t *testing.T) {
	sendVestingTest(t, true, true, delegatedEnoughToSend)
}

func TestSendVestingDelegationAllowedNoVestingRestartWithDelegatedNotEnoughToSend(t *testing.T) {
	sendVestingTest(t, true, false, delegatedNotEnoughToSend)
}

func TestSendVestingDelegationAllowedVestingRestartWithDelegatedNotEnoughToSend(t *testing.T) {
	sendVestingTest(t, true, true, delegatedNotEnoughToSend)
}

const noWithdrawableNoDelegatedEnoughToSend = 0
const withdrawableEnoughToSend = 1
const withdrawableNotEnoughToSend = 2
const allWithdrawableEnoughToSend = 3
const delegatedEnoughToSend = 4
const delegatedNotEnoughToSend = 5

func sendVestingTest(t *testing.T, delegationAllowed bool, restartVesting bool, testType int) {
	addHelperModuleAccountPerms()
	const vested = 1000
	const accInitBalance = 10000
	app, ctx := setupApp(1000)
	setupStakingBondDenom(ctx, app)

	acountsAddresses, validatorsAddresses := commontestutils.CreateAccounts(2, 1)
	accSrcAddr := acountsAddresses[0]
	accDestAddr := acountsAddresses[1]

	addCoinsToAccount(accInitBalance, ctx, app, accSrcAddr)

	if testType == delegatedEnoughToSend || testType == delegatedNotEnoughToSend {
		setupValidators(t, ctx, app, validatorsAddresses, vested/2)
	}

	valAddr := validatorsAddresses[0]

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, delegationAllowed, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	if delegationAllowed {
		makeVesting(t, ctx, app, accSrcAddr, false, true, false, true, *usedVestingType, vested, accInitBalance, 0, 0, accInitBalance-vested, vested, 0)
	} else {
		makeVesting(t, ctx, app, accSrcAddr, false, true, false, false, *usedVestingType, vested, accInitBalance, 0, 0, accInitBalance-vested, 0, vested)
	}

	verifyAccountVestings(t, ctx, app, accSrcAddr, []types.VestingType{*usedVestingType}, []int64{vested}, []int64{0})

	addTime := testutils.CreateDurationFromNumOfHours(500)
	if testType == withdrawableEnoughToSend {
		addTime = testutils.CreateDurationFromNumOfHours(2000)
	} else if testType == withdrawableNotEnoughToSend {
		addTime = testutils.CreateDurationFromNumOfHours(5600)
	} else if testType == allWithdrawableEnoughToSend {
		addTime = testutils.CreateDurationFromNumOfHours(7000)
	}
	oldCtx := ctx
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1000).WithBlockTime(ctx.BlockTime().Add(addTime))

	witdrawable := int64(0)
	if testType == withdrawableEnoughToSend {
		witdrawable = int64(200)
	} else if testType == withdrawableNotEnoughToSend {
		witdrawable = int64(920)
	} else if testType == allWithdrawableEnoughToSend {
		witdrawable = int64(1000)
	}
	verifyAccountVestings(t, oldCtx, app, accSrcAddr, []types.VestingType{*usedVestingType}, []int64{vested}, []int64{0})

	vestingData := getVestings(t, ctx, app, accSrcAddr)
	verifyVestingResponseWithStoredAccountVestings(t, ctx, app, vestingData, accSrcAddr, ctx.BlockTime(), delegationAllowed)

	sentAmount := int64(vested / 10)

	var delegableAccAddr sdk.AccAddress
	if delegationAllowed {
		accVestings, found := app.CfevestingKeeper.GetAccountVestings(ctx, accSrcAddr.String())
		require.EqualValues(t, true, found)

		delegableAccAddrLocal, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
		require.EqualValues(t, nil, err)
		delegableAccAddr = delegableAccAddrLocal
	}


	delegateAmount := uint64(0)
	if testType == delegatedEnoughToSend || testType == delegatedNotEnoughToSend {
		delegateAmount = uint64(300)

		if testType == delegatedNotEnoughToSend {
			delegateAmount = uint64(950)
		}
		
		delegate(t, ctx, app, accSrcAddr, delegableAccAddr, valAddr, delegateAmount, accInitBalance-vested, vested, accInitBalance-vested, int64(vested-delegateAmount))
		verifyDelegations(t, ctx, app, delegableAccAddr, []sdk.ValAddress{valAddr}, []int64{int64(delegateAmount)})

	}

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msgSendVesting := types.MsgSendVesting{FromAddress: accSrcAddr.String(), ToAddress: accDestAddr.String(), VestingId: 1, Amount: sdk.NewInt(sentAmount), RestartVesting: restartVesting}
	_, err := msgServer.SendVesting(msgServerCtx, &msgSendVesting)

	if testType == withdrawableNotEnoughToSend || testType == allWithdrawableEnoughToSend {
		require.NotEqualValues(t, nil, err)
		// require.EqualError()
		require.EqualError(t, err,
			"vesting available: "+strconv.FormatInt(vested-witdrawable, 10)+" is smaller than "+strconv.FormatInt(sentAmount, 10)+": insufficient funds")
	} else if testType == delegatedNotEnoughToSend {
		require.NotEqualValues(t, nil, err)
		// require.EqualError()
		require.EqualError(t, err,
			"vesting available: "+strconv.FormatUint(uint64(vested)-delegateAmount, 10)+
				" is smaller than "+strconv.FormatInt(sentAmount, 10)+" - probably delageted to validator.: insufficient funds")
	} else {
		require.EqualValues(t, nil, err)
	}
	wasSent := true
	if testType == withdrawableNotEnoughToSend || testType == allWithdrawableEnoughToSend || testType == delegatedNotEnoughToSend {
		wasSent = false
		sentAmount = int64(0)
	}
	if wasSent {
		verifyAccountVestingsWithModification(t, oldCtx, app, accSrcAddr, 2, []types.VestingType{*usedVestingType}, []time.Time{oldCtx.BlockTime()}, []int64{vested}, []int64{witdrawable},
			[]int64{sentAmount}, []time.Time{ctx.BlockTime()}, []int64{vested-sentAmount-witdrawable}, []int64{0})
	} else {
		verifyAccountVestings(t, oldCtx, app, accSrcAddr, []types.VestingType{*usedVestingType}, []int64{vested}, []int64{witdrawable})
	}
	
	verifyAccountBalance(t, app, ctx, accSrcAddr, sdk.NewInt(accInitBalance-vested+witdrawable))
	if delegationAllowed {
		verifyModuleAccount(t, ctx, app, sdk.NewInt(0))
		verifyAccountBalance(t, app, ctx, delegableAccAddr, sdk.NewInt(vested-sentAmount-witdrawable-int64(delegateAmount)))
	} else {
		verifyModuleAccount(t, ctx, app, sdk.NewInt(vested-witdrawable))

	}

	accVestings, accFound := app.CfevestingKeeper.GetAccountVestings(ctx, accDestAddr.String())
	require.EqualValues(t, wasSent, accFound)
	if !wasSent {
		return
	}

	if restartVesting {
		verifyAccountVestingsWithModification(t, ctx, app, accDestAddr, 2, []types.VestingType{*usedVestingType}, []time.Time{ctx.BlockTime()}, []int64{sentAmount}, []int64{0},
			[]int64{0}, []time.Time{ctx.BlockTime()}, []int64{sentAmount}, []int64{0})
	} else {
		verifyAccountVestingsWithModification(t, oldCtx, app, accDestAddr, 2, []types.VestingType{*usedVestingType}, []time.Time{ctx.BlockTime()}, []int64{sentAmount}, []int64{0},
			[]int64{0}, []time.Time{ctx.BlockTime()}, []int64{sentAmount}, []int64{0})
	}
	verifyAccountBalance(t, app, ctx, accDestAddr, sdk.NewInt(0))
	if delegationAllowed {
		verifyModuleAccount(t, ctx, app, sdk.NewInt(0))
		delegableAccAddrDst, _ := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
		verifyAccountBalance(t, app, ctx, delegableAccAddrDst, sdk.NewInt(sentAmount))
	} else {
		verifyModuleAccount(t, ctx, app, sdk.NewInt(vested-witdrawable))

	}

}
