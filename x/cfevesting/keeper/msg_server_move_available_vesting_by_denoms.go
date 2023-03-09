package keeper

import (
	"context"
	"fmt"

	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MoveAvailableVestingByDenoms(goCtx context.Context, msg *types.MsgMoveAvailableVestingByDenoms) (*types.MsgMoveAvailableVestingByDenomsResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "move available vesting by denoms message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrParam, fmt.Errorf("move available vesting by denoms - error parsing from address: %s: %w", msg.FromAddress, err).Error())
	}
	locked := k.bank.LockedCoins(ctx, from)
	amount := sdk.NewCoins()
	for _, denom := range msg.Denoms {
		if len(denom) == 0 {
			return nil, sdkerrors.Wrapf(types.ErrParam, "move available vesting by denoms - empty denom")
		}
		denAmount := locked.AmountOf(denom)
		if denAmount.IsPositive() {
			amount = amount.Add(sdk.NewCoin(denom, denAmount))
		}
	}
	if err := k.splitVestingCoins(ctx, from, msg.ToAddress, amount); err != nil {
		return nil, sdkerrors.Wrap(err, "move available vesting by denoms")
	}
	for _, a := range amount {
		if a.Amount.IsInt64() {
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(a.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
			)
		}
	}
	return &types.MsgMoveAvailableVestingByDenomsResponse{}, nil

}
