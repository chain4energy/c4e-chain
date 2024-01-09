package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgBurn{}, "cfeminter/Burn")
	cdc.RegisterInterface((*MinterConfigI)(nil), nil)
	cdc.RegisterConcrete(&ExponentialStepMinting{}, "c4e-chain/ExponentialStepMinting", nil)
	cdc.RegisterConcrete(&LinearMinting{}, "c4e-chain/LinearMinting", nil)
	cdc.RegisterConcrete(&NoMinting{}, "c4e-chain/NoMinting", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"chain4energy.c4echain.cfeminter.MinterConfigI",
		(*MinterConfigI)(nil),
		&LinearMinting{},
		&ExponentialStepMinting{},
		&NoMinting{},
	)

	registry.RegisterImplementations(
		(*MinterConfigI)(nil),
		&LinearMinting{},
		&ExponentialStepMinting{},
		&NoMinting{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurn{},
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
