package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
)

func SimulateMsgMoveAvailableVestingByDenoms(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		if !spendable.IsAllPositive() {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateVestingAccount, "balance is negative"), nil, nil
		}

		amount, err := simtypes.RandPositiveInt(r, spendable.AmountOf(sdk.DefaultBondDenom))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateVestingAccount, "balance is negative"), nil, nil
		}

		createVestingAccCoin := sdk.NewCoin(sdk.DefaultBondDenom, amount)
		createVestingAccountCoins := sdk.NewCoins(createVestingAccCoin)
		startTime := ctx.BlockTime().Add(-utils.RandDurationBetween(r, 1, 10))
		endTime := ctx.BlockTime().Add(utils.RandDurationBetween(r, 1, 10))
		simAccount2 := simtypes.RandomAccounts(r, 1)[0]

		msgCreateVestingAccount := &types.MsgCreateVestingAccount{
			FromAddress: simAccount.Address.String(),
			ToAddress:   simAccount2.Address.String(),
			StartTime:   startTime.Unix(),
			EndTime:     endTime.Unix(),
			Amount:      createVestingAccountCoins,
		}
		err = simulation.SendMessageWithFees(ctx, r, ak, app, simAccount, msgCreateVestingAccount, spendable.Sub(sdk.NewCoin(sdk.DefaultBondDenom, amount)), chainID)
		if err != nil {
			return simtypes.NewOperationMsg(msgCreateVestingAccount, false, "", nil), nil, nil
		}

		simAccount3 := simtypes.RandomAccounts(r, 1)[0]
		msgMoveAvailableVestingByDenoms := &types.MsgMoveAvailableVestingByDenoms{
			FromAddress: simAccount2.Address.String(),
			ToAddress:   simAccount3.Address.String(),
			Denoms: []string{
				GetRandomDenom(utils.RandInt64(r, 2)),
			},
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak, bk, app, simAccount2, msgMoveAvailableVestingByDenoms, chainID); err != nil {
			return simtypes.NewOperationMsg(msgMoveAvailableVestingByDenoms, false, "", nil), nil, nil
		}
		return simtypes.NewOperationMsg(msgMoveAvailableVestingByDenoms, true, "", nil), nil, nil
	}
}

func GetRandomDenom(n int64) string {
	if n == 0 {
		return sdk.DefaultBondDenom
	}
	return "abcd"
}
