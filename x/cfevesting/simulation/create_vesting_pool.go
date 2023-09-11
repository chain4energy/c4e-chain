package simulation

import (
	"github.com/chain4energy/c4e-chain/v2/testutil/simulation"
	"github.com/chain4energy/c4e-chain/v2/testutil/utils"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"math/rand"
	"strconv"
	"time"
)

func SimulateCreateVestingPool(
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

		msg := types.NewMsgCreateVestingPool(simAccount.Address.String(), randVestingPoolName, vestingPoolAmount.Amount, time.Hour,
			randomVestingType)

		if err = simulation.SendMessageWithFees(ctx, r, ak.(authkeeper.AccountKeeper), app, simAccount, msg, spendable.Sub(vestingPoolAmount), chainID); err != nil {
			return simtypes.NewOperationMsg(msg, false, "", nil), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
