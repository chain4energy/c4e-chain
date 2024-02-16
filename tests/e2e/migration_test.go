package e2e

import (
	"cosmossdk.io/math"
	v131 "github.com/chain4energy/c4e-chain/app/upgrades/v131"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
	"time"
)

type MainnetMigrationSetupSuite struct {
	BaseSetupSuite
}

func TestMainnetMigrationSuite(t *testing.T) {
	suite.Run(t, new(MainnetMigrationSetupSuite))
}

func (s *MainnetMigrationSetupSuite) SetupSuite() {
	bytes, err := os.ReadFile("./resources/mainnet-v1.3.1-migration-app-state.json")
	if err != nil {
		panic(err)
	}
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, false, false, bytes)
}

func (s *MainnetMigrationSetupSuite) TestMainnetMigration() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	// verify vesting types
	vestingTypes := node.QueryVestingTypes()
	s.Equal(7, len(vestingTypes))

	s.ElementsMatch(createMainnetVestingTypes(), vestingTypes)

	// verify community pool
	s.Equal(node.QueryCommunityPool().AmountOf(testenv.DefaultTestDenom), sdk.NewDec(40_000_000_000_000))

	// verify strategic reserve account
	acc := node.QueryAccount(v131.StrategicReserveAccount)
	s.NotNil(acc)
	vAcc, ok := acc.(*vestexported.ContinuousVestingAccount)
	s.True(ok)
	s.EqualValues(v131.StrategicReserveAccount, vAcc.Address)
	s.EqualValues(1727222400, vAcc.StartTime)
	s.EqualValues(1821830400, vAcc.EndTime)
	s.EqualValues(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(50_000_000_000_000))), vAcc.OriginalVesting)

	// verify strategic reserve short term pool
	vestingPools := node.QueryVestingPoolsInfo(v131.StrategicReservceShortTermPoolAccount)

	s.Equal(3, len(vestingPools))
	s.ElementsMatch(s.createVestingPools(), vestingPools)
}

type NonMainnetMigrationSetupSuite struct {
	BaseSetupSuite
}

func TestNonMainnetMigrationSuite(t *testing.T) {
	suite.Run(t, new(NonMainnetMigrationSetupSuite))
}

func (s *NonMainnetMigrationSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuiteWithUpgradeAppState(true, false, false, nil)
}

func (s *NonMainnetMigrationSetupSuite) TestNonMainnetVestingsMigration() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	vestingTypes := node.QueryVestingTypes()
	s.Equal(2, len(vestingTypes))

	s.ElementsMatch(createNonMainnetVestingTypes(), vestingTypes)

	node.QueryVestingPoolsNotFound(v131.StrategicReservceShortTermPoolAccount)

	node.QueryAccountNotFound(v131.StrategicReserveAccount)
}

func (s *MainnetMigrationSetupSuite) createVestingPools() []*types.VestingPoolInfo {
	earlyBirdLockStart, err := time.Parse(time.RFC3339, "2022-03-30T00:00:00Z")
	s.NoError(err)
	earlyBirdLockEnd, err := time.Parse(time.RFC3339, "2025-03-30T00:00:00Z")
	s.NoError(err)
	earlyBirdCoin := sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(8000000000000))
	earlyBirdPool := types.VestingPoolInfo{
		Name:            "Early-bird round pool",
		VestingType:     "Early-bird round",
		InitiallyLocked: &earlyBirdCoin,
		LockStart:       earlyBirdLockStart,
		LockEnd:         earlyBirdLockEnd,
		Withdrawable:    math.ZeroInt().String(),
		SentAmount:      math.ZeroInt().String(),
		Reservations:    []*types.VestingPoolReservation{},
		CurrentlyLocked: earlyBirdCoin.Amount.String(),
	}

	publicRoundLockStart, err := time.Parse(time.RFC3339, "2022-03-30T00:00:00Z")
	s.NoError(err)
	publicRoundLockEnd, err := time.Parse(time.RFC3339, "2024-03-30T00:00:00Z")
	s.NoError(err)
	publicRoundCoin := sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(9000000000000))
	publicRoundPool := types.VestingPoolInfo{
		Name:            "Public round pool",
		VestingType:     "Public round",
		InitiallyLocked: &publicRoundCoin,
		LockStart:       publicRoundLockStart,
		LockEnd:         publicRoundLockEnd,
		Withdrawable:    math.ZeroInt().String(),
		SentAmount:      math.ZeroInt().String(),
		Reservations:    []*types.VestingPoolReservation{},
		CurrentlyLocked: publicRoundCoin.Amount.String(),
	}

	strategicReserveLockStart, err := time.Parse(time.RFC3339, "2022-09-22T14:00:00Z")
	strategicReserveLockEnd, err := time.Parse(time.RFC3339, "2024-09-22T14:00:00Z")
	strategicReserveCoin := sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(20000000000000))
	strategicReservePool := types.VestingPoolInfo{
		Name:            "Strategic reserve short term round pool",
		VestingType:     "Strategic reserve short term round",
		InitiallyLocked: &strategicReserveCoin,
		LockStart:       strategicReserveLockStart,
		LockEnd:         strategicReserveLockEnd,
		Withdrawable:    math.ZeroInt().String(),
		SentAmount:      math.NewInt(250000000000).String(),
		Reservations:    []*types.VestingPoolReservation{},
		CurrentlyLocked: strategicReserveCoin.Amount.Sub(math.NewInt(250000000000)).String(),
	}

	return []*types.VestingPoolInfo{&earlyBirdPool, &publicRoundPool, &strategicReservePool}
}

func createNonMainnetVestingTypes() []types.GenesisVestingType {
	vt1 := types.GenesisVestingType{
		Name:              "Advisors",
		LockupPeriod:      365,
		LockupPeriodUnit:  "day",
		VestingPeriod:     730,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.000000000000000000"),
	}
	vt2 := types.GenesisVestingType{
		Name:              "TestVestingPool",
		LockupPeriod:      30,
		LockupPeriodUnit:  "second",
		VestingPeriodUnit: "second",
		VestingPeriod:     30,
		Free:              sdk.MustNewDecFromStr("0.000000000000000000"),
	}

	return []types.GenesisVestingType{vt1, vt2}
}
func createMainnetVestingTypes() []types.GenesisVestingType {
	vt1 := types.GenesisVestingType{
		Name:              "Validator round",
		LockupPeriod:      122,
		LockupPeriodUnit:  "day",
		VestingPeriod:     305,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.080000000000000000"),
	}

	vt2 := types.GenesisVestingType{
		Name:              "Strategic reserve short term round",
		LockupPeriod:      365,
		LockupPeriodUnit:  "day",
		VestingPeriod:     365,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.200000000000000000"),
	}

	vt3 := types.GenesisVestingType{
		Name:              "VC round",
		LockupPeriod:      122,
		LockupPeriodUnit:  "day",
		VestingPeriod:     305,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.080000000000000000"),
	}

	vt4 := types.GenesisVestingType{
		Name:              "Advisors",
		LockupPeriod:      365,
		LockupPeriodUnit:  "day",
		VestingPeriod:     730,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.000000000000000000"),
	}
	vt5 := types.GenesisVestingType{
		Name:              "TestVestingPool",
		LockupPeriod:      30,
		LockupPeriodUnit:  "second",
		VestingPeriodUnit: "second",
		VestingPeriod:     30,
		Free:              sdk.MustNewDecFromStr("0.000000000000000000"),
	}

	vt6 := types.GenesisVestingType{
		Name:              "Early-bird round",
		LockupPeriod:      61,
		LockupPeriodUnit:  "day",
		VestingPeriod:     213,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.150000000000000000"),
	}

	vt7 := types.GenesisVestingType{
		Name:              "Public round",
		LockupPeriod:      30,
		LockupPeriodUnit:  "day",
		VestingPeriod:     152,
		VestingPeriodUnit: "day",
		Free:              sdk.MustNewDecFromStr("0.200000000000000000"),
	}

	return []types.GenesisVestingType{vt1, vt2, vt3, vt4, vt5, vt6, vt7}
}
