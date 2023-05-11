package keeper

import (
	"context"

	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VestingsSummary(goCtx context.Context, req *types.QueryVestingsSummaryRequest) (*types.QueryVestingsSummaryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	summary, err := k.createVestingsSummary(ctx, false)
	result := types.QueryVestingsSummaryResponse(*summary)
	return &result, err
}

func (k Keeper) createVestingsSummary(ctx sdk.Context, genesisOnly bool) (*Summary, error) {
	mAcc := k.account.GetModuleAccount(ctx, types.ModuleName)
	denom := k.GetParams(ctx).Denom
	var vestingInPoolsAmount math.Int
	if genesisOnly {
		allAvpPolls := k.GetAllAccountVestingPools(ctx)
		vestingInPoolsAmount = allAvpPolls.GetGenesisAmount()
	} else {
		modBalance := k.bank.GetBalance(ctx, mAcc.GetAddress(), denom)
		vestingInPoolsAmount = modBalance.Amount
	}

	allAcc := k.GetAllVestingAccountTrace(ctx)
	allVestingInAccounts := sdk.ZeroInt()
	allLockedNotDelegated := sdk.ZeroInt()
	for _, accFromList := range allAcc {
		if genesisOnly && !accFromList.IsGenesisOrFromGenesis() {
			continue
		}
		accAddr, err := sdk.AccAddressFromBech32(accFromList.Address)
		if err != nil {
			return &Summary{}, status.Error(codes.Internal, err.Error())
		}

		vestingAccount := k.account.GetAccount(ctx, accAddr)
		if continuousVestingAccount, ok := vestingAccount.(*vestingtypes.ContinuousVestingAccount); ok {
			vestingCoins := continuousVestingAccount.GetVestingCoins(ctx.BlockTime())
			lockedCoins := continuousVestingAccount.BaseVestingAccount.LockedCoinsFromVesting(vestingCoins) // TODO ??? this should be reduced by amount of burned coins by jailing if burned amount before moment in time when thore burned vested
			allVestingInAccounts = allVestingInAccounts.Add(vestingCoins.AmountOf(denom))
			allLockedNotDelegated = allLockedNotDelegated.Add(lockedCoins.AmountOf(denom))
		} else if periodcVestingAccount, ok := vestingAccount.(*types.PeriodicContinuousVestingAccount); ok {
			vestingCoins := periodcVestingAccount.GetVestingCoinsForSpecyficPeriods(ctx.BlockTime(), accFromList.PeriodsToTrace)
			lockedCoins := periodcVestingAccount.GetLockedCoins(vestingCoins)
			allVestingInAccounts = allVestingInAccounts.Add(vestingCoins.AmountOf(denom))
			allLockedNotDelegated = allLockedNotDelegated.Add(lockedCoins.AmountOf(denom))
		}
	}

	return &Summary{
		VestingAllAmount:        allVestingInAccounts.Add(vestingInPoolsAmount),
		VestingInPoolsAmount:    vestingInPoolsAmount,
		VestingInAccountsAmount: allVestingInAccounts,
		DelegatedVestingAmount:  allVestingInAccounts.Sub(allLockedNotDelegated),
	}, nil
}

type Summary struct {
	VestingAllAmount        sdk.Int
	VestingInPoolsAmount    sdk.Int
	VestingInAccountsAmount sdk.Int
	DelegatedVestingAmount  sdk.Int
}
