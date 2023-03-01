package keeper

import (
	v110cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v110"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	if err := v110cfedistributor.MigrateParams(ctx, &m.keeper.paramstore); err != nil {
		return err
	}

	return v110cfedistributor.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	var oldSubDistributors []types.SubDistributor
	oldSubDistributorsRaw := m.keeper.paramstore.GetRaw(ctx, types.KeySubDistributors)
	if err := codec.NewLegacyAmino().UnmarshalJSON(oldSubDistributorsRaw, &oldSubDistributors); err != nil {
		panic(err)
	}
	if err := m.keeper.SetParams(ctx, types.NewParams(oldSubDistributors)); err != nil {
		return err
	}
	return nil
}
