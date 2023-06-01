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
	"math/rand"
)

func SimulateMsgEnableCampaign(
	k keeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	cfevestingKeeper cfevestingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, err := createCampaign(ak, bk, cfevestingKeeper, app, r, ctx, accs, chainID)
		if err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Enable campaign - create campaign failure", "", false, nil), nil, nil
		}

		var EnableCampaignOwner string
		if utils.RandInt64(r, 2) == 1 {
			simAccount2, _ := simtypes.RandomAcc(r, accs)
			EnableCampaignOwner = simAccount2.Address.String()
		} else {
			EnableCampaignOwner = simAccount.Address.String()
		}

		campaigns := k.GetAllCampaigns(ctx)
		msgEnableCampaign := &types.MsgEnableCampaign{
			Owner:      EnableCampaignOwner,
			CampaignId: uint64(len(campaigns) - 1),
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak, bk, app, *simAccount, msgEnableCampaign, chainID); err != nil {
			return simtypes.NewOperationMsg(msgEnableCampaign, false, "", nil), nil, nil
		}

		return simtypes.NewOperationMsg(msgEnableCampaign, true, "", nil), nil, nil
	}
}
