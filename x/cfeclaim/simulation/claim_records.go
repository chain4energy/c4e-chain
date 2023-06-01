package simulation

import (
	"cosmossdk.io/math"
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

func SimulateMsgAddClaimRecords(
	k keeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	cfevestingKeeper cfevestingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, err := createCampaign(ak, bk, cfevestingKeeper, app, r, ctx, accs, chainID)
		if err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Add claim records - create campaign failure", "", false, nil), nil, nil
		}

		campaigns := k.GetAllCampaigns(ctx)
		msgEnableCampaign := &types.MsgEnableCampaign{
			Owner:      simAccount.Address.String(),
			CampaignId: uint64(len(campaigns) - 1),
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak, bk, app, *simAccount, msgEnableCampaign, chainID); err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Add claim records - enable campaign failure", "", false, nil), nil, nil
		}

		addClaimRecordsMsg := &types.MsgAddClaimRecords{
			Owner:              simAccount.Address.String(),
			CampaignId:         uint64(len(campaigns) - 1),
			ClaimRecordEntries: createNClaimRecordEntries(r, accs, utils.RandIntBetween(r, 10, 100)),
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak, bk, app, *simAccount, addClaimRecordsMsg, chainID); err != nil {
			return simtypes.NewOperationMsg(addClaimRecordsMsg, false, "", nil), nil, nil
		}

		return simtypes.NewOperationMsg(addClaimRecordsMsg, true, "", nil), nil, nil
	}
}

func SimulateMsgDeleteClaimRecord(
	k keeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	cfevestingKeeper cfevestingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, err := createCampaign(ak, bk, cfevestingKeeper, app, r, ctx, accs, chainID)
		if err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Delete claim record - create campaign failure", "", false, nil), nil, nil
		}

		campaigns := k.GetAllCampaigns(ctx)
		msgEnableCampaign := &types.MsgEnableCampaign{
			Owner:      simAccount.Address.String(),
			CampaignId: uint64(len(campaigns) - 1),
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak, bk, app, *simAccount, msgEnableCampaign, chainID); err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Delete claim record - enable campaign failure", "", false, nil), nil, nil
		}

		claimRecordEntries := createNClaimRecordEntries(r, accs, utils.RandIntBetween(r, 10, 100))
		addClaimRecordsMsg := &types.MsgAddClaimRecords{
			Owner:              simAccount.Address.String(),
			CampaignId:         uint64(len(campaigns) - 1),
			ClaimRecordEntries: claimRecordEntries,
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak, bk, app, *simAccount, addClaimRecordsMsg, chainID); err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Delete claim record - add claim record entries failure", "", false, nil), nil, nil
		}

		deleteClaimRecordMsg := &types.MsgDeleteClaimRecord{
			Owner:       simAccount.Address.String(),
			CampaignId:  uint64(len(campaigns) - 1),
			UserAddress: claimRecordEntries[utils.RandInt64(r, len(claimRecordEntries))].UserEntryAddress,
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak, bk, app, *simAccount, deleteClaimRecordMsg, chainID); err != nil {
			return simtypes.NewOperationMsg(deleteClaimRecordMsg, false, "", nil), nil, nil
		}

		return simtypes.NewOperationMsg(deleteClaimRecordMsg, true, "", nil), nil, nil
	}
}

func createNClaimRecordEntries(r *rand.Rand, accs []simtypes.Account, n int) []*types.ClaimRecordEntry {
	var claimRecords []*types.ClaimRecordEntry
	for i := 0; i < n; i++ {
		claimRecordAccount, _ := simtypes.RandomAcc(r, accs)
		claimRecords = append(claimRecords, &types.ClaimRecordEntry{
			UserEntryAddress: claimRecordAccount.Address.String(),
			Amount:           sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(utils.RandInt64(r, 1000)))),
		})
	}
	return claimRecords
}
