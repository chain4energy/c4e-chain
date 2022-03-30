package keeper_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

)

func TestVestDelegationNotAllowed(t *testing.T) {
	vestTest(t, false)
}

func TestVestDelegationAllowed(t *testing.T) {
	vestTest(t, true)
}

func TestVestLockupZeroDelegationNotAllowed(t *testing.T) {
	vestTestLockupZero(t, false)
}

func TestVestLockupZeroDelegationAllowed(t *testing.T) {
	vestTestLockupZero(t, true)
}

func vestTest(t *testing.T, delegationAllowed bool) {
	addHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := setupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, delegationAllowed, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	
	if delegationAllowed {
		makeVesting(t, ctx, app, accAddr, false, true, false, true, *usedVestingType, vested, accInitBalance, 0, 0, accInitBalance-vested, vested, 0)
	} else {
		makeVesting(t, ctx, app, accAddr, false, true, false, false, *usedVestingType, vested, accInitBalance, 0, 0, accInitBalance-vested, 0, vested)
	}

	verifyAccountVestings(t, ctx, app, accAddr, []types.VestingType{*usedVestingType}, []int64{vested}, []int64{0})

	if delegationAllowed {
		makeVesting(t, ctx, app, accAddr, true, true, true, true, *usedVestingType, vested, accInitBalance-vested, vested, 0, accInitBalance-2*vested, 2*vested, 0)
	} else {
		makeVesting(t, ctx, app, accAddr, true, true, false, false, *usedVestingType, vested, accInitBalance-vested, 0, vested, accInitBalance-2*vested, 0, 2*vested)
	}

	verifyAccountVestings(t, ctx, app, accAddr, []types.VestingType{*usedVestingType, *usedVestingType}, []int64{vested, vested}, []int64{0, 0})

}

func vestTestLockupZero(t *testing.T, delegationAllowed bool) {
	addHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := setupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	modifyVestingType := func(vt *types.VestingType) {
		vt.LockupPeriod = 0
	}
	// vestingTypes := setupVestingTypesWithModification(ctx, app, modifyVestingType, 1, 1, false, 1)

	vestingTypes := setupVestingTypesWithModification(ctx, app, modifyVestingType, 2, 1, delegationAllowed, 1)
	usedVestingType := vestingTypes.VestingTypes[0]
	
	if delegationAllowed {
		makeVesting(t, ctx, app, accAddr, false, true, false, true, *usedVestingType, vested, accInitBalance, 0, 0, accInitBalance-vested, vested, 0)
	} else {
		makeVesting(t, ctx, app, accAddr, false, true, false, false, *usedVestingType, vested, accInitBalance, 0, 0, accInitBalance-vested, 0, vested)
	}

	verifyAccountVestings(t, ctx, app, accAddr, []types.VestingType{*usedVestingType}, []int64{vested}, []int64{0})

	if delegationAllowed {
		makeVesting(t, ctx, app, accAddr, true, true, true, true, *usedVestingType, vested, accInitBalance-vested, vested, 0, accInitBalance-2*vested, 2*vested, 0)
	} else {
		makeVesting(t, ctx, app, accAddr, true, true, false, false, *usedVestingType, vested, accInitBalance-vested, 0, vested, accInitBalance-2*vested, 0, 2*vested)
	}

	verifyAccountVestings(t, ctx, app, accAddr, []types.VestingType{*usedVestingType, *usedVestingType}, []int64{vested, vested}, []int64{0, 0})

}

func TestVestFirstDelagtionNotAllowedSecondAllowed(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := setupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	// vestingTypes := setupVestingTypesWithModification(ctx, app, modifyVestingType, 1, 1, false, 1)

	// vestingTypes := setupVestingTypes(ctx, app, 2, 2, delegationAllowed, 1)

	i := 1
	modifyVestingType := func(vt *types.VestingType) {
		if (i == 1) {
			vt.DelegationsAllowed = true
		}
		i++
	}

	vestingTypes := setupVestingTypesWithModification(ctx, app, modifyVestingType, 2, 2, false, 1)
	delegableVestingType := vestingTypes.VestingTypes[0]
	nonDelegableVestingType := vestingTypes.VestingTypes[1]

	makeVesting(t, ctx, app, accAddr, false, true, false, false, *nonDelegableVestingType, vested, accInitBalance, 0, 0, accInitBalance-vested, 0, vested)

	verifyAccountVestings(t, ctx, app, accAddr, []types.VestingType{*nonDelegableVestingType}, []int64{vested}, []int64{0})

	makeVesting(t, ctx, app, accAddr, true, true, true, true, *delegableVestingType, vested, accInitBalance-vested, 0, vested, accInitBalance-2*vested, vested, vested)

	verifyAccountVestings(t, ctx, app, accAddr, []types.VestingType{*nonDelegableVestingType, *delegableVestingType}, []int64{vested, vested}, []int64{0, 0})

}

func TestVestingId(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000
	const accInitBalance = 10000
	app, ctx := setupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, true, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	addr := accAddr.String()

	k := app.CfevestingKeeper

	k.SetVestingTypes(ctx, vestingTypes)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgVest{Creator: addr, Amount: sdk.NewInt(vested), VestingType: usedVestingType.Name}
	_, error := msgServer.Vest(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVesting, accFound := k.GetAccountVestings(ctx, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 1, len(accVesting.Vestings))

	vesting := accVesting.Vestings[0]
	require.EqualValues(t, 1, vesting.Id)

	_, error = msgServer.Vest(msgServerCtx, &msg)

	require.EqualValues(t, nil, error)

	accVesting, accFound = k.GetAccountVestings(ctx, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 2, len(accVesting.Vestings))

	vesting = accVesting.Vestings[0]
	require.EqualValues(t, 1, vesting.Id)

	vesting = accVesting.Vestings[1]
	require.EqualValues(t, 2, vesting.Id)

	_, error = msgServer.Vest(msgServerCtx, &msg)

	require.EqualValues(t, nil, error)

	accVesting, accFound = k.GetAccountVestings(ctx, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 3, len(accVesting.Vestings))

	vesting = accVesting.Vestings[0]
	require.EqualValues(t, 1, vesting.Id)

	vesting = accVesting.Vestings[1]
	require.EqualValues(t, 2, vesting.Id)

	vesting = accVesting.Vestings[2]
	require.EqualValues(t, 3, vesting.Id)
}
