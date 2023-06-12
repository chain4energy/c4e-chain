package types_test

import (
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/testutil/sample"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgBurn_ValidateBasic(t *testing.T) {
	correctAmount := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(100)))
	tests := []struct {
		name   string
		msg    types.MsgBurn
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgBurn{
				Address: "abcd",
				Amount:  correctAmount,
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid address (decoding bech32 failed: invalid bech32 string length 4): invalid address",
		},
		{
			name: "nil amount",
			msg: types.MsgBurn{
				Address: sample.AccAddress(),
				Amount:  nil,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "amount is nil: wrong param value",
		},
		{
			name: "empty amount",
			msg: types.MsgBurn{
				Address: sample.AccAddress(),
				Amount:  sdk.NewCoins(),
			},
			err:    c4eerrors.ErrParam,
			errMsg: "amount is not positive: wrong param value",
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
