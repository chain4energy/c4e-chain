package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Vestings(goCtx context.Context, req *types.QueryVestingsRequest) (*types.QueryVestingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
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
			return &types.QueryVestingsResponse{}, err
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
