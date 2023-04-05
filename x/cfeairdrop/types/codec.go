package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgClaim{}, "cfeairdrop/Claim", nil)
	cdc.RegisterConcrete(&MsgInitialClaim{}, "cfeairdrop/InitialClaim", nil)
	cdc.RegisterConcrete(&MsgCreateCampaign{}, "cfeairdrop/CreateCampaign", nil)
	cdc.RegisterConcrete(&MsgAddMissionToCampaign{}, "cfeairdrop/AddMissionToCampaign", nil)
	cdc.RegisterConcrete(&MsgAddClaimRecords{}, "cfeairdrop/AddUsersEntries", nil)
	cdc.RegisterConcrete(&MsgDeleteClaimRecord{}, "cfeairdrop/DeleteClaimRecord", nil)
	cdc.RegisterConcrete(&MsgCloseCampaign{}, "cfeairdrop/CloseCampaign", nil)
	cdc.RegisterConcrete(&MsgStartCampaign{}, "cfeairdrop/StartCampaign", nil)
	cdc.RegisterConcrete(&MsgEditCampaign{}, "cfeairdrop/EditCampaign", nil)
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
