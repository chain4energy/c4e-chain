package e2e

import (
	"encoding/json"
	"fmt"
	cfeevutils "github.com/chain4energy/c4e-chain/testutil/module/cfeev"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type EvSetupSuite struct {
	BaseSetupSuite
}

func TestEvSuite(t *testing.T) {
	suite.Run(t, new(EvSetupSuite))
}

func (s *EvSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(false, false, false)
}

func (s *EvSetupSuite) TestAllEVLogic() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	offerPublisherWalletName := utils.RandStringOfLength(10)
	offerPublisherAddress := node.CreateWallet(offerPublisherWalletName)
	node.BankSendBaseBalanceFromNode(offerPublisherAddress)

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	tariff := 100
	chargerId := utils.RandStringOfLength(10)
	name := utils.RandStringOfLength(10)
	plugType := types.PlugType(utils.RandInt64(r, 4))
	location := cfeevutils.RandomLocation(r)
	locationJson, _ := json.Marshal(location)
	offerId := node.PublishEnergyTransferOffer(chargerId, strconv.Itoa(tariff), string(locationJson), name, plugType.String(), offerPublisherWalletName)
	stringOfferId := strconv.Itoa(int(offerId))

	driverWalletName := utils.RandStringOfLength(10)
	driverAddress := node.CreateWallet(driverWalletName)
	node.BankSendBaseBalanceFromNode(driverAddress)

	energyToTransfer := 100
	node.StartEnergyTransferError(stringOfferId, strconv.Itoa(tariff+100), strconv.Itoa(energyToTransfer), driverWalletName, fmt.Sprintf("wrong tariff expected %d got %d", tariff, tariff+100))
	node.StartEnergyTransferError(stringOfferId, strconv.Itoa(tariff), strconv.Itoa(energyToTransfer*100000000), driverWalletName, "owner balance is too small")
	transferId := node.StartEnergyTransfer(stringOfferId, strconv.Itoa(tariff), strconv.Itoa(energyToTransfer), driverWalletName)

	stringTransferId := strconv.Itoa(int(transferId))
	node.RemoveEnergyTransferError(stringTransferId, driverWalletName, fmt.Sprintf("energy transfer status must be PAID or CANCELLED not REQUESTED"))

	node.EnergyTransferStarted(stringTransferId, stringOfferId, "", driverWalletName)
	node.CancelEnergyTransferdError(stringTransferId, "", "", driverWalletName, "energy transfer status must be REQUESTED not ONGOING")

	node.EnergyTransferCompleted(stringTransferId, stringOfferId, strconv.Itoa(75), "", driverWalletName)

	node.RemoveEnergyTransfer(stringTransferId, driverWalletName)
	node.RemoveEnergyOffer(stringOfferId, offerPublisherWalletName)
	node.StartEnergyTransferError(stringOfferId, strconv.Itoa(tariff), strconv.Itoa(energyToTransfer), driverWalletName, fmt.Sprintf("energy transfer offer with id %d not found", offerId))
}

func (s *EvSetupSuite) TestCancelEnergyTransfer() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	offerPublisherWalletName := utils.RandStringOfLength(10)
	offerPublisherAddress := node.CreateWallet(offerPublisherWalletName)
	node.BankSendBaseBalanceFromNode(offerPublisherAddress)

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	tariff := 100
	chargerId := utils.RandStringOfLength(10)
	name := utils.RandStringOfLength(10)
	plugType := types.PlugType(utils.RandInt64(r, 4))
	location := cfeevutils.RandomLocation(r)
	locationJson, _ := json.Marshal(location)
	offerId := node.PublishEnergyTransferOffer(chargerId, strconv.Itoa(tariff), string(locationJson), name, plugType.String(), offerPublisherWalletName)
	stringOfferId := strconv.Itoa(int(offerId))

	driverWalletName := utils.RandStringOfLength(10)
	driverAddress := node.CreateWallet(driverWalletName)
	node.BankSendBaseBalanceFromNode(driverAddress)

	energyToTransfer := 100
	transferId := node.StartEnergyTransfer(stringOfferId, strconv.Itoa(tariff), strconv.Itoa(energyToTransfer), driverWalletName)
	stringTransferId := strconv.Itoa(int(transferId))
	node.RemoveEnergyTransferError(stringTransferId, driverWalletName, fmt.Sprintf("energy transfer status must be PAID or CANCELLED not REQUESTED"))
	node.RemoveEnergyOfferError(stringTransferId, offerPublisherWalletName, fmt.Sprintf("energy transfer offer charger status must be ACTIVE or INACTIVE not BUSY"))
	node.CancelEnergyTransfer(stringTransferId, "", "", driverWalletName)
	node.RemoveEnergyOffer(stringTransferId, offerPublisherWalletName)

	node.EnergyTransferStartedError(stringTransferId, "", driverWalletName, "energy transfer status must be REQUESTED not CANCELLED")
	node.CancelEnergyTransferdError(stringTransferId, "", "", driverWalletName, fmt.Sprintf("energy transfer status must be REQUESTED not CANCELLED"))

	node.EnergyTransferCompletedError(stringTransferId, strconv.Itoa(75), "", driverWalletName, "energy transfer status must be REQUESTED or ONGOING not CANCELLED")
	node.RemoveEnergyTransfer(stringTransferId, driverWalletName)
}
