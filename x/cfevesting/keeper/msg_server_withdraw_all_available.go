package keeper

import (
	"context"

	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawAllAvailable(goCtx context.Context, msg *types.MsgWithdrawAllAvailable) (*types.MsgWithdrawAllAvailableResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper
	withdrawn, err := keeper.WithdrawAllAvailable(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	if withdrawn.Amount.IsInt64() {
		defer func() {
			telemetry.IncrCounter(1, types.ModuleName, "withdraw_available")
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(withdrawn.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", withdrawn.Denom)},
			)
		}()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawVesting,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAmount, withdrawn.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgWithdrawAllAvailableResponse{}, nil
}
