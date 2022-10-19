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
	_ types.AccountKeeper,
	_ types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		randInt := helpers.RandomInt(r, 100000)
		simAccount2Address := helpers.CreateRandomAccAddressNoBalance(randInt)

		randCoinsAmount := sdk.NewInt(helpers.RandomInt(r, 1000))
		coin := sdk.NewCoin("stake", randCoinsAmount)
		coins := sdk.NewCoins(coin)

		randomStartDurationAdd := time.Duration(helpers.RandomInt(r, 1000000))
		randomStartDurationEnd := time.Duration(helpers.RandIntBetween(r, 1000000, 10000000))

		msg := &types.MsgCreateVestingAccount{
			FromAddress: simAccount.Address.String(),
			ToAddress:   simAccount2Address,
			EndTime:     time.Now().Add(randomStartDurationAdd).Unix(),
			StartTime:   time.Now().Add(randomStartDurationEnd).Unix(),
			Amount:      coins,
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CreateVestingAccount(msgServerCtx, msg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create vesting account error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), ""), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Create vesting account - CREATED")
		return simtypes.NewOperationMsg(msg, true, "Create vesting account simulation completed", nil), nil, nil
	}
}
