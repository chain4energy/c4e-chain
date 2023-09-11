package keeper

import (
	"context"
	"cosmossdk.io/errors"
	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MoveAvailableVesting(goCtx context.Context, msg *types.MsgMoveAvailableVesting) (*types.MsgMoveAvailableVestingResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "move available vesting message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	fromAccAddress, toAccAddress, err := types.ValidateMsgMoveAvailableVesting(msg.FromAddress, msg.ToAddress)
	if err != nil {
		k.Logger(ctx).Debug("move available vesting - validation error", "error", err)
		return nil, err
	}
	amount := k.bank.LockedCoins(ctx, fromAccAddress)

	if err := k.splitVestingCoins(ctx, fromAccAddress, toAccAddress, amount); err != nil {
		return nil, errors.Wrap(err, "move available vesting")
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
