package simulation

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/testutil/simulation"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
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
	ak types.AccountKeeper,
	bk types.BankKeeper,
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
			Enabled:                true,
			Free:                   sdk.ZeroDec(),
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
			if err := cfevestingKeeper.CreateVestingPool(ctx, simAccount.Address.String(), randomVestingPoolName, math.NewInt(10000), time.Hour, randomVestingType); err != nil {
				return simtypes.NewOperationMsg(&types.MsgCloseCampaign{}, false, "", nil), nil, nil
			}

			campaign.VestingPoolName = randomVestingPoolName
		}
		k.AppendNewCampaign(ctx, campaign)

		campaigns := k.GetAllCampaigns(ctx)
		campaignId := uint64(len(campaigns) - 1)

		closeCampaignMsg := &types.MsgCloseCampaign{
			Owner:      simAccount.Address.String(),
			CampaignId: campaignId,
		}

		if err := simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, closeCampaignMsg, chainID); err != nil {
			return simtypes.NewOperationMsg(closeCampaignMsg, false, "", nil), nil, nil
		}

		return simtypes.NewOperationMsg(closeCampaignMsg, true, "", nil), nil, nil
	}
}
