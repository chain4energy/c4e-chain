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

func SimulateMsgStartCampaign(
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

		var startCampaignOwner string
		if helpers.RandomInt(r, 2) == 1 {
			simAccount2, _ := simtypes.RandomAcc(r, accs)
			startCampaignOwner = simAccount2.Address.String()
		} else {
			startCampaignOwner = ownerAddress.String()
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