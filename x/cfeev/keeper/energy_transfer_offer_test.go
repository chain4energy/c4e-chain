package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	cfeevutils "github.com/chain4energy/c4e-chain/testutil/module/cfeev"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	"github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEnergyTransferOfferGet(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createAndAppendNEnergyTransfersOffer(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEnergyTransferOffer(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEnergyTransferOfferRemove(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createAndAppendNEnergyTransfersOffer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEnergyTransferOffer(ctx, item.Id)
		_, found := keeper.GetEnergyTransferOffer(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEnergyTransferOfferGetAll(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createAndAppendNEnergyTransfersOffer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEnergyTransferOffers(ctx)),
	)
}

func TestEnergyTransferOfferCount(t *testing.T) {
	keeper, ctx, _ := keepertest.CfeevKeeper(t)
	items := createAndAppendNEnergyTransfersOffer(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEnergyTransferOfferCount(ctx))
}

func TestCreateEnergyTransferOffer(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
}

func TestCreateEnergyTransferOfferWrongOwner(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	testHelper.C4eEvUtils.PublishEnergyTransferOfferError("ivalid owner", offer,
		"invalid creator address (decoding bech32 failed: invalid character in string: ' '): invalid address")
}

func TestCreateEnergyTransferOfferWrongLongitude(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	wrongLongitude := sdk.NewDec(-200)
	offer.Location.Longitude = &wrongLongitude
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.C4eEvUtils.PublishEnergyTransferOfferError(acountsAddresses[0].String(), offer,
		"longitude must be between 180.000000000000000000 and -180.000000000000000000: wrong param value")
}

func TestCreateEnergyTransferOfferWrongLatitude(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	wrongLatitude := sdk.NewDec(-200)
	offer.Location.Latitude = &wrongLatitude
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.C4eEvUtils.PublishEnergyTransferOfferError(acountsAddresses[0].String(), offer,
		"latitude must be between 90.000000000000000000 and -90.000000000000000000: wrong param value")
}

func TestCreateEnergyTransferOfferWrongPlugType(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	offer.PlugType = 10
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.C4eEvUtils.PublishEnergyTransferOfferError(acountsAddresses[0].String(), offer,
		"invalid plug type (10): wrong param value")
}

func TestCreateEnergyTransferOfferEmptyName(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	offer.Name = ""
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.C4eEvUtils.PublishEnergyTransferOfferError(acountsAddresses[0].String(), offer,
		"charger name cannot be empty: wrong param value")
}

func TestCreateEnergyTransferOfferEmptyChargerid(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	offer.ChargerId = ""
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.C4eEvUtils.PublishEnergyTransferOfferError(acountsAddresses[0].String(), offer,
		"charger id cannot be empty: wrong param value")
}

func TestRemoveEnergyTransferOffer(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	testHelper.C4eEvUtils.RemoveEnergyTransferOffer(acountsAddresses[0].String(), offerId)
}

func TestRemoveEnergyTransferOfferWrongOwner(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	testHelper.C4eEvUtils.RemoveEnergyTransferOfferError(acountsAddresses[1].String(), offerId,
		fmt.Sprintf("address %s is not a creator of energy offer with id %d: tx intended signer does not match the given signer", acountsAddresses[1], offerId))
}

func TestRemoveEnergyTransferOfferDoesntExist(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	testHelper.C4eEvUtils.RemoveEnergyTransferOfferError(acountsAddresses[0].String(), 123, "energy transfer offer with id 123 not found: not found")
}

func TestRemoveEnergyTransferOfferStarted(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(math.NewIntFromUint64(offer.Tariff*energyTransfer.EnergyToTransfer), acountsAddresses[1])
	testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.RemoveEnergyTransferOfferError(acountsAddresses[0].String(), offerId, "energy transfer offer charger status must be ACTIVE or INACTIVE not BUSY: invalid type")
}

func TestRemoveEnergyTransferOfferStartedAndCompleted(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompleted(transferId, energyTransfer.EnergyToTransfer, offerId, math.NewInt(0), collateral)
	testHelper.C4eEvUtils.RemoveEnergyTransferOffer(acountsAddresses[0].String(), offerId)
}

func prepareTestTransferOffer() types.EnergyTransferOffer {
	r := utils.RandomSource()
	location := cfeevutils.RandomLocation(r)
	return types.EnergyTransferOffer{
		ChargerId: utils.RandStringOfLength(15),
		Location:  location,
		Tariff:    utils.RandUint64(r, 200),
		Name:      utils.RandStringOfLength(10),
		PlugType:  types.PlugType(utils.RandInt64(r, 4)),
	}
}

func createAndAppendNEnergyTransfersOffer(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EnergyTransferOffer {
	items := make([]types.EnergyTransferOffer, n)
	for i := range items {
		items[i].Id = keeper.AppendEnergyTransferOffer(ctx, items[i])
	}
	return items
}
