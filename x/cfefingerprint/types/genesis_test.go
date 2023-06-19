package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/stretchr/testify/require"
)

const (
	validReferenceKey   = "abcs123123"
	validReferenceValue = "abcs123123"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		name     string
		genState *types.GenesisState
		errMsg   string
	}{
		{
			name:     "default is valid",
			genState: types.DefaultGenesis(),
		},
		{
			name:     "empty genesis state",
			genState: &types.GenesisState{},
		},
		{
			name: "valid genesis state",
			genState: &types.GenesisState{
				PayloadLinks: []*types.GenesisPayloadLink{
					{
						ReferenceKey:   validReferenceKey,
						ReferenceValue: validReferenceValue,
					},
				},
			},
		},
		{
			name: "invalid genesis state - empty reference key",
			genState: &types.GenesisState{
				PayloadLinks: []*types.GenesisPayloadLink{
					{
						ReferenceKey:   "",
						ReferenceValue: validReferenceValue,
					},
				},
			},
			errMsg: "referance key cannot be empty",
		},
		{
			name: "invalid genesis state - empty reference value",
			genState: &types.GenesisState{
				PayloadLinks: []*types.GenesisPayloadLink{
					{
						ReferenceKey:   validReferenceKey,
						ReferenceValue: "",
					},
				},
			},
			errMsg: "referance value cannot be empty",
		},
		{
			name: "invalid genesis state - dyplicated reference key",
			genState: &types.GenesisState{
				PayloadLinks: []*types.GenesisPayloadLink{
					{
						ReferenceKey:   validReferenceKey,
						ReferenceValue: validReferenceValue,
					},
					{
						ReferenceKey:   validReferenceKey,
						ReferenceValue: validReferenceValue,
					},
				},
			},
			errMsg: "duplicated reference key abcs123123",
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.genState.Validate()
			if tt.errMsg != "" {
				require.EqualError(t, err, tt.errMsg)
				return
			}
			require.NoError(t, err)
		})
	}
}
