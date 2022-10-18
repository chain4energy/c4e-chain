package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) Vestings(goCtx context.Context, req *types.QueryVestingsRequest) (*types.QueryVestingsResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	mAcc := k.account.GetModuleAccount(ctx, types.ModuleName)
	denom := k.GetParams(ctx).Denom
	modBalance := k.bank.GetBalance(ctx, mAcc.GetAddress(), denom)
	vestingInPoolsAmount := modBalance.Amount

	allAcc := k.GetAllVestingAccount(ctx)
	allVestingInAccounts := sdk.ZeroInt()
	allLockedNotDelegated := sdk.ZeroInt()
	for _, accFromList := range allAcc {
		accAddr, err := sdk.AccAddressFromBech32(accFromList.Address)
		if err != nil {
			return &types.QueryVestingsResponse{}, sdkerrors.Wrap(types.ErrParsing, err.Error())
		}

		vestingAccount := k.account.GetAccount(ctx, accAddr)
		if continuousVestingAccount, ok := vestingAccount.(*vestingtypes.ContinuousVestingAccount); ok {
			lockedCoins := continuousVestingAccount.LockedCoins(ctx.BlockTime())
			vestingCoins := continuousVestingAccount.GetVestingCoins(ctx.BlockTime())
			allVestingInAccounts = allVestingInAccounts.Add(vestingCoins.AmountOf(denom))
			allLockedNotDelegated = allLockedNotDelegated.Add(lockedCoins.AmountOf(denom))
		}
	}

	return &types.QueryVestingsResponse{
		VestingAllAmount:        allVestingInAccounts.Add(vestingInPoolsAmount),
		VestingInPoolsAmount:    vestingInPoolsAmount,
		VestingInAccountsAmount: allVestingInAccounts,
		DelegatedVestingAmount:  allVestingInAccounts.Sub(allLockedNotDelegated),
	}, nil
}
