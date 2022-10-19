package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
	"time"
)

func SimulateMsgCreateVestingAccount(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		simAccount2, _ := simtypes.RandomAcc(r, accs)
		randAmount := sdk.NewInt(helpers.RandomInt(r, 10000000000))
		coin := sdk.NewCoin("stake", randAmount)
		coins := sdk.NewCoins(coin)
		msg := &types.MsgCreateVestingAccount{
			FromAddress: simAccount.Address.String(),
			ToAddress:   simAccount2.Address.String(),
			EndTime:     time.Now().Add(1).Unix(),
			StartTime:   time.Now().Unix(),
			Amount:      coins,
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CreateVestingAccount(msgServerCtx, msg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create vesting account error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "Create vesting account simulation completed", nil), nil, nil
	}
}
