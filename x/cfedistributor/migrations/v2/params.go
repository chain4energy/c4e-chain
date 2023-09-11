package v2

import (
	"github.com/chain4energy/c4e-chain/v2/types/subspace"
	"github.com/chain4energy/c4e-chain/v2/x/cfedistributor/migrations/v1"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MigrateParams performs in-place store migrations from v1.0.1 to v1.1.0.
// The migration includes:
// - SubDistributor params structure changed.
// - BurnShare and Share now must be set between 0 and 1, not 0 and 100.
func MigrateParams(ctx sdk.Context, paramStore subspace.Subspace) error {
	var oldSubDistributors []v1.SubDistributor
	if !paramStore.HasKeyTable() {
		paramStore.WithKeyTable(ParamKeyTable())
	}
	oldSubDistributorsRaw := paramStore.GetRaw(ctx, KeySubDistributors)
	if err := codec.NewLegacyAmino().UnmarshalJSON(oldSubDistributorsRaw, &oldSubDistributors); err != nil {
		panic(err)
	}

	var newSubDistributors []SubDistributor

	for _, oldSubDistributor := range oldSubDistributors {
		var newShares []*DestinationShare
		for _, oldShare := range oldSubDistributor.Destination.Share {
			newShare := DestinationShare{
				Share: oldShare.Percent.Quo(sdk.NewDec(100)),
				Destination: Account{
					Id:   oldShare.Account.Id,
					Type: oldShare.Account.Type,
				},
				Name: oldShare.Name,
			}
			newShares = append(newShares, &newShare)
		}

		var newSources []*Account
		for _, oldSource := range oldSubDistributor.Sources {
			newSource := Account{
				Id:   oldSource.Id,
				Type: oldSource.Type,
			}
			newSources = append(newSources, &newSource)
		}

		newSubDistributor := SubDistributor{
			Name: oldSubDistributor.Name,
			Destinations: Destinations{
				Shares:    newShares,
				BurnShare: oldSubDistributor.Destination.BurnShare.Percent.Quo(sdk.NewDec(100)),
				PrimaryShare: Account{
					Id:   oldSubDistributor.Destination.Account.Id,
					Type: oldSubDistributor.Destination.Account.Type,
				},
			},
			Sources: newSources,
		}

		newSubDistributors = append(newSubDistributors, newSubDistributor)
	}
	paramStore.Set(ctx, KeySubDistributors, newSubDistributors)

	return nil
}
