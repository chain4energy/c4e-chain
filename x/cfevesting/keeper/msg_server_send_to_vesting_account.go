package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	metrics "github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
)

func (k msgServer) SendToVestingAccount(goCtx context.Context, msg *types.MsgSendToVestingAccount) (*types.MsgSendToVestingAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	defer func() {
		telemetry.IncrCounter(1, "new", "account")

		if msg.Amount.IsInt64() {
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(msg.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", k.Keeper.Denom(ctx))},
			)
		}
		
	}()

	keeper := k.Keeper
	_, err := keeper.SendToNewVestingAccount(ctx, msg.FromAddress, msg.ToAddress, msg.VestingId, msg.Amount, msg.RestartVesting)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgSendToVestingAccountResponse{}, nil
}
