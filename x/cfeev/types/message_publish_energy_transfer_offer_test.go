package types_test

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

const validChargerName = "valid_charger_name"

var (
	validLatitude    = sdk.MustNewDecFromStr("0.123")
	invalidLatitude  = sdk.MustNewDecFromStr("95")
	validLongitude   = sdk.MustNewDecFromStr("0.123")
	invalidLongitude = sdk.MustNewDecFromStr("190")
)
var validLocation = types.Location{
	Latitude:  &validLatitude,
	Longitude: &validLongitude,
}

func TestMsgPublishEnergyTransferOffer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgPublishEnergyTransferOffer
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgPublishEnergyTransferOffer{
				Creator:   "invalid_address",
				ChargerId: validChargerId,
				Tariff:    0,
				Location:  nil,
				Name:      "",
				PlugType:  0,
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid creator address (decoding bech32 failed: invalid separator index -1): invalid address",
		}, {
			name: "valid message",
			msg: types.MsgPublishEnergyTransferOffer{
				Creator:   sample.AccAddress(),
				ChargerId: validChargerId,
				Tariff:    0,
				Location:  &validLocation,
				Name:      validChargerName,
				PlugType:  0,
			},
		},
		{
			name: "empty charger id",
			msg: types.MsgPublishEnergyTransferOffer{
				Creator:   sample.AccAddress(),
				ChargerId: "",
				Tariff:    0,
				Location:  &validLocation,
				Name:      validChargerName,
				PlugType:  0,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "charger id cannot be empty: wrong param value",
		},
		{
			name: "empty name",
			msg: types.MsgPublishEnergyTransferOffer{
				Creator:   sample.AccAddress(),
				ChargerId: validChargerName,
				Tariff:    0,
				Location:  &validLocation,
				Name:      "",
				PlugType:  0,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "charger name cannot be empty: wrong param value",
		},
		{
			name: "Latitude nil",
			msg: types.MsgPublishEnergyTransferOffer{
				Creator:   sample.AccAddress(),
				ChargerId: validChargerName,
				Tariff:    0,
				Location: &types.Location{
					Latitude:  nil,
					Longitude: nil,
				},
				Name:     validChargerName,
				PlugType: 0,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "latitude cannot be nil: wrong param value",
		},
		{
			name: "Longitude nil",
			msg: types.MsgPublishEnergyTransferOffer{
				Creator:   sample.AccAddress(),
				ChargerId: validChargerName,
				Tariff:    0,
				Location: &types.Location{
					Latitude:  &validLatitude,
					Longitude: nil,
				},
				Name:     validChargerName,
				PlugType: 0,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "longitude cannot be nil: wrong param value",
		},
		{
			name: "latitude to big",
			msg: types.MsgPublishEnergyTransferOffer{
				Creator:   sample.AccAddress(),
				ChargerId: validChargerName,
				Tariff:    0,
				Location: &types.Location{
					Latitude:  &invalidLatitude,
					Longitude: &validLongitude,
				},
				Name:     validChargerName,
				PlugType: 0,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "latitude must be between 90.000000000000000000 and -90.000000000000000000: wrong param value",
		},
		{
			name: "Longitude to big",
			msg: types.MsgPublishEnergyTransferOffer{
				Creator:   sample.AccAddress(),
				ChargerId: validChargerName,
				Tariff:    0,
				Location: &types.Location{
					Latitude:  &validLatitude,
					Longitude: &invalidLongitude,
				},
				Name:     validChargerName,
				PlugType: 0,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "longitude must be between 180.000000000000000000 and -180.000000000000000000: wrong param value",
		},
		{
			name: "location nil",
			msg: types.MsgPublishEnergyTransferOffer{
				Creator:   sample.AccAddress(),
				ChargerId: validChargerName,
				Tariff:    0,
				Location:  nil,
				Name:      validChargerName,
				PlugType:  0,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "charger location cannot be nil: wrong param value",
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
