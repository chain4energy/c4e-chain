package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
)

func SimulateWithdrawAllAvailable(
	_ types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		allVestingAccounts := k.GetAllAccountVestings(ctx)
		randInt := helpers.RandomInt(r, len(allVestingAccounts))
		accAddress := allVestingAccounts[randInt].Address
		msgWithdrawAllAvailable := &types.MsgWithdrawAllAvailable{
			Creator: accAddress,
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		withdraw, err := msgServer.WithdrawAllAvailable(msgServerCtx, msgWithdrawAllAvailable)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Withdraw all available error", err.Error())
			return simtypes.NewOperationMsg(msgWithdrawAllAvailable, false, "", nil), nil, nil
		}
		if withdraw.Withdrawn.Amount.Int64() > 0 {
			k.Logger(ctx).Debug("SIMULATION: Withdraw operations - FINISHED")
			return simtypes.NewOperationMsg(msgWithdrawAllAvailable, true, "", nil), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Withdraw operations - FINISHED")
		return simtypes.NewOperationMsgBasic(types.ModuleName, "withdraw_all_available", "", false, nil), nil, nil
	}
}
