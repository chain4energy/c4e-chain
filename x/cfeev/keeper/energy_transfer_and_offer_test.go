package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func TestStartEnergyTransfer(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
}

func TestStartEnergyTransferBalanceToSmall(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff*energyTransfer.EnergyToTransfer - 1)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	testHelper.C4eEvUtils.StartEnergyTransferError(acountsAddresses[1].String(), energyTransfer, offerId,
		fmt.Sprintf("owner balance is too small (%s < %s): insufficient funds", sdk.NewCoin(testenv.DefaultTestDenom, collateral), sdk.NewCoin(testenv.DefaultTestDenom, collateral.AddRaw(1))))
}

func TestStartEnergyTransferTarifChanged(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	energyTransfer.OfferedTariff = offer.Tariff + 1
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	testHelper.C4eEvUtils.StartEnergyTransferError(acountsAddresses[1].String(), energyTransfer, offerId,
		fmt.Sprintf("wrong tariff expected %d got %d: wrong param value", offer.Tariff, energyTransfer.OfferedTariff))
}

func TestStartEnergyTransferWrongOfferID(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	energyTransfer.OfferedTariff = offer.Tariff + 1
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	testHelper.C4eEvUtils.StartEnergyTransferError(acountsAddresses[1].String(), energyTransfer, offerId+1, "energy transfer offer with id 1 not found: not found")
}

func TestStartEnergyTransferInvalidAddress(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	testHelper.C4eEvUtils.StartEnergyTransferError("invalid_address", energyTransfer, offerId, "invalid creator address (decoding bech32 failed: invalid separator index -1): invalid address")
}

func TestEnergyTransferStarted(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
}

func TestEnergyTransferStartedTwice(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferStartedError(transferId, offerId, "energy transfer status must be REQUESTED not ONGOING: invalid type")
}

func TestEnergyTransferStartedWrongTransferId(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	wrongTransferId := transferId + 1
	testHelper.C4eEvUtils.EnergyTransferStartedError(wrongTransferId, offerId, fmt.Sprintf("energy transfer with id %d not found: not found", wrongTransferId))
}

func TestEnergyTransferCompletedFull(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompleted(transferId, energyTransfer.EnergyToTransfer, offerId, math.ZeroInt(), collateral)
}

func TestEnergyTransferCompletedHalf(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompleted(transferId, energyTransfer.EnergyToTransfer/2, offerId, collateral.QuoRaw(2), collateral.QuoRaw(2))
}

func TestEnergyTransferCompletedNothingTransfered(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompleted(transferId, 0, offerId, collateral, math.ZeroInt())
}

func TestEnergyTransferCompletedMoreTransferedLessThan4(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompleted(transferId, energyTransfer.EnergyToTransfer+3, offerId, math.ZeroInt(), collateral)
}

func TestEnergyTransferCompletedMoreTransferedMoreThan4(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompleted(transferId, energyTransfer.EnergyToTransfer+50, offerId, math.ZeroInt(), collateral)
}

func TestEnergyTransferCompletedWrongTransferId(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompletedError(transferId+1, energyTransfer.EnergyToTransfer, offerId, "energy transfer with id 1 not found: not found")
}

func TestEnergyTransferCompletedWrongNotStarted(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompleted(transferId, energyTransfer.EnergyToTransfer, offerId, math.ZeroInt(), collateral)
}

func TestEnergyTransferCompletedTwice(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompleted(transferId, energyTransfer.EnergyToTransfer, offerId, math.ZeroInt(), collateral)
	testHelper.C4eEvUtils.EnergyTransferCompletedError(transferId, energyTransfer.EnergyToTransfer, offerId, "energy transfer status must be REQUESTED or ONGOING not PAID: invalid type")
}

func TestEnergyTransferCompletedCanceled(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.CancelEnergyTransfer(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompletedError(transferId, energyTransfer.EnergyToTransfer, offerId, "energy transfer status must be REQUESTED or ONGOING not CANCELLED: invalid type")
}

func TestCancelEnergyTransfer(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.CancelEnergyTransfer(transferId, offerId)
}

func TestCancelEnergyTransferWrongTransferId(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	wrongTransferId := transferId + 1
	testHelper.C4eEvUtils.CancelEnergyTransferError(wrongTransferId, offerId, fmt.Sprintf("energy transfer with id %d not found: not found", wrongTransferId))
}

func TestCancelEnergyTransferAlreadyStarted(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.CancelEnergyTransferError(transferId, offerId, "energy transfer status must be REQUESTED not ONGOING: invalid type")
}

func TestCancelEnergyTransferAlreadyCompleted(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompleted(transferId, energyTransfer.EnergyToTransfer, offerId, math.ZeroInt(), collateral)
	testHelper.C4eEvUtils.CancelEnergyTransferError(transferId, offerId, "energy transfer status must be REQUESTED not PAID: invalid type")
}

func TestCancelEnergyTransferTwice(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.CancelEnergyTransfer(transferId, offerId)
	testHelper.C4eEvUtils.CancelEnergyTransferError(transferId, offerId, "energy transfer status must be REQUESTED not CANCELLED: invalid type")
}

func TestRemoveEnergyTransfer(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.CancelEnergyTransfer(transferId, offerId)
	testHelper.C4eEvUtils.RemoveEnergyTransfer(acountsAddresses[1].String(), transferId)
}

func TestRemoveEnergyTransferStart(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.RemoveEnergyTransferError(acountsAddresses[1].String(), transferId, "energy transfer status must be PAID or CANCELLED not REQUESTED: invalid type")
}

func TestRemoveEnergyTransferStarted(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.RemoveEnergyTransferError(acountsAddresses[1].String(), transferId, "energy transfer status must be PAID or CANCELLED not ONGOING: invalid type")
}

func TestRemoveEnergyTransferCompleted(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.EnergyTransferStarted(transferId, offerId)
	testHelper.C4eEvUtils.EnergyTransferCompleted(transferId, energyTransfer.EnergyToTransfer, offerId, math.ZeroInt(), collateral)
	testHelper.C4eEvUtils.RemoveEnergyTransfer(acountsAddresses[1].String(), transferId)
}

func TestRemoveEnergyTransferCanceled(t *testing.T) {
	testHelper := app.SetupTestApp(t)
	offer := prepareTestTransferOffer()
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	offerId := testHelper.C4eEvUtils.PublishEnergyTransferOffer(acountsAddresses[0].String(), offer)
	energyTransfer := prepareTestEnergyTransfer(offer.Tariff)
	collateral := math.NewIntFromUint64(offer.Tariff * energyTransfer.EnergyToTransfer)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(collateral, acountsAddresses[1])
	transferId := testHelper.C4eEvUtils.StartEnergyTransfer(acountsAddresses[1].String(), energyTransfer, offerId)
	testHelper.C4eEvUtils.CancelEnergyTransfer(transferId, offerId)
	testHelper.C4eEvUtils.RemoveEnergyTransfer(acountsAddresses[1].String(), transferId)
}

func prepareTestEnergyTransfer(oferedTarif uint64) types.EnergyTransfer {
	return types.EnergyTransfer{
		Id:               0,
		OfferedTariff:    oferedTarif,
		EnergyToTransfer: 100,
	}
}
