package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateAirdropCampaign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateAirdropCampaign
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateAirdropCampaign{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
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
