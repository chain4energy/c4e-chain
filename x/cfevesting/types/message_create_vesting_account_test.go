package types_test

import (
	"cosmossdk.io/math"
	testenv "github.com/chain4energy/c4e-chain/v2/testutil/env"
	"github.com/chain4energy/c4e-chain/v2/testutil/sample"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgCreateVestingAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgCreateVestingAccount
		err    error
		errMsg string
	}{
		{
			name: "invalid from address",
			msg: types.MsgCreateVestingAccount{
				FromAddress: "invalidaddress",
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000))),
				StartTime:   0,
				EndTime:     0,
			},
			err:    types.ErrParsing,
			errMsg: "create vesting account - from-address parsing error: invalidaddress: decoding bech32 failed: invalid separator index -1: failed to parse",
		},
		{
			name: "invalid to address",
			msg: types.MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   "invalidaddress",
				Amount:      sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000))),
				StartTime:   0,
				EndTime:     0,
			},
			err:    types.ErrParsing,
			errMsg: "create vesting account - to-address parsing error: invalidaddress: decoding bech32 failed: invalid separator index -1: failed to parse",
		},
		{
			name: "start time is after end time",
			msg: types.MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000))),
				StartTime:   100,
				EndTime:     0,
			},
			err:    types.ErrParam,
			errMsg: "create vesting account - start time is after end time error (1970-01-01 01:01:40 +0100 CET > 1970-01-01 01:00:00 +0100 CET): wrong param value",
		},
		{
			name: "amount is nil",
			msg: types.MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Amount:      nil,
				StartTime:   0,
				EndTime:     0,
			},
			err:    types.ErrParam,
			errMsg: "create vesting account - coin amount cannot be nil: wrong param value",
		},
		{
			name: "negative amount",
			msg: types.MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.Coins{sdk.Coin{Denom: testenv.DefaultTestDenom, Amount: sdk.NewInt(-10000)}},
				StartTime:   0,
				EndTime:     0,
			},
			err:    types.ErrParam,
			errMsg: "create vesting account - invalid amount (coin -10000uc4e amount is not positive): wrong param value",
		},
		{
			name: "nil coin",
			msg: types.MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.Coins{sdk.Coin{Denom: testenv.DefaultTestDenom, Amount: sdk.NewInt(10000)}, sdk.Coin{}},
				StartTime:   0,
				EndTime:     0,
			},
			err:    types.ErrParam,
			errMsg: "create vesting account - coin amount cannot be nil: wrong param value",
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
