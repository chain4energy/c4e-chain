package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MoveAvailableVestingByDenoms(goCtx context.Context, msg *types.MsgMoveAvailableVestingByDenoms) (*types.MsgMoveAvailableVestingByDenomsResponse, error) {
	// TODO events and telemetry
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "move available vesting by denoms - error parsing from address: %s", msg.FromAddress)
	}
	locked := k.bank.LockedCoins(ctx, from)
	amount := sdk.NewCoins()
	for _, denom := range msg.Denoms {
		if len(denom) == 0 {
			return nil, sdkerrors.Wrapf(types.ErrParam, "move available vesting by denoms - empty denom") // TODO error type
		}
		denAmount := locked.AmountOf(denom)
		if denAmount.IsPositive() {
			amount = amount.Add(sdk.NewCoin(denom, denAmount))
		}
	}
	if err := k.splitVestingCoins(ctx, from, msg.ToAddress, amount); err != nil {
		return nil, sdkerrors.Wrap(err, "move available vesting by denoms")
	}
	return &types.MsgMoveAvailableVestingByDenomsResponse{}, nil

}
