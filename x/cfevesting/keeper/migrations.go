package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/exported"
	v110cfevesting "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v2"
	v120cfevesting "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v3"
	v4 "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v4"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper         Keeper
	legacySubspace exported.Subspace
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper, legacySubspace exported.Subspace) Migrator {
	return Migrator{keeper: keeper, legacySubspace: legacySubspace}
}

// Migrate2to3 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v110cfevesting.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v120cfevesting.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

// Migrate2to3 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	return v4.MigrateStore(ctx, m.keeper.storeKey, m.legacySubspace, m.keeper.cdc)
}
