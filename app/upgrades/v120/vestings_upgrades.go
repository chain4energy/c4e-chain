package v120

import (
	"fmt"
	"time"

	math "cosmossdk.io/math"

	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ValidatorsVestingPoolOwner = "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54"

	oldValidatorTypeName                   = "Validators"
	validatorRoundTypeName                 = "Validator round"
	vcRoundTypeName                        = "VC round"
	earlyBirdRoundTypeName                 = "Early-bird round"
	publicRoundTypeName                    = "Public round"
	strategicReserveShortTermRoundTypeName = "Strategic reserve short term round"

	oldValidatorPoolName   = "Validators pool"
	validatorRoundPoolName = "Validator round pool"

	vcRoundPoolName                        = "VC round pool"
	earlyBirdRoundPoolName                 = "Early-bird round pool"
	publicRoundPoolName                    = "Public round pool"
	strategicReserveShortTermRoundPoolName = "Strategic reserve short term round pool"
	toUc4e                                 = 1000000
)

var (
	vcRoundUc4e                        = math.NewInt(15000000).MulRaw(toUc4e)
	earlyBirdRoundUc4e                 = math.NewInt(8000000).MulRaw(toUc4e)
	publicRoundUc4e                    = math.NewInt(9000000).MulRaw(toUc4e)
	strategicReserveShortTermRoundUc4e = math.NewInt(40000000).MulRaw(toUc4e)
)

func ModifyVestingPoolsState(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	if err := modifyAndAddVestingTypes(ctx, appKeepers); err != nil {
		return err
	}
	return modifyAndAddVestingPools(ctx, appKeepers)
}

func modifyAndAddVestingTypes(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	vestingType, err := appKeepers.GetC4eVestingKeeper().GetVestingType(ctx, oldValidatorTypeName)
	if err != nil {
		return err
	}
	appKeepers.GetC4eVestingKeeper().RemoveVestingType(ctx, oldValidatorTypeName)
	vestingType.Name = validatorRoundTypeName
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, vestingType)

	vcRoundType := cfevestingtypes.VestingType{
		Name:          vcRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  548 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, vcRoundType)

	earlyBirdRoundType := cfevestingtypes.VestingType{
		Name:          earlyBirdRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.10"),
		LockupPeriod:  (365 + 91) * 24 * time.Hour,
		VestingPeriod: 365 * 24 * time.Hour,
	}
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, earlyBirdRoundType)

	publicRoundType := cfevestingtypes.VestingType{
		Name:          publicRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 274 * 24 * time.Hour,
	}
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, publicRoundType)

	strategicReserveShortTermRoundType := cfevestingtypes.VestingType{
		Name:          strategicReserveShortTermRoundTypeName,
		Free:          sdk.MustNewDecFromStr("0.20"),
		LockupPeriod:  365 * 24 * time.Hour,
		VestingPeriod: 365 * 24 * time.Hour,
	}
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, strategicReserveShortTermRoundType)
	return nil
}

func modifyAndAddVestingPools(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	poolsOwnerAddress, err := sdk.AccAddressFromBech32(ValidatorsVestingPoolOwner)
	if err != nil {
		return err
	}
	vestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, poolsOwnerAddress.String())
	if !found {
		return fmt.Errorf("vesting pools of %s not found", poolsOwnerAddress.String())
	}
	vestingPoolsP := &vestingPools
	var validatorsVestingPools *cfevestingtypes.VestingPool = nil
	for _, vp := range vestingPools.VestingPools {
		if vp.Name == oldValidatorPoolName {
			validatorsVestingPools = vp
		}
	}
	if validatorsVestingPools == nil {
		return fmt.Errorf("validators vesting pool of %s not found", poolsOwnerAddress.String())
	}
	validatorsVestingPools.Name = validatorRoundPoolName
	validatorsVestingPools.VestingType = validatorRoundTypeName

	_, err = splitVestingPool(vestingPoolsP, validatorsVestingPools, vcRoundPoolName, vcRoundTypeName, vcRoundUc4e, 3, 0)
	if err != nil {
		return err
	}

	_, err = splitVestingPool(vestingPoolsP, validatorsVestingPools, earlyBirdRoundPoolName, earlyBirdRoundTypeName, earlyBirdRoundUc4e, 2, 3)
	if err != nil {
		return err
	}

	_, err = splitVestingPool(vestingPoolsP, validatorsVestingPools, publicRoundPoolName, publicRoundTypeName, publicRoundUc4e, 1, 6)
	if err != nil {
		return err
	}

	_, err = splitVestingPool(vestingPoolsP, validatorsVestingPools, strategicReserveShortTermRoundPoolName, strategicReserveShortTermRoundTypeName, strategicReserveShortTermRoundUc4e, 2, 0)
	if err != nil {
		return err
	}

	appKeepers.GetC4eVestingKeeper().SetAccountVestingPools(ctx, vestingPools)

	return nil
}

func splitVestingPool(vestingPools *cfevestingtypes.AccountVestingPools, validatorsVestingPools *cfevestingtypes.VestingPool, poolName string, vestingType string, locked math.Int, addYears int, addMonths int) (*cfevestingtypes.AccountVestingPools, error) {
	if validatorsVestingPools.GetCurrentlyLocked().Sub(locked).IsNegative() {
		return nil, fmt.Errorf("not enough coins to send, pool name: %s, currently locked: %s, pool amount: %s", poolName, validatorsVestingPools.GetCurrentlyLocked(), locked)
	}
	validatorsVestingPools.InitiallyLocked = validatorsVestingPools.InitiallyLocked.Sub(locked)

	newPool := cfevestingtypes.VestingPool{
		Name:            poolName,
		VestingType:     vestingType,
		InitiallyLocked: locked,
		LockStart:       validatorsVestingPools.LockStart,
		LockEnd:         validatorsVestingPools.LockStart.AddDate(addYears, addMonths, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
	}

	vestingPools.VestingPools = append(vestingPools.VestingPools, &newPool)
	return vestingPools, nil
}
