package cfeclaim

import (
	"github.com/chain4energy/c4e-chain/x/cfeclaim"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type C4eClaimKeeperUtils struct {
	t                   require.TestingT
	helpeCfeclaimkeeper *keeper.Keeper
}

func NewC4eClaimKeeperUtils(t require.TestingT, helpeCfeclaimmodulekeeper *keeper.Keeper) C4eClaimKeeperUtils {
	return C4eClaimKeeperUtils{t: t, helpeCfeclaimkeeper: helpeCfeclaimmodulekeeper}
}

func (d *C4eClaimKeeperUtils) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	cfeclaim.InitGenesis(ctx, *d.helpeCfeclaimkeeper, genState)
}
