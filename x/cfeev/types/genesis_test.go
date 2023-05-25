package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
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

				EnergyTransferOfferList: []types.EnergyTransferOffer{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				EnergyTransferOfferCount: 2,
				EnergyTransferList: []types.EnergyTransfer{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				EnergyTransferCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated energyTransferOffer",
			genState: &types.GenesisState{
				EnergyTransferOfferList: []types.EnergyTransferOffer{
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
			desc: "invalid energyTransferOffer count",
			genState: &types.GenesisState{
				EnergyTransferOfferList: []types.EnergyTransferOffer{
					{
						Id: 1,
					},
				},
				EnergyTransferOfferCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated energyTransfer",
			genState: &types.GenesisState{
				EnergyTransferList: []types.EnergyTransfer{
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
			desc: "invalid energyTransfer count",
			genState: &types.GenesisState{
				EnergyTransferList: []types.EnergyTransfer{
					{
						Id: 1,
					},
				},
				EnergyTransferCount: 0,
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
