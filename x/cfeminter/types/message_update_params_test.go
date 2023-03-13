package types_test

import (
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgUpdateParams_ValidateBasic(t *testing.T) {
	tests := []struct {
		name         string
		msg          types.MsgUpdateParams
		expectError  bool
		errorMessage string
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateParams{
				Authority: "abcd",
			},
			expectError:  true,
			errorMessage: "expected gov account as only signer for proposal message",
		},
		{
			name: "correct config",
			msg: types.MsgUpdateParams{
				Authority: appparams.GetAuthority(),
				MintDenom: types.DefaultMintDenom,
				StartTime: types.DefaultStartTime,
				Minters:   types.DefaultMinters,
			},
			expectError: false,
		},
		{
			name: "wrong mint denom",
			msg: types.MsgUpdateParams{
				Authority: appparams.GetAuthority(),
				MintDenom: "",
				StartTime: types.DefaultStartTime,
				Minters:   types.DefaultMinters,
			},
			expectError:  true,
			errorMessage: "denom cannot be empty",
		},
		{
			name: "wrong minters",
			msg: types.MsgUpdateParams{
				Authority: appparams.GetAuthority(),
				MintDenom: types.DefaultMintDenom,
				StartTime: types.DefaultStartTime,
				Minters:   WrongMinters(),
			},
			expectError:  true,
			errorMessage: "minter with id 1 validation error: minter config is nil",
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

func TestMsgUpdateMinters_ValidateBasic(t *testing.T) {
	tests := []struct {
		name         string
		msg          types.MsgUpdateMintersParams
		expectError  bool
		errorMessage string
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateMintersParams{
				Authority: "abcd",
			},
			expectError:  true,
			errorMessage: "expected gov account as only signer for proposal message",
		},
		{
			name: "correct config",
			msg: types.MsgUpdateMintersParams{
				Authority: appparams.GetAuthority(),
				StartTime: types.DefaultStartTime,
				Minters:   types.DefaultMinters,
			},
			expectError: false,
		},
		{
			name: "wrong minters",
			msg: types.MsgUpdateMintersParams{
				Authority: appparams.GetAuthority(),
				StartTime: types.DefaultStartTime,
				Minters:   WrongMinters(),
			},
			expectError:  true,
			errorMessage: "minter with id 1 validation error: minter config is nil",
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

func WrongMinters() []*types.Minter {
	return []*types.Minter{
		{
			SequenceId: 1,
			Config:     nil,
		},
	}
}
