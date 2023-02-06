package keeper

import (
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) InitialClaim(ctx sdk.Context, claimer string, campaignId uint64, additionalAddress string) error {
	addressToClaim := claimer
	if additionalAddress != "" {
		if err := k.validateAdditionalAddressToClaim(ctx, additionalAddress); err != nil {
			return err
		}
		addressToClaim = additionalAddress
	}

	campaign, mission, userEntry, err := k.missionFirstStep(ctx, campaignId, types.InitialMissionId, addressToClaim)
	if err != nil {
		return err
	}
	userEntry.ClaimAddress = addressToClaim

	userEntry, err = k.completeMission(ctx, mission, userEntry)
	if err != nil {
		return err
	}

	claimableAmount := k.calculateInitialClaimClaimableAmount(ctx, campaignId, userEntry)

	claimableAmount, err = k.calculateAndSendInitialClaimFreeAmount(ctx, campaignId, userEntry, claimableAmount, campaign.InitialClaimFreeAmount)
	if err != nil {
		return err
	}

	userEntry, err = k.claimMission(ctx, campaign, mission, userEntry, claimableAmount)
	if err != nil {
		return err
	}

	k.SetUserEntry(ctx, *userEntry)

	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		granteeAddr, err := sdk.AccAddressFromBech32(userEntry.Address)
		if err != nil {
			return err
		}
		_, accountAddr := FeegrantAccountAddress(campaignId)
		if err = k.revokeFeeAllowance(ctx, accountAddr, granteeAddr); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) Claim(ctx sdk.Context, campaignId uint64, missionId uint64, claimer string) error {
	campaign, mission, userEntry, err := k.missionFirstStep(ctx, campaignId, missionId, claimer)
	if err != nil {
		return err
	}

	if !userEntry.IsInitialMissionClaimed(campaignId) {
		k.Logger(ctx).Error("complete mission - initial mission not completed", "claimerAddress", userEntry.Address, "campaignId", campaignId, "missionId", missionId)
		return sdkerrors.Wrapf(types.ErrMissionNotCompleted, fmt.Sprintf("initial mission not completed: address %s, campaignId: %d", userEntry.Address, campaignId))
	}

	if mission.MissionType == types.MissionClaim {
		userEntry, err = k.completeMission(ctx, mission, userEntry)
		if err != nil {
			return err
		}
	}

	claimableAmount := userEntry.ClaimableFromMission(mission)
	userEntry, err = k.claimMission(ctx, campaign, mission, userEntry, claimableAmount)
	if err != nil {
		return err
	}

	k.SetUserEntry(ctx, *userEntry)
	return nil
}

func (k Keeper) CompleteMissionFromHook(ctx sdk.Context, campaignId uint64, missionId uint64, address string) error {
	_, mission, userEntry, err := k.missionFirstStep(ctx, campaignId, missionId, address)
	if err != nil {
		return err
	}
	if !userEntry.IsInitialMissionClaimed(campaignId) {
		k.Logger(ctx).Error("complete mission - initial mission not completed", "claimerAddress", address, "campaignId", campaignId, "missionId", missionId)
		return sdkerrors.Wrapf(types.ErrMissionNotCompleted, "initial mission not completed: address %s, campaignId: %d, missionId: %d", address, campaignId, 0)
	}
	userEntry, err = k.completeMission(ctx, mission, userEntry)
	if err != nil {
		return err
	}
	k.SetUserEntry(ctx, *userEntry)
	return nil
}

func (k Keeper) completeMission(ctx sdk.Context, mission *types.Mission, userEntry *types.UserEntry) (*types.UserEntry, error) {
	campaignId := mission.CampaignId
	missionId := mission.Id
	address := userEntry.Address

	if userEntry.IsMissionCompleted(campaignId, missionId) {
		k.Logger(ctx).Debug("complete mission - mission already completed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionCompleted, "mission already completed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if err := userEntry.CompleteMission(campaignId, missionId); err != nil {
		k.Logger(ctx).Debug("complete mission - cannot complete", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionCompletion, err.Error())
	}

	return userEntry, nil
}

func (k Keeper) claimMission(ctx sdk.Context, campaign *types.Campaign, mission *types.Mission, userEntry *types.UserEntry, claimableAmount sdk.Coins) (*types.UserEntry, error) {
	campaignId := mission.CampaignId
	missionId := mission.Id
	address := userEntry.ClaimAddress

	if !userEntry.IsMissionCompleted(campaignId, missionId) {
		k.Logger(ctx).Error("claim mission - mission not completed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionNotCompleted, "mission not completed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if userEntry.IsMissionClaimed(campaignId, missionId) {
		k.Logger(ctx).Error("claim mission - mission already claimed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionClaimed, "mission already claimed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if err := userEntry.ClaimMission(campaignId, missionId); err != nil {
		k.Logger(ctx).Error("claim mission - cannot claime", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionClaiming, err.Error())
	}

	start := ctx.BlockTime().Add(campaign.LockupPeriod)
	end := start.Add(campaign.VestingPeriod)

	if err := k.SendToNewRepeatedContinuousVestingAccount(ctx, userEntry, claimableAmount, start.Unix(), end.Unix(), mission.MissionType); err != nil {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrSendCoins, "send to claiming address %s error: "+err.Error(), userEntry.ClaimAddress)
	}

	k.DecrementCampaignAmountLeft(ctx, campaignId, claimableAmount)
	return userEntry, nil
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

func (k Keeper) calculateInitialClaimClaimableAmount(ctx sdk.Context, campaignId uint64, userEntry *types.UserEntry) sdk.Coins {
	allCampaignMissions, _ := k.AllMissionForCampaign(ctx, campaignId)
	claimRecord := userEntry.GetClaimRecord(campaignId)
	allMissionsAmountSum := sdk.NewCoins()
	for _, mission := range allCampaignMissions {
		for _, amount := range claimRecord.Amount {
			allMissionsAmountSum = allMissionsAmountSum.Add(sdk.NewCoin(amount.Denom, mission.Weight.Mul(sdk.NewDecFromInt(amount.Amount)).TruncateInt()))
		}
	}
	return claimRecord.Amount.Sub(allMissionsAmountSum)
}

func (k Keeper) calculateAndSendInitialClaimFreeAmount(ctx sdk.Context, campaignId uint64, userEntry *types.UserEntry, claimableAmount sdk.Coins, initialClaimFreeAmount sdk.Int) (sdk.Coins, error) {
	userMainAddress, err := sdk.AccAddressFromBech32(userEntry.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: "+err.Error(), userEntry.Address)
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
	k.DecrementCampaignAmountLeft(ctx, campaignId, freeVestingAmount)
	return claimableAmount, nil
}
