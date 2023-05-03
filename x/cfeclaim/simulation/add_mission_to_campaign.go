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

func SimulateMsgAddMissionToAidropCampaign(
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

		for i := int64(0); i < helpers.RandomInt(r, 3); i++ {
			randomWeight := helpers.RandomDecAmount(r, sdk.NewDec(1))
			addMissionToCampaignMsg := &types.MsgAddMissionToCampaign{
				Owner:          simAccount.Address.String(),
				CampaignId:     uint64(len(campaigns) - 1),
				Name:           helpers.RandStringOfLengthCustomSeed(r, 10),
				Description:    helpers.RandStringOfLengthCustomSeed(r, 10),
				MissionType:    types.MissionType(helpers.RandomInt(r, 4)),
				Weight:         &randomWeight,
				ClaimStartDate: nil,
			}

			_, err = msgServer.AddMissionToCampaign(msgServerCtx, addMissionToCampaignMsg)
			if err != nil {
				k.Logger(ctx).Error("SIMULATION: Add mission to campaign error", err.Error())
				return simtypes.NoOpMsg(types.ModuleName, addMissionToCampaignMsg.Type(), ""), nil, nil
			}
		}

		k.Logger(ctx).Debug("SIMULATION: Add mission to campaign campaign - added")
		return simtypes.NewOperationMsg(&types.MsgAddMissionToCampaign{}, true, "Add mission to campaign completed", nil), nil, nil
	}
}
