package helpers

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// SimAppChainID hardcoded chainID for simulation
const (
	DefaultGenTxGas = 10000000
	SimAppChainID   = "simulation-app"
)

func CreateRandomAccAddressNoBalance(i int64) string {
	var buffer bytes.Buffer
	numString := strconv.Itoa(int(i))
	buffer.WriteString("A58856F0FD53BF058B4909A21AEC019107BA6") // base address string
	buffer.WriteString(numString)                               // adding on final two digits to make addresses unique
	res, _ := sdk.AccAddressFromHex(buffer.String())
	return res.String()
}
