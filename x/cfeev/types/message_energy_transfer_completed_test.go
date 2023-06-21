package types_test

import (
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

const validChargerId = "valid_charger_id"

func TestMsgEnergyTransferCompleted_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgEnergyTransferCompleted
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgEnergyTransferCompleted{
				Creator:          "invalid_address",
				EnergyTransferId: 0,
				UsedServiceUnits: 0,
				Info:             "",
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid creator address (decoding bech32 failed: invalid separator index -1): invalid address",
		},
		{
			name: "valid address",
			msg: types.MsgEnergyTransferCompleted{
				Creator:          sample.AccAddress(),
				EnergyTransferId: 0,
				UsedServiceUnits: 0,
				Info:             "",
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
