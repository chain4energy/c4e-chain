package simulation

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/baseapp"
	helpers2 "github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"math/rand"
)

func SendMessageWithFees(ctx sdk.Context, r *rand.Rand, ak authkeeper.AccountKeeper, app *baseapp.BaseApp,
	simAccount simtypes.Account, msg sdk.Msg, spendable sdk.Coins, chainID string) error {
	_, _, err := sendMessage(ctx, r, ak, app, simAccount, msg, spendable, chainID)
	return err
}

func SendMessageWithRandomFees(ctx sdk.Context, r *rand.Rand, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, app *baseapp.BaseApp,
	simAccount simtypes.Account, msg sdk.Msg, chainID string) error {
	spendable := bk.SpendableCoins(ctx, simAccount.Address)
	_, _, err := sendMessage(ctx, r, ak, app, simAccount, msg, spendable, chainID)
	return err
}

func SendMessageWithRandomFeesAndResult(ctx sdk.Context, r *rand.Rand, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, app *baseapp.BaseApp,
	simAccount simtypes.Account, msg sdk.Msg, chainID string) (*sdk.Result, error) {
	spendable := bk.SpendableCoins(ctx, simAccount.Address)
	_, result, err := sendMessage(ctx, r, ak, app, simAccount, msg, spendable, chainID)
	return result, err
}

func sendMessage(ctx sdk.Context, r *rand.Rand, ak authkeeper.AccountKeeper, app *baseapp.BaseApp,
	simAccount simtypes.Account, msg sdk.Msg, spendable sdk.Coins, chainID string) (sdk.GasInfo, *sdk.Result, error) {
	account := ak.GetAccount(ctx, simAccount.Address)
	if !spendable.IsAllPositive() {
		return sdk.GasInfo{}, nil, sdkerrors.ErrInsufficientFunds
	}

	fees, err := simtypes.RandomFees(r, ctx, spendable)
	if err != nil {
		return sdk.GasInfo{}, nil, err
	}

	txConfig := params.MakeTestEncodingConfig().TxConfig
	tx, err := helpers2.GenSignedMockTx(r, params.MakeTestEncodingConfig().TxConfig, []sdk.Msg{msg}, fees, helpers2.DefaultGenTxGas, chainID, []uint64{account.GetAccountNumber()}, []uint64{account.GetSequence()},
		simAccount.PrivKey,
	)
	if err != nil {
		return sdk.GasInfo{}, nil, errors.Wrap(sdkerrors.ErrConflict, "unable to generate mock tx")
	}

	return app.SimDeliver(txConfig.TxEncoder(), tx)
}
