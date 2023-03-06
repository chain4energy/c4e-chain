package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

var (
	_ proto.Message                      = (*LinearMinting)(nil)
	_ proto.Message                      = (*ExponentialStepMinting)(nil)
	_ MinterConfigI                      = (*LinearMinting)(nil)
	_ MinterConfigI                      = (*ExponentialStepMinting)(nil)
	_ codectypes.UnpackInterfacesMessage = (*Minter)(nil)
	_ codectypes.UnpackInterfacesMessage = (*Params)(nil)
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*MinterConfigI)(nil), nil)
	cdc.RegisterConcrete(&ExponentialStepMinting{}, "c4echain/cfeminter/ExponentialStepMinting", nil)
	cdc.RegisterConcrete(&LinearMinting{}, "c4echain/cfeminter/LinearMinting", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"chain4energy.c4echain.cfeminter.MinterConfigI",
		(*MinterConfigI)(nil),
		&LinearMinting{},
		&ExponentialStepMinting{},
	)

	registry.RegisterImplementations(
		(*MinterConfigI)(nil),
		&LinearMinting{},
		&ExponentialStepMinting{},
	)
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
