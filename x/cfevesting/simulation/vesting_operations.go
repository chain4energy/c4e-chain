package simulation

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateVestingPool(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		simAccount2, _ := simtypes.RandomAcc(r, accs)
		randName := helpers.RandStringOfLength(r, 10)
		randAmount := sdk.NewInt(helpers.RandomInt(r, 10000000000))
		randDuration := time.Duration(helpers.RandomInt(r, 3))
		randVesingId := helpers.RandomInt(r, 3)
		randAmountVestingAccount := helpers.RandomInt(r, 10)
		randomVestingType := "New vesting" + strconv.Itoa(int(randVesingId))

		msg := &types.MsgCreateVestingPool{
			Creator:     simAccount.Address.String(),
			Name:        randName,
			Amount:      randAmount,
			VestingType: randomVestingType,
			Duration:    randDuration,
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CreateVestingPool(msgServerCtx, msg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create vesting pool error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, nil
		}
		allAccVestings := k.GetAllAccountVestings(ctx)
		fmt.Println("ALL ACCOUNT VESTINGS", allAccVestings)

		msg2 := types.MsgSendToVestingAccount{FromAddress: simAccount.Address.String(), ToAddress: simAccount2.Address.String(),
			VestingId: 1, Amount: sdk.NewInt(randAmountVestingAccount), RestartVesting: false}
		_, err = msgServer.SendToVestingAccount(msgServerCtx, &msg2)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Send to vesting account error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "Vesting operations simulation completed", nil), nil, nil
	}
}
