package e2e

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"

	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type WasmSetupSuite struct {
	BaseSetupSuite
}

func TestWasmSuite(t *testing.T) {
	suite.Run(t, new(WasmSetupSuite))
}

func (s *WasmSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(false, false, false)
}

func (s *WasmSetupSuite) TestUploadAndInstantiateCounter() {
	chain := s.configurer.GetChainConfig(0)
	chain.WaitUntilHeight(1)
	node, err := chain.GetDefaultNode()
	s.NoError(err)

	codeId := node.StoreWasmCode("bytecode/counter.wasm", initialization.ValidatorWalletName)

	node.InstantiateWasmContract(
		strconv.Itoa(codeId),
		`{"count": 0}`,
		initialization.ValidatorWalletName)
	contracts, err := node.QueryContractsFromId(codeId)
	s.NoError(err)
	s.Require().Len(contracts, 1)
	contractAddr := contracts[0]

	node.WasmExecute(contractAddr, `{"increment":{}}`, initialization.ValidatorWalletName)

	resultObject, err := node.QueryWasmSmartObject(contractAddr, fmt.Sprintf(`{"get_count":{"addr": "%s"}}`, node.PublicAddress))
	s.NoError(err)
	s.Require().EqualValues(1, resultObject["count"])

	node.WasmExecute(contractAddr, `{"increment":{}}`, initialization.ValidatorWalletName)

	resultObject, err = node.QueryWasmSmartObject(contractAddr, fmt.Sprintf(`{"get_count":{"addr": "%s"}}`, node.PublicAddress))
	s.NoError(err)
	s.Require().EqualValues(2, resultObject["count"])
}
