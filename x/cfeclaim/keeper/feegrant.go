package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	"strconv"
)

func validateFeegrantAmount(feeGrantAmount *math.Int) (math.Int, error) {
	if feeGrantAmount == nil {
		return sdk.ZeroInt(), nil
	}

	if feeGrantAmount.IsNil() {
		return sdk.ZeroInt(), nil
	}

	if feeGrantAmount.IsNegative() {
		return sdk.ZeroInt(), errors.Wrapf(c4eerrors.ErrParam, "feegrant amount (%s) cannot be negative", feeGrantAmount.String())
	}

	return *feeGrantAmount, nil
}

func CreateFeegrantAccountAddress(campaignId uint64) (string, sdk.AccAddress) {
	moduleAddressName := types.ModuleName + "-fee-grant-" + strconv.FormatUint(campaignId, 10)
	return moduleAddressName, authtypes.NewModuleAddress(moduleAddressName)
}

func calculateFeegrantFeesSum(feegrantAmount math.Int, claimRecordsNumber int64, feegrantDenom string) (feesSum sdk.Coins) {
	if feegrantAmount.GT(sdk.ZeroInt()) {
		feesSum = feesSum.Add(sdk.NewCoin(feegrantDenom, feegrantAmount.MulRaw(claimRecordsNumber)))
	}
	return
}

func (k Keeper) setupAndSendFeegrant(ctx sdk.Context, ownerAcc sdk.AccAddress, campaign *types.Campaign, feegrantFeesSum sdk.Coins, claimRecords []*types.ClaimRecord, feegrantDenom string) error {
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		acc := k.NewModuleAccountSet(ctx, campaign.Id)
		if err := k.bankKeeper.SendCoins(ctx, ownerAcc, acc.GetAddress(), feegrantFeesSum); err != nil {
			return err
		}
		if err := k.grantFeeAllowanceToAllClaimRecords(ctx, acc.GetAddress(), claimRecords, sdk.NewCoins(sdk.NewCoin(feegrantDenom, campaign.FeegrantAmount))); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) grantFeeAllowanceToAllClaimRecords(ctx sdk.Context, moduleAddress sdk.AccAddress, claimEntries []*types.ClaimRecord, grantAmount sdk.Coins) error {
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

	for _, claimRecord := range claimEntries {
		granteeAddress, _ := sdk.AccAddressFromBech32(claimRecord.Address)
		existingFeeAllowance, _ := k.feeGrantKeeper.GetAllowance(ctx, moduleAddress, granteeAddress)
		if existingFeeAllowance == nil {
			if err = k.feeGrantKeeper.GrantAllowance(ctx, moduleAddress, granteeAddress, &allowedMsgAllowance); err != nil {
				return err
			}
		}
	}

	return nil
}

func (k Keeper) closeCampaignSendFeegrant(ctx sdk.Context, CloseAction types.CloseAction, campaign *types.Campaign) error {
	if !campaign.FeegrantAmount.IsPositive() {
		return nil
	}
	_, feegrantAccountAddress := CreateFeegrantAccountAddress(campaign.Id)
	feegrantTotalAmount := k.bankKeeper.GetAllBalances(ctx, feegrantAccountAddress)

	return k.feegrantCloseActionSwitch(ctx, campaign.Owner, CloseAction, feegrantTotalAmount, feegrantAccountAddress)
}

func (k Keeper) deleteClaimRecordSendFeegrant(ctx sdk.Context, CloseAction types.CloseAction, campaign *types.Campaign, userEntryAddress string) error {
	if !campaign.FeegrantAmount.IsPositive() {
		return nil
	}
	feegrantAccountAddress, granteeAddress, amountLeft, err := k.getFeegrantLeftAmount(ctx, campaign.Id, userEntryAddress)
	if err != nil {
		return err
	}

	k.revokeFeeAllowance(ctx, feegrantAccountAddress, granteeAddress)

	return k.feegrantCloseActionSwitch(ctx, campaign.Owner, CloseAction, amountLeft, feegrantAccountAddress)
}

func (k Keeper) feegrantCloseActionSwitch(ctx sdk.Context, owner string, CloseAction types.CloseAction, amountLeft sdk.Coins, granterAddress sdk.AccAddress) error {
	switch CloseAction {
	case types.CloseSendToCommunityPool:
		if err := k.distributionKeeper.FundCommunityPool(ctx, amountLeft, granterAddress); err != nil {
			return err
		}
	case types.CloseBurn:
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, granterAddress, types.ModuleName, amountLeft); err != nil {
			return err
		}
		if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, amountLeft); err != nil {
			return err
		}
	case types.CloseSendToOwner:
		ownerAddress, _ := sdk.AccAddressFromBech32(owner)
		if err := k.bankKeeper.SendCoins(ctx, granterAddress, ownerAddress, amountLeft); err != nil {
			return err
		}
	default:
		return errors.Wrap(sdkerrors.ErrInvalidType, "wrong close action type")
	}
	return nil
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
			basicAllowance := x.Allowance.GetCachedValue().(feegranttypes.BasicAllowance)
			return granterAddress, granteeAddress, basicAllowance.SpendLimit, nil
		}
	}
	return granterAddress, granteeAddress, nil, errors.Wrap(sdkerrors.ErrInvalidType, "cannot get feegrant left amount")
}
