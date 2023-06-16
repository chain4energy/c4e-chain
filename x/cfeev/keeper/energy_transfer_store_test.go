package keeper_test

import (
	"cosmossdk.io/math"
	"testing"
	"time"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createAndAppendNEnergyTransfers1(n int) (EnergyTransfers []types.EnergyTransfer) {
	for i := 0; i < n; i++ {
		energyTransfer := types.EnergyTransfer{
			Id:                    uint64(i),
			EnergyTransferOfferId: 0,
			ChargerId:             "",
			OwnerAccountAddress:   "",
			DriverAccountAddress:  "",
			OfferedTariff:         0,
			Status:                0,
			Collateral:            math.NewInt(100),
			EnergyToTransfer:      0,
			EnergyTransferred:     0,
			PaidDate:              time.Time{},
		}
		EnergyTransfers = append(EnergyTransfers, energyTransfer)
	}
	return
}

func appendEnergyTransfers(k *keeper.Keeper, ctx sdk.Context, EnergyTransfers []types.EnergyTransfer) {
	for _, energyTransfer := range EnergyTransfers {
		k.AppendEnergyTransfer(ctx, energyTransfer)
	}

	return
}

func createAndAppendNEnergyTransfers(k *keeper.Keeper, ctx sdk.Context, n int) []types.EnergyTransfer {
	EnergyTransfers := createAndAppendNEnergyTransfers1(n)
	appendEnergyTransfers(k, ctx, EnergyTransfers)
	return EnergyTransfers
}

func TestEnergyTransferGet(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createAndAppendNEnergyTransfers(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEnergyTransfer(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEnergyTransferRemove(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createAndAppendNEnergyTransfers(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEnergyTransfer(ctx, item.Id)
		_, found := keeper.GetEnergyTransfer(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEnergyTransferGetAll(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createAndAppendNEnergyTransfers(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEnergyTransfers(ctx)),
	)
}

func TestEnergyTransferCount(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createAndAppendNEnergyTransfers(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEnergyTransferCount(ctx))
}
