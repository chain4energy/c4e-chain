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
	"github.com/tendermint/tendermint/libs/log"
)

func (k Keeper) SendToNewRepeatedContinuousVestingAccount(ctx sdk.Context, userEntry *types.UserEntry,
	amount sdk.Coins, startTime int64, endTime int64, missionType types.MissionType) error {
	logger := ctx.Logger().With("send to airdrop account", "userEntry", userEntry,
		"amount", amount, "startTime", startTime, "endTime", endTime, "missionType", missionType)

	claimerAddress, err := k.getClaimerAddress(logger, userEntry.ClaimAddress)
	if err != nil {
		return err
	}

	claimerAccount, err := k.getClaimerAccount(logger, ctx, claimerAddress, missionType, startTime, endTime)
	if err != nil {
		return err
	}

	if k.bankKeeper.GetAllBalances(ctx, k.accountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress()).IsAllLT(amount) {
		logger.Debug("account insufficient funds error", "claimerAddress", claimerAddress, "amount", amount)
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, fmt.Sprintf(
			"send to airdrop account - send coins to airdrop account insufficient funds error (to: %s, amount: %s)", claimerAddress, amount))
	}

	claimerAccount = setupNewContinousVestingPeriods(claimerAccount, startTime, endTime, amount)

	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimerAddress, amount); err != nil {
		logger.Debug(" send coins to vesting account error", "toAddress", claimerAddress,
			"amount", amount, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrSendCoins, sdkerrors.Wrapf(err,
			"send to airdrop account - send coins to airdrop account error (to: %s, amount: %s)", claimerAddress, amount).Error())
	}

	k.accountKeeper.SetAccount(ctx, claimerAccount)
	return nil
}

func (k Keeper) getClaimerAddress(logger log.Logger, claimer string) (sdk.AccAddress, error) {
	claimerAddress, err := sdk.AccAddressFromBech32(claimer)
	if err != nil {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: "+err.Error(), claimer)
	}
	if k.bankKeeper.BlockedAddr(claimerAddress) {
		logger.Debug("account is not allowed to receive funds error")
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "send to airdrop account - account address: %s is not allowed to receive funds error", claimerAddress)
	}
	return claimerAddress, nil
}

func (k Keeper) getClaimerAccount(logger log.Logger, ctx sdk.Context, claimerAddress sdk.AccAddress, missionType types.MissionType, startTime, endTime int64) (*cfevestingtypes.RepeatedContinuousVestingAccount, error) {
	claimerAccount := k.accountKeeper.GetAccount(ctx, claimerAddress)
	_, ok := claimerAccount.(*cfevestingtypes.RepeatedContinuousVestingAccount)
	if missionType == types.MissionInitialClaim && !ok {
		var err error
		claimerAccount, err = k.setupNewClaimerAccountForInitialClaim(ctx, claimerAddress, startTime, endTime)
		if err != nil {
			return nil, err
		}
	}

	if claimerAccount == nil {
		logger.Debug("account does not exist error")
		return nil, sdkerrors.Wrapf(c4eerrors.ErrNotExists, "send to airdrop account - account does not exist: %s", claimerAddress)
	}

	repeatedContinousVestingAccount, ok := claimerAccount.(*cfevestingtypes.RepeatedContinuousVestingAccount)
	if !ok {
		logger.Debug("invalid account type; expected: RepeatedContinuousVestingAccount", "notExpectedAccount", claimerAccount)
		return nil, sdkerrors.Wrapf(c4eerrors.ErrInvalidAccountType, "send to airdrop account - expected RepeatedContinuousVestingAccount, got: %T", claimerAccount)
	}
	return repeatedContinousVestingAccount, nil
}

func setupNewContinousVestingPeriods(claimerAccount *cfevestingtypes.RepeatedContinuousVestingAccount, startTime int64, endTime int64, amount sdk.Coins) *cfevestingtypes.RepeatedContinuousVestingAccount {
	hadPariods := len(claimerAccount.VestingPeriods) > 0

	claimerAccount.VestingPeriods = append(claimerAccount.VestingPeriods,
		cfevestingtypes.ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: amount})

	claimerAccount.BaseVestingAccount.OriginalVesting = claimerAccount.BaseVestingAccount.OriginalVesting.Add(amount...)
	if !hadPariods || endTime > claimerAccount.BaseVestingAccount.EndTime {
		claimerAccount.BaseVestingAccount.EndTime = endTime
	}
	if !hadPariods || startTime < claimerAccount.StartTime {
		claimerAccount.StartTime = startTime
	}
	return claimerAccount
}

func (k Keeper) setupNewClaimerAccountForInitialClaim(ctx sdk.Context, claimerAddress sdk.AccAddress, startTime int64, endTime int64) (*cfevestingtypes.RepeatedContinuousVestingAccount, error) {
	baseAccount := k.accountKeeper.NewAccountWithAddress(ctx, claimerAddress)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Error("send to airdrop account invalid account type; expected: BaseAccount", "notExpectedAccount", baseAccount)
		return nil, sdkerrors.Wrapf(c4eerrors.ErrInvalidAccountType, "send to airdrop account - expected BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), endTime)

	newAcc := cfevestingtypes.NewRepeatedContinuousVestingAccountRaw(baseVestingAccount, startTime)
	newAcc.EndTime = endTime
	return newAcc, nil
}
