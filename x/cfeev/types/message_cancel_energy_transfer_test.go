package types_test

import (
	"github.com/chain4energy/c4e-chain/testutil/sample"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCancelEnergyTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgCancelEnergyTransfer
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgCancelEnergyTransfer{
				Creator:          "invalid_address",
				EnergyTransferId: 0,
				ChargerId:        validChargerId,
				ErrorInfo:        "",
				ErrorCode:        "",
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid creator address (decoding bech32 failed: invalid separator index -1): invalid address",
		},
		{
			name: "valid address",
			msg: types.MsgCancelEnergyTransfer{
				Creator:          sample.AccAddress(),
				EnergyTransferId: 0,
				ChargerId:        validChargerId,
				ErrorInfo:        "",
				ErrorCode:        "",
			},
		},
		{
			name: "empty charger id",
			msg: types.MsgCancelEnergyTransfer{
				Creator:          sample.AccAddress(),
				EnergyTransferId: 0,
				ChargerId:        "",
				ErrorInfo:        "",
				ErrorCode:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "charger id cannot be empty: wrong param value",
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
