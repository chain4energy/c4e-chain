package types_test

import (
	cfedistributortestutils "github.com/chain4energy/c4e-chain/v2/testutil/module/cfedistributor"
	"testing"

	"github.com/chain4energy/c4e-chain/v2/x/cfedistributor/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	cfedistributortestutils.SetTestMaccPerms()
	for _, tc := range []struct {
		desc        string
		genState    *types.GenesisState
		expectError bool
	}{
		{
			desc:        "default is valid",
			genState:    types.DefaultGenesis(),
			expectError: false,
		},
		{
			desc:     "invalid genesis state - there must be at least one  subdistributor with the source main type",
			genState: &types.GenesisState{
				// this line is used by starport scaffolding # types/genesis/validField
			},
			expectError: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.expectError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
