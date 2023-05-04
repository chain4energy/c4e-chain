package types_test

import (
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/testutil/sample"

	"github.com/stretchr/testify/require"
)

func TestMsgCreateCampaign_ValidateBasic(t *testing.T) {
	startTime := time.Now().Add(time.Hour * 24 * 30) // start time 30 days from now
	endTime := startTime.Add(time.Hour * 24 * 30)    // end time 30 days from start time

	tests := []struct {
		name   string
		msg    types.MsgCreateCampaign
		err    error
		errMsg string
	}{
		{
			name: "empty name",
			msg: types.MsgCreateCampaign{
				Name:         "",
				Description:  "Valid description",
				StartTime:    &startTime,
				EndTime:      &endTime,
				CampaignType: types.CampaignDefault,
				Owner:        sample.AccAddress(),
			},
			err:    c4eerrors.ErrParam,
			errMsg: "campaign name is empty: wrong param value",
		},
		{
			name: "empty description",
			msg: types.MsgCreateCampaign{
				Name:         "Valid name",
				Description:  "",
				StartTime:    &startTime,
				EndTime:      &endTime,
				CampaignType: types.CampaignDefault,
				Owner:        sample.AccAddress(),
			},
			err:    c4eerrors.ErrParam,
			errMsg: "description is empty: wrong param value",
		},
		{
			name: "end time equal start time",
			msg: types.MsgCreateCampaign{
				Name:         "Valid name",
				Description:  "Valid description",
				StartTime:    &startTime,
				EndTime:      &startTime,
				CampaignType: types.CampaignDefault,
				Owner:        sample.AccAddress(),
			},
			err:    c4eerrors.ErrParam,
			errMsg: fmt.Sprintf("start time is equal to end time (%s = %s): wrong param value", startTime, startTime),
		},
		{
			name: "invalid campaign type",
			msg: types.MsgCreateCampaign{
				Name:         "Valid name",
				Description:  "Valid description",
				StartTime:    &startTime,
				EndTime:      &endTime,
				CampaignType: types.CampaignType_CAMPAIGN_TYPE_UNSPECIFIED,
				Owner:        sample.AccAddress(),
			},
			err:    sdkerrors.ErrInvalidType,
			errMsg: "wrong campaign type: invalid type",
		},
		{
			name: "invalid campaign type for Teamdrop",
			msg: types.MsgCreateCampaign{
				Name:         "Valid name",
				Description:  "Valid description",
				StartTime:    &startTime,
				EndTime:      &endTime,
				CampaignType: types.CampaignTeamdrop,
				Owner:        sample.AccAddress(),
			},
			err:    sdkerrors.ErrorInvalidSigner,
			errMsg: "TeamDrop campaigns can be created only by specific accounts: tx intended signer does not match the given signer",
		},
		{
			name: "valid campaign",
			msg: types.MsgCreateCampaign{
				Name:         "Valid name",
				Description:  "Valid description",
				StartTime:    &startTime,
				EndTime:      &endTime,
				CampaignType: types.CampaignDefault,
				Owner:        sample.AccAddress(),
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
