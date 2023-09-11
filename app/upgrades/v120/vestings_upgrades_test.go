package v120_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	v120 "github.com/chain4energy/c4e-chain/v2/app/upgrades/v120"
	testapp "github.com/chain4energy/c4e-chain/v2/testutil/app"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var (
	advisorsType = types.VestingType{
		Name:          "Advisors",
		Free:          sdk.ZeroDec(),
		LockupPeriod:  365 * 24 * time.Hour,
		VestingPeriod: 730 * 24 * time.Hour,
	}

	oldValidatorsType = types.VestingType{
		Name:          "Validators",
		Free:          sdk.ZeroDec(),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}

	newValidatorRoundType = types.VestingType{
		Name:          "Validator round",
		Free:          sdk.ZeroDec(),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}

	newVcRoundType = types.VestingType{
		Name:          "VC round",
		Free:          sdk.MustNewDecFromStr("0.05"),
		LockupPeriod:  548 * 24 * time.Hour,
		VestingPeriod: 548 * 24 * time.Hour,
	}

	newEarlyBirdRoundType = types.VestingType{
		Name:          "Early-bird round",
		Free:          sdk.MustNewDecFromStr("0.10"),
		LockupPeriod:  (365 + 91) * 24 * time.Hour,
		VestingPeriod: 365 * 24 * time.Hour,
	}

	newPublicRoundType = types.VestingType{
		Name:          "Public round",
		Free:          sdk.MustNewDecFromStr("0.15"),
		LockupPeriod:  274 * 24 * time.Hour,
		VestingPeriod: 274 * 24 * time.Hour,
	}

	newStrategicReserveShortTermRoundType = types.VestingType{
		Name:          "Strategic reserve short term round",
		Free:          sdk.MustNewDecFromStr("0.20"),
		LockupPeriod:  365 * 24 * time.Hour,
		VestingPeriod: 365 * 24 * time.Hour,
	}
)

var (
	timeLayout             = "2006-01-02T15:04:05.000Z"
	validatorsLockStart, _ = time.Parse(timeLayout, "2022-09-22T14:00:00.000Z")
	validatorsLockEnd, _   = time.Parse(timeLayout, "2024-12-26T00:00:00.000Z")

	advisorsLockStart, _ = time.Parse(timeLayout, "2022-09-22T14:00:00.000Z")
	advisorsLockEnd, _   = time.Parse(timeLayout, "2025-09-25T00:00:00.000Z")

	oldAdvisorsPool = types.VestingPool{
		Name:            "Advisors pool",
		VestingType:     advisorsType.Name,
		InitiallyLocked: math.NewInt(12087500000000),
		LockStart:       advisorsLockStart,
		LockEnd:         advisorsLockEnd,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.NewInt(500000000000),
		GenesisPool:     false,
	}

	oldValidatorsPool = types.VestingPool{
		Name:            "Validators pool",
		VestingType:     oldValidatorsType.Name,
		InitiallyLocked: math.NewInt(80498690000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockEnd,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.NewInt(95000000000),
		GenesisPool:     false,
	}

	newAdvisorsPool = types.VestingPool{
		Name:            "Advisors pool",
		VestingType:     advisorsType.Name,
		InitiallyLocked: math.NewInt(12087500000000),
		LockStart:       advisorsLockStart,
		LockEnd:         advisorsLockEnd,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.NewInt(500000000000),
		GenesisPool:     true,
	}

	newValidatorsRoundPool = types.VestingPool{
		Name:            "Validator round pool",
		VestingType:     newValidatorRoundType.Name,
		InitiallyLocked: math.NewInt(8498690000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockEnd,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.NewInt(95000000000),
		GenesisPool:     true,
	}

	newVcRoundPool = types.VestingPool{
		Name:            "VC round pool",
		VestingType:     newVcRoundType.Name,
		InitiallyLocked: math.NewInt(15000000000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(3, 0, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
	}

	newEarlyBirdRoundPool = types.VestingPool{
		Name:            "Early-bird round pool",
		VestingType:     newEarlyBirdRoundType.Name,
		InitiallyLocked: math.NewInt(8000000000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(2, 3, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
	}

	newPublicRoundPool = types.VestingPool{
		Name:            "Public round pool",
		VestingType:     newPublicRoundType.Name,
		InitiallyLocked: math.NewInt(9000000000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(1, 6, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
	}

	newStrategicRoundPool = types.VestingPool{
		Name:            "Strategic reserve short term round pool",
		VestingType:     newStrategicReserveShortTermRoundType.Name,
		InitiallyLocked: math.NewInt(40000000000000),
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(2, 0, 0),
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
	}
)

func TestSplitVestingPools(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addVestingPools(testHelper)

	avps, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	sumBefore := math.ZeroInt()
	for _, vp := range avps.VestingPools {
		sumBefore = sumBefore.Add(vp.GetLockedNotReserved())
	}

	err := v120.ModifyVestingPoolsState(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vts := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 6, len(vts.VestingTypes))
	expectedTypes := []*types.VestingType{&advisorsType, &newValidatorRoundType, &newVcRoundType,
		&newEarlyBirdRoundType, &newPublicRoundType, &newStrategicReserveShortTermRoundType}
	require.ElementsMatch(t, expectedTypes, vts.VestingTypes)

	avps, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	sumAfter := math.ZeroInt()
	for _, vp := range avps.VestingPools {
		sumAfter = sumAfter.Add(vp.GetLockedNotReserved())
	}
	require.Equal(t, sumBefore, sumAfter)

	require.True(t, found)
	require.Equal(t, v120.ValidatorsVestingPoolOwner, avps.Owner)
	require.Equal(t, 6, len(avps.VestingPools))

	expectedPools := []*types.VestingPool{&newAdvisorsPool, &newValidatorsRoundPool, &newVcRoundPool,
		&newEarlyBirdRoundPool, &newPublicRoundPool, &newStrategicRoundPool}
	require.ElementsMatch(t, expectedPools, avps.VestingPools)

}

func TestSplitVestingPoolsNoVestingPools(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)

	_, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.False(t, found)

	err := v120.ModifyVestingPoolsState(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vts := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 2, len(vts.VestingTypes))
	expectedTypes := []*types.VestingType{&advisorsType, &oldValidatorsType}
	require.ElementsMatch(t, expectedTypes, vts.VestingTypes)

	_, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.False(t, found)

}

func TestSplitVestingPoolsNoValidatorsVestingPool(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	addAdvisorsVestingPools(testHelper)

	avps, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, 1, len(avps.VestingPools))

	err := v120.ModifyVestingPoolsState(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vts := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 2, len(vts.VestingTypes))
	expectedTypes := []*types.VestingType{&advisorsType, &oldValidatorsType}
	require.ElementsMatch(t, expectedTypes, vts.VestingTypes)

	avps, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, v120.ValidatorsVestingPoolOwner, avps.Owner)
	require.Equal(t, 1, len(avps.VestingPools))

	expectedPools := []*types.VestingPool{&oldAdvisorsPool}
	require.ElementsMatch(t, expectedPools, avps.VestingPools)

}

func TestSplitVestingPoolsNoEnoughValidatorsVestingPoolToSplit(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addVestingTypes(testHelper)
	vestinngPoolNotEnough := addVestingPoolsNotEnoughCoins(testHelper)

	avps, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, 2, len(avps.VestingPools))

	err := v120.ModifyVestingPoolsState(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vts := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 2, len(vts.VestingTypes))
	expectedTypes := []*types.VestingType{&advisorsType, &oldValidatorsType}
	require.ElementsMatch(t, expectedTypes, vts.VestingTypes)

	avps, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, v120.ValidatorsVestingPoolOwner, avps.Owner)
	require.Equal(t, 2, len(avps.VestingPools))

	expectedPools := []*types.VestingPool{&oldAdvisorsPool, &vestinngPoolNotEnough}
	require.ElementsMatch(t, expectedPools, avps.VestingPools)

}

func TestSplitVestingPoolsNoVestingType(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	addAdvisorsVestingTypes(testHelper)
	addVestingPools(testHelper)

	avps, found := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, 2, len(avps.VestingPools))

	err := v120.ModifyVestingPoolsState(testHelper.Context, testHelper.App)
	require.NoError(t, err)

	vts := testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAllVestingTypes(testHelper.Context)
	require.Equal(t, 1, len(vts.VestingTypes))
	expectedTypes := []*types.VestingType{&advisorsType}
	require.ElementsMatch(t, expectedTypes, vts.VestingTypes)

	avps, found = testHelper.C4eVestingUtils.GetC4eVestingKeeper().GetAccountVestingPools(testHelper.Context, v120.ValidatorsVestingPoolOwner)
	require.True(t, found)
	require.Equal(t, v120.ValidatorsVestingPoolOwner, avps.Owner)
	require.Equal(t, 2, len(avps.VestingPools))

	expectedPools := []*types.VestingPool{&oldAdvisorsPool, &oldValidatorsPool}
	require.ElementsMatch(t, expectedPools, avps.VestingPools)

}

func TestUpdateVestingAccountTraces(t *testing.T) {
	addresses := []string{
		"c4e1z5h0squtynr8rhwl0mzqdcd0wgmfyvpqmx3y2r",
		"c4e1x6umuffxgcrgqqqdncwn2t8qdnc2muvultxmza",
		"c4e1wrhuuwjjmkjx3lxs08ych9ddgdzvujgdr6hnwv",
		"c4e12rxujjj4th90t8z30gnre5tv4zmguuqvtn2u02",
		"c4e1zvkxuvk8t6wju76pxkp3f4kk447sjm2kdsgvwy",
		"c4e13qamrx863pa72ku88d3ykypdh0ar6rjycnpkl2",
		"c4e1f57wax48ttw068e6lgag9fse62d4m3e24u0sph",
		"c4e1jxlv64qf8rvy8zayl7m2m8a0jzhxkfj9aw96f3",
		"c4e1cpnh73765mx3q87lxacqwvwxn4s8ppry458xp4",
		"c4e1argfhnzzxjft426tnj4crjsu8lqp0av3x8gjey",
		"c4e1w8hdxd6g7vzupll9ynmenjkln9rs4kcq0mdesf",
		"c4e12znccp5u8zx9qy4u9gmpxjge9reaxy80qfm295",
		"c4e1t45l2pnk5uwj2qqjw4f6rcy6jw5f9lkplmp49e",
		"c4e1nmfgexjj3yvvrnc2n7yyahgxsm0vqcm57dqx5f",
		"c4e1ej2es5fjztqjcd4pwa0zyvaevtjd2y5wq2vaaq",
		"c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8",
		"c4e10wjj2qmn4zjg2sdxq9mfyj5v4yukwyhzdtf2zp",
		"c4e1zrd0783g8qa5659apw5tpuqmz2ct6j20t4ymx3",
		"c4e1y8lndj6jz5z93g4xd05nmwyc3wtn39dfgfx7r7",
		"c4e12845qa79cwlvf3jdcnfq2jy2jfmzslcg52lv3g",
		"c4e13e303u43k7mng4927axuhve0plgsyxc4xky63k",
		"c4e1twh6302lzcvn7lr3x0fjwfkgryn9ac5c6v2zaj",
		"c4e19je7lmu4yzrpzh7gksj3uhku4as8at6lk36qe7",
		"c4e1nm50zycnm9yf33rv8n6lpks24usxzahk5usl7e",
		"c4e5fdsycnm9yf33rvewwfdvs4usdfwer34fwefc",
		"c4ejsdfdfycnm9yf33rv8n6lpks2sdfdsdfssdfsd6",
	}
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	for _, addr := range addresses {
		vat := types.VestingAccountTrace{
			Address:            addr,
			Genesis:            false,
			FromGenesisPool:    false,
			FromGenesisAccount: false,
		}
		testHelper.App.GetC4eVestingKeeper().AppendVestingAccountTrace(testHelper.Context, vat)
	}

	v120.UpdateVestingAccountTraces(testHelper.Context, testHelper.App)
	for i, addr := range addresses[:20] {
		vat := types.VestingAccountTrace{
			Id:                 uint64(i),
			Address:            addr,
			Genesis:            true,
			FromGenesisPool:    false,
			FromGenesisAccount: false,
		}
		newAvt, found := testHelper.App.GetC4eVestingKeeper().GetVestingAccountTrace(testHelper.Context, addr)
		require.True(t, found)
		require.EqualValues(t, vat, newAvt)
	}
	for i, addr := range addresses[20:24] {
		vat := types.VestingAccountTrace{
			Id:                 uint64(i + 20),
			Address:            addr,
			Genesis:            false,
			FromGenesisPool:    true,
			FromGenesisAccount: false,
		}
		newAvt, found := testHelper.App.GetC4eVestingKeeper().GetVestingAccountTrace(testHelper.Context, addr)
		require.True(t, found)
		require.EqualValues(t, vat, newAvt)
	}

	for i, addr := range addresses[24:] {
		vat := types.VestingAccountTrace{
			Id:                 uint64(i + 24),
			Address:            addr,
			Genesis:            false,
			FromGenesisPool:    false,
			FromGenesisAccount: false,
		}
		newAvt, found := testHelper.App.GetC4eVestingKeeper().GetVestingAccountTrace(testHelper.Context, addr)
		require.True(t, found)
		require.EqualValues(t, vat, newAvt)
	}
}

func addAdvisorsVestingTypes(testHelper *testapp.TestHelper) {
	vestingTypes := types.VestingTypes{
		VestingTypes: []*types.VestingType{&advisorsType},
	}
	testHelper.App.CfevestingKeeper.SetVestingTypes(testHelper.Context, vestingTypes)
}

func addVestingTypes(testHelper *testapp.TestHelper) {
	vestingTypes := types.VestingTypes{
		VestingTypes: []*types.VestingType{&advisorsType, &oldValidatorsType},
	}
	testHelper.App.CfevestingKeeper.SetVestingTypes(testHelper.Context, vestingTypes)
}

func addVestingPools(testHelper *testapp.TestHelper) {
	vpools := types.AccountVestingPools{
		Owner:        v120.ValidatorsVestingPoolOwner,
		VestingPools: []*types.VestingPool{&oldAdvisorsPool, &oldValidatorsPool},
	}
	testHelper.App.CfevestingKeeper.SetAccountVestingPools(testHelper.Context, vpools)
}

func addVestingPoolsNotEnoughCoins(testHelper *testapp.TestHelper) types.VestingPool {
	oldValidatorsPoolNotEnough := oldValidatorsPool
	oldValidatorsPoolNotEnough.InitiallyLocked = oldValidatorsPool.Sent.Add(newVcRoundPool.InitiallyLocked).
		Add(newEarlyBirdRoundPool.InitiallyLocked).Add(newPublicRoundPool.InitiallyLocked).
		Add(newStrategicRoundPool.InitiallyLocked).SubRaw(1)
	vpools := types.AccountVestingPools{
		Owner:        v120.ValidatorsVestingPoolOwner,
		VestingPools: []*types.VestingPool{&oldAdvisorsPool, &oldValidatorsPoolNotEnough},
	}
	testHelper.App.CfevestingKeeper.SetAccountVestingPools(testHelper.Context, vpools)
	return oldValidatorsPoolNotEnough
}

func addAdvisorsVestingPools(testHelper *testapp.TestHelper) {
	vpools := types.AccountVestingPools{
		Owner:        v120.ValidatorsVestingPoolOwner,
		VestingPools: []*types.VestingPool{&oldAdvisorsPool},
	}
	testHelper.App.CfevestingKeeper.SetAccountVestingPools(testHelper.Context, vpools)
}
