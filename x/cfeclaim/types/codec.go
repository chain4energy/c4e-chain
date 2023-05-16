package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgClaim{}, "cfeclaim/Claim")
	legacy.RegisterAminoMsg(cdc, &MsgInitialClaim{}, "cfeclaim/InitialClaim")
	legacy.RegisterAminoMsg(cdc, &MsgCreateCampaign{}, "cfeclaim/CreateCampaign")
	legacy.RegisterAminoMsg(cdc, &MsgAddMissionToCampaign{}, "cfeclaim/AddMissionToCampaign")
	legacy.RegisterAminoMsg(cdc, &MsgAddClaimRecords{}, "cfeclaim/AddClaimRecords")
	legacy.RegisterAminoMsg(cdc, &MsgDeleteClaimRecord{}, "cfeclaim/DeleteClaimRecord")
	legacy.RegisterAminoMsg(cdc, &MsgCloseCampaign{}, "cfeclaim/CloseCampaign")
	legacy.RegisterAminoMsg(cdc, &MsgEnableCampaign{}, "cfeclaim/EnableCampaign")
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
		&MsgEnableCampaign{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterLegacyAminoCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
