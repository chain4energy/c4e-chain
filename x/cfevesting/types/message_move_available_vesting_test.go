package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
)

func TestMsgMoveAvailableVesting_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgMoveAvailableVesting
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgMoveAvailableVesting{
				FromAddress: "invalid_address",
				ToAddress:   sample.AccAddress(),
			},
			err:    types.ErrParsing,
			errMsg: "move available vesting - from acc address error: decoding bech32 failed: invalid separator index -1: failed to parse",
		},
		{
			name: "invalid address",
			msg: types.MsgMoveAvailableVesting{
				FromAddress: sample.AccAddress(),
				ToAddress:   "invalid_address",
			},
			err:    types.ErrParsing,
			errMsg: "move available vesting - to acc address error: decoding bech32 failed: invalid separator index -1: failed to parse",
		},
		{
			name: "valid address",
			msg: types.MsgMoveAvailableVesting{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				require.EqualError(t, err, tt.errMsg)
				return
			}
			require.NoError(t, err)
		})
	}
}
