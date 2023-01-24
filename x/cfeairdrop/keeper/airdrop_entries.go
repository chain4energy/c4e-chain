package keeper

import (
	"fmt"
	errortypes "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	"strconv"
)

func (k Keeper) AddUserAirdropEntries(ctx sdk.Context, owner string, campaignId uint64, airdropEntries []*types.AirdropEntry) error {
	ownerAddress, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Error("add campaign entries owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(errortypes.ErrParsing, sdkerrors.Wrapf(err, "add campaign entries  - owner parsing error: %s", owner).Error())
	}
	campaign, found := k.GetCampaign(
		ctx,
		campaignId,
	)
	if !found {
		k.Logger(ctx).Error("add campaign entries campaign doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(errortypes.ErrParsing, "add campaign entries -  campaign with id %d doesn't exist", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("add campaign entries you are not the owner of this campaign", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add campaign entries - you are not the owner of campaign with id %d", campaignId)
	}
	if campaign.Enabled == false {
		k.Logger(ctx).Error("add campaign entries campaign is disabled", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add campaign entries - campaign %d is disabled", campaignId)
	}
	var usersAirdropEntries []*types.UserAirdropEntries
	entriesAmountSum := sdk.NewCoins()

	for i, airdropEntry := range airdropEntries {
		if airdropEntry.Address == "" {
			k.Logger(ctx).Error("add campaign entries airdrop entry empty address", "airdropEntryIndex", i)
			return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add campaign entries - airdrop entry empty address on index %d", i)
		}
		for _, coin := range airdropEntry.AirdropCoins {
			if coin.Amount.LT(campaign.InitialClaimFreeAmount) {
				k.Logger(ctx).Error("add campaign entries airdrop entry amount < campaign initial claim free amount", "amount", coin.Amount, "initialClaimFreeAmount", campaign.InitialClaimFreeAmount, "airdropEntryIndex", i)
				return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add campaign entries - airdrop entry at index %d amount %s  < campaign initial claim free amount (%s)", i, coin.Amount.String(), campaign.InitialClaimFreeAmount.String())
			}
			entriesAmountSum = entriesAmountSum.Add(coin)
		}

		userAirdropEntries, err := k.addUserAirdropEntry(ctx, campaignId, airdropEntry.Address, airdropEntry.AirdropCoins)
		if err != nil {
			return err
		}
		usersAirdropEntries = append(usersAirdropEntries, userAirdropEntries)
	}

	feesAndEntriesSum := sdk.NewCoins()
	feesSum := sdk.NewCoins()

	bondedDenom := "uc4e"
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		feesSum = sdk.NewCoins(sdk.NewCoin(bondedDenom, campaign.FeegrantAmount.MulRaw(int64(len(airdropEntries)))))
		feesAndEntriesSum.Add(entriesAmountSum...)
		feesAndEntriesSum.Add(feesSum...)
	}

	allBalances := k.bankKeeper.GetAllBalances(ctx, ownerAddress)
	if !allBalances.IsAllGT(feesAndEntriesSum) {
		k.Logger(ctx).Error("add campaign entries airdrop entry owner balance is too small")
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add campaign entries - owner balance is too small")
	}

	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		acc := k.NewModuleAccountSet(ctx, campaignId)
		if err = k.bankKeeper.SendCoins(ctx, ownerAddress, acc.GetAddress(), feesSum); err != nil {
			return err
		}
		if err = k.grantAllFeeAllowance(ctx, acc.GetAddress(), airdropEntries, sdk.NewCoins(types.OneForthC4e)); err != nil {
			return err
		}
	}

	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAddress, types.ModuleName, entriesAmountSum); err != nil {
		return err
	}

	k.IncrementAirdropDistrubitions(ctx, types.AirdropDistrubitions{
		CampaignId:   campaignId,
		AirdropCoins: entriesAmountSum,
	})
	k.IncrementAirdropClaimsLeft(ctx, types.AirdropClaimsLeft{
		CampaignId:   campaignId,
		AirdropCoins: entriesAmountSum,
	})
	for _, userAirdropEntries := range usersAirdropEntries {
		k.SetUserAirdropEntries(ctx, *userAirdropEntries)
	}
	return nil
}

func (k Keeper) addUserAirdropEntry(ctx sdk.Context, campaignId uint64, address string, allCoins sdk.Coins) (*types.UserAirdropEntries, error) {
	userAirdropEntries, found := k.GetUserAirdropEntries(ctx, address)
	if !found {
		userAirdropEntries = types.UserAirdropEntries{Address: address}
	}
	if userAirdropEntries.HasCampaign(campaignId) {
		return nil, sdkerrors.Wrapf(errortypes.ErrAlreadyExists, "campaignId %d already exists for address: %s", campaignId, address)
	}
	userAirdropEntries.AirdropEntries = append(userAirdropEntries.AirdropEntries, &types.AirdropEntry{CampaignId: campaignId, AirdropCoins: allCoins})
	return &userAirdropEntries, nil
}

func (k Keeper) DeleteUserAirdropEntry(ctx sdk.Context, owner string, campaignId uint64, userAddress string) error {
	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Error("delete user airdrop entry owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(errortypes.ErrParsing, sdkerrors.Wrapf(err, "delete user airdrop entry - owner parsing error: %s", owner).Error())
	}
	campaign, found := k.GetCampaign(
		ctx,
		campaignId,
	)
	if !found {
		k.Logger(ctx).Error("delete user airdrop entry campaign doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(errortypes.ErrParsing, "delete user airdrop entry -  campaign with id %d doesn't exist", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("delete user airdrop entry you are not the owner of this campaign", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "delete user airdrop entry - you are not the owner of campaign with id %d", campaignId)
	}
	userAirdropEntries, found := k.GetUserAirdropEntries(
		ctx,
		userAddress,
	)
	if !found {
		k.Logger(ctx).Error("delete user airdrop entry userAirdropEntries doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(errortypes.ErrParsing, "delete user airdrop entry -  campaign id %d userAirdropEntries doesn't exist", campaignId)
	}
	airdropEntryAmount := sdk.NewCoins()
	airdropEntryFound := false
	for i, airdropEntry := range userAirdropEntries.AirdropEntries {
		if airdropEntry.CampaignId == campaignId {
			airdropEntryFound = true
			airdropEntryAmount = airdropEntry.AirdropCoins
			userAirdropEntries.AirdropEntries = append(userAirdropEntries.AirdropEntries[:i], userAirdropEntries.AirdropEntries[i+1:]...)
			break
		}
	}
	if !airdropEntryFound {
		k.Logger(ctx).Error("delete user airdrop entry airdrop entry doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(errortypes.ErrParsing, "delete user airdrop entry -  campaign id %d airdrop entry doesn't exist", campaignId)
	}

	k.SetUserAirdropEntries(ctx, userAirdropEntries)
	k.DecrementAirdropDistrubitions(ctx, campaignId, airdropEntryAmount)

	return nil
}

func (k Keeper) grantAllFeeAllowance(ctx sdk.Context, moduleAddress sdk.AccAddress, airdropEntries []*types.AirdropEntry, grantAmount sdk.Coins) error {
	basicAllowance, err := codectypes.NewAnyWithValue(&feegranttypes.BasicAllowance{
		SpendLimit: grantAmount,
	})
	if err != nil {
		return err // TODO
	}
	allowedMsgAllowance := feegranttypes.AllowedMsgAllowance{
		Allowance:       basicAllowance,
		AllowedMessages: []string{"/chain4energy.c4echain.cfeairdrop.MsgInitialClaim"},
	}
	for _, airdropEntry := range airdropEntries {
		granteeAddress, err := sdk.AccAddressFromBech32(airdropEntry.Address)
		if err != nil {
			return err // TODO
		}
		f, err := k.feeGrantKeeper.GetAllowance(ctx, moduleAddress, granteeAddress)
		if f != nil {
			fmt.Println(err.Error())
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
		}

		if err = k.feeGrantKeeper.GrantAllowance(ctx, moduleAddress, granteeAddress, &allowedMsgAllowance); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) revokeFeeAllowance(ctx sdk.Context, granter sdk.Address, grantee sdk.AccAddress) error {
	keeper, _ := (k.feeGrantKeeper).(feegrantkeeper.Keeper)
	feegrantMsgServer := feegrantkeeper.NewMsgServerImpl(keeper)
	msg := feegranttypes.MsgRevokeAllowance{
		Granter: granter.String(),
		Grantee: grantee.String(),
	}
	_, err := feegrantMsgServer.RevokeAllowance(sdk.WrapSDKContext(ctx), &msg)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) NewModuleAccountSet(ctx sdk.Context, campaignId uint64) *authtypes.ModuleAccount {
	moduleAddressName := "fee-grant-" + strconv.FormatUint(campaignId, 10)
	accountAddr := authtypes.NewModuleAddress(moduleAddressName)
	macc := &authtypes.ModuleAccount{
		BaseAccount: &authtypes.BaseAccount{
			Address: accountAddr.String(),
		},
		Name: moduleAddressName,
	}
	k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccount(ctx, macc))
	return macc
}
