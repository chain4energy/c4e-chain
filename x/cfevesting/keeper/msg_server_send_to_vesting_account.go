package keeper

import (
	"context"
	"cosmossdk.io/errors"
	metrics "github.com/armon/go-metrics"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendToVestingAccount(goCtx context.Context, msg *types.MsgSendToVestingAccount) (*types.MsgSendToVestingAccountResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "send to vesting account message")
	if msg.Amount.IsNil() {
		return nil, errors.Wrap(c4eerrors.ErrParam, "send to new vesting account - amount is nil")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	defer func() {
		if msg.Amount.IsInt64() {
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(msg.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", k.Keeper.Denom(ctx))},
			)
		}
	}()

	keeper := k.Keeper
	if err := keeper.SendToNewVestingAccount(ctx, msg.Owner, msg.ToAddress, msg.VestingPoolName, msg.Amount, msg.RestartVesting); err != nil {
		return nil, err
	}

	return &types.MsgSendToVestingAccountResponse{}, nil
}
