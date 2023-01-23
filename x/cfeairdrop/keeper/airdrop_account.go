package keeper

import (
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) SendToAirdropAccount(ctx sdk.Context, userAirdropEntries *types.UserAirdropEntries,
	amount sdk.Coins, startTime int64, endTime int64, freeamount sdk.Int, initialClaim bool) error {
	k.Logger(ctx).Debug("send to airdrop account", "userAirdropEntries", userAirdropEntries,
		"amount", amount, "startTime", startTime, "endTime", endTime, "initialClaim", initialClaim)
	ak := k.accountKeeper
	bk := k.bankKeeper
	claimer, err := sdk.AccAddressFromBech32(userAirdropEntries.ClaimAddress)
	if err != nil {
		return sdkerrors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: "+err.Error(), userAirdropEntries.ClaimAddress)
	}
	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		k.Logger(ctx).Error("send to airdrop account send coins disabled", "error", err.Error())
		return sdkerrors.Wrap(err, "send to airdrop account - send coins disabled")
	}

	if bk.BlockedAddr(claimer) {
		k.Logger(ctx).Error("send to airdrop account account is not allowed to receive funds error", "claimer", claimer)
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "send to airdrop account - account address: %s is not allowed to receive funds error", claimer)
	}

	acc := ak.GetAccount(ctx, claimer)
	_, ok := acc.(*cfevestingtypes.RepeatedContinuousVestingAccount)
	if initialClaim && !ok {
		baseAccount := ak.NewAccountWithAddress(ctx, claimer)
		if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
			k.Logger(ctx).Error("send to airdrop account invalid account type; expected: BaseAccount", "notExpectedAccount", baseAccount)
			return sdkerrors.Wrapf(c4eerrors.ErrInvalidAccountType, "send to airdrop account - expected BaseAccount, got: %T", baseAccount)
		}

		baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), endTime)

		newAirdropAcc := cfevestingtypes.NewRepeatedContinuousVestingAccountRaw(baseVestingAccount, startTime)
		newAirdropAcc.EndTime = endTime
		acc = newAirdropAcc
		k.Logger(ctx).Debug("send to airdrop account", "baseAccount", baseVestingAccount.BaseAccount, "originalVesting",
			baseVestingAccount.OriginalVesting, "delegatedFree", baseVestingAccount.DelegatedFree, "delegatedVesting",
			baseVestingAccount.DelegatedVesting, "endTime", baseVestingAccount.EndTime, "startTime", startTime)
	}

	if acc == nil {
		k.Logger(ctx).Error("send to airdrop account - account not exists error", "claimer", claimer)
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "create airdrop account - account does not exist: %s", claimer)
	}
	airdropAccount, ok := acc.(*cfevestingtypes.RepeatedContinuousVestingAccount)
	if !ok {
		k.Logger(ctx).Error("send to airdrop account invalid account type; expected: RepeatedContinuousVestingAccount", "notExpectedAccount", acc)
		return sdkerrors.Wrapf(c4eerrors.ErrInvalidAccountType, "send to airdrop account - expected RepeatedContinuousVestingAccount, got: %T", acc)
	}
	if bk.GetAllBalances(ctx, ak.GetModuleAccount(ctx, types.ModuleName).GetAddress()).IsAllLT(amount) {
		k.Logger(ctx).Debug("send to airdrop account send coins to vesting account error", "toAddress", claimer,
			"amount", amount)
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, fmt.Sprintf(
			"send to airdrop account - send coins to airdrop account insufficient funds error (to: %s, amount: %s)", claimer, amount))
	}
	vestingAmount := amount
	if initialClaim {
		for _, coin := range amount {
			if coin.Sub(types.OneForthC4e).IsNegative() {
				k.Logger(ctx).Error("send to airdrop account wrong send coins amount. Amount < 1 token (1000000)", "amount", coin.Amount, "denom", coin.Denom)
				return sdkerrors.Wrapf(c4eerrors.ErrSendCoins, "send to airdrop account  wrong send coins amount. %s < 1 token (1000000 %s)", coin.String(), coin.Denom)
			}
			freeVestingAmount := sdk.NewCoins()
			for _, amount := range vestingAmount {
				coin := sdk.NewCoins(sdk.NewCoin(amount.Denom, freeamount))
				freeVestingAmount.Add(coin...)
				vestingAmount = vestingAmount.Sub(coin)
			}

			mainAddress, err := sdk.AccAddressFromBech32(userAirdropEntries.Address)
			if err != nil {
				return sdkerrors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: "+err.Error(), userAirdropEntries.Address)
			}
			if bk.BlockedAddr(mainAddress) {
				k.Logger(ctx).Error("send to airdrop account account is not allowed to receive funds error", "mainAddress", mainAddress)
				return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "send to airdrop account - account address: %s is not allowed to receive funds error", mainAddress)
			}
			if err := bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, mainAddress, freeVestingAmount); err != nil {
				k.Logger(ctx).Debug("send to airdrop account send coins to vesting account error", "toAddress", mainAddress,
					"amount", amount, "error", err.Error())
				return sdkerrors.Wrap(c4eerrors.ErrSendCoins, sdkerrors.Wrapf(err,
					"send to airdrop account - send coins to airdrop account error (to: %s, amount: %s)", mainAddress, amount).Error())
			}
		}
	}
	ak.SetAccount(ctx, airdropAccount)
	hadPariods := len(airdropAccount.VestingPeriods) > 0
	airdropAccount.VestingPeriods = append(airdropAccount.VestingPeriods,
		cfevestingtypes.ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: vestingAmount})
	airdropAccount.BaseVestingAccount.OriginalVesting = airdropAccount.BaseVestingAccount.OriginalVesting.Add(vestingAmount...)
	if !hadPariods || endTime > airdropAccount.BaseVestingAccount.EndTime {
		airdropAccount.BaseVestingAccount.EndTime = endTime
	}
	if !hadPariods || startTime < airdropAccount.StartTime {
		airdropAccount.StartTime = startTime
	}

	if err := bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimer, vestingAmount); err != nil {
		k.Logger(ctx).Debug("send to airdrop account send coins to vesting account error", "toAddress", claimer,
			"vestingAmount", vestingAmount, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrSendCoins, sdkerrors.Wrapf(err,
			"send to airdrop account - send coins to airdrop account error (to: %s, amount: %s)", claimer, amount).Error())
	}
	ak.SetAccount(ctx, airdropAccount)
	return nil
}
