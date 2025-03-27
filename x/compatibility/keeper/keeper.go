package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/chain4energy/c4e-chain/x/compatibility/types"
)

type (
	Keeper struct {
		cdc                codec.BinaryCodec
		stakingStoreKey    storetypes.StoreKey
		c4eVestingStoreKey storetypes.StoreKey
		// memKey     storetypes.StoreKey
		// paramstore paramtypes.Subspace

		stakingKeeper    types.StakingKeeper
		cfevestingKeeper types.CfevestingKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	stakingStoreKey storetypes.StoreKey,
	c4eVestingStoreKey storetypes.StoreKey,
	// storeKey,
	// memKey storetypes.StoreKey,
	// ps paramtypes.Subspace,

	stakingKeeper types.StakingKeeper,
	cfevestingKeeper types.CfevestingKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	// if !ps.HasKeyTable() {
	// 	ps = ps.WithKeyTable(types.ParamKeyTable())
	// }

	return &Keeper{
		cdc:                cdc,
		stakingStoreKey:    stakingStoreKey,
		c4eVestingStoreKey: c4eVestingStoreKey,
		// storeKey:   storeKey,
		// memKey:     memKey,
		// paramstore: ps,

		stakingKeeper:    stakingKeeper,
		cfevestingKeeper: cfevestingKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
