package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"math/rand"
)

func SimulateMsgAddMission(
	k keeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	cfevestingKeeper cfevestingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, err := createCampaign(ak, bk, cfevestingKeeper, app, r, ctx, accs, chainID)
		if err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Add mission - create campaign failure", "", false, nil), nil, nil
		}

		campaignsLen := len(k.GetAllCampaigns(ctx))
		for i := int64(0); i < utils.RandInt64(r, 3); i++ {
			randomWeight := utils.RandomDecAmount(r, sdk.NewDec(1))
			msgAddMission := &types.MsgAddMission{
				Owner:          simAccount.Address.String(),
				CampaignId:     uint64(campaignsLen - 1),
				Name:           simtypes.RandStringOfLength(r, 10),
				Description:    simtypes.RandStringOfLength(r, 10),
				MissionType:    types.MissionType(utils.RandInt64(r, 4)),
				Weight:         &randomWeight,
				ClaimStartDate: nil,
			}
			if err = simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, *simAccount, msgAddMission, chainID); err != nil {
				return simtypes.NewOperationMsg(msgAddMission, false, "", nil), nil, nil
			}

		}

		return simtypes.NewOperationMsg(&types.MsgAddMission{}, true, "", nil), nil, nil
	}
}
