package types_test

import (
	appparams "github.com/chain4energy/c4e-chain/v2/app/params"
	testenv "github.com/chain4energy/c4e-chain/v2/testutil/env"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgUpdateDenomParam_ValidateBasic(t *testing.T) {
	tests := []struct {
		name         string
		msg          types.MsgUpdateDenomParam
		expectError  bool
		errorMessage string
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateDenomParam{
				Authority: "abcd",
			},
			expectError:  true,
			errorMessage: "expected gov account as only signer for proposal message",
		},
		{
			name: "empty denom",
			msg: types.MsgUpdateDenomParam{
				Authority: appparams.GetAuthority(),
				Denom:     "",
			},
			expectError:  true,
			errorMessage: "denom cannot be empty",
		},
		{
			name: "correct denom",
			msg: types.MsgUpdateDenomParam{
				Authority: appparams.GetAuthority(),
				Denom:     testenv.DefaultTestDenom,
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
