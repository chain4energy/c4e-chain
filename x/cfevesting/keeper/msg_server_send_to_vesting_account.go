package keeper

import (
	"context"
	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SendToVestingAccount(goCtx context.Context, msg *types.MsgSendToVestingAccount) (*types.MsgSendToVestingAccountResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "send to vesting account message")
	if msg.Amount.IsNil() {
		return nil, sdkerrors.Wrap(types.ErrParam, "send to new vesting account - amount is nil")
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
	_, err := keeper.SendToNewVestingAccount(ctx, msg.FromAddress, msg.ToAddress, msg.VestingPoolName, msg.Amount, msg.RestartVesting)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendToVestingAccountResponse{}, nil
}
