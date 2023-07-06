package v200

import (
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

const (
	EarlyBirdRoundTypeName = "Early-bird round"
	PublicRoundTypeName    = "Public round"
	FairdropTypeName       = "Fairdrop"
	TeamdropTypeName       = "Teamdrop"
)

func ModifyAndAddVestingTypes(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("modifying and adding new vesting types")

	fairdropVestingType := cfevestingtypes.VestingType{
		Name:          FairdropTypeName,
		Free:          sdk.MustNewDecFromStr("0.01"),
		LockupPeriod:  183 * 24 * time.Hour,
		VestingPeriod: 91 * 24 * time.Hour,
	}
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, fairdropVestingType)

	teamdropVestingType := cfevestingtypes.VestingType{
		Name:          TeamdropTypeName,
		Free:          sdk.MustNewDecFromStr("0.10"),
		LockupPeriod:  730 * 24 * time.Hour,
		VestingPeriod: 730 * 24 * time.Hour,
	}
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, teamdropVestingType)

	_, err := appKeepers.GetC4eVestingKeeper().MustGetVestingType(ctx, EarlyBirdRoundTypeName)
	if err != nil {
		ctx.Logger().Info("vesting type not found", "vestingType", EarlyBirdRoundTypeName)
		return nil
	}

	_, err = appKeepers.GetC4eVestingKeeper().MustGetVestingType(ctx, PublicRoundTypeName)
	if err != nil {
		ctx.Logger().Info("vesting type not found", "vestingType", PublicRoundTypeName)
		return nil
	}

	earlyBirdRoundVestingType := cfevestingtypes.VestingType{
		Name:          EarlyBirdRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  0,
		VestingPeriod: 274 * 24 * time.Hour,
	}
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, earlyBirdRoundVestingType)

	publicRoundVestingType := cfevestingtypes.VestingType{
		Name:          PublicRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.20"),
		LockupPeriod:  0,
		VestingPeriod: 183 * 24 * time.Hour,
	}
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, publicRoundVestingType)

	return nil
}
