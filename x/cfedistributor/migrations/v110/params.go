package v110

import (
	"github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v101"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// MigrateParams performs in-place store migrations from v1.0.0 to v1.0.1. The
// migration includes:
//
// - SubDistributor params structure changed.
// - BurnShare and Share now must be set between 0 and 1, not 0 and 100.
func MigrateParams(ctx sdk.Context, storeKey storetypes.StoreKey, paramStore *paramtypes.Subspace) error {
	var oldSubdistributors []v101.SubDistributor
	oldSubdistributorsRaw := paramStore.GetRaw(ctx, types.KeySubDistributors)
	if err := codec.NewLegacyAmino().UnmarshalJSON(oldSubdistributorsRaw, &oldSubdistributors); err != nil {
		panic(err)
	}

	var newSubDistributors []types.SubDistributor

	for _, oldSubDistributor := range oldSubdistributors {
		var newShares []*types.DestinationShare
		for _, oldShare := range oldSubDistributor.Destination.Share {
			newShare := types.DestinationShare{
				Share: oldShare.Percent.Quo(sdk.NewDec(100)),
				Destination: types.Account{
					Id:   oldShare.Account.Id,
					Type: oldShare.Account.Type,
				},
				Name: oldShare.Name,
			}
			newShares = append(newShares, &newShare)
		}

		var newSources []*types.Account
		for _, oldSource := range oldSubDistributor.Sources {
			newSource := types.Account{
				Id:   oldSource.Id,
				Type: oldSource.Type,
			}
			newSources = append(newSources, &newSource)
		}

		newSubDistributor := types.SubDistributor{
			Name: oldSubDistributor.Name,
			Destinations: types.Destinations{
				Shares:    newShares,
				BurnShare: oldSubDistributor.Destination.BurnShare.Percent.Quo(sdk.NewDec(100)),
				PrimaryShare: types.Account{
					Id:   oldSubDistributor.Destination.Account.Id,
					Type: oldSubDistributor.Destination.Account.Type,
				},
			},
			Sources: newSources,
		}
		if err := newSubDistributor.Validate(); err != nil {
			return err
		}
		newSubDistributors = append(newSubDistributors, newSubDistributor)
	}
	if err := types.ValidateSubDistributors(newSubDistributors); err != nil {
		return err
	}
	paramStore.Set(ctx, types.KeySubDistributors, newSubDistributors)

	return nil
}
