package v120

import (
	"fmt"
	"time"

	math "cosmossdk.io/math"

	cfeupgradetypes "github.com/chain4energy/c4e-chain/v2/app/upgrades"
	cfevestingtypes "github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
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
	oldAdvisorsPoolName    = "Advisors pool"
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
	sum                                = vcRoundUc4e.Add(earlyBirdRoundUc4e).Add(publicRoundUc4e).Add(strategicReserveShortTermRoundUc4e)
)

func ModifyVestingPoolsState(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	poolsOwnerAddress, err := sdk.AccAddressFromBech32(ValidatorsVestingPoolOwner)
	if err != nil {
		return err
	}
	vestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, poolsOwnerAddress.String())
	if !found {
		ctx.Logger().Info("vesting pools not found", "owner", poolsOwnerAddress.String())
		return nil
	}
	vestingPoolsP := &vestingPools
	var validatorsVestingPools *cfevestingtypes.VestingPool = nil
	for _, vp := range vestingPoolsP.VestingPools {
		if vp.Name == oldValidatorPoolName {
			validatorsVestingPools = vp
		}
		if vp.Name == oldAdvisorsPoolName {
			vp.GenesisPool = true
		}
	}
	if validatorsVestingPools == nil {
		ctx.Logger().Info("validators vesting pool of not found", "owner", poolsOwnerAddress.String())
		return nil
	}

	if validatorsVestingPools.GetCurrentlyLocked().LT(sum) {
		ctx.Logger().Info("validators vesting pool not enough locked to split", "owner", poolsOwnerAddress.String())
		return nil
	}
	if !modifyAndAddVestingTypes(ctx, appKeepers) {
		return nil
	}

	return modifyAndAddVestingPools(ctx, appKeepers, vestingPoolsP, validatorsVestingPools)
}

func modifyAndAddVestingTypes(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) bool {
	vestingType, err := appKeepers.GetC4eVestingKeeper().MustGetVestingType(ctx, oldValidatorTypeName)
	if err != nil {
		ctx.Logger().Info("vesting type not found", "vestingType", oldValidatorTypeName)
		return false
	}
	appKeepers.GetC4eVestingKeeper().RemoveVestingType(ctx, oldValidatorTypeName)
	vestingType.Name = validatorRoundTypeName
	appKeepers.GetC4eVestingKeeper().SetVestingType(ctx, *vestingType)

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
	return true
}

func modifyAndAddVestingPools(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers, vestingPoolsP *cfevestingtypes.AccountVestingPools, validatorsVestingPools *cfevestingtypes.VestingPool) error {

	validatorsVestingPools.Name = validatorRoundPoolName
	validatorsVestingPools.VestingType = validatorRoundTypeName
	validatorsVestingPools.GenesisPool = true
	_, err := splitVestingPool(vestingPoolsP, validatorsVestingPools, vcRoundPoolName, vcRoundTypeName, vcRoundUc4e, 3, 0)
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

	appKeepers.GetC4eVestingKeeper().SetAccountVestingPools(ctx, *vestingPoolsP)

	return nil
}

func splitVestingPool(vestingPools *cfevestingtypes.AccountVestingPools, validatorsVestingPools *cfevestingtypes.VestingPool, poolName string, vestingType string, locked math.Int, addYears int, addMonths int) (*cfevestingtypes.AccountVestingPools, error) {
	if validatorsVestingPools.GetLockedNotReserved().Sub(locked).IsNegative() {
		return nil, fmt.Errorf("not enough coins to send, pool name: %s, currently locked: %s, pool amount: %s", poolName, validatorsVestingPools.GetLockedNotReserved(), locked)
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
		GenesisPool:     true,
	}

	vestingPools.VestingPools = append(vestingPools.VestingPools, &newPool)
	return vestingPools, nil
}

func UpdateVestingAccountTraces(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) {
	traces := appKeepers.GetC4eVestingKeeper().GetAllVestingAccountTrace(ctx)
	genesisAddress := map[string]struct{}{
		"c4e1z5h0squtynr8rhwl0mzqdcd0wgmfyvpqmx3y2r": {},
		"c4e1x6umuffxgcrgqqqdncwn2t8qdnc2muvultxmza": {},
		"c4e1wrhuuwjjmkjx3lxs08ych9ddgdzvujgdr6hnwv": {},
		"c4e12rxujjj4th90t8z30gnre5tv4zmguuqvtn2u02": {},
		"c4e1zvkxuvk8t6wju76pxkp3f4kk447sjm2kdsgvwy": {},
		"c4e13qamrx863pa72ku88d3ykypdh0ar6rjycnpkl2": {},
		"c4e1f57wax48ttw068e6lgag9fse62d4m3e24u0sph": {},
		"c4e1jxlv64qf8rvy8zayl7m2m8a0jzhxkfj9aw96f3": {},
		"c4e1cpnh73765mx3q87lxacqwvwxn4s8ppry458xp4": {},
		"c4e1argfhnzzxjft426tnj4crjsu8lqp0av3x8gjey": {},
		"c4e1w8hdxd6g7vzupll9ynmenjkln9rs4kcq0mdesf": {},
		"c4e12znccp5u8zx9qy4u9gmpxjge9reaxy80qfm295": {},
		"c4e1t45l2pnk5uwj2qqjw4f6rcy6jw5f9lkplmp49e": {},
		"c4e1nmfgexjj3yvvrnc2n7yyahgxsm0vqcm57dqx5f": {},
		"c4e1ej2es5fjztqjcd4pwa0zyvaevtjd2y5wq2vaaq": {},
		"c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8": {},
		"c4e10wjj2qmn4zjg2sdxq9mfyj5v4yukwyhzdtf2zp": {},
		"c4e1zrd0783g8qa5659apw5tpuqmz2ct6j20t4ymx3": {},
		"c4e1y8lndj6jz5z93g4xd05nmwyc3wtn39dfgfx7r7": {},
		"c4e12845qa79cwlvf3jdcnfq2jy2jfmzslcg52lv3g": {},
	}
	fromPool := map[string]struct{}{
		"c4e13e303u43k7mng4927axuhve0plgsyxc4xky63k": {},
		"c4e1twh6302lzcvn7lr3x0fjwfkgryn9ac5c6v2zaj": {},
		"c4e19je7lmu4yzrpzh7gksj3uhku4as8at6lk36qe7": {},
		"c4e1nm50zycnm9yf33rv8n6lpks24usxzahk5usl7e": {},
	}

	for i, trace := range traces {
		_, found := genesisAddress[trace.Address]
		if found {
			traces[i].Genesis = true
			continue
		}
		_, found = fromPool[trace.Address]
		if found {
			traces[i].FromGenesisPool = true
			continue
		}
	}
	for _, trace := range traces {
		appKeepers.GetC4eVestingKeeper().SetVestingAccountTrace(ctx, trace)
	}
}
