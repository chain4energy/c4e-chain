package keeper

import (
	"context"
	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateVestingPool(goCtx context.Context, msg *types.MsgCreateVestingPool) (*types.MsgCreateVestingPoolResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create vesting pool message")
	if msg.Amount.IsNil() {
		return nil, sdkerrors.Wrap(types.ErrParam, "add vesting pool - amount is nil")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	err := keeper.CreateVestingPool(ctx, msg.Owner, msg.Name, msg.Amount, msg.Duration, msg.VestingType)
	if err != nil {
		return nil, err
	}

	denom := keeper.Denom(ctx)
	if msg.Amount.IsInt64() {
		defer func() {
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(msg.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", denom)},
			)
		}()
	}

	event := &types.NewVestingPool{
		Owner:       msg.Owner,
		Name:        msg.Name,
		Amount:      msg.Amount.String() + denom,
		Duration:    msg.Duration.String(),
		VestingType: msg.VestingType,
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("new vesting pool emit event error", "event", event, "error", err.Error())
	}

	return &types.MsgCreateVestingPoolResponse{}, nil
}
