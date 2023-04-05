package cfeairdrop

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type C4eAirdropKeeperUtils struct {
	t                     require.TestingT
	helpeCfeairdropkeeper *keeper.Keeper
}

func NewC4eAirdropKeeperUtils(t require.TestingT, helpeCfeairdropmodulekeeper *keeper.Keeper) C4eAirdropKeeperUtils {
	return C4eAirdropKeeperUtils{t: t, helpeCfeairdropkeeper: helpeCfeairdropmodulekeeper}
}

func (d *C4eAirdropKeeperUtils) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	cfeairdrop.InitGenesis(ctx, *d.helpeCfeairdropkeeper, genState)
}
