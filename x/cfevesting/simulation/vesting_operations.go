package simulation

import (
	"cosmossdk.io/math"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateVestingOperations(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount1, _ := simtypes.RandomAcc(r, accs)
		randVestingPoolName := helpers.RandStringOfLengthCustomSeed(r, 10)
		randVestingAmount := math.NewInt(helpers.RandomInt(r, 10000000000))
		randVestingDuration := time.Duration(helpers.RandomInt(r, 3))
		randVesingId := helpers.RandomInt(r, 3)
		randomVestingType := "New vesting" + strconv.Itoa(int(randVesingId))

		msgCreateVestingPool := &types.MsgCreateVestingPool{
			Owner:       simAccount1.Address.String(),
			Name:        randVestingPoolName,
			Amount:      randVestingAmount,
			VestingType: randomVestingType,
			Duration:    randVestingDuration,
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CreateVestingPool(msgServerCtx, msgCreateVestingPool)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create vesting pool error", err.Error())
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting operations - create vesting pool", "", false, nil), nil, nil
		}
		simulationAddress2, _ := testcosmos.CreateRandomAccAddressHexAndBechNoBalance(helpers.RandomInt(r, 1000000))
		randMsgSendToVestinAccAmount := math.NewInt(helpers.RandomInt(r, 10000000))
		msgSendToVestingAccount := &types.MsgSendToVestingAccount{
			Owner:           simAccount1.Address.String(),
			ToAddress:       simulationAddress2,
			VestingPoolName: randVestingPoolName,
			Amount:          randMsgSendToVestinAccAmount,
			RestartVesting:  true,
		}
		_, err = msgServer.SendToVestingAccount(msgServerCtx, msgSendToVestingAccount)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Send to vesting account error", err.Error())
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting operations - send to vesting account", "", false, nil), nil, nil
		}

		msgWithdrawAllAvailable := &types.MsgWithdrawAllAvailable{
			Owner: simAccount1.Address.String(),
		}
		_, err = msgServer.WithdrawAllAvailable(msgServerCtx, msgWithdrawAllAvailable)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Withdraw all available error", err.Error())
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting operations - withdraw all available", "", false, nil), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Vesting operations - FINISHED")
		return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting operations simulation completed", "", true, nil), nil, nil
	}
}
