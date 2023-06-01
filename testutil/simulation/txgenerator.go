package simulation

import (
	"cosmossdk.io/errors"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	helpers2 "github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
)

func SendMessageWithFees(ctx sdk.Context, r *rand.Rand, ak types.AccountKeeper, app *baseapp.BaseApp,
	simAccount simtypes.Account, msg sdk.Msg, spendable sdk.Coins, chainID string) error {
	account := ak.GetAccount(ctx, simAccount.Address)
	if !spendable.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}

	fees, err := simtypes.RandomFees(r, ctx, spendable)
	if err != nil {
		return err
	}

	txConfig := params.MakeTestEncodingConfig().TxConfig
	tx, err := helpers2.GenSignedMockTx(r, params.MakeTestEncodingConfig().TxConfig, []sdk.Msg{msg}, fees, helpers2.DefaultGenTxGas, chainID, []uint64{account.GetAccountNumber()}, []uint64{account.GetSequence()},
		simAccount.PrivKey,
	)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrConflict, "unable to generate mock tx")
	}

	_, _, err = app.SimDeliver(txConfig.TxEncoder(), tx)
	return err
}

func SendMessageWithRandomFees(ctx sdk.Context, r *rand.Rand, ak types.AccountKeeper, bk types.BankKeeper, app *baseapp.BaseApp,
	simAccount simtypes.Account, msg sdk.Msg, chainID string) error {
	account := ak.GetAccount(ctx, simAccount.Address)
	spendable := bk.SpendableCoins(ctx, simAccount.Address)

	fees, err := simtypes.RandomFees(r, ctx, spendable)
	if err != nil {
		return err
	}

	txConfig := params.MakeTestEncodingConfig().TxConfig
	tx, err := helpers2.GenSignedMockTx(r, params.MakeTestEncodingConfig().TxConfig, []sdk.Msg{msg}, fees, helpers2.DefaultGenTxGas, chainID, []uint64{account.GetAccountNumber()}, []uint64{account.GetSequence()},
		simAccount.PrivKey,
	)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrConflict, "unable to generate mock tx")
	}

	_, _, err = app.SimDeliver(txConfig.TxEncoder(), tx)
	return err
}
