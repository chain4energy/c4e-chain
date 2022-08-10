package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				EnergyTokenList: []types.EnergyToken{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				EnergyTokenCount: 2,
				TokenParamsList: []types.TokenParams{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated energyToken",
			genState: &types.GenesisState{
				EnergyTokenList: []types.EnergyToken{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid energyToken count",
			genState: &types.GenesisState{
				EnergyTokenList: []types.EnergyToken{
					{
						Id: 1,
					},
				},
				EnergyTokenCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated tokenParams",
			genState: &types.GenesisState{
				TokenParamsList: []types.TokenParams{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
