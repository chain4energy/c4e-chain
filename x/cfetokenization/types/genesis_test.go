package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
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

				CertificateTypeList: []types.CertificateType{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				CertificateTypeCount: 2,
				UserDevicesList: []types.UserDevices{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				UserDevicesCount: 2,
				UserCertificatesList: []types.UserCertificates{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				UserCertificatesCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated certificateType",
			genState: &types.GenesisState{
				CertificateTypeList: []types.CertificateType{
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
			desc: "invalid certificateType count",
			genState: &types.GenesisState{
				CertificateTypeList: []types.CertificateType{
					{
						Id: 1,
					},
				},
				CertificateTypeCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated userDevices",
			genState: &types.GenesisState{
				UserDevicesList: []types.UserDevices{
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
			desc: "invalid userDevices count",
			genState: &types.GenesisState{
				UserDevicesList: []types.UserDevices{
					{
						Id: 1,
					},
				},
				UserDevicesCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated userCertificates",
			genState: &types.GenesisState{
				UserCertificatesList: []types.UserCertificates{
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
			desc: "invalid userCertificates count",
			genState: &types.GenesisState{
				UserCertificatesList: []types.UserCertificates{
					{
						Id: 1,
					},
				},
				UserCertificatesCount: 0,
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
