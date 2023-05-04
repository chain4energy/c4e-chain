package simulation

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgClaim(
	k keeper.Keeper,
	cfevestingKeeper cfevestingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		startTime := ctx.BlockTime()
		endTime := startTime.Add(time.Duration(helpers.RandIntBetween(r, 1000000, 10000000)))

		lockupPeriod := time.Duration(helpers.RandIntBetween(r, 1000000, 10000000))
		vestingPeriod := time.Duration(helpers.RandIntBetween(r, 1000000, 10000000))
		msg := &types.MsgCreateCampaign{
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
		if msg.CampaignType == types.VestingPoolCampaign {
			randomVestingPoolName := helpers.RandStringOfLengthCustomSeed(r, 10)
			randVesingTypeId := helpers.RandomInt(r, 3)
			randomVestingType := "New vesting" + strconv.Itoa(int(randVesingTypeId))
			_ = cfevestingKeeper.CreateVestingPool(ctx, simAccount.Address.String(), randomVestingPoolName, math.NewInt(1000000), time.Hour, randomVestingType)
			msg.VestingPoolName = randomVestingPoolName
		}
		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CreateCampaign(msgServerCtx, msg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), ""), nil, nil
		}

		campaigns := k.GetCampaigns(ctx)
		campaignId := uint64(len(campaigns) - 1)
		randomWeight := helpers.RandomDecAmount(r, sdk.NewDec(1))
		addMissionToCampaignMsg := &types.MsgAddMissionToCampaign{
			Owner:          simAccount.Address.String(),
			CampaignId:     campaignId,
			Name:           helpers.RandStringOfLengthCustomSeed(r, 10),
			Description:    helpers.RandStringOfLengthCustomSeed(r, 10),
			MissionType:    types.MissionType(helpers.RandIntBetween(r, 2, 6)),
			Weight:         &randomWeight,
			ClaimStartDate: nil,
		}

		_, err = msgServer.AddMissionToCampaign(msgServerCtx, addMissionToCampaignMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Add mission to campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, addMissionToCampaignMsg.Type(), ""), nil, nil
		}

		startCampaignMsg := &types.MsgStartCampaign{
			Owner:      simAccount.Address.String(),
			CampaignId: campaignId,
		}

		_, err = msgServer.StartCampaign(msgServerCtx, startCampaignMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Start campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, startCampaignMsg.Type(), ""), nil, nil
		}

		addClaimRecordsMsg := &types.MsgAddClaimRecords{
			Owner:        simAccount.Address.String(),
			CampaignId:   campaignId,
			ClaimRecords: createNClaimRecords(100, accs),
		}

		_, err = msgServer.AddClaimRecords(msgServerCtx, addClaimRecordsMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Add claim records campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, addClaimRecordsMsg.Type(), ""), nil, nil
		}

		initialClaimMsg := &types.MsgInitialClaim{
			Claimer:        addClaimRecordsMsg.ClaimRecords[helpers.RandomInt(r, len(addClaimRecordsMsg.ClaimRecords))].Address,
			CampaignId:     campaignId,
			AddressToClaim: "",
		}

		_, err = msgServer.InitialClaim(msgServerCtx, initialClaimMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Initial claim error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, initialClaimMsg.Type(), ""), nil, nil
		}

		claimMsg := &types.MsgClaim{
			Claimer:    initialClaimMsg.Claimer,
			CampaignId: campaignId,
			MissionId:  1,
		}

		_, err = msgServer.Claim(msgServerCtx, claimMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Claim mission error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, claimMsg.Type(), ""), nil, nil
		}
		k.Logger(ctx).Debug("SIMULATION: Claim mission - claimed")
		return simtypes.NewOperationMsg(claimMsg, true, "Claim simulation completed", nil), nil, nil
	}
}
