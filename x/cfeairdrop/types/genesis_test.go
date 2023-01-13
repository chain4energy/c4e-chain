package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
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

				ClaimRecords: []types.ClaimRecord{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				InitialClaims: []types.InitialClaim{
					{
						CampaignId: 0,
					},
					{
						CampaignId: 1,
					},
				},
				Missions: []types.Mission{
					{
						CampaignId: 0,
						MissionId:  0,
					},
					{
						CampaignId: 1,
						MissionId:  1,
					},
				},
				AirdropEntryList: []types.AirdropEntry{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				AirdropEntryCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated claimRecord",
			genState: &types.GenesisState{
				ClaimRecords: []types.ClaimRecord{
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
			desc: "duplicated initialClaim",
			genState: &types.GenesisState{
				InitialClaims: []types.InitialClaim{
					{
						CampaignId: 0,
					},
					{
						CampaignId: 0,
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
						MissionId:  0,
					},
					{
						CampaignId: 0,
						MissionId:  0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated airdropEntry",
			genState: &types.GenesisState{
				AirdropEntryList: []types.AirdropEntry{
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
			desc: "invalid airdropEntry count",
			genState: &types.GenesisState{
				AirdropEntryList: []types.AirdropEntry{
					{
						Id: 1,
					},
				},
				AirdropEntryCount: 0,
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
