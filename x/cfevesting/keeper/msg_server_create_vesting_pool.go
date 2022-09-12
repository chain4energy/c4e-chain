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
			telemetry.IncrCounter(1, types.ModuleName, "vesting_pools")
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(msg.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", denom)},
			)
		}()
	}

	ctx.EventManager().EmitTypedEvent(&types.NewVestingPool{
		Creator:     msg.Creator,
		Name:        msg.Name,
		Amount:      msg.Amount.String() + denom,
		Duration:    msg.Duration.String(),
		VestingType: msg.VestingType,
	})

	return &types.MsgCreateVestingPoolResponse{}, nil
}
