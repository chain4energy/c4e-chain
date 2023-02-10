package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) UnlockUnbondedContinuousVestingAccountCoins(ctx sdk.Context, ownerAddress sdk.AccAddress, amountToUnlock sdk.Coins) error {
	k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins", "ownerAddress", ownerAddress, "amountToUnlock", amountToUnlock)
	ownerAccount := k.account.GetAccount(ctx, ownerAddress)
	if ownerAccount == nil {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - account doesn't exist", "ownerAddress", ownerAddress)
		return sdkerrors.Wrapf(types.ErrNotExists, "account %s doesn't exist", ownerAddress) // TODO ErrNotExists to c4eerrors namespace and remove from this module
	}

	vestingAcc, ok := ownerAccount.(*vestingtypes.ContinuousVestingAccount)
	if !ok {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - account is not ContinuousVestingAccount", "account", ownerAccount)
		return sdkerrors.Wrapf(types.ErrNotExists, "account %s is not ContinuousVestingAccount", ownerAddress) // TODO some other error
	}

	lockedCoins := vestingAcc.LockedCoins(ctx.BlockTime())

	if amountToUnlock.IsAnyGT(lockedCoins) {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - not enough to unlock", "account", ownerAccount, "lockedCoins", lockedCoins, "amountToUnlock", amountToUnlock)
		return sdkerrors.Wrapf(types.ErrNotExists, "account %s: not enough to unlock. locked: %s, to unlock: %s", ownerAddress, lockedCoins, amountToUnlock) // TODO some other error
	}

	vestingCoins := vestingAcc.GetVestingCoins(ctx.BlockTime())
	orignalVestings := vestingAcc.OriginalVesting

	for _, coin := range amountToUnlock {
		orignalVesting := orignalVestings.AmountOf(coin.Denom)
		vestingCoin := vestingCoins.AmountOf(coin.Denom)
		originalVestingDiffDec := coin.Amount.ToDec().Mul(orignalVesting.ToDec()).Quo(vestingCoin.ToDec())
		originalVestingDiff := originalVestingDiffDec.TruncateInt()
		vestingAcc.OriginalVesting = vestingAcc.OriginalVesting.Sub(sdk.NewCoins(sdk.NewCoin(coin.Denom, originalVestingDiff)))
		if !vestingCoin.Equal(coin.Amount.Add(vestingAcc.GetVestingCoins(ctx.BlockTime()).AmountOf(coin.Denom))) {
			vestingAcc.OriginalVesting = vestingAcc.OriginalVesting.Sub(sdk.NewCoins(sdk.NewCoin(coin.Denom, sdk.NewInt(1)))) // TODO is this enough ??
		}

	}
	k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins", "ownerAddress", ownerAddress,
		"amountToUnlock", amountToUnlock, "vestingCoins", vestingCoins, "orignalVestings", orignalVestings,
		"newOrignalVestings", vestingAcc.OriginalVesting, "newVestingCoins", vestingAcc.GetVestingCoins(ctx.BlockTime()))

	k.account.SetAccount(ctx, vestingAcc)
	return nil
}
