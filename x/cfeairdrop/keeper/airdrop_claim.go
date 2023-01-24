package keeper

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) InitialClaim(ctx sdk.Context, claimer string, campaignId uint64, additionalAddress string) error {
	var addressToClaim = claimer
	if additionalAddress != "" {
		if err := k.validateAdditionalAddressToClaim(ctx, additionalAddress); err != nil {
			return err
		}
		addressToClaim = additionalAddress
	}

	campaign, mission, userAirdropEntries, err := k.missionFirstStep(ctx, "claim initial mission", campaignId, types.InitialMissionId, addressToClaim, false)
	if err != nil {
		return err
	}
	userAirdropEntries.ClaimAddress = addressToClaim

	userAirdropEntries, err = k.completeMission(ctx, mission, userAirdropEntries)
	if err != nil {
		return err
	}

	claimableAmount := k.calculateInitialClaimClaimableAmount(ctx, campaignId, userAirdropEntries)

	claimableAmount, err = k.calculateAndSendInitialClaimFreeAmount(ctx, campaignId, userAirdropEntries, claimableAmount, campaign.InitialClaimFreeAmount)
	if err != nil {
		return err
	}

	userAirdropEntries, err = k.claimMission(ctx, campaign, mission, userAirdropEntries, claimableAmount)
	if err != nil {
		return err
	}

	k.SetUserAirdropEntries(ctx, *userAirdropEntries)
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		granteeAddr, err := sdk.AccAddressFromBech32(userAirdropEntries.Address)
		if err != nil {
			return err
		}
		_, accountAddr := feegrantAccountAddress(campaignId)
		if err = k.revokeFeeAllowance(ctx, accountAddr, granteeAddr); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) Claim(ctx sdk.Context, campaignId uint64, missionId uint64, claimer string) error {
	campaign, mission, userAirdropEntries, err := k.missionFirstStep(ctx, "claim mission", campaignId, missionId, claimer, false)
	if err != nil {
		return err
	}

	if !userAirdropEntries.IsInitialMissionClaimed(campaignId) {
		k.Logger(ctx).Error("complete mission - initial mission not completed", "claimerAddress", userAirdropEntries.ClaimAddress, "campaignId", campaignId, "missionId", missionId)
		return sdkerrors.Wrapf(types.ErrMissionNotCompleted, "initial mission not completed: address %s, campaignId: %d, missionId: %d", userAirdropEntries.ClaimAddress, campaignId, 0)
	}

	if mission.MissionType == types.MissionClaim {
		userAirdropEntries, err = k.completeMission(ctx, mission, userAirdropEntries)
		if err != nil {
			return err
		}
	}

	claimableAmount := userAirdropEntries.ClaimableFromMission(mission)
	userAirdropEntries, err = k.claimMission(ctx, campaign, mission, userAirdropEntries, claimableAmount)
	if err != nil {
		return err
	}

	k.SetUserAirdropEntries(ctx, *userAirdropEntries)
	return nil
}

func (k Keeper) validateAdditionalAddressToClaim(ctx sdk.Context, additionalAddress string) error {
	addititonalAccAddress, err := sdk.AccAddressFromBech32(additionalAddress)
	if err != nil {
		k.Logger(ctx).Error("add mission to airdrop campaign claimer parsing error", "additionalAddress", additionalAddress, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "add mission to airdrop campaign - additionalAddress parsing error: %s", additionalAddress).Error())
	}
	if k.bankKeeper.BlockedAddr(addititonalAccAddress) {
		k.Logger(ctx).Error("new vesting account is not allowed to receive funds error", "address", additionalAddress)
		return sdkerrors.Wrapf(c4eerrors.ErrAccountNotAllowedToReceiveFunds, "new vesting account - account address: %s", additionalAddress)
	}

	account := k.accountKeeper.GetAccount(ctx, addititonalAccAddress)
	_, ok := account.(*vestingtypes.BaseVestingAccount)
	if ok {
		k.Logger(ctx).Error("new vesting account is not allowed to receive funds error", "address", additionalAddress)
		return sdkerrors.Wrapf(c4eerrors.ErrAccountNotAllowedToReceiveFunds, "new vesting account - account address: %s", additionalAddress)
	}

	return nil
}

func (k Keeper) calculateInitialClaimClaimableAmount(ctx sdk.Context, campaignId uint64, userAirdropEntries *types.UserAirdropEntries) sdk.Coins {
	allCampaignMissions, _ := k.AllMissionForCampaign(ctx, campaignId)
	airdropEntry := userAirdropEntries.GetAidropEntry(campaignId)
	allMissionsAmountSum := sdk.NewCoins()
	for _, mission := range allCampaignMissions {
		for _, amount := range airdropEntry.AirdropCoins {
			if mission.Weight != nil {
				allMissionsAmountSum = allMissionsAmountSum.Add(sdk.NewCoin(amount.Denom, mission.Weight.Mul(sdk.NewDecFromInt(amount.Amount)).TruncateInt()))
			}
		}
	}
	return airdropEntry.AirdropCoins.Sub(allMissionsAmountSum)
}

func (k Keeper) calculateAndSendInitialClaimFreeAmount(ctx sdk.Context, campaignId uint64, userAirdropEntries *types.UserAirdropEntries, claimableAmount sdk.Coins, initialClaimFreeAmount sdk.Int) (sdk.Coins, error) {
	userMainAddress, err := sdk.AccAddressFromBech32(userAirdropEntries.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: "+err.Error(), userAirdropEntries.Address)
	}
	if k.bankKeeper.BlockedAddr(userMainAddress) {
		k.Logger(ctx).Error("send to airdrop account account is not allowed to receive funds error", "userMainAddress", userMainAddress)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "send to airdrop account - account address: %s is not allowed to receive funds error", userMainAddress)
	}

	freeVestingAmount := sdk.NewCoins()
	for _, claimableAmountCoin := range claimableAmount {
		if claimableAmountCoin.Sub(sdk.NewCoin(claimableAmountCoin.Denom, initialClaimFreeAmount)).IsNegative() {
			k.Logger(ctx).Error("send to airdrop account wrong send coins amount. Amount < 1 token (1000000)", "amount", claimableAmountCoin.Amount, "denom", claimableAmountCoin.Denom)
			return nil, sdkerrors.Wrapf(c4eerrors.ErrSendCoins, "send to airdrop account  wrong send coins amount. %s < 1 token (1000000 %s)", claimableAmountCoin.String(), claimableAmountCoin.Denom)
		}
		coin := sdk.NewCoins(sdk.NewCoin(claimableAmountCoin.Denom, initialClaimFreeAmount))
		freeVestingAmount = freeVestingAmount.Add(coin...)

	}
	claimableAmount = claimableAmount.Sub(freeVestingAmount)

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userMainAddress, freeVestingAmount); err != nil {
		k.Logger(ctx).Debug("send to airdrop account send coins to vesting account error", "toAddress", userMainAddress,
			"freeVestingAmount", freeVestingAmount, "error", err.Error())
		return nil, sdkerrors.Wrap(c4eerrors.ErrSendCoins, sdkerrors.Wrapf(err,
			"send to airdrop account - send coins to airdrop account error (to: %s, amount: %s)", userMainAddress, freeVestingAmount.String()).Error())
	}
	k.DecrementAirdropClaimsLeft(ctx, campaignId, freeVestingAmount)
	return claimableAmount, nil
}
