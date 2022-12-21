package common

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// CreateTestPubKeys returns a total of numPubKeys public keys in ascending order.
func CreateTestPubKeys(numPubKeys int) []cryptotypes.PubKey {
	var publicKeys []cryptotypes.PubKey
	var buffer bytes.Buffer

	// start at 10 to avoid changing 1 to 01, 2 to 02, etc
	for i := 100; i < (numPubKeys + 100); i++ {
		numString := strconv.Itoa(i)
		buffer.WriteString("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AF") // base pubkey string
		buffer.WriteString(numString)                                                       // adding on final two digits to make pubkeys unique
		publicKeys = append(publicKeys, NewPubKeyFromHex(buffer.String()))
		buffer.Reset()
	}

	return publicKeys
}

// NewPubKeyFromHex returns a PubKey from a hex string.
func NewPubKeyFromHex(pk string) (res cryptotypes.PubKey) {
	pkBytes, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}
	if len(pkBytes) != ed25519.PubKeySize {
		panic(errors.Wrap(errors.ErrInvalidPubKey, "invalid pubkey size"))
	}
	return &ed25519.PubKey{Key: pkBytes}
}

func CreateAccounts(accNum int, valAccNum int) (acountsAddresses []sdk.AccAddress, validatorsAddresses []sdk.ValAddress) {
	if accNum > 0 {
		acountsAddresses = CreateIncrementalAccounts(accNum, 0)
	}
	if valAccNum > 0 {
		validatorsAddresses = ConvertAddrsToValAddrs(CreateIncrementalAccounts(valAccNum, accNum))
	}
	return acountsAddresses, validatorsAddresses
}

func CreateIncrementalAccounts(accNum int, genInitNumber int) []sdk.AccAddress {
	var addresses []sdk.AccAddress
	// var buffer bytes.Buffer

	// start at 100 so we can make up to 999 test addresses with valid test addresses
	for i := 100; i < (accNum + 100); i++ {
		hex, bech := createRandomAccAddressHexAndBechNoBalance(int64(i + genInitNumber))
		addr, _ := TestAddr(hex, bech)
		addresses = append(addresses, addr)
	}

	return addresses
}

func createRandomAccAddressHexAndBechNoBalance(i int64) (hex string, bech string) {
	var buffer bytes.Buffer
	numString := strconv.Itoa(int(i))
	buffer.WriteString("A58856F0FD53BF058B4909A21AEC019107BA6") // base address string
	buffer.WriteString(numString)                               // adding on final two digits to make addresses unique
	res, _ := sdk.AccAddressFromHex(buffer.String())
	return buffer.String(), res.String()
}

func CreateRandomAccAddressNoBalance(i int64) string {
	_, bech := createRandomAccAddressHexAndBechNoBalance(i)
	return bech
}

func CreateRandomAccAddress() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	_, bech := createRandomAccAddressHexAndBechNoBalance(r.Int63())
	return bech
}

func ConvertAddrsToValAddrs(addrs []sdk.AccAddress) []sdk.ValAddress {
	valAddrs := make([]sdk.ValAddress, len(addrs))

	for i, addr := range addrs {
		valAddrs[i] = sdk.ValAddress(addr)
	}

	return valAddrs
}

func TestAddr(addr string, bech string) (sdk.AccAddress, error) {
	res, err := sdk.AccAddressFromHex(addr)
	if err != nil {
		return nil, err
	}
	bechexpected := res.String()
	if bech != bechexpected {
		return nil, fmt.Errorf("bech encoding doesn't match reference")
	}

	bechres, err := sdk.AccAddressFromBech32(bech)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(bechres, res) {
		return nil, err
	}

	return res, nil
}
