package keeper

import (
	"context"
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Delegate(goCtx context.Context, msg *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper
	// stakingKeeper := keeper.staking

	accVestings, found := keeper.GetAccountVestings(ctx, msg.DelegatorAddress)
	if !found {
		return nil, fmt.Errorf("No vestings for account: %q", msg.DelegatorAddress)
	}
	if len(accVestings.DelegableAddress) == 0 {
		return nil, fmt.Errorf("No delegable vestings for account: %q", msg.DelegatorAddress)
	}

	delegableAddress, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
	if err != nil {
		return nil, err
	}

	accountAddress, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	// TODO !!!! getting all reward coins, not only bond denom
	// balanceBefore := k.bank.GetBalance(ctx, delegableAddress, keeper.staking.BondDenom(ctx))

	delagateMsg := stakingtypes.MsgDelegate{DelegatorAddress: accVestings.DelegableAddress,
		ValidatorAddress: msg.ValidatorAddress, Amount: msg.Amount}
	_, err = k.stakingMsgServer.Delegate(goCtx, &delagateMsg)
	if err != nil {
		return nil, err
	}

	// balanceAfter := k.bank.GetBalance(ctx, delegableAddress, keeper.staking.BondDenom(ctx))
	// balanceAfter = balanceAfter.Add(msg.Amount)
	// coinToSend := balanceAfter.Sub(balanceBefore)

	// accVestings.Delegated += msg.Amount.Amount.Uint64()
	// keeper.SetAccountVestings(ctx, accVestings)

	// k.bank.SendCoins(ctx, delegableAddress, accountAddress, sdk.NewCoins(coinToSend))

	if k.distribution.GetDelegatorWithdrawAddr(ctx, delegableAddress).Equals(delegableAddress) {
		k.distribution.SetDelegatorWithdrawAddr(ctx, delegableAddress, accountAddress)
	}
	// valAddr, valErr := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	// if valErr != nil {
	// 	return nil, valErr
	// }

	// validator, found := stakingKeeper.GetValidator(ctx, valAddr)
	// if !found {
	// 	return nil, stakingtypes.ErrNoValidatorFound
	// }

	// delegatorAddress, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
	// if err != nil {
	// 	return nil, err
	// }

	// bondDenom := stakingKeeper.BondDenom(ctx)
	// if msg.Amount.Denom != bondDenom {
	// 	return nil, sdkerrors.Wrapf(
	// 		sdkerrors.ErrInvalidRequest, "invalid coin denomination: got %s, expected %s", msg.Amount.Denom, bondDenom,
	// 	)
	// }

	// // NOTE: source funds are always unbonded
	// newShares, err := stakingKeeper.Delegate(ctx, delegatorAddress, msg.Amount.Amount, stakingtypes.Unbonded, validator, true)
	// if err != nil {
	// 	return nil, err
	// }
	// // TODO: Handling the message
	// _ = newShares

	return &types.MsgDelegateResponse{}, nil
}
