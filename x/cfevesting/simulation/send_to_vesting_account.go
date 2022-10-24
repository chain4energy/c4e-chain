package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

)

func SimulateSendToVestingAccount(
	_ types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		allVestingPools := k.GetAllAccountVestingPools(ctx)
		randVestingPoolId := helpers.RandomInt(r, len(allVestingPools))
		accAddress := allVestingPools[randVestingPoolId].Address
		randMsgSendToVestinAccAmount := sdk.NewInt(helpers.RandomInt(r, 10))
		randInt := helpers.RandomInt(r, 1000000000)
		simAccount2Address := commontestutils.CreateRandomAccAddressNoBalance(randInt)
		randVestingId := helpers.RandomInt(r, len(allVestingPools[randVestingPoolId].VestingPools))
		msgSendToVestingAccount := &types.MsgSendToVestingAccount{
			FromAddress:    accAddress,
			ToAddress:      simAccount2Address,
			VestingId:      int32(randVestingId),
			Amount:         randMsgSendToVestinAccAmount,
			RestartVesting: true,
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.SendToVestingAccount(msgServerCtx, msgSendToVestingAccount)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Send to vesting account error", err.Error())
			return simtypes.NewOperationMsg(msgSendToVestingAccount, false, "", nil), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Send to vesting account - FINISHED")
		return simtypes.NewOperationMsg(msgSendToVestingAccount, true, "", nil), nil, nil
	}
}
