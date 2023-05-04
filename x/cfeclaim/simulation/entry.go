package simulation

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"math/rand"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddClaimRecords(
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
		startCampaignMsg := &types.MsgStartCampaign{
			Owner:      ownerAddress.String(),
			CampaignId: uint64(len(campaigns) - 1),
		}

		_, err = msgServer.StartCampaign(msgServerCtx, startCampaignMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Start campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, startCampaignMsg.Type(), ""), nil, nil
		}

		addClaimRecordsMsg := &types.MsgAddClaimRecords{
			Owner:        ownerAddress.String(),
			CampaignId:   uint64(len(campaigns) - 1),
			ClaimRecords: createNClaimRecords(100, accs),
		}

		_, err = msgServer.AddClaimRecords(msgServerCtx, addClaimRecordsMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Add claim records campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, addClaimRecordsMsg.Type(), ""), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Add claim records - added")
		return simtypes.NewOperationMsg(addClaimRecordsMsg, true, "Add claim records simulation completed", nil), nil, nil
	}
}

func createNClaimRecords(n int, accs []simtypes.Account) []*types.ClaimRecord {
	var claimRecords []*types.ClaimRecord
	for i := 0; i < n; i++ {
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		claimRecordAccount, _ := simtypes.RandomAcc(r, accs)
		claimRecords = append(claimRecords, &types.ClaimRecord{
			Address: claimRecordAccount.Address.String(),
			Amount:  sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(helpers.RandomInt(r, 1000)))),
		})
	}
	return claimRecords
}
