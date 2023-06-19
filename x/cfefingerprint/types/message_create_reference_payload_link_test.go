package types_test

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

const (
	validPayloadHash = "YWFhYWZmZjQ0cnJmZmZkc2RlZGVmZXRldAo="
)

func TestMsgCreateReferencePayloadLink_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgCreateReferencePayloadLink
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgCreateReferencePayloadLink{
				Creator:     "invalid_address",
				PayloadHash: validPayloadHash,
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid creator address (decoding bech32 failed: invalid separator index -1): invalid address",
		},
		{
			name: "valid address",
			msg: types.MsgCreateReferencePayloadLink{
				Creator:     sample.AccAddress(),
				PayloadHash: validPayloadHash,
			},
		},
		{
			name: "empty payload hash",
			msg: types.MsgCreateReferencePayloadLink{
				Creator:     sample.AccAddress(),
				PayloadHash: "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "payload hash cannot be empty: wrong param value",
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
