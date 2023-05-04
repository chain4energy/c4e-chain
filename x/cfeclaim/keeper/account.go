package keeper

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"
)

func (k Keeper) SendToNewPeriodicContinuousVestingAccount(ctx sdk.Context, userEntry *types.UserEntry,
	amount sdk.Coins, startTime int64, endTime int64, missionType types.MissionType) error {
	logger := ctx.Logger().With("send to claim account", "userEntry", userEntry,
		"amount", amount, "startTime", startTime, "endTime", endTime, "missionType", missionType)

	claimerAddress, err := k.getClaimerAddress(logger, userEntry.ClaimAddress)
	if err != nil {
		return err
	}

	claimerAccount, err := k.getOrCreateClaimerAccount(logger, ctx, claimerAddress, missionType, startTime, endTime)
	if err != nil {
		return err
	}

	claimerAccount = cfevestingtypes.AddNewContinousVestingPeriods(claimerAccount, startTime, endTime, amount)

	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimerAddress, amount); err != nil {
		return err
	}

	k.accountKeeper.SetAccount(ctx, claimerAccount)
	return nil
}

func (k Keeper) getClaimerAddress(logger log.Logger, claimer string) (sdk.AccAddress, error) {
	claimerAddress, err := sdk.AccAddressFromBech32(claimer)
	if err != nil {
		return nil, errors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: "+err.Error(), claimer)
	}
	if k.bankKeeper.BlockedAddr(claimerAddress) {
		logger.Debug("account is not allowed to receive funds error")
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized, "send to claim account - account address: %s is not allowed to receive funds error", claimerAddress)
	}
	return claimerAddress, nil
}

func (k Keeper) getOrCreateClaimerAccount(logger log.Logger, ctx sdk.Context, claimerAddress sdk.AccAddress, missionType types.MissionType, startTime, endTime int64) (*cfevestingtypes.PeriodicContinuousVestingAccount, error) {
	claimerAccount := k.accountKeeper.GetAccount(ctx, claimerAddress)
	_, ok := claimerAccount.(*cfevestingtypes.PeriodicContinuousVestingAccount)
	if missionType == types.MissionInitialClaim && !ok {
		var err error
		claimerAccount, err = k.vestingKeeper.SetupNewPeriodicContinousVestingAccount(ctx, claimerAddress, startTime, endTime)
		if err != nil {
			return nil, err
		}
	}

	if claimerAccount == nil {
		logger.Debug("account does not exist error")
		return nil, errors.Wrapf(c4eerrors.ErrNotExists, "send to claim account - account does not exist: %s", claimerAddress)
	}

	periodicContinuousVestingAccount, ok := claimerAccount.(*cfevestingtypes.PeriodicContinuousVestingAccount)
	if !ok {
		logger.Debug("invalid account type; expected: PeriodicContinuousVestingAccount", "notExpectedAccount", claimerAccount)
		return nil, errors.Wrapf(c4eerrors.ErrInvalidAccountType, "send to claim account - expected PeriodicContinuousVestingAccount, got: %T", claimerAccount)
	}
	return periodicContinuousVestingAccount, nil
}
