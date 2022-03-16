package keeper_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/chain4energy/c4e-chain/app"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestWithdrawAllAvailableOnVestingStart(t *testing.T) {
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accAddr, _ := sdk.AccAddressFromBech32(addr)

	const vt1 = "test1"
	const initBlock = 1000
	const vested = 1000000
	addHelperModuleAccountPerms()

	accountVestings, vesting1 := createAccountVestings(addr, vt1, vested, 0)

	app, ctx := setupApp(initBlock)

	bank := app.BankKeeper
	auth := app.AccountKeeper

	denom := addCoinsToModule(vested, helperModuleAccount, ctx, bank)

	k := app.CfevestingKeeper

	k.SetAccountVestings(ctx, accountVestings)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: addr}
	_, error := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVestings := k.GetAllAccountVestings(ctx)

	verifyAcountVestings(k, ctx, addr, t, accVestings, 1)
	vesting := accVestings[0].Vestings[0]
	verifyVesting(t, *vesting1, *vesting)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, 0)

	verifyModuleAccount(auth, ctx, bank, denom, t, vested)

}

func TestWithdrawAllAvailableManyVestingsOnVestingStart(t *testing.T) {
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accAddr, _ := sdk.AccAddressFromBech32(addr)

	const vt1 = "test1"
	const vt2 = "test2"
	const vt3 = "test3"
	const initBlock = 1000
	const vested = 1000000
	addHelperModuleAccountPerms()

	accountVestings, vesting1, vesting2, vesting3 := createAccountVestingsMany(addr, vt1, vt2, vt3, vested, 0)

	app, ctx := setupApp(initBlock)

	bank := app.BankKeeper
	auth := app.AccountKeeper

	denom := addCoinsToModule(3*vested, helperModuleAccount, ctx, bank)

	k := app.CfevestingKeeper

	k.SetAccountVestings(ctx, accountVestings)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: addr}
	_, error := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVestings := k.GetAllAccountVestings(ctx)

	verifyAcountVestings(k, ctx, addr, t, accVestings, 3)
	vesting := accVestings[0].Vestings[0]
	verifyVesting(t, *vesting1, *vesting)

	vesting = accVestings[0].Vestings[1]
	verifyVesting(t, *vesting2, *vesting)

	vesting = accVestings[0].Vestings[2]
	verifyVesting(t, *vesting3, *vesting)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, 0)

	verifyModuleAccount(auth, ctx, bank, denom, t, 3*vested)

}

func TestWithdrawAllAvailableSomeToWithdraw(t *testing.T) {
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accAddr, _ := sdk.AccAddressFromBech32(addr)

	const vt1 = "test1"
	const initBlock = 10100
	const vested = 1000000
	addHelperModuleAccountPerms()

	accountVestings, vesting1 := createAccountVestings(addr, vt1, vested, 0)

	app, ctx := setupApp(initBlock)

	bank := app.BankKeeper
	auth := app.AccountKeeper

	denom := addCoinsToModule(vested, helperModuleAccount, ctx, bank)

	k := app.CfevestingKeeper

	k.SetAccountVestings(ctx, accountVestings)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: addr}
	_, error := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVestings := k.GetAllAccountVestings(ctx)

	verifyAcountVestings(k, ctx, addr, t, accVestings, 1)
	vesting := accVestings[0].Vestings[0]
	vesting1.Withdrawn = 1000
	verifyVesting(t, *vesting1, *vesting)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, vesting1.Withdrawn)

	verifyModuleAccount(auth, ctx, bank, denom, t, vested-vesting1.Withdrawn)

}

func TestWithdrawAllAvailableManyVestedSomeToWithdraw(t *testing.T) {
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accAddr, _ := sdk.AccAddressFromBech32(addr)

	const vt1 = "test1"
	const vt2 = "test2"
	const vt3 = "test3"
	const initBlock = 10100
	const vested = 1000000
	addHelperModuleAccountPerms()

	accountVestings, vesting1, vesting2, vesting3 := createAccountVestingsMany(addr, vt1, vt2, vt3, vested, 0)

	app, ctx := setupApp(initBlock)

	bank := app.BankKeeper
	auth := app.AccountKeeper

	denom := addCoinsToModule(3*vested, helperModuleAccount, ctx, bank)

	k := app.CfevestingKeeper

	k.SetAccountVestings(ctx, accountVestings)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: addr}
	_, error := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVestings := k.GetAllAccountVestings(ctx)

	verifyAcountVestings(k, ctx, addr, t, accVestings, 3)
	vesting1.Withdrawn = 1000
	vesting2.Withdrawn = 1000
	vesting3.Withdrawn = 1000
	vesting := accVestings[0].Vestings[0]
	verifyVesting(t, *vesting1, *vesting)

	vesting = accVestings[0].Vestings[1]
	verifyVesting(t, *vesting2, *vesting)

	vesting = accVestings[0].Vestings[2]
	verifyVesting(t, *vesting3, *vesting)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, 3*vesting1.Withdrawn)

	verifyModuleAccount(auth, ctx, bank, denom, t, 3*vested-3*vesting1.Withdrawn)

}

func TestWithdrawAllAvailableSomeToWithdrawAndSomeWithdrawn(t *testing.T) {
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accAddr, _ := sdk.AccAddressFromBech32(addr)

	const vt1 = "test1"
	const initBlock = 10100
	const vested = 1000000
	const withdrawn = 300
	addHelperModuleAccountPerms()

	accountVestings, vesting1 := createAccountVestings(addr, vt1, vested, withdrawn)

	app, ctx := setupApp(initBlock)

	bank := app.BankKeeper
	auth := app.AccountKeeper

	denom := addCoinsToModule(vested, helperModuleAccount, ctx, bank)

	k := app.CfevestingKeeper

	k.SetAccountVestings(ctx, accountVestings)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: addr}
	_, error := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVestings := k.GetAllAccountVestings(ctx)

	verifyAcountVestings(k, ctx, addr, t, accVestings, 1)
	vesting := accVestings[0].Vestings[0]
	vesting1.Withdrawn = 1000
	verifyVesting(t, *vesting1, *vesting)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, vesting1.Withdrawn-withdrawn)

	verifyModuleAccount(auth, ctx, bank, denom, t, vested-vesting1.Withdrawn+withdrawn)

}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAndSomeWithdrawn(t *testing.T) {
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accAddr, _ := sdk.AccAddressFromBech32(addr)

	const vt1 = "test1"
	const vt2 = "test2"
	const vt3 = "test3"
	const initBlock = 10100
	const vested = 1000000
	const withdrawn = 300
	addHelperModuleAccountPerms()

	accountVestings, vesting1, vesting2, vesting3 := createAccountVestingsMany(addr, vt1, vt2, vt3, vested, withdrawn)

	app, ctx := setupApp(initBlock)

	bank := app.BankKeeper
	auth := app.AccountKeeper

	denom := addCoinsToModule(3*vested, helperModuleAccount, ctx, bank)

	k := app.CfevestingKeeper

	k.SetAccountVestings(ctx, accountVestings)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: addr}
	_, error := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVestings := k.GetAllAccountVestings(ctx)

	verifyAcountVestings(k, ctx, addr, t, accVestings, 3)
	vesting := accVestings[0].Vestings[0]
	vesting1.Withdrawn = 1000
	vesting2.Withdrawn = 1000
	vesting3.Withdrawn = 1000
	verifyVesting(t, *vesting1, *vesting)

	vesting = accVestings[0].Vestings[1]
	verifyVesting(t, *vesting2, *vesting)

	vesting = accVestings[0].Vestings[2]
	verifyVesting(t, *vesting3, *vesting)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, 3*(vesting1.Withdrawn-withdrawn))

	verifyModuleAccount(auth, ctx, bank, denom, t, 3*vested-3*(vesting1.Withdrawn-withdrawn))

}

func TestVestAndWithdrawAllAvailable(t *testing.T) {
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"

	accAddr, _ := sdk.AccAddressFromBech32(addr)

	const vt1 = "test1"
	const initBlock = 1000

	const vested = 1000000
	addHelperModuleAccountPerms()

	app, ctx := setupApp(initBlock)

	bank := app.BankKeeper
	auth := app.AccountKeeper

	denom := addCoinsToAccount(vested, helperModuleAccount, ctx, bank, accAddr)

	k := app.CfevestingKeeper

	vestingTypes := types.VestingTypes{}
	vestingType1 := types.VestingType{
		Name:                 vt1,
		LockupPeriod:         9000,
		VestingPeriod:        100000,
		TokenReleasingPeriod: 10,
		DelegationsAllowed:   false,
	}
	vestingTypesArray := []*types.VestingType{&vestingType1}
	vestingTypes.VestingTypes = vestingTypesArray
	k.SetVestingTypes(ctx, vestingTypes)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgVest{Creator: addr, Amount: vested, VestingType: vt1}
	_, error := msgServer.Vest(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	msgWithdraw := types.MsgWithdrawAllAvailable{Creator: addr}
	_, error = msgServer.WithdrawAllAvailable(msgServerCtx, &msgWithdraw)
	require.EqualValues(t, nil, error)

	accVestings := k.GetAllAccountVestings(ctx)

	verifyAcountVestings(k, ctx, addr, t, accVestings, 1)
	vesting := accVestings[0].Vestings[0]
	vesting1 := types.Vesting{
		VestingType:          vt1,
		VestingStartBlock:    1000,
		LockEndBlock:         10000,
		VestingEndBlock:      110000,
		Vested:               vested,
		Claimable:            0,
		LastFreeingBlock:     0,
		FreeCoinsBlockPeriod: 10,
		FreeCoinsPerPeriod:   100,
		DelegationAllowed:    false,
		Withdrawn:            0,
	}
	verifyVesting(t, vesting1, *vesting)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, 0)

	verifyModuleAccount(auth, ctx, bank, denom, t, vested)

	ctx = ctx.WithBlockHeight(int64(10100))
	msgServerCtx = sdk.WrapSDKContext(ctx)

	msgWithdraw = types.MsgWithdrawAllAvailable{Creator: addr}
	_, error = msgServer.WithdrawAllAvailable(msgServerCtx, &msgWithdraw)
	require.EqualValues(t, nil, error)

	accVestings = k.GetAllAccountVestings(ctx)

	verifyAcountVestings(k, ctx, addr, t, accVestings, 1)
	vesting = accVestings[0].Vestings[0]
	vesting1 = types.Vesting{
		VestingType:          vt1,
		VestingStartBlock:    1000,
		LockEndBlock:         10000,
		VestingEndBlock:      110000,
		Vested:               vested,
		Claimable:            0,
		LastFreeingBlock:     0,
		FreeCoinsBlockPeriod: 10,
		FreeCoinsPerPeriod:   100,
		DelegationAllowed:    false,
		Withdrawn:            1000,
	}
	verifyVesting(t, vesting1, *vesting)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, 1000)

	verifyModuleAccount(auth, ctx, bank, denom, t, vested-1000)
}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllDelegable(t *testing.T) {
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"
	const delagableAddr = "cosmos1dfugyfm087qa3jrdglkeaew0wkn59jk8mgw6x6"

	accAddr, _ := sdk.AccAddressFromBech32(addr)
	delegableAccAddr, _ := sdk.AccAddressFromBech32(delagableAddr)

	const vt1 = "test1"
	const vt2 = "test2"
	const vt3 = "test3"
	const initBlock = 10100
	const vested = 1000000
	accountVestings, vesting1, vesting2, vesting3 := createAccountVestingsMany(addr, vt1, vt2, vt3, vested, 0)
	accountVestings.DelegableAddress = delagableAddr
	vesting1.DelegationAllowed = true
	vesting2.DelegationAllowed = true
	vesting3.DelegationAllowed = true
	addHelperModuleAccountPerms()

	app, ctx := setupApp(initBlock)

	bank := app.BankKeeper
	// mint := app.MintKeeper
	auth := app.AccountKeeper

	// denom := addCoinsToModule(3*vested, mint, ctx, bank)
	denom := addCoinsToAccount(3*vested, helperModuleAccount, ctx, bank, delegableAccAddr)

	k := app.CfevestingKeeper

	k.SetAccountVestings(ctx, accountVestings)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: addr}
	_, error := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVestings := k.GetAllAccountVestings(ctx)

	verifyAcountVestings(k, ctx, addr, t, accVestings, 3)
	vesting1.Withdrawn = 1000
	vesting2.Withdrawn = 1000
	vesting3.Withdrawn = 1000
	vesting := accVestings[0].Vestings[0]
	verifyVesting(t, *vesting1, *vesting)

	vesting = accVestings[0].Vestings[1]
	verifyVesting(t, *vesting2, *vesting)

	vesting = accVestings[0].Vestings[2]
	verifyVesting(t, *vesting3, *vesting)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, 3*vesting1.Withdrawn)
	verifyAccountBalance(t, bank, ctx, delegableAccAddr, denom, 3*vested-3*vesting1.Withdrawn)

	verifyModuleAccount(auth, ctx, bank, denom, t, 0)

}

func TestWithdrawAllAvailableManyVestedSomeToWithdrawAllSomeDelegable(t *testing.T) {
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"
	const delagableAddr = "cosmos1dfugyfm087qa3jrdglkeaew0wkn59jk8mgw6x6"

	accAddr, _ := sdk.AccAddressFromBech32(addr)
	delegableAccAddr, _ := sdk.AccAddressFromBech32(delagableAddr)

	const vt1 = "test1"
	const vt2 = "test2"
	const vt3 = "test3"
	const initBlock = 10100
	const vested = 1000000
	addHelperModuleAccountPerms()

	accountVestings, vesting1, vesting2, vesting3 := createAccountVestingsMany(addr, vt1, vt2, vt3, vested, 0)
	accountVestings.DelegableAddress = delagableAddr

	vesting3.DelegationAllowed = true

	app, ctx := setupApp(initBlock)

	bank := app.BankKeeper
	// mint := app.MintKeeper
	auth := app.AccountKeeper

	addCoinsToModule(2*vested, helperModuleAccount, ctx, bank)
	denom := addCoinsToAccount(vested, helperModuleAccount, ctx, bank, delegableAccAddr)

	k := app.CfevestingKeeper

	k.SetAccountVestings(ctx, accountVestings)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: addr}
	_, error := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVestings := k.GetAllAccountVestings(ctx)

	verifyAcountVestings(k, ctx, addr, t, accVestings, 3)
	vesting1.Withdrawn = 1000
	vesting2.Withdrawn = 1000
	vesting3.Withdrawn = 1000
	vesting := accVestings[0].Vestings[0]
	verifyVesting(t, *vesting1, *vesting)

	vesting = accVestings[0].Vestings[1]
	verifyVesting(t, *vesting2, *vesting)

	vesting = accVestings[0].Vestings[2]
	verifyVesting(t, *vesting3, *vesting)

	verifyAccountBalance(t, bank, ctx, accAddr, denom, 3*vesting1.Withdrawn)
	verifyAccountBalance(t, bank, ctx, delegableAccAddr, denom, vested-vesting1.Withdrawn)

	verifyModuleAccount(auth, ctx, bank, denom, t, 2*vested-2*vesting1.Withdrawn)

}

func verifyAcountVestings(k keeper.Keeper, ctx sdk.Context, addr string, t *testing.T, accVestings []types.AccountVestings, numOfVestings int) {
	_, accFound := k.GetAccountVestings(ctx, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 1, len(accVestings))
	require.EqualValues(t, numOfVestings, len(accVestings[0].Vestings))

	require.EqualValues(t, addr, accVestings[0].Address)
}

func verifyVesting(t *testing.T, vestingExpected types.Vesting, vestingActual types.Vesting) {
	require.EqualValues(t, vestingExpected, vestingActual)

	require.EqualValues(t, vestingExpected.VestingType, vestingActual.VestingType)
	require.EqualValues(t, vestingExpected.VestingStartBlock, vestingActual.VestingStartBlock)
	require.EqualValues(t, vestingExpected.LockEndBlock, vestingActual.LockEndBlock)

	require.EqualValues(t, vestingExpected.VestingEndBlock, vestingActual.VestingEndBlock)

	require.EqualValues(t, vestingExpected.Vested, vestingActual.Vested)
	require.EqualValues(t, vestingExpected.Claimable, vestingActual.Claimable)
	require.EqualValues(t, vestingExpected.LastFreeingBlock, vestingActual.LastFreeingBlock)
	require.EqualValues(t, vestingExpected.FreeCoinsBlockPeriod, vestingActual.FreeCoinsBlockPeriod)
	require.EqualValues(t, vestingExpected.FreeCoinsPerPeriod, vestingActual.FreeCoinsPerPeriod)
	require.EqualValues(t, vestingExpected.DelegationAllowed, vestingActual.DelegationAllowed)
}

func setupApp(initBlock int64) (*app.App, sdk.Context) {
	app := app.Setup(false)
	header := tmproto.Header{}
	header.Height = initBlock
	ctx := app.BaseApp.NewContext(false, header)
	return app, ctx
}

func createAccountVestingsMany(addr string, vt1 string, vt2 string, vt3 string, vested uint64, withdrawn uint64) (types.AccountVestings, *types.Vesting, *types.Vesting, *types.Vesting) {
	accountVestings := types.AccountVestings{}
	accountVestings.Address = addr
	vesting1 := types.Vesting{
		VestingType:          vt1,
		VestingStartBlock:    1000,
		LockEndBlock:         10000,
		VestingEndBlock:      110000,
		Vested:               vested,
		Claimable:            0,
		LastFreeingBlock:     0,
		FreeCoinsBlockPeriod: 10,
		FreeCoinsPerPeriod:   100,
		DelegationAllowed:    false,
		Withdrawn:            withdrawn,
	}
	vesting2 := types.Vesting{
		VestingType:          vt1,
		VestingStartBlock:    1000,
		LockEndBlock:         10000,
		VestingEndBlock:      110000,
		Vested:               vested,
		Claimable:            0,
		LastFreeingBlock:     0,
		FreeCoinsBlockPeriod: 10,
		FreeCoinsPerPeriod:   100,
		DelegationAllowed:    false,
		Withdrawn:            withdrawn,
	}
	vesting3 := types.Vesting{
		VestingType:          vt1,
		VestingStartBlock:    1000,
		LockEndBlock:         10000,
		VestingEndBlock:      110000,
		Vested:               vested,
		Claimable:            0,
		LastFreeingBlock:     0,
		FreeCoinsBlockPeriod: 10,
		FreeCoinsPerPeriod:   100,
		DelegationAllowed:    false,
		Withdrawn:            withdrawn,
	}

	vestingsArray := []*types.Vesting{&vesting1, &vesting2, &vesting3}
	accountVestings.Vestings = vestingsArray
	return accountVestings, &vesting1, &vesting2, &vesting3
}

func addCoinsToModule(vested uint64, mintTo string, ctx sdk.Context, bank bankkeeper.Keeper) string {
	return addCoinsToModuleByName(vested, types.ModuleName, mintTo, ctx, bank)
}
