package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
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
				Params: types.DefaultParams(),

				UsersEntries: []types.UserEntry{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				Missions: []types.Mission{
					{
						CampaignId: 0,
						Id:         0,
					},
					{
						CampaignId: 1,
						Id:         1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated userEntry",
			genState: &types.GenesisState{
				UsersEntries: []types.UserEntry{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated mission",
			genState: &types.GenesisState{
				Missions: []types.Mission{
					{
						CampaignId: 0,
						Id:         0,
					},
					{
						CampaignId: 0,
						Id:         0,
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
