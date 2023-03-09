package types_test

import (
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

var correctSubDistributor = CreateSubDistributor(MAIN_SOURCE)

var correctSubDistributors = []types.SubDistributor{
	CreateSubDistributor(MAIN_SOURCE),
}

func TestMsgUpdateAllSubDistributorsParams_ValidateBasic(t *testing.T) {
	tests := []struct {
		name         string
		msg          types.MsgUpdateAllSubDistributorsParams
		expectError  bool
		errorMessage string
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateAllSubDistributorsParams{
				Authority: "abcd",
			},
			expectError:  true,
			errorMessage: "expected gov account as only signer for proposal message",
		},
		{
			name: "empty sub distributors",
			msg: types.MsgUpdateAllSubDistributorsParams{
				Authority: testenv.GetAuthority(),
				SubDistributors: []types.SubDistributor{
					{
						Name:         "",
						Sources:      nil,
						Destinations: types.Destinations{},
					},
				},
			},
			expectError:  true,
			errorMessage: "validation error: subdistributor name cannot be empty: invalid proposal content",
		},
		{
			name: "correct denom",
			msg: types.MsgUpdateAllSubDistributorsParams{
				Authority:       testenv.GetAuthority(),
				SubDistributors: correctSubDistributors,
			},
			expectError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateSubDistributorParam_ValidateBasic(t *testing.T) {
	tests := []struct {
		name         string
		msg          types.MsgUpdateSubDistributorParam
		expectError  bool
		errorMessage string
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateSubDistributorParam{
				Authority: "abcd",
			},
			expectError:  true,
			errorMessage: "expected gov account as only signer for proposal message",
		},
		{
			name: "empty sub distributor",
			msg: types.MsgUpdateSubDistributorParam{
				Authority:      testenv.GetAuthority(),
				SubDistributor: &types.SubDistributor{},
			},
			expectError:  true,
			errorMessage: "validation error: subdistributor name cannot be empty: invalid proposal content",
		},
		{
			name: "correct sub distributor",
			msg: types.MsgUpdateSubDistributorParam{
				Authority:      testenv.GetAuthority(),
				SubDistributor: &correctSubDistributor,
			},
			expectError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateSubDistributorBurnShareParam_ValidateBasic(t *testing.T) {
	tests := []struct {
		name         string
		msg          types.MsgUpdateSubDistributorBurnShareParam
		expectError  bool
		errorMessage string
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateSubDistributorBurnShareParam{
				Authority: "abcd",
			},
			expectError:  true,
			errorMessage: "expected gov account as only signer for proposal message",
		},
		{
			name: "empty sub distributor name",
			msg: types.MsgUpdateSubDistributorBurnShareParam{
				Authority:          testenv.GetAuthority(),
				SubDistributorName: "",
				BurnShare:          sdk.MustNewDecFromStr("0.5"),
			},
			expectError:  true,
			errorMessage: "empty sub distributor name: invalid proposal content",
		},
		{
			name: "negative burn share",
			msg: types.MsgUpdateSubDistributorBurnShareParam{
				Authority:          testenv.GetAuthority(),
				SubDistributorName: "Abc",
				BurnShare:          sdk.MustNewDecFromStr("-0.5"),
			},
			expectError:  true,
			errorMessage: "burn share must be between 0 and 1: invalid proposal content",
		},
		{
			name: "correct message",
			msg: types.MsgUpdateSubDistributorBurnShareParam{
				Authority:          testenv.GetAuthority(),
				SubDistributorName: "Abc",
				BurnShare:          sdk.MustNewDecFromStr("0.5"),
			},
			expectError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateSubDistributorDestinationShareParam_ValidateBasic(t *testing.T) {
	tests := []struct {
		name         string
		msg          types.MsgUpdateSubDistributorDestinationShareParam
		expectError  bool
		errorMessage string
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateSubDistributorDestinationShareParam{
				Authority: "abcd",
			},
			expectError:  true,
			errorMessage: "expected gov account as only signer for proposal message",
		},
		{
			name: "empty sub distributor name",
			msg: types.MsgUpdateSubDistributorDestinationShareParam{
				Authority:          testenv.GetAuthority(),
				SubDistributorName: "",
				DestinationName:    "123",
				Share:              sdk.MustNewDecFromStr("0.5"),
			},
			expectError:  true,
			errorMessage: "empty sub distributor name: invalid proposal content",
		},
		{
			name: "empty destination name",
			msg: types.MsgUpdateSubDistributorDestinationShareParam{
				Authority:          testenv.GetAuthority(),
				SubDistributorName: "Abc",
				DestinationName:    "",
				Share:              sdk.MustNewDecFromStr("0.5"),
			},
			expectError:  true,
			errorMessage: "empty destination name: invalid proposal content",
		},
		{
			name: "negative burn share",
			msg: types.MsgUpdateSubDistributorDestinationShareParam{
				Authority:          testenv.GetAuthority(),
				SubDistributorName: "Abc",
				DestinationName:    "Abc",
				Share:              sdk.MustNewDecFromStr("-0.5"),
			},
			expectError:  true,
			errorMessage: "share must be between 0 and 1: invalid proposal content",
		},
		{
			name: "correct message",
			msg: types.MsgUpdateSubDistributorDestinationShareParam{
				Authority:          testenv.GetAuthority(),
				SubDistributorName: "Abc",
				DestinationName:    "Abc",
				Share:              sdk.MustNewDecFromStr("0.5"),
			},
			expectError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}
