package keeper

import (
	errortypes "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
)

// SetClaimRecordXX set a specific claimRecordXX in the store from its index
func (k Keeper) SetUserAirdropEntries(ctx sdk.Context, userAirdropEntries types.UserAirdropEntries) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserAirdropEntriesKeyPrefix))
	b := k.cdc.MustMarshal(&userAirdropEntries)
	store.Set(types.UserAirdropEntriesKey(
		userAirdropEntries.Address,
	), b)
}

// GetClaimRecordXX returns a claimRecordXX from its index
func (k Keeper) GetUserAirdropEntries(
	ctx sdk.Context,
	address string,
) (val types.UserAirdropEntries, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserAirdropEntriesKeyPrefix))

	b := store.Get(types.UserAirdropEntriesKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllClaimRecordXX returns all claimRecordXX
func (k Keeper) GetUsersAirdropEntries(ctx sdk.Context) (list []types.UserAirdropEntries) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserAirdropEntriesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserAirdropEntries
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

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
	var usersAirdropEntries []*types.UserAirdropEntries
	sum := sdk.ZeroInt()
	for i, airdropEntry := range airdropEntries {
		if airdropEntry.Amount.LT(types.OneToken) {
			k.Logger(ctx).Error("add campaign entries airdrop entry amount < 1000000 (One token)", "amount", airdropEntry.Amount, "airdropEntryIndex", i)
			return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add campaign entries - airdrop entry at index %d amount %s < 1000000 (One token)", i, airdropEntry.Amount.String())
		}
		userAirdropEntries, err := k.addUserAirdropEntry(ctx, campaignId, airdropEntry.Address, airdropEntry.Amount)
		if err != nil {
			return err
		}
		usersAirdropEntries = append(usersAirdropEntries, userAirdropEntries)
		sum = sum.Add(airdropEntry.Amount)
	}
	coin := sdk.NewCoin(campaign.Denom, sum)
	coins := sdk.NewCoins(coin)
	k.IncrementAirdropDistrubitions(ctx, campaignId, coin)
	k.IncrementAirdropClaimsLeft(ctx, campaignId, coin)
	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAddress, types.ModuleName, coins); err != nil {
		return err
	}
	for _, userAirdropEntries := range usersAirdropEntries {
		k.SetUserAirdropEntries(ctx, *userAirdropEntries)
	}
	return nil
}

func (k Keeper) addUserAirdropEntry(ctx sdk.Context, campaignId uint64, address string, totalAmount sdk.Int) (*types.UserAirdropEntries, error) {
	userAirdropEntries, found := k.GetUserAirdropEntries(ctx, address)
	if !found {
		userAirdropEntries = types.UserAirdropEntries{Address: address}
		// k.grantFeeAllowance(ctx, address)
	}
	if userAirdropEntries.HasCampaign(campaignId) {
		return nil, sdkerrors.Wrapf(errortypes.ErrAlreadyExists, "campaignId %d already exists for address: %s", campaignId, address)
	}
	userAirdropEntries.AirdropEntries = append(userAirdropEntries.AirdropEntries, &types.AirdropEntry{CampaignId: campaignId, Amount: totalAmount})
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
	airdropEntryAmount := sdk.ZeroInt()
	for i, airdropEntry := range userAirdropEntries.AirdropEntries {
		if airdropEntry.CampaignId == campaignId {
			airdropEntryAmount = airdropEntry.Amount
			userAirdropEntries.AirdropEntries = append(userAirdropEntries.AirdropEntries[:i], userAirdropEntries.AirdropEntries[i+1:]...)
			break
		}
	}
	k.SetUserAirdropEntries(ctx, userAirdropEntries)
	coin := sdk.NewCoin(campaign.Denom, airdropEntryAmount)
	k.DecrementAirdropDistrubitions(ctx, campaignId, coin)

	return nil
}

func (k Keeper) grantFeeAllowance(ctx sdk.Context, grantee string) error {
	allowance := feegranttypes.BasicAllowance{}
	address, err := sdk.AccAddressFromBech32(grantee)
	if err != nil {
		return nil // TODO
	}
	modAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	k.feeGrantKeeper.GrantAllowance(ctx, modAcc.GetAddress(), address, &allowance)
	return nil // TODO error handling
}

// func (k Keeper) revokeFeeAllowance(ctx sdk.Context, grantee sdk.AccAddress) error  {
// 	allowance := feegranttypes.BasicAllowance{}

// 	modAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
// 	k.feeGrantKeeper.GrantAllowance(ctx, modAcc.GetAddress(), grantee, &allowance)
// 	return nil // TODO error handling
// }
