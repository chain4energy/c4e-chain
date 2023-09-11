package keeper

import (
	"github.com/chain4energy/c4e-chain/v2/types/subspace"
	v2 "github.com/chain4energy/c4e-chain/v2/x/cfevesting/migrations/v2"
	v3 "github.com/chain4energy/c4e-chain/v2/x/cfevesting/migrations/v3"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper         Keeper
	legacySubspace subspace.Subspace
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper, legacySubspace subspace.Subspace) Migrator {
	return Migrator{keeper: keeper, legacySubspace: legacySubspace}
}

// Migrate2to3 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := v3.MigrateParams(ctx, m.keeper.storeKey, m.legacySubspace, m.keeper.cdc); err != nil {
		return err
	}
	return v3.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
