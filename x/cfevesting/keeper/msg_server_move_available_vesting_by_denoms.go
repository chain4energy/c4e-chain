package keeper

import (
	"context"
	"cosmossdk.io/errors"
	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MoveAvailableVestingByDenoms(goCtx context.Context, msg *types.MsgMoveAvailableVestingByDenoms) (*types.MsgMoveAvailableVestingByDenomsResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "move available vesting by denoms message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	fromAccAddress, toAccAddress, err := types.ValidateMsgMoveAvailableVestingByDenom(msg.FromAddress, msg.ToAddress, msg.Denoms)
	if err != nil {
		k.Logger(ctx).Debug("move available vesting by denoms - validation error", "error", err)
		return nil, err
	}

	locked := k.bank.LockedCoins(ctx, fromAccAddress)
	amount := sdk.NewCoins()
	for _, denom := range msg.Denoms {
		if len(denom) == 0 {
			return nil, errors.Wrapf(types.ErrParam, "move available vesting by denoms - empty denom")
		}
		denAmount := locked.AmountOf(denom)
		if denAmount.IsPositive() {
			amount = amount.Add(sdk.NewCoin(denom, denAmount))
		}
	}
	if err := k.splitVestingCoins(ctx, fromAccAddress, toAccAddress, amount); err != nil {
		return nil, errors.Wrap(err, "move available vesting by denoms")
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
