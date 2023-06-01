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

func SimulateSendToVestingAccount(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		allVestingPools := k.GetAllAccountVestingPools(ctx)
		if len(allVestingPools) == 0 {
			return simtypes.NewOperationMsg(&types.MsgSendToVestingAccount{}, false, "", nil), nil, nil
		}
		randVestingPoolId := utils.RandInt64(r, len(allVestingPools))

		accAddress := allVestingPools[randVestingPoolId].Owner
		simAccount, _ := simtypes.FindAccount(accs, sdk.MustAccAddressFromBech32(accAddress))
		numOfPools := len(allVestingPools[randVestingPoolId].VestingPools)
		var randVestingId int64 = 0
		if numOfPools > 1 {
			randVestingId = utils.RandInt64(r, numOfPools-1)
		}

		spendable := bk.SpendableCoins(ctx, simAccount.Address)
		if !spendable.IsAllPositive() {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSendToVestingAccount, "balance is negative"), nil, nil
		}
		amount, err := simtypes.RandPositiveInt(r, spendable.AmountOf(sdk.DefaultBondDenom))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSendToVestingAccount, "unable to generate positive amount"), nil, err
		}

		simAccount2 := simtypes.RandomAccounts(r, 1)[0]
		msg := types.NewMsgSendToVestingAccount(accAddress, simAccount2.Address.String(),
			allVestingPools[randVestingPoolId].VestingPools[randVestingId].Name, amount, false)

		if err = simulation.SendMessageWithFees(ctx, r, ak, app, simAccount, msg, spendable.Sub(sdk.NewCoin(sdk.DefaultBondDenom, amount)), chainID); err != nil {
			return simtypes.NewOperationMsg(msg, false, "", nil), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
