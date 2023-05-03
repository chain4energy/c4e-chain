package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"math/rand"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateCampaign(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		startTime := ctx.BlockTime().Add(time.Duration(helpers.RandIntBetween(r, 1000000, 10000000)))
		endTime := startTime.Add(time.Duration(helpers.RandIntBetween(r, 1000000, 10000000)))

		lockupPeriod := time.Duration(helpers.RandIntBetween(r, 1000000, 10000000))
		vestingPeriod := time.Duration(helpers.RandIntBetween(r, 1000000, 10000000))
		msg := &types.MsgCreateCampaign{
			Owner:                  simAccount.Address.String(),
			Name:                   helpers.RandStringOfLengthCustomSeed(r, 10),
			Description:            helpers.RandStringOfLengthCustomSeed(r, 10),
			CampaignType:           types.CampaignType(helpers.RandomInt(r, 4)),
			FeegrantAmount:         nil,
			InitialClaimFreeAmount: nil,
			StartTime:              &startTime,
			EndTime:                &endTime,
			LockupPeriod:           &lockupPeriod,
			VestingPeriod:          &vestingPeriod,
			VestingPoolName:        "",
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CreateCampaign(msgServerCtx, msg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), ""), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Create campaign - CREATED")
		return simtypes.NewOperationMsg(msg, true, "Create campaign simulation completed", nil), nil, nil
	}
}
