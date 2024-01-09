package keeper

import (
	"context"
	"cosmossdk.io/errors"
	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateVestingAccount(goCtx context.Context, msg *types.MsgCreateVestingAccount) (*types.MsgCreateVestingAccountResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create vesting account message")
	if msg.Amount == nil || msg.Amount.IsAnyNil() {
		return nil, errors.Wrap(types.ErrParam, "create vesting account - amount is nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	defer func() {
		for _, a := range msg.Amount {
			if a.Amount.IsInt64() {
				telemetry.SetGaugeWithLabels(
					[]string{"tx", "msg", types.ModuleName, msg.Type()},
					float32(a.Amount.Int64()),
					[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
				)
			}
		}
	}()
	keeper := k.Keeper
	err := keeper.CreateVestingAccount(ctx, msg.FromAddress, msg.ToAddress, msg.Amount, msg.StartTime, msg.EndTime)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateVestingAccountResponse{}, nil
}
