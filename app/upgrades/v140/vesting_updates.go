package v140

import (
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

const (
	VcRoundTypeName        = "VC round"
	ValidatorRoundTypeName = "Valdiator round"
)

func ModifyVestingTypes(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("modifying and adding new vesting types")

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

	return nil
}
