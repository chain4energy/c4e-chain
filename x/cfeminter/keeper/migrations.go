package keeper

import (
	v110cfeminter "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v110"
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
	if err := v110cfeminter.MigrateParams(ctx, m.keeper.storeKey, &m.keeper.paramstore); err != nil {
		return err
	}

	return v110cfeminter.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
