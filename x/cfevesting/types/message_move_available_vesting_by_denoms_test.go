package types_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"

	"github.com/chain4energy/c4e-chain/v2/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgMoveAvailableVestingByDenoms_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgMoveAvailableVestingByDenoms
		err    error
		errMsg string
	}{
		{
			name: "invalid from address",
			msg: types.MsgMoveAvailableVestingByDenoms{
				FromAddress: "invalid_address",
				ToAddress:   sample.AccAddress(),
				Denoms:      []string{"denom1", "denom2"},
			},
			err:    types.ErrParsing,
			errMsg: "move available vesting by denoms: from acc address error: decoding bech32 failed: invalid separator index -1: failed to parse",
		},
		{
			name: "invalid to address",
			msg: types.MsgMoveAvailableVestingByDenoms{
				FromAddress: sample.AccAddress(),
				ToAddress:   "invalid_address",
				Denoms:      []string{"denom1", "denom2"},
			},
			err:    types.ErrParsing,
			errMsg: "move available vesting by denoms: to acc address error: decoding bech32 failed: invalid separator index -1: failed to parse",
		},
		{
			name: "invalid denoms - no denoms",
			msg: types.MsgMoveAvailableVestingByDenoms{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
			},
			err:    types.ErrParam,
			errMsg: "move available vesting by denoms - no denominations: wrong param value",
		},
		{
			name: "invalid denoms - empty denoms",
			msg: types.MsgMoveAvailableVestingByDenoms{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Denoms:      []string{},
			},
			err:    types.ErrParam,
			errMsg: "move available vesting by denoms - no denominations: wrong param value",
		},
		{
			name: "invalid denoms - empty denom only",
			msg: types.MsgMoveAvailableVestingByDenoms{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Denoms:      []string{""},
			},
			err:    types.ErrParam,
			errMsg: "move available vesting by denoms - empty denomination at position 0: wrong param value",
		},
		{
			name: "invalid denoms - empty denom at pos 0",
			msg: types.MsgMoveAvailableVestingByDenoms{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Denoms:      []string{"", "denom"},
			},
			err:    types.ErrParam,
			errMsg: "move available vesting by denoms - empty denomination at position 0: wrong param value",
		},
		{
			name: "invalid denoms - empty denom at pos 0",
			msg: types.MsgMoveAvailableVestingByDenoms{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Denoms:      []string{"denom1", "denom2", "", "denom3"},
			},
			err:    types.ErrParam,
			errMsg: "move available vesting by denoms - empty denomination at position 2: wrong param value",
		},
		{
			name: "invalid denoms - duplicated",
			msg: types.MsgMoveAvailableVestingByDenoms{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Denoms:      []string{"denom1", "denom2", "denom2", "denom3"},
			},
			err:    types.ErrParam,
			errMsg: "move available vesting by denoms - duplicate denomination denom2: wrong param value",
		},
		{
			name: "invalid denoms - first duplicated",
			msg: types.MsgMoveAvailableVestingByDenoms{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Denoms:      []string{"denom1", "denom2", "denom1", "denom3"},
			},
			err:    types.ErrParam,
			errMsg: "move available vesting by denoms - duplicate denomination denom1: wrong param value",
		},
		{
			name: "valid address",
			msg: types.MsgMoveAvailableVestingByDenoms{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Denoms:      []string{"denom1", "denom2"},
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
