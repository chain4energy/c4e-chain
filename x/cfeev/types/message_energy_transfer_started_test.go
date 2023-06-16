package types_test

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgEnergyTransferStarted_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgEnergyTransferStarted
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgEnergyTransferStarted{
				Creator:          "invalid_address",
				EnergyTransferId: 0,
				ChargerId:        validChargerId,
				Info:             "",
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid creator address (decoding bech32 failed: invalid separator index -1): invalid address",
		},
		{
			name: "valid address",
			msg: types.MsgEnergyTransferStarted{
				Creator:          sample.AccAddress(),
				EnergyTransferId: 0,
				ChargerId:        validChargerId,
				Info:             "",
			},
		},
		{
			name: "empty charger id",
			msg: types.MsgEnergyTransferStarted{
				Creator:          sample.AccAddress(),
				EnergyTransferId: 0,
				ChargerId:        "",
				Info:             "",
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
