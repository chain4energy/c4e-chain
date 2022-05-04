package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	// distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	// distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	// govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	// govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	// stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	// stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type msgServer struct {
	Keeper
	// stakingMsgServer stakingtypes.MsgServer
	// distrMsgServer   distrtypes.MsgServer
	// govMsgServer     govtypes.MsgServer
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// func NewMsgServerImpl(keeper Keeper) types.MsgServer {
// 	return &msgServer{Keeper: keeper,
// 		stakingMsgServer: stakingkeeper.NewMsgServerImpl(keeper.staking.(stakingkeeper.Keeper)),
// 		distrMsgServer:   distrkeeper.NewMsgServerImpl(keeper.distribution.(distrkeeper.Keeper)),
// 		govMsgServer:     govkeeper.NewMsgServerImpl(keeper.gov.(govkeeper.Keeper))}
// }

var _ types.MsgServer = msgServer{}
