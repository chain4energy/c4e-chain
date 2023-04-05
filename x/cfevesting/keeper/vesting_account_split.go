package keeper

import (
	errortypes "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

const bankersRoundingCompensation = 1

func (k Keeper) UnlockUnbondedContinuousVestingAccountCoins(ctx sdk.Context, ownerAddress sdk.AccAddress, amountToUnlock sdk.Coins) (*vestingtypes.ContinuousVestingAccount, error) {
	k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins", "ownerAddress", ownerAddress, "amountToUnlock", amountToUnlock)
	if err := amountToUnlock.Validate(); err != nil {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - amountToUnlock validation error", "error", err)
		return nil, sdkerrors.Wrap(err, "amount to unlock validation error")

	}

	ownerAccount := k.account.GetAccount(ctx, ownerAddress)
	if ownerAccount == nil {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - account doesn't exist", "ownerAddress", ownerAddress)
		return nil, sdkerrors.Wrapf(errortypes.ErrNotExists, "account %s doesn't exist", ownerAddress)
	}

	vestingAcc, ok := ownerAccount.(*vestingtypes.ContinuousVestingAccount)
	if !ok {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - account is not ContinuousVestingAccount", "ownerAddress", ownerAddress)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "account %s is not ContinuousVestingAccount", ownerAddress)
	}

	lockedCoins := vestingAcc.LockedCoins(ctx.BlockTime())

	if !amountToUnlock.IsAllLTE(lockedCoins) {
		k.Logger(ctx).Debug("unlock unbonded continuous vesting account coins - not enough to unlock", "account", ownerAccount, "lockedCoins", lockedCoins, "amountToUnlock", amountToUnlock)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "account %s: not enough to unlock. locked: %s, to unlock: %s", ownerAddress, lockedCoins, amountToUnlock)
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
				vestingAcc.OriginalVesting = vestingAcc.OriginalVesting.Sub(sdk.NewCoin(coin.Denom, sdk.NewInt(bankersRoundingCompensation)))
				// Subtracting 1 is done to compensate bankers rounding of vesting coins calculation nad truncating of original vesting difference.
				// TODO Some better explanation maybe?
				// Variables used in the calculations are:
				// OV - Oryginal Vesting
				// U - the amount of time left until unlocking
				// VgC - the Vesting Amount that corresponds to the current moment before unlocking
				// OVnew - the newly calculated OV that allows for unlocking U
				// a - a factor indicating what percentage of OV is currently in vesting (VgC/OV)
				//
				// Expected value of OVnew:
				// OVnew = OV - (U * OV / VgC)
				//
				// However, since bankers rounding and truncation are used when calculating VgC, we have:
				//
				// OVnew = OV - truncate(U * OV / (VgC ± 0.5))
				//
				// Vesting Amount value before unlocking:
				// VgC = a*OV
				//
				// Vesting Amount value after unlocking::
				// VgC-after = a*OVnew
				// VgC-after = a*(OV - truncate(U * OV / (VgC ± 0.5)))
				//
				// The difference between the old and new VgC is equal to amount unlocked.
				// Unlocked = a*OV - a*(OV - truncate(U * OV / (VgC ± 0.5))) = a* truncate(U * OV / (VgC ± 0.5))
				//
				// Continuing by substituting
				// Unlocked = a*truncate(U * OV / (a*OV ± 0.5))
				//
				// The calculation error Δ can be represented by U - Unlocked.
				// Δ = U - a*truncate(U * OV / (a*OV ± 0.5))
				//
				// We must further take into account the following constraints:
				// OV is positive integer
				// 1 <= toUnlock <= a*OV
				// 0 < a <=1
				// a*OV >= 1
				// OV >= 1
				//
				// Let us consider extreme cases:
				//
				// 1. a = 1, U = OV
				//
				// Δ = OV - truncate(OV * OV / (OV ± 0.5))
				// For OV >= 1 Δ ∈ <-1 ; 0.333333> when truncate is not used
				//
				// -1 can only occur when OV = 1, which is an extreme case where bankers rounding does nothing (there is no really rounding and we can skip ± 0.5). Therefore, -1 will not actually occur. We have a range of (-1; 0.333333>. For truncated values <OV-0.33333; OV+1), the result after truncation is <OV-1; OV>, so:
				// Δ ∈ <0; 1>
				//
				// 2. a = 1, U = 1
				// Δ = 1 - truncate(OV / (OV ± 0.5))
				// For OV >= 1 Δ ∈ <-1 ; 0.333333> when truncate is not used
				// Continuing identically to case 1.
				// Δ ∈ <0; 1>
				//
				// 3. a = 1/OV, U = 1
				// This is specific case where bankers rounding return original value, because OV*a = 1 and bankers rounding of 1 is 1
				//
				// Δ = 1 - truncate(OV / (1))/OV = 0
				//
				// 4. a > 1/OV, U = 1
				//
				// Δ = 1 - truncate(OV / (a*OV ± 0.5))/OV
				//
				// a*OV > 1 => OV / (a*OV ± 0.5) < 2 *OV => truncate(OV / (a*OV ± 0.5))/OV < 2
				// Δ > -1, OV is positive =>
				// Δ ∈ (-1 ; 1>
				//
				// But for a*V ≈ 1 rounding is to 1 so truncate(OV / (a*OV ± 0.5))/OV = truncate(OV / 1)/OV = 1
				// for a*V ≈ 1.5 rounding is to 2 so truncate(OV / (a*OV ± 0.5))/OV = truncate(OV / 2)/OV < 1
				// For a*V ≈ 2 rounding is to 2 so truncate(OV / (a*OV ± 0.5))/OV = truncate(OV / 2)/OV < 1
				// And farther truncate(OV / n)/OV  < 1
				// Δ = 1 - truncate(OV / n)/OV => 0 and OV is positive integer => Δ ∈ <0 ;1>
				//
				// So finally:
				// Δ ∈ <0 ;1>, as Δ can be integer only then there are only 2 possible values 0 and 1.
				//
				// To additionally unlock missing 1 we need to subtract it from OVnew

				// Examples:
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

				// example: OV = 8999999999999999999, 1 hour after half of vesting time (vesting time = 1000h) passed and we want to unlock 1.
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
