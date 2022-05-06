package keeper

import (
	"context"

	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateVestingPool(goCtx context.Context, msg *types.MsgCreateVestingPool) (*types.MsgCreateVestingPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	err := keeper.CreateVestingPool(ctx, msg.Creator, msg.Name, msg.Amount, msg.Duration, msg.VestingType)
	if err != nil {
		return nil, err
	}

	denom := keeper.Denom(ctx)
	if msg.Amount.IsInt64() {
		defer func() {
			telemetry.IncrCounter(1, types.ModuleName, "vest")
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(msg.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", denom)},
			)
		}()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeVest,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()+denom),
			sdk.NewAttribute(types.AttributeKeyVestingType, msg.VestingType),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgCreateVestingPoolResponse{}, nil
}
