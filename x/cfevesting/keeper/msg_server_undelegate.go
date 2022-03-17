package keeper

import (
	"context"
	"fmt"
	"time"

	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (k msgServer) Undelegate(goCtx context.Context, msg *types.MsgUndelegate) (*types.MsgUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	accVestings, found := keeper.GetAccountVestings(ctx, msg.DelegatorAddress)
	if !found {
		return nil, fmt.Errorf("no vestings for account: %q", msg.DelegatorAddress)
	}
	if len(accVestings.DelegableAddress) == 0 {
		return nil, fmt.Errorf("no delegable vestings for account: %q", msg.DelegatorAddress)
	}

	delagateMsg := stakingtypes.MsgUndelegate{DelegatorAddress: accVestings.DelegableAddress,
		ValidatorAddress: msg.ValidatorAddress, Amount: msg.Amount}
	resp, err := k.stakingMsgServer.Undelegate(goCtx, &delagateMsg)
	if err != nil {
		return nil, err
	}

	if msg.Amount.Amount.IsInt64() {
		defer func() {
			telemetry.IncrCounter(1, types.ModuleName, "undelegate")
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(msg.Amount.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", msg.Amount.Denom)},
			)
		}()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnbond,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyDelegator, msg.DelegatorAddress),
			sdk.NewAttribute(types.AttributeKeyDelegableAddress, accVestings.DelegableAddress),
			sdk.NewAttribute(types.AttributeKeyCompletionTime, resp.CompletionTime.Format(time.RFC3339)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress),
		),
	})

	return &types.MsgUndelegateResponse{CompletionTime: resp.CompletionTime}, nil
}
