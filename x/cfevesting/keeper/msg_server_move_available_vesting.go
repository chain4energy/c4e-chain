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

func (k msgServer) MoveAvailableVesting(goCtx context.Context, msg *types.MsgMoveAvailableVesting) (*types.MsgMoveAvailableVestingResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "move available vesting message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrParam, fmt.Errorf("move available vesting - error parsing from address: %s: %w", msg.FromAddress, err).Error())
	}
	amount := k.bank.LockedCoins(ctx, from)

	if err := k.splitVestingCoins(ctx, from, msg.ToAddress, amount); err != nil {
		return nil, sdkerrors.Wrap(err, "move available vesting")
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
	return &types.MsgMoveAvailableVestingResponse{}, nil

}
