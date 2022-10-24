package types_test

import (
	"testing"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
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
	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

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
		validVestingPoolsTest(acountsAddresses),
		invalidVestingPoolsNoVestingTypes(acountsAddresses),
		invalidVestingPoolsVestingTypeNotFound(acountsAddresses),
		invalidVestingPoolsMoreThanOneIdError(acountsAddresses),
		invalidVestingPoolsMoreThanOneNameError(acountsAddresses),
		invalidVestingPoolsMoreThanOneAddressError(acountsAddresses),
		invalidVestingTypesWrongLockupPeriodUnitTest(),
		invalidVestingTypesWrongVestingPeriodUnitTest(),
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

			VestingAccountList: []types.VestingAccount{
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

			VestingAccountList: []types.VestingAccount{
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
			VestingAccountList: []types.VestingAccount{
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
			VestingAccountList: []types.VestingAccount{
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
			VestingAccountList: []types.VestingAccount{
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
			VestingAccountList: []types.VestingAccount{
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

func validVestingPoolsTest(acountsAddresses []sdk.AccAddress) TcData {
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

func invalidVestingPoolsNoVestingTypes(acountsAddresses []sdk.AccAddress) TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)

	return TcData{
		desc: "invalid VestingPools no vesting types",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        []types.GenesisVestingType{},
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting with id: 1 defined for account: " + accountVestingPoolsArray[0].Address + " - vesting type not found: test-vesting-account-1-1",
	}

}

func invalidVestingPoolsVestingTypeNotFound(acountsAddresses []sdk.AccAddress) TcData {
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
		errorMassage: "vesting with id: 8 defined for account: " + accountVestingPoolsArray[4].Address + " - vesting type not found: " + accountVestingPoolsArray[4].VestingPools[7].VestingType,
	}

}

func invalidVestingPoolsMoreThanOneIdError(acountsAddresses []sdk.AccAddress) TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	accountVestingPoolsArray[4].VestingPools[3].Id = accountVestingPoolsArray[4].VestingPools[6].Id
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)

	return TcData{
		desc: "invalid VestingPools more than one id",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting with id: 7 defined more than once for account: " + accountVestingPoolsArray[4].Address,
	}

}

func invalidVestingPoolsMoreThanOneNameError(acountsAddresses []sdk.AccAddress) TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	accountVestingPoolsArray[4].VestingPools[3].Name = accountVestingPoolsArray[4].VestingPools[6].Name
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)

	return TcData{
		desc: "invalid VestingPools more than one name",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "vesting with name: " + accountVestingPoolsArray[4].VestingPools[3].Name + " defined more than once for account: " + accountVestingPoolsArray[4].Address,
	}
}

func invalidVestingPoolsMoreThanOneAddressError(acountsAddresses []sdk.AccAddress) TcData {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)
	accountVestingPoolsArray[3].Address = accountVestingPoolsArray[7].Address
	vestingTypes := testutils.GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPoolsArray)

	return TcData{
		desc: "invalid VestingPools more than one address",
		genState: &types.GenesisState{
			Params: types.NewParams("test_denom"),

			VestingTypes:        vestingTypes,
			AccountVestingPools: accountVestingPoolsArray,
		},
		valid:        false,
		errorMassage: "account vesting pools with address: " + accountVestingPoolsArray[3].Address + " defined more than once",
	}

}

func invalidVestingTypesWrongLockupPeriodUnitTest() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[7].LockupPeriodUnit = "err-unit"
	return TcData{
		desc: "valid vestingTypes",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid:        false,
		errorMassage: "Unknown PeriodUnit: " + vestingTypes[7].LockupPeriodUnit + ": invalid type",
	}
}

func invalidVestingTypesWrongVestingPeriodUnitTest() TcData {
	vestingTypes := testutils.GenerateGenesisVestingTypes(10, 1)
	vestingTypes[7].VestingPeriodUnit = "err-unit"
	return TcData{
		desc: "valid vestingTypes",
		genState: &types.GenesisState{
			Params:       types.NewParams("test_denom"),
			VestingTypes: vestingTypes,
		},
		valid:        false,
		errorMassage: "Unknown PeriodUnit: " + vestingTypes[7].VestingPeriodUnit + ": invalid type",
	}
}
