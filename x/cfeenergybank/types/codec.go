package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateTokenParams{}, "cfeenergybank/CreateTokenParams", nil)
	cdc.RegisterConcrete(&MsgMintToken{}, "cfeenergybank/MintToken", nil)
	cdc.RegisterConcrete(&MsgTransferTokens{}, "cfeenergybank/TransferTokens", nil)
	cdc.RegisterConcrete(&MsgTransferTokensOptimally{}, "cfeenergybank/TransferTokensOptimally", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTokenParams{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferTokens{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferTokensOptimally{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
