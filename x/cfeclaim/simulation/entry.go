package simulation

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"math/rand"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddClaimRecords(
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
		msg := &types.MsgCreateCampaign{
			Owner:                  simAccount.Address.String(),
			Name:                   helpers.RandStringOfLengthCustomSeed(r, 10),
			Description:            helpers.RandStringOfLengthCustomSeed(r, 10),
			CampaignType:           types.CampaignType(helpers.RandomInt(r, 3)),
			FeegrantAmount:         nil,
			InitialClaimFreeAmount: nil,
			StartTime:              &startTime,
			EndTime:                &endTime,
			LockupPeriod:           &lockupPeriod,
			VestingPeriod:          &vestingPeriod,
			VestingPoolName:        "",
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CreateCampaign(msgServerCtx, msg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), ""), nil, nil
		}

		campaigns := k.GetCampaigns(ctx)
		startCampaignMsg := &types.MsgStartCampaign{
			Owner:      simAccount.Address.String(),
			CampaignId: uint64(len(campaigns) - 1),
		}

		_, err = msgServer.StartCampaign(msgServerCtx, startCampaignMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Start campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, startCampaignMsg.Type(), ""), nil, nil
		}

		addClaimRecordsMsg := &types.MsgAddClaimRecords{
			Owner:        simAccount.Address.String(),
			CampaignId:   uint64(len(campaigns) - 1),
			ClaimRecords: createNUserEntries(100, accs),
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

func createNUserEntries(n int, accs []simtypes.Account) []*types.ClaimRecord {
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
