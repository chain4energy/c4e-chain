package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgClaim_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgClaim
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgClaim{
				Claimer: "invalid_address",
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid claimer address (decoding bech32 failed: invalid separator index -1): invalid address",
		}, {
			name: "valid address",
			msg: types.MsgClaim{
				Claimer: sample.AccAddress(),
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
