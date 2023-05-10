package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
)

func SimulateMsgAddMissionToCampaign(
	k keeper.Keeper,
	cfevestingKeeper cfevestingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		ownerAddress, err := createCampaign(k, cfevestingKeeper, msgServer, msgServerCtx, r, ctx, accs)
		if err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Create campaign", "", false, nil), nil, nil
		}
		campaigns := k.GetCampaigns(ctx)

		for i := int64(0); i < helpers.RandomInt(r, 3); i++ {
			randomWeight := helpers.RandomDecAmount(r, sdk.NewDec(1))
			addMissionToCampaignMsg := &types.MsgAddMissionToCampaign{
				Owner:          ownerAddress.String(),
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
