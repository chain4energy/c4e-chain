package simulation

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"math/rand"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMoveAvailableVesting(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		simAccount2, _ := simtypes.RandomAcc(r, accs)
		simAccount3, _ := simtypes.RandomAcc(r, accs)

		randCoinsAmount := math.NewInt(helpers.RandomInt(r, 1000))
		coin := sdk.NewCoin(sdk.DefaultBondDenom, randCoinsAmount)
		coins := sdk.NewCoins(coin)

		randomEndDurationToAdd := time.Duration(helpers.RandomInt(r, 10000000000))
		randomStartDurationToSub := time.Duration(helpers.RandomInt(r, 1000000000000))
		startTime := ctx.BlockTime()
		msg := &types.MsgCreateVestingAccount{
			FromAddress: simAccount.Address.String(),
			ToAddress:   simAccount2.Address.String(),
			StartTime:   startTime.Add(-randomStartDurationToSub).Unix(),
			EndTime:     startTime.Add(randomEndDurationToAdd).Unix(),
			Amount:      coins,
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CreateVestingAccount(msgServerCtx, msg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create vesting account error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), ""), nil, nil
		}

		msgSplitVesting := &types.MsgMoveAvailableVesting{
			FromAddress: simAccount2.Address.String(),
			ToAddress:   simAccount3.Address.String(),
		}

		_, err = msgServer.MoveAvailableVesting(msgServerCtx, msgSplitVesting)
		if err != nil {
			if err != nil {
				k.Logger(ctx).Error("SIMULATION: Move available vesting error", err.Error())
			}

			return simtypes.NewOperationMsg(msgSplitVesting, false, "", nil), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Move available vesting - FINISHED")
		return simtypes.NewOperationMsg(msgSplitVesting, true, "", nil), nil, nil
	}
}
