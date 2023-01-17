package keeper

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) SendToAirdropAccount(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Coins, startTime int64, endTime int64, initialClaim bool) error {
	k.Logger(ctx).Debug("send to airdrop account", "toAddress", toAddress,
		"amount", amount, "startTime", startTime, "endTime", endTime, "initialClaim", initialClaim)
	ak := k.accountKeeper
	bk := k.bankKeeper

	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		k.Logger(ctx).Error("send to airdrop account send coins disabled", "error", err.Error())
		return sdkerrors.Wrap(err, "send to airdrop account - send coins disabled")
	}

	if bk.BlockedAddr(toAddress) {
		k.Logger(ctx).Error("send to airdrop account account is not allowed to receive funds error", "toAddress", toAddress)
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "send to airdrop account - account address: %s is not allowed to receive funds error", toAddress)
	}

	acc := ak.GetAccount(ctx, toAddress)
	if initialClaim {
		baseAccount := ak.NewAccountWithAddress(ctx, toAddress)
		if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
			k.Logger(ctx).Error("send to airdrop account invalid account type; expected: BaseAccount", "notExpectedAccount", baseAccount)
			return sdkerrors.Wrapf(c4eerrors.ErrInvalidAccountType, "send to airdrop account - expected BaseAccount, got: %T", baseAccount)
		}

		baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), endTime)

		newAirdropAcc := types.NewAirdropVestingAccountRaw(baseVestingAccount, startTime)
		newAirdropAcc.EndTime = endTime
		acc = newAirdropAcc
		k.Logger(ctx).Debug("send to airdrop account", "baseAccount", baseVestingAccount.BaseAccount, "originalVesting",
			baseVestingAccount.OriginalVesting, "delegatedFree", baseVestingAccount.DelegatedFree, "delegatedVesting",
			baseVestingAccount.DelegatedVesting, "endTime", baseVestingAccount.EndTime, "startTime", startTime)
	}

	if acc == nil {
		k.Logger(ctx).Error("send to airdrop account - account not exists error", "toAddress", toAddress)
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "create airdrop account - account does not exist: %s", toAddress)
	}
	airdropAccount, ok := acc.(*types.AirdropVestingAccount)
	if !ok {
		k.Logger(ctx).Error("send to airdrop account invalid account type; expected: AirdropVestingAccount", "notExpectedAccount", acc)
		return sdkerrors.Wrapf(c4eerrors.ErrInvalidAccountType, "send to airdrop account - expected AirdropVestingAccount, got: %T", acc)
	}
	ak.SetAccount(ctx, airdropAccount)
	hadPariods := len(airdropAccount.VestingPeriods) > 0
	airdropAccount.VestingPeriods = append(airdropAccount.VestingPeriods,
		types.ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: amount})
	airdropAccount.BaseVestingAccount.OriginalVesting = airdropAccount.BaseVestingAccount.OriginalVesting.Add(amount...)
	if !hadPariods || endTime > airdropAccount.BaseVestingAccount.EndTime {
		airdropAccount.BaseVestingAccount.EndTime = endTime
	}
	if !hadPariods || startTime < airdropAccount.StartTime {
		airdropAccount.StartTime = startTime
	}

	if err := bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddress, amount); err != nil {
		k.Logger(ctx).Debug("send to airdrop account send coins to vesting account error", "toAddress", toAddress,
			"amount", amount, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrSendCoins, sdkerrors.Wrapf(err,
			"send to airdrop account - send coins to airdrop account error (to: %s, amount: %s)", toAddress, amount).Error())
	}
	ak.SetAccount(ctx, airdropAccount)
	return nil
}
