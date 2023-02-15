package simulation

import (
	"math/rand"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateSendToVestingAccount(
	_ types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		allVestingPools := k.GetAllAccountVestingPools(ctx)
		if len(allVestingPools) == 0 {
			return simtypes.NewOperationMsg(&types.MsgSendToVestingAccount{}, false, "", nil), nil, nil
		}
		randVestingPoolId := helpers.RandomInt(r, len(allVestingPools))
		accAddress := allVestingPools[randVestingPoolId].Owner
		randMsgSendToVestinAccAmount := sdk.NewInt(helpers.RandomInt(r, 10))
		randInt := helpers.RandomInt(r, 1000000000)
		simAccount2Address := testcosmos.CreateRandomAccAddressNoBalance(randInt)
		numOfPools := len(allVestingPools[randVestingPoolId].VestingPools)
		var randVestingId int64 = 0
		if numOfPools > 1 {
			randVestingId = helpers.RandomInt(r, numOfPools-1)
		}
		msgSendToVestingAccount := &types.MsgSendToVestingAccount{
			Owner:           accAddress,
			ToAddress:       simAccount2Address,
			VestingPoolName: allVestingPools[randVestingPoolId].VestingPools[randVestingId].Name,
			Amount:          randMsgSendToVestinAccAmount,
			RestartVesting:  true,
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
