package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	//	cdc.RegisterConcrete(&MsgCreateUserDevices{}, "cfetokenization/CreateUserDevices", nil)
	//cdc.RegisterConcrete(&MsgUpdateUserDevices{}, "cfetokenization/UpdateUserDevices", nil)
	//cdc.RegisterConcrete(&MsgDeleteUserDevices{}, "cfetokenization/DeleteUserDevices", nil)
	cdc.RegisterConcrete(&MsgCreateUserCertificates{}, "cfetokenization/CreateUserCertificates", nil)
	cdc.RegisterConcrete(&MsgUpdateUserCertificates{}, "cfetokenization/UpdateUserCertificates", nil)
	cdc.RegisterConcrete(&MsgDeleteUserCertificates{}, "cfetokenization/DeleteUserCertificates", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil))
	//&MsgCreateUserDevices{},
	//&MsgUpdateUserDevices{},
	//&MsgDeleteUserDevices{},

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUserCertificates{},
		&MsgUpdateUserCertificates{},
		&MsgDeleteUserCertificates{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
