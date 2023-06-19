package types_test

import (
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateNewAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgCreateNewAccount
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgCreateNewAccount{
				Creator: "invalid_address",
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid creator address (decoding bech32 failed: invalid separator index -1): invalid address",
		}, {
			name: "valid address",
			msg: types.MsgCreateNewAccount{
				Creator: sample.AccAddress(),
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
