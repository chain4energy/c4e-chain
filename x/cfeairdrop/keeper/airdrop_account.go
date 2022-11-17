package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) SendToAirdropAccount(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Coins, startTime int64, endTime int64, createAccount bool) error {
	k.Logger(ctx).Debug("create airdrop account", "toAddress", toAddress,
		"amount", amount, "startTime", startTime, "endTime", endTime, "createAccount", createAccount)
	ak := k.accountKeeper
	bk := k.bankKeeper

	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		k.Logger(ctx).Error("create vesting account send coins disabled", "error", err.Error())
		return sdkerrors.Wrap(err, "create vesting account - send coins disabled")
	}

	// from, err := sdk.AccAddressFromBech32(fromAddress)
	// if err != nil {
	// 	k.Logger(ctx).Error("create vesting account from-address parsing error", "fromAddress", fromAddress, "error", err.Error())
	// 	return sdkerrors.Wrap(types.ErrSample /* TODO */, sdkerrors.Wrapf(err, "create vesting account - from-address parsing error: %s", fromAddress).Error())
	// }
	// to, err := sdk.AccAddressFromBech32(toAddress)
	// if err != nil {
	// 	k.Logger(ctx).Error("create vesting account to-address parsing error", "toAddress", toAddress, "error", err.Error())
	// 	return sdkerrors.Wrap(types.ErrSample /* TODO */, sdkerrors.Wrapf(err, "create vesting account - to-address parsing error: %s", toAddress).Error())
	// }

	if bk.BlockedAddr(toAddress) {
		k.Logger(ctx).Error("create vesting account account is not allowed to receive funds error", "toAddress", toAddress)
		return sdkerrors.Wrapf(types.ErrSample /* TODO */, "create vesting account - account address: %s", toAddress)
	}

	acc := ak.GetAccount(ctx, toAddress)
	if createAccount && acc == nil {
		// if acc != nil {
		// 	k.Logger(ctx).Error("create vesting account account already exists error", "toAddress", toAddress)
		// 	return sdkerrors.Wrapf(types.ErrSample /* TODO */, "create vesting account - account address: %s", toAddress)
		// }
		baseAccount := ak.NewAccountWithAddress(ctx, toAddress)
		if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
			k.Logger(ctx).Error("create airdrop account invalid account type; expected: BaseAccount", "notExpectedAccount", baseAccount)
			return sdkerrors.Wrapf(types.ErrSample /* TODO */, "create vesting account - expected BaseAccount, got: %T", baseAccount)
		}

		baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), endTime)

		acc = types.NewAirdropVestingAccountRaw(baseVestingAccount, startTime)
		k.Logger(ctx).Debug("create airdrop account", "baseAccount", baseVestingAccount.BaseAccount, "originalVesting",
			baseVestingAccount.OriginalVesting, "delegatedFree", baseVestingAccount.DelegatedFree, "delegatedVesting",
			baseVestingAccount.DelegatedVesting, "endTime", baseVestingAccount.EndTime, "startTime", startTime)
	}

	if acc == nil {
		k.Logger(ctx).Error("create vesting account account already exists error", "toAddress", toAddress)
		return sdkerrors.Wrapf(types.ErrSample /* TODO */, "create vesting account - account address: %s", toAddress)
	}
	airdropAccount, ok := acc.(*types.AirdropVestingAccount)
	if !ok {
		k.Logger(ctx).Error("create fffffff account invalid account type; expected: BaseAccount", "notExpectedAccount", acc)
		return sdkerrors.Wrapf(types.ErrSample /* TODO */, "create vesting account - expected BaseAccount, got: %T", airdropAccount)
	}

	airdropAccount.VestingPeriods = append(airdropAccount.VestingPeriods,
		types.ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: amount})
	airdropAccount.BaseVestingAccount.OriginalVesting = airdropAccount.BaseVestingAccount.OriginalVesting.Add(amount...)
	if endTime > airdropAccount.BaseVestingAccount.EndTime {
		airdropAccount.BaseVestingAccount.EndTime = endTime
	}
	if startTime < airdropAccount.StartTime {
		airdropAccount.StartTime = startTime
	}
	ak.SetAccount(ctx, airdropAccount)

	// err = ctx.EventManager().EmitTypedEvent(&types.NewVestingAccount{
	// 	Address: acc.Address,
	// })
	// if err != nil {
	// 	k.Logger(ctx).Error("new vestig account emit event error", "error", err.Error())
	// }

	if err := bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddress, amount); err != nil {
		k.Logger(ctx).Debug("create vesting account send coins to vesting account error", "toAddress", toAddress,
			"amount", amount, "error", err.Error())
		return sdkerrors.Wrap(types.ErrSample /* TODO */, sdkerrors.Wrapf(err,
			"create vesting account - send coins to vesting account error (to: %s, amount: %s)", toAddress, amount).Error())
	}
	return nil
}
