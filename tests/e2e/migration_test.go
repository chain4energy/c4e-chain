package e2e

import (
	"os"
	"testing"
	"time"

	"cosmossdk.io/math"
	v120 "github.com/chain4energy/c4e-chain/app/upgrades/v120"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type MigrationSetupSuite struct {
	BaseSetupSuite
}

func TestMigrationSuite(t *testing.T) {
	suite.Run(t, new(MigrationSetupSuite))
}

func (s *MigrationSetupSuite) SetupSuite() {
	bytes, err := os.ReadFile("./resources/mainnet-vestings-migration-app-state.json")
	if err != nil {
		panic(err)
	}
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, false, bytes)
}

func (s *MigrationSetupSuite) TestVestingsMigration() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	vestingTypes := node.QueryVestingTypes()
	s.Equal(7, len(vestingTypes))

	s.ElementsMatch(createVestingTypes(), vestingTypes)

	vestingPools := node.QueryVestingPools(v120.ValidatorsVestingPoolOwner)
	s.Equal(6, len(vestingPools))
	s.ElementsMatch(createVestingPools(), vestingPools)

	endTimeUnix := int64(1758758400)
	startTimeUnix := int64(1695686400)
	newEndTimeUnix := endTimeUnix + 365*24*3600
	newStartTimeUnix := startTimeUnix + 366*24*3600

	acc := node.QueryAccount(v120.Account1)
	s.NotNil(acc)
	vAcc, ok := acc.(*vestexported.ContinuousVestingAccount)
	s.True(ok)
	s.EqualValues(v120.Account1, vAcc.Address)
	s.EqualValues(newStartTimeUnix, vAcc.StartTime)
	s.EqualValues(newEndTimeUnix, vAcc.EndTime)
	s.EqualValues(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(8899990000000))), vAcc.OriginalVesting)
	s.True(vAcc.DelegatedVesting.IsZero())
	s.True(vAcc.DelegatedFree.IsZero())

	acc = node.QueryAccount(v120.Account2)
	s.NotNil(acc)
	vAcc, ok = acc.(*vestexported.ContinuousVestingAccount)
	s.True(ok)
	s.EqualValues(v120.Account2, vAcc.Address)
	s.EqualValues(newStartTimeUnix, vAcc.StartTime)
	s.EqualValues(newEndTimeUnix, vAcc.EndTime)
	s.EqualValues(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(6574990000000))), vAcc.OriginalVesting)
	s.True(vAcc.DelegatedVesting.IsZero())
	s.True(vAcc.DelegatedFree.IsZero())

	acc = node.QueryAccount(v120.Account3)
	s.NotNil(acc)
	vAcc, ok = acc.(*vestexported.ContinuousVestingAccount)
	s.True(ok)
	s.EqualValues(v120.Account3, vAcc.Address)
	s.EqualValues(newStartTimeUnix, vAcc.StartTime)
	s.EqualValues(newEndTimeUnix, vAcc.EndTime)
	s.EqualValues(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(6574990000000))), vAcc.OriginalVesting)
	s.EqualValues(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(100000000))), vAcc.DelegatedVesting)
	s.True(vAcc.DelegatedFree.IsZero())

	acc = node.QueryAccount(v120.Account4)
	s.NotNil(acc)
	vAcc, ok = acc.(*vestexported.ContinuousVestingAccount)
	s.True(ok)
	s.EqualValues(v120.Account4, vAcc.Address)
	s.EqualValues(newStartTimeUnix, vAcc.StartTime)
	s.EqualValues(newEndTimeUnix, vAcc.EndTime)
	s.EqualValues(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(2949990000000))), vAcc.OriginalVesting)
	s.True(vAcc.DelegatedVesting.IsZero())
	s.True(vAcc.DelegatedFree.IsZero())

}

func createVestingTypes() []types.GenesisVestingType {
	vt1 := types.GenesisVestingType{
		Name:              "Advisors",
		LockupPeriod:      365,
		LockupPeriodUnit:  "day",
		VestingPeriod:     730,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.000000000000000000"),
	}
	vt2 := types.GenesisVestingType{
		Name:              "Early-bird round",
		LockupPeriod:      456,
		LockupPeriodUnit:  "day",
		VestingPeriod:     365,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.100000000000000000"),
	}
	vt3 := types.GenesisVestingType{
		Name:              "Public round",
		LockupPeriod:      274,
		LockupPeriodUnit:  "day",
		VestingPeriod:     274,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.150000000000000000"),
	}
	vt4 := types.GenesisVestingType{
		Name:              "Strategic reserve short term round",
		LockupPeriod:      365,
		LockupPeriodUnit:  "day",
		VestingPeriod:     365,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.200000000000000000"),
	}
	vt5 := types.GenesisVestingType{
		Name:              "TestVestingPool",
		LockupPeriod:      30,
		LockupPeriodUnit:  "second",
		VestingPeriod:     30,
		VestingPeriodUnit: "second",
		Free:              sdk.MustNewDecFromStr("0.000000000000000000"),
	}
	vt6 := types.GenesisVestingType{
		Name:              "VC round",
		LockupPeriod:      548,
		LockupPeriodUnit:  "day",
		VestingPeriod:     548,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.050000000000000000"),
	}
	vt7 := types.GenesisVestingType{
		Name:              "Validator round",
		LockupPeriod:      274,
		LockupPeriodUnit:  "day",
		VestingPeriod:     548,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.050000000000000000"),
	}
	return []types.GenesisVestingType{vt1, vt2, vt3, vt4, vt5, vt6, vt7}
}

func createVestingPools() []*types.VestingPoolInfo {
	validatorsLockStart, _ := time.Parse("2006-01-02T15:04:05.000Z", "2022-09-22T14:00:00.000Z")
	validatorsLockEnd, _ := time.Parse("2006-01-02T15:04:05.000Z", "2024-12-26T00:00:00.000Z")

	advisorsLockStart, _ := time.Parse("2006-01-02T15:04:05.000Z", "2022-09-22T14:00:00.000Z")
	advisorsLockEnd, _ := time.Parse("2006-01-02T15:04:05.000Z", "2025-09-25T00:00:00.000Z")

	coin1 := sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(12087500000000))
	advisorsPool := types.VestingPoolInfo{
		Name:            "Advisors pool",
		VestingType:     "Advisors",
		InitiallyLocked: &coin1,
		LockStart:       advisorsLockStart,
		LockEnd:         advisorsLockEnd,
		Withdrawable:    math.ZeroInt().String(),
		SentAmount:      math.NewInt(500000000000).String(),
		CurrentlyLocked: coin1.Amount.Sub(math.NewInt(500000000000)).String(),
	}

	coin2 := sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(8498690000000))
	newValidatorsRoundPool := types.VestingPoolInfo{
		Name:            "Validator round pool",
		VestingType:     "Validator round",
		InitiallyLocked: &coin2,
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockEnd,
		Withdrawable:    math.ZeroInt().String(),
		SentAmount:      math.NewInt(95000000000).String(),
		CurrentlyLocked: coin2.Amount.Sub(math.NewInt(95000000000)).String(),
	}

	coin3 := sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(15000000000000))
	newVcRoundPool := types.VestingPoolInfo{
		Name:            "VC round pool",
		VestingType:     "VC round",
		InitiallyLocked: &coin3,
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(3, 0, 0),
		Withdrawable:    math.ZeroInt().String(),
		SentAmount:      math.ZeroInt().String(),
		CurrentlyLocked: coin3.Amount.String(),
	}

	coin4 := sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(8000000000000))
	newEarlyBirdRoundPool := types.VestingPoolInfo{
		Name:            "Early-bird round pool",
		VestingType:     "Early-bird round",
		InitiallyLocked: &coin4,
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(2, 3, 0),
		Withdrawable:    math.ZeroInt().String(),
		SentAmount:      math.ZeroInt().String(),
		CurrentlyLocked: coin4.Amount.String(),
	}

	coin5 := sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(9000000000000))
	newPublicRoundPool := types.VestingPoolInfo{
		Name:            "Public round pool",
		VestingType:     "Public round",
		InitiallyLocked: &coin5,
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(1, 6, 0),
		Withdrawable:    math.ZeroInt().String(),
		SentAmount:      math.ZeroInt().String(),
		CurrentlyLocked: coin5.Amount.String(),
	}

	coin6 := sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(40000000000000))
	newStrategicRoundPool := types.VestingPoolInfo{
		Name:            "Strategic reserve short term round pool",
		VestingType:     "Strategic reserve short term round",
		InitiallyLocked: &coin6,
		LockStart:       validatorsLockStart,
		LockEnd:         validatorsLockStart.AddDate(2, 0, 0),
		Withdrawable:    math.ZeroInt().String(),
		SentAmount:      math.ZeroInt().String(),
		CurrentlyLocked: coin6.Amount.String(),
	}
	return []*types.VestingPoolInfo{&advisorsPool, &newValidatorsRoundPool, &newVcRoundPool, &newEarlyBirdRoundPool, &newPublicRoundPool, &newStrategicRoundPool}
}
