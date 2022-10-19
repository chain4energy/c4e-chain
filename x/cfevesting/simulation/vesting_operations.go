package simulation

import (
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
		simAccount1, _ := simtypes.RandomAcc(r, accs)
		randInt := helpers.RandomInt(r, 10000000)
		simAccount2Address := helpers.CreateRandomAccAddressNoBalance(randInt)

		randVestingName := helpers.RandStringOfLength(r, 10)
		randVestingAmount := sdk.NewInt(helpers.RandomInt(r, 10000000000))
		randVestingDuration := time.Duration(helpers.RandomInt(r, 3))
		randVesingId := helpers.RandomInt(r, 3)
		randomVestingType := "New vesting" + strconv.Itoa(int(randVesingId))

		msg := &types.MsgCreateVestingPool{
			Creator:     simAccount1.Address.String(),
			Name:        randVestingName,
			Amount:      randVestingAmount,
			VestingType: randomVestingType,
			Duration:    randVestingDuration,
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CreateVestingPool(msgServerCtx, msg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create vesting pool error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, nil
		}

		randMsgSendAmount := sdk.NewInt(helpers.RandomInt(r, 10))
		msg2 := types.MsgSendToVestingAccount{
			FromAddress:    simAccount1.Address.String(),
			ToAddress:      simAccount2Address,
			VestingId:      1,
			Amount:         randMsgSendAmount,
			RestartVesting: true,
		}
		_, err = msgServer.SendToVestingAccount(msgServerCtx, &msg2)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Send to vesting account error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, nil
		}

		msg3 := types.MsgWithdrawAllAvailable{
			Creator: simAccount1.Address.String(),
		}
		_, err = msgServer.WithdrawAllAvailable(msgServerCtx, &msg3)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Withdraw all available error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, nil
		}

		k.Logger(ctx).Info("SIMULATION: Vesting operations - FINISHED")
		return simtypes.NewOperationMsg(msg, true, "Vesting operations simulation completed", nil), nil, nil
	}
}
