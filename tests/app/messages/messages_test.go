package messages_test

import (
	"fmt"
	"time"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"

	"testing"

	cfeclaimmoduletypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfedistributormoduletypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfemintermoduletypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfevestingmoduletypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var accountsAddresses, _ = testcosmos.CreateAccounts(22, 0)

type MsgCommon interface {
	Route() string
	Type() string
	GetSigners() []sdk.AccAddress

	GetSignBytes() []byte
}

func TestMsgMethods(t *testing.T) {
	tests := []struct {
		name            string
		msg             MsgCommon
		moduleCdc       codec.Codec
		expectedRoute   string
		expectedType    string
		expectedSigners []sdk.AccAddress
		errorMessage    string
	}{
		// ---- cfeclaim module
		{
			name:            fmt.Sprintf("message %T test", cfeclaimmoduletypes.MsgAddMission{}),
			msg:             cfeclaimmoduletypes.NewMsgAddMission(accountsAddresses[0].String(), 3, "name", "desc", cfeclaimmoduletypes.MissionType_INITIAL_CLAIM, &sdk.Dec{}, &time.Time{}),
			moduleCdc:       cfeclaimmoduletypes.ModuleCdc,
			expectedRoute:   cfeclaimmoduletypes.RouterKey,
			expectedType:    cfeclaimmoduletypes.TypeMsgAddMission,
			expectedSigners: []sdk.AccAddress{accountsAddresses[0]},
			errorMessage:    fmt.Sprintf("message %T error", cfeclaimmoduletypes.MsgAddMission{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfeclaimmoduletypes.MsgClaim{}),
			msg:             cfeclaimmoduletypes.NewMsgClaim(accountsAddresses[1].String(), 3, 5),
			moduleCdc:       cfeclaimmoduletypes.ModuleCdc,
			expectedRoute:   cfeclaimmoduletypes.RouterKey,
			expectedType:    cfeclaimmoduletypes.TypeMsgClaim,
			expectedSigners: []sdk.AccAddress{accountsAddresses[1]},
			errorMessage:    fmt.Sprintf("message %T error", cfeclaimmoduletypes.MsgClaim{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfeclaimmoduletypes.MsgCloseCampaign{}),
			msg:             cfeclaimmoduletypes.NewMsgCloseCampaign(accountsAddresses[2].String(), 3),
			moduleCdc:       cfeclaimmoduletypes.ModuleCdc,
			expectedRoute:   cfeclaimmoduletypes.RouterKey,
			expectedType:    cfeclaimmoduletypes.TypeMsgCloseCampaign,
			expectedSigners: []sdk.AccAddress{accountsAddresses[2]},
			errorMessage:    fmt.Sprintf("message %T error", cfeclaimmoduletypes.MsgCloseCampaign{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfeclaimmoduletypes.MsgCreateCampaign{}),
			msg:             cfeclaimmoduletypes.NewMsgCreateCampaign(accountsAddresses[2].String(), "3", "", cfeclaimmoduletypes.CampaignType_DEFAULT, true, nil, nil, nil, nil, nil, nil, nil, ""),
			moduleCdc:       cfeclaimmoduletypes.ModuleCdc,
			expectedRoute:   cfeclaimmoduletypes.RouterKey,
			expectedType:    cfeclaimmoduletypes.TypeMsgCreateCampaign,
			expectedSigners: []sdk.AccAddress{accountsAddresses[2]},
			errorMessage:    fmt.Sprintf("message %T error", cfeclaimmoduletypes.MsgCreateCampaign{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfeclaimmoduletypes.MsgEnableCampaign{}),
			msg:             cfeclaimmoduletypes.NewMsgEnableCampaign(accountsAddresses[3].String(), 3, nil, nil),
			moduleCdc:       cfeclaimmoduletypes.ModuleCdc,
			expectedRoute:   cfeclaimmoduletypes.RouterKey,
			expectedType:    cfeclaimmoduletypes.TypeMsgEnableCampaign,
			expectedSigners: []sdk.AccAddress{accountsAddresses[3]},
			errorMessage:    fmt.Sprintf("message %T error", cfeclaimmoduletypes.MsgEnableCampaign{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfeclaimmoduletypes.MsgInitialClaim{}),
			msg:             cfeclaimmoduletypes.NewMsgInitialClaim(accountsAddresses[4].String(), 3, ""),
			moduleCdc:       cfeclaimmoduletypes.ModuleCdc,
			expectedRoute:   cfeclaimmoduletypes.RouterKey,
			expectedType:    cfeclaimmoduletypes.TypeMsgInitialClaim,
			expectedSigners: []sdk.AccAddress{accountsAddresses[4]},
			errorMessage:    fmt.Sprintf("message %T error", cfeclaimmoduletypes.MsgInitialClaim{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfeclaimmoduletypes.MsgRemoveCampaign{}),
			msg:             cfeclaimmoduletypes.NewMsgRemoveCampaign(accountsAddresses[5].String(), 3),
			moduleCdc:       cfeclaimmoduletypes.ModuleCdc,
			expectedRoute:   cfeclaimmoduletypes.RouterKey,
			expectedType:    cfeclaimmoduletypes.TypeMsgRemoveCampaign,
			expectedSigners: []sdk.AccAddress{accountsAddresses[5]},
			errorMessage:    fmt.Sprintf("message %T error", cfeclaimmoduletypes.MsgRemoveCampaign{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfeclaimmoduletypes.MsgAddClaimRecords{}),
			msg:             cfeclaimmoduletypes.NewMsgAddClaimRecords(accountsAddresses[6].String(), 3, nil),
			moduleCdc:       cfeclaimmoduletypes.ModuleCdc,
			expectedRoute:   cfeclaimmoduletypes.RouterKey,
			expectedType:    cfeclaimmoduletypes.TypeMsgAddClaimRecords,
			expectedSigners: []sdk.AccAddress{accountsAddresses[6]},
			errorMessage:    fmt.Sprintf("message %T error", cfeclaimmoduletypes.MsgAddClaimRecords{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfeclaimmoduletypes.MsgDeleteClaimRecord{}),
			msg:             cfeclaimmoduletypes.NewMsgDeleteClaimRecord(accountsAddresses[7].String(), 3, ""),
			moduleCdc:       cfeclaimmoduletypes.ModuleCdc,
			expectedRoute:   cfeclaimmoduletypes.RouterKey,
			expectedType:    cfeclaimmoduletypes.TypeMsgDeleteClaimRecord,
			expectedSigners: []sdk.AccAddress{accountsAddresses[7]},
			errorMessage:    fmt.Sprintf("message %T error", cfeclaimmoduletypes.MsgDeleteClaimRecord{}),
		},
		// ---- cfedistributor module
		{
			name:            fmt.Sprintf("message %T test", cfedistributormoduletypes.MsgUpdateParams{}),
			msg:             &cfedistributormoduletypes.MsgUpdateParams{accountsAddresses[8].String(), nil},
			moduleCdc:       cfedistributormoduletypes.ModuleCdc,
			expectedRoute:   cfedistributormoduletypes.RouterKey,
			expectedType:    cfedistributormoduletypes.TypeMsgUpdateParams,
			expectedSigners: []sdk.AccAddress{accountsAddresses[8]},
			errorMessage:    fmt.Sprintf("message %T error", cfedistributormoduletypes.MsgUpdateParams{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfedistributormoduletypes.MsgUpdateSubDistributorBurnShareParam{}),
			msg:             &cfedistributormoduletypes.MsgUpdateSubDistributorBurnShareParam{accountsAddresses[9].String(), "", sdk.ZeroDec()},
			moduleCdc:       cfedistributormoduletypes.ModuleCdc,
			expectedRoute:   cfedistributormoduletypes.RouterKey,
			expectedType:    cfedistributormoduletypes.TypeMsgUpdateSubDistributorBurnShareParam,
			expectedSigners: []sdk.AccAddress{accountsAddresses[9]},
			errorMessage:    fmt.Sprintf("message %T error", cfedistributormoduletypes.MsgUpdateSubDistributorBurnShareParam{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfedistributormoduletypes.MsgUpdateSubDistributorDestinationShareParam{}),
			msg:             &cfedistributormoduletypes.MsgUpdateSubDistributorDestinationShareParam{accountsAddresses[10].String(), "", "", sdk.ZeroDec()},
			moduleCdc:       cfedistributormoduletypes.ModuleCdc,
			expectedRoute:   cfedistributormoduletypes.RouterKey,
			expectedType:    cfedistributormoduletypes.TypeMsgUpdateSubDistributorDestinationShareParam,
			expectedSigners: []sdk.AccAddress{accountsAddresses[10]},
			errorMessage:    fmt.Sprintf("message %T error", cfedistributormoduletypes.MsgUpdateSubDistributorDestinationShareParam{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfedistributormoduletypes.MsgUpdateSubDistributorParam{}),
			msg:             &cfedistributormoduletypes.MsgUpdateSubDistributorParam{accountsAddresses[11].String(), nil},
			moduleCdc:       cfedistributormoduletypes.ModuleCdc,
			expectedRoute:   cfedistributormoduletypes.RouterKey,
			expectedType:    cfedistributormoduletypes.TypeMsgUpdateSubDistributorParam,
			expectedSigners: []sdk.AccAddress{accountsAddresses[11]},
			errorMessage:    fmt.Sprintf("message %T error", cfedistributormoduletypes.MsgUpdateSubDistributorParam{}),
		},
		// ---- cfeminter module
		{
			name:            fmt.Sprintf("message %T test", cfemintermoduletypes.MsgUpdateParams{}),
			msg:             &cfemintermoduletypes.MsgUpdateParams{accountsAddresses[12].String(), "", time.Now(), nil},
			moduleCdc:       cfemintermoduletypes.ModuleCdc,
			expectedRoute:   cfemintermoduletypes.RouterKey,
			expectedType:    cfemintermoduletypes.TypeMsgUpdateParams,
			expectedSigners: []sdk.AccAddress{accountsAddresses[12]},
			errorMessage:    fmt.Sprintf("message %T error", cfemintermoduletypes.MsgUpdateParams{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfemintermoduletypes.MsgUpdateMintersParams{}),
			msg:             &cfemintermoduletypes.MsgUpdateMintersParams{accountsAddresses[12].String(), time.Now(), nil},
			moduleCdc:       cfemintermoduletypes.ModuleCdc,
			expectedRoute:   cfemintermoduletypes.RouterKey,
			expectedType:    cfemintermoduletypes.TypeMsgUpdateMintersParams,
			expectedSigners: []sdk.AccAddress{accountsAddresses[12]},
			errorMessage:    fmt.Sprintf("message %T error", cfemintermoduletypes.MsgUpdateMintersParams{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfemintermoduletypes.MsgBurn{}),
			msg:             cfemintermoduletypes.NewMsgBurn(accountsAddresses[13].String(), sdk.NewCoins()),
			moduleCdc:       cfemintermoduletypes.ModuleCdc,
			expectedRoute:   cfemintermoduletypes.RouterKey,
			expectedType:    cfemintermoduletypes.TypeMsgBurn,
			expectedSigners: []sdk.AccAddress{accountsAddresses[13]},
			errorMessage:    fmt.Sprintf("message %T error", cfemintermoduletypes.MsgBurn{}),
		},
		// ---- cfevesting module
		{
			name:            fmt.Sprintf("message %T test", cfevestingmoduletypes.MsgCreateVestingAccount{}),
			msg:             cfevestingmoduletypes.NewMsgCreateVestingAccount(accountsAddresses[14].String(), "", sdk.NewCoins(), 5, 6),
			moduleCdc:       cfevestingmoduletypes.ModuleCdc,
			expectedRoute:   cfevestingmoduletypes.RouterKey,
			expectedType:    cfevestingmoduletypes.TypeMsgCreateVestingAccount,
			expectedSigners: []sdk.AccAddress{accountsAddresses[14]},
			errorMessage:    fmt.Sprintf("message %T error", cfevestingmoduletypes.MsgCreateVestingAccount{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfevestingmoduletypes.MsgCreateVestingPool{}),
			msg:             cfevestingmoduletypes.NewMsgCreateVestingPool(accountsAddresses[15].String(), "", sdk.ZeroInt(), 5, ""),
			moduleCdc:       cfevestingmoduletypes.ModuleCdc,
			expectedRoute:   cfevestingmoduletypes.RouterKey,
			expectedType:    cfevestingmoduletypes.TypeMsgCreateVestingPool,
			expectedSigners: []sdk.AccAddress{accountsAddresses[15]},
			errorMessage:    fmt.Sprintf("message %T error", cfevestingmoduletypes.MsgCreateVestingPool{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfevestingmoduletypes.MsgMoveAvailableVesting{}),
			msg:             cfevestingmoduletypes.NewMsgMoveAvailableVesting(accountsAddresses[16].String(), ""),
			moduleCdc:       cfevestingmoduletypes.ModuleCdc,
			expectedRoute:   cfevestingmoduletypes.RouterKey,
			expectedType:    cfevestingmoduletypes.TypeMsgMoveAvailableVesting,
			expectedSigners: []sdk.AccAddress{accountsAddresses[16]},
			errorMessage:    fmt.Sprintf("message %T error", cfevestingmoduletypes.MsgMoveAvailableVesting{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfevestingmoduletypes.MsgMoveAvailableVestingByDenoms{}),
			msg:             cfevestingmoduletypes.NewMsgMoveAvailableVestingByDenoms(accountsAddresses[17].String(), "", nil),
			moduleCdc:       cfevestingmoduletypes.ModuleCdc,
			expectedRoute:   cfevestingmoduletypes.RouterKey,
			expectedType:    cfevestingmoduletypes.TypeMsgMoveAvailableVestingByDenoms,
			expectedSigners: []sdk.AccAddress{accountsAddresses[17]},
			errorMessage:    fmt.Sprintf("message %T error", cfevestingmoduletypes.MsgMoveAvailableVestingByDenoms{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfevestingmoduletypes.MsgSendToVestingAccount{}),
			msg:             cfevestingmoduletypes.NewMsgSendToVestingAccount(accountsAddresses[18].String(), "", "", sdk.ZeroInt(), true),
			moduleCdc:       cfevestingmoduletypes.ModuleCdc,
			expectedRoute:   cfevestingmoduletypes.RouterKey,
			expectedType:    cfevestingmoduletypes.TypeMsgSendToVestingAccount,
			expectedSigners: []sdk.AccAddress{accountsAddresses[18]},
			errorMessage:    fmt.Sprintf("message %T error", cfevestingmoduletypes.MsgSendToVestingAccount{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfevestingmoduletypes.MsgSplitVesting{}),
			msg:             cfevestingmoduletypes.NewMsgSplitVesting(accountsAddresses[19].String(), "", sdk.NewCoins()),
			moduleCdc:       cfevestingmoduletypes.ModuleCdc,
			expectedRoute:   cfevestingmoduletypes.RouterKey,
			expectedType:    cfevestingmoduletypes.TypeMsgSplitVesting,
			expectedSigners: []sdk.AccAddress{accountsAddresses[19]},
			errorMessage:    fmt.Sprintf("message %T error", cfevestingmoduletypes.MsgSplitVesting{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfevestingmoduletypes.MsgUpdateDenomParam{}),
			msg:             &cfevestingmoduletypes.MsgUpdateDenomParam{accountsAddresses[20].String(), ""},
			moduleCdc:       cfevestingmoduletypes.ModuleCdc,
			expectedRoute:   cfevestingmoduletypes.RouterKey,
			expectedType:    cfevestingmoduletypes.TypeMsgUpdateDenomParam,
			expectedSigners: []sdk.AccAddress{accountsAddresses[20]},
			errorMessage:    fmt.Sprintf("message %T error", cfevestingmoduletypes.MsgUpdateDenomParam{}),
		},
		{
			name:            fmt.Sprintf("message %T test", cfevestingmoduletypes.MsgWithdrawAllAvailable{}),
			msg:             cfevestingmoduletypes.NewMsgWithdrawAllAvailable(accountsAddresses[21].String()),
			moduleCdc:       cfevestingmoduletypes.ModuleCdc,
			expectedRoute:   cfevestingmoduletypes.RouterKey,
			expectedType:    cfevestingmoduletypes.TypeMsgWithdrawAllAvailable,
			expectedSigners: []sdk.AccAddress{accountsAddresses[21]},
			errorMessage:    fmt.Sprintf("message %T error", cfevestingmoduletypes.MsgWithdrawAllAvailable{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.EqualValues(t, tt.expectedRoute, tt.msg.Route(), tt.errorMessage+": wrong route")
			require.EqualValues(t, tt.expectedType, tt.msg.Type(), tt.errorMessage+": wrong type")
			require.ElementsMatch(t, tt.expectedSigners, tt.msg.GetSigners(), tt.errorMessage+": wrong signers")

			bz := tt.moduleCdc.MustMarshalJSON(tt.msg.(proto.Message))
			expectedSignBytes := sdk.MustSortJSON(bz)
			require.ElementsMatch(t, expectedSignBytes, tt.msg.GetSignBytes(), tt.errorMessage+": wrong signers")
		})
	}
}