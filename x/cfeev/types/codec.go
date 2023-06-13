package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPublishEnergyTransferOffer{}, "cfeev/PublishEnergyTransferOffer", nil)
	cdc.RegisterConcrete(&MsgStartEnergyTransfer{}, "cfeev/StartEnergyTransfer", nil)
	cdc.RegisterConcrete(&MsgEnergyTransferStarted{}, "cfeev/EnergyTransferStarted", nil)
	cdc.RegisterConcrete(&MsgEnergyTransferCompleted{}, "cfeev/EnergyTransferCompleted", nil)
	cdc.RegisterConcrete(&MsgCancelEnergyTransfer{}, "cfeev/CancelEnergyTransfer", nil)
	cdc.RegisterConcrete(&MsgRemoveEnergyOffer{}, "cfeev/RemoveEnergyOffer", nil)
	cdc.RegisterConcrete(&MsgRemoveTransfer{}, "cfeev/RemoveTransfer", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPublishEnergyTransferOffer{},
		&MsgStartEnergyTransfer{},
		&MsgEnergyTransferStarted{},
		&MsgEnergyTransferCompleted{},
		&MsgCancelEnergyTransfer{},
		&MsgRemoveEnergyOffer{},
		&MsgRemoveTransfer{},
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
