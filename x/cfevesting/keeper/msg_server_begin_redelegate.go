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

func (k msgServer) BeginRedelegate(goCtx context.Context, msg *types.MsgBeginRedelegate) (*types.MsgBeginRedelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	accVestings, found := keeper.GetAccountVestings(ctx, msg.DelegatorAddress)
	if !found {
		return nil, fmt.Errorf("no vestings for account: %q", msg.DelegatorAddress)
	}
	if len(accVestings.DelegableAddress) == 0 {
		return nil, fmt.Errorf("no delegable vestings for account: %q", msg.DelegatorAddress)
	}

	delagateMsg := stakingtypes.MsgBeginRedelegate{DelegatorAddress: accVestings.DelegableAddress,
		ValidatorSrcAddress: msg.ValidatorSrcAddress, ValidatorDstAddress: msg.ValidatorDstAddress, Amount: msg.Amount}
	resp, err := k.stakingMsgServer.BeginRedelegate(goCtx, &delagateMsg)
	if err != nil {
		return nil, err
	}

	if msg.Amount.Amount.IsInt64() {
		defer func() {
			telemetry.IncrCounter(1, types.ModuleName, "redelegate")
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(msg.Amount.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", msg.Amount.Denom)},
			)
		}()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRedelegate,
			sdk.NewAttribute(types.AttributeKeySrcValidator, msg.ValidatorSrcAddress),
			sdk.NewAttribute(types.AttributeKeyDstValidator, msg.ValidatorDstAddress),
			sdk.NewAttribute(types.AttributeKeyDelegator, msg.DelegatorAddress),
			sdk.NewAttribute(types.AttributeKeyDelegableAddress, accVestings.DelegableAddress),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyCompletionTime, resp.CompletionTime.Format(time.RFC3339)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress),
		),
	})

	return &types.MsgBeginRedelegateResponse{CompletionTime: resp.CompletionTime}, nil
}
