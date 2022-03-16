package keeper

import (
	"context"
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

)

func (k msgServer) WithdrawDelegatorReward(goCtx context.Context, msg *types.MsgWithdrawDelegatorReward) (*types.MsgWithdrawDelegatorRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	keeper.Logger(ctx).Info("WithdrawDelegatorReward: " + msg.DelegatorAddress)

	accVestings, found := keeper.GetAccountVestings(ctx, msg.DelegatorAddress)
	if !found {
		return nil, fmt.Errorf("no vestings for account: %q", msg.DelegatorAddress)
	}
	if len(accVestings.DelegableAddress) == 0 {
		return nil, fmt.Errorf("no delegable vestings for account: %q", msg.DelegatorAddress)
	}

	withdrawMsg := distrtypes.MsgWithdrawDelegatorReward{DelegatorAddress: accVestings.DelegableAddress,
		ValidatorAddress: msg.ValidatorAddress}
	_, err := k.distrMsgServer.WithdrawDelegatorReward(goCtx, &withdrawMsg)
	if err != nil {
		return nil, err
	}

	// TODO check if this should be in telemetry if vesting module
	// defer func() {
	// 	for _, a := range amount {
	// 		if a.Amount.IsInt64() {
	// 			telemetry.SetGaugeWithLabels(
	// 				[]string{"tx", "msg", "withdraw_reward"},
	// 				float32(a.Amount.Int64()),
	// 				[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
	// 			)
	// 		}
	// 	}
	// }()

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawRewards,
			sdk.NewAttribute(types.AttributeKeyDelegator, msg.DelegatorAddress),
			sdk.NewAttribute(types.AttributeKeyDelegableAddress, accVestings.DelegableAddress),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress),
		),
	})

	return &types.MsgWithdrawDelegatorRewardResponse{}, nil
}
