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
	"time"
)

func SimulateMsgClaim(
	k keeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	cfevestingKeeper cfevestingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, err := createCampaign(ak, bk, cfevestingKeeper, app, r, ctx, accs, chainID)
		if err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Claim - create campaign failure", "", false, nil), nil, nil
		}

		campaigns := k.GetAllCampaigns(ctx)
		campaignId := uint64(len(campaigns) - 1)
		randomWeight := utils.RandomDecAmount(r, sdk.NewDec(1))
		msgAddMission := &types.MsgAddMission{
			Owner:          simAccount.Address.String(),
			CampaignId:     campaignId,
			Name:           simtypes.RandStringOfLength(r, 10),
			Description:    simtypes.RandStringOfLength(r, 10),
			MissionType:    types.MissionType(utils.RandIntBetween(r, 2, 5)),
			Weight:         &randomWeight,
			ClaimStartDate: nil,
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, *simAccount, msgAddMission, chainID); err != nil {
			return simtypes.NewOperationMsg(msgAddMission, false, "", nil), nil, nil
		}
		startTime := ctx.BlockTime().Add(-time.Minute)
		msgEnableCampaign := &types.MsgEnableCampaign{
			Owner:      simAccount.Address.String(),
			CampaignId: campaignId,
			StartTime:  &startTime,
			EndTime:    nil,
		}
		if err = simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, *simAccount, msgEnableCampaign, chainID); err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Claim - enable campaign failure", "", false, nil), nil, nil
		}

		claimRecordEntries := createNClaimRecordEntries(r, accs, utils.RandIntBetween(r, 10, 100))
		addClaimRecordsMsg := &types.MsgAddClaimRecords{
			Owner:              simAccount.Address.String(),
			CampaignId:         campaignId,
			ClaimRecordEntries: claimRecordEntries,
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, *simAccount, addClaimRecordsMsg, chainID); err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Claim - add claim records failure", "", false, nil), nil, nil
		}

		claimerAddress := claimRecordEntries[utils.RandInt64(r, len(claimRecordEntries))].UserEntryAddress
		initialClaimMsg := &types.MsgInitialClaim{
			Claimer:            claimerAddress,
			CampaignId:         campaignId,
			DestinationAddress: claimerAddress,
		}

		simAccount2, found := simtypes.FindAccount(accs, sdk.MustAccAddressFromBech32(claimerAddress))
		if !found {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Claim - initial claim failure", "", false, nil), nil, nil
		}
		if err = simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount2, initialClaimMsg, chainID); err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Claim - initial claim failure", "", false, nil), nil, nil
		}

		claimMsg := &types.MsgClaim{
			Claimer:    initialClaimMsg.Claimer,
			CampaignId: campaignId,
			MissionId:  1,
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount2, claimMsg, chainID); err != nil {
			return simtypes.NewOperationMsg(claimMsg, false, "", nil), nil, nil
		}

		return simtypes.NewOperationMsg(claimMsg, true, "", nil), nil, nil
	}
}
