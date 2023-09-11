package simulation

import (
	"github.com/chain4energy/c4e-chain/v2/testutil/simulation"
	"github.com/chain4energy/c4e-chain/v2/testutil/utils"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateVestingOperations(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		randVestingPoolName := simtypes.RandStringOfLength(r, 10)
		randVesingId := utils.RandInt64(r, 5)
		randomVestingType := "New vesting" + strconv.Itoa(int(randVesingId))

		simAccount, _ := simtypes.RandomAcc(r, accs)
		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		if !spendable.IsAllPositive() {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateVestingPool, "balance is negative"), nil, nil
		}
		amount, err := simtypes.RandPositiveInt(r, spendable.AmountOf(sdk.DefaultBondDenom))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateVestingPool, "unable to generate positive amount"), nil, err
		}
		vestingPoolAmount := sdk.NewCoin(sdk.DefaultBondDenom, amount)

		msgCreateVestingPool := types.NewMsgCreateVestingPool(simAccount.Address.String(), randVestingPoolName, amount, time.Hour,
			randomVestingType)

		err = simulation.SendMessageWithFees(ctx, r, ak.(authkeeper.AccountKeeper), app, simAccount, msgCreateVestingPool, spendable.Sub(vestingPoolAmount), chainID)
		if err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting operations - create vesting pool", "", false, nil), nil, nil
		}

		simAccount2 := simtypes.RandomAccounts(r, 1)[0]
		spendable = bk.SpendableCoins(ctx, simAccount2.Address)
		amount, err = simtypes.RandPositiveInt(r, amount)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateVestingPool, "unable to generate positive amount"), nil, err
		}

		msgSendToVestingAccount := &types.MsgSendToVestingAccount{
			Owner:           simAccount.Address.String(),
			ToAddress:       simAccount2.Address.String(),
			VestingPoolName: randVestingPoolName,
			Amount:          amount,
			RestartVesting:  true,
		}
		err = simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgSendToVestingAccount, chainID)
		if err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting operations - send to vesting account", "", false, nil), nil, nil
		}

		msgWithdrawAllAvailable := &types.MsgWithdrawAllAvailable{
			Owner: simAccount.Address.String(),
		}
		if err = simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgWithdrawAllAvailable, chainID); err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting operations - withdraw all available", "", false, nil), nil, nil
		}

		return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting operations simulation completed", "", true, nil), nil, nil
	}
}
