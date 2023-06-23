package keeper

import (
	"cosmossdk.io/errors"
	"encoding/binary"
	"github.com/chain4energy/c4e-chain/types/util"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetEnergyTransferCount get the total number of energyTransfer
func (k Keeper) GetEnergyTransferCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(types.EnergyTransferCountKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetEnergyTransferCount set the total number of energyTransfer
func (k Keeper) SetEnergyTransferCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.EnergyTransferCountKey
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendEnergyTransfer appends a energyTransfer in the store with a new id and update the count
func (k Keeper) AppendEnergyTransfer(
	ctx sdk.Context,
	energyTransfer types.EnergyTransfer,
) uint64 {
	// Create the energyTransfer
	count := k.GetEnergyTransferCount(ctx)

	// Set the ID of the appended value
	energyTransfer.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferKey)
	appendedValue := k.cdc.MustMarshal(&energyTransfer)
	store.Set(util.GetUint64Key(energyTransfer.Id), appendedValue)

	// Update energyTransfer count
	k.SetEnergyTransferCount(ctx, count+1)

	return count
}

// SetEnergyTransfer set a specific energyTransfer in the store
func (k Keeper) SetEnergyTransfer(ctx sdk.Context, energyTransfer types.EnergyTransfer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferKey)
	b := k.cdc.MustMarshal(&energyTransfer)
	store.Set(util.GetUint64Key(energyTransfer.Id), b)
}

// GetEnergyTransfer returns a energyTransfer from its id
func (k Keeper) GetEnergyTransfer(ctx sdk.Context, id uint64) (val types.EnergyTransfer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferKey)
	b := store.Get(util.GetUint64Key(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) MustGetEnergyTransfer(ctx sdk.Context, id uint64) (*types.EnergyTransfer, error) {
	energyTransfer, found := k.GetEnergyTransfer(ctx, id)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "energy transfer with id %d not found", id)
	}
	return &energyTransfer, nil
}

// RemoveEnergyTransfer removes a energyTransfer from the store
func (k Keeper) DeleteEnergyTransfer(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferKey)
	store.Delete(util.GetUint64Key(id))
}

// GetAllEnergyTransfers returns all energyTransfer
func (k Keeper) GetAllEnergyTransfers(ctx sdk.Context) (list []types.EnergyTransfer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EnergyTransfer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
