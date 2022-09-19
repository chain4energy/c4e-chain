package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

)

func TestGenesisState_Validate(t *testing.T) {
	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	for _, tc := range []struct {
		desc         string
		genState     *types.GenesisState
		valid        bool
		errorMassage string
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				VestingAccountList: []types.VestingAccount{
					{
						Id: 0,
						Address: acountsAddresses[0].String(),
					},
					{
						Id: 1,
						Address: acountsAddresses[1].String(),
					},
				},
				VestingAccountCount: 2,
				Params: types.Params{Denom: "uc4e"},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        true,
		},
		{
			desc: "empty genesis state",
			genState: &types.GenesisState{

				VestingAccountList: []types.VestingAccount{
					{
						Id: 0,
						Address: acountsAddresses[0].String(),
					},
					{
						Id: 1,
						Address: acountsAddresses[1].String(),
					},
				},
				VestingAccountCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "denom cannot be empty",
		},
		{
			desc: "duplicated vestingAccount",
			genState: &types.GenesisState{
				VestingAccountList: []types.VestingAccount{
					{
						Id: 0,
						Address: acountsAddresses[0].String(),
					},
					{
						Id: 0,
						Address: acountsAddresses[0].String(),
					},
				},
				VestingAccountCount: 2,
			},
			valid: false,
			errorMassage: "duplicated id for vestingAccount",
		},
		{
			desc: "invalid vestingAccount count",
			genState: &types.GenesisState{
				VestingAccountList: []types.VestingAccount{
					{
						Id: 1,
						Address: acountsAddresses[0].String(),
					},
				},
				VestingAccountCount: 0,
			},
			valid: false,
			errorMassage: "vestingAccount id should be lower or equal than the last id",
		},
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
