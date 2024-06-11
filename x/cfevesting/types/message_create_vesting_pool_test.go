package types_test

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/testutil/sample"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateVestingPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgCreateVestingPool
		err    error
		errMsg string
	}{
		{
			name: "empty pool name",
			msg: types.MsgCreateVestingPool{
				Name:        "",
				Owner:       sample.AccAddress(),
				VestingType: "VestingType",
				Amount:      math.NewInt(100),
				Duration:    time.Hour,
			},
			err:    types.ErrParam,
			errMsg: "add vesting pool empty name: wrong param value",
		},
		{
			name: "amount is nil",
			msg: types.MsgCreateVestingPool{
				Name:        "Pool name",
				Owner:       sample.AccAddress(),
				VestingType: "VestingType",
				Amount:      math.Int{},
				Duration:    time.Hour,
			},
			err:    types.ErrAmount,
			errMsg: "add vesting pool - amount cannot be nil: wrong amount value",
		},
		{
			name: "amount is negative",
			msg: types.MsgCreateVestingPool{
				Name:        "Pool name",
				Owner:       sample.AccAddress(),
				VestingType: "VestingType",
				Amount:      math.NewInt(-100),
				Duration:    time.Hour,
			},
			err:    types.ErrAmount,
			errMsg: "add vesting pool - amount is <= 0: wrong amount value",
		},
		{
			name: "duration is 0",
			msg: types.MsgCreateVestingPool{
				Name:        "Pool name",
				Owner:       sample.AccAddress(),
				VestingType: "VestingType",
				Amount:      math.NewInt(100),
				Duration:    0,
			},
			err:    types.ErrParam,
			errMsg: "add vesting pool - duration is <= 0 or nil: wrong param value",
		},
		{
			name: "invalid owner address",
			msg: types.MsgCreateVestingPool{
				Name:        "Pool name",
				Owner:       "invalidaddress",
				VestingType: "VestingType",
				Amount:      math.NewInt(100),
				Duration:    time.Hour,
			},
			err:    types.ErrParsing,
			errMsg: "add vesting pool - vesting acc address error: decoding bech32 failed: invalid separator index -1: failed to parse",
		},
		{
			name: "correct message",
			msg: types.MsgCreateVestingPool{
				Name:        "Pool name",
				Owner:       sample.AccAddress(),
				VestingType: "VestingType",
				Amount:      math.NewInt(100),
				Duration:    time.Hour,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				require.EqualError(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
