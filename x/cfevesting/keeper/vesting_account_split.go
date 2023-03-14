package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) UnlockUnbondedContinuousVestingAccountCoins(ctx sdk.Context, ownerAddress sdk.AccAddress, amountToUnlock sdk.Coins) (*vestingtypes.ContinuousVestingAccount, error) {
	k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins", "ownerAddress", ownerAddress, "amountToUnlock", amountToUnlock)
	if err := amountToUnlock.Validate(); err != nil {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - amountToUnlock validation error", "error", err)
		return nil, sdkerrors.Wrap(err, "amount to unlock validation error")

	}

	ownerAccount := k.account.GetAccount(ctx, ownerAddress)
	if ownerAccount == nil {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - account doesn't exist", "ownerAddress", ownerAddress)
		return nil, sdkerrors.Wrapf(types.ErrNotExists, "account %s doesn't exist", ownerAddress) // TODO ErrNotExists to c4eerrors namespace and remove from this module
	}

	vestingAcc, ok := ownerAccount.(*vestingtypes.ContinuousVestingAccount)
	if !ok {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - account is not ContinuousVestingAccount", "account", ownerAccount)
		return nil, sdkerrors.Wrapf(types.ErrNotExists, "account %s is not ContinuousVestingAccount", ownerAddress) // TODO some other error
	}

	lockedCoins := vestingAcc.LockedCoins(ctx.BlockTime())

	if !amountToUnlock.IsAllLTE(lockedCoins) {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - not enough to unlock", "account", ownerAccount, "lockedCoins", lockedCoins, "amountToUnlock", amountToUnlock)
		return nil, sdkerrors.Wrapf(types.ErrNotExists, "account %s: not enough to unlock. locked: %s, to unlock: %s", ownerAddress, lockedCoins, amountToUnlock) // TODO some other error
	}

	vestingCoins := vestingAcc.GetVestingCoins(ctx.BlockTime())
	orignalVestings := vestingAcc.OriginalVesting

	for _, coin := range amountToUnlock {
		if coin.Amount.GT(sdk.ZeroInt()) {
			orignalVesting := orignalVestings.AmountOf(coin.Denom)
			vestingCoin := vestingCoins.AmountOf(coin.Denom)
			originalVestingDiffDec := sdk.NewDecFromInt(coin.Amount).Mul(sdk.NewDecFromInt(orignalVesting)).Quo(sdk.NewDecFromInt(vestingCoin))
			originalVestingDiff := originalVestingDiffDec.TruncateInt()
			vestingAcc.OriginalVesting = vestingAcc.OriginalVesting.Sub(sdk.NewCoin(coin.Denom, originalVestingDiff))
			if vestingCoin.Sub(vestingAcc.GetVestingCoins(ctx.BlockTime()).AmountOf(coin.Denom)).LT(coin.Amount) {
				vestingAcc.OriginalVesting = vestingAcc.OriginalVesting.Sub(sdk.NewCoin(coin.Denom, sdk.NewInt(1)))
				// This is to compensate bankers rounding of vesting coins calculation nad truncating of original vesting difference.
				// VgC - vesting coins, VdC - vested coins, OV - original vesting, bt - block time
				// VdCdec - decimal vested coins amount
				// VgC(bt) = OV - VdC(bt) where VdC(bt) = bankersRound(VdCdec(bt))

				// example: OV = 8999999999999999999, half of vesting time passed and we want to unlock 1.
				// VdC(half vesting time) = VgC(half vesting time) = 8999999999999999999 / 2 = 4499999999999999999.5
				// OVnew = OV - (toUnlock * OV / VgC)
				// OVnew = 8999999999999999999 - 1 * 8999999999999999999 / (8999999999999999999 - 4499999999999999999.5) )
				//         8999999999999999999 - 2 = 8999999999999999997
				// New VgC(half vesting time) = 4499999999999999998.5
				// unlocked = VgC(half vesting time) - New VgC(half vesting time) =
				//            4499999999999999999.5 - 4499999999999999998.5 = 1, 1 is unlocked

				// but with bankers rounding
				// VdC(half vesting time) = bankers-round(4499999999999999999.5) = 4500000000000000000
				// VgC(half vesting time) = 8999999999999999999 - 4500000000000000000 = 4499999999999999999
				// OVnew = 8999999999999999999 - trunc(1 * 8999999999999999999 / 4499999999999999999)
				// OVnew = 8999999999999999999 - trunc(2) = 8999999999999999997
				//                   [2 is in fact 2.00000000000000000022, but decimal precision is 18, so in fact it is truncating and Ceil is not enough]
				// New VgC(half vesting time) = 8999999999999999997 - bankers-round(8999999999999999997 / 2) =
				//                              8999999999999999997 - bankers-round(4499999999999999998.5) =
				//                              8999999999999999997 - 4499999999999999998 = 4499999999999999999
				// unlocked = VgC(half vesting time) - New VgC(half vesting time) =
				//            4499999999999999999 - 4499999999999999999 = 0, so nothing was unlocked, 1 missing

				// second example, we want to unlock 2249999999999999999 in teh same situation
				// VdC(half vesting time) = VgC(half vesting time) = 8999999999999999999 / 2 = 4499999999999999999.5
				// OVnew = 8999999999999999999 - 2249999999999999999 * 8999999999999999999 / (8999999999999999999 - 4499999999999999999.5) =
				//         8999999999999999999 - 2249999999999999999 * 2 =
				//         4500000000000000001
				// New VgC(half vesting time) = 2250000000000000000.5
				// unlocked = VgC(half vesting time) - New VgC(half vesting time) =
				//            4499999999999999999.5 - 2250000000000000000.5 = 2249999999999999999

				// but with bankers rounding:
				// VdC(half vesting time) = bankers-round(4499999999999999999.5) = 4500000000000000000
				// VgC(half vesting time) = 8999999999999999999 - 4500000000000000000 = 4499999999999999999

				// OVnew = 8999999999999999999 - trunc(2249999999999999999 * 8999999999999999999 / 4499999999999999999) =
				//         8999999999999999999 - trunc(4499999999999999998.5) =
				//         4500000000000000001
				// New VgC(half vesting time) = 4500000000000000001 - bankers-round(4500000000000000001 / 2) =
				//                              4500000000000000001 - bankers-round(2250000000000000000.5) =
				//                              4500000000000000001 - 2250000000000000000 = 2250000000000000001
				// unlocked = VgC(half vesting time) - New VgC(half vesting time) =
				//            4499999999999999999 - 2250000000000000001 = 2249999999999999998, 1 is missing

				// OV = 1999
				// VdC(half vesting time) = 999.5
				// VgC(half vesting time) = 999.5
				// unlocking = 499
				// OVnew = 1999 - 499 * 1999 / 999.5) =
				//         1999 - 998 =
				//         1001
				// New VgC(half vesting time) = 1001 - 1001 / 2 =
				//                              500.5
				// unlocked = VgC(half vesting time) - New VgC(half vesting time) =
				//            999.5 - 500.5 = 499

				// OV = 1999
				// VdC(half vesting time) = bankers-round(999.5) = 1000
				// VgC(half vesting time) = 1999 - 1000 = 999
				// unlocking = 499
				// OVnew = 1999 - trunc(499 * 1999 / 999) =
				//         1999 - trunc(998.499499499) =
				//         1001
				// New VgC(half vesting time) = 1001 - bankers-round(1001 / 2) =
				//                              1001 - bankers-round(500.5) =
				//                              501
				// unlocked = VgC(half vesting time) - New VgC(half vesting time) =
				//            999 - 501 = 498, 1 is missing

				// OV = 1997
				// VdC(half vesting time) = 998.5
				// VgC(half vesting time) = 998.5
				// unlocking = 499
				// OVnew = 1997 - 499 * 1997 / 998.5) =
				//         1997 - 998 =
				//         999
				// New VgC(half vesting time) = 999 - 999 / 2 =
				//                              499.5
				// unlocked = VgC(half vesting time) - New VgC(half vesting time) =
				//            998.5 - 499.5 = 499

				// OV = 1997
				// VdC(half vesting time) = bankers-round(998.5) = 998
				// VgC(half vesting time) = 1997 - 998 = 999
				// unlocking = 499
				// OVnew = 1997 - trunc(499 * 1997 / 999) =
				//         1997 - trunc(997.500500501) =
				//         1000
				// New VgC(half vesting time) = 1000 - bankers-round(1000 / 2) =
				//                              1000 - bankers-round(500) =
				//                              500
				// unlocked = VgC(half vesting time) - New VgC(half vesting time) =
				//            999 - 500 = 499

				// example: OV = 8999999999999999999, 1 hour after half of vesting time (vesting time = 1000h) passed and we want to unlock 1.
				// VdC(half vesting time + 1h) = 8999999999999999999 * 501/1000 = 4508999999999999999.499
				// VgC(half vesting time + 1h) = 4490999999999999999.501
				// unlocking = 1
				// OVnew = 8999999999999999999 - 1 * 8999999999999999999 / 4490999999999999999.501
				//         8999999999999999999 - 2.004008016032064128 = 8999999999999999996.995991983967935872
				// New VdC(half vesting time + 1h) = 8999999999999999996.995991983967935872 * 501/1000 = 4508999999999999998.494991983967935872
				// New VgC(half vesting time + 1h) = 8999999999999999996.995991983967935872 - 4508999999999999998.494991983967935872 =
				//                                   4490999999999999998.501
				// unlocked = VgC(half vesting time + 1h) - New VgC(half vesting time) =
				//            4490999999999999999.501 - 4490999999999999998.501 = 1, 1 is unlocked

				// bankers rounding:
				// VdC(half vesting time + 1h) = bankers-round(8999999999999999999 * 501/1000) = bankers-round(4508999999999999999.499)
				//                               4508999999999999999
				// VgC(half vesting time + 1h) = 8999999999999999999 - 4508999999999999999 - 4491000000000000000
				// unlocking = 1
				// OVnew = 8999999999999999999 - trunc(1 * 8999999999999999999 / 4491000000000000000)
				//         8999999999999999999 - trunc(2.004008016032064128) = 8999999999999999997
				// New VdC(half vesting time + 1h) = bankers-round(8999999999999999997 * 501/1000) = bankers-round(4508999999999999998.497) =
				//                                   4508999999999999998
				// New VgC(half vesting time + 1h) = 8999999999999999997 - 4508999999999999998 =
				//                              4490999999999999999
				// unlocked = VgC(half vesting time + 1h) - New VgC(half vesting time) =
				//            4491000000000000000 - 4490999999999999999 = 1, 1 is unlocked

				// bankers rounding if we use Ceil instead of truncate: just to prove that ceiling is not good solution
				// VdC(half vesting time + 1h) = bankers-round(8999999999999999999 * 501/1000) = bankers-round(4508999999999999999.499)
				//                               4508999999999999999
				// VgC(half vesting time + 1h) = 8999999999999999999 - 4508999999999999999 - 4491000000000000000
				// unlocking = 1
				// OVnew = 8999999999999999999 - ceil(1 * 8999999999999999999 / 4491000000000000000)
				//         8999999999999999999 - ceil(2.004008016032064128) = 8999999999999999996
				// New VdC(half vesting time + 1h) = bankers-round(8999999999999999996 * 501/1000) = bankers-round(4508999999999999997.996) =
				//                                   4508999999999999998
				// New VgC(half vesting time + 1h) = 8999999999999999996 - 4508999999999999998 =
				//                              4490999999999999998
				// unlocked = VgC(half vesting time + 1h) - New VgC(half vesting time) =
				//            4491000000000000000 - 4490999999999999998 = 2, 1 too much is unlocked
			}
		}
	}
	k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins", "ownerAddress", ownerAddress,
		"amountToUnlock", amountToUnlock, "vestingCoins", vestingCoins, "orignalVestings", orignalVestings,
		"newOrignalVestings", vestingAcc.OriginalVesting, "newVestingCoins", vestingAcc.GetVestingCoins(ctx.BlockTime()))

	k.account.SetAccount(ctx, vestingAcc)
	return vestingAcc, nil
}
