package types_test

import (
	"github.com/chain4energy/c4e-chain/testutil/sample"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

const (
	validTariff           = 100
	validEnergyToTransfer = 100
)

func TestMsgStartEnergyTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgStartEnergyTransfer
		err    error
		errMsg string
	}{
		{
			name: "invalid creator address",
			msg: types.MsgStartEnergyTransfer{
				Creator:               "invalid_address",
				EnergyTransferOfferId: 0,
				OfferedTariff:         validTariff,
				EnergyToTransfer:      validEnergyToTransfer,
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid creator address (decoding bech32 failed: invalid separator index -1): invalid address",
		},
		{
			name: "valid message",
			msg: types.MsgStartEnergyTransfer{
				Creator:               sample.AccAddress(),
				EnergyTransferOfferId: 0,
				OfferedTariff:         validTariff,
				EnergyToTransfer:      validEnergyToTransfer,
			},
		},
		{
			name: "invalid offered tariff",
			msg: types.MsgStartEnergyTransfer{
				Creator:               sample.AccAddress(),
				EnergyTransferOfferId: 0,
				OfferedTariff:         0,
				EnergyToTransfer:      validEnergyToTransfer,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "offered tariff cannot be empty: wrong param value",
		},
		{
			name: "invalid energy to transfer",
			msg: types.MsgStartEnergyTransfer{
				Creator:               sample.AccAddress(),
				EnergyTransferOfferId: 0,
				OfferedTariff:         validTariff,
				EnergyToTransfer:      0,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "cannot transfer zero [kWh] energy: wrong param value",
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
