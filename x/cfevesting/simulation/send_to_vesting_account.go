package simulation

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simhelpers "github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
)

func SimulateSendToVestingAccount(
	_ types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		allVestingPools := k.GetAllAccountVestingPools(ctx)
		if len(allVestingPools) == 0 {
			return simtypes.NewOperationMsg(&types.MsgSendToVestingAccount{}, false, "", nil), nil, nil
		}
		randVestingPoolId := helpers.RandomInt(r, len(allVestingPools))
		accAddress := allVestingPools[randVestingPoolId].Owner
		randMsgSendToVestinAccAmount := math.NewInt(helpers.RandomInt(r, 10))
		claimRecordAccount, _ := simtypes.RandomAcc(r, accs)
		numOfPools := len(allVestingPools[randVestingPoolId].VestingPools)
		var randVestingId int64 = 0
		if numOfPools > 1 {
			randVestingId = helpers.RandomInt(r, numOfPools-1)
		}
		msgSendToVestingAccount := &types.MsgSendToVestingAccount{
			Owner:           accAddress,
			ToAddress:       claimRecordAccount.Address.String(),
			VestingPoolName: allVestingPools[randVestingPoolId].VestingPools[randVestingId].Name,
			Amount:          randMsgSendToVestinAccAmount,
			RestartVesting:  false,
		}

		msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
		_, err := msgServer.SendToVestingAccount(msgServerCtx, msgSendToVestingAccount)
		if err != nil {
			k.Logger(ctx).Error("SIMULATION: Send to vesting account error", err.Error())
			return simtypes.NewOperationMsg(msgSendToVestingAccount, false, "", nil), nil, nil
		}

		k.Logger(ctx).Debug("SIMULATION: Send to vesting account - FINISHED")
		return simtypes.NewOperationMsg(msgSendToVestingAccount, true, "", nil), nil, nil
	}
}

// sendMsgSend sends a transaction with a MsgSend from a provided random account.
func sendMsgSend(
	r *rand.Rand, app *baseapp.BaseApp, k keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper,
	msg *types.MsgSendToVestingAccount, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	var (
		fees sdk.Coins
		err  error
	)

	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return err
	}

	account := ak.GetAccount(ctx, from)
	spendable := bk.SpendableCoins(ctx, account.GetAddress())

	fees, err = simtypes.RandomFees(r, ctx, spendable)
	if err != nil {
		return err
	}

	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := simhelpers.GenSignedMockTx(
		r,
		txGen,
		[]sdk.Msg{msg},
		fees,
		helpers.DefaultGenTxGas,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privkeys...,
	)
	if err != nil {
		return err
	}

	_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
	if err != nil {
		return err
	}

	return nil
}
