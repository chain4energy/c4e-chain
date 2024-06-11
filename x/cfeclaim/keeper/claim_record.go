package keeper

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k Keeper) AddClaimRecords(ctx sdk.Context, owner string, campaignId uint64, claimRecordEntries []*types.ClaimRecordEntry) error {
	k.Logger(ctx).Debug("add user entries", "owner", owner, "campaignId", campaignId, "claimRecordsLength", len(claimRecordEntries))
	feegrantDenom := k.stakingKeeper.BondDenom(ctx)

	campaign, err := k.ValidateAddClaimRecords(ctx, owner, campaignId, claimRecordEntries)
	if err != nil {
		return err
	}

	usersEntries, amountSum, err := k.validateClaimRecordEntries(ctx, campaign, claimRecordEntries)
	if err != nil {
		return err
	}

	feegrantSum := calculateFeegrantFeesSum(campaign.FeegrantAmount, int64(len(claimRecordEntries)), feegrantDenom)
	ownerAddress, _ := sdk.AccAddressFromBech32(owner)

	if campaign.CampaignType == types.VestingPoolCampaign {
		err = k.addClaimRecordsToVestingPoolCampaign(ctx, campaign, ownerAddress, feegrantSum, amountSum)
	} else {
		err = k.addClaimRecordsToDefaultCampaign(ctx, ownerAddress, feegrantSum, amountSum)
	}
	if err != nil {
		return err
	}

	campaign.IncrementCampaignCurrentAmount(amountSum)
	campaign.IncrementCampaignTotalAmount(amountSum)
	k.SetCampaign(ctx, *campaign)
	k.Logger(ctx).Debug("increment campaign amounts", "campaignId", campaignId, "amount", amountSum.String())
	err = k.setupAndSendFeegrant(ctx, ownerAddress, campaign, feegrantSum, claimRecordEntries, feegrantDenom)
	if err != nil {
		return err
	}

	for _, userEntry := range usersEntries {
		userEntryAddress, err := sdk.AccAddressFromBech32(userEntry.Address)
		if err != nil {
			return errors.Wrapf(c4eerrors.ErrParam, "claim record entry user entry address parsing error (%s)", err)
		}
		if k.accountKeeper.GetAccount(ctx, userEntryAddress) == nil {
			acc := k.accountKeeper.NewAccountWithAddress(ctx, userEntryAddress)
			k.accountKeeper.SetAccount(ctx, acc)
		}
		k.SetUserEntry(ctx, *userEntry)
	}

	event := &types.EventAddClaimRecords{
		Owner:                   owner,
		CampaignId:              campaignId,
		ClaimRecordsTotalAmount: amountSum.String(),
		ClaimRecordsNumber:      int64(len(claimRecordEntries)),
	}

	if err = ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Debug("add claim records emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) addClaimRecordsToVestingPoolCampaign(ctx sdk.Context, campaign *types.Campaign, ownerAddress sdk.AccAddress,
	feegrantFeesSum sdk.Coins, amount sdk.Coins) error {

	vestingDenom := k.vestingKeeper.Denom(ctx)
	spendableCoins := k.bankKeeper.SpendableCoins(ctx, ownerAddress)
	if !feegrantFeesSum.IsAllLTE(spendableCoins) {
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "owner balance is too small (%s < %s)", spendableCoins, feegrantFeesSum)
	}
	return k.vestingKeeper.AddVestingPoolReservation(ctx, ownerAddress.String(), campaign.VestingPoolName, campaign.Id, amount.AmountOf(vestingDenom))
}

func (k Keeper) addClaimRecordsToDefaultCampaign(ctx sdk.Context, ownerAddress sdk.AccAddress, feegrantFeesSum sdk.Coins, amount sdk.Coins) error {
	feesAndClaimRecordsAmountSum := amount.Add(feegrantFeesSum...)
	spendableCoins := k.bankKeeper.SpendableCoins(ctx, ownerAddress)
	if !feesAndClaimRecordsAmountSum.IsAllLTE(spendableCoins) {
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "owner balance is too small (%s < %s)", spendableCoins, feesAndClaimRecordsAmountSum)
	}
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAddress, types.ModuleName, amount)
}

func (k Keeper) DeleteClaimRecord(ctx sdk.Context, owner string, campaignId uint64, userAddress string) error {
	k.Logger(ctx).Debug("delete claim record", "owner", owner, "campaignId", campaignId, "userAddress", userAddress)
	campaign, err := k.MustGetCampaign(ctx, campaignId)
	if err != nil {
		return err
	}

	userEntry, claimRecord, err := k.validateDeleteClaimRecord(ctx, owner, campaign, userAddress)
	if err != nil {
		return err
	}

	amount, err := k.calculateDeleteClaimRecordAmount(ctx, campaign, claimRecord)
	if err != nil {
		return err
	}
	if err = k.sendCampaignCurrentAmountToOwner(ctx, campaign, amount); err != nil {
		return err
	}

	err = k.deleteClaimRecordSendFeegrant(ctx, campaign, userAddress)
	if err != nil {
		k.Logger(ctx).Debug("delete claim record send feegrant err", "err", err.Error())
	}

	userEntry.DeleteClaimRecord(campaignId)
	k.SetUserEntry(ctx, *userEntry)

	campaign.DecrementCampaignCurrentAmount(amount)
	campaign.DecrementCampaignTotalAmount(amount)
	k.SetCampaign(ctx, *campaign)
	k.Logger(ctx).Debug("delete claim record decrement campaign amounts", "campaignId", campaignId, "amount", amount)

	event := &types.EventDeleteClaimRecord{
		Owner:             owner,
		CampaignId:        campaignId,
		UserAddress:       userAddress,
		ClaimRecordAmount: amount.String(),
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Debug("delete claim record emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) ValidateAddClaimRecords(ctx sdk.Context, owner string, campaignId uint64, claimRecordEntries []*types.ClaimRecordEntry) (*types.Campaign, error) {
	campaign, err := k.MustGetCampaign(ctx, campaignId)
	if err != nil {
		return nil, err
	}
	if err = campaign.ValidateOwner(owner); err != nil {
		return nil, err
	}
	if err = campaign.ValidateNotEnded(ctx.BlockTime()); err != nil {
		return nil, err
	}
	if campaign.CampaignType == types.VestingPoolCampaign {
		if err = types.ValidateVestingPoolCampaignClaimRecordEntries(claimRecordEntries, k.vestingKeeper.Denom(ctx)); err != nil {
			return nil, err
		}
	} else {
		if err = types.ValidateClaimRecordEntries(claimRecordEntries); err != nil {
			return nil, err
		}
	}

	return campaign, nil
}

func (k Keeper) validateClaimRecordEntries(ctx sdk.Context, campaign *types.Campaign, claimRecordEntries []*types.ClaimRecordEntry) (usersEntries []*types.UserEntry, amountSum sdk.Coins, err error) {
	for i, claimRecord := range claimRecordEntries {
		amountSum = amountSum.Add(claimRecord.Amount...)

		userEntry, err := k.addClaimRecordToUserEntry(ctx, campaign.Id, claimRecord.UserEntryAddress, claimRecord.Amount)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "claim record entry index %d", i)
		}
		usersEntries = append(usersEntries, userEntry)
	}
	return
}

func (k Keeper) addClaimRecordToUserEntry(ctx sdk.Context, campaignId uint64, address string, allCoins sdk.Coins) (*types.UserEntry, error) {
	userEntry, found := k.GetUserEntry(ctx, address)
	if !found {
		userEntry = types.UserEntry{Address: address}
	}
	if userEntry.HasCampaign(campaignId) {
		return nil, errors.Wrapf(c4eerrors.ErrAlreadyExists, "campaignId %d already exists for address: %s", campaignId, address)
	}
	userEntry.ClaimRecords = append(userEntry.ClaimRecords, &types.ClaimRecord{CampaignId: campaignId, Amount: allCoins})
	return &userEntry, nil
}

func (k Keeper) SetupNewFeegrantAccount(ctx sdk.Context, campaignId uint64) sdk.AccAddress {
	moduleName, accAddress := CreateFeegrantAccountAddress(campaignId)
	account := k.accountKeeper.GetAccount(ctx, accAddress)
	if account != nil {
		return accAddress
	}
	baseAccount := authtypes.NewBaseAccountWithAddress(accAddress)
	moduleAccount := authtypes.NewModuleAccount(baseAccount, moduleName)
	k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccount(ctx, moduleAccount))
	return accAddress
}

func (k Keeper) validateDeleteClaimRecord(ctx sdk.Context, owner string, campaign *types.Campaign, userAddress string) (*types.UserEntry, *types.ClaimRecord, error) {
	if err := campaign.ValidateRemovableClaimRecords(); err != nil {
		return nil, nil, err
	}

	if err := campaign.ValidateOwner(owner); err != nil {
		return nil, nil, err
	}

	userEntry, err := k.MustGetUserEntry(ctx, userAddress)
	if err != nil {
		return nil, nil, err
	}

	claimRecord, err := userEntry.MustGetClaimRecord(campaign.Id)
	if err != nil {
		return nil, nil, err
	}

	return &userEntry, claimRecord, nil
}

func (k Keeper) calculateDeleteClaimRecordAmount(ctx sdk.Context, campaign *types.Campaign, claimRecord *types.ClaimRecord) (sdk.Coins, error) {
	amount := claimRecord.Amount
	for _, claimedMissionId := range claimRecord.ClaimedMissions {
		mission, err := k.MustGetMission(ctx, campaign.Id, claimedMissionId)
		if err != nil {
			return nil, err
		}
		if mission.MissionType == types.MissionInitialClaim {
			initiialClaimClaimableAmount, _ := k.calculateInitialClaimClaimableAmount(ctx, campaign, claimRecord)
			amount = amount.Sub(initiialClaimClaimableAmount...)
			continue
		}
		claimableFromMission, err := claimRecord.ClaimableFromMission(mission)
		if err != nil {
			return nil, err
		}
		amount = amount.Sub(claimableFromMission...)
	}

	return amount, nil
}
