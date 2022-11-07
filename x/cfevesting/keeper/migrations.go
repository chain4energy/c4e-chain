package keeper

import (
	v101cfevesting "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v101"
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
	logger := m.keeper.Logger(ctx)
	logger.Info("Starting migration cfevesting - migration 1 to 2")
	return v101cfevesting.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
