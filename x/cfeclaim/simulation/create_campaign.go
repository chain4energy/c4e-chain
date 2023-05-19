package simulation

import (
	"context"
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

func SimulateMsgCreateCampaign(
	k keeper.Keeper,
	cfevestingKeeper cfevestingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := createCampaign(k, cfevestingKeeper, msgServer, msgServerCtx, r, ctx, accs)
		if err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Create campaign", "", false, nil), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Create campaign - CREATED")
		return simtypes.NewOperationMsgBasic(types.ModuleName, "Create campaign completed", "", true, nil), nil, nil
	}
}

var nanoSecondsInDay = 1000000000 * 24 * 60 * 60

func createCampaign(k keeper.Keeper, cfevestingKeeper cfevestingkeeper.Keeper, msgServer types.MsgServer, msgServerCtx context.Context, r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) (sdk.AccAddress, error) {
	simAccount, _ := simtypes.RandomAcc(r, accs)
	startTime := ctx.BlockTime().Add(time.Duration(helpers.RandIntBetween(r, 1000000, 10000000)))
	endTime := startTime.Add(time.Duration(helpers.RandIntBetween(r, 1000000, 10000000)))

	lockupPeriod := time.Duration(helpers.RandomInt(r, nanoSecondsInDay))
	vestingPeriod := time.Duration(helpers.RandomInt(r, nanoSecondsInDay))

	randomMathInt := helpers.RandomAmount(r, math.NewInt(1000000))
	msg := &types.MsgCreateCampaign{
		Owner:                  simAccount.Address.String(),
		Name:                   helpers.RandStringOfLengthCustomSeed(r, 10),
		Description:            helpers.RandStringOfLengthCustomSeed(r, 10),
		CampaignType:           types.CampaignType(helpers.RandomInt(r, 3)),
		FeegrantAmount:         nil,
		RemovableClaimRecords:  helpers.RandomBool(r),
		InitialClaimFreeAmount: &randomMathInt,
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

	_, err := msgServer.CreateCampaign(msgServerCtx, msg)
	if err != nil {
		k.Logger(ctx).Error("SIMULATION: Create campaign error", err.Error())
		return simAccount.Address, err
	}
	return simAccount.Address, nil
}
