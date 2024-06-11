package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SplitVesting(goCtx context.Context, msg *types.MsgSplitVesting) (*types.MsgSplitVestingResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "split vesting message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	fromAccAddress, toAccAddress, err := types.ValidateMsgSplitVesting(msg.FromAddress, msg.ToAddress, msg.Amount)
	if err != nil {
		k.Logger(ctx).Debug("split vesting - validation error", "error", err)
		return nil, err
	}

	if err = k.splitVestingCoins(ctx, fromAccAddress, toAccAddress, msg.Amount); err != nil {
		return nil, errors.Wrap(err, "split vesting")
	}

	for _, a := range msg.Amount {
		if a.Amount.IsInt64() {
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(a.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
			)
		}
	}
	return &types.MsgSplitVestingResponse{}, nil
}

func (k msgServer) splitVestingCoins(ctx sdk.Context, from sdk.AccAddress, toAddress sdk.AccAddress,
	amount sdk.Coins) error {

	if len(amount) == 0 {
		return errors.Wrapf(types.ErrParam, "split vesting coins - no coins to split %s", amount)
	}

	if amount.IsAnyNil() {
		return errors.Wrapf(types.ErrParam, "split vesting coins - all coins of amount must not be nil: %s", amount)
	}

	if err := k.bank.IsSendEnabledCoins(ctx, amount...); err != nil {
		return errors.Wrapf(types.ErrParam, "send is disabled")
	}

	if k.bank.BlockedAddr(toAddress) {
		return errors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", toAddress)
	}

	if acc := k.account.GetAccount(ctx, toAddress); acc != nil {
		k.Logger(ctx).Debug("split vesting coins - to account already exists error", "toAddress", toAddress)
		return errors.Wrapf(types.ErrAlreadyExists, "split vesting coins - account address: %s", toAddress)
	}

	vestingAcc, err := k.UnlockUnbondedContinuousVestingAccountCoins(ctx, from, amount)
	if err != nil {
		return errors.Wrap(err, "split vesting coins")
	}

	startTime := ctx.BlockTime().Unix()
	if vestingAcc.StartTime > startTime {
		startTime = vestingAcc.StartTime
	}

	event := &types.EventVestingSplit{
		Source:      from.String(),
		Destination: toAddress.String(),
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("vesting split emit event error", "event", event, "error", err.Error())
	}
	if _, err = k.newContinuousVestingAccount(ctx, toAddress, amount, startTime, vestingAcc.EndTime); err != nil {
		return errors.Wrap(err, "split vesting coins")
	}

	if err = k.bank.SendCoins(ctx, from, toAddress, amount); err != nil {
		return errors.Wrap(err, "split vesting coins")
	}

	vAcc, found := k.GetVestingAccountTrace(ctx, from.String())
	if found {
		k.AppendVestingAccountTrace(ctx, types.VestingAccountTrace{
			Address:            toAddress.String(),
			Genesis:            false,
			FromGenesisPool:    vAcc.FromGenesisPool,
			FromGenesisAccount: vAcc.Genesis || vAcc.FromGenesisAccount,
		})
	}
	return nil
}
