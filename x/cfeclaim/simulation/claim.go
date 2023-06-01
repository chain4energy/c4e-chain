package simulation

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/chain4energy/c4e-chain/testutil/utils"
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
		endTime := startTime.Add(time.Duration(utils.RandIntBetween(r, 1000000, 10000000)))

		lockupPeriod := time.Duration(utils.RandInt64(r, nanoSecondsInDay))
		vestingPeriod := time.Duration(utils.RandInt64(r, nanoSecondsInDay))
		randomMathInt := utils.RandomAmount(r, math.NewInt(1000000))
		msgCreateCampaign := &types.MsgCreateCampaign{
			Owner:                  simAccount.Address.String(),
			Name:                   simtypes.RandStringOfLength(r, 10),
			Description:            simtypes.RandStringOfLength(r, 10),
			CampaignType:           types.CampaignType(utils.RandInt64(r, 3)),
			FeegrantAmount:         &randomMathInt,
			InitialClaimFreeAmount: &randomMathInt,
			StartTime:              &startTime,
			EndTime:                &endTime,
			LockupPeriod:           &lockupPeriod,
			VestingPeriod:          &vestingPeriod,
			VestingPoolName:        "",
		}
		if msgCreateCampaign.CampaignType == types.VestingPoolCampaign {
			randomVestingPoolName := simtypes.RandStringOfLength(r, 10)
			randVesingTypeId := utils.RandInt64(r, 3)
			randomVestingType := "New vesting" + strconv.Itoa(int(randVesingTypeId))
			_ = cfevestingKeeper.CreateVestingPool(ctx, simAccount.Address.String(), randomVestingPoolName, math.NewInt(10000000), time.Hour, randomVestingType)
			msgCreateCampaign.VestingPoolName = randomVestingPoolName
		}
		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CreateCampaign(msgServerCtx, msgCreateCampaign)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Create campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, msgCreateCampaign.Type(), ""), nil, nil
		}

		campaigns := k.GetAllCampaigns(ctx)
		campaignId := uint64(len(campaigns) - 1)
		randomWeight := utils.RandomDecAmount(r, sdk.NewDec(1))
		AddMissionMsg := &types.MsgAddMission{
			Owner:          simAccount.Address.String(),
			CampaignId:     campaignId,
			Name:           simtypes.RandStringOfLength(r, 10),
			Description:    simtypes.RandStringOfLength(r, 10),
			MissionType:    types.MissionType(utils.RandIntBetween(r, 2, 5)),
			Weight:         &randomWeight,
			ClaimStartDate: nil,
		}

		_, err = msgServer.AddMission(msgServerCtx, AddMissionMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Add mission to campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, AddMissionMsg.Type(), ""), nil, nil
		}

		EnableCampaignMsg := &types.MsgEnableCampaign{
			Owner:      simAccount.Address.String(),
			CampaignId: campaignId,
		}

		_, err = msgServer.EnableCampaign(msgServerCtx, EnableCampaignMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Start campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, EnableCampaignMsg.Type(), ""), nil, nil
		}

		addClaimRecordsMsg := &types.MsgAddClaimRecords{
			Owner:              simAccount.Address.String(),
			CampaignId:         campaignId,
			ClaimRecordEntries: createNClaimRecordEntries(100, accs),
		}

		_, err = msgServer.AddClaimRecords(msgServerCtx, addClaimRecordsMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Add claim records campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, addClaimRecordsMsg.Type(), ""), nil, nil
		}
		claimerAddress := addClaimRecordsMsg.ClaimRecordEntries[utils.RandInt64(r, len(addClaimRecordsMsg.ClaimRecordEntries))].UserEntryAddress
		initialClaimMsg := &types.MsgInitialClaim{
			Claimer:            claimerAddress,
			CampaignId:         campaignId,
			DestinationAddress: "",
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
		return simtypes.NewOperationMsg(claimMsg, true, fmt.Sprintf("Claim simulation completed for %s", msgCreateCampaign.CampaignType.String()), nil), nil, nil
	}
}
