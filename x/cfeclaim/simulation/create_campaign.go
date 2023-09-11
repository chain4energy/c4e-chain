package simulation

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/v2/testutil/simulation"
	"github.com/chain4energy/c4e-chain/v2/testutil/utils"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/v2/x/cfevesting/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateCampaign(
	k keeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	cfevestingKeeper cfevestingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, err := createCampaign(ak, bk, cfevestingKeeper, app, r, ctx, accs, chainID)
		if err != nil {
			return simtypes.NewOperationMsg(&types.MsgCreateCampaign{}, false, "", nil), nil, nil
		}

		return simtypes.NewOperationMsg(&types.MsgCreateCampaign{}, true, "", nil), nil, nil
	}
}

var secondsInADay = 24 * 60 * 60

func createCampaign(ak types.AccountKeeper, bk types.BankKeeper, cfevestingKeeper cfevestingkeeper.Keeper,
	app *baseapp.BaseApp, r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, chainID string) (*simtypes.Account, error) {
	simAccount, _ := simtypes.RandomAcc(r, accs)
	spendable := bk.SpendableCoins(ctx, simAccount.Address)
	startTime := ctx.BlockTime()
	endTime := startTime.Add(utils.RandDurationBetween(r, 10, 100))

	lockupPeriod := utils.RandDurationBetween(r, 1, secondsInADay)
	vestingPeriod := utils.RandDurationBetween(r, 1, secondsInADay)
	randomMathInt, _ := simtypes.RandPositiveInt(r, math.NewInt(100000))
	randomDec := simtypes.RandomDecAmount(r, sdk.MustNewDecFromStr("0.8"))

	msgCreateCampaign := &types.MsgCreateCampaign{
		Owner:                  simAccount.Address.String(),
		Name:                   simtypes.RandStringOfLength(r, 10),
		Description:            simtypes.RandStringOfLength(r, 10),
		CampaignType:           types.CampaignType(utils.RandIntBetween(r, 1, 3)),
		FeegrantAmount:         &randomMathInt,
		RemovableClaimRecords:  utils.RandomBool(r),
		InitialClaimFreeAmount: &randomMathInt,
		Free:                   &randomDec,
		StartTime:              &startTime,
		EndTime:                &endTime,
		LockupPeriod:           &lockupPeriod,
		VestingPeriod:          &vestingPeriod,
		VestingPoolName:        "",
	}

	if msgCreateCampaign.CampaignType == types.VestingPoolCampaign {
		vestingPoolAmount, err := simtypes.RandPositiveInt(r, spendable.AmountOf(sdk.DefaultBondDenom))
		if err != nil {
			return nil, err
		}
		randomVestingPoolName := simtypes.RandStringOfLength(r, 10)
		randVesingTypeId := utils.RandInt64(r, 3)
		randomVestingType := "New vesting" + strconv.Itoa(int(randVesingTypeId))
		if err = cfevestingKeeper.CreateVestingPool(ctx, simAccount.Address.String(), randomVestingPoolName, vestingPoolAmount, time.Hour, randomVestingType); err != nil {
			return &simAccount, err
		}
		msgCreateCampaign.VestingPoolName = randomVestingPoolName
	}

	if err := simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgCreateCampaign, chainID); err != nil {
		return &simAccount, err
	}
	return &simAccount, nil
}
