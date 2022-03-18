package keeper

import (
	"context"

	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendVesting(goCtx context.Context, msg *types.MsgSendVesting) (*types.MsgSendVestingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	withdrawn, err := keeper.SendVesting(ctx, msg.FromAddress, msg.ToAddress, msg.VestingId, msg.Amount, msg.RestartVesting)
	if err != nil {
		return nil, err
	}

	denom := keeper.Denom(ctx)
	if msg.Amount.IsInt64() {
		defer func() {
			telemetry.IncrCounter(1, types.ModuleName, "trasfer_vesting")
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(msg.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", denom)},
			)
		}()
	}

	// if msg.Amount.IsInt64() {
	// 	defer func() {
	// 		telemetry.SetGaugeWithLabels(
	// 			[]string{"tx", "msg", types.ModuleName, "send"},
	// 			float32(msg.Amount.Int64()),
	// 			[]metrics.Label{telemetry.NewLabel("denom", denom)},
	// 		)
	// 	}()
	// }

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransfer,
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.ToAddress),
			sdk.NewAttribute(types.AttributeKeySender, msg.FromAddress),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()+denom),
			sdk.NewAttribute(types.AttributeKeyWithdrawn, withdrawn.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.FromAddress),
		),
	})

	return &types.MsgSendVestingResponse{}, nil
}
