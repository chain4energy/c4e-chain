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

func SimulateVestingMultiOperations(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount1, _ := simtypes.RandomAcc(r, accs)

		randVestingDuration := time.Duration(helpers.RandomInt(r, 3))
		randVesingTypeId := helpers.RandomInt(r, 3)

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		multiOperationsCount := 4

		for i := 0; i < multiOperationsCount; i++ {
			randVestingName := helpers.RandStringOfLength(r, 10)
			randomVestingType := "New vesting" + strconv.Itoa(int(randVesingTypeId))
			randVestingAmount := sdk.NewInt(helpers.RandomInt(r, 100000000))
			msgCreateVestingPool := &types.MsgCreateVestingPool{
				Creator:     simAccount1.Address.String(),
				Name:        randVestingName,
				Amount:      randVestingAmount,
				VestingType: randomVestingType,
				Duration:    randVestingDuration,
			}
			_, err := msgServer.CreateVestingPool(msgServerCtx, msgCreateVestingPool)
			if err != nil {
				k.Logger(ctx).Error("SIMULATION: Create vesting pool error", err.Error())
				return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting multi operations - create vesting pool", "", false, nil), nil, nil
			}
		}

		for i := 1; i < multiOperationsCount; i++ {
			randMsgSendToVestinAccAmount := sdk.NewInt(helpers.RandomInt(r, 100))
			randInt := helpers.RandomInt(r, 10000000)
			simAccount2Address := helpers.CreateRandomAccAddressNoBalance(randInt)
			msgSendToVestingAccount := &types.MsgSendToVestingAccount{
				FromAddress:    simAccount1.Address.String(),
				ToAddress:      simAccount2Address,
				VestingId:      int32(i),
				Amount:         randMsgSendToVestinAccAmount,
				RestartVesting: true,
			}
			_, err := msgServer.SendToVestingAccount(msgServerCtx, msgSendToVestingAccount)
			if err != nil {
				k.Logger(ctx).Error("SIMULATION: Send to vesting account error", err.Error())
				return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting multi operations - send to vesting account", "", false, nil), nil, nil
			}
		}

		msgWithdrawAllAvailable := &types.MsgWithdrawAllAvailable{
			Creator: simAccount1.Address.String(),
		}
		_, err := msgServer.WithdrawAllAvailable(msgServerCtx, msgWithdrawAllAvailable)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Withdraw all available error", err.Error())
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting multi operations - withdraw all available", "", false, nil), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Vesting multi operations - FINISHED")
		return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting multi operations simulation completed", "", true, nil), nil, nil
	}
}
