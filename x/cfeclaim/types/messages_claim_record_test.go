package types_test

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAddClaimRecords_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgAddClaimRecords
		err    error
		errMsg string
	}{
		{
			name: "Valid MsgAddClaimRecords",
			msg: types.MsgAddClaimRecords{
				Owner:              sample.AccAddress(),
				ClaimRecordEntries: []*types.ClaimRecordEntry{{UserEntryAddress: sample.AccAddress(), Amount: sample.Coins()}},
			},
		},
		{
			name: "Invalid Owner",
			msg: types.MsgAddClaimRecords{
				Owner:              "invalid_address",
				ClaimRecordEntries: []*types.ClaimRecordEntry{{UserEntryAddress: sample.AccAddress(), Amount: sample.Coins()}},
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid owner address (decoding bech32 failed: invalid separator index -1): invalid address",
		},
		{
			name: "Invalid ClaimRecord Address",
			msg: types.MsgAddClaimRecords{
				Owner:              sample.AccAddress(),
				ClaimRecordEntries: []*types.ClaimRecordEntry{{UserEntryAddress: "", Amount: sample.Coins()}},
			},
			err:    c4eerrors.ErrParam,
			errMsg: "claim record entry index 0: claim record entry empty user entry address: wrong param value",
		},
		{
			name: "Invalid ClaimRecord Amount",
			msg: types.MsgAddClaimRecords{
				Owner:              sample.AccAddress(),
				ClaimRecordEntries: []*types.ClaimRecordEntry{{UserEntryAddress: sample.AccAddress(), Amount: sdk.Coins{}}},
			},
			err:    c4eerrors.ErrParam,
			errMsg: "claim record entry index 0: claim record entry must has at least one coin and all amounts must be positive: wrong param value",
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

func TestMsgDeleteClaimRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgDeleteClaimRecord
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: types.MsgDeleteClaimRecord{
				Owner: "invalid_address",
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid owner address (decoding bech32 failed: invalid separator index -1): invalid address",
		}, {
			name: "valid address",
			msg: types.MsgDeleteClaimRecord{
				Owner: sample.AccAddress(),
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
