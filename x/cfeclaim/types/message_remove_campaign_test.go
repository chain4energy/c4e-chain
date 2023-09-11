package types_test

import (
	"github.com/chain4energy/c4e-chain/v2/testutil/sample"
	"testing"

	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRemoveCampaign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgRemoveCampaign
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgRemoveCampaign{
				Owner:      "abcd",
				CampaignId: 0,
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid creator address (decoding bech32 failed: invalid bech32 string length 4): invalid address",
		}, {
			name: "valid message",
			msg: types.MsgRemoveCampaign{
				Owner:      sample.AccAddress(),
				CampaignId: 0,
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
