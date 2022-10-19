package simulation

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateVestingPool(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, position := simtypes.RandomAcc(r, accs)
		fmt.Println("POSITION! " + strconv.Itoa(position))
		allAccVestings := k.GetAllAccountVestings(ctx)
		fmt.Println("ALL ACCOUNT VESTINGS", len(allAccVestings))
		msg := &types.MsgCreateVestingPool{
			Creator:     simAccount.Address.String(),
			Name:        "New Vesting Pool",
			Amount:      sdk.NewInt(100000),
			VestingType: "NEW_VESTING_TYEPE",
			Duration:    time.Duration(10000000),
		}
		err := sendMsgSend(r, app, bk, ak, msg, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "Vest simulation implemented", nil), nil, nil
	}
}

func SimulateMsgCreateVestingAccount(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateVestingAccount{
			FromAddress: simAccount.Address.String(),
		}

		// TODO: Handling the CreateVestingAccount simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateVestingAccount simulation not implemented"), nil, nil
	}
}

func sendMsgSend(
	r *rand.Rand, app *baseapp.BaseApp, bk types.BankKeeper, ak types.AccountKeeper,
	msg *types.MsgCreateVestingPool, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	var (
		fees sdk.Coins
		err  error
	)

	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}

	account := ak.GetAccount(ctx, from)
	spendable := bk.SpendableCoins(ctx, account.GetAddress())
	fmt.Println("Spendable:", spendable.String())
	coins := sdk.Coins{sdk.Coin{Amount: sdk.NewInt(100000), Denom: "stake"}}
	fees, err = simtypes.RandomFees(r, ctx, coins)
	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenSignedMockTx(
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

	_, _, err = app.Deliver(txGen.TxEncoder(), tx)
	if err != nil {
		return err
	}

	return nil
}
