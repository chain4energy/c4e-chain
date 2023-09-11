package keeper

import (
	"context"

	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawAllAvailable(goCtx context.Context, msg *types.MsgWithdrawAllAvailable) (*types.MsgWithdrawAllAvailableResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "withdraw all available message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper
	withdrawn, err := keeper.WithdrawAllAvailable(ctx, msg.Owner)
	if err != nil {
		return nil, err
	}

	if withdrawn.Amount.IsInt64() {
		defer func() {
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(withdrawn.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", withdrawn.Denom)},
			)
		}()
	}

	return &types.MsgWithdrawAllAvailableResponse{Withdrawn: withdrawn}, nil
}
