package v131

import (
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

const (
	VcRoundTypeName        = "VC round"
	ValidatorRoundTypeName = "Valdiator round"
	PublicRoundTypeName    = "Public round"
	EarlyBirdRoundTypeName = "Early-bird round"
)

func ModifyVestingTypes(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("modifying vesting types")

	_, err := appKeepers.GetC4eVestingKeeper().MustGetVestingType(ctx, VcRoundTypeName)
	if err != nil {
		ctx.Logger().Info("vesting type not found", "vestingType", VcRoundTypeName)
		return nil
	}

	_, err = appKeepers.GetC4eVestingKeeper().MustGetVestingType(ctx, ValidatorRoundTypeName)
	if err != nil {
		ctx.Logger().Info("vesting type not found", "vestingType", ValidatorRoundTypeName)
		return nil
	}

	_, err = appKeepers.GetC4eVestingKeeper().MustGetVestingType(ctx, PublicRoundTypeName)
	if err != nil {
		ctx.Logger().Info("vesting type not found", "vestingType", PublicRoundTypeName)
		return nil
	}

	_, err = appKeepers.GetC4eVestingKeeper().MustGetVestingType(ctx, EarlyBirdRoundTypeName)
	if err != nil {
		ctx.Logger().Info("vesting type not found", "vestingType", EarlyBirdRoundTypeName)
		return nil
	}

	vcRoundPoolType := cfevestingtypes.VestingType{
		Name:          VcRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.08"),
		LockupPeriod:  122 * 24 * time.Hour,
		VestingPeriod: 305 * 24 * time.Hour,
	}
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, vcRoundPoolType)

	validatorRoundType := cfevestingtypes.VestingType{
		Name:          ValidatorRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.08"),
		LockupPeriod:  122 * 24 * time.Hour,
		VestingPeriod: 305 * 24 * time.Hour,
	}
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, validatorRoundType)

	publicRoundType := cfevestingtypes.VestingType{
		Name:          PublicRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.2"),
		LockupPeriod:  30 * 24 * time.Hour,
		VestingPeriod: 152 * 24 * time.Hour,
	}

	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, publicRoundType)

	earlyBirdRoundType := cfevestingtypes.VestingType{
		Name:          EarlyBirdRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  61 * 24 * time.Hour,
		VestingPeriod: 213 * 24 * time.Hour,
	}

	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, earlyBirdRoundType)

	return nil
}
