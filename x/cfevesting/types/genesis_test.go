package types_test

import (
	"fmt"
	"testing"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type TcData struct {
	desc         string
	genState     *types.GenesisState
	valid        bool
	errorMassage string
}

func TestGenesisState_Validate(t *testing.T) {
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	for _, tc := range []TcData{
		defaultGenesisTest(),
		validVestingAccountsTest(acountsAddresses),
		emptyDenomTest(acountsAddresses),
		duplicatedVestingAccountTest(acountsAddresses),
		invalidVestingAccountCountTest(acountsAddresses),
		validVestingTypesTest(),
		validVestingAccountsWithVestingTypesTest(acountsAddresses),
		invalidVestingTypesNameMoreThanOnceError(),
		invalidVestingAccountsWrongIdTest(acountsAddresses),
		validVestingPoolsTest(),
		invalidVestingPoolsNoVestingTypes(),
		invalidVestingPoolsVestingTypeNotFound(),
		invalidVestingPoolsOneNameMoreThanOnceError(),
		invalidVestingPoolsMoreThanOneAddressError(),
		invalidVestingPoolsEmptyName(),
		invalidVestingPoolsNegativeInitiallyLocked(),
		invalidVestingPoolsNegativeWithdrawn(),
		invalidVestingPoolsNegativeSent(),
		invalidVestingPoolsNegativeCurrentlyLocked(),
		invalidVestingTypesWrongLockupPeriodUnitTest(),
		invalidVestingTypesWrongVestingPeriodUnitTest(),
		invalidVestingTypesEmptyNameTest(),
		invalidVestingTypesNegativeLockupPeriodTest(),
		invalidVestingTypesNegativeVestingPeriodTest(),
		validVestingTypesFreeEquals1(),
		validVestingTypesFreeEquals0(),
		invalidVestingTypesFreeGreaterThan1(),
		invalidVestingTypesFreeLowerThan0(),
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.errorMassage)
			}
		})
	}
}

func defaultGenesisTest() TcData {
	return TcData{
		desc:     "default is valid",
		genState: types.DefaultGenesis(),
		valid:    true,
	}
}

func validVestingAccountsTest(acountsAddresses []sdk.AccAddress) TcData {
	return TcData{
		desc: "valid vesting accounts",
		genState: &types.GenesisState{

			VestingAccountTraces: []types.VestingAccountTrace{
				{
					Id:      0,
					Address: acountsAddresses[0].String(),
				},
				{
					Id:      1,
					Address: acountsAddresses[1].String(),
				},
			},
			VestingAccountCount: 2,
			Params:              types.Params{Denom: "uc4e"},
			// this line is used by starport scaffolding # types/genesis/validField
		},
		valid: true,
	}
}

func emptyDenomTest(acountsAddresses []sdk.AccAddress) TcData {
	return TcData{
		desc: "empty denom",
		genState: &types.GenesisState{

			VestingAccountTraces: []types.VestingAccountTrace{
				{
					Id:      0,
					Address: acountsAddresses[0].String(),
				},
				{
					Id:      1,
					Address: acountsAddresses[1].String(),
				},
			},
			VestingAccountCount: 2,
			// this line is used by starport scaffolding # types/genesis/validField
		},
		valid:        false,
		errorMassage: "denom cannot be empty",
	}
}

func duplicatedVestingAccountTest(acountsAddresses []sdk.AccAddress) TcData {
	return TcData{
		desc: "duplicated vestingAccount",
		genState: &types.GenesisState{
			VestingAccountTraces: []types.VestingAccountTrace{
				{
					Id:      0,
					Address: acountsAddresses[0].String(),
				},
				{
					Id:      0,
					Address: acountsAddresses[0].String(),
				},
			},
			VestingAccountCount: 2,
		},
		valid:        false,
		errorMassage: "duplicated id for vestingAccount",
	}
}

func invalidVestingAccountCountTest(acountsAddresses []sdk.AccAddress) TcData {
	return TcData{
		desc: "invalid vestingAccount count",
		genState: &types.GenesisState{
			VestingAccountTraces: []types.VestingAccountTrace{
				{
					Id:      1,
					Address: acountsAddresses[0].String(),
				},
			},
			VestingAccountCount: 0,
		},
		valid:        false,
		errorMassage: "vestingAccount id should be lower or equal than the last id",
	}
}

func validVestingTypesTest() TcData {
	return TcData{
		desc: "valid vestingTypes",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: testutils.GenerateGenesisVestingTypes(10, 1),
		},
		valid: true,
	}
}

func validVestingAccountsWithVestingTypesTest(acountsAddresses []sdk.AccAddress) TcData {
	return TcData{
		desc: "Valid VestingAccounts",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),
			VestingAccountTraces: []types.VestingAccountTrace{
				{
					Id:      0,
					Address: acountsAddresses[0].String(),
				},
				{
					Id:      1,
					Address: acountsAddresses[1].String(),
				},
			},
			VestingAccountCount: 2,
			VestingTypes:        testutils.GenerateGenesisVestingTypes(10, 1),
		},
		valid: true,
	}
}

func invalidVestingTypesNameMoreThanOnceError() TcData {
	vestingTypesArray := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypesArray[3].Name = vestingTypesArray[6].Name
	return TcData{
		desc: "invalid VestingTypes name more than once",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypesArray,
		},
		valid:        false,
		errorMassage: "vesting type with name: test-vesting-type-7 defined more than once",
	}
}

func invalidVestingAccountsWrongIdTest(acountsAddresses []sdk.AccAddress) TcData {
	return TcData{
		desc: "invalid VestingAccounts wrong id",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),
			VestingAccountTraces: []types.VestingAccountTrace{
				{
					Id:      0,
					Address: acountsAddresses[0].String(),
				},
				{
					Id:      1,
					Address: acountsAddresses[1].String(),
				},
			},
			VestingAccountCount: 1,
			VestingTypes:        testutils.GenerateGenesisVestingTypes(10, 1),
		},
		valid:        false,
		errorMassage: "vestingAccount id should be lower or equal than the last id",
	}
}

func validVestingPoolsTest() TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)

	return TcData{
		desc: "valid Vesting Pools",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid: true,
	}
}

func invalidVestingPoolsNoVestingTypes() TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)

	return TcData{
		desc: "invalid VestingPools no vesting types",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        []types.GenesisVestingType{},
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting pool with name: test-vesting-account-name1-1 defined for account: " + accountVestingPoolsArray[0].Owner + " - vesting type not found: test-vesting-account-1-1",
	}

}

func invalidVestingPoolsVestingTypeNotFound() TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)
	accountVestingPoolsArray[4].VestingPools[7].VestingType = "wrong type"

	return TcData{
		desc: "invalid VestingPools vesting type not found",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting pool with name: test-vesting-account-name5-8 defined for account: " + accountVestingPoolsArray[4].Owner + " - vesting type not found: " + accountVestingPoolsArray[4].VestingPools[7].VestingType,
	}

}

func invalidVestingPoolsOneNameMoreThanOnceError() TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	accountVestingPoolsArray[4].VestingPools[3].Name = accountVestingPoolsArray[4].VestingPools[6].Name
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)

	return TcData{
		desc: "invalid VestingPools name more than once",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting pool with name: " + accountVestingPoolsArray[4].VestingPools[3].Name + " defined more than once for account: " + accountVestingPoolsArray[4].Owner,
	}
}

func invalidVestingPoolsMoreThanOneAddressError() TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	accountVestingPoolsArray[3].Owner = accountVestingPoolsArray[7].Owner
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)

	return TcData{
		desc: "invalid VestingPools more than one address",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "account vesting pools with address: " + accountVestingPoolsArray[3].Owner + " defined more than once",
	}

}

func invalidVestingPoolsEmptyName() TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)
	accountVestingPoolsArray[4].VestingPools[3].Name = ""

	return TcData{
		desc: "invalid VestingPools empty name",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting pool defined for account: " + accountVestingPoolsArray[4].Owner + " has no name",
	}

}

func invalidVestingPoolsNegativeInitiallyLocked() TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)
	accountVestingPoolsArray[4].VestingPools[3].InitiallyLocked = sdk.NewInt(-1)

	return TcData{
		desc: "invalid VestingPools negative initially locked",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting pool with name: " + accountVestingPoolsArray[4].VestingPools[3].Name + " defined for account: " + accountVestingPoolsArray[4].Owner + " has InitiallyLocked value negative -1",
	}

}

func invalidVestingPoolsNegativeWithdrawn() TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)
	accountVestingPoolsArray[4].VestingPools[3].Withdrawn = sdk.NewInt(-1)

	return TcData{
		desc: "invalid VestingPools negative withdrawn",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting pool with name: " + accountVestingPoolsArray[4].VestingPools[3].Name + " defined for account: " + accountVestingPoolsArray[4].Owner + " has Withdrawn value negative -1",
	}

}

func invalidVestingPoolsNegativeSent() TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)
	accountVestingPoolsArray[4].VestingPools[3].Sent = sdk.NewInt(-1)

	return TcData{
		desc: "invalid VestingPools negative sent",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting pool with name: " + accountVestingPoolsArray[4].VestingPools[3].Name + " defined for account: " + accountVestingPoolsArray[4].Owner + " has Sent value negative -1",
	}

}

func invalidVestingPoolsNegativeCurrentlyLocked() TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)
	accountVestingPoolsArray[4].VestingPools[3].InitiallyLocked = sdk.NewInt(100)
	accountVestingPoolsArray[4].VestingPools[3].Withdrawn = sdk.NewInt(50)
	accountVestingPoolsArray[4].VestingPools[3].Sent = sdk.NewInt(51)

	return TcData{
		desc: "invalid VestingPools empty name",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting pool with name: " + accountVestingPoolsArray[4].VestingPools[3].Name + " defined for account: " + accountVestingPoolsArray[4].Owner + " has InitiallyLocked (100) < Withdrawn (50) + Sent (51)",
	}

}

func invalidVestingTypesWrongLockupPeriodUnitTest() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[7].LockupPeriodUnit = "err-unit"
	return TcData{
		desc: "invalid vestingTypes wrong lockup period unit",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid:        false,
		errorMassage: "LockupPeriodUnit of veting type: " + vestingTypes[7].Name + " error: Unknown PeriodUnit: " + vestingTypes[7].LockupPeriodUnit + ": invalid type" + getWrongUnitMessageCodeLineInfo(),
	}
}

func invalidVestingTypesWrongVestingPeriodUnitTest() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[7].VestingPeriodUnit = "err-unit"
	return TcData{
		desc: "invalid vestingTypes wrong vesting period unit",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid:        false,
		errorMassage: "VestingPeriodUnit of veting type: " + vestingTypes[7].Name + " error: Unknown PeriodUnit: " + vestingTypes[7].VestingPeriodUnit + ": invalid type" + getWrongUnitMessageCodeLineInfo(),
	}
}

func invalidVestingTypesNegativeLockupPeriodTest() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[7].LockupPeriod = -1
	return TcData{
		desc: "invalid vestingTypes wrong lockup period unit",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid:        false,
		errorMassage: "LockupPeriod of veting type: " + vestingTypes[7].Name + " less than 0",
	}
}

func invalidVestingTypesNegativeVestingPeriodTest() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[7].VestingPeriod = -1
	return TcData{
		desc: "invalid vestingTypes wrong lockup period unit",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid:        false,
		errorMassage: "VestingPeriod of veting type: " + vestingTypes[7].Name + " less than 0",
	}
}

func invalidVestingTypesEmptyNameTest() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[7].Name = ""
	return TcData{
		desc: "valid vestingTypes empty name",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid:        false,
		errorMassage: "vesting type has no name",
	}
}

func getWrongUnitMessageCodeLineInfo() string {

	unit := types.PeriodUnit("unit")
	_, err := types.DurationFromUnits(unit, 0)
	err = fmt.Errorf("%w", err)
	startLen := len("Unknown PeriodUnit: " + unit + ": invalid type")
	errLen := len(err.Error())
	return err.Error()[startLen:errLen]
}

func invalidVestingTypesFreeGreaterThan1() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[6].Free = sdk.MustNewDecFromStr("1.1")
	return TcData{
		desc: "invalid vestingTypes initial bonus greater than 100",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid:        false,
		errorMassage: "Free of veting type " + vestingTypes[6].Name + " must be set between 0 and 1",
	}
}

func invalidVestingTypesFreeLowerThan0() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[6].Free = sdk.MustNewDecFromStr("-1")
	return TcData{
		desc: "invalid vestingTypes initial bonus lower than 100",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid:        false,
		errorMassage: "Free of veting type " + vestingTypes[6].Name + " must be set between 0 and 1",
	}
}

func validVestingTypesFreeEquals1() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[6].Free = sdk.MustNewDecFromStr("1")
	return TcData{
		desc: "valid vestingTypes initial bonus equals 1",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid: true,
	}
}

func validVestingTypesFreeEquals0() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[6].Free = sdk.MustNewDecFromStr("0")
	return TcData{
		desc: "valid vestingTypes initial bonus equals 0",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid: true,
	}
}
