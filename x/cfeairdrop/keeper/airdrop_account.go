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

func (k Keeper) SendToNewRepeatedContinuousVestingAccount(ctx sdk.Context, userEntry *types.UserEntry,
	amount sdk.Coins, startTime int64, endTime int64, missionType types.MissionType) error {
	k.Logger(ctx).Debug("send to airdrop account", "userEntry", userEntry,
		"amount", amount, "startTime", startTime, "endTime", endTime, "missionType", missionType)
	ak := k.accountKeeper
	bk := k.bankKeeper
	claimerAddress, err := sdk.AccAddressFromBech32(userEntry.ClaimAddress)
	if err != nil {
		return sdkerrors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: "+err.Error(), userEntry.ClaimAddress)
	}
	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		k.Logger(ctx).Error("send to airdrop account send coins disabled", "error", err.Error())
		return sdkerrors.Wrap(err, "send to airdrop account - send coins disabled")
	}

	if bk.BlockedAddr(claimerAddress) {
		k.Logger(ctx).Error("send to airdrop account account is not allowed to receive funds error", "claimerAddress", claimerAddress)
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "send to airdrop account - account address: %s is not allowed to receive funds error", claimerAddress)
	}

	claimerAccount := ak.GetAccount(ctx, claimerAddress)
	_, ok := claimerAccount.(*cfevestingtypes.RepeatedContinuousVestingAccount)
	if missionType == types.MissionInitialClaim && !ok {
		baseAccount := ak.NewAccountWithAddress(ctx, claimerAddress)
		if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
			k.Logger(ctx).Error("send to airdrop account invalid account type; expected: BaseAccount", "notExpectedAccount", baseAccount)
			return sdkerrors.Wrapf(c4eerrors.ErrInvalidAccountType, "send to airdrop account - expected BaseAccount, got: %T", baseAccount)
		}

		baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), endTime)

		newAcc := cfevestingtypes.NewRepeatedContinuousVestingAccountRaw(baseVestingAccount, startTime)
		newAcc.EndTime = endTime
		claimerAccount = newAcc
		k.Logger(ctx).Debug("send to airdrop account", "baseAccount", baseVestingAccount.BaseAccount, "originalVesting",
			baseVestingAccount.OriginalVesting, "delegatedFree", baseVestingAccount.DelegatedFree, "delegatedVesting",
			baseVestingAccount.DelegatedVesting, "endTime", baseVestingAccount.EndTime, "startTime", startTime)
	}

	if claimerAccount == nil {
		k.Logger(ctx).Error("send to airdrop account - account not exists error", "claimerAddress", claimerAddress)
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "send to airdrop account - account does not exist: %s", claimerAddress)
	}
	airdropAccount, ok := claimerAccount.(*cfevestingtypes.RepeatedContinuousVestingAccount)
	if !ok {
		k.Logger(ctx).Error("send to airdrop account invalid account type; expected: RepeatedContinuousVestingAccount", "notExpectedAccount", claimerAccount)
		return sdkerrors.Wrapf(c4eerrors.ErrInvalidAccountType, "send to airdrop account - expected RepeatedContinuousVestingAccount, got: %T", claimerAccount)
	}
	if bk.GetAllBalances(ctx, ak.GetModuleAccount(ctx, types.ModuleName).GetAddress()).IsAllLT(amount) {
		k.Logger(ctx).Debug("send to airdrop account send coins to vesting account error", "toAddress", claimerAddress,
			"amount", amount)
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, fmt.Sprintf(
			"send to airdrop account - send coins to airdrop account insufficient funds error (to: %s, amount: %s)", claimerAddress, amount))
	}
	ak.SetAccount(ctx, airdropAccount)
	hadPariods := len(airdropAccount.VestingPeriods) > 0

	airdropAccount.VestingPeriods = append(airdropAccount.VestingPeriods,
		cfevestingtypes.ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: amount})

	airdropAccount.BaseVestingAccount.OriginalVesting = airdropAccount.BaseVestingAccount.OriginalVesting.Add(amount...)
	if !hadPariods || endTime > airdropAccount.BaseVestingAccount.EndTime {
		airdropAccount.BaseVestingAccount.EndTime = endTime
	}
	if !hadPariods || startTime < airdropAccount.StartTime {
		airdropAccount.StartTime = startTime
	}

	if err := bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimerAddress, amount); err != nil {
		k.Logger(ctx).Debug("send to airdrop account send coins to vesting account error", "toAddress", claimerAddress,
			"amount", amount, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrSendCoins, sdkerrors.Wrapf(err,
			"send to airdrop account - send coins to airdrop account error (to: %s, amount: %s)", claimerAddress, amount).Error())
	}

	ak.SetAccount(ctx, airdropAccount)
	return nil
}
