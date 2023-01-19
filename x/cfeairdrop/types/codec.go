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
	cdc.RegisterConcrete(&MsgCreateAirdropCampaign{}, "cfeairdrop/CreateAirdropCampaign", nil)
	cdc.RegisterConcrete(&MsgAddMissionToAidropCampaign{}, "cfeairdrop/AddMissionToAidropCampaign", nil)
	cdc.RegisterConcrete(&MsgAddAirdropEntries{}, "cfeairdrop/AddUserAirdropEntries", nil)
	cdc.RegisterConcrete(&MsgDeleteAirdropEntry{}, "cfeairdrop/DeleteAirdropEntry", nil)
	cdc.RegisterConcrete(&MsgCloseAirdropCampaign{}, "cfeairdrop/CloseAirdropCampaign", nil)
	cdc.RegisterConcrete(&MsgStartAirdropCampaign{}, "cfeairdrop/StartAirdropCampaign", nil)
	cdc.RegisterConcrete(&MsgEditAirdropCampaign{}, "cfeairdrop/EditAirdropCampaign", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaim{},
		&MsgInitialClaim{},
		&MsgCreateAirdropCampaign{},
		&MsgAddMissionToAidropCampaign{},
		&MsgAddAirdropEntries{},
		&MsgDeleteAirdropEntry{},
		&MsgCloseAirdropCampaign{},
		&MsgStartAirdropCampaign{},
		&MsgEditAirdropCampaign{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
