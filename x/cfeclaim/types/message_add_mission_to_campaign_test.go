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

func TestMsgAddMissionToCampaign_ValidateBasic(t *testing.T) {
	correctWeight := sdk.MustNewDecFromStr("0.5")
	incorrectWeight := sdk.MustNewDecFromStr("1.5")
	tests := []struct {
		name   string
		msg    types.MsgAddMissionToCampaign
		err    error
		errMsg string
	}{
		{
			name: "invalid owner address",
			msg: types.MsgAddMissionToCampaign{
				Owner:       "invalid_address",
				Name:        "Test Mission",
				Description: "Test Mission Description",
				MissionType: types.MissionClaim,
				Weight:      &correctWeight,
			},
			err:    sdkerrors.ErrInvalidAddress,
			errMsg: "invalid owner address (decoding bech32 failed: invalid separator index -1): invalid address",
		},
		{
			name: "invalid mission weight",
			msg: types.MsgAddMissionToCampaign{
				Owner:       sample.AccAddress(),
				Name:        "Test Mission",
				Description: "Test Mission Description",
				MissionType: types.MissionClaim,
				Weight:      &incorrectWeight,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "weight (1.500000000000000000) is not between 0 and 1 error: wrong param value",
		},
		{
			name: "invalid mission weight (nil)",
			msg: types.MsgAddMissionToCampaign{
				Owner:       sample.AccAddress(),
				Name:        "Test Mission",
				Description: "Test Mission Description",
				MissionType: types.MissionClaim,
				Weight:      nil,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "add mission to claim campaign weight is nil error: wrong param value",
		},
		{
			name: "empty mission name",
			msg: types.MsgAddMissionToCampaign{
				Owner:       sample.AccAddress(),
				Name:        "",
				Description: "Test Mission Description",
				MissionType: types.MissionClaim,
				Weight:      &correctWeight,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "empty name error: wrong param value",
		},
		{
			name: "empty mission description",
			msg: types.MsgAddMissionToCampaign{
				Owner:       sample.AccAddress(),
				Name:        "Test Mission",
				Description: "",
				MissionType: types.MissionClaim,
				Weight:      &correctWeight,
			},
			err:    c4eerrors.ErrParam,
			errMsg: "mission empty description error: wrong param value",
		},
		{
			name: "invalid mission type",
			msg: types.MsgAddMissionToCampaign{
				Owner:       sample.AccAddress(),
				Name:        "Test Mission",
				Description: "Test Mission Description",
				MissionType: types.MissionType_MISSION_TYPE_UNSPECIFIED,
				Weight:      &correctWeight,
			},
			err:    sdkerrors.ErrInvalidType,
			errMsg: "wrong mission type: invalid type",
		},
		{
			name: "valid mission",
			msg: types.MsgAddMissionToCampaign{
				Owner:       sample.AccAddress(),
				Name:        "Test Mission",
				Description: "Test Mission Description",
				MissionType: types.MissionClaim,
				Weight:      &correctWeight,
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
