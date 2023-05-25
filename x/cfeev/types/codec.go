package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPublishEnergyTransferOffer{}, "cfeev/PublishEnergyTransferOffer", nil)
	cdc.RegisterConcrete(&MsgStartEnergyTransferRequest{}, "cfeev/StartEnergyTransferRequest", nil)
	cdc.RegisterConcrete(&MsgEnergyTransferStartedRequest{}, "cfeev/EnergyTransferStartedRequest", nil)
	cdc.RegisterConcrete(&MsgEnergyTransferCompletedRequest{}, "cfeev/EnergyTransferCompletedRequest", nil)
	cdc.RegisterConcrete(&MsgCancelEnergyTransferRequest{}, "cfeev/CancelEnergyTransferRequest", nil)
	cdc.RegisterConcrete(&MsgRemoveEnergyOffer{}, "cfeev/RemoveEnergyOffer", nil)
	cdc.RegisterConcrete(&MsgRemoveTransfer{}, "cfeev/RemoveTransfer", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPublishEnergyTransferOffer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStartEnergyTransferRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnergyTransferStartedRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnergyTransferCompletedRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelEnergyTransferRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveEnergyOffer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveTransfer{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
