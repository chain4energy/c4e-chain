package types_test

import (
	"github.com/chain4energy/c4e-chain/testutil/module/cfedistributor"
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	types.SetMaccPerms(cfedistributor.TestMaccPerms)
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
			desc:     "invalid genesis state - there must be at least one subdistributor with the source main type",
			genState: &types.GenesisState{
				// this line is used by starport scaffolding # types/genesis/validField
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
