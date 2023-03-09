package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSendToVestingAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgSendToVestingAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgSendToVestingAccount{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgSendToVestingAccount{
				Owner: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
