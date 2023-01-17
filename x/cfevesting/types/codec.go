package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateVestingPool{}, "cfevesting/CreateVestingPool", nil)
	cdc.RegisterConcrete(&MsgWithdrawAllAvailable{}, "cfevesting/WithdrawAllAvailable", nil)
	cdc.RegisterConcrete(&MsgCreateVestingAccount{}, "cfevesting/CreateVestingAccount", nil)
	cdc.RegisterConcrete(&MsgSendToVestingAccount{}, "cfevesting/SendToVestingAccount", nil)
	cdc.RegisterConcrete(&RepeatedContinuousVestingAccount{}, "c4e/RepeatedContinuousVestingAccount", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVestingPool{},
		&MsgWithdrawAllAvailable{},
		&MsgCreateVestingAccount{},
		&MsgSendToVestingAccount{},
	)

	registry.RegisterImplementations(
		(*vestexported.VestingAccount)(nil),
		&RepeatedContinuousVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&RepeatedContinuousVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&RepeatedContinuousVestingAccount{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	Amino.Seal()
}
