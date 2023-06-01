package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/utils"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
)

func SimulateMsgAddMission(
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
		campaigns := k.GetAllCampaigns(ctx)

		for i := int64(0); i < utils.RandInt64(r, 3); i++ {
			randomWeight := utils.RandomDecAmount(r, sdk.NewDec(1))
			AddMissionMsg := &types.MsgAddMission{
				Owner:          ownerAddress.String(),
				CampaignId:     uint64(len(campaigns) - 1),
				Name:           simtypes.RandStringOfLength(r, 10),
				Description:    simtypes.RandStringOfLength(r, 10),
				MissionType:    types.MissionType(utils.RandInt64(r, 4)),
				Weight:         &randomWeight,
				ClaimStartDate: nil,
			}

			_, err = msgServer.AddMission(msgServerCtx, AddMissionMsg)
			if err != nil {
				k.Logger(ctx).Error("SIMULATION: Add mission to campaign error", err.Error())
				return simtypes.NoOpMsg(types.ModuleName, AddMissionMsg.Type(), ""), nil, nil
			}
		}

		k.Logger(ctx).Debug("SIMULATION: Add mission to campaign campaign - added")
		return simtypes.NewOperationMsg(&types.MsgAddMission{}, true, "Add mission to campaign completed", nil), nil, nil
	}
}
