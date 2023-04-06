package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgClaim{}, "cfeclaim/Claim", nil)
	cdc.RegisterConcrete(&MsgInitialClaim{}, "cfeclaim/InitialClaim", nil)
	cdc.RegisterConcrete(&MsgCreateCampaign{}, "cfeclaim/CreateCampaign", nil)
	cdc.RegisterConcrete(&MsgAddMissionToCampaign{}, "cfeclaim/AddMissionToCampaign", nil)
	cdc.RegisterConcrete(&MsgAddClaimRecords{}, "cfeclaim/AddUsersEntries", nil)
	cdc.RegisterConcrete(&MsgDeleteClaimRecord{}, "cfeclaim/DeleteClaimRecord", nil)
	cdc.RegisterConcrete(&MsgCloseCampaign{}, "cfeclaim/CloseCampaign", nil)
	cdc.RegisterConcrete(&MsgStartCampaign{}, "cfeclaim/StartCampaign", nil)
	cdc.RegisterConcrete(&MsgEditCampaign{}, "cfeclaim/EditCampaign", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaim{},
		&MsgInitialClaim{},
		&MsgCreateCampaign{},
		&MsgAddMissionToCampaign{},
		&MsgAddClaimRecords{},
		&MsgDeleteClaimRecord{},
		&MsgCloseCampaign{},
		&MsgStartCampaign{},
		&MsgEditCampaign{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
