package keeper

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types2 "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) InitialClaim(ctx sdk.Context, owner string, campaignId uint64, additionalAddress string) error {
	ak := k.accountKeeper
	bk := k.bankKeeper
	var addressToClaim = owner
	if additionalAddress != "" {
		accAddress, err := sdk.AccAddressFromBech32(additionalAddress)
		if err != nil {
			k.Logger(ctx).Error("add mission to airdrop campaign owner parsing error", "owner", owner, "error", err.Error())
			return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "add mission to airdrop campaign - owner parsing error: %s", owner).Error())
		}
		if bk.BlockedAddr(accAddress) {
			k.Logger(ctx).Error("new vesting account is not allowed to receive funds error", "address", additionalAddress)
			return sdkerrors.Wrapf(c4eerrors.ErrAccountNotAllowedToReceiveFunds, "new vesting account - account address: %s", additionalAddress)
		}

		account := ak.GetAccount(ctx, accAddress)
		_, ok := account.(*types2.BaseVestingAccount)
		if ok {
			k.Logger(ctx).Error("new vesting account is not allowed to receive funds error", "address", additionalAddress)
			return sdkerrors.Wrapf(c4eerrors.ErrAccountNotAllowedToReceiveFunds, "new vesting account - account address: %s", additionalAddress)

		}
		addressToClaim = additionalAddress
	}

	campaign, mission, userAirdropEntries, err := k.missionFirstStep(ctx, "claim initial mission", campaignId, types.InitialMissionId, addressToClaim, false)
	if err != nil {
		return err
	}
	if err = mission.IsEnabled(ctx.BlockTime()); err != nil {
		k.Logger(ctx).Error("claim initial mission - mission disabled", "campaignId", campaignId, "missionId", mission.Id, "err", err)
		return sdkerrors.Wrapf(err, "mission disabled - campaignId %d, missionId %d", campaignId, mission.Id)
	}
	userAirdropEntries.ClaimAddress = addressToClaim
	userAirdropEntries, err = k.completeMission(ctx, true, mission, userAirdropEntries)
	if err != nil {
		return err
	}
	userAirdropEntries, err = k.claimMission(ctx, true, campaign, mission, userAirdropEntries)
	if err != nil {
		return err
	}
	k.SetUserAirdropEntries(ctx, *userAirdropEntries)
	return nil
}

func (k Keeper) Claim(ctx sdk.Context, campaignId uint64, missionId uint64, claimer string) error {
	campaign, mission, userAirdropEntries, err := k.missionFirstStep(ctx, "claim mission", campaignId, missionId, claimer, false)
	if err != nil {
		return err
	}
	if err = mission.IsEnabled(ctx.BlockTime()); err != nil {
		k.Logger(ctx).Error("claim mission - mission disabled", "campaignId", campaignId, "missionId", missionId, "err", err)
		return sdkerrors.Wrapf(err, "mission disabled - campaignId %d, missionId %d", campaignId, missionId)
	}
	userAirdropEntries, err = k.claimMission(ctx, false, campaign, mission, userAirdropEntries)
	if err != nil {
		return err
	}

	k.SetUserAirdropEntries(ctx, *userAirdropEntries)
	return nil
}
