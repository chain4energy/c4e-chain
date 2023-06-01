package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"strconv"
)

func (k Keeper) InitialClaim(ctx sdk.Context, claimer string, campaignId uint64, destinationAddress string) (sdk.Coins, error) {
	k.Logger(ctx).Debug("initial claim", "claimer", claimer, "campaignId", campaignId, "destinationAddress", destinationAddress)

	campaign, mission, userEntry, claimRecord, err := k.prepareClaimData(ctx, campaignId, types.InitialMissionId, claimer)
	if err != nil {
		return nil, err
	}

	if err = k.validateDestinationAddress(ctx, destinationAddress); err != nil {
		return nil, err
	}
	claimRecord.Address = destinationAddress

	if err = claimRecord.CompleteMission(campaignId, mission.Id); err != nil {
		return nil, err
	}

	claimableAmount, updatedFree := k.calculateInitialClaimClaimableAmount(ctx, campaign, claimRecord)

	err = k.claimMission(ctx, campaign, mission, claimRecord, claimableAmount, updatedFree)
	if err != nil {
		return nil, err
	}

	if campaign.FeegrantAmount.GT(math.ZeroInt()) {
		granteeAddr, err := sdk.AccAddressFromBech32(userEntry.Address)
		if err != nil {
			return nil, err
		}
		_, accountAddr := CreateFeegrantAccountAddress(campaignId)
		if err = k.revokeFeeAllowance(ctx, accountAddr, granteeAddr); err != nil {
			return nil, err
		}
	}

	k.SetUserEntry(ctx, *userEntry)

	event := &types.InitialClaim{
		Claimer:        claimer,
		CampaignId:     strconv.FormatUint(campaignId, 10),
		AddressToClaim: destinationAddress,
		Amount:         claimableAmount.String(),
	}
	if err = ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Debug("initial claim emit event error", "event", event, "error", err.Error())
	}

	return claimableAmount, nil
}

func (k Keeper) Claim(ctx sdk.Context, campaignId uint64, missionId uint64, claimer string) (sdk.Coins, error) {
	k.Logger(ctx).Debug("claim", "claimer", claimer, "campaignId", campaignId, "missionId", missionId)

	campaign, mission, userEntry, claimRecord, err := k.prepareClaimData(ctx, campaignId, missionId, claimer)
	if err != nil {
		return nil, err
	}

	if !claimRecord.IsInitialMissionClaimed() {
		return nil, errors.Wrapf(types.ErrMissionNotCompleted, "initial mission not completed: address %s, campaignId: %d", userEntry.Address, campaignId)
	}

	if mission.MissionType == types.MissionClaim {
		if err = claimRecord.CompleteMission(campaignId, mission.Id); err != nil {
			return nil, err
		}
	}

	claimableAmount, err := claimRecord.ClaimableFromMission(mission)
	if err != nil {
		return nil, err
	}
	err = k.claimMission(ctx, campaign, mission, claimRecord, claimableAmount, campaign.Free)
	if err != nil {
		return nil, err
	}

	k.SetUserEntry(ctx, *userEntry)

	event := &types.Claim{
		Claimer:    claimer,
		CampaignId: strconv.FormatUint(campaignId, 10),
		MissionId:  strconv.FormatUint(missionId, 10),
		Amount:     claimableAmount.String(),
	}
	if err = ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Debug("claim emit event error", "event", event, "error", err.Error())
	}
	return claimableAmount, nil
}

func (k Keeper) CompleteMissionFromHook(ctx sdk.Context, campaignId uint64, missionId uint64, address string) error {
	_, mission, userEntry, claimRecord, err := k.prepareClaimData(ctx, campaignId, missionId, address)
	if err != nil {
		k.Logger(ctx).Debug("complete mission from hook", "err", err.Error())
		return err
	}

	if err = claimRecord.CompleteMission(campaignId, mission.Id); err != nil {
		return err
	}

	k.SetUserEntry(ctx, *userEntry)

	event := &types.CompleteMission{
		CampaignId:  strconv.FormatUint(campaignId, 10),
		MissionId:   strconv.FormatUint(missionId, 10),
		UserAddress: address,
	}
	if err = ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Debug("complete mission from hook event error", "event", event, "error", err.Error())
	}
	return nil
}

func (k Keeper) claimMission(ctx sdk.Context, campaign *types.Campaign, mission *types.Mission, claimRecord *types.ClaimRecord,
	claimableAmount sdk.Coins, free sdk.Dec) error {

	if err := claimRecord.ClaimMission(campaign.Id, mission.Id); err != nil {
		return errors.Wrapf(types.ErrMissionClaiming, err.Error())
	}

	if campaign.CampaignType == types.VestingPoolCampaign {
		if err := k.vestingKeeper.SendReservedToNewVestingAccount(ctx, campaign.Owner, claimRecord.Address, campaign.VestingPoolName,
			claimableAmount.AmountOf(k.vestingKeeper.Denom(ctx)), campaign.Id, free, campaign.LockupPeriod, campaign.VestingPeriod); err != nil {
			return err
		}
	} else {
		start := ctx.BlockTime().Add(campaign.LockupPeriod)
		end := start.Add(campaign.VestingPeriod)
		if _, _, err := k.vestingKeeper.SendToPeriodicContinuousVestingAccountFromModule(ctx, types.ModuleName, claimRecord.Address,
			claimableAmount, free, start.Unix(), end.Unix()); err != nil {
			return errors.Wrapf(c4eerrors.ErrSendCoins, "send to claiming address %s error: "+err.Error(), claimRecord.Address)
		}
	}

	campaign.CampaignCurrentAmount = campaign.CampaignCurrentAmount.Sub(claimableAmount...)
	k.SetCampaign(ctx, *campaign)
	return nil
}

func (k Keeper) validateDestinationAddress(ctx sdk.Context, destAddress string) error {
	destAccAddress, err := sdk.AccAddressFromBech32(destAddress)
	if err != nil {
		return errors.Wrapf(c4eerrors.ErrParsing, "destAddress parsing error: %s", destAddress)
	}

	if k.bankKeeper.BlockedAddr(destAccAddress) {
		return errors.Wrapf(c4eerrors.ErrAccountNotAllowedToReceiveFunds, "new vesting account - account address: %s", destAddress)
	}

	account := k.accountKeeper.GetAccount(ctx, destAccAddress)
	if account == nil {
		return nil
	}

	_, baseAccountOk := account.(*authtypes.BaseAccount)
	_, periodicContinuousVestingAccountOk := account.(*cfevestingtypes.PeriodicContinuousVestingAccount)
	if !baseAccountOk && !periodicContinuousVestingAccountOk {
		return errors.Wrapf(c4eerrors.ErrInvalidAccountType, "account already exists and is not of PeriodicContinuousVestingAccount nor BaseAccount type, got: %T", account)
	}

	return nil
}

func (k Keeper) calculateInitialClaimClaimableAmount(ctx sdk.Context, campaign *types.Campaign, claimRecord *types.ClaimRecord) (sdk.Coins, sdk.Dec) {
	_, weightSum := k.AllMissionForCampaign(ctx, campaign.Id)
	claimableAmount := claimRecord.CalculateInitialClaimClaimableAmount(weightSum)
	free := calculateInitialClaimFree(claimableAmount, campaign)
	return claimableAmount, free
}

func calculateInitialClaimFree(claimableAmount sdk.Coins, campaign *types.Campaign) sdk.Dec {
	minFreeAmount := campaign.Free
	for _, claimableAmountCoin := range claimableAmount {
		free := sdk.NewDecFromInt(campaign.InitialClaimFreeAmount).Quo(sdk.NewDecFromInt(claimableAmountCoin.Amount))
		if minFreeAmount.LT(free) {
			minFreeAmount = free
		}
	}
	maxMinFreeAmount := sdk.NewDec(1)
	if minFreeAmount.GT(sdk.NewDec(1)) {
		minFreeAmount = maxMinFreeAmount
	}
	return minFreeAmount
}
