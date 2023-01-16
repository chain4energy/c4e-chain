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
	cdc.RegisterConcrete(&MsgClaim{}, "cfeairdrop/Claim", nil)
	cdc.RegisterConcrete(&MsgCreateAirdropCampaign{}, "cfeairdrop/CreateAirdropCampaign", nil)
	cdc.RegisterConcrete(&MsgAddMissionToAidropCampaign{}, "cfeairdrop/AddMissionToAidropCampaign", nil)
	cdc.RegisterConcrete(&MsgAddAirdropEntries{}, "cfeairdrop/AddAirdropEntries", nil)
	cdc.RegisterConcrete(&MsgDeleteAirdropEntry{}, "cfeairdrop/DeleteAirdropEntry", nil)
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&AirdropVestingAccount{}, "c4e/AirdropVestingAccount", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaim{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAirdropCampaign{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddMissionToAidropCampaign{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAirdropEntries{},
		&MsgDeleteAirdropEntry{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations(
		(*vestexported.VestingAccount)(nil),
		&AirdropVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&AirdropVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&AirdropVestingAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
