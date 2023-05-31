package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	"strconv"
)

func CreateFeegrantAccountAddress(campaignId uint64) (string, sdk.AccAddress) {
	moduleAddressName := types.ModuleName + "-fee-grant-" + strconv.FormatUint(campaignId, 10)
	return moduleAddressName, authtypes.NewModuleAddress(moduleAddressName)
}

func calculateFeegrantFeesSum(feegrantAmount math.Int, claimRecordEntriesNumber int64, feegrantDenom string) sdk.Coins {
	if feegrantAmount.GT(math.ZeroInt()) {
		return sdk.NewCoins(sdk.NewCoin(feegrantDenom, feegrantAmount.MulRaw(claimRecordEntriesNumber)))
	}
	return nil
}

func (k Keeper) setupAndSendFeegrant(ctx sdk.Context, ownerAcc sdk.AccAddress, campaign *types.Campaign, feegrantFeesSum sdk.Coins, claimRecordEntries []*types.ClaimRecordEntry, feegrantDenom string) error {
	if campaign.FeegrantAmount.GT(math.ZeroInt()) {
		acc := k.SetupNewFeegrantAccount(ctx, campaign.Id)

		if err := k.bankKeeper.SendCoins(ctx, ownerAcc, acc, feegrantFeesSum); err != nil {
			return err
		}
		if err := k.grantFeeAllowanceToAllClaimRecords(ctx, acc, claimRecordEntries, sdk.NewCoins(sdk.NewCoin(feegrantDenom, campaign.FeegrantAmount))); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) grantFeeAllowanceToAllClaimRecords(ctx sdk.Context, moduleAddress sdk.AccAddress, claimRecordEntries []*types.ClaimRecordEntry, grantAmount sdk.Coins) error {
	basicAllowance, err := codectypes.NewAnyWithValue(&feegranttypes.BasicAllowance{
		SpendLimit: grantAmount,
	})
	if err != nil {
		return err
	}

	allowedMsgAllowance := feegranttypes.AllowedMsgAllowance{
		Allowance:       basicAllowance,
		AllowedMessages: []string{sdk.MsgTypeURL(&types.MsgInitialClaim{})},
	}

	for _, claimRecord := range claimRecordEntries {
		granteeAddress, _ := sdk.AccAddressFromBech32(claimRecord.UserEntryAddress)
		existingFeeAllowance, _ := k.feeGrantKeeper.GetAllowance(ctx, moduleAddress, granteeAddress)
		if existingFeeAllowance == nil {
			if err = k.feeGrantKeeper.GrantAllowance(ctx, moduleAddress, granteeAddress, &allowedMsgAllowance); err != nil {
				return err
			}
		}
	}

	return nil
}

func (k Keeper) closeCampaignSendFeegrant(ctx sdk.Context, campaign *types.Campaign) error {
	if !campaign.FeegrantAmount.IsPositive() {
		return nil
	}
	_, feegrantAccountAddress := CreateFeegrantAccountAddress(campaign.Id)
	feegrantTotalAmount := k.bankKeeper.GetAllBalances(ctx, feegrantAccountAddress)
	ownerAddress, _ := sdk.AccAddressFromBech32(campaign.Owner)

	return k.bankKeeper.SendCoins(ctx, feegrantAccountAddress, ownerAddress, feegrantTotalAmount)
}

func (k Keeper) deleteClaimRecordSendFeegrant(ctx sdk.Context, campaign *types.Campaign, userEntryAddress string) error {
	if !campaign.FeegrantAmount.IsPositive() {
		return nil
	}
	feegrantAccountAddress, granteeAddress, amountLeft, err := k.getFeegrantLeftAmount(ctx, campaign.Id, userEntryAddress)
	if err != nil {
		return err
	}

	k.revokeFeeAllowance(ctx, feegrantAccountAddress, granteeAddress)
	campaignOwnerAccAddress, err := sdk.AccAddressFromBech32(campaign.Owner)
	if err != nil {
		return err
	}
	return k.bankKeeper.SendCoins(ctx, feegrantAccountAddress, campaignOwnerAccAddress, amountLeft)
}

func (k Keeper) getFeegrantLeftAmount(ctx sdk.Context, campaignId uint64, userEntryAddress string) (sdk.AccAddress, sdk.AccAddress, sdk.Coins, error) {
	_, granterAddress := CreateFeegrantAccountAddress(campaignId)
	granteeAddress, _ := sdk.AccAddressFromBech32(userEntryAddress)

	allowance, err := k.feeGrantKeeper.GetAllowance(ctx, granterAddress, granteeAddress)
	if err != nil {
		return nil, nil, nil, err
	}
	x, ok := allowance.(*feegranttypes.AllowedMsgAllowance)
	if !ok {
		return nil, nil, nil, errors.Wrap(sdkerrors.ErrInvalidType, "cannot get AllowedMsgAllowance")
	}
	for _, msg := range x.AllowedMessages {
		if msg == sdk.MsgTypeURL(&types.MsgInitialClaim{}) {
			basicAllowance := x.Allowance.GetCachedValue().(*feegranttypes.BasicAllowance)
			return granterAddress, granteeAddress, basicAllowance.SpendLimit, nil
		}
	}
	return granterAddress, granteeAddress, nil, errors.Wrap(sdkerrors.ErrInvalidType, "cannot get feegrant left amount")
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
