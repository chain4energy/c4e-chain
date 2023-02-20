package keeper

import (
	v120cfevesting "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v120"
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

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v120cfevesting.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
