package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgVest{}, "cfevesting/Vest", nil)
	cdc.RegisterConcrete(&MsgWithdrawAllAvailable{}, "cfevesting/WithdrawAllAvailable", nil)
	cdc.RegisterConcrete(&MsgDelegate{}, "cfevesting/Delegate", nil)
	cdc.RegisterConcrete(&MsgUndelegate{}, "cfevesting/Undelegate", nil)
	cdc.RegisterConcrete(&MsgBeginRedelegate{}, "cfevesting/BeginRedelegate", nil)
	cdc.RegisterConcrete(&MsgWithdrawDelegatorReward{}, "cfevesting/WithdrawDelegatorReward", nil)
	cdc.RegisterConcrete(&MsgSendVesting{}, "cfevesting/SendVesting", nil)
	cdc.RegisterConcrete(&MsgVote{}, "cfevesting/Vote", nil)
	cdc.RegisterConcrete(&MsgVoteWeighted{}, "cfevesting/VoteWeighted", nil)
	cdc.RegisterConcrete(&MsgCreateVestingAccount{}, "cfevesting/CreateVestingAccount", nil)
	cdc.RegisterConcrete(&MsgSendToVestingAccount{}, "cfevesting/SendToVestingAccount", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawAllAvailable{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDelegate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUndelegate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBeginRedelegate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawDelegatorReward{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendVesting{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVote{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVoteWeighted{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVestingAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendToVestingAccount{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
