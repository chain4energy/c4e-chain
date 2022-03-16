package keeper

import (
	"context"
	"fmt"

	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (k msgServer) Delegate(goCtx context.Context, msg *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper
	// stakingKeeper := keeper.staking

	accVestings, found := keeper.GetAccountVestings(ctx, msg.DelegatorAddress)
	if !found {
		return nil, fmt.Errorf("no vestings for account: %q", msg.DelegatorAddress)
	}
	if len(accVestings.DelegableAddress) == 0 {
		return nil, fmt.Errorf("no delegable vestings for account: %q", msg.DelegatorAddress)
	}

	delegableAddress, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
	if err != nil {
		return nil, err
	}

	accountAddress, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	delagateMsg := stakingtypes.MsgDelegate{DelegatorAddress: accVestings.DelegableAddress,
		ValidatorAddress: msg.ValidatorAddress, Amount: msg.Amount}
	_, err = k.stakingMsgServer.Delegate(goCtx, &delagateMsg)
	if err != nil {
		return nil, err
	}

	if k.distribution.GetDelegatorWithdrawAddr(ctx, delegableAddress).Equals(delegableAddress) {
		k.distribution.SetDelegatorWithdrawAddr(ctx, delegableAddress, accountAddress)
	}

	if msg.Amount.Amount.IsInt64() {
		defer func() {
			telemetry.IncrCounter(1, types.ModuleName, "vesting-delegate")
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(msg.Amount.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", msg.Amount.Denom)},
			)
		}()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDelegate,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyDelegator, msg.DelegatorAddress),
			sdk.NewAttribute(types.AttributeKeyDelegableAddress, accVestings.DelegableAddress),
			// sdk.NewAttribute(types.AttributeKeyNewShares, newShares.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress),
		),
	})

	return &types.MsgDelegateResponse{}, nil
}
