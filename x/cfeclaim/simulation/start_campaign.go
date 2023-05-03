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

func SimulateMsgStartCampaign(
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

		var startCampaignOwner string
		if helpers.RandomInt(r, 2) == 1 {
			simAccount2, _ := simtypes.RandomAcc(r, accs)
			startCampaignOwner = simAccount2.Address.String()
		} else {
			startCampaignOwner = simAccount.Address.String()
		}

		campaigns := k.GetCampaigns(ctx)
		startCampaignMsg := &types.MsgStartCampaign{
			Owner:      startCampaignOwner,
			CampaignId: uint64(len(campaigns) - 1),
		}

		_, err = msgServer.StartCampaign(msgServerCtx, startCampaignMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Start campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, startCampaignMsg.Type(), ""), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Start campaign - started")
		return simtypes.NewOperationMsg(startCampaignMsg, true, "Start campaign simulation completed", nil), nil, nil
	}
}
