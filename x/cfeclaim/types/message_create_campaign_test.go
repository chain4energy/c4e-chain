package types_test

import (
	"cosmossdk.io/math"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/v2/types/errors"
	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/v2/testutil/sample"

	"github.com/stretchr/testify/require"
)

const (
	validDescription = "Valid description"
	validName        = "Valid name"
)

var (
	startTime    = time.Now().Add(time.Hour * 24 * 30) // start time 30 days from now
	endTime      = startTime.Add(time.Hour * 24 * 30)  // end time 30 days from start time
	zeroInt      = math.ZeroInt()
	zeroDec      = sdk.ZeroDec()
	zeroDuration = time.Duration(0)
)

func TestMsgCreateCampaign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgCreateCampaign
		err    error
		errMsg string
	}{
		{
			name: "empty name",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   "",
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "campaign name is empty: wrong param value",
		},
		{
			name: "empty description",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            "",
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
		},
		{
			name: "end time equal start time",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &startTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: fmt.Sprintf("start time is equal to end time (%s = %s): wrong param value", startTime, startTime),
		},
		{
			name: "invalid campaign type",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.CampaignType_CAMPAIGN_TYPE_UNSPECIFIED,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    sdkerrors.ErrInvalidType,
			errMsg: "wrong campaign type: invalid type",
		},
		{
			name: "invalid vesting pool campaign - no vesting pool name",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.VestingPoolCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "for VESTING_POOL type campaigns, the vesting pool name must be provided: wrong param value",
		},
		{
			name: "invalid default pool campaign - vesting pool name set",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "abcd",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "vesting pool name can be set only for VESTING_POOL type campaigns: wrong param value",
		},
		{
			name: "valid vesting pool campaign",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.VestingPoolCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "abcd",
			},
		},
		{
			name: "valid campaign",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
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

func TestMsgCreateCampaignNilValues_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgCreateCampaign
		err    error
		errMsg string
	}{
		{
			name: "nil feegrant",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         nil,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "feegrant amount cannot be nil: wrong param value",
		},
		{
			name: "nil InitialClaimFreeAmount",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: nil,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "initital claim free amount cannot be nil: wrong param value",
		},
		{
			name: "nil Free",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   nil,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "free decimal cannot be nil: wrong param value",
		},
		{
			name: "nil start time",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              nil,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "start time cannot be nil: wrong param value",
		},
		{
			name: "nil end time",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                nil,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "end time cannot be nil: wrong param value",
		},
		{
			name: "nil lockup period",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           nil,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "lockup period cannot be nil: wrong param value",
		},
		{
			name: "nil vesting period",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          nil,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "vesting period cannot be nil: wrong param value",
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

func TestMsgCreateCampaignWrongValues_ValidateBasic(t *testing.T) {
	negativeInt := math.ZeroInt().SubRaw(1)
	negativeDec := sdk.ZeroDec().Sub(sdk.NewDec(1))
	negativeDuration := time.Duration(-100)
	tests := []struct {
		name   string
		msg    types.MsgCreateCampaign
		err    error
		errMsg string
	}{
		{
			name: "negative feegrant",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &negativeInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "feegrant amount (-1) cannot be negative: wrong param value",
		},
		{
			name: "negative InitialClaimFreeAmount",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &negativeInt,
				Free:                   &zeroDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "initial claim free amount (-1) cannot be negative: wrong param value",
		},
		{
			name: "negative free",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &negativeDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "free amount (-1.000000000000000000) cannot be negative: wrong param value",
		},
		{
			name: "negative vesting period",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &negativeDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &zeroDuration,
				VestingPeriod:          &negativeDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "vesting period cannot be negative: wrong param value",
		},
		{
			name: "negative lockup period",
			msg: types.MsgCreateCampaign{
				Owner:                  sample.AccAddress(),
				Name:                   validName,
				Description:            validDescription,
				CampaignType:           types.DefaultCampaign,
				RemovableClaimRecords:  false,
				FeegrantAmount:         &zeroInt,
				InitialClaimFreeAmount: &zeroInt,
				Free:                   &negativeDec,
				StartTime:              &startTime,
				EndTime:                &endTime,
				LockupPeriod:           &negativeDuration,
				VestingPeriod:          &zeroDuration,
				VestingPoolName:        "",
			},
			err:    c4eerrors.ErrParam,
			errMsg: "lockup period cannot be negative: wrong param value",
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
