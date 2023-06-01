package simulation

import (
	"cosmossdk.io/math"
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

func SimulateMsgCloseCampaign(
	k keeper.Keeper,
	cfevestingKeeper cfevestingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		startTime := ctx.BlockTime().Add(-time.Hour)
		endTime := startTime.Add(time.Second)

		lockupPeriod := time.Hour * 10
		vestingPeriod := time.Hour * 10
		randomMathInt := utils.RandomAmount(r, math.NewInt(1000000))
		campaign := types.Campaign{
			Owner:                  simAccount.Address.String(),
			Name:                   simtypes.RandStringOfLength(r, 10),
			Description:            simtypes.RandStringOfLength(r, 10),
			CampaignType:           types.CampaignType(utils.RandIntBetween(r, 1, 3)),
			RemovableClaimRecords:  utils.RandomBool(r),
			FeegrantAmount:         randomMathInt,
			InitialClaimFreeAmount: randomMathInt,
			StartTime:              startTime,
			EndTime:                endTime,
			LockupPeriod:           lockupPeriod,
			VestingPeriod:          vestingPeriod,
			VestingPoolName:        "",
		}

		if campaign.CampaignType == types.VestingPoolCampaign {
			randomVestingPoolName := simtypes.RandStringOfLength(r, 10)
			randVesingTypeId := utils.RandInt64(r, 3)
			randomVestingType := "New vesting" + strconv.Itoa(int(randVesingTypeId))
			_ = cfevestingKeeper.CreateVestingPool(ctx, simAccount.Address.String(), randomVestingPoolName, math.NewInt(10000), time.Hour, randomVestingType)
			campaign.VestingPoolName = randomVestingPoolName
		}

		k.AppendNewCampaign(ctx, campaign)

		campaigns := k.GetAllCampaigns(ctx)
		campaignId := uint64(len(campaigns) - 1)

		closeCampaignMsg := &types.MsgCloseCampaign{
			Owner:      simAccount.Address.String(),
			CampaignId: campaignId,
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.CloseCampaign(msgServerCtx, closeCampaignMsg)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Close campaign error", err.Error())
			return simtypes.NoOpMsg(types.ModuleName, closeCampaignMsg.Type(), ""), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Close campaign - closed")
		return simtypes.NewOperationMsg(closeCampaignMsg, true, "Close campaign simulation completed", nil), nil, nil
	}
}
