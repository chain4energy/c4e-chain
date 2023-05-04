package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/bank interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgCreateVestingPool{}, "cfevesting/CreateVestingPool")
	legacy.RegisterAminoMsg(cdc, &MsgWithdrawAllAvailable{}, "cfevesting/WithdrawAllAvailable")
	legacy.RegisterAminoMsg(cdc, &MsgCreateVestingAccount{}, "cfevesting/CreateVestingAccount")
	legacy.RegisterAminoMsg(cdc, &MsgSendToVestingAccount{}, "cfevesting/SendToVestingAccount")
	legacy.RegisterAminoMsg(cdc, &MsgSplitVesting{}, "cfevesting/SplitVesting")
	legacy.RegisterAminoMsg(cdc, &MsgMoveAvailableVesting{}, "cfevesting/MoveAvailableVesting")
	legacy.RegisterAminoMsg(cdc, &MsgMoveAvailableVestingByDenoms{}, "cfevesting/MoveAvailableVestingByDenoms")
	cdc.RegisterConcrete(&PeriodicContinuousVestingAccount{}, "c4e/PeriodicContinuousVestingAccount", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVestingPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawAllAvailable{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVestingAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendToVestingAccount{},
	)

	registry.RegisterImplementations(
		(*vestexported.VestingAccount)(nil),
		&PeriodicContinuousVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&PeriodicContinuousVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&PeriodicContinuousVestingAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSplitVesting{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMoveAvailableVesting{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMoveAvailableVestingByDenoms{},
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
