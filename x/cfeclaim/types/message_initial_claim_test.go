package types_test

import (
	"github.com/chain4energy/c4e-chain/v2/testutil/sample"
	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgInitialClaim_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgInitialClaim
		err    error
		errMsg string
	}{
		{
			name: "invalid claimer address",
			msg: types.MsgInitialClaim{
				Claimer:            "invalid_address",
				CampaignId:         0,
				DestinationAddress: sample.AccAddress(),
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid claimer address (decoding bech32 failed: invalid separator index -1): invalid address",
		},
		{
			name: "invalid destination address",
			msg: types.MsgInitialClaim{
				Claimer:            sample.AccAddress(),
				CampaignId:         0,
				DestinationAddress: "",
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid destination address (empty address string is not allowed): invalid address",
		},
		{
			name: "valid message",
			msg: types.MsgInitialClaim{
				Claimer:            sample.AccAddress(),
				CampaignId:         0,
				DestinationAddress: sample.AccAddress(),
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
