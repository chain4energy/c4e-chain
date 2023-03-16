package types_test

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

var sampleAccAddress = sample.AccAddress()

func TestMsgSendToVestingAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgSendToVestingAccount
		err    error
		errMsg string
	}{
		{
			name: "empty vesting pool name",
			msg: types.MsgSendToVestingAccount{
				Owner:           sample.AccAddress(),
				ToAddress:       sample.AccAddress(),
				VestingPoolName: "",
				Amount:          math.NewInt(100),
			},
			err:    types.ErrParam,
			errMsg: "send to new vesting account - empty name: wrong param value",
		},
		{
			name: "nil amount",
			msg: types.MsgSendToVestingAccount{
				Owner:           sample.AccAddress(),
				ToAddress:       sample.AccAddress(),
				VestingPoolName: "pool",
				Amount:          math.Int{},
			},
			err:    types.ErrAmount,
			errMsg: "send to new vesting account - amount cannot be nil: wrong amount value",
		},
		{
			name: "negative amount",
			msg: types.MsgSendToVestingAccount{
				Owner:           sample.AccAddress(),
				ToAddress:       sample.AccAddress(),
				VestingPoolName: "pool",
				Amount:          math.NewInt(-100),
			},
			err:    types.ErrAmount,
			errMsg: "send to new vesting account - amount is <= 0: wrong amount value",
		},
		{
			name: "identical owner and to address",
			msg: types.MsgSendToVestingAccount{
				Owner:           sampleAccAddress,
				ToAddress:       sampleAccAddress,
				VestingPoolName: "pool",
				Amount:          math.NewInt(100),
			},
			err:    types.ErrIdenticalAccountsAddresses,
			errMsg: "send to new vesting account - identical from address (" + sampleAccAddress + ") and to address (" + sampleAccAddress + "): account addresses cannot be identical",
		},
		{
			name: "invalid owner address",
			msg: types.MsgSendToVestingAccount{
				Owner:           "invalidaddress",
				ToAddress:       sample.AccAddress(),
				VestingPoolName: "pool",
				Amount:          math.NewInt(100),
			},
			err:    types.ErrParsing,
			errMsg: "send to new vesting account - owner acc address error: decoding bech32 failed: invalid separator index -1: failed to parse",
		},
		{
			name: "invalid to address",
			msg: types.MsgSendToVestingAccount{
				Owner:           sample.AccAddress(),
				ToAddress:       "invalidaddress",
				VestingPoolName: "pool",
				Amount:          math.NewInt(100),
			},
			err:    types.ErrParsing,
			errMsg: "send to new vesting account - to acc address error: decoding bech32 failed: invalid separator index -1: failed to parse",
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
