package cfeclaim

import (
	"github.com/chain4energy/c4e-chain/testutil/nullify"
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

func (d *C4eClaimKeeperUtils) ExportGenesis(ctx sdk.Context, genState types.GenesisState) {
	got := cfeclaim.ExportGenesis(ctx, *d.helpeCfeclaimkeeper)
	require.NotNil(d.t, got)

	nullify.Fill(&genState)
	nullify.Fill(got)

	require.ElementsMatch(d.t, genState.UsersEntries, got.UsersEntries)
	require.ElementsMatch(d.t, genState.Missions, got.Missions)
	require.ElementsMatch(d.t, genState.Campaigns, got.Campaigns)
}
