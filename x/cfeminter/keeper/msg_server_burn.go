package keeper

import (
	"context"
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Burn message")
	if msg.Amount == nil || msg.Amount.IsAnyNil() {
		return nil, errors.Wrap(c4eerrors.ErrParam, "amount is nil")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.Burn(ctx, msg.Address, msg.Amount); err != nil {
		k.Logger(ctx).Debug("Burn", "err", err.Error())
		return nil, err
	}
	return &types.MsgBurnResponse{}, nil
}

func (k Keeper) Burn(ctx sdk.Context, address string, amount sdk.Coins) error {
	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return errors.Wrap(c4eerrors.ErrParsing, err.Error())
	}

	balances := k.bankKeeper.GetAllBalances(ctx, accAddress)
	if !amount.IsAllLTE(balances) {
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "balance is too small (%s < %s)", balances, amount)
	}

	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddress, types.ModuleName, amount); err != nil {
		return err
	}
	if err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, amount); err != nil {
		return err
	}
	return nil
}
