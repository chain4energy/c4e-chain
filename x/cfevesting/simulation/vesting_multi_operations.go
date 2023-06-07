package simulation

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/testutil/simulation"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateVestingMultiOperations(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		randVestingDuration := time.Duration(utils.RandInt64(r, 3))

		multiOperationsCount := utils.RandInt64(r, 10)
		var poolsNames []string
		for i := int64(0); i < multiOperationsCount; i++ {
			randVesingTypeId := utils.RandInt64(r, 3)
			randVestingPoolName := simtypes.RandStringOfLength(r, 10)
			poolsNames = append(poolsNames, randVestingPoolName)
			randomVestingType := "New vesting" + strconv.Itoa(int(randVesingTypeId))

			spendable := bk.SpendableCoins(ctx, simAccount.Address)
			if !spendable.IsAllPositive() {
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateVestingAccount, "balance is negative"), nil, nil
			}
			amount, err := simtypes.RandPositiveInt(r, spendable.AmountOf(sdk.DefaultBondDenom))
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateVestingPool, "unable to generate positive amount"), nil, err
			}
			msgCreateVestingPool := &types.MsgCreateVestingPool{
				Owner:       simAccount.Address.String(),
				Name:        randVestingPoolName,
				Amount:      amount,
				VestingType: randomVestingType,
				Duration:    randVestingDuration,
			}
			if err = simulation.SendMessageWithFees(ctx, r, ak.(authkeeper.AccountKeeper), app, simAccount, msgCreateVestingPool, spendable.Sub(sdk.NewCoin(sdk.DefaultBondDenom, amount)), chainID); err != nil {
				return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting multi operations - create vesting pool", "", false, nil), nil, nil
			}
		}

		for i := int64(0); i < multiOperationsCount; i++ {
			randMsgSendToVestinAccAmount := simtypes.RandomAmount(r, math.NewInt(100))

			simAccount2 := simtypes.RandomAccounts(r, 1)[0]
			msgSendToVestingAccount := &types.MsgSendToVestingAccount{
				Owner:           simAccount.Address.String(),
				ToAddress:       simAccount2.Address.String(),
				VestingPoolName: poolsNames[i],
				Amount:          randMsgSendToVestinAccAmount,
				RestartVesting:  true,
			}
			if err := simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgSendToVestingAccount, chainID); err != nil {
				return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting multi operations - send to vesting account", "", false, nil), nil, nil
			}
		}

		msgWithdrawAllAvailable := &types.MsgWithdrawAllAvailable{
			Owner: simAccount.Address.String(),
		}
		if err := simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgWithdrawAllAvailable, chainID); err != nil {
			return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting multi operations - withdraw all available", "", false, nil), nil, nil
		}

		return simtypes.NewOperationMsgBasic(types.ModuleName, "Vesting multi operations simulation completed", "", true, nil), nil, nil
	}
}
