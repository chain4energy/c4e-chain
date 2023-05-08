package types_test

import (
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgStartCampaign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgStartCampaign
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgStartCampaign{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgStartCampaign{
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
