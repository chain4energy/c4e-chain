package types_test

import (
	"testing"

	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/testutil/sample"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSplitVesting_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgSplitVesting
		err    error
		errMsg string
	}{
		{
			name: "invalid from address",
			msg: types.MsgSplitVesting{
				FromAddress: "invalid_address",
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(2))),
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid fromAddress address (decoding bech32 failed: invalid separator index -1): invalid address",
		},
		{
			name: "invalid to address",
			msg: types.MsgSplitVesting{
				FromAddress: sample.AccAddress(),
				ToAddress:   "invalid_address",
				Amount:      sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(2))),
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid toAddress address (decoding bech32 failed: invalid separator index -1): invalid address",
		},
		{
			name: "invalid Amount",
			msg: types.MsgSplitVesting{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.Coins{sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(2)), sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(2))},
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid amount (duplicate denomination uc4e): invalid address",
		},
		{
			name: "valid address",
			msg: types.MsgSplitVesting{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(2))),
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
