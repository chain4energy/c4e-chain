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
	legacy.RegisterAminoMsg(cdc, &MsgCreateUserCertificates{}, "cfetokenization/CreateUserCertificates")
	legacy.RegisterAminoMsg(cdc, &MsgBuyCertificate{}, "cfetokenization/BuyCertificate")
	legacy.RegisterAminoMsg(cdc, &MsgAddCertificateToMarketplace{}, "cfetokenization/AddCertificateToMarketplace")
	legacy.RegisterAminoMsg(cdc, &MsgAcceptDevice{}, "cfetokenization/AcceptDevice")
	legacy.RegisterAminoMsg(cdc, &MsgAssignDeviceToUser{}, "cfetokenization/AssignDeviceToUser")
	legacy.RegisterAminoMsg(cdc, &MsgBurnCertificate{}, "cfetokenization/AddMeasurement")
	legacy.RegisterAminoMsg(cdc, &MsgAddMeasurement{}, "cfetokenization/AddMeasurement")
	legacy.RegisterAminoMsg(cdc, &MsgAuthorizeCertificate{}, "cfetokenization/AuthorizeCertificate")
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil))

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUserCertificates{},
		&MsgBuyCertificate{},
		&MsgAddCertificateToMarketplace{},
		&MsgAcceptDevice{},
		&MsgAssignDeviceToUser{},
		&MsgBurnCertificate{},
		&MsgAddMeasurement{},
		&MsgAuthorizeCertificate{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
