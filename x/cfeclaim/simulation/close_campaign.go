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

func SimulateMsgCloseCampaign(
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
		msgCreateCampaign := &types.MsgCreateCampaign{
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
		_, err := msgServer.CreateCampaign(msgServerCtx, msgCreateCampaign)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msgCreateCampaign.Type(), ""), nil, nil
		}
		campaigns := k.GetCampaigns(ctx)
		closeCampaignMsg := &types.MsgCloseCampaign{
			Owner:               simAccount.Address.String(),
			CampaignId:          uint64(len(campaigns) - 1),
			CampaignCloseAction: types.CloseAction(helpers.RandomInt(r, 4)),
		}

		_, err = msgServer.CloseCampaign(msgServerCtx, closeCampaignMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Close campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, closeCampaignMsg.Type(), ""), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Close campaign - closed")
		return simtypes.NewOperationMsg(closeCampaignMsg, true, "Close campaign simulation completed", nil), nil, nil
	}
}
