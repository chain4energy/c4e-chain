package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
)

func SimulateWithdrawAllAvailable(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		allVestingAccounts := k.GetAllAccountVestingPools(ctx)
		if len(allVestingAccounts) == 0 {
			return simtypes.NewOperationMsg(&types.MsgWithdrawAllAvailable{}, false, "", nil), nil, nil
		}
		randInt := utils.RandInt64(r, len(allVestingAccounts))
		accAddress := allVestingAccounts[randInt].Owner
		sdkAccAddress := sdk.MustAccAddressFromBech32(accAddress)
		simAccount, _ := simtypes.FindAccount(accs, sdkAccAddress)
		msgWithdrawAllAvailable := &types.MsgWithdrawAllAvailable{
			Owner: accAddress,
		}

		if err := simulation.SendMessageWithRandomFees(ctx, r, ak, bk, app, simAccount, msgWithdrawAllAvailable, chainID); err != nil {
			return simtypes.NewOperationMsg(msgWithdrawAllAvailable, false, "", nil), nil, nil
		}
		return simtypes.NewOperationMsg(msgWithdrawAllAvailable, true, "", nil), nil, nil
	}
}
